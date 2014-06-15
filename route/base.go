package route

import "net/http"

func ResponseNotFound() (int, map[string]interface{}) {
	return http.StatusNotFound, map[string]interface{}{"status": "failure"}
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
