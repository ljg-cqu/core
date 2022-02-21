package openchain

// Transaction types
const (
	MethodDeposit Method = "DEPOSIT"

	MethodDeployWASMContract Method = "DEPLOYWASMCONTRACT"
	MethodCallWASMContract   Method = "CALLWASMCONTRACT"

	MethodUpdateContractForBiz Method = "UPDATECONTRACTFORBIZ"
	MethodDeployContractForBiz Method = "DEPLOYCONTRACTFORBIZ"

	MethodCallContractBizAsync Method = "CALLCONTRACTBIZASYNC"
	//MethodUpdateContractForBiz  Method = "UPDATECONTRACTFORBIZ"

	MethodTenantCreateAccount Method = "TENANTCREATEACCUNT"
	MethodParseOutput         Method = "PARSEOUTPUT"

	MethodQueryTransaction Method = "QUERYTRANSACTION"

	MethodQueryReceipt   Method = "QUERYRECEIPT"
	MethodQueryBlock     Method = "QUERYBLOCK"
	MethodQueryBlockBody Method = "QUERYBLOCKBODY"
	MethodQueryLastBlock Method = "QUERYLASTBLOCK"
	MethodQueryAccount   Method = "QUERYACCOUNT"
)

// Method stands for transaction method for interaction with open chain.
type Method string
