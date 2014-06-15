package route

import "net/http"

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
