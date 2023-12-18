package connectors

// Connector interface defines the methods that each connector must implement.
type Connector interface {
	Connect() error
	Disconnect() error
	FetchData() ([]map[string]interface{}, error)
	// Additional methods for configuration, metadata retrieval, etc.
}
