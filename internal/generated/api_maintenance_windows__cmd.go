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

type DeferMaintenanceWindowOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store store.DeferMaintenanceWindowOperation
	groupId string
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeferMaintenanceWindowOpts) Run() error {
	params := &atlasv2.DeferMaintenanceWindowApiParams{
		GroupId: opts.groupId,
	}
	_, err := opts.store.DeferMaintenanceWindow(params)
	if err != nil {
		return err
	}

	return opts.Print(nil)
}

const DeferMaintenanceWindowTemplate = "<<some template>>"

func DeferMaintenanceWindowBuilder() cobra.Command {
	opts := DeferMaintenanceWindowOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), DeferMaintenanceWindowTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")

	return cmd
}
type GetMaintenanceWindowOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store store.GetMaintenanceWindowOperation
	groupId string
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *GetMaintenanceWindowOpts) Run() error {
	params := &atlasv2.GetMaintenanceWindowApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.store.GetMaintenanceWindow(params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const GetMaintenanceWindowTemplate = "<<some template>>"

func GetMaintenanceWindowBuilder() cobra.Command {
	opts := GetMaintenanceWindowOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), GetMaintenanceWindowTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")

	return cmd
}
type ResetMaintenanceWindowOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store store.ResetMaintenanceWindowOperation
	groupId string
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *ResetMaintenanceWindowOpts) Run() error {
	params := &atlasv2.ResetMaintenanceWindowApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.store.ResetMaintenanceWindow(params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const ResetMaintenanceWindowTemplate = "<<some template>>"

func ResetMaintenanceWindowBuilder() cobra.Command {
	opts := ResetMaintenanceWindowOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), ResetMaintenanceWindowTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")

	return cmd
}
type ToggleMaintenanceAutoDeferOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store store.ToggleMaintenanceAutoDeferOperation
	groupId string
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *ToggleMaintenanceAutoDeferOpts) Run() error {
	params := &atlasv2.ToggleMaintenanceAutoDeferApiParams{
		GroupId: opts.groupId,
	}
	_, err := opts.store.ToggleMaintenanceAutoDefer(params)
	if err != nil {
		return err
	}

	return opts.Print(nil)
}

const ToggleMaintenanceAutoDeferTemplate = "<<some template>>"

func ToggleMaintenanceAutoDeferBuilder() cobra.Command {
	opts := ToggleMaintenanceAutoDeferOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), ToggleMaintenanceAutoDeferTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")

	return cmd
}
type UpdateMaintenanceWindowOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store store.UpdateMaintenanceWindowOperation
	groupId string
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *UpdateMaintenanceWindowOpts) Run() error {
	params := &atlasv2.UpdateMaintenanceWindowApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.store.UpdateMaintenanceWindow(params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const UpdateMaintenanceWindowTemplate = "<<some template>>"

func UpdateMaintenanceWindowBuilder() cobra.Command {
	opts := UpdateMaintenanceWindowOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), UpdateMaintenanceWindowTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")

	return cmd
}
