package health_check

import (
	"encoding/json"
	"net/http"
)

type healthJson struct {
	Status string `json:"status"`
}

func Initialize() {

	http.HandleFunc("/health", healthHandler)

}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(health()))
}


func health() string {
	healthJson := new(healthJson)

	healthJson.Status = "UP"

	b, err := json.Marshal(healthJson)
	if err != nil {
		return err.Error()
	}
	return string(b)
}