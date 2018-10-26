package rhvsdk

import (
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

// func TestFakeAPIResponse(t *testing.T) {
// 	httpmock.Activate()
// 	defer httpmock.DeactivateAndReset()

// 	httpmock.RegisterResponder("GET", "https://rhvserver/ovirt-engine/api/",
// 		httpmock.NewStringResponder(200, `[{"id": 1, "name": "server"}]`))

// 	resp, err := http.Get("https://rhvserver/ovirt-engine/api/")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	assert.Equal(t, string(body), `[{"id": 1, "name": "server"}]`)
// }
