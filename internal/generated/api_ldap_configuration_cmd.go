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

type DeleteLDAPConfigurationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store store.DeleteLDAPConfigurationOperation
	groupId string
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeleteLDAPConfigurationOpts) Run() error {
	params := &atlasv2.DeleteLDAPConfigurationApiParams{
		GroupId: opts.groupId,
	}
	_, err := opts.store.DeleteLDAPConfiguration(params)
	if err != nil {
		return err
	}

	return opts.Print(nil)
}

const DeleteLDAPConfigurationTemplate = "<<some template>>"

func DeleteLDAPConfigurationBuilder() cobra.Command {
	opts := DeleteLDAPConfigurationOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), DeleteLDAPConfigurationTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")

	return cmd
}
type GetLDAPConfigurationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store store.GetLDAPConfigurationOperation
	groupId string
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *GetLDAPConfigurationOpts) Run() error {
	params := &atlasv2.GetLDAPConfigurationApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.store.GetLDAPConfiguration(params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const GetLDAPConfigurationTemplate = "<<some template>>"

func GetLDAPConfigurationBuilder() cobra.Command {
	opts := GetLDAPConfigurationOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), GetLDAPConfigurationTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")

	return cmd
}
type GetLDAPConfigurationStatusOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store store.GetLDAPConfigurationStatusOperation
	groupId string
	requestId string
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *GetLDAPConfigurationStatusOpts) Run() error {
	params := &atlasv2.GetLDAPConfigurationStatusApiParams{
		GroupId: opts.groupId,
		RequestId: opts.requestId,
	}
	resp, _, err := opts.store.GetLDAPConfigurationStatus(params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const GetLDAPConfigurationStatusTemplate = "<<some template>>"

func GetLDAPConfigurationStatusBuilder() cobra.Command {
	opts := GetLDAPConfigurationStatusOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), GetLDAPConfigurationStatusTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	cmd.Flags().StringVar(&opts.requestId, "requestId", "", "usage description")

	return cmd
}
type SaveLDAPConfigurationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store store.SaveLDAPConfigurationOperation
	groupId string
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *SaveLDAPConfigurationOpts) Run() error {
	params := &atlasv2.SaveLDAPConfigurationApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.store.SaveLDAPConfiguration(params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const SaveLDAPConfigurationTemplate = "<<some template>>"

func SaveLDAPConfigurationBuilder() cobra.Command {
	opts := SaveLDAPConfigurationOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), SaveLDAPConfigurationTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")

	return cmd
}
type VerifyLDAPConfigurationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store store.VerifyLDAPConfigurationOperation
	groupId string
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *VerifyLDAPConfigurationOpts) Run() error {
	params := &atlasv2.VerifyLDAPConfigurationApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.store.VerifyLDAPConfiguration(params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const VerifyLDAPConfigurationTemplate = "<<some template>>"

func VerifyLDAPConfigurationBuilder() cobra.Command {
	opts := VerifyLDAPConfigurationOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), VerifyLDAPConfigurationTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")

	return cmd
}
