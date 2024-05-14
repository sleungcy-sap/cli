package ccv3

import (
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/internal"
	"code.cloudfoundry.org/cli/resources"
)

func (client *Client) GetProcessSidecars(processGuid string) ([]resources.Sidecar, Warnings, error) {
	var sidecars []resources.Sidecar

	_, warnings, err := client.MakeListRequest(RequestParams{
		RequestName:  internal.GetProcessSidecarsRequest,
		URIParams:    internal.Params{"process_guid": processGuid},
		ResponseBody: resources.Sidecar{},
		AppendToList: func(item interface{}) error {
			sidecars = append(sidecars, item.(resources.Sidecar))
			return nil
		},
	})

	return sidecars, warnings, err
}

func (client *Client) GetApplicationSidecars(appGuid string) ([]resources.Sidecar, Warnings, error) {
	var sidecars []resources.Sidecar

	_, warnings, err := client.MakeListRequest(RequestParams{
		RequestName:  internal.GetApplicationSidecarsRequest,
		URIParams:    internal.Params{"app_guid": appGuid},
		ResponseBody: resources.Sidecar{},
		AppendToList: func(item interface{}) error {
			sidecars = append(sidecars, item.(resources.Sidecar))
			return nil
		},
	})

	return sidecars, warnings, err
}

func (client *Client) CreateApplicationSidecar(appGUID string, sidecar resources.Sidecar) (resources.Sidecar, Warnings, error) {
	var responseBody resources.Sidecar
	_, warnings, err := client.MakeRequest(RequestParams{
		RequestName:  internal.PostApplicationSidecarRequest,
		URIParams:    internal.Params{"app_guid": appGUID},
		RequestBody:  sidecar,
		ResponseBody: &responseBody,
	})

	return responseBody, warnings, err
}

func (client *Client) GetSidecar(sidecarGuid string) (resources.Sidecar, Warnings, error) {
	var responseBody resources.Sidecar

	_, warnings, err := client.MakeRequest(RequestParams{
		RequestName:  internal.GetSidecarRequest,
		URIParams:    internal.Params{"sidecar_guid": sidecarGuid},
		ResponseBody: &responseBody,
	})

	return responseBody, warnings, err
}

func (client *Client) UpdateSidecar(sidecar resources.Sidecar) (resources.Sidecar, Warnings, error) {
	var responseBody resources.Sidecar

	_, warnings, err := client.MakeRequest(RequestParams{
		RequestName:  internal.PatchSidecarRequest,
		URIParams:    internal.Params{"sidecar_guid": sidecar.GUID},
		RequestBody:  sidecar,
		ResponseBody: &responseBody,
	})

	return responseBody, warnings, err
}

func (client Client) DeleteSidecar(sidecarGuid string) (JobURL, Warnings, error) {
	jobURL, warnings, err := client.MakeRequest(RequestParams{
		RequestName: internal.DeleteSidecarRequest,
		URIParams:   internal.Params{"sidecar_guid": sidecarGuid},
	})

	return jobURL, warnings, err
}
