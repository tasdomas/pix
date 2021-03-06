package config_test

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	gc "gopkg.in/check.v1"

	"github.com/tasdomas/pix/config"
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
secret: thisisasecret
`))
	cfg, err := config.Load(fname)
	c.Assert(err, gc.IsNil)
	c.Assert(cfg, gc.DeepEquals, &config.Config{
		Root:    "/opt",
		Storage: "/var/media",
		Port:    ":8080",
		Secret:  "thisisasecret",
		Name:    "3pxls",
		GAID:    "",
	})
}

func (t *TSuite) TestReadEmptyConfig(c *gc.C) {
	temp := c.MkDir()
	fname := path.Join(temp, "cfg.yml")

	mustWrite(fname, []byte(``))
	cfg, err := config.Load(fname)
	c.Assert(err, gc.IsNil)
	c.Assert(cfg, gc.DeepEquals, &config.Config{
		Root:    "./",
		Storage: "./files",
		Port:    ":8080",
		Secret:  "",
		Name:    "3pxls",
		GAID:    "",
	})
}

func (t *TSuite) TestReadFromEnv(c *gc.C) {
	os.Setenv("PIX_NAME", "blahdiblah")
	cfg, err := config.LoadFromEnv()
	c.Assert(err, gc.IsNil)
	c.Assert(cfg.Name, gc.Equals, "blahdiblah")
}
