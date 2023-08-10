/*
Web Modeler REST API

Testing MilestonesAPIService

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

func Test_openapi_MilestonesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MilestonesAPIService CompareMilestones", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var milestone1Id string
		var milestone2Id string

		resp, httpRes, err := apiClient.MilestonesAPI.CompareMilestones(context.Background(), milestone1Id, milestone2Id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MilestonesAPIService CreateMilestone", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MilestonesAPI.CreateMilestone(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MilestonesAPIService DeleteMilestone", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var milestoneId string

		httpRes, err := apiClient.MilestonesAPI.DeleteMilestone(context.Background(), milestoneId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MilestonesAPIService GetMilestone", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var milestoneId string

		resp, httpRes, err := apiClient.MilestonesAPI.GetMilestone(context.Background(), milestoneId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MilestonesAPIService PatchMilestone", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var milestoneId string

		resp, httpRes, err := apiClient.MilestonesAPI.PatchMilestone(context.Background(), milestoneId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MilestonesAPIService RestoreMilestone", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var milestoneId string

		resp, httpRes, err := apiClient.MilestonesAPI.RestoreMilestone(context.Background(), milestoneId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MilestonesAPIService SearchMilestones", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MilestonesAPI.SearchMilestones(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
