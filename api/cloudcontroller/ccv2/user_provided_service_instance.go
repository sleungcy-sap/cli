// begin:==kil--sl---sl==

package ccv2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"code.cloudfoundry.org/cli/api/cloudcontroller"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccerror"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2/constant"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2/internal"
	"code.cloudfoundry.org/cli/types"
)

// UserProvidedServiceInstance represents a Cloud Controller user provided service instance.
type UserProvidedServiceInstance struct {
	// GUID is the unique user provided service instance identifier.
	GUID string

	// Name is the name given to the user provided service instance.
	Name string
}

// MarshalJSON converts an user provided service instance into a Cloud Controller user provided service instance.
func (o UserProvidedServiceInstance) MarshalJSON() ([]byte, error) {
	ccObj := struct {
		Name string `json:"name"`
	}{
		Name: o.Name,
	}

	return json.Marshal(ccObj)
}

// UnmarshalJSON helps unmarshal a Cloud Controller user provided service instance response.
func (o *UserProvidedServiceInstance) UnmarshalJSON(data []byte) error {
	var ccObj struct {
		Metadata internal.Metadata `json:"metadata"`
		Entity   struct {
			Name string `json:"name"`
		} `json:"entity"`
	}
	err := cloudcontroller.DecodeJSON(data, &ccObj)
	if err != nil {
		return err
	}

	o.Name = ccObj.Entity.Name
	o.GUID = ccObj.Metadata.GUID

	return nil
}

// CreateUserProvidedServiceInstance creates a cloud controller user provided service instance in with the given settings.
func (client *Client) CreateUserProvidedServiceInstance(userProvidedServiceInstance UserProvidedServiceInstance) (UserProvidedServiceInstance, Warnings, error) {
	body, err := json.Marshal(userProvidedServiceInstance)
	if err != nil {
		return UserProvidedServiceInstance{}, nil, err
	}

	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.PostUserProvidedServiceInstancesRequest,
		Body:        bytes.NewReader(body),
	})
	if err != nil {
		return UserProvidedServiceInstance{}, nil, err
	}

	var updatedObj UserProvidedServiceInstance
	response := cloudcontroller.Response{
		DecodeJSONResponseInto: &updatedObj,
	}

	err = client.connection.Make(request, &response)
	return updatedObj, response.Warnings, err
}

// UpdateUserProvidedServiceInstance updates the user provided service instance with the given GUID.
func (client *Client) UpdateUserProvidedServiceInstance(userProvidedServiceInstance UserProvidedServiceInstance) (UserProvidedServiceInstance, Warnings, error) {
	body, err := json.Marshal(userProvidedServiceInstance)
	if err != nil {
		return UserProvidedServiceInstance{}, nil, err
	}

	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.PutUserProvidedServiceInstanceRequest,
		URIParams:   Params{"user_provided_service_instance_guid": userProvidedServiceInstance.GUID},
		Body:        bytes.NewReader(body),
	})
	if err != nil {
		return UserProvidedServiceInstance{}, nil, err
	}

	var updatedObj UserProvidedServiceInstance
	response := cloudcontroller.Response{
		DecodeJSONResponseInto: &updatedObj,
	}

	err = client.connection.Make(request, &response)
	return updatedObj, response.Warnings, err
}

// GetUserProvidedServiceInstance returns back a user provided service instance.
func (client *Client) GetUserProvidedServiceInstance(guid string) (UserProvidedServiceInstance, Warnings, error) {
	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.GetUserProvidedServiceInstanceRequest,
		URIParams: Params{
			"user_provided_service_instance_guid": guid,
		},
	})
	if err != nil {
		return UserProvidedServiceInstance{}, nil, err
	}

	var obj UserProvidedServiceInstance
	response := cloudcontroller.Response{
		DecodeJSONResponseInto: &obj,
	}

	err = client.connection.Make(request, &response)
	return obj, response.Warnings, err
}

// DeleteUserProvidedServiceInstance delete a user provided service instance
func (client *Client) DeleteUserProvidedServiceInstance(guid string) (Warnings, error) {
	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.DeleteUserProvidedServiceInstanceRequest,
		URIParams: Params{
			"user_provided_service_instance_guid": guid,
		},
	})
	if err != nil {
		return nil, err
	}

	response := cloudcontroller.Response{}

	err = client.connection.Make(request, &response)
	return response.Warnings, err
}

// end:==kil--sl---sl==
