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

//go:generate mockgen -destination=../../mocks/atlas/api_cloud_backups_store_mock.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas CancelBackupRestoreJobOperation, CreateBackupExportJobOperation, CreateBackupRestoreJobOperation, CreateExportBucketOperation, CreateServerlessBackupRestoreJobOperation, DeleteAllBackupSchedulesOperation, DeleteExportBucketOperation, DeleteReplicaSetBackupOperation, DeleteShardedClusterBackupOperation, GetBackupExportJobOperation, GetBackupRestoreJobOperation, GetBackupScheduleOperation, GetDataProtectionSettingsOperation, GetExportBucketOperation, GetReplicaSetBackupOperation, GetServerlessBackupOperation, GetServerlessBackupRestoreJobOperation, GetShardedClusterBackupOperation, ListBackupExportJobsOperation, ListBackupRestoreJobsOperation, ListExportBucketsOperation, ListReplicaSetBackupsOperation, ListServerlessBackupRestoreJobsOperation, ListServerlessBackupsOperation, ListShardedClusterBackupsOperation, TakeSnapshotOperation, UpdateBackupScheduleOperation, UpdateDataProtectionSettingsOperation, UpdateSnapshotRetentionOperation

type CancelBackupRestoreJobOperation interface {
	CancelBackupRestoreJob (*atlasv2.CancelBackupRestoreJobApiParams) (*atlasv2.map[string]interface{}, error)
}

// CancelBackupRestoreJob encapsulates the logic to manage different cloud providers.
func (s *Store) CancelBackupRestoreJob(params *atlasv2.CancelBackupRestoreJobApiParams) (*atlasv2.map[string]interface{}, error) {
	result, _, err := s.clientv2.CloudBackupsApi.CancelBackupRestoreJobWithParams(s.ctx, params).Execute()
	return &result, err
}

type CreateBackupExportJobOperation interface {
	CreateBackupExportJob (*atlasv2.CreateBackupExportJobApiParams) (*atlasv2.DiskBackupExportJob, error)
}

// CreateBackupExportJob encapsulates the logic to manage different cloud providers.
func (s *Store) CreateBackupExportJob(params *atlasv2.CreateBackupExportJobApiParams) (*atlasv2.DiskBackupExportJob, error) {
	result, _, err := s.clientv2.CloudBackupsApi.CreateBackupExportJobWithParams(s.ctx, params).Execute()
	return &result, err
}

type CreateBackupRestoreJobOperation interface {
	CreateBackupRestoreJob (*atlasv2.CreateBackupRestoreJobApiParams) (*atlasv2.DiskBackupRestoreJob, error)
}

// CreateBackupRestoreJob encapsulates the logic to manage different cloud providers.
func (s *Store) CreateBackupRestoreJob(params *atlasv2.CreateBackupRestoreJobApiParams) (*atlasv2.DiskBackupRestoreJob, error) {
	result, _, err := s.clientv2.CloudBackupsApi.CreateBackupRestoreJobWithParams(s.ctx, params).Execute()
	return &result, err
}

type CreateExportBucketOperation interface {
	CreateExportBucket (*atlasv2.CreateExportBucketApiParams) (*atlasv2.DiskBackupSnapshotAWSExportBucket, error)
}

// CreateExportBucket encapsulates the logic to manage different cloud providers.
func (s *Store) CreateExportBucket(params *atlasv2.CreateExportBucketApiParams) (*atlasv2.DiskBackupSnapshotAWSExportBucket, error) {
	result, _, err := s.clientv2.CloudBackupsApi.CreateExportBucketWithParams(s.ctx, params).Execute()
	return &result, err
}

type CreateServerlessBackupRestoreJobOperation interface {
	CreateServerlessBackupRestoreJob (*atlasv2.CreateServerlessBackupRestoreJobApiParams) (*atlasv2.ServerlessBackupRestoreJob, error)
}

// CreateServerlessBackupRestoreJob encapsulates the logic to manage different cloud providers.
func (s *Store) CreateServerlessBackupRestoreJob(params *atlasv2.CreateServerlessBackupRestoreJobApiParams) (*atlasv2.ServerlessBackupRestoreJob, error) {
	result, _, err := s.clientv2.CloudBackupsApi.CreateServerlessBackupRestoreJobWithParams(s.ctx, params).Execute()
	return &result, err
}

type DeleteAllBackupSchedulesOperation interface {
	DeleteAllBackupSchedules (*atlasv2.DeleteAllBackupSchedulesApiParams) (*atlasv2.DiskBackupSnapshotSchedule, error)
}

// DeleteAllBackupSchedules encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteAllBackupSchedules(params *atlasv2.DeleteAllBackupSchedulesApiParams) (*atlasv2.DiskBackupSnapshotSchedule, error) {
	result, _, err := s.clientv2.CloudBackupsApi.DeleteAllBackupSchedulesWithParams(s.ctx, params).Execute()
	return &result, err
}

