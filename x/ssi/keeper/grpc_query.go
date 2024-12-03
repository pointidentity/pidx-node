package keeper

import (
	"github.com/pointidentity/pidx-node/x/ssi/types"
)

var _ types.QueryServer = Keeper{}
