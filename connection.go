// Package rhvlib will handle all of the details around the interactions
// with the RHEV managers
package rhvlib

import (
	"crypto/tls"
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/Sirupsen/logrus"
)

// Connection represents the connection to a RHEV engine
type Connection struct {
	url        *url.URL
	server     string
	username   string
	password   string
	token      string
	insecure   bool
	caFile     string
	caContents []byte
	headers    map[string]string

//	kerberos bool
//	timeout  time.Duration
//	compress bool
	// http client
	client *http.Client
	// SSO attributes
	ssoToken     string
	ssoTokenName string
	err          error
}


// NewConnection will create a new connection object and setup logging
func NewConnection(server, user, pass, level string) (*Connection, error) {
	var (
		err      error
	)

	c := Connection{
		err:      nil,
		username: user,
		password: pass,
		server:   server,
	}

	if c.server == "" {
		c.err = errors.New("server cannot be empty")
	} else {
		c.url, c.err = url.Parse(server)
	}
	c.insecure = false

	return &c, c.err
}

// SetInsecure sets the connection to ignore the cert file, should only be
// used for testing.
func (c *Connection) SetInsecure() {
	c.insecure = true
}

func (c *Connection) Connect() error {
	if c.err != nil {
		return c.err
	}
	c.validateOptions()
	c.loadCACert()
	c.getToken()
	//c.build()
	return nil
}

// validateOptions makes sure we have some decent values to work with
// basically, no empty servers or credentials
func (c *Connection) validateOptions() {
	// if there is an error already, just return
	if c.err != nil {
		return
	}
	if len(c.username) == 0 {
		c.err = errors.New("Username cannot be empty")
	}
	if len(c.password) == 0 {
		c.err = errors.New("Password cannot be empty")
	}
}

func (c *Connection) loadCACert() {
	var tlsConfig *tls.Config
	// if there is an error already, just return
	if c.err != nil {
		return
	}
	if c.insecure {
		tlsConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
		return
	}
	tlsConfig = &tls.Config{
		InsecureSkipVerify: false,
	}
    // if we have contents, we will asuming this doesn't need to be done
	if len(caContents) > 0 {
		return
	}

	if c.caFile == "" {
		c.err = errors.New("caFile is required")
		return
	}
	c.caContents, c.err = ioutils.ReadFile(c.caFile)
}

// SetCAFilePath set the path for the CA cert file
// This function or the SetCAFileContents needs to be run to set
// Cert details
func (c *Connection) SetCAFilePath(path string) {
	c.caFile = path
}

// SetCAFileContents sets the content of a ca cert
// Either this function or the SetCAFilePath need to be called
// when using a cert
func (c *Connection)SetCAFileContents(content []byte) {
	c.caContents = content
}
// getToken does some cool stuff
func (c *Connection) getToken() {
	// TODO work here next, full from line 526 of go-ovirt
}
