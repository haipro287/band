// Code generated by protoconstructorgen.py. DO NOT EDIT.
package types

import github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
import "time"

func NewMsgRequestData(
	OracleScriptID OracleScriptID,
	Calldata []byte,
	AskCount uint64,
	MinCount uint64,
	ClientID string,
	Sender github_com_cosmos_cosmos_sdk_types.AccAddress,
) *MsgRequestData {
	return &MsgRequestData{
		OracleScriptID: OracleScriptID,
		Calldata:       Calldata,
		AskCount:       AskCount,
		MinCount:       MinCount,
		ClientID:       ClientID,
		Sender:         Sender.String(),
	}
}

func NewMsgReportData(
	RequestID RequestID,
	RawReports []RawReport,
	Validator github_com_cosmos_cosmos_sdk_types.ValAddress,
	Reporter github_com_cosmos_cosmos_sdk_types.AccAddress,
) MsgReportData {
	return MsgReportData{
		RequestID:  RequestID,
		RawReports: RawReports,
		Validator:  Validator.String(),
		Reporter:   Reporter.String(),
	}
}

func NewMsgCreateDataSource(
	Owner github_com_cosmos_cosmos_sdk_types.AccAddress,
	Name string,
	Description string,
	Executable []byte,
	Sender github_com_cosmos_cosmos_sdk_types.AccAddress,
) MsgCreateDataSource {
	return MsgCreateDataSource{
		Owner:       Owner.String(),
		Name:        Name,
		Description: Description,
		Executable:  Executable,
		Sender:      Sender.String(),
	}
}

func NewMsgEditDataSource(
	DataSourceID DataSourceID,
	Owner github_com_cosmos_cosmos_sdk_types.AccAddress,
	Name string,
	Description string,
	Executable []byte,
	Sender github_com_cosmos_cosmos_sdk_types.AccAddress,
) MsgEditDataSource {
	return MsgEditDataSource{
		DataSourceID: DataSourceID,
		Owner:        Owner.String(),
		Name:         Name,
		Description:  Description,
		Executable:   Executable,
		Sender:       Sender.String(),
	}
}

func NewMsgCreateOracleScript(
	Owner github_com_cosmos_cosmos_sdk_types.AccAddress,
	Name string,
	Description string,
	Code []byte,
	Schema string,
	SourceCodeURL string,
	Sender github_com_cosmos_cosmos_sdk_types.AccAddress,
) MsgCreateOracleScript {
	return MsgCreateOracleScript{
		Owner:         Owner.String(),
		Name:          Name,
		Description:   Description,
		Code:          Code,
		Schema:        Schema,
		SourceCodeURL: SourceCodeURL,
		Sender:        Sender.String(),
	}
}

func NewMsgEditOracleScript(
	OracleScriptID OracleScriptID,
	Owner github_com_cosmos_cosmos_sdk_types.AccAddress,
	Name string,
	Description string,
	Code []byte,
	Schema string,
	SourceCodeURL string,
	Sender github_com_cosmos_cosmos_sdk_types.AccAddress,
) MsgEditOracleScript {
	return MsgEditOracleScript{
		OracleScriptID: OracleScriptID,
		Owner:          Owner.String(),
		Name:           Name,
		Description:    Description,
		Code:           Code,
		Schema:         Schema,
		SourceCodeURL:  SourceCodeURL,
		Sender:         Sender.String(),
	}
}

func NewMsgActivate(
	Validator github_com_cosmos_cosmos_sdk_types.ValAddress,
) MsgActivate {
	return MsgActivate{
		Validator: Validator.String(),
	}
}

func NewMsgAddReporter(
	Validator github_com_cosmos_cosmos_sdk_types.ValAddress,
	Reporter github_com_cosmos_cosmos_sdk_types.AccAddress,
) MsgAddReporter {
	return MsgAddReporter{
		Validator: Validator.String(),
		Reporter:  Reporter.String(),
	}
}

func NewMsgRemoveReporter(
	Validator github_com_cosmos_cosmos_sdk_types.ValAddress,
	Reporter github_com_cosmos_cosmos_sdk_types.AccAddress,
) MsgRemoveReporter {
	return MsgRemoveReporter{
		Validator: Validator.String(),
		Reporter:  Reporter.String(),
	}
}

