package tailslide

import (
	flagManager "github.com/tailslide-io/tailslide.go/lib/flagmanager"
	toggler "github.com/tailslide-io/tailslide.go/lib/toggler"
	tailslideTypes "github.com/tailslide-io/tailslide.go/lib/types"
)

func NewFlagManager(config tailslideTypes.FlagManagerConfig) *flagManager.FlagManager {
	return flagManager.New(config)
}

type FlagManagerConfig = tailslideTypes.FlagManagerConfig
type TogglerConfig = toggler.TogglerConfig
