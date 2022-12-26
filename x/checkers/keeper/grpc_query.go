package keeper

import (
	"github.com/terry/checkers/x/checkers/types"
)

var _ types.QueryServer = Keeper{}
