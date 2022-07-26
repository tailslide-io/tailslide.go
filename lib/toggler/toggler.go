package toggler

import (
	"crypto/md5"
	"errors"
	"fmt"
	"strconv"
	"strings"

	tailslideTypes "github.com/tailslide-io/tailslide/lib/types"
)

type TogglerConfig struct {
	FlagName        string
	UserContext     string
	GetFlags        tailslideTypes.GetFlags
	EmitRedisSignal tailslideTypes.EmitRedisSignal
	// FeatureCB
	// DefaultCB
	// ErrorCondition

}

type Toggler struct {
	flagName        string
	flagId          int
	appId           int
	userContext     string
	getFlags        tailslideTypes.GetFlags
	emitRedisSignal tailslideTypes.EmitRedisSignal
	// FeatureCB
	// DefaultCB
	// ErrorCondition
}

func New(config TogglerConfig) (*Toggler, error) {
	toggler := Toggler{
		flagName:        config.FlagName,
		userContext:     config.UserContext,
		getFlags:        config.GetFlags,
		emitRedisSignal: config.EmitRedisSignal,
	}
	err := toggler.setFlagIdAndAppId()
	if err != nil {
		return nil, err
	}
	return &toggler, nil
}

func (toggler *Toggler) IsFlagActive() bool {
	flag, err := toggler.getMatchingFlag()
	if err != nil {
		return false
	}
	return flag.IsActive && (toggler.isUserWhiteListed(flag) || toggler.validateUserRollout(flag))
}

func (toggler *Toggler) EmitSuccess() {
	if toggler.flagId == 0 {
		return
	}
	toggler.emitRedisSignal(toggler.flagId, toggler.appId, "success")
}

func (toggler *Toggler) EmitFailiure() {
	if toggler.flagId == 0 {
		return
	}
	toggler.emitRedisSignal(toggler.flagId, toggler.appId, "failure")
}

func (toggler *Toggler) setFlagIdAndAppId() error {
	matchingFlag, err := toggler.getMatchingFlag()
	if err != nil {
		return err
	}
	toggler.flagId = matchingFlag.AppId
	toggler.appId = matchingFlag.AppId
	return nil
}

func (toggler *Toggler) getMatchingFlag() (tailslideTypes.Flag, error) {
	flags := toggler.getFlags()
	for _, flag := range flags {
		if flag.Title == toggler.flagName {
			return flag, nil
		}
	}
	return tailslideTypes.Flag{}, errors.New(fmt.Sprintf("Cannot find flag with name: %s\n", toggler.flagName))
}

func (toggler *Toggler) isUserWhiteListed(flag tailslideTypes.Flag) bool {
	for _, user := range strings.Split(flag.WhiteListedUsers, ",") {
		if user == toggler.userContext {
			return true
		}
	}
	return false
}

func (toggler *Toggler) validateUserRollout(flag tailslideTypes.Flag) bool {
	rollout := flag.RolloutPercentage / 100.0
	if toggler.circuitInRecovery(flag) {
		rollout = rollout * (flag.CircuitRecoveryPercentage / 100.0)
	}
	return toggler.IsUserInRollout(rollout)
}

func (toggler *Toggler) circuitInRecovery(flag tailslideTypes.Flag) bool {
	return flag.IsRecoverable && flag.CircuitStatus == "recovery"
}

func (toggler *Toggler) IsUserInRollout(rollout float32) bool {
	fmt.Println("Rollout is ", rollout, "Hashed user value is ", toggler.hashUserContext())
	return toggler.hashUserContext() <= rollout
}

func (toggler *Toggler) hashUserContext() float32 {
	hash := md5.Sum([]byte(toggler.userContext))
	hashString := fmt.Sprintf("%x", hash)
	value, _ := strconv.ParseInt(hashString, 16, 8)
	finaValue := float32(value%100) / 100.0
	return float32(finaValue)
}
