package rhvsdk

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConnectionNoServerError(t *testing.T) {
	c, err := NewConnection("", "", "")
	assert.EqualError(t, err, "server cannot be empty")
	assert.EqualError(t, c.err, "server cannot be empty")
}

func TestNewConnectionDefaults(t *testing.T) {
	c, err := NewConnection("server", "", "")
	// allow NewConnection without a username/pass?
	assert.Nil(t, err)
	assert.Nil(t, c.err)
	// should be secure by default
	assert.False(t, c.insecure)
	// ssoTokenName should be "access_token" by default
	assert.Equal(t, c.ssoTokenName, "access_token")
}

func TestSetInsecure(t *testing.T) {
	c, _ := NewConnection("server", "", "")
	c.SetInsecure()
	assert.True(t, c.insecure)
}

func TestValidateOptions(t *testing.T) {
	// ensure err 'short-circuit' when no username provided
	c, _ := NewConnection("server", "", "")
	c.validateOptions()
	assert.EqualError(t, c.err, "Username cannot be empty")
	// ensure err 'short-circuit' when no username provided
	c, _ = NewConnection("server", "username", "")
	c.validateOptions()
	assert.EqualError(t, c.err, "Password cannot be empty")
}

func TestLoadCACert(t *testing.T) {
	c, err := NewConnection("server", "username", "password")
	assert.Nil(t, err)
	c.loadCACert()
	assert.EqualError(t, c.err, "caFile is required")

	// clear err for next test
	c.err = nil

	// write a fake caFile
	caContent := []byte("123456789")
	err = ioutil.WriteFile("/tmp/caFile", caContent, 0644)
	assert.Nil(t, err)

	// test that c.caContent is populated with file contents
	c.caFile = "/tmp/caFile"
	c.loadCACert()
	assert.Equal(t, c.caContents, []byte("123456789"))
	assert.Nil(t, c.err)

	// remove fake caFile
	os.Remove("/tmp/caFile")

	// with c.caContent populated, the file should not be read again
	c.caFile = ""
	c.loadCACert()
	assert.Nil(t, c.err)
}

func TestSetCAFIlePath(t *testing.T) {
	c, err := NewConnection("server", "username", "password")
	assert.Nil(t, err)
	assert.Equal(t, c.caFile, "")
	c.SetCAFilePath("/tmp/caFile")
	assert.Equal(t, c.caFile, "/tmp/caFile")
}

func TestSetCAFileContents(t *testing.T) {
	c, err := NewConnection("server", "username", "password")
	assert.Nil(t, err)
	assert.Nil(t, c.caContents)
	c.SetCAFileContents([]byte("abcxyz1234"))
	assert.Equal(t, c.caContents, []byte("abcxyz1234"))
}
