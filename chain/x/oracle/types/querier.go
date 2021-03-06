package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Query endpoints supported by the oracle Querier.
const (
	QueryParams            = "params"
	QueryCounts            = "counts"
	QueryData              = "data"
	QueryDataSources       = "data_sources"
	QueryOracleScripts     = "oracle_scripts"
	QueryRequests          = "requests"
	QueryValidatorStatus   = "validator_status"
	QueryReporters         = "reporters"
	QueryActiveValidators  = "active_validators"
	QueryPendingRequests   = "pending_requests"
	QueryDataProvidersPool = "data_providers_pool"
)

// QueryCountsResult is the struct for the result of query counts.
type QueryCountsResult struct {
	DataSourceCount   int64 `json:"data_source_count"`
	OracleScriptCount int64 `json:"oracle_script_count"`
	RequestCount      int64 `json:"request_count"`
}

// QueryRequestResult is the struct for the result of request query.
type QueryRequestResult struct {
	Request Request  `json:"request"`
	Reports []Report `json:"reports"`
	Result  *Result  `json:"result"`
}

// QueryActiveValidatorResult is the struct for the result of request active validators.
type QueryActiveValidatorResult struct {
	Address sdk.ValAddress `json:"address"`
	Power   uint64         `json:"power"`
}
