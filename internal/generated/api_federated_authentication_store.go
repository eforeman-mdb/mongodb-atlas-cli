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

//go:generate mockgen -destination=../../mocks/atlas/api_federated_authentication_store_mock.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas CreateRoleMappingOperation, DeleteFederationAppOperation, DeleteRoleMappingOperation, GetConnectedOrgConfigOperation, GetFederationSettingsOperation, GetIdentityProviderOperation, GetIdentityProviderMetadataOperation, GetRoleMappingOperation, ListConnectedOrgConfigsOperation, ListIdentityProvidersOperation, ListRoleMappingsOperation, RemoveConnectedOrgConfigOperation, UpdateConnectedOrgConfigOperation, UpdateIdentityProviderOperation, UpdateRoleMappingOperation

type CreateRoleMappingOperation interface {
	CreateRoleMapping (*atlasv2.CreateRoleMappingApiParams) (*atlasv2.RoleMapping, error)
}

// CreateRoleMapping encapsulates the logic to manage different cloud providers.
func (s *Store) CreateRoleMapping(params *atlasv2.CreateRoleMappingApiParams) (*atlasv2.RoleMapping, error) {
	result, _, err := s.clientv2.FederatedAuthenticationApi.CreateRoleMappingWithParams(s.ctx, params).Execute()
	return &result, err
}

type DeleteFederationAppOperation interface {
	DeleteFederationApp (*atlasv2.DeleteFederationAppApiParams) (, error)
}

// DeleteFederationApp encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteFederationApp(params *atlasv2.DeleteFederationAppApiParams) (error) {
	_, err := s.clientv2.FederatedAuthenticationApi.DeleteFederationAppWithParams(s.ctx, params).Execute()
	return err
}

type DeleteRoleMappingOperation interface {
	DeleteRoleMapping (*atlasv2.DeleteRoleMappingApiParams) (, error)
}

// DeleteRoleMapping encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteRoleMapping(params *atlasv2.DeleteRoleMappingApiParams) (error) {
	_, err := s.clientv2.FederatedAuthenticationApi.DeleteRoleMappingWithParams(s.ctx, params).Execute()
	return err
}

type GetConnectedOrgConfigOperation interface {
	GetConnectedOrgConfig (*atlasv2.GetConnectedOrgConfigApiParams) (*atlasv2.ConnectedOrgConfig, error)
}

// GetConnectedOrgConfig encapsulates the logic to manage different cloud providers.
func (s *Store) GetConnectedOrgConfig(params *atlasv2.GetConnectedOrgConfigApiParams) (*atlasv2.ConnectedOrgConfig, error) {
	result, _, err := s.clientv2.FederatedAuthenticationApi.GetConnectedOrgConfigWithParams(s.ctx, params).Execute()
	return &result, err
}

type GetFederationSettingsOperation interface {
	GetFederationSettings (*atlasv2.GetFederationSettingsApiParams) (*atlasv2.OrgFederationSettings, error)
}

// GetFederationSettings encapsulates the logic to manage different cloud providers.
func (s *Store) GetFederationSettings(params *atlasv2.GetFederationSettingsApiParams) (*atlasv2.OrgFederationSettings, error) {
	result, _, err := s.clientv2.FederatedAuthenticationApi.GetFederationSettingsWithParams(s.ctx, params).Execute()
	return &result, err
}

type GetIdentityProviderOperation interface {
	GetIdentityProvider (*atlasv2.GetIdentityProviderApiParams) (*atlasv2.IdentityProvider, error)
}

// GetIdentityProvider encapsulates the logic to manage different cloud providers.
func (s *Store) GetIdentityProvider(params *atlasv2.GetIdentityProviderApiParams) (*atlasv2.IdentityProvider, error) {
	result, _, err := s.clientv2.FederatedAuthenticationApi.GetIdentityProviderWithParams(s.ctx, params).Execute()
	return &result, err
}

type GetIdentityProviderMetadataOperation interface {
	GetIdentityProviderMetadata (*atlasv2.GetIdentityProviderMetadataApiParams) (*atlasv2.string, error)
}

// GetIdentityProviderMetadata encapsulates the logic to manage different cloud providers.
func (s *Store) GetIdentityProviderMetadata(params *atlasv2.GetIdentityProviderMetadataApiParams) (*atlasv2.string, error) {
	result, _, err := s.clientv2.FederatedAuthenticationApi.GetIdentityProviderMetadataWithParams(s.ctx, params).Execute()
	return &result, err
}

type GetRoleMappingOperation interface {
	GetRoleMapping (*atlasv2.GetRoleMappingApiParams) (*atlasv2.RoleMapping, error)
}

