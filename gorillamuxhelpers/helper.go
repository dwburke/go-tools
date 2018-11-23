package gorillamuxhelper

import (
	"encoding/json"
	"net/http"
)

func DecodeJsonBody(w http.ResponseWriter, r *http.Request) (obj map[string]interface{}, err error) {
	err = json.NewDecoder(r.Body).Decode(&obj)
	if err != nil {
		RespondWithError(w, 500, err.Error())
		return
	}

	return
}

func CheckRequiredVar(w http.ResponseWriter, vars interface{}, name string) bool {

	if v, ok := vars.(map[string]string); ok {
		if _, ok := v[name]; ok {
			if len(v[name]) == 0 {
				RespondWithError(w, 500, "required param '"+name+"' is blank")
				return false
			}
			return true
		}
	} else if v, ok := vars.(map[string]interface{}); ok {
		if _, ok := v[name]; ok {
			return true
		}
	}

	// exists?
	RespondWithError(w, 500, "required param '"+name+"' is missing")
	return false
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
