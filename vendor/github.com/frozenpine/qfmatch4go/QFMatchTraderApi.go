package qfmatch4go

/*
#cgo CFLAGS: -I${SRCDIR} -fPIC
#cgo LDFLAGS: -L${SRCDIR} -lcQFMatchSuperApi

#include <string.h>

#include "cgoTraderApi.h"
*/
import "C"
import "fmt"

func transformGoCQFMatchOrderField(order *GoCQFMatchOrderField) *C.struct_CQFMatchInputOrderField {
	req := C.struct_CQFMatchInputOrderField{}

	C.strncpy(&req.ParticipantID[0], C.CString(order.ParticipantID), C.sizeof_TQFMatchParticipantIDType-1)
	C.strncpy(&req.ClientID[0], C.CString(order.ClientID), C.sizeof_TQFMatchClientIDType-1)
	C.strncpy(&req.UserID[0], C.CString(order.UserID), C.sizeof_TQFMatchUserIDType-1)
	C.strncpy(&req.InstrumentID[0], C.CString(order.InstrumentID), C.sizeof_TQFMatchInstrumentIDType-1)
	req.OrderPriceType = ToCChar(order.OrderPriceType)
	req.Direction = ToCChar(order.Direction)

	for idx, value := range order.CombOffsetFlag {
		req.CombOffsetFlag[idx] = ToCChar(value)
	}

	for idx, value := range order.CombHedgeFlag {
		req.CombHedgeFlag[idx] = ToCChar(value)
	}

	req.LimitPrice = C.TQFMatchPriceType(order.LimitPrice)
	req.VolumeTotalOriginal = C.TQFMatchVolumeType(order.VolumeTotalOriginal)
	req.TimeCondition = ToCChar(order.TimeCondition)
	C.strncpy(&req.GTDDate[0], C.CString(order.GTDDate), C.sizeof_TQFMatchDateType-1)
	req.VolumeCondition = ToCChar(order.VolumeCondition)
	req.MinVolume = C.TQFMatchVolumeType(order.MinVolume)
	req.ContingentCondition = ToCChar(order.ContingentCondition)
	req.StopPrice = C.TQFMatchPriceType(order.StopPrice)

	orderLocalIDFormat := fmt.Sprint("%", C.sizeof_TQFMatchOrderLocalIDType-1, "d")
	req.ForceCloseReason = ToCChar(order.ForceCloseReason)
	C.strncpy(&req.OrderLocalID[0], C.CString(fmt.Sprintf(orderLocalIDFormat, order.OrderLocalID)), C.sizeof_TQFMatchOrderLocalIDType-1)

	if order.IsAutoSuspend {
		req.IsAutoSuspend = 1
	} else {
		req.IsAutoSuspend = 0
	}

	C.strncpy(&req.BusinessUnit[0], C.CString(order.BusinessUnit), C.sizeof_TQFMatchBusinessUnitType)

	return &req
}

func transformGoCQFMatchOrderActionField(orderAction *GoCQFMatchOrderActionField) *C.struct_CQFMatchOrderActionField {
	action := C.struct_CQFMatchOrderActionField{}

	orderSysIDFormat := fmt.Sprint("%", C.sizeof_TQFMatchOrderSysIDType-1, "d")
	C.strncpy(&action.OrderSysID[0], C.CString(fmt.Sprintf(orderSysIDFormat, orderAction.OrderSysID)), C.sizeof_TQFMatchOrderSysIDType-1)

	orderLocalIDFormat := fmt.Sprint("%", C.sizeof_TQFMatchOrderLocalIDType-1, "d")
	C.strncpy(&action.OrderLocalID[0], C.CString(fmt.Sprintf(orderLocalIDFormat, orderAction.OrderLocalID)), C.sizeof_TQFMatchOrderLocalIDType-1)

	action.ActionFlag = ToCChar(orderAction.ActionFlag)
	C.strncpy(&action.ParticipantID[0], C.CString(orderAction.ParticipantID), C.sizeof_TQFMatchParticipantIDType-1)
	C.strncpy(&action.ClientID[0], C.CString(orderAction.ClientID), C.sizeof_TQFMatchClientIDType-1)
	C.strncpy(&action.UserID[0], C.CString(orderAction.UserID), C.sizeof_TQFMatchUserIDType-1)
	action.LimitPrice = C.TQFMatchPriceType(orderAction.LimitPrice)
	action.VolumeChange = C.TQFMatchVolumeType(orderAction.VolumeChange)

	actionLocalIDFormat := fmt.Sprint("%", C.sizeof_TQFMatchOrderLocalIDType-1, "d")
	C.strncpy(&action.ActionLocalID[0], C.CString(fmt.Sprintf(actionLocalIDFormat, orderAction.ActionLocalID)), C.sizeof_TQFMatchOrderLocalIDType-1)

	C.strncpy(&action.BusinessUnit[0], C.CString(orderAction.BusinessUnit), C.sizeof_TQFMatchBusinessUnitType-1)

	return &action
}

