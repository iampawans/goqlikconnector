package registry

import (
	"github.com/iampawans/goqlikconnector/connectors"
)

// ConnectorRegistry keeps track of available connectors and their factories.
type ConnectorRegistry struct {
	factories map[string]connectors.ConnectorFactory
}

// RegisterConnectorFactory registers a connector factory with the registry.
func (r *ConnectorRegistry) RegisterConnectorFactory(name string, factory connectors.ConnectorFactory) {
	r.factories[name] = factory
}

// GetConnectorFactory returns the connector factory for a given connector type.
func (r *ConnectorRegistry) GetConnectorFactory(name string) connectors.ConnectorFactory {
	return r.factories[name]
}
