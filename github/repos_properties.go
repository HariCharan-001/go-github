// Copyright 2023 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
	"fmt"
)

// CustomPropertyNonNullValue represents a custom property, with a non-null value.
type CustomPropertyNonNullValue struct {
	PropertyName string `json:"property_name"`
	Value        string `json:"value"`
}

// GetAllCustomPropertyValues gets all custom property values that are set for a repository.
//
// GitHub API docs: https://docs.github.com/rest/repos/custom-properties#get-all-custom-property-values-for-a-repository
//
//meta:operation GET /repos/{owner}/{repo}/properties/values
func (s *RepositoriesService) GetAllCustomPropertyValues(ctx context.Context, org, repo string) ([]*CustomPropertyValue, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/properties/values", org, repo)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var customPropertyValues []*CustomPropertyValue
	resp, err := s.client.Do(ctx, req, &customPropertyValues)
	if err != nil {
		return nil, resp, err
	}

	return customPropertyValues, resp, nil
}

// CreateOrUpdateCustomPropertyValues creates new or updates existing custom property values for a repository.
//
// GitHub API docs: https://docs.github.com/en/rest/repos/custom-properties#create-or-update-custom-property-values-for-a-repository
//
//meta:operation PATCH /repos/{owner}/{repo}/properties/values
func (s *RepositoriesService) CreateOrUpdateCustomPropertyValues(ctx context.Context, org, repo string, customPropertyValues []*CustomPropertyNonNullValue) (*Response, error) {
	u := fmt.Sprintf("repos/%v/%v/properties/values", org, repo)

	params := struct {
		CustomPropertyValues []*CustomPropertyNonNullValue `json:"properties"`
	}{
		CustomPropertyValues: customPropertyValues,
	}

	req, err := s.client.NewRequest("PATCH", u, params)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
