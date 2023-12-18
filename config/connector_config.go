package config

// ConnectorConfig stores configuration parameters for a connector.
type ConnectorConfig struct {
	Type         string                 // Type of the connector (API, Database, Streaming, etc.)
	Name         string                 // Name of the connector instance
	ConfigParams map[string]interface{} // Connector-specific configuration parameters
}
