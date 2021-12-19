// Copyright 2019 Squeeze Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"encoding/json"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}, indent bool) {
	var (
		response []byte
		err      error
	)

	if indent {
		response, err = json.MarshalIndent(payload, "", "  ")
	} else {
		response, err = json.Marshal(payload)
	}

	if err != nil {
		response, _ = json.Marshal(map[string]string{"error": err.Error()})
		code = http.StatusInternalServerError
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data := map[string]interface{}{
		"error": "",
		"data":  payload,
	}
	respondWithJSON(w, code, data, false)
}

func RespondWithJSONIndent(w http.ResponseWriter, code int, payload interface{}) {
	respondWithJSON(w, code, payload, true)
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	data := map[string]interface{}{
		"error": message,
		"data":  nil,
	}
	respondWithJSON(w, code, data, false)
}
