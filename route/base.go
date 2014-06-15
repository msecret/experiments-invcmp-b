package route

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

var validSchemaFields = []string{
	"symbol",
	"group",
	"cap",
	"price",
}

func ResponseNotFound() (int, map[string]interface{}) {
	return http.StatusNotFound, map[string]interface{}{"status": "failure"}
}

func ResponseBadRequest(err error) (int, map[string]interface{}) {
	return http.StatusBadRequest, map[string]interface{}{
		"status":        "failure",
		"error_message": err.Error()}
}

func ResponseInternalServerError(err error) (int, map[string]interface{}) {
	return http.StatusInternalServerError, map[string]interface{}{
		"status":        "failure",
		"error_message": err.Error()}
}

func ResponseSuccess(resource interface{}, resourceName string) (int, map[string]interface{}) {
	return http.StatusAccepted, map[string]interface{}{
		"status": "success",
		"data": map[string]interface{}{
			resourceName: resource},
	}
}

// ResponseSuccessNoData will return a 200 success response for when no data has
// to be sent back to the client. The struct will still include a data attribute
// for consistency.
func ResponseSuccessNoData() (int, map[string]interface{}) {
	return http.StatusAccepted, map[string]interface{}{
		"status": "success",
		"data":   map[string]interface{}{},
	}
}

// TransformQueryToMapping will transform a url query to a mapping that is
// searchable in the database.
func TransformQueryToMapping(query url.Values) (map[string]interface{}, []error) {
	fields := map[string]interface{}{}
	returnErrors := make([]error, 0)

	for key, value := range query {
		// TODO take care of all fields that could be further namespaced
		// Take care of group keys
		// TODO move this to function
		if strings.HasPrefix(key, "group") {
			if _, ok := fields["group"]; !ok {
				fields["group"] = map[string]interface{}{}
			}
			// TODO dangerous splitting in case group not formatted correctly
			keySplit := strings.Split(key, "-")
			if len(keySplit) < 2 {
				returnErrors = append(returnErrors, errors.New(key))
				continue
			}
			fields["group"].(map[string]interface{})[keySplit[1]] = value[0]
			continue
		}
		if err := checkQueryKey(key); err != nil {
			returnErrors = append(returnErrors, errors.New(key))
			continue
		} else {
			if _, ok := fields["fields"]; !ok {
				fields["fields"] = map[string]interface{}{}
			}
			fields["fields"].(map[string]interface{})[key] = value[0]
		}
	}

	return fields, returnErrors
}

// checkQueryKey checks that the key is valid as in there is a definition for
// in it the schema. Returns an error if the key is not defined in the schema
//
// TODO: make an easy way to get these valid keys rather then having to maintain
// in code here
func checkQueryKey(key string) error {
	for _, validField := range validSchemaFields {
		if key == validField {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Invalid field: %s", key))
}
