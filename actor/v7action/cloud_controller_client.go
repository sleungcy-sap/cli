package v7action

import (
	"io"

	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3"
)

//go:generate counterfeiter . CloudControllerClient

// CloudControllerClient is the interface to the cloud controller V3 API.
type CloudControllerClient interface {
	AppSSHEndpoint() string
	AppSSHHostKeyFingerprint() string
	CheckRoute(domainGUID string, hostname string, path string) (bool, ccv3.Warnings, error)
	CloudControllerAPIVersion() string
	CancelDeployment(deploymentGUID string) (ccv3.Warnings, error)
	CreateApplication(app ccv3.Application) (ccv3.Application, ccv3.Warnings, error)
	CreateApplicationDeployment(appGUID string, dropletGUID string) (string, ccv3.Warnings, error)
	CreateApplicationProcessScale(appGUID string, process ccv3.Process) (ccv3.Process, ccv3.Warnings, error)
	CreateApplicationTask(appGUID string, task ccv3.Task) (ccv3.Task, ccv3.Warnings, error)
	CreateBuild(build ccv3.Build) (ccv3.Build, ccv3.Warnings, error)
	CreateBuildpack(bp ccv3.Buildpack) (ccv3.Buildpack, ccv3.Warnings, error)
	CreateDomain(domain ccv3.Domain) (ccv3.Domain, ccv3.Warnings, error)
	CreateDroplet(appGUID string) (ccv3.Droplet, ccv3.Warnings, error)
	CreateIsolationSegment(isolationSegment ccv3.IsolationSegment) (ccv3.IsolationSegment, ccv3.Warnings, error)
	CreatePackage(pkg ccv3.Package) (ccv3.Package, ccv3.Warnings, error)
	CreateRoute(route ccv3.Route) (ccv3.Route, ccv3.Warnings, error)
	CreateServiceBroker(name, username, password, url, spaceGUID string) (ccv3.Warnings, error)
	CreateSpace(space ccv3.Space) (ccv3.Space, ccv3.Warnings, error)
	DeleteApplication(guid string) (ccv3.JobURL, ccv3.Warnings, error)
	DeleteApplicationProcessInstance(appGUID string, processType string, instanceIndex int) (ccv3.Warnings, error)
	DeleteBuildpack(buildpackGUID string) (ccv3.JobURL, ccv3.Warnings, error)
	DeleteDomain(domainGUID string) (ccv3.JobURL, ccv3.Warnings, error)
	DeleteIsolationSegment(guid string) (ccv3.Warnings, error)
	DeleteIsolationSegmentOrganization(isolationSegmentGUID string, organizationGUID string) (ccv3.Warnings, error)
	DeleteOrphanedRoutes(spaceGUID string) (ccv3.JobURL, ccv3.Warnings, error)
	DeleteRoute(routeGUID string) (ccv3.JobURL, ccv3.Warnings, error)
	DeleteServiceInstanceRelationshipsSharedSpace(serviceInstanceGUID string, sharedToSpaceGUID string) (ccv3.Warnings, error)
	DeleteSpace(guid string) (ccv3.JobURL, ccv3.Warnings, error)
	EntitleIsolationSegmentToOrganizations(isoGUID string, orgGUIDs []string) (ccv3.RelationshipList, ccv3.Warnings, error)
	GetApplicationDropletCurrent(appGUID string) (ccv3.Droplet, ccv3.Warnings, error)
	GetApplicationEnvironment(appGUID string) (ccv3.Environment, ccv3.Warnings, error)
	GetApplicationManifest(appGUID string) ([]byte, ccv3.Warnings, error)
	GetApplicationProcessByType(appGUID string, processType string) (ccv3.Process, ccv3.Warnings, error)
	GetApplicationProcesses(appGUID string) ([]ccv3.Process, ccv3.Warnings, error)
	GetApplicationRoutes(appGUID string) ([]ccv3.Route, ccv3.Warnings, error)
	GetApplicationTasks(appGUID string, query ...ccv3.Query) ([]ccv3.Task, ccv3.Warnings, error)
	GetApplications(query ...ccv3.Query) ([]ccv3.Application, ccv3.Warnings, error)
	GetBuild(guid string) (ccv3.Build, ccv3.Warnings, error)
	GetBuildpacks(query ...ccv3.Query) ([]ccv3.Buildpack, ccv3.Warnings, error)
	GetDefaultDomain(orgGuid string) (ccv3.Domain, ccv3.Warnings, error)
	GetDeployment(guid string) (ccv3.Deployment, ccv3.Warnings, error)
	GetDeployments(query ...ccv3.Query) ([]ccv3.Deployment, ccv3.Warnings, error)
	GetDomain(GUID string) (ccv3.Domain, ccv3.Warnings, error)
	GetDomains(query ...ccv3.Query) ([]ccv3.Domain, ccv3.Warnings, error)
	GetDroplet(guid string) (ccv3.Droplet, ccv3.Warnings, error)
	GetDroplets(query ...ccv3.Query) ([]ccv3.Droplet, ccv3.Warnings, error)
	GetFeatureFlag(featureFlagName string) (ccv3.FeatureFlag, ccv3.Warnings, error)
	GetFeatureFlags() ([]ccv3.FeatureFlag, ccv3.Warnings, error)
	GetIsolationSegment(guid string) (ccv3.IsolationSegment, ccv3.Warnings, error)
	GetIsolationSegmentOrganizations(isolationSegmentGUID string) ([]ccv3.Organization, ccv3.Warnings, error)
	GetIsolationSegments(query ...ccv3.Query) ([]ccv3.IsolationSegment, ccv3.Warnings, error)
	GetNewApplicationProcesses(appGUID string, deploymentGUID string) ([]ccv3.Process, ccv3.Warnings, error)
	GetOrganizationDefaultIsolationSegment(orgGUID string) (ccv3.Relationship, ccv3.Warnings, error)
	GetOrganizationDomains(orgGUID string, query ...ccv3.Query) ([]ccv3.Domain, ccv3.Warnings, error)
	GetOrganizations(query ...ccv3.Query) ([]ccv3.Organization, ccv3.Warnings, error)
	GetPackage(guid string) (ccv3.Package, ccv3.Warnings, error)
	GetPackages(query ...ccv3.Query) ([]ccv3.Package, ccv3.Warnings, error)
	GetProcess(processGUID string) (ccv3.Process, ccv3.Warnings, error)
	GetProcessInstances(processGUID string) ([]ccv3.ProcessInstance, ccv3.Warnings, error)
	GetRouteDestinations(routeGUID string) ([]ccv3.RouteDestination, ccv3.Warnings, error)
	GetRoutes(query ...ccv3.Query) ([]ccv3.Route, ccv3.Warnings, error)
	GetServiceBrokers() ([]ccv3.ServiceBroker, ccv3.Warnings, error)
	GetServiceInstances(query ...ccv3.Query) ([]ccv3.ServiceInstance, ccv3.Warnings, error)
	GetSpaceIsolationSegment(spaceGUID string) (ccv3.Relationship, ccv3.Warnings, error)
	GetSpaces(query ...ccv3.Query) ([]ccv3.Space, ccv3.Warnings, error)
	GetStacks(query ...ccv3.Query) ([]ccv3.Stack, ccv3.Warnings, error)
	MapRoute(routeGUID string, appGUID string) (ccv3.Warnings, error)
	PollJob(jobURL ccv3.JobURL) (ccv3.Warnings, error)
	ResourceMatch(resources []ccv3.Resource) ([]ccv3.Resource, ccv3.Warnings, error)
	SetApplicationDroplet(appGUID string, dropletGUID string) (ccv3.Relationship, ccv3.Warnings, error)
	SharePrivateDomainToOrgs(domainGuid string, sharedOrgs ccv3.SharedOrgs) (ccv3.Warnings, error)
	ShareServiceInstanceToSpaces(serviceInstanceGUID string, spaceGUIDs []string) (ccv3.RelationshipList, ccv3.Warnings, error)
	UnmapRoute(routeGUID string, destinationGUID string) (ccv3.Warnings, error)
	UnsharePrivateDomainFromOrg(domainGUID string, sharedOrgGUID string) (ccv3.Warnings, error)
	UpdateApplication(app ccv3.Application) (ccv3.Application, ccv3.Warnings, error)
	UpdateApplicationApplyManifest(appGUID string, rawManifest []byte) (ccv3.JobURL, ccv3.Warnings, error)
	UpdateApplicationEnvironmentVariables(appGUID string, envVars ccv3.EnvironmentVariables) (ccv3.EnvironmentVariables, ccv3.Warnings, error)
	UpdateApplicationRestart(appGUID string) (ccv3.Application, ccv3.Warnings, error)
	UpdateApplicationStart(appGUID string) (ccv3.Application, ccv3.Warnings, error)
	UpdateApplicationStop(appGUID string) (ccv3.Application, ccv3.Warnings, error)
	UpdateBuildpack(buildpack ccv3.Buildpack) (ccv3.Buildpack, ccv3.Warnings, error)
	UpdateFeatureFlag(flag ccv3.FeatureFlag) (ccv3.FeatureFlag, ccv3.Warnings, error)
	UpdateOrganization(org ccv3.Organization) (ccv3.Organization, ccv3.Warnings, error)
	UpdateOrganizationDefaultIsolationSegmentRelationship(orgGUID string, isolationSegmentGUID string) (ccv3.Relationship, ccv3.Warnings, error)
	UpdateProcess(process ccv3.Process) (ccv3.Process, ccv3.Warnings, error)
	UpdateResourceMetadata(resource string, resourceGUID string, metadata ccv3.Metadata) (ccv3.ResourceMetadata, ccv3.Warnings, error)
	UpdateSpace(space ccv3.Space) (ccv3.Space, ccv3.Warnings, error)
	UpdateSpaceApplyManifest(spaceGUID string, rawManifest []byte, query ...ccv3.Query) (ccv3.JobURL, ccv3.Warnings, error)
	UpdateSpaceIsolationSegmentRelationship(spaceGUID string, isolationSegmentGUID string) (ccv3.Relationship, ccv3.Warnings, error)
	UpdateTaskCancel(taskGUID string) (ccv3.Task, ccv3.Warnings, error)
	UploadBitsPackage(pkg ccv3.Package, matchedResources []ccv3.Resource, newResources io.Reader, newResourcesLength int64) (ccv3.Package, ccv3.Warnings, error)
	UploadBuildpack(buildpackGUID string, buildpackPath string, buildpack io.Reader, buildpackLength int64) (ccv3.JobURL, ccv3.Warnings, error)
	UploadDropletBits(dropletGUID string, dropletPath string, droplet io.Reader, dropletLength int64) (ccv3.JobURL, ccv3.Warnings, error)
	UploadPackage(pkg ccv3.Package, zipFilepath string) (ccv3.Package, ccv3.Warnings, error)
}
