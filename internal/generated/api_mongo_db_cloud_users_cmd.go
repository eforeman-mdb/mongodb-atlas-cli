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

type CreateUserOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	
}

func (opts *CreateUserOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *CreateUserOpts) Run(ctx context.Context) error {
	params := &admin.CreateUserApiParams{
		
	}
	resp, _, err := opts.client.MongoDBCloudUsersApi.CreateUserWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func CreateUserBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := CreateUserOpts{}
	cmd := &cobra.Command{
		Use:     "createUser",
		// Aliases: []string{"?"},
		Short:   "Create One MongoDB Cloud User",
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
	

	return cmd
}
type GetUserOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	userId string
}

func (opts *GetUserOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetUserOpts) Run(ctx context.Context) error {
	params := &admin.GetUserApiParams{
		UserId: opts.userId,
	}
	resp, _, err := opts.client.MongoDBCloudUsersApi.GetUserWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func GetUserBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := GetUserOpts{}
	cmd := &cobra.Command{
		Use:     "getUser",
		// Aliases: []string{"?"},
		Short:   "Return One MongoDB Cloud User using Its ID",
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
	cmd.Flags().StringVar(&opts.userId, "userId", "", "Unique 24-hexadecimal digit string that identifies this user.")

	_ = cmd.MarkFlagRequired("userId")
	return cmd
}
type GetUserByUsernameOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	userName string
}

func (opts *GetUserByUsernameOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetUserByUsernameOpts) Run(ctx context.Context) error {
	params := &admin.GetUserByUsernameApiParams{
		UserName: opts.userName,
	}
	resp, _, err := opts.client.MongoDBCloudUsersApi.GetUserByUsernameWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func GetUserByUsernameBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := GetUserByUsernameOpts{}
	cmd := &cobra.Command{
		Use:     "getUserByUsername",
		// Aliases: []string{"?"},
		Short:   "Return One MongoDB Cloud User using Their Username",
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
	cmd.Flags().StringVar(&opts.userName, "userName", "", "Email address that belongs to the MongoDB Cloud user account. You cannot modify this address after creating the user.")

	_ = cmd.MarkFlagRequired("userName")
	return cmd
}

func MongoDBCloudUsersBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "mongoDBCloudUsers",
		Short:   "Returns, adds, and edits MongoDB Cloud users.",
	}
	cmd.AddCommand(
		CreateUserBuilder(),
		GetUserBuilder(),
		GetUserByUsernameBuilder(),
	)
	return cmd
}

