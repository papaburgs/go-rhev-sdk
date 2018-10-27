// Package rhvsdk will handle all of the details around the interactions
// with the RHEV managers
package rhvsdk

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type ssoResponseJSON struct {
	AccessToken string `json:"access_token"`
	// do not know if these exist in RHEV
	//	SsoError     string `json:"error"`
	//	SsoErrorCode string `json:"error_code"`
	Scope     string `json:"scope"`
	Expire    string `json:"exp"`
	TokenType string `json:"token_type"`
}

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

// NewConnection will create a new connection object
func NewConnection(server, user, pass string) (*Connection, error) {

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
	c.ssoTokenName = "access_token"

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
	c.buildClient()
	c.getToken()
	return c.err
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
		return
	}
	if len(c.password) == 0 {
		c.err = errors.New("Password cannot be empty")
		return
	}
}

func (c *Connection) loadCACert() {
	// if there is an error already, just return
	if c.err != nil {
		return
	}
	// if we have contents, we will asuming this doesn't need to be done
	if len(c.caContents) > 0 {
		return
	}

	if c.caFile == "" {
		c.err = errors.New("caFile is required")
		return
	}
	c.caContents, c.err = ioutil.ReadFile(c.caFile)
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
func (c *Connection) SetCAFileContents(content []byte) {
	c.caContents = content
}

// buildClient creates the client
func (c *Connection) buildClient() {
	var tlsConfig *tls.Config
	c.loadCACert()
	if c.err != nil {
		return
	}

	tlsConfig = &tls.Config{
		InsecureSkipVerify: c.insecure,
	}
	if c.url.Scheme == "https" {
		pool := x509.NewCertPool()
		if !pool.AppendCertsFromPEM(c.caContents) {
			c.err = errors.New("Failed to parse CA Certificate in contents")
		}
		tlsConfig.RootCAs = pool
	}
	c.client = &http.Client{
		Timeout: time.Duration(time.Second * 30),
		Transport: &http.Transport{
			// Close the http connection after calling resp.Body.Close()
			DisableKeepAlives: true,
			TLSClientConfig:   tlsConfig,
		},
	}
	return
}

func (c *Connection) getToken() {
	var ssoResp *ssoResponseJSON
	if c.err != nil {
		return
	}
	if c.ssoToken != "" {
		// assume we have a token we are ok
		fmt.Println("have token")
		return
	}

	// Build the URL and parameters required for the request:
	parameters := c.buildSsoAuthRequest()
	// Send the response and wait for the request:
	ssoResp, c.err = c.getSsoResponse(parameters)
	if c.err != nil {
		return
	}
	c.ssoToken = ssoResp.AccessToken
}

func (c *Connection) buildSsoAuthRequest() map[string]string {
	// Compute the entry point and the parameters:
	parameters := map[string]string{
		"scope":      "ovirt-app-api",
		"grant_type": "password",
		"username":   c.username,
		"password":   c.password,
	}

	// Compute the URL:
	c.url.Path = fmt.Sprintf("/ovirt-engine/sso/oauth/token")

	// Return the URL and the parameters:
	return parameters
}

func (c *Connection) getSsoResponse(parameters map[string]string) (*ssoResponseJSON, error) {
	var err error
	// POST request body:
	formValues := make(url.Values)
	for k1, v1 := range parameters {
		formValues[k1] = []string{v1}
	}
	// Build the net/http request:
	req, err := http.NewRequest("POST",
		c.url.String(),
		strings.NewReader(formValues.Encode()),
	)
	if err != nil {
		return nil, err
	}

	// Add request headers:
	req.Header.Add("User-Agent", "GoSDK/1")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	// Send the request and wait for the response:
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse and return the JSON response:
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var jsonObj ssoResponseJSON
	err = json.Unmarshal(body, &jsonObj)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse non-array sso with response %v", string(body))
	}
	// Unmarshal successfully
	if jsonObj.AccessToken == "" {
		return nil, fmt.Errorf("Did not get a token. Response: %s", string(body))
	}
	return &jsonObj, nil
}

// getResponseBody takes an API endpoint path and returns the body as a slice of bytes
func (c *Connection) getResponseBody(path string) ([]byte, error) {
	var err error

	// Compute the URL:
	c.url.Path = fmt.Sprintf(path)

	// Build the net/http request:
	req, err := http.NewRequest("GET",
		c.url.String(),
		nil,
	)
	if err != nil {
		return nil, err
	}
	// Add request headers:
	req.Header.Add("User-Agent", "GoSDK/1")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer "+c.ssoToken))

	// Send the request and wait for the response:
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse and return the JSON response:
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}
