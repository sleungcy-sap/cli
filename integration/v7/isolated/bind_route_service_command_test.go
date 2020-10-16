package isolated

import (
	"code.cloudfoundry.org/cli/integration/helpers"
	"code.cloudfoundry.org/cli/resources"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("bind-route-service command", func() {
	const command = "v3-bind-route-service"

	Describe("help", func() {
		matchHelpMessage := SatisfyAll(
			Say(`NAME:\n`),
			Say(`\s+v3-bind-route-service - Bind a service instance to an HTTP route\n`),
			Say(`\n`),
			Say(`USAGE:\n`),
			Say(`\s+cf bind-route-service DOMAIN \[--hostname HOSTNAME\] \[--path PATH\] SERVICE_INSTANCE \[-c PARAMETERS_AS_JSON\]\n`),
			Say(`\n`),
			Say(`EXAMPLES:\n`),
			Say(`\s+cf bind-route-service example.com --hostname myapp --path foo myratelimiter\n`),
			Say(`\s+cf bind-route-service example.com myratelimiter -c file.json\n`),
			Say(`\s+cf bind-route-service example.com myratelimiter -c '{"valid":"json"}'\n`),
			Say(`\n`),
			Say(`\s+In Windows PowerShell use double-quoted, escaped JSON: "\{\\"valid\\":\\"json\\"\}"\n`),
			Say(`\s+In Windows Command Line use single-quoted, escaped JSON: '\{\\"valid\\":\\"json\\"\}'\n`),
			Say(`\n`),
			Say(`ALIAS:\n`),
			Say(`\s+brs\n`),
			Say(`\n`),
			Say(`OPTIONS:\n`),
			Say(`\s+-c\s+Valid JSON object containing service-specific configuration parameters, provided inline or in a file. For a list of supported configuration parameters, see documentation for the particular service offering.\n`),
			Say(`\s+--hostname, -n\s+Hostname used in combination with DOMAIN to specify the route to bind\n`),
			Say(`\s+--path\s+Path used in combination with HOSTNAME and DOMAIN to specify the route to bind\n`),
			Say(`\s+--wait, -w\s+Wait for the bind operation to complete\n`),
			Say(`\n`),
			Say(`SEE ALSO:\n`),
			Say(`\s+routes, services\n`),
		)

		When("the -h flag is specified", func() {
			It("succeeds and prints help", func() {
				session := helpers.CF(command, "-h")
				Eventually(session).Should(Exit(0))
				Expect(session.Out).To(matchHelpMessage)
			})
		})

		When("the --help flag is specified", func() {
			It("succeeds and prints help", func() {
				session := helpers.CF(command, "--help")
				Eventually(session).Should(Exit(0))
				Expect(session.Out).To(matchHelpMessage)
			})
		})

		When("no arguments are provided", func() {
			It("displays a warning, the help text, and exits 1", func() {
				session := helpers.CF(command)
				Eventually(session).Should(Exit(1))
				Expect(session.Err).To(Say("Incorrect Usage: the required arguments `DOMAIN` and `SERVICE_INSTANCE` were not provided"))
				Expect(session.Out).To(matchHelpMessage)
			})
		})

		When("unknown flag is passed", func() {
			It("displays a warning, the help text, and exits 1", func() {
				session := helpers.CF(command, "-u")
				Eventually(session).Should(Exit(1))
				Expect(session.Err).To(Say("Incorrect Usage: unknown flag `u"))
				Expect(session.Out).To(matchHelpMessage)
			})
		})
	})

	When("the environment is not setup correctly", func() {
		It("fails with the appropriate errors", func() {
			helpers.CheckEnvironmentTargetedCorrectly(true, true, ReadOnlyOrg, command, "foo", "bar")
		})
	})

	Context("user-provided route service", func() {
		var (
			orgName             string
			spaceName           string
			routeServiceURL     string
			serviceInstanceName string
			domain              string
			hostname            string
			path                string
			username            string
		)

		BeforeEach(func() {
			orgName = helpers.NewOrgName()
			spaceName = helpers.NewSpaceName()
			helpers.SetupCF(orgName, spaceName)

			routeServiceURL = helpers.RandomURL()
			serviceInstanceName = helpers.NewServiceInstanceName()
			Eventually(helpers.CF("cups", serviceInstanceName, "-r", routeServiceURL)).Should(Exit(0))

			domain = helpers.DefaultSharedDomain()
			hostname = helpers.NewHostName()
			path = helpers.PrefixedRandomName("path")
			Eventually(helpers.CF("create-route", domain, "--hostname", hostname, "--path", path)).Should(Exit(0))

			username, _ = helpers.GetCredentials()
		})

		AfterEach(func() {
			helpers.QuickDeleteOrg(orgName)
		})

		It("creates a route binding", func() {
			session := helpers.CF(command, domain, "--hostname", hostname, "--path", path, serviceInstanceName)
			Eventually(session).Should(Exit(0))

			Expect(session.Out).To(SatisfyAll(
				Say(`Binding route %s.%s/%s to service instance %s in org %s / space %s as %s\.\.\.\n`, hostname, domain, path, serviceInstanceName, orgName, spaceName, username),
				Say(`\n`),
				Say(`Route binding created\.\n`),
				Say(`OK\n`),
			))

			Expect(string(session.Err.Contents())).To(BeEmpty())

			var receiver struct {
				Resources []resources.RouteBinding `json:"resources"`
			}
			helpers.Curl(&receiver, "/v3/service_route_bindings?service_instance_names=%s", serviceInstanceName)
			Expect(receiver.Resources).To(HaveLen(1))
		})
	})
})
