package tests

import (
	"client_api/src/domains"
	"client_api/src/portpb"
	"client_api/src/services"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"io/ioutil"
	"strings"
	"testing"
)

var importFunc func(<-chan portpb.Port, *grpc.ClientConn) (*domains.ImportResponse, error)

// Example end to end test using the port_service mock
func TestController_Import(t *testing.T) {

	// mocking PortService
	services.PortService = &portServiceMock{}
	// mocking PortService Import method
	importFunc = SuccessfulImportOneInserted

	res, err := client.Post(fmt.Sprintf("%s/ports", testServer.URL), "application/json",
		strings.NewReader(`{
  		"AEDXB": {
    	"name": "Dubai",
    	"coordinates": [
      		55.27,
      		25.25
    	],
    	"city": "Dubai",
    	"province": "Dubayy [Dubai]",
    	"country": "United Arab Emirates",
    	"alias": [],
    	"regions": [],
    	"timezone": "Asia/Dubai",
    	"unlocs": [
      		"AEDXB"
		],
    	"code": "52005"}
		}`))
	if err != nil {
		t.Fatal("unable to execute request", err)
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal("unable to read response body", err)
	}
	var respImported domains.ImportResponse
	err = json.Unmarshal(bodyBytes, &respImported)
	if err != nil {
		t.Fatal("unable to unmarshal response", err)
	}
	assert.Equal(t, 201, res.StatusCode)
	assert.Equal(t, int32(1), respImported.NumberInserted)
}

type portServiceMock struct {}

func (p portServiceMock) Import(portCh <-chan portpb.Port, conn *grpc.ClientConn) (*domains.ImportResponse, error) {
	return importFunc(portCh, conn)
}

func (p portServiceMock) Get(key string, conn *grpc.ClientConn) (*domains.Port, error) {
	panic("implement me")
}

func SuccessfulImportOneInserted(portCh <-chan portpb.Port, conn *grpc.ClientConn) (*domains.ImportResponse, error) {
	resp := &domains.ImportResponse{
		NumberInserted:  1,
		NumberUpdated:   0,
		EncounterErrors: false,
	}
	return resp, nil
}
