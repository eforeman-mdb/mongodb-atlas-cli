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

package store

import (
	"context"
	atlasv2 "go.mongodb.org/atlas/mongodbatlasv2"
)

//go:generate mockgen -destination=../../mocks/atlas/api_aws_clusters_dns_store_mock.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas GetAWSCustomDNSOperation, ToggleAWSCustomDNSOperation

type GetAWSCustomDNSOperation interface {
	GetAWSCustomDNS (*atlasv2.GetAWSCustomDNSApiParams) (*atlasv2.AWSCustomDNSEnabled, error)
}

// GetAWSCustomDNS encapsulates the logic to manage different cloud providers.
func (s *Store) GetAWSCustomDNS(params *atlasv2.GetAWSCustomDNSApiParams) (*atlasv2.AWSCustomDNSEnabled, error) {
	result, _, err := s.clientv2.AWSClustersDNSApi.GetAWSCustomDNSWithParams(s.ctx, params).Execute()
	return &result, err
}

type ToggleAWSCustomDNSOperation interface {
	ToggleAWSCustomDNS (*atlasv2.ToggleAWSCustomDNSApiParams) (*atlasv2.AWSCustomDNSEnabled, error)
}

// ToggleAWSCustomDNS encapsulates the logic to manage different cloud providers.
func (s *Store) ToggleAWSCustomDNS(params *atlasv2.ToggleAWSCustomDNSApiParams) (*atlasv2.AWSCustomDNSEnabled, error) {
	result, _, err := s.clientv2.AWSClustersDNSApi.ToggleAWSCustomDNSWithParams(s.ctx, params).Execute()
	return &result, err
}

