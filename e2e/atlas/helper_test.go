// Copyright 2021 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//go:build e2e || atlas

package atlas_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"testing"

	"github.com/mongodb/mongocli/e2e"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas/mongodbatlas"
)

const (
	atlasEntity                  = "atlas"
	eventsEntity                 = "events"
	clustersEntity               = "clusters"
	processesEntity              = "processes"
	metricsEntity                = "metrics"
	searchEntity                 = "search"
	indexEntity                  = "index"
	datalakeEntity               = "datalake"
	alertsEntity                 = "alerts"
	configEntity                 = "settings"
	dbusersEntity                = "dbusers"
	certsEntity                  = "certs"
	privateEndpointsEntity       = "privateendpoints"
	onlineArchiveEntity          = "onlineArchives"
	iamEntity                    = "iam"
	projectEntity                = "project"
	orgEntity                    = "org"
	maintenanceEntity            = "maintenanceWindows"
	integrationsEntity           = "integrations"
	securityEntity               = "security"
	ldapEntity                   = "ldap"
	awsEntity                    = "aws"
	azureEntity                  = "azure"
	customDNSEntity              = "customDns"
	logsEntity                   = "logs"
	cloudProvidersEntity         = "cloudProviders"
	accessRolesEntity            = "accessRoles"
	customDBRoleEntity           = "customDbRoles"
	regionalModeEntity           = "regionalModes"
	serverlessEntity             = "serverless"
	liveMigrationsEntity         = "liveMigrations"
	accessLogsEntity             = "accessLogs"
	accessListEntity             = "accessList"
	performanceAdvisorEntity     = "performanceAdvisor"
	slowQueryLogsEntity          = "slowQueryLogs"
	namespacesEntity             = "namespaces"
	suggestedIndexesEntity       = "suggestedIndexes"
	slowoperationThresholdEntity = "slowOperationThreshold"
)

// Cluster settings.
const (
	e2eClusterTier     = "M30"
	e2eClusterProvider = "AWS" // e2eClusterProvider preferred provider for e2e testing.
	e2eMDBVer          = "4.4"
)

func getHostnameAndPort() (string, error) {
	processes, err := getProcesses()
	if err != nil {
		return "", err
	}

	// The first element may not be the created cluster but that is fine since
	// we just need one cluster up and running
	return processes[0].Hostname + ":" + strconv.Itoa(processes[0].Port), nil
}

func getHostname() (string, error) {
	processes, err := getProcesses()
	if err != nil {
		return "", err
	}

	return processes[0].Hostname, nil
}

func getProcesses() ([]*mongodbatlas.Process, error) {
	cliPath, err := e2e.Bin()
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(cliPath,
		atlasEntity,
		processesEntity,
		"list",
		"-o=json",
	)

	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()

	if err != nil {
		return nil, err
	}

	var processes []*mongodbatlas.Process
	err = json.Unmarshal(resp, &processes)

	if err != nil {
		return nil, err
	}

	if len(processes) == 0 {
		return nil, fmt.Errorf("got=%#v\nwant=%#v", 0, "len(processes) > 0")
	}

	return processes, nil
}

func deployClusterForProject(projectID string) (string, error) {
	cliPath, err := e2e.Bin()
	if err != nil {
		return "", err
	}
	clusterName, err := RandClusterName()
	if err != nil {
		return "", err
	}
	region, err := newAvailableRegion(projectID, e2eClusterTier, e2eClusterProvider)
	if err != nil {
		return "", err
	}
	args := []string{
		atlasEntity,
		clustersEntity,
		"create",
		clusterName,
		"--mdbVersion", e2eMDBVer,
		"--region", region,
		"--tier", e2eClusterTier,
		"--provider", e2eClusterProvider,
		"--diskSizeGB=30",
	}
	if projectID != "" {
		args = append(args, "--projectId", projectID)
	}
	create := exec.Command(cliPath, args...)
	create.Env = os.Environ()
	if resp, err := create.CombinedOutput(); err != nil {
		return "", fmt.Errorf("error creating cluster %w: %s", err, string(resp))
	}

	watchArgs := []string{
		atlasEntity,
		clustersEntity,
		"watch",
		clusterName,
	}
	if projectID != "" {
		watchArgs = append(watchArgs, "--projectId", projectID)
	}
	watch := exec.Command(cliPath, watchArgs...)
	watch.Env = os.Environ()
	if resp, err := watch.CombinedOutput(); err != nil {
		return "", fmt.Errorf("error watching cluster %w: %s", err, string(resp))
	}
	return clusterName, nil
}

func deployCluster() (string, error) {
	return deployClusterForProject("")
}

