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
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	store "github.com/mongodb/mongodb-atlas-cli/internal/store/atlas"
)

type CreateAtlasSearchIndexOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store store.CreateAtlasSearchIndexOperation
	groupId string
	clusterName string
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *CreateAtlasSearchIndexOpts) Run() error {
	params := &atlasv2.CreateAtlasSearchIndexApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
	}
	resp, _, err := opts.store.CreateAtlasSearchIndex(params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const CreateAtlasSearchIndexTemplate = "<<some template>>"

func CreateAtlasSearchIndexBuilder() cobra.Command {
	opts := CreateAtlasSearchIndexOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		Short:   "<<decription>>",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		// Aliases: []string{"ls"},
		Args:    require.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), CreateAtlasSearchIndexTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", "usage description")

	return cmd
}
type DeleteAtlasSearchIndexOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store store.DeleteAtlasSearchIndexOperation
	groupId string
	clusterName string
	indexId string
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeleteAtlasSearchIndexOpts) Run() error {
	params := &atlasv2.DeleteAtlasSearchIndexApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
		IndexId: opts.indexId,
	}
	resp, _, err := opts.store.DeleteAtlasSearchIndex(params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const DeleteAtlasSearchIndexTemplate = "<<some template>>"

func DeleteAtlasSearchIndexBuilder() cobra.Command {
	opts := DeleteAtlasSearchIndexOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		Short:   "<<decription>>",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		// Aliases: []string{"ls"},
		Args:    require.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), DeleteAtlasSearchIndexTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", "usage description")
	cmd.Flags().StringVar(&opts.indexId, "indexId", "", "usage description")

	return cmd
}
type GetAtlasSearchIndexOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store store.GetAtlasSearchIndexOperation
	groupId string
	clusterName string
	indexId string
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *GetAtlasSearchIndexOpts) Run() error {
	params := &atlasv2.GetAtlasSearchIndexApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
		IndexId: opts.indexId,
	}
	resp, _, err := opts.store.GetAtlasSearchIndex(params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const GetAtlasSearchIndexTemplate = "<<some template>>"

func GetAtlasSearchIndexBuilder() cobra.Command {
	opts := GetAtlasSearchIndexOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		Short:   "<<decription>>",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		// Aliases: []string{"ls"},
		Args:    require.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), GetAtlasSearchIndexTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", "usage description")
	cmd.Flags().StringVar(&opts.indexId, "indexId", "", "usage description")

	return cmd
}
type ListAtlasSearchIndexesOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store store.ListAtlasSearchIndexesOperation
	groupId string
	clusterName string
	collectionName string
	databaseName string
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *ListAtlasSearchIndexesOpts) Run() error {
	params := &atlasv2.ListAtlasSearchIndexesApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
		CollectionName: opts.collectionName,
		DatabaseName: opts.databaseName,
	}
	resp, _, err := opts.store.ListAtlasSearchIndexes(params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const ListAtlasSearchIndexesTemplate = "<<some template>>"

func ListAtlasSearchIndexesBuilder() cobra.Command {
	opts := ListAtlasSearchIndexesOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		Short:   "<<decription>>",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		// Aliases: []string{"ls"},
		Args:    require.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), ListAtlasSearchIndexesTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", "usage description")
	cmd.Flags().StringVar(&opts.collectionName, "collectionName", "", "usage description")
	cmd.Flags().StringVar(&opts.databaseName, "databaseName", "", "usage description")

	return cmd
}
type UpdateAtlasSearchIndexOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store store.UpdateAtlasSearchIndexOperation
	groupId string
	clusterName string
	indexId string
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *UpdateAtlasSearchIndexOpts) Run() error {
	params := &atlasv2.UpdateAtlasSearchIndexApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
		IndexId: opts.indexId,
	}
	resp, _, err := opts.store.UpdateAtlasSearchIndex(params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const UpdateAtlasSearchIndexTemplate = "<<some template>>"

func UpdateAtlasSearchIndexBuilder() cobra.Command {
	opts := UpdateAtlasSearchIndexOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		Short:   "<<decription>>",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		// Aliases: []string{"ls"},
		Args:    require.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), UpdateAtlasSearchIndexTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", "usage description")
	cmd.Flags().StringVar(&opts.indexId, "indexId", "", "usage description")

	return cmd
}
