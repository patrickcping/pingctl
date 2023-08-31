package internal

import (
	"github.com/pingidentity/pingctl/connector"
	"github.com/pingidentity/pingctl/connectors/pingone"
)

func Connectors() []connector.ConnectorType {
	return []connector.ConnectorType{
		pingone.NewPingOneCmdConnector,
	}
}
