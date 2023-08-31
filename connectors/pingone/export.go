package pingone

import (
	"context"

	"github.com/pingidentity/pingctl/connector"
	"github.com/pingidentity/pingctl/internal/logger"
)

var (
	_ connector.CmdConnectorWithExport = &pingOneCmdConnector{}
)

func (p *pingOneCmdConnector) Export(ctx context.Context, opts connector.GenerateHCLOpts) error {

	// generate := pingtf.NewPingOneTerraformProvider(p.apiClient)
	// fmt.Println(generate.GenerateImportBlocks(ctx))

	l := logger.Get()

	l.Debug().Str("output directory", opts.OutputDirectoryPath).Msg("Run export")

	return nil
}
