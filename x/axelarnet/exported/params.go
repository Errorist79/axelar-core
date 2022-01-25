package exported

import (
	"github.com/axelarnetwork/axelar-core/x/nexus/exported"
	tss "github.com/axelarnetwork/axelar-core/x/tss/exported"
)

var (
	// Uaxl is the native asset on Axelarnet
	Uaxl = "uaxl"
	// Axelarnet defines properties of the Axelar chain
	Axelarnet = exported.Chain{
		Name:                  "Axelarnet",
		SupportsForeignAssets: true,
		KeyType:               tss.None,
		Module:                "axelarnet", // cannot use constant due to import cycle
	}
)
