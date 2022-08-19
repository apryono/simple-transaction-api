package helper

import (
	"encoding/json"
	"net/http"
)

// Response function response to json
func Response(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(data)
}

// Message function response
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"success": status, "message": message}
}

// ResponseErr return error response
func ResponseErr(w http.ResponseWriter, code int, data interface{}) {
	msg := Message(false, data.(string))
	Response(w, code, msg)
}
