package storage

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"image"
	_ "image/jpeg"
	"io"
	"os"
	"path"

	"github.com/disintegration/imaging"
	"github.com/fvbock/trie"
	"github.com/juju/errgo"
)

const (
	THB_DIM     = 200
	PREVIEW_DIM = 1000
	IDX_FILE    = ".idx"
)

// storage will implement on-disk storage of images.
// To keep it simple at the moment it will implement the following interface:
type Storage interface {
	Put(f io.ReadSeeker) (string, error)
	Get(string, string) (io.ReadCloser, error)
	List() ([]string, error)
}

type storage struct {
	entries *trie.Trie
	dir     string
}

var _ Storage = (*storage)(nil)

func New(dir string) (*storage, error) {
	os.MkdirAll(dir, 0700)
	// Check for reload.
	idx := trie.NewTrie()
	idxFile := path.Join(dir, IDX_FILE)
	if _, err := os.Stat(idxFile); err == nil {
		idx, err = trie.LoadFromFile(idxFile)
		if err != nil {
			return nil, errgo.Mask(err)
		}
	}

	return &storage{
		entries: idx,
		dir:     dir,
	}, nil
}

func (s *storage) Get(id string, mod string) (io.ReadCloser, error) {
	if mod != "" && mod != "full" {
		id = fmt.Sprintf("%s_%s", id, mod)
	}
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

	thbName := path.Join(s.dir, fmt.Sprintf("%s_thb", hashName))
	thb, err := os.OpenFile(thbName, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return "", errgo.Mask(err)
	}
	defer thb.Close()
	f.Seek(0, 0)
	err = generateThumbnail(f, thb)
	if err != nil {
		return "", errgo.Mask(err)
	}

	prvName := path.Join(s.dir, fmt.Sprintf("%s_large", hashName))
	prv, err := os.OpenFile(prvName, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return "", errgo.Mask(err)
	}
	defer prv.Close()
	f.Seek(0, 0)
	err = generatePreview(f, prv)
	if err != nil {
		return "", errgo.Mask(err)
	}

	s.entries.Add(hashName)
	// Ignoring this error for now.
	s.entries.DumpToFile(path.Join(s.dir, IDX_FILE))
	return hashName, nil
}

func generateThumbnail(in io.Reader, out io.Writer) error {
	img, _, err := image.Decode(in)
	if err != nil {
		return err
	}

	bounds := img.Bounds().Max
	var w int
	var h int
	if bounds.X >= bounds.Y {
		w = bounds.Y
		h = bounds.Y
	} else {
		w = bounds.X
		h = bounds.X
	}
	thb := imaging.CropCenter(img, w, h)
	thb = imaging.Resize(thb, THB_DIM, THB_DIM, imaging.Lanczos)
	return imaging.Encode(out, thb, imaging.JPEG)
}

func generatePreview(in io.Reader, out io.Writer) error {
	img, _, err := image.Decode(in)
	if err != nil {
		return err
	}

	bounds := img.Bounds().Max
	var w int
	var h int
	if bounds.X < PREVIEW_DIM && bounds.Y < PREVIEW_DIM {
		_, err := io.Copy(out, in)
		return err
	}
	if bounds.X >= bounds.Y {
		w = PREVIEW_DIM
	} else {
		h = PREVIEW_DIM
	}
	prv := imaging.Resize(img, w, h, imaging.Lanczos)
	return imaging.Encode(out, prv, imaging.JPEG)
}

func (s *storage) List() ([]string, error) {
	return s.entries.MembersList(), nil
}

func randomName() (string, error) {
	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		return "", errgo.Mask(err)
	}

	return hex.EncodeToString(b), nil
}
