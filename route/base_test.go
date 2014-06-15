package route

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RouteBaseTests struct {
	suite.Suite
}

func (suite *RouteBaseTests) SetupTest() {

}

func (suite *RouteBaseTests) TestResponseNotFound() {
	expectedCode := http.StatusNotFound
	expectedData := map[string]interface{}{"status": "failure"}

	actualCode, actualData := ResponseNotFound()

	assert.Equal(suite.T(), actualCode, expectedCode, "Status code returned should "+
		"be 404, not found")
	assert.Equal(suite.T(), actualData, expectedData, "Data should be status "+
		"failure")
}

func (suite *RouteBaseTests) TestInternalServerError_WithError() {
	testError := errors.New("Test error")
	expectedCode := http.StatusInternalServerError
	expectedData := map[string]interface{}{
		"status":        "failure",
		"error_message": testError.Error(),
	}

	actualCode, actualData := ResponseInternalServerError(testError)

	assert.Equal(suite.T(), actualCode, expectedCode, "Status code should be "+
		"500, internal server error")
	assert.Equal(suite.T(), actualData, expectedData, "The data sent back in "+
		"response should be status message and error_message")
}

func (suite *RouteBaseTests) TestResponseSuccessWithData() {
	testData := struct{}{}
	testResourceName := "resource"
	expectedCode := http.StatusAccepted
	expectedData := map[string]interface{}{
		"status": "success",
		"data": map[string]interface{}{
			testResourceName: testData,
		},
	}

	actualCode, actualData := ResponseSuccess(testData, testResourceName)

	assert.Equal(suite.T(), actualCode, expectedCode, "Status code should be 200")
	assert.Equal(suite.T(), actualData, expectedData, "Data will have a success "+
		"status with the data added along with the resource name")

}
func (suite *RouteBaseTests) TestResponseSuccessWithoutData() {
	testData := map[string]interface{}{}
	testResourceName := "resource"
	expectedCode := http.StatusAccepted
	expectedData := map[string]interface{}{
		"status": "success",
		"data": map[string]interface{}{
			testResourceName: testData,
		},
	}

	actualCode, actualData := ResponseSuccess(testData, testResourceName)

	assert.Equal(suite.T(), actualCode, expectedCode, "Status code should be 200")
	assert.Equal(suite.T(), actualData, expectedData, "Data will have a success "+
		"status with the data added along with the resource name")
}

func TestRouteBaseTests(t *testing.T) {
	suite.Run(t, new(RouteBaseTests))

}
