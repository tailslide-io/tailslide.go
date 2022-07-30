package tailslide

import (
	flagManager "github.com/tailslide-io/tailslide/lib/flagmanager"
	toggler "github.com/tailslide-io/tailslide/lib/toggler"
	tailslideTypes "github.com/tailslide-io/tailslide/lib/types"
)

type Tailslide struct {
	flagManager.FlagManager
	tailslideTypes.FlagManagerConfig
	toggler.TogglerConfig
}
