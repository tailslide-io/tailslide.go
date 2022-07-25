package toggler

import (
	"errors"
	"fmt"

	tailslideTypes "github.com/tailslide-io/tailslide/lib/types"
)

type TogglerConfig struct {
	FlagName string
	UserContext string
	GetFlags tailslideTypes.GetFlags
	// emitRedisSignal
	// FeatureCB
	// DefaultCB
	// ErrorCondition

}

type Toggler struct {
	flagName string
	flagId int
	appId int
	userContext string
	getFlags tailslideTypes.GetFlags
	// emitRedisSignal
	// FeatureCB
	// DefaultCB
	// ErrorCondition
}

func New(config TogglerConfig) (*Toggler, error) {
	toggler := Toggler{
		flagName: config.FlagName,
		userContext: config.UserContext,
		getFlags: config.GetFlags,
	}
	err := toggler.setFlagIdAndAppId()
	if err != nil {
		return nil, err
	}
	return &toggler, nil
}


func (toggler *Toggler) IsFlagActive() bool {
	return true
}

// func (toggler *Toggler) EmitSuccess(){}
// func (toggler *Toggler) EmitFailiure(){}

func (toggler *Toggler) setFlagIdAndAppId() error {
	matchingFlag := toggler.getMatchingFlag()
	if matchingFlag == nil {
		return errors.New(fmt.Sprintf("Cannot find flag with name: %s\n", toggler.flagName))
	}
	return nil
}

func (toggler *Toggler) getMatchingFlag() *tailslideTypes.Flag{
	flags := toggler.getFlags()
	for _, flag := range flags {
		if flag.Title == toggler.flagName {
			return &flag
		}
	}
	return nil
}

func (toggler *Toggler) EmitFailiure(){}
