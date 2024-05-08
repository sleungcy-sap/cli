package v7action

import (
	"io"

	"code.cloudfoundry.org/cli/resources"

	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/constant"
)

//go:generate counterfeiter . CloudControllerClient

// CloudControllerClient is the interface to the cloud controller V3 API.
type CloudControllerClient interface {
	AppSSHEndpoint() string
	AppSSHHostKeyFingerprint() string
	CloudControllerAPIVersion() string
	CreateApplication(app resources.Application) (resources.Application, ccv3.Warnings, error)
	CreateApplicationProcessScale(appGUID string, process resources.Process) (resources.Process, ccv3.Warnings, error)
	CreateApplicationTask(appGUID string, task resources.Task) (resources.Task, ccv3.Warnings, error)
	CreateBuild(build resources.Build) (resources.Build, ccv3.Warnings, error)
	CreateBuildpack(bp resources.Buildpack) (resources.Buildpack, ccv3.Warnings, error)
	CreateDomain(domain resources.Domain) (resources.Domain, ccv3.Warnings, error)
	CreateIsolationSegment(isolationSegment resources.IsolationSegment) (resources.IsolationSegment, ccv3.Warnings, error)
	CreatePackage(pkg resources.Package) (resources.Package, ccv3.Warnings, error)
	DeleteApplication(guid string) (ccv3.JobURL, ccv3.Warnings, error)
	DeleteApplicationProcessInstance(appGUID string, processType string, instanceIndex int) (ccv3.Warnings, error)
	DeleteBuildpack(buildpackGUID string) (ccv3.JobURL, ccv3.Warnings, error)
	DeleteDomain(domainGUID string) (ccv3.JobURL, ccv3.Warnings, error)
	DeleteIsolationSegment(guid string) (ccv3.Warnings, error)
	DeleteIsolationSegmentOrganization(isolationSegmentGUID string, organizationGUID string) (ccv3.Warnings, error)
	DeleteServiceInstanceRelationshipsSharedSpace(serviceInstanceGUID string, sharedToSpaceGUID string) (ccv3.Warnings, error)
	EntitleIsolationSegmentToOrganizations(isoGUID string, orgGUIDs []string) (resources.RelationshipList, ccv3.Warnings, error)
	GetApplicationDropletCurrent(appGUID string) (resources.Droplet, ccv3.Warnings, error)
	GetApplicationEnvironment(appGUID string) (ccv3.Environment, ccv3.Warnings, error)
	GetApplicationManifest(appGUID string) ([]byte, ccv3.Warnings, error)
	GetApplicationProcessByType(appGUID string, processType string) (resources.Process, ccv3.Warnings, error)
	GetApplicationProcesses(appGUID string) ([]resources.Process, ccv3.Warnings, error)
	GetApplications(query ...ccv3.Query) ([]resources.Application, ccv3.Warnings, error)
	GetApplicationTasks(appGUID string, query ...ccv3.Query) ([]resources.Task, ccv3.Warnings, error)
	GetBuild(guid string) (resources.Build, ccv3.Warnings, error)
	GetBuildpacks(query ...ccv3.Query) ([]resources.Buildpack, ccv3.Warnings, error)
	GetDomains(query ...ccv3.Query) ([]resources.Domain, ccv3.Warnings, error)
	GetDroplet(guid string) (resources.Droplet, ccv3.Warnings, error)
	GetDroplets(query ...ccv3.Query) ([]resources.Droplet, ccv3.Warnings, error)
	GetEnvironmentVariableGroup(group constant.EnvironmentVariableGroupName) (ccv3.EnvironmentVariables, ccv3.Warnings, error)
	GetEvents(query ...ccv3.Query) ([]ccv3.Event, ccv3.Warnings, error)
	GetFeatureFlag(featureFlagName string) (resources.FeatureFlag, ccv3.Warnings, error)
	GetInfo() (ccv3.Info, ccv3.Warnings, error)
	GetIsolationSegment(guid string) (resources.IsolationSegment, ccv3.Warnings, error)
	GetIsolationSegmentOrganizations(isolationSegmentGUID string) ([]resources.Organization, ccv3.Warnings, error)
	GetIsolationSegments(query ...ccv3.Query) ([]resources.IsolationSegment, ccv3.Warnings, error)
	GetOrganizationDefaultIsolationSegment(orgGUID string) (resources.Relationship, ccv3.Warnings, error)
	GetOrganizationDomains(orgGUID string, query ...ccv3.Query) ([]resources.Domain, ccv3.Warnings, error)
	GetOrganizations(query ...ccv3.Query) ([]resources.Organization, ccv3.Warnings, error)
	GetPackage(guid string) (resources.Package, ccv3.Warnings, error)
	GetPackages(query ...ccv3.Query) ([]resources.Package, ccv3.Warnings, error)
	GetProcessInstances(processGUID string) ([]ccv3.ProcessInstance, ccv3.Warnings, error)
	GetServiceInstances(query ...ccv3.Query) ([]resources.ServiceInstance, ccv3.Warnings, error)
	GetSpaceIsolationSegment(spaceGUID string) (resources.Relationship, ccv3.Warnings, error)
	GetSpaces(query ...ccv3.Query) ([]resources.Space, ccv3.Warnings, error)
	GetStacks(query ...ccv3.Query) ([]resources.Stack, ccv3.Warnings, error)
	PollJob(jobURL ccv3.JobURL) (ccv3.Warnings, error)
	ResourceMatch(resources []ccv3.Resource) ([]ccv3.Resource, ccv3.Warnings, error)
	SetApplicationDroplet(appGUID string, dropletGUID string) (resources.Relationship, ccv3.Warnings, error)
	SharePrivateDomainToOrgs(domainGuid string, sharedOrgs ccv3.SharedOrgs) (ccv3.Warnings, error)
	ShareServiceInstanceToSpaces(serviceInstanceGUID string, spaceGUIDs []string) (resources.RelationshipList, ccv3.Warnings, error)
	UnsharePrivateDomainFromOrg(domainGUID string, sharedOrgGUID string) (ccv3.Warnings, error)
	UpdateApplication(app resources.Application) (resources.Application, ccv3.Warnings, error)
	UpdateApplicationApplyManifest(appGUID string, rawManifest []byte) (ccv3.JobURL, ccv3.Warnings, error)
	UpdateApplicationEnvironmentVariables(appGUID string, envVars resources.EnvironmentVariables) (resources.EnvironmentVariables, ccv3.Warnings, error)
	UpdateApplicationRestart(appGUID string) (resources.Application, ccv3.Warnings, error)
	UpdateApplicationStart(appGUID string) (resources.Application, ccv3.Warnings, error)
	UpdateApplicationStop(appGUID string) (resources.Application, ccv3.Warnings, error)
	UpdateBuildpack(buildpack resources.Buildpack) (resources.Buildpack, ccv3.Warnings, error)
	UpdateFeatureFlag(flag resources.FeatureFlag) (resources.FeatureFlag, ccv3.Warnings, error)
	UpdateOrganization(org resources.Organization) (resources.Organization, ccv3.Warnings, error)
	UpdateOrganizationDefaultIsolationSegmentRelationship(orgGUID string, isolationSegmentGUID string) (resources.Relationship, ccv3.Warnings, error)
	UpdateProcess(process resources.Process) (resources.Process, ccv3.Warnings, error)
	UpdateSpaceApplyManifest(spaceGUID string, rawManifest []byte, query ...ccv3.Query) (ccv3.JobURL, ccv3.Warnings, error)
	UpdateSpaceIsolationSegmentRelationship(spaceGUID string, isolationSegmentGUID string) (resources.Relationship, ccv3.Warnings, error)
	UpdateTaskCancel(taskGUID string) (resources.Task, ccv3.Warnings, error)
	UploadBitsPackage(pkg resources.Package, matchedResources []ccv3.Resource, newResources io.Reader, newResourcesLength int64) (resources.Package, ccv3.Warnings, error)
	UploadBuildpack(buildpackGUID string, buildpackPath string, buildpack io.Reader, buildpackLength int64) (ccv3.JobURL, ccv3.Warnings, error)
	UploadPackage(pkg resources.Package, zipFilepath string) (resources.Package, ccv3.Warnings, error)
}
