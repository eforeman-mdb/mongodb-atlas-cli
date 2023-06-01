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

type VersionedExampleOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	additionalInfo bool
}

func (opts *VersionedExampleOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *VersionedExampleOpts) Run(ctx context.Context) error {
	params := &admin.VersionedExampleApiParams{
		AdditionalInfo: opts.additionalInfo,
	}
	resp, _, err := opts.client.TestApi.VersionedExampleWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func VersionedExampleBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := VersionedExampleOpts{}
	cmd := &cobra.Command{
		Use:     "versionedExample",
		// Aliases: []string{"?"},
		Short:   "Example resource info for versioning of the Atlas API",
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
	cmd.Flags().BoolVar(&opts.additionalInfo, "additionalInfo", false, "")

	return cmd
}

func TestBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "test",
		Short:   "Atlas test endpoints.",
	}
	cmd.AddCommand(
		VersionedExampleBuilder(),
	)
	return cmd
}

