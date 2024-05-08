package v3action

import "code.cloudfoundry.org/cli/resources"

type RelationshipList resources.RelationshipList

func (actor Actor) ShareServiceInstanceToSpaces(serviceInstanceGUID string, spaceGUIDs []string) (RelationshipList, Warnings, error) {
	relationshipList, warnings, err := actor.CloudControllerClient.ShareServiceInstanceToSpaces(serviceInstanceGUID, spaceGUIDs)
	return RelationshipList(relationshipList), Warnings(warnings), err
}
