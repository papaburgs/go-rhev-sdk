## Go RHEV sdk
This is a project to wrap the rhev api to make it easier to do basic task/reporting.

Note, this is not working at the moment and is being actively developed.

## Goals
Intend this to be more than a simple wrapper, but instead a collection of functions
that will be used to drive reports.  For instants, instead of just a function to get 
a list of VMs, there can be a function that will get all the details around that (assuming
that function doesn't already exist)

With this package, one would be able to make a stand alone app that could drive reports,
or periodically update a caching database to drive dashboards.

Also want to do it 'right', as in using oauth and not ignoring certs.

### First steps 
1. make a connection
2. define what a data structure will look like
2. get a list of vms
3. for a vm, get details
    * this could be all or a specific subset
    * a 'GetAll...' type of functions can be used to run all the smaller subsections

### Data structs
This will be populated as created.
To start, there will be a connection type (with a possible config type).
A VM will be a type, probably with nested types like interface or hard drives.
This will be modeled closely to the api's return values

### Usage
can clone this down in the normal go src directory using `go get`.

#### example
```go
cert := []byte("<cert text>")
api_url := "https://hostname"
user := "my-username"
pass := "my password"
```
