package ccv3_test

import (
	"net/http"

	"code.cloudfoundry.org/cli/resources"
	"code.cloudfoundry.org/cli/types"

	"code.cloudfoundry.org/cli/api/cloudcontroller/ccerror"
	. "code.cloudfoundry.org/cli/api/cloudcontroller/ccv3"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/constant"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/ghttp"
)

var _ = Describe("Sidecars", func() {
	var client *Client

	BeforeEach(func() {
		client, _ = NewTestClient()
	})

	Describe("Patch & Delete", func() {
		var (
			executeErr error
		)

		When("updating a sidecar", func() {
			It("supplies all needed parameters", func() {
				server.AppendHandlers(
					CombineHandlers(
						VerifyRequest(http.MethodPatch, "/v3/sidecars/some-sidecar-guid"),
						RespondWith(http.StatusOK, `{
							"guid": "some-sidecar-guid",
							"name": "auth-sidecar",
							"command": "bundle exec rackup",
							"process_types": ["web", "worker"],
							"memory_in_mb": 10,
							"origin": "user",
							"relationships": {
							  "app": {
								"data": {
								  "guid": "some-app-guid"
								}
							  }
							},
							"created_at": "2017-02-01T01:33:58Z",
							"updated_at": "2017-02-01T01:33:58Z"
						  }
						  `),
					),
				)
				var memoryInMb = 10
				var sidecarRequest = resources.Sidecar{
					GUID: "some-sidecar-guid",
					Name: "auth-sidecar",
					Command: types.FilteredString{
						IsSet: true,
						Value: "bundle exec rackup;",
					},
					MemoryInMB: &memoryInMb,
				}
				_, _, executeErr = client.UpdateSidecar(sidecarRequest)
				Expect(executeErr).NotTo(HaveOccurred())
			})
		})
		When("deleting a sidecar", func() {
			It("supplies all needed parameters", func() {
				server.AppendHandlers(
					CombineHandlers(
						VerifyRequest(http.MethodDelete, "/v3/sidecars/some-sidecar-guid"),
						RespondWith(http.StatusNoContent, ""),
					),
				)
				_, _, executeErr := client.DeleteSidecar("some-sidecar-guid")
				Expect(executeErr).NotTo(HaveOccurred())
			})
		})
	})
	Describe("Post", func() {
		var (
			sidecar    resources.Sidecar
			executeErr error
		)

		When("creating a sidecar", func() {
			It("supplies all needed parameters", func() {
				server.AppendHandlers(
					CombineHandlers(
						VerifyRequest(http.MethodPost, "/v3/apps/some-app-guid/sidecars"),
						RespondWith(http.StatusOK, `{
							"guid": "some-sidecar-guid",
							"name": "auth-sidecar",
							"command": "bundle exec rackup",
							"process_types": ["web", "worker"],
							"memory_in_mb": 10,
							"origin": "user",
							"relationships": {
							  "app": {
								"data": {
								  "guid": "some-app-guid"
								}
							  }
							},
							"created_at": "2017-02-01T01:33:58Z",
							"updated_at": "2017-02-01T01:33:58Z"
						  }
						  `),
					),
				)
				var memoryInMb = 10
				var sidecarRequest = resources.Sidecar{
					Name: "auth-sidecar",
					Command: types.FilteredString{
						IsSet: true,
						Value: "bundle exec rackup;",
					},
					MemoryInMB: &memoryInMb,
				}
				sidecar, _, executeErr = client.CreateApplicationSidecar("some-app-guid", sidecarRequest)
				Expect(executeErr).NotTo(HaveOccurred())

				Expect(sidecar.GUID).To(Equal("some-sidecar-guid"))
				Expect(sidecar.Name).To(Equal("auth-sidecar"))
				Expect(sidecar.Command.Value).To(Equal("bundle exec rackup"))
				Expect(sidecar.ProcessTypes).To(Equal([]string{"web", "worker"}))
				Expect(*sidecar.MemoryInMB).To(Equal(10))
				Expect(*sidecar.Origin).To(Equal("user"))
				Expect(sidecar.Relationships[constant.RelationshipTypeApplication].GUID).To(Equal("some-app-guid"))
			})
			It("misssed required parameter", func() {
				server.AppendHandlers(
					CombineHandlers(
						VerifyRequest(http.MethodPost, "/v3/apps/some-app-guid/sidecars"),
						RespondWith(http.StatusUnprocessableEntity, `{"errors":[{"detail":"Name can't be blank, Name must be a string, Process types must be an array, Process types must have at least 1 process_type","title":"CF-UnprocessableEntity","code":10008}]}`),
					),
				)
				var memoryInMb = 10
				var sidecarRequest = resources.Sidecar{
					Command: types.FilteredString{
						IsSet: true,
						Value: "bundle exec rackup;",
					},
					MemoryInMB: &memoryInMb,
				}
				sidecar, _, executeErr = client.CreateApplicationSidecar("some-app-guid", sidecarRequest)
				Expect(executeErr).To(HaveOccurred())
				Expect(executeErr.Error()).To(ContainSubstring("Name can't be blank, Name must be a string"))
				Expect(executeErr.Error()).To(ContainSubstring("Process types must be an array"))

			})
		})
	})
	Describe("Get", func() {
		var (
			sidecar resources.Sidecar

			warnings   Warnings
			executeErr error
		)

		When("sidecar exist", func() {
			It("finds a sidecar which exists", func() {
				server.AppendHandlers(
					CombineHandlers(
						VerifyRequest(http.MethodGet, "/v3/sidecars/some-sidecar-guid"),
						RespondWith(http.StatusOK, `{
							"guid": "some-sidecar-guid",
							"name": "auth-sidecar",
							"command": "bundle exec rackup",
							"process_types": ["web", "worker"],
							"memory_in_mb": 300,
							"origin": "user",
							"relationships": {
							  "app": {
								"data": {
								  "guid": "some-app-guid"
								}
							  }
							},
							"created_at": "2017-02-01T01:33:58Z",
							"updated_at": "2017-02-01T01:33:58Z"
						  }
						  `),
					),
				)
				sidecar, _, executeErr = client.GetSidecar("some-sidecar-guid")
				Expect(executeErr).NotTo(HaveOccurred())

				Expect(sidecar.GUID).To(Equal("some-sidecar-guid"))
				Expect(sidecar.Name).To(Equal("auth-sidecar"))
				Expect(sidecar.Command.Value).To(Equal("bundle exec rackup"))
				Expect(sidecar.ProcessTypes).To(Equal([]string{"web", "worker"}))
				Expect(*sidecar.MemoryInMB).To(Equal(300))
				Expect(*sidecar.Origin).To(Equal("user"))
				Expect(sidecar.Relationships[constant.RelationshipTypeApplication].GUID).To(Equal("some-app-guid"))
			})
			It("finds sidecars of an application", func() {
				server.AppendHandlers(
					CombineHandlers(
						VerifyRequest(http.MethodGet, "/v3/apps/some-app-guid/sidecars"),
						RespondWith(http.StatusOK, `{
							"pagination":
							{
								"total_results": 1,
								"total_pages": 1,
								"first":
								{
									"href": "https://api.cf.example.com/v3/apps/0047e113-8f65-4b67-9c3c-3095635bea2f/sidecars?page=1&per_page=50"
								},
								"last":
								{
									"href": "https://api.cf.example.com/v3/apps/0047e113-8f65-4b67-9c3c-3095635bea2f/sidecars?page=1&per_page=50"
								},
								"next": null,
								"previous": null
							},
							"resources":
							[
								{
									"guid": "69c263e8-2af0-4776-bf42-7099740a45e7",
									"name": "test",
									"command": "while true; do echo \"sidecar event\"; sleep 1; done;",
									"process_types":
									[
										"web"
									],
									"memory_in_mb": 10,
									"origin": "user",
									"relationships":
									{
										"app":
										{
											"data":
											{
												"guid": "0047e113-8f65-4b67-9c3c-3095635bea2f"
											}
										}
									},
									"created_at": "2024-05-10T21:45:58Z",
									"updated_at": "2024-05-10T21:45:58Z"
								},
								{
									"guid": "69c263e8-2af0-4776-bf42-7099740a45e8",
									"name": "test2",
									"command": "while true; do echo \"sidecar event2\"; sleep 1; done;",
									"process_types":
									[
										"web"
									],
									"memory_in_mb": 10,
									"origin": "user",
									"relationships":
									{
										"app":
										{
											"data":
											{
												"guid": "0047e113-8f65-4b67-9c3c-3095635bea2f"
											}
										}
									},
									"created_at": "2024-05-10T21:45:58Z",
									"updated_at": "2024-05-10T21:45:58Z"
								}
							]
						}
						  `),
					),
				)

				var sidecars []resources.Sidecar

				sidecars, _, executeErr = client.GetApplicationSidecars("some-app-guid")
				Expect(executeErr).NotTo(HaveOccurred())
				Expect(len(sidecars)).To(Equal(2))
				Expect(sidecars[0].Name).To(Equal("test"))
				Expect(sidecars[1].Name).To(Equal("test2"))
			})
			It("finds sidecars of a process", func() {
				server.AppendHandlers(
					CombineHandlers(
						VerifyRequest(http.MethodGet, "/v3/processes/some-process-guid/sidecars"),
						RespondWith(http.StatusOK, `{
							"pagination":
							{
								"total_results": 1,
								"total_pages": 1,
								"first":
								{
									"href": "https://api.cf.example.com/v3/processes/0047e113-8f65-4b67-9c3c-3095635bea2f/sidecars?page=1&per_page=50"
								},
								"last":
								{
									"href": "https://api.cf.example.com/v3/processes/0047e113-8f65-4b67-9c3c-3095635bea2f/sidecars?page=1&per_page=50"
								},
								"next": null,
								"previous": null
							},
							"resources":
							[
								{
									"guid": "69c263e8-2af0-4776-bf42-7099740a45e7",
									"name": "test",
									"command": "while true; do echo \"sidecar event\"; sleep 1; done;",
									"process_types":
									[
										"web"
									],
									"memory_in_mb": 10,
									"origin": "user",
									"relationships":
									{
										"app":
										{
											"data":
											{
												"guid": "0047e113-8f65-4b67-9c3c-3095635bea2f"
											}
										}
									},
									"created_at": "2024-05-10T21:45:58Z",
									"updated_at": "2024-05-10T21:45:58Z"
								},
								{
									"guid": "69c263e8-2af0-4776-bf42-7099740a45e8",
									"name": "test2",
									"command": "while true; do echo \"sidecar event2\"; sleep 1; done;",
									"process_types":
									[
										"web"
									],
									"memory_in_mb": 10,
									"origin": "user",
									"relationships":
									{
										"app":
										{
											"data":
											{
												"guid": "0047e113-8f65-4b67-9c3c-3095635bea2f"
											}
										}
									},
									"created_at": "2024-05-10T21:45:58Z",
									"updated_at": "2024-05-10T21:45:58Z"
								}
							]
						}
						  `),
					),
				)

				var sidecars []resources.Sidecar

				sidecars, _, executeErr = client.GetProcessSidecars("some-process-guid")
				Expect(executeErr).NotTo(HaveOccurred())
				Expect(len(sidecars)).To(Equal(2))
				Expect(sidecars[0].Name).To(Equal("test"))
				Expect(sidecars[1].Name).To(Equal("test2"))
			})

			It("cannot find a sidecar", func() {
				server.AppendHandlers(
					CombineHandlers(
						VerifyRequest(http.MethodGet, "/v3/sidecars/some-missing-sidecar-guid"),
						RespondWith(http.StatusNotFound, `{
							"errors": [
							  {
								"code": 10010,
								"detail": "Sidecar not found",
								"title": "CF-ResourceNotFound"
							  }
							]
						  }`),
					),
				)
				sidecar, _, executeErr = client.GetSidecar("some-missing-sidecar-guid")
				Expect(executeErr).To(HaveOccurred())
				Expect(sidecar).To(Equal(resources.Sidecar{}))
			})
		})

		When("the cloud controller returns errors and warnings", func() {
			It("returns the error and all warnings", func() {
				server.AppendHandlers(
					CombineHandlers(
						VerifyRequest(http.MethodGet, "/v3/sidecars/some-missing-sidecar-guid"),
						RespondWith(http.StatusTeapot, `{
								"errors": [
								  {
									"code": 10008,
									"detail": "The request is semantically invalid: command presence",
									"title": "CF-UnprocessableEntity"
								  },
								  {
									"code": 10010,
									"detail": "Space not found",
									"title": "CF-SidecarNotFound"
								  }
								]
							  }`, http.Header{"X-Cf-Warnings": {"this is a warning"}}),
					),
				)
				sidecar, warnings, executeErr = client.GetSidecar("some-missing-sidecar-guid")
				Expect(executeErr).To(MatchError(ccerror.MultiError{
					ResponseCode: http.StatusTeapot,
					Errors: []ccerror.V3Error{
						{
							Code:   10008,
							Detail: "The request is semantically invalid: command presence",
							Title:  "CF-UnprocessableEntity",
						},
						{
							Code:   10010,
							Detail: "Space not found",
							Title:  "CF-SidecarNotFound",
						},
					},
				}))
				Expect(warnings).To(ConsistOf("this is a warning"))
			})
		})
	})
})
