// Copyright 2023 MongoDB Inc
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

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package generated

import (
	"context"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/admin"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
)

type EndOutageSimulationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	clusterName string
}

func (opts *EndOutageSimulationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *EndOutageSimulationOpts) Run(ctx context.Context) error {
	params := &admin.EndOutageSimulationApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
	}
	resp, _, err := opts.client.ClusterOutageSimulationApi.EndOutageSimulationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func EndOutageSimulationBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := EndOutageSimulationOpts{}
	cmd := &cobra.Command{
		Use:     "endOutageSimulation",
		// Aliases: []string{"?"},
		Short:   "End an Outage Simulation",
		Annotations: map[string]string{
			"output":      template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				//opts.ValidateProjectID,
				opts.initClient(),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", "Human-readable label that identifies the cluster that is undergoing outage simulation.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}
type GetOutageSimulationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	clusterName string
}

func (opts *GetOutageSimulationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetOutageSimulationOpts) Run(ctx context.Context) error {
	params := &admin.GetOutageSimulationApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
	}
	resp, _, err := opts.client.ClusterOutageSimulationApi.GetOutageSimulationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func GetOutageSimulationBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := GetOutageSimulationOpts{}
	cmd := &cobra.Command{
		Use:     "getOutageSimulation",
		// Aliases: []string{"?"},
		Short:   "Return One Outage Simulation",
		Annotations: map[string]string{
			"output":      template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				//opts.ValidateProjectID,
				opts.initClient(),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", "Human-readable label that identifies the cluster that is undergoing outage simulation.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}
type StartOutageSimulationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	clusterName string
	
}

func (opts *StartOutageSimulationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *StartOutageSimulationOpts) Run(ctx context.Context) error {
	params := &admin.StartOutageSimulationApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
		
	}
	resp, _, err := opts.client.ClusterOutageSimulationApi.StartOutageSimulationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func StartOutageSimulationBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := StartOutageSimulationOpts{}
	cmd := &cobra.Command{
		Use:     "startOutageSimulation",
		// Aliases: []string{"?"},
		Short:   "Start an Outage Simulation",
		Annotations: map[string]string{
			"output":      template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				//opts.ValidateProjectID,
				opts.initClient(),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", "Human-readable label that identifies the cluster to undergo an outage simulation.")
	

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}

func ClusterOutageSimulationBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "clusterOutageSimulation",
		Short:   "Returns, starts, or ends a cluster outage simulation.",
	}
	cmd.AddCommand(
		EndOutageSimulationBuilder(),
		GetOutageSimulationBuilder(),
		StartOutageSimulationBuilder(),
	)
	return cmd
}

