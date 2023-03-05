package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_readJSON(t *testing.T) {
	// create sample JSON
	sampleJSON := map[string]interface{}{
		"foo": "bar",
	}

	body, _ := json.Marshal(sampleJSON)

	var decodedJSON struct {
		FOO string `json:"foo"`
	}

	// create request
	req, err := http.NewRequest("POST", "/", bytes.NewReader(body))
	if err != nil {
		t.Log(err)
	}

	// create a test response recorder
	rr := httptest.NewRecorder()
	defer req.Body.Close()

	// call readJSON
	err = testApp.readJSON(rr, req, &decodedJSON)
	if err != nil {
		t.Log("Failed to decode the json", err)
	}

}

func Test_writeJSON(t *testing.T) {
	// create a test response recorder
	rr := httptest.NewRecorder()

	payload := jsonResponse{
		Error:   false,
		Message: "Howdy",
	}

	headers := make(http.Header)
	headers.Add("FOO", "BAR")

	// call writeJSON
	err := testApp.writeJSON(rr, http.StatusOK, payload, headers)
	if err != nil {
		t.Errorf("Failed to write the JSON: %v", err)
	}

	// test production
	testApp.environment = "production"

	err = testApp.writeJSON(rr, http.StatusOK, payload, headers)
	if err != nil {
		t.Errorf("Failed to write the JSON in Production: %v", err)
	}

	testApp.environment = "development"

}
