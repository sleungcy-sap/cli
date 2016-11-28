package common

import (
	"code.cloudfoundry.org/cli/actor/v3action"
	"code.cloudfoundry.org/cli/api/cloudcontroller"
)

func HandleError(err error) error {
	switch e := err.(type) {
	case cloudcontroller.RequestError:
		return APIRequestError{Err: e.Err}
	case cloudcontroller.UnverifiedServerError:
		return InvalidSSLCertError{API: e.URL}
	case cloudcontroller.UnprocessableEntityError:
		if e.Message == "The request is semantically invalid: Task must have a droplet. Specify droplet or assign current droplet to app." {
			return RunTaskError{
				Message: "App is not staged."}
		}
	case cloudcontroller.APINotFoundError:
		return APINotFoundError{URL: e.URL}
	case v3action.ApplicationNotFoundError:
		return ApplicationNotFoundError{Name: e.Name}
	case v3action.TaskWorkersUnavailableError:
		return RunTaskError{Message: "Task workers are unavailable."}
	}

	return err
}
