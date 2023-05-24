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

//go:generate mockgen -destination=../../mocks/atlas/api_alert_configurations_store_mock.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas CreateAlertConfigurationOperation, DeleteAlertConfigurationOperation, GetAlertConfigurationOperation, ListAlertConfigurationMatchersFieldNamesOperation, ListAlertConfigurationsOperation, ListAlertConfigurationsByAlertIdOperation, ToggleAlertConfigurationOperation, UpdateAlertConfigurationOperation

type CreateAlertConfigurationOperation interface {
	CreateAlertConfiguration (*atlasv2.CreateAlertConfigurationApiParams) (*atlasv2.AlertConfigViewForNdsGroup, error)
}

// CreateAlertConfiguration encapsulates the logic to manage different cloud providers.
func (s *Store) CreateAlertConfiguration(params *atlasv2.CreateAlertConfigurationApiParams) (*atlasv2.AlertConfigViewForNdsGroup, error) {
	result, _, err := s.clientv2.AlertConfigurationsApi.CreateAlertConfigurationWithParams(s.ctx, params).Execute()
	return &result, err
}

type DeleteAlertConfigurationOperation interface {
	DeleteAlertConfiguration (*atlasv2.DeleteAlertConfigurationApiParams) (, error)
}

// DeleteAlertConfiguration encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteAlertConfiguration(params *atlasv2.DeleteAlertConfigurationApiParams) (error) {
	_, err := s.clientv2.AlertConfigurationsApi.DeleteAlertConfigurationWithParams(s.ctx, params).Execute()
	return err
}

type GetAlertConfigurationOperation interface {
	GetAlertConfiguration (*atlasv2.GetAlertConfigurationApiParams) (*atlasv2.AlertConfigViewForNdsGroup, error)
}

// GetAlertConfiguration encapsulates the logic to manage different cloud providers.
func (s *Store) GetAlertConfiguration(params *atlasv2.GetAlertConfigurationApiParams) (*atlasv2.AlertConfigViewForNdsGroup, error) {
	result, _, err := s.clientv2.AlertConfigurationsApi.GetAlertConfigurationWithParams(s.ctx, params).Execute()
	return &result, err
}

type ListAlertConfigurationMatchersFieldNamesOperation interface {
	ListAlertConfigurationMatchersFieldNames (*atlasv2.ListAlertConfigurationMatchersFieldNamesApiParams) (*atlasv2.[]MatcherField, error)
}

// ListAlertConfigurationMatchersFieldNames encapsulates the logic to manage different cloud providers.
func (s *Store) ListAlertConfigurationMatchersFieldNames(params *atlasv2.ListAlertConfigurationMatchersFieldNamesApiParams) (*atlasv2.[]MatcherField, error) {
	result, _, err := s.clientv2.AlertConfigurationsApi.ListAlertConfigurationMatchersFieldNamesWithParams(s.ctx, params).Execute()
	return &result, err
}

type ListAlertConfigurationsOperation interface {
	ListAlertConfigurations (*atlasv2.ListAlertConfigurationsApiParams) (*atlasv2.PaginatedAlertConfig, error)
}

// ListAlertConfigurations encapsulates the logic to manage different cloud providers.
func (s *Store) ListAlertConfigurations(params *atlasv2.ListAlertConfigurationsApiParams) (*atlasv2.PaginatedAlertConfig, error) {
	result, _, err := s.clientv2.AlertConfigurationsApi.ListAlertConfigurationsWithParams(s.ctx, params).Execute()
	return &result, err
}

type ListAlertConfigurationsByAlertIdOperation interface {
	ListAlertConfigurationsByAlertId (*atlasv2.ListAlertConfigurationsByAlertIdApiParams) (*atlasv2.PaginatedAlertConfig, error)
}

// ListAlertConfigurationsByAlertId encapsulates the logic to manage different cloud providers.
func (s *Store) ListAlertConfigurationsByAlertId(params *atlasv2.ListAlertConfigurationsByAlertIdApiParams) (*atlasv2.PaginatedAlertConfig, error) {
	result, _, err := s.clientv2.AlertConfigurationsApi.ListAlertConfigurationsByAlertIdWithParams(s.ctx, params).Execute()
	return &result, err
}

type ToggleAlertConfigurationOperation interface {
	ToggleAlertConfiguration (*atlasv2.ToggleAlertConfigurationApiParams) (*atlasv2.AlertConfigViewForNdsGroup, error)
}

// ToggleAlertConfiguration encapsulates the logic to manage different cloud providers.
func (s *Store) ToggleAlertConfiguration(params *atlasv2.ToggleAlertConfigurationApiParams) (*atlasv2.AlertConfigViewForNdsGroup, error) {
	result, _, err := s.clientv2.AlertConfigurationsApi.ToggleAlertConfigurationWithParams(s.ctx, params).Execute()
	return &result, err
}

type UpdateAlertConfigurationOperation interface {
	UpdateAlertConfiguration (*atlasv2.UpdateAlertConfigurationApiParams) (*atlasv2.AlertConfigViewForNdsGroup, error)
}

// UpdateAlertConfiguration encapsulates the logic to manage different cloud providers.
func (s *Store) UpdateAlertConfiguration(params *atlasv2.UpdateAlertConfigurationApiParams) (*atlasv2.AlertConfigViewForNdsGroup, error) {
	result, _, err := s.clientv2.AlertConfigurationsApi.UpdateAlertConfigurationWithParams(s.ctx, params).Execute()
	return &result, err
}

