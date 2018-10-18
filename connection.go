// Package rhvlib will handle all of the details around the interactions
// with the RHEV managers
// wrapping it in a package so different front ends can use it
package rhvlib

import (
	"net/http"
	"net/url"
	"time"

	"github.com/Sirupsen/logrus"
)

type Connection struct {
	url      *url.URL
	server   string
	username string
	password string
	token    string
	insecure bool
	caFile   string
	headers  map[string]string

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
	c.url, c.err = url.Parse(server)

	return &Connection, nil
}

// getToken does some cool stuff
func (c *Connection) getToken() {
}
