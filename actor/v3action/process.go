package v3action

import (
	"code.cloudfoundry.org/cli/actor/actionerror"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccerror"
	"code.cloudfoundry.org/cli/resources"
)

// Process represents a V3 actor process.
type Process resources.Process

func (actor Actor) ScaleProcessByApplication(appGUID string, process Process) (Warnings, error) {
	_, warnings, err := actor.CloudControllerClient.CreateApplicationProcessScale(appGUID, resources.Process(process))
	allWarnings := Warnings(warnings)
	if err != nil {
		if _, ok := err.(ccerror.ProcessNotFoundError); ok {
			return allWarnings, actionerror.ProcessNotFoundError{ProcessType: process.Type}
		}
		return allWarnings, err
	}

	return allWarnings, nil
}
