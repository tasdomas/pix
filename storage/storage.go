package storage

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/juju/errgo"
)

// storage will implement on-disk storage of images.
// To keep it simple at the moment it will implement the following interface:
type Storage interface {
	Put(f io.ReadSeeker) (string, error)
	Get(string) (io.ReadCloser, error)
	//GetThumb(string) (io.Read, error)
}

type storage struct {
	dir string
}

var _ Storage = (*storage)(nil)

func New(dir string) (*storage, error) {
	return &storage{
		dir: dir,
	}, nil
}

func (s *storage) Get(id string) (io.ReadCloser, error) {
	f, err := os.Open(path.Join(s.dir, id))
	if err != nil {
		return nil, errgo.Mask(err)
	}
	return f, nil
}

func (s *storage) Put(f io.ReadSeeker) (string, error) {
	_, err := f.Seek(0, 0)
	if err != nil {
		return "", errgo.Mask(err)
	}

	tempName, err := randomName()
	if err != nil {
		return "", errgo.Mask(err)
	}
	tempName = path.Join(s.dir, tempName)

	dst, err := os.OpenFile(tempName, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		return "", errgo.Mask(err)
	}
	defer dst.Close()
	h := md5.New()

	w := io.MultiWriter(dst, h)
	_, err = io.Copy(w, f)
	if err != nil {
		return "", errgo.Mask(err)
	}

	hashName := fmt.Sprintf("%x", h.Sum(nil))
	err = os.Rename(tempName, path.Join(s.dir, hashName))
	if err != nil {
		return "", errgo.Mask(err)
	}
	return hashName, nil
}

func randomName() (string, error) {
	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		return "", errgo.Mask(err)
	}

	return hex.EncodeToString(b), nil
}
