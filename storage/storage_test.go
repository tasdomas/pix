package storage_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/tasdomas/pix/storage"
)

func TestStorage(t *testing.T) {
	stPath, err := ioutil.TempDir(".testdata/tmp", "")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(stPath)
	st, err := storage.New(stPath)
	if err != nil {
		t.Error(err)
	}
	f, err := os.Open(".testdata/pic.jpg")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	id, err := st.Put(f)
	if err != nil {
		t.Error(err)
	}
	if id == "" {
		t.Errorf("no valid id generated")
	}
	f.Seek(0, 0)
	stored, err := st.Get(id, "")
	if err != nil {
		t.Error(err)
	}
	defer stored.Close()
	bytes_f, err := ioutil.ReadAll(f)
	if err != nil {
		t.Error(err)
	}
	bytes_stored, err := ioutil.ReadAll(stored)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(bytes_f, bytes_stored) {
		t.Errorf("retrieved from storage does not match stored")
	}
}

func TestStorageIndex(t *testing.T) {
	stPath, err := ioutil.TempDir(".testdata/tmp", "")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(stPath)
	st, err := storage.New(stPath)
	if err != nil {
		t.Error(err)
	}
	f, err := os.Open(".testdata/pic.jpg")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	id, err := st.Put(f)
	if err != nil {
		t.Error(err)
	}
	if id == "" {
		t.Errorf("no valid id generated")
	}
	l, err := st.List()
	if err != nil {
		t.Error(err)
	}
	if l[0] != id {
		t.Errorf("list does not contain inserted id")
	}

	// Test index reloading.
	st2, err := storage.New(stPath)
	if err != nil {
		t.Error(err)
	}
	l, err = st2.List()
	if err != nil {
		t.Error(err)
	}
	if l[0] != id {
		t.Errorf("list does not contain inserted id")
	}

}
