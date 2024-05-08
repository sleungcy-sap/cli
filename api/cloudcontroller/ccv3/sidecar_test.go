package ccv3_test

import (
	"net/http"

	"code.cloudfoundry.org/cli/resources"

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

	Describe("Get", func() {
		var (
			sidecar resources.Sidecar

			warnings   Warnings
			executeErr error
		)

		When("sidecar exist", func() {
			BeforeEach(func() {
				response1 := `{
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
				  `
				response2 := `{
					"errors": [
					  {
						"code": 10010,
						"detail": "Sidecar not found",
						"title": "CF-ResourceNotFound"
					  }
					]
				  }`

				server.AppendHandlers(
					CombineHandlers(
						VerifyRequest(http.MethodGet, "/v3/sidecars/some-sidecar-guid"),
						RespondWith(http.StatusOK, response1),
					),
				)
				server.AppendHandlers(
					CombineHandlers(
						VerifyRequest(http.MethodGet, "/v3/sidecars/some-missing-sidecar-guid"),
						RespondWith(http.StatusNotFound, response2),
					),
				)
			})

			It("finds a sidecar which exists", func() {

				sidecar, _, executeErr = client.GetSidecar("some-sidecar-guid")
				Expect(executeErr).NotTo(HaveOccurred())

				Expect(sidecar.GUID).To(Equal("some-sidecar-guid"))
				Expect(sidecar.Name).To(Equal("auth-sidecar"))
				Expect(sidecar.Command).To(Equal("bundle exec rackup"))
				Expect(sidecar.ProcessTypes).To(Equal([]string{"web", "worker"}))
				Expect(sidecar.MemoryInMB).To(Equal(300))
				Expect(sidecar.Origin).To(Equal("user"))
				Expect(sidecar.Relationships[constant.RelationshipTypeApplication].GUID).To(Equal("some-app-guid"))
			})
			It("cannot find a sidecar", func() {

				sidecar, _, executeErr = client.GetSidecar("some-missing-sidecar-guid")
				Expect(executeErr).To(HaveOccurred())
				Expect(sidecar).To(Equal(resources.Sidecar{}))
			})
		})

		When("the cloud controller returns errors and warnings", func() {
			BeforeEach(func() {
				response := `{
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
}`
				server.AppendHandlers(
					CombineHandlers(
						VerifyRequest(http.MethodGet, "/v3/sidecars/some-sidecar-guid"),
						RespondWith(http.StatusTeapot, response, http.Header{"X-Cf-Warnings": {"this is a warning"}}),
					),
				)
			})

			It("returns the error and all warnings", func() {
				sidecar, _, executeErr = client.GetSidecar("some-missing-sidecar-guid")
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