// GetRoleMapping encapsulates the logic to manage different cloud providers.
func (s *Store) GetRoleMapping(params *atlasv2.GetRoleMappingApiParams) (*atlasv2.RoleMapping, error) {
	result, _, err := s.clientv2.FederatedAuthenticationApi.GetRoleMappingWithParams(s.ctx, params).Execute()
	return &result, err
}

type ListConnectedOrgConfigsOperation interface {
	ListConnectedOrgConfigs (*atlasv2.ListConnectedOrgConfigsApiParams) (*atlasv2.[]ConnectedOrgConfig, error)
}

// ListConnectedOrgConfigs encapsulates the logic to manage different cloud providers.
func (s *Store) ListConnectedOrgConfigs(params *atlasv2.ListConnectedOrgConfigsApiParams) (*atlasv2.[]ConnectedOrgConfig, error) {
	result, _, err := s.clientv2.FederatedAuthenticationApi.ListConnectedOrgConfigsWithParams(s.ctx, params).Execute()
	return &result, err
}

type ListIdentityProvidersOperation interface {
	ListIdentityProviders (*atlasv2.ListIdentityProvidersApiParams) (*atlasv2.[]IdentityProvider, error)
}

// ListIdentityProviders encapsulates the logic to manage different cloud providers.
func (s *Store) ListIdentityProviders(params *atlasv2.ListIdentityProvidersApiParams) (*atlasv2.[]IdentityProvider, error) {
	result, _, err := s.clientv2.FederatedAuthenticationApi.ListIdentityProvidersWithParams(s.ctx, params).Execute()
	return &result, err
}

type ListRoleMappingsOperation interface {
	ListRoleMappings (*atlasv2.ListRoleMappingsApiParams) (*atlasv2.[]RoleMapping, error)
}

// ListRoleMappings encapsulates the logic to manage different cloud providers.
func (s *Store) ListRoleMappings(params *atlasv2.ListRoleMappingsApiParams) (*atlasv2.[]RoleMapping, error) {
	result, _, err := s.clientv2.FederatedAuthenticationApi.ListRoleMappingsWithParams(s.ctx, params).Execute()
	return &result, err
}

type RemoveConnectedOrgConfigOperation interface {
	RemoveConnectedOrgConfig (*atlasv2.RemoveConnectedOrgConfigApiParams) (*atlasv2.map[string]interface{}, error)
}

// RemoveConnectedOrgConfig encapsulates the logic to manage different cloud providers.
func (s *Store) RemoveConnectedOrgConfig(params *atlasv2.RemoveConnectedOrgConfigApiParams) (*atlasv2.map[string]interface{}, error) {
	result, _, err := s.clientv2.FederatedAuthenticationApi.RemoveConnectedOrgConfigWithParams(s.ctx, params).Execute()
	return &result, err
}

type UpdateConnectedOrgConfigOperation interface {
	UpdateConnectedOrgConfig (*atlasv2.UpdateConnectedOrgConfigApiParams) (*atlasv2.ConnectedOrgConfig, error)
}

// UpdateConnectedOrgConfig encapsulates the logic to manage different cloud providers.
func (s *Store) UpdateConnectedOrgConfig(params *atlasv2.UpdateConnectedOrgConfigApiParams) (*atlasv2.ConnectedOrgConfig, error) {
	result, _, err := s.clientv2.FederatedAuthenticationApi.UpdateConnectedOrgConfigWithParams(s.ctx, params).Execute()
	return &result, err
}

type UpdateIdentityProviderOperation interface {
	UpdateIdentityProvider (*atlasv2.UpdateIdentityProviderApiParams) (*atlasv2.IdentityProvider, error)
}

// UpdateIdentityProvider encapsulates the logic to manage different cloud providers.
func (s *Store) UpdateIdentityProvider(params *atlasv2.UpdateIdentityProviderApiParams) (*atlasv2.IdentityProvider, error) {
	result, _, err := s.clientv2.FederatedAuthenticationApi.UpdateIdentityProviderWithParams(s.ctx, params).Execute()
	return &result, err
}

type UpdateRoleMappingOperation interface {
	UpdateRoleMapping (*atlasv2.UpdateRoleMappingApiParams) (*atlasv2.RoleMapping, error)
}

// UpdateRoleMapping encapsulates the logic to manage different cloud providers.
func (s *Store) UpdateRoleMapping(params *atlasv2.UpdateRoleMappingApiParams) (*atlasv2.RoleMapping, error) {
	result, _, err := s.clientv2.FederatedAuthenticationApi.UpdateRoleMappingWithParams(s.ctx, params).Execute()
	return &result, err
}