type tradeAPICallbacks interface {
	OnRspOrderInsert(rspOrderInsert *GoCQFMatchInputOrderField, err *GoCQFMatchRspInfoField, requestID int, isLast bool)
	OnRspOrderAction(rspOrderAction *GoCQFMatchOrderActionField, err *GoCQFMatchRspInfoField, requestID int, isLast bool)
	// OnRspQryPartAccount
	OnRspQryOrder(rspQryOrder *GoCQFMatchOrderField, err *GoCQFMatchRspInfoField, requestID int, isLast bool)
	OnRspQryTrade(rspQryTrade *GoCQFMatchTradeField, err *GoCQFMatchRspInfoField, requestID int, isLast bool)
	// OnRspQryClient
	// OnRspQryClientPosition
	OnRtnTrade(rtnTrade *GoCQFMatchTradeField)
	OnRtnOrder(rtnOrder *GoCQFMatchOrderField)
	// OnRtnClientPosition
	// OnRtnFlowMessageCancel
	// OnErrRtnOrderInsert
	// OnErrRtnOrderAction
	// OnRspQryPartClientFee
}

//export goOnRspOrderInsert
func goOnRspOrderInsert(instance C.QFMatchSuperApiInstance, pInputOrder *C.struct_CQFMatchInputOrderField, pRspInfo *C.struct_CQFMatchRspInfoField, nRequestID C.int, bIsLast C.bool) {

}

//export goOnRspOrderAction
func goOnRspOrderAction(instance C.QFMatchSuperApiInstance, pOrderAction *C.struct_CQFMatchOrderActionField, pRspInfo *C.struct_CQFMatchRspInfoField, nRequestID C.int, bIsLast C.bool) {

}

//export goOnRspQryPartAccount
// func goOnRspQryPartAccount(instance C.QFMatchSuperApiInstance, pRspPartAccount *C.struct_CQFMatchRspPartAccountField, pRspInfo *C.struct_CQFMatchRspInfoField, nRequestID C.int, bIsLast C.bool) {}

//export goOnRspQryOrder
func goOnRspQryOrder(instance C.QFMatchSuperApiInstance, pOrder *C.struct_CQFMatchOrderField, pRspInfo *C.struct_CQFMatchRspInfoField, nRequestID C.int, bIsLast C.bool) {

}

//export goOnRspQryTrade
func goOnRspQryTrade(instance C.QFMatchSuperApiInstance, pTrade *C.struct_CQFMatchTradeField, pRspInfo *C.struct_CQFMatchRspInfoField, nRequestID C.int, bIsLast C.bool) {

}

//export goOnRspQryClient
// func goOnRspQryClient(instance C.QFMatchSuperApiInstance, pRspClient *C.struct_CQFMatchRspClientField, pRspInfo *C.struct_CQFMatchRspInfoField, nRequestID C.int, bIsLast C.bool) {}

//export goOnRspQryClientPosition
// func goOnRspQryClientPosition(instance C.QFMatchSuperApiInstance, pRspClientPosition *C.struct_CQFMatchRspClientPositionField, pRspInfo *C.struct_CQFMatchRspInfoField, nRequestID C.int, bIsLast C.bool) {}

//export goOnRtnTrade
func goOnRtnTrade(instance C.QFMatchSuperApiInstance, pTrade *C.struct_CQFMatchTradeField) {

}

//export goOnRtnOrder
func goOnRtnOrder(instance C.QFMatchSuperApiInstance, pOrder *C.struct_CQFMatchOrderField) {

}

//export OnRtnClientPosition
// func OnRtnClientPosition(instance C.QFMatchSuperApiInstance, pClientPosition *C.struct_CQFMatchClientPositionField) {}

//export goOnRtnFlowMessageCancel
// func goOnRtnFlowMessageCancel(instance C.QFMatchSuperApiInstance, pFlowMessageCancel *C.struct_CQFMatchFlowMessageCancelField) {}

//export goOnErrRtnOrderInsert
// func goOnErrRtnOrderInsert(instance C.QFMatchSuperApiInstance, pInputOrder *C.struct_CQFMatchInputOrderField, pRspInfo *C.struct_CQFMatchRspInfoField) {}

//export goOnErrRtnOrderAction
// func goOnErrRtnOrderAction(instance C.QFMatchSuperApiInstance, pOrderAction *C.struct_CQFMatchOrderActionField, pRspInfo *C.struct_CQFMatchRspInfoField) {}

