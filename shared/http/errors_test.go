package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRespondWithJSONError(t *testing.T) {
	w := httptest.NewRecorder()
	RespondWithJSONError(w, errors.New("does not work"), http.StatusInternalServerError)
	if w.Code != http.StatusInternalServerError {
		t.Errorf("Unexpected status code %d", w.Code)
	}
	if w.Body.String() != `{"error":"does not work","status":500}` {
		t.Errorf("Unexpected response body %s", w.Body.String())
	}
}
