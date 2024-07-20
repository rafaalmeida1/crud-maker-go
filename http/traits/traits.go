package traits

import (
    "encoding/json"
    "net/http"
)

func JsonResponse(w http.ResponseWriter, data interface{}, message string, status int) {
    response := map[string]interface{}{
        "data":    data,
        "message": message,
        "status":  status,
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(response)
}

func ErrorResponse(w http.ResponseWriter, message string, status int) {
    JsonResponse(w, nil, message, status)
}