//export goOnRspQryPartClientFee
// func goOnRspQryPartClientFee(instance C.QFMatchSuperApiInstance, pPartClientFee *C.struct_CQFMatchPartClientFeeField, pRspInfo *C.struct_CQFMatchRspInfoField, nRequestID C.int, bIsLast C.bool) {}

type traderAPI struct {
	common *commonAPI
}

// QFMatchTraderAPICallbacks Quantfair撮合引擎交易API回调接口
type QFMatchTraderAPICallbacks interface {
	commonAPICallbacks
	marketAPICallbacks
	tradeAPICallbacks
}

// QFMatchTraderAPI Quantfair撮合引擎交易接口
type QFMatchTraderAPI struct {
	commonAPI
	marketAPI
	traderAPI
}

func (api *traderAPI) registerTraderCallbacks(ptrVtCallbacks *C.Callbacks) {
	ptrVtCallbacks.ptrOnRspOrderInsert = C.FuncOnRspOrderInsert(C.cgoOnRspOrderInsert)
	ptrVtCallbacks.ptrOnRspOrderAction = C.FuncOnRspOrderAction(C.cgoOnRspOrderAction)
	// ptrVtCallbacks.ptrOnRspQryPartAccount = C.FuncOnRspQryPartAccount(C.cgoOnRspQryPartAccount)
	ptrVtCallbacks.ptrOnRspQryOrder = C.FuncOnRspQryOrder(C.cgoOnRspQryOrder)
	ptrVtCallbacks.ptrOnRspQryTrade = C.FuncOnRspQryTrade(C.cgoOnRspQryTrade)
	// ptrVtCallbacks.ptrOnRspQryClient = C.FuncOnRspQryClient(C.cgoOnRspQryClient)
	// ptrVtCallbacks.ptrOnRspQryClientPosition = C.FuncOnRspQryClientPosition(C.cgoOnRspQryClientPosition)
	ptrVtCallbacks.ptrOnRtnTrade = C.FuncOnRtnTrade(C.cgoOnRtnTrade)
	ptrVtCallbacks.ptrOnRtnOrder = C.FuncOnRtnOrder(C.cgoOnRtnOrder)
	// ptrVtCallbacks.ptrOnRtnClientPosition = C.FuncOnRtnClientPosition(C.cgoOnRtnClientPosition)
	// ptrVtCallbacks.ptrOnRtnFlowMessageCancel = C.FuncOnRtnFlowMessageCancel(C.cgoOnRtnFlowMessageCancel)
	// ptrVtCallbacks.ptrOnErrRtnOrderInsert = C.FuncOnErrRtnOrderInsert(C.cgoOnErrRtnOrderInsert)
	// ptrVtCallbacks.ptrOnErrRtnOrderAction = C.FuncOnErrRtnOrderAction(C.cgoOnErrRtnOrderAction)
	// ptrVtCallbacks.ptrOnRspQryPartClientFee = C.FuncOnRspQryPartClientFee(C.cgoOnRspQryPartClientFee)
}

// RegisterCallbacks 注册回调接口
func (api *QFMatchTraderAPI) RegisterCallbacks(cb interface{}) {
	vtCallbacks := C.Callbacks{}

	switch cb.(type) {
	case QFMatchMarketAPICallbacks:
		api.registerMarketCallbacks(&vtCallbacks)
		api.marketAPI.common = &api.commonAPI
	case QFMatchTraderAPICallbacks:
		api.registerMarketCallbacks(&vtCallbacks)
		api.registerTraderCallbacks(&vtCallbacks)
		api.marketAPI.common = &api.commonAPI
		api.traderAPI.common = &api.commonAPI
	default:
		panic("Invalid callback interface.")
	}

	api.registerCommonCallbacks(&vtCallbacks)

	C.RegisterCallbacks(C.QFMatchSuperApiInstance(api.apiInstance), &vtCallbacks)

	apiMap[api.apiInstance] = cb
}

// ReqOrderInsert 报单录入请求
func (api *traderAPI) ReqOrderInsert(order *GoCQFMatchOrderField) int {
	rtn := C.ReqOrderInsert(C.QFMatchSuperApiInstance(api.common.apiInstance), transformGoCQFMatchOrderField(order), C.int(api.common.getRequestID()))

	return int(rtn)
}

// ReqOrderAction 报单操作请求
func (api *traderAPI) ReqOrderAction(orderAction *GoCQFMatchOrderActionField) int {
	rtn := C.ReqOrderAction(C.QFMatchSuperApiInstance(api.common.apiInstance), transformGoCQFMatchOrderActionField(orderAction), C.int(api.common.getRequestID()))

	return int(rtn)
}