type DeleteExportBucketOperation interface {
	DeleteExportBucket (*atlasv2.DeleteExportBucketApiParams) (*atlasv2.map[string]interface{}, error)
}

// DeleteExportBucket encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteExportBucket(params *atlasv2.DeleteExportBucketApiParams) (*atlasv2.map[string]interface{}, error) {
	result, _, err := s.clientv2.CloudBackupsApi.DeleteExportBucketWithParams(s.ctx, params).Execute()
	return &result, err
}

type DeleteReplicaSetBackupOperation interface {
	DeleteReplicaSetBackup (*atlasv2.DeleteReplicaSetBackupApiParams) (*atlasv2.map[string]interface{}, error)
}

// DeleteReplicaSetBackup encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteReplicaSetBackup(params *atlasv2.DeleteReplicaSetBackupApiParams) (*atlasv2.map[string]interface{}, error) {
	result, _, err := s.clientv2.CloudBackupsApi.DeleteReplicaSetBackupWithParams(s.ctx, params).Execute()
	return &result, err
}

type DeleteShardedClusterBackupOperation interface {
	DeleteShardedClusterBackup (*atlasv2.DeleteShardedClusterBackupApiParams) (*atlasv2.map[string]interface{}, error)
}

// DeleteShardedClusterBackup encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteShardedClusterBackup(params *atlasv2.DeleteShardedClusterBackupApiParams) (*atlasv2.map[string]interface{}, error) {
	result, _, err := s.clientv2.CloudBackupsApi.DeleteShardedClusterBackupWithParams(s.ctx, params).Execute()
	return &result, err
}

type GetBackupExportJobOperation interface {
	GetBackupExportJob (*atlasv2.GetBackupExportJobApiParams) (*atlasv2.DiskBackupExportJob, error)
}

// GetBackupExportJob encapsulates the logic to manage different cloud providers.
func (s *Store) GetBackupExportJob(params *atlasv2.GetBackupExportJobApiParams) (*atlasv2.DiskBackupExportJob, error) {
	result, _, err := s.clientv2.CloudBackupsApi.GetBackupExportJobWithParams(s.ctx, params).Execute()
	return &result, err
}

type GetBackupRestoreJobOperation interface {
	GetBackupRestoreJob (*atlasv2.GetBackupRestoreJobApiParams) (*atlasv2.DiskBackupRestoreJob, error)
}

// GetBackupRestoreJob encapsulates the logic to manage different cloud providers.
func (s *Store) GetBackupRestoreJob(params *atlasv2.GetBackupRestoreJobApiParams) (*atlasv2.DiskBackupRestoreJob, error) {
	result, _, err := s.clientv2.CloudBackupsApi.GetBackupRestoreJobWithParams(s.ctx, params).Execute()
	return &result, err
}

type GetBackupScheduleOperation interface {
	GetBackupSchedule (*atlasv2.GetBackupScheduleApiParams) (*atlasv2.DiskBackupSnapshotSchedule, error)
}

// GetBackupSchedule encapsulates the logic to manage different cloud providers.
func (s *Store) GetBackupSchedule(params *atlasv2.GetBackupScheduleApiParams) (*atlasv2.DiskBackupSnapshotSchedule, error) {
	result, _, err := s.clientv2.CloudBackupsApi.GetBackupScheduleWithParams(s.ctx, params).Execute()
	return &result, err
}

type GetDataProtectionSettingsOperation interface {
	GetDataProtectionSettings (*atlasv2.GetDataProtectionSettingsApiParams) (*atlasv2.DataProtectionSettings, error)
}

// GetDataProtectionSettings encapsulates the logic to manage different cloud providers.
func (s *Store) GetDataProtectionSettings(params *atlasv2.GetDataProtectionSettingsApiParams) (*atlasv2.DataProtectionSettings, error) {
	result, _, err := s.clientv2.CloudBackupsApi.GetDataProtectionSettingsWithParams(s.ctx, params).Execute()
	return &result, err
}

type GetExportBucketOperation interface {
	GetExportBucket (*atlasv2.GetExportBucketApiParams) (*atlasv2.DiskBackupSnapshotAWSExportBucket, error)
}

// GetExportBucket encapsulates the logic to manage different cloud providers.
func (s *Store) GetExportBucket(params *atlasv2.GetExportBucketApiParams) (*atlasv2.DiskBackupSnapshotAWSExportBucket, error) {
	result, _, err := s.clientv2.CloudBackupsApi.GetExportBucketWithParams(s.ctx, params).Execute()
	return &result, err
}