func deleteClusterForProject(projectID, clusterName string) error {
	cliPath, err := e2e.Bin()
	if err != nil {
		return err
	}
	args := []string{
		atlasEntity,
		clustersEntity,
		"delete",
		clusterName,
		"--force",
	}
	if projectID != "" {
		args = append(args, "--projectId", projectID)
	}
	deleteCmd := exec.Command(cliPath, args...)
	deleteCmd.Env = os.Environ()
	if resp, err := deleteCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("error deleting cluster %w: %s", err, string(resp))
	}

	watchArgs := []string{
		atlasEntity,
		clustersEntity,
		"watch",
		clusterName,
	}
	if projectID != "" {
		watchArgs = append(watchArgs, "--projectId", projectID)
	}
	watchCmd := exec.Command(cliPath, watchArgs...)
	watchCmd.Env = os.Environ()
	// this command will fail with 404 once the cluster is deleted
	// we just need to wait for this to close the project
	_ = watchCmd.Run()
	return nil
}

func deleteCluster(clusterName string) error {
	return deleteClusterForProject("", clusterName)
}

func newAvailableRegion(projectID, tier, provider string) (string, error) {
	cliPath, err := e2e.Bin()
	if err != nil {
		return "", err
	}
	args := []string{
		atlasEntity,
		clustersEntity,
		"availableRegions",
		"ls",
		"--provider", provider,
		"--tier", tier,
		"-o=json",
	}
	if projectID != "" {
		args = append(args, "--projectId", projectID)
	}
	cmd := exec.Command(cliPath, args...)
	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()

	if err != nil {
		return "", fmt.Errorf("error getting regions %w: %s", err, string(resp))
	}

	var cloudProviders mongodbatlas.CloudProviders
	err = json.Unmarshal(resp, &cloudProviders)
	if err != nil {
		return "", fmt.Errorf("error unmashaling response %w: %s", err, string(resp))
	}

	if len(cloudProviders.Results) == 0 || len(cloudProviders.Results[0].InstanceSizes) == 0 {
		return "", errors.New("no regions available")
	}

	return cloudProviders.Results[0].InstanceSizes[0].AvailableRegions[0].Name, nil
}

func RandClusterName() (string, error) {
	n, err := e2e.RandInt(1000)
	if err != nil {
		return "", err
	}
	if revision, ok := os.LookupEnv("revision"); ok {
		return fmt.Sprintf("cluster-%v-%s", n, revision), nil
	}
	return fmt.Sprintf("cluster-%v", n), nil
}

func RandProjectName() (string, error) {
	n, err := e2e.RandInt(1000)
	if err != nil {
		return "", err
	}
	if revision, ok := os.LookupEnv("revision"); ok {
		return fmt.Sprintf("%v-%s", n, revision), nil
	}
	return fmt.Sprintf("e2e-%v", n), nil
}

func RandProjectNameWithPrefix(prefix string) (string, error) {
	name, err := RandProjectName()
	if err != nil {
		return "", err
	}
	prefixedName := fmt.Sprintf("%s-%s", prefix, name)
	if len(prefixedName) > 64 {
		return prefixedName[:64], nil
	}
	return prefixedName, nil
}

func integrationExists(name string, thirdPartyIntegrations mongodbatlas.ThirdPartyIntegrations) bool {
	services := thirdPartyIntegrations.Results
	for i := range services {
		if services[i].Type == name {
			return true
		}
	}
	return false
}

func createProject(projectName string) (string, error) {
	cliPath, err := e2e.Bin()
	if err != nil {
		return "", fmt.Errorf("%w: invalid bin", err)
	}
	cmd := exec.Command(cliPath,
		iamEntity,
		projectEntity,
		"create",
		projectName,
		"-o=json")
	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s (%w)", string(resp), err)
	}

	var project mongodbatlas.Project
	if err := json.Unmarshal(resp, &project); err != nil {
		return "", fmt.Errorf("invalid response: %s (%w)", string(resp), err)
	}

	return project.ID, nil
}

func deleteProject(projectID string) error {
	cliPath, err := e2e.Bin()
	if err != nil {
		return err
	}
	cmd := exec.Command(cliPath,
		iamEntity,
		projectEntity,
		"delete",
		projectID,
		"--force")
	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s (%w)", string(resp), err)
	}
	return nil
}

func ensureCluster(t *testing.T, cluster *mongodbatlas.AdvancedCluster, clusterName, version string, diskSizeGB float64) {
	t.Helper()
	a := assert.New(t)
	a.Equal(clusterName, cluster.Name)
	a.Equal(version, cluster.MongoDBMajorVersion)
	a.Equal(diskSizeGB, *cluster.DiskSizeGB)
}
