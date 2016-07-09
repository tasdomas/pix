package ui_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tasdomas/pix/ui"
)

type NilStorage struct{}

func (*NilStorage) Put(_ io.ReadSeeker) (string, error)           { return "", nil }
func (*NilStorage) Get(_ string, _ string) (io.ReadCloser, error) { return nil, nil }
func (*NilStorage) List() ([]string, error)                       { return nil, nil }

func TestServeRoot(t *testing.T) {
	srv, err := ui.NewServer(&NilStorage{})
	if err != nil {
		t.Errorf("failed to create server: %v", err)
	}
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("failed to create request: %v", err)
	}

	rec := httptest.NewRecorder()
	srv.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("unexpected status: %d", rec.Code)
	}

}
