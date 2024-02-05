package keeper

import (
	"errors"
	"fmt"

	state "github.com/upshot-tech/protocol-state-machine-module"
)

var ErrIntegerUnderflowDelegator = errors.New(state.Err_ErrIntegerUnderflowDelegator.String())
var ErrIntegerUnderflowBonds = errors.New(state.Err_ErrIntegerUnderflowBonds.String())
var ErrIntegerUnderflowTarget = errors.New(state.Err_ErrIntegerUnderflowTarget.String())
var ErrIntegerUnderflowTopicStake = errors.New(state.Err_ErrIntegerUnderflowTopicStake.String())
var ErrIntegerUnderflowTotalStake = errors.New(state.Err_ErrIntegerUnderflowTotalStake.String())
var ErrIterationLengthDoesNotMatch = errors.New(state.Err_ErrIterationLengthDoesNotMatch.String())
var ErrInvalidTopicId = fmt.Errorf(state.Err_ErrInvalidTopicId.String())
var ErrReputerAlreadyRegistered = fmt.Errorf(state.Err_ErrReputerAlreadyRegistered.String())
var ErrWorkerAlreadyRegistered = fmt.Errorf(state.Err_ErrWorkerAlreadyRegistered.String())
var ErrInsufficientStakeToRegister = fmt.Errorf(state.Err_ErrInsufficientStakeToRegister.String())
var ErrLibP2PKeyRequired = fmt.Errorf(state.Err_ErrLibP2PKeyRequired.String())
var ErrSenderNotRegistered = fmt.Errorf(state.Err_ErrSenderNotRegistered.String())
var ErrStakeTargetNotRegistered = fmt.Errorf(state.Err_ErrStakeTargetNotRegistered.String())
var ErrTopicIdOfStakerAndTargetDoNotMatch = fmt.Errorf(state.Err_ErrInvalidTopicId.String())
var ErrInsufficientStakeToRemove = fmt.Errorf(state.Err_ErrInsufficientStakeToRemove.String())
var ErrDoNotSetMapValueToZero = fmt.Errorf(state.Err_ErrDoNotSetMapValueToZero.String())
var ErrBlockHeightNegative = fmt.Errorf(state.Err_ErrBlockHeightNegative.String())
