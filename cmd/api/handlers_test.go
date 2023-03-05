package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestApplication_AllUsers(t *testing.T) {
	// mock some rows, then add one
	var mockedRows = mockedDB.NewRows([]string{"id", "email", "first_name", "last_name", "password", "active", "created_at", "updated_at", "has_token"})

	mockedRows.AddRow("1", "me@here.com", "Stuart", "Little", "abc123", "1", time.Now(), time.Now(), "0")

	// tell mockDB how many queries to expect
	mockedDB.ExpectQuery("select \\\\* ").WillReturnRows(mockedRows)

	// create a test response recorder
	rr := httptest.NewRecorder()

	// create a request
	req, _ := http.NewRequest("POST", "/admin/users", nil)

	// call handler
	handler := http.HandlerFunc(testApp.AllUsers)
	handler.ServeHTTP(rr, req)

	// check for status code
	if rr.Code != http.StatusOK {
		t.Error("AllUsers returned wrong status code of ", rr.Code)
	}

}

func Test_errorJSON(t *testing.T) {
	rr := httptest.NewRecorder()
	err := testApp.errorJSON(rr, errors.New("Failed to write JSON"))
	if err != nil {
		t.Error(err)
	}

	testJSONPayload(t, rr)

	errSlice := []string{
		"{SQLSTATE 23505}",
		"{SQLSTATE 22001}",
		"{SQLSTATE 23503}",
	}

	for _, x := range errSlice {
		customErr := testApp.errorJSON(rr, errors.New(x), http.StatusUnauthorized)
		if customErr != nil {
			t.Error(customErr)
		}
		testJSONPayload(t, rr)
	}

}

func testJSONPayload(t *testing.T, rr *httptest.ResponseRecorder) {
	var requestPayload jsonResponse
	decoder := json.NewDecoder(rr.Body)
	err := decoder.Decode(&requestPayload)
	if err != nil {
		t.Error("received error when decoding errorJSON payload: ", err)
	}

	if !requestPayload.Error {
		t.Error("error set to false from errorJSON payload, but should be true")
	}

}
