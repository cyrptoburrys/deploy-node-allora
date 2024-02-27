package emissions

import "cosmossdk.io/errors"

var (
	ErrIntegerUnderflowDelegator                         = errors.Register(ModuleName, 1, "integer underflow for delegator")
	ErrIntegerUnderflowBonds                             = errors.Register(ModuleName, 2, "integer underflow for bonds")
	ErrIntegerUnderflowTarget                            = errors.Register(ModuleName, 3, "integer underflow for target")
	ErrIntegerUnderflowTopicStake                        = errors.Register(ModuleName, 4, "integer underflow for topic stake")
	ErrIntegerUnderflowTotalStake                        = errors.Register(ModuleName, 5, "integer underflow for total stake")
	ErrIterationLengthDoesNotMatch                       = errors.Register(ModuleName, 6, "iteration length does not match")
	ErrInvalidTopicId                                    = errors.Register(ModuleName, 7, "invalid topic ID")
	ErrReputerAlreadyRegisteredInTopic                   = errors.Register(ModuleName, 8, "reputer already registered in topic")
	ErrWorkerAlreadyRegisteredInTopic                    = errors.Register(ModuleName, 9, "worker already registered in topic")
	ErrAddressAlreadyRegisteredInATopic                  = errors.Register(ModuleName, 10, "address already registered in a topic")
	ErrAddressIsNotRegisteredInAnyTopic                  = errors.Register(ModuleName, 11, "address is not registered in any topic")
	ErrAddressIsNotRegisteredInThisTopic                 = errors.Register(ModuleName, 12, "address is not registered in this topic")
	ErrInsufficientStakeToRegister                       = errors.Register(ModuleName, 13, "insufficient stake to register")
	ErrLibP2PKeyRequired                                 = errors.Register(ModuleName, 14, "libp2p key required")
	ErrAddressNotRegistered                              = errors.Register(ModuleName, 15, "address not registered")
	ErrStakeTargetNotRegistered                          = errors.Register(ModuleName, 16, "stake target not registered")
	ErrTopicIdOfStakerAndTargetDoNotMatch                = errors.Register(ModuleName, 17, "topic ID of staker and target do not match")
	ErrInsufficientStakeToRemove                         = errors.Register(ModuleName, 18, "insufficient stake to remove")
	ErrNoStakeToRemove                                   = errors.Register(ModuleName, 19, "no stake to remove")
	ErrDoNotSetMapValueToZero                            = errors.Register(ModuleName, 20, "do not set map value to zero")
	ErrBlockHeightNegative                               = errors.Register(ModuleName, 21, "block height negative")
	ErrBlockHeightLessThanPrevious                       = errors.Register(ModuleName, 22, "block height less than previous")
	ErrModifyStakeBeforeBondLessThanAmountModified       = errors.Register(ModuleName, 23, "modify stake before bond less than amount modified")
	ErrModifyStakeBeforeSumGreaterThanSenderStake        = errors.Register(ModuleName, 24, "modify stake before sum greater than sender stake")
	ErrModifyStakeSumBeforeNotEqualToSumAfter            = errors.Register(ModuleName, 25, "modify stake sum before not equal to sum after")
	ErrConfirmRemoveStakeNoRemovalStarted                = errors.Register(ModuleName, 26, "confirm remove stake no removal started")
	ErrConfirmRemoveStakeTooEarly                        = errors.Register(ModuleName, 27, "confirm remove stake too early")
	ErrConfirmRemoveStakeTooLate                         = errors.Register(ModuleName, 28, "confirm remove stake too late")
	ErrScalarMultiplyNegative                            = errors.Register(ModuleName, 29, "scalar multiply negative")
	ErrDivideMapValuesByZero                             = errors.Register(ModuleName, 30, "divide map values by zero")
	ErrTopicIdListValueDecodeInvalidLength               = errors.Register(ModuleName, 31, "topic ID list value decode invalid length")
	ErrTopicIdListValueDecodeJsonInvalidLength           = errors.Register(ModuleName, 32, "topic ID list value decode JSON invalid length")
	ErrTopicIdListValueDecodeJsonInvalidFormat           = errors.Register(ModuleName, 33, "topic ID list value decode JSON invalid format")
	ErrTopicDoesNotExist                                 = errors.Register(ModuleName, 34, "topic does not exist")
	ErrCannotRemoveMoreStakeThanStakedInTopic            = errors.Register(ModuleName, 35, "cannot remove more stake than staked in topic")
	ErrInferenceRequestAlreadyInMempool                  = errors.Register(ModuleName, 36, "inference request already in mempool")
	ErrInferenceRequestBidAmountLessThanPrice            = errors.Register(ModuleName, 37, "inference request bid amount less than price")
	ErrInferenceRequestTimestampValidUntilInPast         = errors.Register(ModuleName, 38, "inference request timestamp valid until in past")
	ErrInferenceRequestTimestampValidUntilTooFarInFuture = errors.Register(ModuleName, 39, "inference request timestamp valid until too far in future")
	ErrInferenceRequestCadenceTooFast                    = errors.Register(ModuleName, 40, "inference request cadence too fast")
	ErrInferenceRequestCadenceTooSlow                    = errors.Register(ModuleName, 41, "inference request cadence too slow")
	ErrInferenceRequestWillNeverBeScheduled              = errors.Register(ModuleName, 42, "inference request will never be scheduled")
	ErrOwnerCannotBeEmpty                                = errors.Register(ModuleName, 43, "owner cannot be empty")
	ErrInsufficientStakeAfterRemoval                     = errors.Register(ModuleName, 44, "insufficient stake after removal")
	ErrInferenceRequestBidAmountTooLow                   = errors.Register(ModuleName, 45, "inference request bid amount too low")
	ErrIntegerUnderflowUnmetDemand                       = errors.Register(ModuleName, 46, "integer underflow for unmet demand")
	ErrInferenceCadenceBelowMinimum                      = errors.Register(ModuleName, 47, "inference cadence must be at least 60 seconds (1 minute)")
	ErrWeightCadenceBelowMinimum                         = errors.Register(ModuleName, 48, "weight cadence must be at least 10800 seconds (3 hours)")
	ErrNotWhitelistAdmin                                 = errors.Register(ModuleName, 49, "not whitelist admin")
	ErrNotInTopicCreationWhitelist                       = errors.Register(ModuleName, 50, "not in topic creation whitelist")
	ErrNotInWeightSettingWhitelist                       = errors.Register(ModuleName, 51, "not in topic weight setting whitelist")
	ErrTopicNotEnoughDemand                              = errors.Register(ModuleName, 52, "topic not enough demand")
	ErrSetParamsVersion                                  = errors.Register(ModuleName, 53, "Error Setting Params: Version")
	ErrSetParamsEpochLength                              = errors.Register(ModuleName, 54, "Error Setting Params: EpochLength")
	ErrSetParamsEmissionsPerEpoch                        = errors.Register(ModuleName, 55, "Error Setting Params: EmissionsPerEpoch")
	ErrSetParamsMinTopicUnmetDemand                      = errors.Register(ModuleName, 56, "Error Setting Params: MinTopicUnmetDemand")
	ErrSetParamsMaxTopicsPerBlock                        = errors.Register(ModuleName, 57, "Error Setting Params: MaxTopicsPerBlock")
	ErrSetParamsMinRequestUnmetDemand                    = errors.Register(ModuleName, 59, "Error Setting Params: MinRequestUnmetDemand")
	ErrSetParamsMaxAllowableMissingInferencePercent      = errors.Register(ModuleName, 60, "Error Setting Params: MaxAllowableMissingInferencePercent")
	ErrSetParamsRequiredMinimumStake                     = errors.Register(ModuleName, 61, "Error Setting Params: RequiredMinimumStake")
	ErrSetParamsRemoveStakeDelayWindow                   = errors.Register(ModuleName, 62, "Error Setting Params: RemoveStakeDelayWindow")
	ErrSetParamsMinFastestAllowedCadence                 = errors.Register(ModuleName, 63, "Error Setting Params: MinFastestAllowedCadence")
	ErrSetParamsMaxInferenceRequestValidity              = errors.Register(ModuleName, 64, "Error Setting Params: MaxInferenceRequestValidity")
	ErrSetParamsMaxSlowestAllowedCadence                 = errors.Register(ModuleName, 65, "Error Setting Params: MaxSlowestAllowedCadence")
)
