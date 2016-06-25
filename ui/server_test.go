package ui_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tasdomas/pix/ui"
)

func TestServeRoot(t *testing.T) {
	srv, err := ui.NewServer(nil)
	if err != nil {
		t.Errorf("failed to create server: %v", err)
	}
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("failed to create request: %v", err)
	}

	rec := httptest.NewRecorder()
	srv.ServeHTTP(rec, req)
	if rec.Code != http.StatusNotFound {
		t.Errorf("unexpected status: %d", rec.Code)
	}

}
