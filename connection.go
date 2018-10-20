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

	kerberos bool
	timeout  time.Duration
	compress bool
	// http client
	client *http.Client
	// SSO attributes
	ssoToken     string
	ssoTokenName string
	err          error
	log          *logrus.Logger
}

// NewConnection will create a new connection object and setup logging
func NewConnection(server, user, pass, level string) (*Connection, error) {
	var (
		levelVal logrus.Level
		err      error
	)

	c := Connection{
		err:      nil,
		username: user,
		password: pass,
		server:   server,
	}
	c.log = logrus.New()
	levelVal, err = logrus.ParseLevel(level)
	if err != nil {
		c.log.SetLevel(logrus.InfoLevel)
		c.log.Warn("bad level passed in, default to Info level but continuing")
	}

	c.log.SetLevel(levelVal)

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
}

///TODO, working here!!!
// getToken does some cool stuff
func (c *Connection) getToken() {
}
