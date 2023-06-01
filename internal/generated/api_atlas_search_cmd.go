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

type CreateAtlasSearchIndexOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	clusterName string
	
}

func (opts *CreateAtlasSearchIndexOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *CreateAtlasSearchIndexOpts) Run(ctx context.Context) error {
	params := &admin.CreateAtlasSearchIndexApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
		
	}
	resp, _, err := opts.client.AtlasSearchApi.CreateAtlasSearchIndexWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func CreateAtlasSearchIndexBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := CreateAtlasSearchIndexOpts{}
	cmd := &cobra.Command{
		Use:     "createAtlasSearchIndex",
		// Aliases: []string{"?"},
		Short:   "Create One Atlas Search Index",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		Args:    require.NoArgs,
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
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", "Name of the cluster that contains the collection on which to create an Atlas Search index.")
	

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}
type DeleteAtlasSearchIndexOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	clusterName string
	indexId string
}

func (opts *DeleteAtlasSearchIndexOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *DeleteAtlasSearchIndexOpts) Run(ctx context.Context) error {
	params := &admin.DeleteAtlasSearchIndexApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
		IndexId: opts.indexId,
	}
	resp, _, err := opts.client.AtlasSearchApi.DeleteAtlasSearchIndexWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func DeleteAtlasSearchIndexBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := DeleteAtlasSearchIndexOpts{}
	cmd := &cobra.Command{
		Use:     "deleteAtlasSearchIndex",
		// Aliases: []string{"?"},
		Short:   "Remove One Atlas Search Index",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		Args:    require.NoArgs,
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
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", "Name of the cluster that contains the database and collection with one or more Application Search indexes.")
	cmd.Flags().StringVar(&opts.indexId, "indexId", "", "Unique 24-hexadecimal digit string that identifies the Atlas Search index. Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	_ = cmd.MarkFlagRequired("indexId")
	return cmd
}
type GetAtlasSearchIndexOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	clusterName string
	indexId string
}

func (opts *GetAtlasSearchIndexOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetAtlasSearchIndexOpts) Run(ctx context.Context) error {
	params := &admin.GetAtlasSearchIndexApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
		IndexId: opts.indexId,
	}
	resp, _, err := opts.client.AtlasSearchApi.GetAtlasSearchIndexWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func GetAtlasSearchIndexBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := GetAtlasSearchIndexOpts{}
	cmd := &cobra.Command{
		Use:     "getAtlasSearchIndex",
		// Aliases: []string{"?"},
		Short:   "Return One Atlas Search Index",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		Args:    require.NoArgs,
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
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", "Name of the cluster that contains the collection with one or more Atlas Search indexes.")
	cmd.Flags().StringVar(&opts.indexId, "indexId", "", "Unique 24-hexadecimal digit string that identifies the Application Search [index](https://docs.atlas.mongodb.com/reference/atlas-search/index-definitions/). Use the [Get All Application Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Application Search indexes.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	_ = cmd.MarkFlagRequired("indexId")
	return cmd
}
type ListAtlasSearchIndexesOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	clusterName string
	collectionName string
	databaseName string
}

func (opts *ListAtlasSearchIndexesOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListAtlasSearchIndexesOpts) Run(ctx context.Context) error {
	params := &admin.ListAtlasSearchIndexesApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
		CollectionName: opts.collectionName,
		DatabaseName: opts.databaseName,
	}
	resp, _, err := opts.client.AtlasSearchApi.ListAtlasSearchIndexesWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func ListAtlasSearchIndexesBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := ListAtlasSearchIndexesOpts{}
	cmd := &cobra.Command{
		Use:     "listAtlasSearchIndexes",
		// Aliases: []string{"?"},
		Short:   "Return All Atlas Search Indexes for One Collection",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		Args:    require.NoArgs,
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
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", "Name of the cluster that contains the collection with one or more Atlas Search indexes.")
	cmd.Flags().StringVar(&opts.collectionName, "collectionName", "", "Name of the collection that contains one or more Atlas Search indexes.")
	cmd.Flags().StringVar(&opts.databaseName, "databaseName", "", "Human-readable label that identifies the database that contains the collection with one or more Atlas Search indexes.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	_ = cmd.MarkFlagRequired("collectionName")
	_ = cmd.MarkFlagRequired("databaseName")
	return cmd
}
type UpdateAtlasSearchIndexOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	clusterName string
	indexId string
	
}

func (opts *UpdateAtlasSearchIndexOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *UpdateAtlasSearchIndexOpts) Run(ctx context.Context) error {
	params := &admin.UpdateAtlasSearchIndexApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
		IndexId: opts.indexId,
		
	}
	resp, _, err := opts.client.AtlasSearchApi.UpdateAtlasSearchIndexWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func UpdateAtlasSearchIndexBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := UpdateAtlasSearchIndexOpts{}
	cmd := &cobra.Command{
		Use:     "updateAtlasSearchIndex",
		// Aliases: []string{"?"},
		Short:   "Update One Atlas Search Index",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		Args:    require.NoArgs,
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
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", "Name of the cluster that contains the collection whose Atlas Search index to update.")
	cmd.Flags().StringVar(&opts.indexId, "indexId", "", "Unique 24-hexadecimal digit string that identifies the Atlas Search [index](https://docs.atlas.mongodb.com/reference/atlas-search/index-definitions/). Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes.")
	

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	_ = cmd.MarkFlagRequired("indexId")
	return cmd
}

func AtlasSearchBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "atlasSearch",
		Short:   "Returns, adds, edits, and removes Atlas Search indexes for the specified cluster. Also returns and updates user-defined analyzers for the specified cluster.",
	}
	cmd.AddCommand(
		CreateAtlasSearchIndexBuilder(),
		DeleteAtlasSearchIndexBuilder(),
		GetAtlasSearchIndexBuilder(),
		ListAtlasSearchIndexesBuilder(),
		UpdateAtlasSearchIndexBuilder(),
	)
	return cmd
}

