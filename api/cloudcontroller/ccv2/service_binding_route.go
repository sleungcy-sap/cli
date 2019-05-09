package ccv2

import (
	"bytes"
	"code.cloudfoundry.org/cli/api/cloudcontroller"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccerror"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2/internal"
	"encoding/json"
	"net/url"
)

// ServiceBindingRoute represents a Cloud Controller service binding route.
type ServiceBindingRoute struct {
	// GUID is the unique service binding route identifier.
	GUID string

	// Name is the name given to the service binding route.
	Name string

	// The guid of the service instance to bind
	ServiceInstanceGuid string

	// The guid of the app to bind
	AppGuid string

	// Arbitrary parameters to pass along to the service broker.
	Parameters map[string]interface{}
}

// MarshalJSON converts an service binding route into a Cloud Controller service binding route.
func (o ServiceBindingRoute) MarshalJSON() ([]byte, error) {
	ccObj := struct {
		Name                string                 `json:"name"`
		ServiceInstanceGuid string                 `json:"service_instance_guid"`
		AppGuid             string                 `json:"app_guid"`
		Parameters          map[string]interface{} `json:"parameters,omitempty"`
	}{
		Name:                o.Name,
		ServiceInstanceGuid: o.ServiceInstanceGuid,
		AppGuid:             o.AppGuid,
		Parameters:          o.Parameters,
	}

	return json.Marshal(ccObj)
}

// UnmarshalJSON helps unmarshal a Cloud Controller service binding route response.
func (o *ServiceBindingRoute) UnmarshalJSON(data []byte) error {
	var ccObj struct {
		Metadata internal.Metadata `json:"metadata"`
		Entity   struct {
			Name                string                 `json:"name"`
			ServiceInstanceGuid string                 `json:"service_instance_guid"`
			AppGuid             string                 `json:"app_guid"`
			Parameters          map[string]interface{} `json:"parameters"`
		} `json:"entity"`
	}
	err := cloudcontroller.DecodeJSON(data, &ccObj)
	if err != nil {
		return err
	}

	o.Name = ccObj.Entity.Name
	o.GUID = ccObj.Metadata.GUID
	o.ServiceInstanceGuid = ccObj.Entity.ServiceInstanceGuid
	o.AppGuid = ccObj.Entity.AppGuid
	o.Parameters = ccObj.Entity.Parameters

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
		Query: url.Values{
			"accepts_incomplete": []string{"true"},
		},
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
