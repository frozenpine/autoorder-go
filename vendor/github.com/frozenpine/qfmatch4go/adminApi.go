package qfmatch4go

/*
#cgo CFLAGS: -I${SRCDIR} -fPIC
#cgo LDFLAGS: -L${SRCDIR} -lcQFMatchSuperApi

#include "cgoAdminApi.h"
*/
import "C"

type adminAPICallbacks interface {
	// OnRspQuoteInsert
	// OnRspQuoteAction
	// OnRspOTCOrderInsert
	// OnRspUserPasswordUpdate
	// OnRspExecOrderInsert
	// OnRspExecOrderAction
	// OnRspAdminOrderInsert
	// OnRspCombOrderInsert
	// OnRspInformation
	// OnRspInstrumentStatusUpdate
	// OnRspForceUserExit
	// OnRspForceUserLogout
	// OnRspExchangeDataSyncStatusUpdate
	// OnRspSGDataSyncStatusUpdate
	// OnRspTBCommand
	// OnRspUpdateMarketData
	// OnRspPositionMarginAction
	// OnRspQryPartAccount
	// OnRspQryQuote
	// OnRspQryPartPosition
	// OnRspQrySGDataSyncStatus
	// OnRspQryHedgeVolume
	// OnRspQryCreditLimit
	// OnRspQryCombOrder
	// OnRspQryUserSession
	// OnRspQryUser
	// OnRspQryParticipant
	// OnRspQryInformation
	// OnRspQryClientPositionDetail
	// OnRtnExecOrder
	// OnRtnQuote
	// OnRtnCombOrder
	// OnRtnInsCombinationLeg
	// OnRtnDelCombinationLeg
	// OnRtnTradingAccount
	// OnRtnAliasDefine
	// OnRtnFlowMessageCancel
	// OnRtnAdminOrder
	// OnErrRtnQuoteInsert
	// OnErrRtnQuoteAction
	// OnErrRtnExecOrderInsert
	// OnErrRtnExecOrderAction
	// OnErrRtnCombOrderInsert
	// OnRspDumpMemTable
	// OnRspOTCTradeInsert
	// OnRtnOTCTrade
	// OnErrRtnOTCTrade
	// OnRspAccountDeposit
	// OnRspQryAccountDepositDetail
	// OnRspQryTBCommandDetail
	// OnRspQryPartClientFee
	// OnRtnMultiHeartbeat
}

type adminAPI struct {
	common *commonAPI
}

func (api *adminAPI) registerAdminCallbacks(ptrVtCallbacks *C.Callbacks) {
	// todo: register admin callback apis
}