func NewDataSource(
	Owner github_com_cosmos_cosmos_sdk_types.AccAddress,
	Name string,
	Description string,
	Filename string,
) DataSource {
	return DataSource{
		Owner:       Owner.String(),
		Name:        Name,
		Description: Description,
		Filename:    Filename,
	}
}

func NewOracleScript(
	Owner github_com_cosmos_cosmos_sdk_types.AccAddress,
	Name string,
	Description string,
	Filename string,
	Schema string,
	SourceCodeURL string,
) OracleScript {
	return OracleScript{
		Owner:         Owner.String(),
		Name:          Name,
		Description:   Description,
		Filename:      Filename,
		Schema:        Schema,
		SourceCodeURL: SourceCodeURL,
	}
}

func NewRawRequest(
	ExternalID ExternalID,
	DataSourceID DataSourceID,
	Calldata []byte,
) RawRequest {
	return RawRequest{
		ExternalID:   ExternalID,
		DataSourceID: DataSourceID,
		Calldata:     Calldata,
	}
}

func NewRawReport(
	ExternalID ExternalID,
	ExitCode uint32,
	Data []byte,
) RawReport {
	return RawReport{
		ExternalID: ExternalID,
		ExitCode:   ExitCode,
		Data:       Data,
	}
}

func NewRequest(
	OracleScriptID OracleScriptID,
	Calldata []byte,
	RequestedValidators []github_com_cosmos_cosmos_sdk_types.ValAddress,
	MinCount uint64,
	RequestHeight int64,
	RequestTime time.Time,
	ClientID string,
	RawRequests []RawRequest,
) Request {
	requestedVals := make([]string, len(RequestedValidators))
	for idx, reqVal := range RequestedValidators {
		requestedVals[idx] = reqVal.String()
	}
	return Request{
		OracleScriptID:      OracleScriptID,
		Calldata:            Calldata,
		RequestedValidators: requestedVals,
		MinCount:            MinCount,
		RequestHeight:       RequestHeight,
		RequestTime:         uint64(RequestTime.UnixNano()),
		ClientID:            ClientID,
		RawRequests:         RawRequests,
	}
}

func NewReport(
	Validator github_com_cosmos_cosmos_sdk_types.ValAddress,
	InBeforeResolve bool,
	RawReports []RawReport,
) Report {
	return Report{
		Validator:       Validator.String(),
		InBeforeResolve: InBeforeResolve,
		RawReports:      RawReports,
	}
}

func NewOracleRequestPacketData(
	ClientID string,
	OracleScriptID OracleScriptID,
	Calldata []byte,
	AskCount uint64,
	MinCount uint64,
) OracleRequestPacketData {
	return OracleRequestPacketData{
		ClientID:       ClientID,
		OracleScriptID: OracleScriptID,
		Calldata:       Calldata,
		AskCount:       AskCount,
		MinCount:       MinCount,
	}
}

func NewOracleResponsePacketData(
	ClientID string,
	RequestID RequestID,
	AnsCount uint64,
	RequestTime int64,
	ResolveTime int64,
	ResolveStatus ResolveStatus,
	Result []byte,
) OracleResponsePacketData {
	return OracleResponsePacketData{
		ClientID:      ClientID,
		RequestID:     RequestID,
		AnsCount:      AnsCount,
		RequestTime:   RequestTime,
		ResolveTime:   ResolveTime,
		ResolveStatus: ResolveStatus,
		Result:        Result,
	}
}

func NewValidatorStatus(
	IsActive bool,
	Since time.Time,
) ValidatorStatus {
	return ValidatorStatus{
		IsActive: IsActive,
		Since:    Since,
	}
}

func NewParams(
	MaxRawRequestCount uint64,
	MaxAskCount uint64,
	ExpirationBlockCount uint64,
	BaseRequestGas uint64,
	PerValidatorRequestGas uint64,
	SamplingTryCount uint64,
	OracleRewardPercentage uint64,
	InactivePenaltyDuration uint64,
) Params {
	return Params{
		MaxRawRequestCount:      MaxRawRequestCount,
		MaxAskCount:             MaxAskCount,
		ExpirationBlockCount:    ExpirationBlockCount,
		BaseRequestGas:          BaseRequestGas,
		PerValidatorRequestGas:  PerValidatorRequestGas,
		SamplingTryCount:        SamplingTryCount,
		OracleRewardPercentage:  OracleRewardPercentage,
		InactivePenaltyDuration: InactivePenaltyDuration,
	}
}
