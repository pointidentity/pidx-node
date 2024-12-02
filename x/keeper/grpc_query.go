package keeper

import (
	"github.com/pointidentity/pidx-node/x/types"
)

var _ types.QueryServer = Keeper{}
