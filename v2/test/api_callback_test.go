/*
HW Mux Reservation System

Testing CallbackAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package hwmux

import (
	"context"
	"testing"

	openapiclient "github.com/Silabs-UTF/hwmux-client-golang/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_hwmux_CallbackAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CallbackAPIService CallbackRetrieve", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.CallbackAPI.CallbackRetrieve(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
