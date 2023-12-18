package main

import (
	"fmt"

	"github.com/iampawans/goqlikconnector/config"
	"github.com/iampawans/goqlikconnector/connectors"
	"github.com/iampawans/goqlikconnector/registry"
)

func main() {
	// Initialize the connector registry
	registry := &registry.ConnectorRegistry{
		factories: make(map[string]connectors.ConnectorFactory),
	}

	// Register API connector factory
	registry.RegisterConnectorFactory("API", &connectors.APIConnectorFactory{})

	// Create a configuration for an API connector
	apiConfig := config.ConnectorConfig{
		Type: "API",
		Name: "GithubAPI",
		ConfigParams: map[string]interface{}{
			"Endpoint": "https://github.com/iampawans/golang",
			// Additional API-specific parameters
		},
	}

	// Create an instance of the API connector
	apiFactory := registry.GetConnectorFactory(apiConfig.Type)
	apiConnector, err := apiFactory.CreateConnector(apiConfig)
	if err != nil {
		fmt.Println("Error creating API connector:", err)
		return
	}

	// Connect to the API
	err = apiConnector.Connect()
	if err != nil {
		fmt.Println("Error connecting to API:", err)
		return
	}

	// Fetch data from the API
	data, err := apiConnector.FetchData()
	if err != nil {
		fmt.Println("Error fetching data from API:", err)
		return
	}

	fmt.Println("Fetched data:", data)

	// Disconnect from the API
	err = apiConnector.Disconnect()
	if err != nil {
		fmt.Println("Error disconnecting from API:", err)
		return
	}
}
