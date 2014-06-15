package route

import (
	"errors"
	"net/http"
	"net/url"
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
	expectedCode := http.StatusAccepted
	expectedData := map[string]interface{}{
		"status": "success",
		"data":   map[string]interface{}{},
	}

	actualCode, actualData := ResponseSuccessNoData()

	assert.Equal(suite.T(), actualCode, expectedCode, "Status code should be 200")
	assert.Equal(suite.T(), actualData, expectedData, "Data will have a success "+
		"status with the data added along with the resource name")
}

func (suite *RouteBaseTests) TestTransformQueryToMappingWithOneField() {
	testUrlQuery, _ := url.Parse("http://test.com/search?cap=tony")
	expected := map[string]interface{}{
		"fields": map[string]interface{}{
			"cap": "tony",
		},
	}
	urlValues := testUrlQuery.Query()

	actual, actualErr := TransformQueryToMapping(urlValues)
	assert.Equal(suite.T(), len(actualErr), 0, "No error was thrown")

	assert.Equal(suite.T(), actual, expected)
}

func (suite *RouteBaseTests) TestTransformQueryToMappingWithManyFields() {
	testUrlQuery, _ := url.Parse("http://test.com/search?cap=tony&price=12")
	expected := map[string]interface{}{
		"fields": map[string]interface{}{
			"cap":   "tony",
			"price": "12",
		},
	}
	urlValues := testUrlQuery.Query()

	actual, actualErr := TransformQueryToMapping(urlValues)
	assert.Equal(suite.T(), len(actualErr), 0, "No error was thrown")

	assert.Equal(suite.T(), actual, expected)
}

func (suite *RouteBaseTests) TestTransformQueryToMappingWithOnlyGroup() {
	testUrlQuery, _ := url.Parse("http://test.com/search?group-name=poop")
	expected := map[string]interface{}{
		"group": map[string]interface{}{
			"name": "poop",
		},
	}
	urlValues := testUrlQuery.Query()

	actual, actualErr := TransformQueryToMapping(urlValues)
	assert.Equal(suite.T(), len(actualErr), 0, "No error was thrown")

	assert.Equal(suite.T(), actual, expected)
}

func (suite *RouteBaseTests) TestTransformQueryToMappingWithGroupAndFields() {
	testUrlQuery, _ := url.Parse(
		"http://test.com/search?group-name=poop&cap=tony&price=12")
	expected := map[string]interface{}{
		"group": map[string]interface{}{
			"name": "poop",
		},
		"fields": map[string]interface{}{
			"cap":   "tony",
			"price": "12",
		},
	}
	urlValues := testUrlQuery.Query()

	actual, actualErr := TransformQueryToMapping(urlValues)
	assert.Equal(suite.T(), len(actualErr), 0, "No error was thrown")

	assert.Equal(suite.T(), actual, expected)
}

func (suite *RouteBaseTests) TestTransformQueryToMapping_WithInvalidFields() {
	testUrlQuery, _ := url.Parse(
		"http://test.com/search?group-name=poop&cap=tony&poop=12")
	urlValues := testUrlQuery.Query()
	expected := []error{errors.New("poop")}

	_, actualErr := TransformQueryToMapping(urlValues)

	assert.NotEqual(suite.T(), len(actualErr), 0, "An error was returned")
	assert.Equal(suite.T(), actualErr, expected)
}

func (suite *RouteBaseTests) TestTransformQueryToMapping_WithInvalidGroup() {
	testUrlQuery, _ := url.Parse(
		"http://test.com/search?group-name=poop&cap=tony&group=thename")
	urlValues := testUrlQuery.Query()
	expected := []error{errors.New("group")}

	_, actualErr := TransformQueryToMapping(urlValues)

	assert.NotEqual(suite.T(), len(actualErr), 0, "An error was returned")
	assert.Equal(suite.T(), actualErr, expected)
}

func TestRouteBaseTests(t *testing.T) {
	suite.Run(t, new(RouteBaseTests))

}
