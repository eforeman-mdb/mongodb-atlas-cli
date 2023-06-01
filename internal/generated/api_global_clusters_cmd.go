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

type CreateCustomZoneMappingOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	clusterName string
	
}

func (opts *CreateCustomZoneMappingOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *CreateCustomZoneMappingOpts) Run(ctx context.Context) error {
	params := &admin.CreateCustomZoneMappingApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
		
	}
	resp, _, err := opts.client.GlobalClustersApi.CreateCustomZoneMappingWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func CreateCustomZoneMappingBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := CreateCustomZoneMappingOpts{}
	cmd := &cobra.Command{
		Use:     "createCustomZoneMapping",
		// Aliases: []string{"?"},
		Short:   "Add One Entry to One Custom Zone Mapping",
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
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", "Human-readable label that identifies this advanced cluster.")
	

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}
type CreateManagedNamespaceOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	clusterName string
	
}

func (opts *CreateManagedNamespaceOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *CreateManagedNamespaceOpts) Run(ctx context.Context) error {
	params := &admin.CreateManagedNamespaceApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
		
	}
	resp, _, err := opts.client.GlobalClustersApi.CreateManagedNamespaceWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func CreateManagedNamespaceBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := CreateManagedNamespaceOpts{}
	cmd := &cobra.Command{
		Use:     "createManagedNamespace",
		// Aliases: []string{"?"},
		Short:   "Create One Managed Namespace in One Global Multi-Cloud Cluster",
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
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", "Human-readable label that identifies this advanced cluster.")
	

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}
type DeleteAllCustomZoneMappingsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	clusterName string
}

func (opts *DeleteAllCustomZoneMappingsOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *DeleteAllCustomZoneMappingsOpts) Run(ctx context.Context) error {
	params := &admin.DeleteAllCustomZoneMappingsApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
	}
	resp, _, err := opts.client.GlobalClustersApi.DeleteAllCustomZoneMappingsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func DeleteAllCustomZoneMappingsBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := DeleteAllCustomZoneMappingsOpts{}
	cmd := &cobra.Command{
		Use:     "deleteAllCustomZoneMappings",
		// Aliases: []string{"?"},
		Short:   "Remove All Custom Zone Mappings from One Global Multi-Cloud Cluster",
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
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", "Human-readable label that identifies this advanced cluster.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}
type DeleteManagedNamespaceOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	clusterName string
	groupId string
	db string
	collection string
}

func (opts *DeleteManagedNamespaceOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *DeleteManagedNamespaceOpts) Run(ctx context.Context) error {
	params := &admin.DeleteManagedNamespaceApiParams{
		ClusterName: opts.clusterName,
		GroupId: opts.groupId,
		Db: opts.db,
		Collection: opts.collection,
	}
	resp, _, err := opts.client.GlobalClustersApi.DeleteManagedNamespaceWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func DeleteManagedNamespaceBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := DeleteManagedNamespaceOpts{}
	cmd := &cobra.Command{
		Use:     "deleteManagedNamespace",
		// Aliases: []string{"?"},
		Short:   "Remove One Managed Namespace from One Global Multi-Cloud Cluster",
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
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", "Human-readable label that identifies this advanced cluster.")
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")
	cmd.Flags().StringVar(&opts.db, "db", "", "Human-readable label that identifies the database that contains the collection.")
	cmd.Flags().StringVar(&opts.collection, "collection", "", "Human-readable label that identifies the collection associated with the managed namespace.")

	_ = cmd.MarkFlagRequired("clusterName")
	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}
type GetManagedNamespaceOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	clusterName string
}

func (opts *GetManagedNamespaceOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetManagedNamespaceOpts) Run(ctx context.Context) error {
	params := &admin.GetManagedNamespaceApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
	}
	resp, _, err := opts.client.GlobalClustersApi.GetManagedNamespaceWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func GetManagedNamespaceBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := GetManagedNamespaceOpts{}
	cmd := &cobra.Command{
		Use:     "getManagedNamespace",
		// Aliases: []string{"?"},
		Short:   "Return One Managed Namespace in One Global Multi-Cloud Cluster",
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
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", "Human-readable label that identifies this advanced cluster.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}

func GlobalClustersBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "globalClusters",
		Short:   "Returns, adds, and removes Global Cluster managed namespaces and custom zone mappings. Each collection in a Global Cluster is associated with a managed namespace. When you create a managed namespace for a Global Cluster, MongoDB Cloud creates an empty collection for that namespace. Creating a managed namespace doesn&#39;t populate a collection with data. Similarly, deleting a managed namespace doesn&#39;t delete the associated collection.
MongoDB Cloud shards the empty collection using the required location field and a custom shard key. For example, if your custom shard key is &#x60;city&#x60;, the compound shard key is &#x60;location, city&#x60;. Each Global Cluster is also associated with one or more Global Writes Zones. When a user creates a Global Cluster, MongoDB Cloud automatically maps each location code to the closest geographical zone. Custom zone mappings allow administrators to override these automatic mappings. For example, a use case might require mapping a location code to a geographically distant zone. Administrators can manage custom zone mappings with the APIs below and the **Global Cluster Configuration** pane when you create or modify your Global Cluster.",
	}
	cmd.AddCommand(
		CreateCustomZoneMappingBuilder(),
		CreateManagedNamespaceBuilder(),
		DeleteAllCustomZoneMappingsBuilder(),
		DeleteManagedNamespaceBuilder(),
		GetManagedNamespaceBuilder(),
	)
	return cmd
}

