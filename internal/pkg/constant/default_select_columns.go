package constant

import (
	"github.com/rodericusifo/echo-template/pkg/types"
)

type Selects []types.SelectOperation

var (
	DEFAULT_SELECT_COLUMNS      = Selects([]types.SelectOperation{})
	DEFAULT_JOIN_SELECT_COLUMNS = Selects([]types.SelectOperation{})
)