type GetReplicaSetBackupOperation interface {
	GetReplicaSetBackup (*atlasv2.GetReplicaSetBackupApiParams) (*atlasv2.DiskBackupReplicaSet, error)
}

// GetReplicaSetBackup encapsulates the logic to manage different cloud providers.
func (s *Store) GetReplicaSetBackup(params *atlasv2.GetReplicaSetBackupApiParams) (*atlasv2.DiskBackupReplicaSet, error) {
	result, _, err := s.clientv2.CloudBackupsApi.GetReplicaSetBackupWithParams(s.ctx, params).Execute()
	return &result, err
}

type GetServerlessBackupOperation interface {
	GetServerlessBackup (*atlasv2.GetServerlessBackupApiParams) (*atlasv2.ServerlessBackupSnapshot, error)
}

// GetServerlessBackup encapsulates the logic to manage different cloud providers.
func (s *Store) GetServerlessBackup(params *atlasv2.GetServerlessBackupApiParams) (*atlasv2.ServerlessBackupSnapshot, error) {
	result, _, err := s.clientv2.CloudBackupsApi.GetServerlessBackupWithParams(s.ctx, params).Execute()
	return &result, err
}

type GetServerlessBackupRestoreJobOperation interface {
	GetServerlessBackupRestoreJob (*atlasv2.GetServerlessBackupRestoreJobApiParams) (*atlasv2.ServerlessBackupRestoreJob, error)
}

// GetServerlessBackupRestoreJob encapsulates the logic to manage different cloud providers.
func (s *Store) GetServerlessBackupRestoreJob(params *atlasv2.GetServerlessBackupRestoreJobApiParams) (*atlasv2.ServerlessBackupRestoreJob, error) {
	result, _, err := s.clientv2.CloudBackupsApi.GetServerlessBackupRestoreJobWithParams(s.ctx, params).Execute()
	return &result, err
}

type GetShardedClusterBackupOperation interface {
	GetShardedClusterBackup (*atlasv2.GetShardedClusterBackupApiParams) (*atlasv2.DiskBackupShardedClusterSnapshot, error)
}

// GetShardedClusterBackup encapsulates the logic to manage different cloud providers.
func (s *Store) GetShardedClusterBackup(params *atlasv2.GetShardedClusterBackupApiParams) (*atlasv2.DiskBackupShardedClusterSnapshot, error) {
	result, _, err := s.clientv2.CloudBackupsApi.GetShardedClusterBackupWithParams(s.ctx, params).Execute()
	return &result, err
}

type ListBackupExportJobsOperation interface {
	ListBackupExportJobs (*atlasv2.ListBackupExportJobsApiParams) (*atlasv2.PaginatedApiAtlasDiskBackupExportJob, error)
}

// ListBackupExportJobs encapsulates the logic to manage different cloud providers.
func (s *Store) ListBackupExportJobs(params *atlasv2.ListBackupExportJobsApiParams) (*atlasv2.PaginatedApiAtlasDiskBackupExportJob, error) {
	result, _, err := s.clientv2.CloudBackupsApi.ListBackupExportJobsWithParams(s.ctx, params).Execute()
	return &result, err
}

type ListBackupRestoreJobsOperation interface {
	ListBackupRestoreJobs (*atlasv2.ListBackupRestoreJobsApiParams) (*atlasv2.PaginatedCloudBackupRestoreJob, error)
}

// ListBackupRestoreJobs encapsulates the logic to manage different cloud providers.
func (s *Store) ListBackupRestoreJobs(params *atlasv2.ListBackupRestoreJobsApiParams) (*atlasv2.PaginatedCloudBackupRestoreJob, error) {
	result, _, err := s.clientv2.CloudBackupsApi.ListBackupRestoreJobsWithParams(s.ctx, params).Execute()
	return &result, err
}

type ListExportBucketsOperation interface {
	ListExportBuckets (*atlasv2.ListExportBucketsApiParams) (*atlasv2.PaginatedBackupSnapshotExportBucket, error)
}

// ListExportBuckets encapsulates the logic to manage different cloud providers.
func (s *Store) ListExportBuckets(params *atlasv2.ListExportBucketsApiParams) (*atlasv2.PaginatedBackupSnapshotExportBucket, error) {
	result, _, err := s.clientv2.CloudBackupsApi.ListExportBucketsWithParams(s.ctx, params).Execute()
	return &result, err
}

type ListReplicaSetBackupsOperation interface {
	ListReplicaSetBackups (*atlasv2.ListReplicaSetBackupsApiParams) (*atlasv2.PaginatedCloudBackupReplicaSet, error)
}

