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

// ServiceBindingRoute represents a Cloud Controller service binding route.
type ServiceBindingRoute struct {
	// GUID is the unique service binding route identifier.
	GUID string

	// Name is the name given to the service binding route.
	Name string
}

// MarshalJSON converts an service binding route into a Cloud Controller service binding route.
func (o ServiceBindingRoute) MarshalJSON() ([]byte, error) {
	ccObj := struct {
		Name string `json:"name"`
	}{
		Name: o.Name,
	}

	return json.Marshal(ccObj)
}

// UnmarshalJSON helps unmarshal a Cloud Controller service binding route response.
func (o *ServiceBindingRoute) UnmarshalJSON(data []byte) error {
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

// GetServiceBindingRoutes returns back a list of service binding routes based off of the
// provided filters.
func (client *Client) GetServiceBindingRoutes(filters ...Filter) ([]ServiceBindingRoute, Warnings, error) {
	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.GetServiceBindingRoutesRequest,
		Query:       ConvertFilterParameters(filters),
	})
	if err != nil {
		return nil, nil, err
	}

	var fullObjList []ServiceBindingRoute
	warnings, err := client.paginate(request, ServiceBindingRoute{}, func(item interface{}) error {
		if app, ok := item.(ServiceBindingRoute); ok {
			fullObjList = append(fullObjList, app)
		} else {
			return ccerror.UnknownObjectInListError{
				Expected:   ServiceBindingRoute{},
				Unexpected: item,
			}
		}
		return nil
	})

	return fullObjList, warnings, err
}

// CreateServiceBindingRoute creates a cloud controller service binding route in with the given settings.
func (client *Client) CreateServiceBindingRoute(serviceBindingRoute ServiceBindingRoute) (ServiceBindingRoute, Warnings, error) {
	body, err := json.Marshal(serviceBindingRoute)
	if err != nil {
		return ServiceBindingRoute{}, nil, err
	}

	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.PostServiceBindingRoutesRequest,
		Body:        bytes.NewReader(body),
	})
	if err != nil {
		return ServiceBindingRoute{}, nil, err
	}

	var updatedObj ServiceBindingRoute
	response := cloudcontroller.Response{
		DecodeJSONResponseInto: &updatedObj,
	}

	err = client.connection.Make(request, &response)
	return updatedObj, response.Warnings, err
}

// DeleteServiceBindingRoute delete a service binding route
func (client *Client) DeleteServiceBindingRoute(serviceBindingGuid, routeGuid string) (Warnings, error) {
	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.DeleteServiceBindingRouteRequest,
		URIParams: Params{
			"service_binding_guid": serviceBindingGuid,
			"route_guid":           routeGuid,
		},
	})
	if err != nil {
		return nil, err
	}

	response := cloudcontroller.Response{}

	err = client.connection.Make(request, &response)
	return response.Warnings, err
}

// GetServiceBindingRoute returns back a service binding route.
func (client *Client) GetServiceBindingRoute(serviceBindingGuid, routeGuid string) (ServiceBindingRoute, Warnings, error) {
	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.GetServiceBindingRouteRequest,
		URIParams: Params{
			"service_binding_guid": serviceBindingGuid,
			"route_guid":           routeGuid,
		},
	})
	if err != nil {
		return ServiceBindingRoute{}, nil, err
	}

	var obj ServiceBindingRoute
	response := cloudcontroller.Response{
		DecodeJSONResponseInto: &obj,
	}

	err = client.connection.Make(request, &response)
	return obj, response.Warnings, err
}

// end:==kil--sl---sl==
