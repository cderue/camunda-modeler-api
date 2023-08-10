/*
Web Modeler REST API

Testing CollaboratorsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_CollaboratorsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CollaboratorsAPIService DeleteCollaborator", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string
		var email string

		httpRes, err := apiClient.CollaboratorsAPI.DeleteCollaborator(context.Background(), projectId, email).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CollaboratorsAPIService ModifyCollaborator", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CollaboratorsAPI.ModifyCollaborator(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CollaboratorsAPIService SearchCollaborators", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CollaboratorsAPI.SearchCollaborators(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
