package config_test

import (
	"io/ioutil"
	"path"
	"testing"

	gc "gopkg.in/check.v1"

	"github.com/tasdomas/pixserver/config"
)

func mustWrite(filename string, contents []byte) {
	err := ioutil.WriteFile(filename, contents, 0666)
	if err != nil {
		panic(err)
	}
}

func Test(t *testing.T) {
	gc.TestingT(t)
}

type TSuite struct{}

var _ = gc.Suite(&TSuite{})

func (t *TSuite) TestReadConfig(c *gc.C) {
	temp := c.MkDir()
	fname := path.Join(temp, "cfg.yml")

	mustWrite(fname, []byte(`
root: /opt
storage: /var/media
port: :8080
`))
	cfg, err := config.Load(fname)
	c.Assert(err, gc.IsNil)
	c.Assert(cfg, gc.DeepEquals, &config.Config{
		Root:    "/opt",
		Storage: "/var/media",
		Port:    ":8080",
	})
}
