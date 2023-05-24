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

//go:generate mockgen -destination=../../mocks/atlas/api_multi_cloud_clusters_store_mock.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas CreateClusterOperation, DeleteClusterOperation, GetClusterOperation, ListClustersOperation, TestFailoverOperation, UpdateClusterOperation

type CreateClusterOperation interface {
	CreateCluster (*atlasv2.CreateClusterApiParams) (*atlasv2.ClusterDescriptionV15, error)
}

// CreateCluster encapsulates the logic to manage different cloud providers.
func (s *Store) CreateCluster(params *atlasv2.CreateClusterApiParams) (*atlasv2.ClusterDescriptionV15, error) {
	result, _, err := s.clientv2.MultiCloudClustersApi.CreateClusterWithParams(s.ctx, params).Execute()
	return &result, err
}

type DeleteClusterOperation interface {
	DeleteCluster (*atlasv2.DeleteClusterApiParams) (, error)
}

// DeleteCluster encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteCluster(params *atlasv2.DeleteClusterApiParams) (error) {
	_, err := s.clientv2.MultiCloudClustersApi.DeleteClusterWithParams(s.ctx, params).Execute()
	return err
}

type GetClusterOperation interface {
	GetCluster (*atlasv2.GetClusterApiParams) (*atlasv2.ClusterDescriptionV15, error)
}

// GetCluster encapsulates the logic to manage different cloud providers.
func (s *Store) GetCluster(params *atlasv2.GetClusterApiParams) (*atlasv2.ClusterDescriptionV15, error) {
	result, _, err := s.clientv2.MultiCloudClustersApi.GetClusterWithParams(s.ctx, params).Execute()
	return &result, err
}

type ListClustersOperation interface {
	ListClusters (*atlasv2.ListClustersApiParams) (*atlasv2.PaginatedClusterDescriptionV15, error)
}

// ListClusters encapsulates the logic to manage different cloud providers.
func (s *Store) ListClusters(params *atlasv2.ListClustersApiParams) (*atlasv2.PaginatedClusterDescriptionV15, error) {
	result, _, err := s.clientv2.MultiCloudClustersApi.ListClustersWithParams(s.ctx, params).Execute()
	return &result, err
}

type TestFailoverOperation interface {
	TestFailover (*atlasv2.TestFailoverApiParams) (, error)
}

// TestFailover encapsulates the logic to manage different cloud providers.
func (s *Store) TestFailover(params *atlasv2.TestFailoverApiParams) (error) {
	_, err := s.clientv2.MultiCloudClustersApi.TestFailoverWithParams(s.ctx, params).Execute()
	return err
}

type UpdateClusterOperation interface {
	UpdateCluster (*atlasv2.UpdateClusterApiParams) (*atlasv2.ClusterDescriptionV15, error)
}

// UpdateCluster encapsulates the logic to manage different cloud providers.
func (s *Store) UpdateCluster(params *atlasv2.UpdateClusterApiParams) (*atlasv2.ClusterDescriptionV15, error) {
	result, _, err := s.clientv2.MultiCloudClustersApi.UpdateClusterWithParams(s.ctx, params).Execute()
	return &result, err
}

