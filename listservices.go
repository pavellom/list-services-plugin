package main

import (
	"code.cloudfoundry.org/cli/plugin"
	"fmt"
	"os"
	"strings"
	"./terminal"
	"./services"
)

type ListServices struct{}

func (c *ListServices) Run(cliConnection plugin.CliConnection, args []string) {
    if args[0] == "CLI-MESSAGE-UNINSTALL" {
        return
    }

    printServices(cliConnection, args)
}

func printServices(cliConnection plugin.CliConnection, args []string) {
    appName := extractAppName(cliConnection, args)
    validateLogin(cliConnection)
    validateTarget(cliConnection)
    extractServices(cliConnection, appName)
}

func extractAppName(cliConnection plugin.CliConnection, args []string) string {
    if len(args) < 2 {
        fmt.Println("Incorrect Usage: the required argument `APP_NAME` was not provided\n")
        help, err := cliConnection.CliCommand("help", "list-services")
        if (err != nil) {
            fmt.Print(help)
        }
        os.Exit(1)
    }
    return args[1]
}

func extractServices(cliConnection plugin.CliConnection, appName string) {
    app, err := cliConnection.GetApp(appName)
    if err != nil {
        terminal.Fail(fmt.Sprintf("App %s not found", appName))
    }

    servicesJsonResponse, err := cliConnection.CliCommandWithoutTerminalOutput("curl", fmt.Sprintf("/v3/service_bindings?app_guids=%s", app.Guid))

    if err != nil {
        terminal.Fail(fmt.Sprintf("Can't retreive services for %s", appName))
    }

    servicesResponse := services.ParseResponse(strings.Join(servicesJsonResponse, ""))
    if err != nil {
        terminal.Fail(fmt.Sprintf("Can't output services for %s", appName))
    }

    // map with results for table
    resources_map := make(map[string]string)

    populateTable(resources_map, servicesResponse.Resources)

    // Iterate over pages
    totalPages := servicesResponse.Pagination.TotalPages
    for i := 2; i <= totalPages; i++ {
        servicesJsonResponse, err := cliConnection.CliCommandWithoutTerminalOutput("curl", fmt.Sprintf("/v3/service_bindings?app_guids=%s&page=%d", app.Guid, i))
        if err != nil {
            terminal.Fail(fmt.Sprintf("Can't output services for %s", appName))
        }
        servicesResponse := services.ParseResponse(strings.Join(servicesJsonResponse, ""))
        populateTable(resources_map, servicesResponse.Resources)
    }


    table := terminal.Table([]string{"Service Name", "Service Instance URL"})
    for key := range resources_map {
        terminal.Add(table, key, resources_map[key])
    }
    terminal.PrintTable(table)
}

func populateTable(m map[string]string, resources []services.Resource) {
    for _, resource := range resources {
        m[resource.Data.Name] = resource.Links.Instance.Href
    }
}

func validateLogin(cliConnection plugin.CliConnection) {
    loggedIn, err := cliConnection.IsLoggedIn()

    if ( !loggedIn || err != nil) {
        terminal.Fail("Not logged in. Use 'cf login' or 'cf login --sso' to log in.")
    }
}

func validateTarget(cliConnection plugin.CliConnection) {
    hasOrg, err := cliConnection.HasOrganization()

    if ( !hasOrg || err != nil) {
        terminal.Fail("No org and space targeted, use 'cf target -o ORG -s SPACE' to target an org and space")
    }

    hasSpace, err := cliConnection.HasSpace()

    if ( !hasSpace || err != nil) {
        terminal.Fail("No space targeted, use 'cf target -s' to target a space.")
    }
}

func (c *ListServices) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "ListServices",
		Version: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 0,
		},
		MinCliVersion: plugin.VersionType{
			Major: 6,
			Minor: 7,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "list-services",
				HelpText: "List services which are bound to tha application",

				UsageDetails: plugin.Usage{
					Usage: "list-services APP_NAME - Output to Console the list of services which are bound to the application.",
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(ListServices))
}