// ListReplicaSetBackups encapsulates the logic to manage different cloud providers.
func (s *Store) ListReplicaSetBackups(params *atlasv2.ListReplicaSetBackupsApiParams) (*atlasv2.PaginatedCloudBackupReplicaSet, error) {
	result, _, err := s.clientv2.CloudBackupsApi.ListReplicaSetBackupsWithParams(s.ctx, params).Execute()
	return &result, err
}

type ListServerlessBackupRestoreJobsOperation interface {
	ListServerlessBackupRestoreJobs (*atlasv2.ListServerlessBackupRestoreJobsApiParams) (*atlasv2.PaginatedApiAtlasServerlessBackupRestoreJob, error)
}

// ListServerlessBackupRestoreJobs encapsulates the logic to manage different cloud providers.
func (s *Store) ListServerlessBackupRestoreJobs(params *atlasv2.ListServerlessBackupRestoreJobsApiParams) (*atlasv2.PaginatedApiAtlasServerlessBackupRestoreJob, error) {
	result, _, err := s.clientv2.CloudBackupsApi.ListServerlessBackupRestoreJobsWithParams(s.ctx, params).Execute()
	return &result, err
}

type ListServerlessBackupsOperation interface {
	ListServerlessBackups (*atlasv2.ListServerlessBackupsApiParams) (*atlasv2.PaginatedApiAtlasServerlessBackupSnapshot, error)
}

// ListServerlessBackups encapsulates the logic to manage different cloud providers.
func (s *Store) ListServerlessBackups(params *atlasv2.ListServerlessBackupsApiParams) (*atlasv2.PaginatedApiAtlasServerlessBackupSnapshot, error) {
	result, _, err := s.clientv2.CloudBackupsApi.ListServerlessBackupsWithParams(s.ctx, params).Execute()
	return &result, err
}

type ListShardedClusterBackupsOperation interface {
	ListShardedClusterBackups (*atlasv2.ListShardedClusterBackupsApiParams) (*atlasv2.PaginatedCloudBackupShardedClusterSnapshot, error)
}

// ListShardedClusterBackups encapsulates the logic to manage different cloud providers.
func (s *Store) ListShardedClusterBackups(params *atlasv2.ListShardedClusterBackupsApiParams) (*atlasv2.PaginatedCloudBackupShardedClusterSnapshot, error) {
	result, _, err := s.clientv2.CloudBackupsApi.ListShardedClusterBackupsWithParams(s.ctx, params).Execute()
	return &result, err
}

type TakeSnapshotOperation interface {
	TakeSnapshot (*atlasv2.TakeSnapshotApiParams) (*atlasv2.DiskBackupSnapshot, error)
}

// TakeSnapshot encapsulates the logic to manage different cloud providers.
func (s *Store) TakeSnapshot(params *atlasv2.TakeSnapshotApiParams) (*atlasv2.DiskBackupSnapshot, error) {
	result, _, err := s.clientv2.CloudBackupsApi.TakeSnapshotWithParams(s.ctx, params).Execute()
	return &result, err
}

type UpdateBackupScheduleOperation interface {
	UpdateBackupSchedule (*atlasv2.UpdateBackupScheduleApiParams) (*atlasv2.DiskBackupSnapshotSchedule, error)
}

// UpdateBackupSchedule encapsulates the logic to manage different cloud providers.
func (s *Store) UpdateBackupSchedule(params *atlasv2.UpdateBackupScheduleApiParams) (*atlasv2.DiskBackupSnapshotSchedule, error) {
	result, _, err := s.clientv2.CloudBackupsApi.UpdateBackupScheduleWithParams(s.ctx, params).Execute()
	return &result, err
}

type UpdateDataProtectionSettingsOperation interface {
	UpdateDataProtectionSettings (*atlasv2.UpdateDataProtectionSettingsApiParams) (*atlasv2.DataProtectionSettings, error)
}

// UpdateDataProtectionSettings encapsulates the logic to manage different cloud providers.
func (s *Store) UpdateDataProtectionSettings(params *atlasv2.UpdateDataProtectionSettingsApiParams) (*atlasv2.DataProtectionSettings, error) {
	result, _, err := s.clientv2.CloudBackupsApi.UpdateDataProtectionSettingsWithParams(s.ctx, params).Execute()
	return &result, err
}

type UpdateSnapshotRetentionOperation interface {
	UpdateSnapshotRetention (*atlasv2.UpdateSnapshotRetentionApiParams) (*atlasv2.DiskBackupReplicaSet, error)
}

// UpdateSnapshotRetention encapsulates the logic to manage different cloud providers.
func (s *Store) UpdateSnapshotRetention(params *atlasv2.UpdateSnapshotRetentionApiParams) (*atlasv2.DiskBackupReplicaSet, error) {
	result, _, err := s.clientv2.CloudBackupsApi.UpdateSnapshotRetentionWithParams(s.ctx, params).Execute()
	return &result, err
}

