package qfmatch4go

/*
#cgo CFLAGS: -I${SRCDIR} -fPIC
#cgo LDFLAGS: -L${SRCDIR} -lcQFMatchSuperApi

#include <string.h>

#include "cgoMarketApi.h"
*/
import "C"
import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/djimenez/iconv-go"
)

type marketAPICallbacks interface {
	OnRspSubscribeTopic(rspSubscribeTopic *GoCQFMatchDisseminationField, err *GoCQFMatchRspInfoField, requestID int, isLast bool)
	OnRspQryMarketData(marketData *GoCQFMatchMarketDataField, err *GoCQFMatchRspInfoField, requestID int, isLast bool)
	OnRspQryInstrument(instrument *GoCQFMatchRspInstrumentField, err *GoCQFMatchRspInfoField, requestID int, isLast bool)
	OnRspQryInstrumentStatus(insStatus *GoCQFMatchInstrumentStatusField, err *GoCQFMatchRspInfoField, requestID int, isLast bool)
	// OnRtnInsInstrument
	// OnRtnDelInstrument
	OnRspQryTopic(dissemination *GoCQFMatchDisseminationField, err *GoCQFMatchRspInfoField, requestID int, isLast bool)
	OnRspQryMBLMarketData(mblMarketData *GoCQFMatchMBLMarketDataField, err *GoCQFMatchRspInfoField, requestID int, isLast bool)
	OnRtnInstrumentStatus(insStatus *GoCQFMatchInstrumentStatusField)
	OnRtnMarketData(marketData *GoCQFMatchMarketDataField)
	OnRtnDepthMarketData(depthMarketData *GoCQFMatchDepthMarketDataField)
}

func convertDisseminationField(pDissemination *C.struct_CQFMatchDisseminationField) *GoCQFMatchDisseminationField {
	data := GoCQFMatchDisseminationField{}

	data.SequenceSeries = int(pDissemination.SequenceSeries)
	data.SequenceNo = int(pDissemination.SequenceNo)

	return &data
}

func convertMarketDataField(pMarketData *C.struct_CQFMatchMarketDataField) *GoCQFMatchMarketDataField {
	data := GoCQFMatchMarketDataField{}

	data.TradingDay = C.GoString(&pMarketData.TradingDay[0])
	data.SettlementGroupID = C.GoString(&pMarketData.SettlementGroupID[0])
	data.SettlementID = int(pMarketData.SettlementID)
	data.LastPrice = float64(pMarketData.LastPrice)
	data.PreSettlementPrice = float64(pMarketData.PreSettlementPrice)
	data.PreClosePrice = float64(pMarketData.PreClosePrice)
	data.PreOpenInterest = float64(pMarketData.PreOpenInterest)
	data.OpenPrice = float64(pMarketData.OpenPrice)
	data.HighestPrice = float64(pMarketData.HighestPrice)
	data.LowestPrice = float64(pMarketData.LowestPrice)
	data.Volume = int(pMarketData.Volume)
	data.Turnover = float64(pMarketData.Turnover)
	data.OpenInterest = float64(pMarketData.OpenInterest)
	data.ClosePrice = float64(pMarketData.ClosePrice)
	data.SettlementPrice = float64(pMarketData.SettlementPrice)
	data.UpperLimitPrice = float64(pMarketData.UpperLimitPrice)
	data.LowerLimitPrice = float64(pMarketData.LowerLimitPrice)
	data.PreDelta = float64(pMarketData.PreDelta)
	data.CurrDelta = float64(pMarketData.CurrDelta)

	updateTime, _ := time.Parse("15:04:05", C.GoString(&pMarketData.UpdateTime[0]))
	data.UpdateTime = updateTime

	data.InstrumentID = C.GoString(&pMarketData.InstrumentID[0])
	data.FundingRate = float64(pMarketData.FundingRate)
	data.FairPrice = float64(pMarketData.FairPrice)

	return &data
}

func convertRspInstrumentField(pRspInstrument *C.struct_CQFMatchRspInstrumentField) *GoCQFMatchRspInstrumentField {
	data := GoCQFMatchRspInstrumentField{}

	data.SettlementGroupID = C.GoString(&pRspInstrument.SettlementGroupID[0])
	data.ProductID = C.GoString(&pRspInstrument.ProductID[0])
	data.ProductGroupID = C.GoString(&pRspInstrument.ProductGroupID[0])
	data.UnderlyingInstrID = C.GoString(&pRspInstrument.UnderlyingInstrID[0])

	prdClass, _ := strconv.ParseUint(string(pRspInstrument.ProductClass), 10, 8)
	data.ProductClass = productClass(prdClass)

	posType, _ := strconv.ParseUint(string(pRspInstrument.PositionType), 10, 8)
	data.PositionType = positionType(posType)

	data.StrikePrice = float64(pRspInstrument.StrikePrice)

	optType, _ := strconv.ParseUint(string(pRspInstrument.OptionsType), 10, 8)
	data.OptionsType = optionsType(optType)

	data.VolumeMultiple = int(pRspInstrument.VolumeMultiple)
	data.UnderlyingMultiple = float64(pRspInstrument.UnderlyingMultiple)
	data.InstrumentID = C.GoString(&pRspInstrument.InstrumentID[0])

	insName, _ := iconv.ConvertString(C.GoString(&pRspInstrument.InstrumentName[0]), "gbk", "utf-8")
	data.InstrumentName = insName

	data.DeliveryYear = int(pRspInstrument.DeliveryYear)
	data.DeliveryMonth = int(pRspInstrument.DeliveryMonth)

	advMonth, _ := strconv.Atoi(C.GoString(&pRspInstrument.AdvanceMonth[0]))
	data.AdvanceMonth = int(advMonth)

	data.IsTrading = pRspInstrument.IsTrading != 0
	data.CreateDate = C.GoString(&pRspInstrument.CreateDate[0])
	data.OpenDate = C.GoString(&pRspInstrument.OpenDate[0])
	data.ExpireDate = C.GoString(&pRspInstrument.ExpireDate[0])
	data.StartDelivDate = C.GoString(&pRspInstrument.StartDelivDate[0])
	data.EndDelivDate = C.GoString(&pRspInstrument.EndDelivDate[0])
	data.BasisPrice = float64(pRspInstrument.BasisPrice)
	data.MaxMarketOrderVolume = int(pRspInstrument.MaxMarketOrderVolume)
	data.MinMarketOrderVolume = int(pRspInstrument.MinMarketOrderVolume)
	data.MaxLimitOrderVolume = int(pRspInstrument.MaxLimitOrderVolume)
	data.MinLimitOrderVolume = int(pRspInstrument.MinLimitOrderVolume)
	data.PriceTick = float64(pRspInstrument.PriceTick)
	data.AllowDelivPersonOpen = pRspInstrument.AllowDelivPersonOpen != 0
	data.Currency = C.GoString(&pRspInstrument.Currency[0])

	return &data
}

func convertInstrumentStatusField(pInstrumentStatus *C.struct_CQFMatchInstrumentStatusField) *GoCQFMatchInstrumentStatusField {
	data := GoCQFMatchInstrumentStatusField{}

	data.SettlementGroupID = C.GoString(&pInstrumentStatus.SettlementGroupID[0])
	data.InstrumentID = C.GoString(&pInstrumentStatus.InstrumentID[0])

	status, _ := strconv.ParseUint(string(pInstrumentStatus.InstrumentStatus), 10, 8)
	data.InstrumentStatus = instrumentStatus(status)

	data.TradingSegmentSN = int(pInstrumentStatus.TradingSegmentSN)

	enterTime, _ := time.Parse("20060102 15:04:05", fmt.Sprintf("%s %s", C.GoString(&pInstrumentStatus.CalendarDate[0]), C.GoString(&pInstrumentStatus.EnterTime[0])))
	data.EnterTime = enterTime

	reason, _ := strconv.ParseUint(string(pInstrumentStatus.EnterReason), 10, 8)
	data.EnterReason = enterReason(reason)

	return &data
}

func convertCQFMatchMBLMarketDataField(pMBLMarketData *C.struct_CQFMatchMBLMarketDataField) *GoCQFMatchMBLMarketDataField {
	data := GoCQFMatchMBLMarketDataField{}

	data.InstrumentID = C.GoString(&pMBLMarketData.InstrumentID[0])

	d, _ := strconv.ParseUint(string(pMBLMarketData.Direction), 10, 8)
	data.Direction = direction(d)

	data.Price = float64(pMBLMarketData.Price)
	data.Volume = int(pMBLMarketData.Volume)

	return &data
}

func convertDepthMarketDataField(pDepthMarketData *C.struct_CQFMatchDepthMarketDataField) *GoCQFMatchDepthMarketDataField {
	data := GoCQFMatchDepthMarketDataField{}

	data.TradingDay = C.GoString(&pDepthMarketData.TradingDay[0])
	data.SettlementGroupID = C.GoString(&pDepthMarketData.SettlementGroupID[0])
	data.SettlementID = int(pDepthMarketData.SettlementID)
	data.LastPrice = float64(pDepthMarketData.LastPrice)
	data.PreSettlementPrice = float64(pDepthMarketData.PreSettlementPrice)
	data.PreClosePrice = float64(pDepthMarketData.PreClosePrice)
	data.PreOpenInterest = float64(pDepthMarketData.PreOpenInterest)
	data.OpenPrice = float64(pDepthMarketData.OpenPrice)
	data.HighestPrice = float64(pDepthMarketData.HighestPrice)
	data.LowestPrice = float64(pDepthMarketData.LowestPrice)
	data.Volume = int(pDepthMarketData.Volume)
	data.Turnover = float64(pDepthMarketData.Turnover)
	data.OpenInterest = float64(pDepthMarketData.OpenInterest)
	data.ClosePrice = float64(pDepthMarketData.ClosePrice)
	data.SettlementPrice = float64(pDepthMarketData.SettlementPrice)
	data.UpperLimitPrice = float64(pDepthMarketData.UpperLimitPrice)
	data.LowerLimitPrice = float64(pDepthMarketData.LowerLimitPrice)
	data.PreDelta = float64(pDepthMarketData.PreDelta)
	data.CurrDelta = float64(pDepthMarketData.CurrDelta)

	updateTime, _ := time.Parse("20060102 15:04:05.000", fmt.Sprintf("%s %s.%03d", C.GoString(&pDepthMarketData.CalendarDate[0]), C.GoString(&pDepthMarketData.UpdateTime[0]), int(pDepthMarketData.UpdateMillisec)))
	data.UpdateTime = updateTime

	data.InstrumentID = C.GoString(&pDepthMarketData.InstrumentID[0])

	data.Bids = make([]priceItem, 0, 5)
	data.Asks = make([]priceItem, 0, 5)

	if pDepthMarketData.BidVolume1 > 0 {
		data.Bids = append(data.Bids, priceItem{Price: float64(pDepthMarketData.BidPrice1), Volume: int(pDepthMarketData.BidVolume1)})

		if pDepthMarketData.BidVolume2 > 0 {
			data.Bids = append(data.Bids, priceItem{Price: float64(pDepthMarketData.BidPrice2), Volume: int(pDepthMarketData.BidVolume2)})

			if pDepthMarketData.BidVolume3 > 0 {
				data.Bids = append(data.Bids, priceItem{Price: float64(pDepthMarketData.BidPrice3), Volume: int(pDepthMarketData.BidVolume3)})

				if pDepthMarketData.BidVolume4 > 0 {
					data.Bids = append(data.Bids, priceItem{Price: float64(pDepthMarketData.BidPrice4), Volume: int(pDepthMarketData.BidVolume4)})

					if pDepthMarketData.BidVolume5 > 0 {
						data.Bids = append(data.Bids, priceItem{Price: float64(pDepthMarketData.BidPrice5), Volume: int(pDepthMarketData.BidVolume5)})
					}
				}
			}
		}
	}

	if pDepthMarketData.AskVolume1 > 0 {
		data.Asks = append(data.Asks, priceItem{Price: float64(pDepthMarketData.AskPrice1), Volume: int(pDepthMarketData.AskVolume1)})

		if pDepthMarketData.AskVolume2 > 0 {
			data.Asks = append(data.Asks, priceItem{Price: float64(pDepthMarketData.AskPrice2), Volume: int(pDepthMarketData.AskVolume2)})

			if pDepthMarketData.AskVolume3 > 0 {
				data.Asks = append(data.Asks, priceItem{Price: float64(pDepthMarketData.AskPrice3), Volume: int(pDepthMarketData.AskVolume3)})

				if pDepthMarketData.AskVolume4 > 0 {
					data.Asks = append(data.Asks, priceItem{Price: float64(pDepthMarketData.AskPrice4), Volume: int(pDepthMarketData.AskVolume4)})

					if pDepthMarketData.AskVolume5 > 0 {
						data.Asks = append(data.Asks, priceItem{Price: float64(pDepthMarketData.AskPrice5), Volume: int(pDepthMarketData.AskVolume5)})
					}
				}
			}
		}
	}

	data.FundingRate = float64(pDepthMarketData.FundingRate)
	data.FairPrice = float64(pDepthMarketData.FairPrice)

	return &data
}

func convertCQFMatchDisseminationField(pDissemination *C.struct_CQFMatchDisseminationField) *GoCQFMatchDisseminationField {
	data := GoCQFMatchDisseminationField{}

	data.SequenceNo = int(pDissemination.SequenceNo)
	data.SequenceSeries = int(pDissemination.SequenceSeries)

	return &data
}

func transformGoCQFMatchDisseminationField(dissemination *GoCQFMatchDisseminationField) *C.struct_CQFMatchDisseminationField {
	diss := C.struct_CQFMatchDisseminationField{}

	diss.SequenceSeries = C.TQFMatchSequenceSeriesType(dissemination.SequenceSeries)
	diss.SequenceNo = C.TQFMatchSequenceNoType(dissemination.SequenceNo)

	return &diss
}

func transformGoCQFMatchQryInstrumentField(qryInstrument *GoCQFMatchQryInstrumentField) *C.struct_CQFMatchQryInstrumentField {
	qry := C.struct_CQFMatchQryInstrumentField{}

	C.strncpy(&qry.SettlementGroupID[0], C.CString(qryInstrument.SettlementGroupID), C.sizeof_TQFMatchSettlementGroupIDType-1)
	C.strncpy(&qry.ProductGroupID[0], C.CString(qryInstrument.ProductGroupID), C.sizeof_TQFMatchProductGroupIDType-1)
	C.strncpy(&qry.ProductID[0], C.CString(qryInstrument.ProductID), C.sizeof_TQFMatchProductIDType-1)
	C.strncpy(&qry.InstrumentID[0], C.CString(qryInstrument.InstrumentID), C.sizeof_TQFMatchInstrumentIDType)

	return &qry
}

func transformGoCQFMatchQryInstrumentStatusField(qryInsStatus *GoCQFMatchQryInstrumentStatusField) *C.struct_CQFMatchQryInstrumentStatusField {
	qry := C.struct_CQFMatchQryInstrumentStatusField{}

	C.strncpy(&qry.InstIDStart[0], C.CString(qryInsStatus.InstIDStart), C.sizeof_TQFMatchInstrumentIDType-1)
	C.strncpy(&qry.InstIDEnd[0], C.CString(qryInsStatus.InstIDEnd), C.sizeof_TQFMatchInstrumentIDType-1)

	return &qry
}

func transformGoCQFMatchQryMarketDataField(qryMarketData *GoCQFMatchQryMarketDataField) *C.struct_CQFMatchQryMarketDataField {
	qry := C.struct_CQFMatchQryMarketDataField{}

	C.strncpy(&qry.ProductID[0], C.CString(qryMarketData.ProductID), C.sizeof_TQFMatchProductIDType-1)
	C.strncpy(&qry.InstrumentID[0], C.CString(qryMarketData.InstrumentID), C.sizeof_TQFMatchInstrumentIDType-1)

	return &qry
}

func transformGoCQFMatchQryMBLMarketDataField(qryMBLMarketData *GoCQFMatchQryMBLMarketDataField) *C.struct_CQFMatchQryMBLMarketDataField {
	qry := C.struct_CQFMatchQryMBLMarketDataField{}

	C.strncpy(&qry.InstIDStart[0], C.CString(qryMBLMarketData.InstIDStart), C.sizeof_TQFMatchInstrumentIDType-1)
	C.strncpy(&qry.InstIDEnd[0], C.CString(qryMBLMarketData.InstIDEnd), C.sizeof_TQFMatchInstrumentIDType-1)
	qry.Direction = ToCChar(qryMBLMarketData.Direction)

	return &qry
}

//export goOnRspSubscribeTopic
func goOnRspSubscribeTopic(instance C.QFMatchSuperApiInstance, pDissemination *C.struct_CQFMatchDisseminationField, pRspInfo *C.struct_CQFMatchRspInfoField, nRequestID C.int, bIsLast C.bool) {
	if pDissemination == nil {
		return
	}

	callback := getInstanceCallback(instance)

	if callback != nil {
		callback.(marketAPICallbacks).OnRspSubscribeTopic(convertDisseminationField(pDissemination), convertRspInfo(pRspInfo), int(nRequestID), bool(bIsLast))
	}
}

//export goOnRspQryMarketData
func goOnRspQryMarketData(instance C.QFMatchSuperApiInstance, pMarketData *C.struct_CQFMatchMarketDataField, pRspInfo *C.struct_CQFMatchRspInfoField, nRequestID C.int, bIsLast C.bool) {
	if pMarketData == nil {
		return
	}

	callback := getInstanceCallback(instance)

	if callback != nil {
		callback.(marketAPICallbacks).OnRspQryMarketData(convertMarketDataField(pMarketData), convertRspInfo(pRspInfo), int(nRequestID), bool(bIsLast))
	}
}

//export goOnRspQryInstrument
func goOnRspQryInstrument(instance C.QFMatchSuperApiInstance, pRspInstrument *C.struct_CQFMatchRspInstrumentField, pRspInfo *C.struct_CQFMatchRspInfoField, nRequestID C.int, bIsLast C.bool) {
	if pRspInstrument == nil {
		return
	}

	callback := getInstanceCallback(instance)

	if callback != nil {
		callback.(marketAPICallbacks).OnRspQryInstrument(convertRspInstrumentField(pRspInstrument), convertRspInfo(pRspInfo), int(nRequestID), bool(bIsLast))
	}
}

//export goOnRspQryInstrumentStatus
func goOnRspQryInstrumentStatus(instance C.QFMatchSuperApiInstance, pInstrumentStatus *C.struct_CQFMatchInstrumentStatusField, pRspInfo *C.struct_CQFMatchRspInfoField, nRequestID C.int, bIsLast C.bool) {
	if pInstrumentStatus == nil {
		return
	}

	callback := getInstanceCallback(instance)

	if callback != nil {
		callback.(marketAPICallbacks).OnRspQryInstrumentStatus(convertInstrumentStatusField(pInstrumentStatus), convertRspInfo(pRspInfo), int(nRequestID), bool(bIsLast))
	}
}

// // export goOnRtnInsInstrument
// func goOnRtnInsInstrument(instance C.QFMatchSuperApiInstance, pInstrument *C.struct_CQFMatchInstrumentField) {}

// //export goOnRtnDelInstrument
// func goOnRtnDelInstrument(instance C.QFMatchSuperApiInstance, pInstrument *C.struct_CQFMatchInstrumentField) {}

//export goOnRspQryTopic
func goOnRspQryTopic(instance C.QFMatchSuperApiInstance, pDissemination *C.struct_CQFMatchDisseminationField, pRspInfo *C.struct_CQFMatchRspInfoField, nRequestID C.int, bIsLast C.bool) {
	if pDissemination == nil {
		return
	}

	callback := getInstanceCallback(instance)

	if callback != nil {
		callback.(marketAPICallbacks).OnRspQryTopic(convertDisseminationField(pDissemination), convertRspInfo(pRspInfo), int(nRequestID), bool(bIsLast))
	}
}

//export goOnRspQryMBLMarketData
func goOnRspQryMBLMarketData(instance C.QFMatchSuperApiInstance, pMBLMarketData *C.struct_CQFMatchMBLMarketDataField, pRspInfo *C.struct_CQFMatchRspInfoField, nRequestID C.int, bIsLast C.bool) {
	if pMBLMarketData == nil {
		return
	}

	callback := getInstanceCallback(instance)

	if callback != nil {
		callback.(marketAPICallbacks).OnRspQryMBLMarketData(convertCQFMatchMBLMarketDataField(pMBLMarketData), convertRspInfo(pRspInfo), int(nRequestID), bool(bIsLast))
	}
}

//export goOnRtnInstrumentStatus
func goOnRtnInstrumentStatus(instance C.QFMatchSuperApiInstance, pInstrumentStatus *C.struct_CQFMatchInstrumentStatusField) {
	if pInstrumentStatus == nil {
		return
	}

	callback := getInstanceCallback(instance)

	if callback != nil {
		callback.(marketAPICallbacks).OnRtnInstrumentStatus(convertInstrumentStatusField(pInstrumentStatus))
	}
}

//export goOnRtnMarketData
func goOnRtnMarketData(instance C.QFMatchSuperApiInstance, pMarketData *C.struct_CQFMatchMarketDataField) {
	if pMarketData == nil {
		return
	}

	callback := getInstanceCallback(instance)

	if callback != nil {
		callback.(marketAPICallbacks).OnRtnMarketData(convertMarketDataField(pMarketData))
	}
}

//export goOnRtnDepthMarketData
func goOnRtnDepthMarketData(instance C.QFMatchSuperApiInstance, pDepthMarketData *C.struct_CQFMatchDepthMarketDataField) {
	if pDepthMarketData == nil {
		return
	}

	callback := getInstanceCallback(instance)

	if callback != nil {
		callback.(marketAPICallbacks).OnRtnDepthMarketData(convertDepthMarketDataField(pDepthMarketData))
	}
}

type marketAPI struct {
	common *commonAPI
}

// QFMatchMarketAPICallbacks Quantfair撮合引擎行情API回调接口
type QFMatchMarketAPICallbacks interface {
	commonAPICallbacks
	marketAPICallbacks
}

// QFMatchMarketAPI Quantfair撮合引擎行情API
type QFMatchMarketAPI struct {
	commonAPI
	marketAPI
}

func (api *marketAPI) registerMarketCallbacks(ptrVtCallbacks *C.Callbacks) {
	ptrVtCallbacks.ptrOnRspSubscribeTopic = C.FuncOnRspSubscribeTopic(C.cgoOnRspSubscribeTopic)
	ptrVtCallbacks.ptrOnRspQryMarketData = C.FuncOnRspQryMarketData(C.cgoOnRspQryMarketData)
	ptrVtCallbacks.ptrOnRspQryInstrument = C.FuncOnRspQryInstrument(C.cgoOnRspQryInstrument)
	ptrVtCallbacks.ptrOnRspQryInstrumentStatus = C.FuncOnRspQryInstrumentStatus(C.cgoOnRspQryInstrumentStatus)
	// ptrVtCallbacks.ptrOnRtnInsInstrument = C.FuncOnRtnInsInstrument(C.cgoOnRtnInsInstrument)
	// ptrVtCallbacks.ptrOnRtnDelInstrument = C.FuncOnRtnDelInstrument(C.cgoOnRtnDelInstrument)
	ptrVtCallbacks.ptrOnRspQryTopic = C.FuncOnRspQryTopic(C.cgoOnRspQryTopic)
	ptrVtCallbacks.ptrOnRspQryMBLMarketData = C.FuncOnRspQryMBLMarketData(C.cgoOnRspQryMBLMarketData)
	ptrVtCallbacks.ptrOnRtnInstrumentStatus = C.FuncOnRtnInstrumentStatus(C.cgoOnRtnInstrumentStatus)
	ptrVtCallbacks.ptrOnRtnMarketData = C.FuncOnRtnMarketData(C.cgoOnRtnMarketData)
	ptrVtCallbacks.ptrOnRtnDepthMarketData = C.FuncOnRtnDepthMarketData(C.cgoOnRtnDepthMarketData)
}

// RegisterCallbacks 注册回调接口
func (api *QFMatchMarketAPI) RegisterCallbacks(cb QFMatchMarketAPICallbacks) {
	vtCallbacks := C.Callbacks{}

	api.registerCommonCallbacks(&vtCallbacks)
	api.registerMarketCallbacks(&vtCallbacks)
	api.marketAPI.common = &api.commonAPI

	C.RegisterCallbacks(C.QFMatchSuperApiInstance(api.apiInstance), &vtCallbacks)

	apiMap[api.apiInstance] = cb
}

// SubscribeMarketDataTopic 订阅市场行情
func (api *marketAPI) SubscribeMarketDataTopic(topicID int, resumeType ResumeType) {
	C.SubscribeMarketDataTopic(C.QFMatchSuperApiInstance(api.common.apiInstance), C.int(topicID), C.enum_QFMATCH_TE_RESUME_TYPE(resumeType))
}

func (api *marketAPI) ReqSubscribeTopic(dissemination *GoCQFMatchDisseminationField) int {
	rtn := C.ReqSubscribeTopic(C.QFMatchSuperApiInstance(api.common.apiInstance), transformGoCQFMatchDisseminationField(dissemination), C.int(api.common.getRequestID()))

	return int(rtn)
}

// ReqQryInstrument 合约查询请求
func (api *marketAPI) ReqQryInstrument(qryInstrument *GoCQFMatchQryInstrumentField) int {
	rtn := C.ReqQryInstrument(
		C.QFMatchSuperApiInstance(api.common.apiInstance),
		transformGoCQFMatchQryInstrumentField(qryInstrument),
		C.int(api.common.getRequestID()))

	return int(rtn)
}

// ReqQryInstrumentStatus 合约交易状态查询请求
func (api *marketAPI) ReqQryInstrumentStatus(qryInsStatus *GoCQFMatchQryInstrumentStatusField) int {
	rtn := C.ReqQryInstrumentStatus(C.QFMatchSuperApiInstance(api.common.apiInstance), transformGoCQFMatchQryInstrumentStatusField(qryInsStatus), C.int(api.common.getRequestID()))

	return int(rtn)
}

// ReqQryMarketData 普通行情查询请求
func (api *marketAPI) ReqQryMarketData(qryMarketData *GoCQFMatchQryMarketDataField) int {
	rtn := C.ReqQryMarketData(C.QFMatchSuperApiInstance(api.common.apiInstance), transformGoCQFMatchQryMarketDataField(qryMarketData), C.int(api.common.getRequestID()))

	return int(rtn)
}

// ReqQryTopic 主题查询请求
func (api *marketAPI) ReqQryTopic(dissemination *GoCQFMatchDisseminationField) int {
	rtn := C.ReqQryTopic(C.QFMatchSuperApiInstance(api.common.apiInstance), transformGoCQFMatchDisseminationField(dissemination), C.int(api.common.getRequestID()))

	return int(rtn)
}

// ReqQryMBLMarketData 合约价位查询
func (api *marketAPI) ReqQryMBLMarketData(qryMBLMarketData *GoCQFMatchQryMBLMarketDataField) int {
	rtn := C.ReqQryMBLMarketData(C.QFMatchSuperApiInstance(api.common.apiInstance), transformGoCQFMatchQryMBLMarketDataField(qryMBLMarketData), C.int(api.common.getRequestID()))

	return int(rtn)
}

// OnRspSubscribeTopic 主题订阅回报
func (api *marketAPI) OnRspSubscribeTopic(rspSubscribeTopic *GoCQFMatchDisseminationField, err *GoCQFMatchRspInfoField, requestID int, isLast bool) {
	if err != nil && err.ErrorID != 0 {
		log.Printf("Faild on subscribe topic: [%d]%s\n", err.ErrorID, err.ErrorMsg)
	} else {
		log.Printf("Subscribe topic[%d] succeded.\n", rspSubscribeTopic.SequenceSeries)
	}
}

func showMarketData(marketData *GoCQFMatchMarketDataField) {
	log.Printf("[%s-%s] %s: H[%f], O[%f], L[%f], C[%f], LMIT[%f], STOP[%f]\n", marketData.TradingDay, marketData.UpdateTime.Format("15:04:05"), marketData.InstrumentID, marketData.HighestPrice, marketData.OpenInterest, marketData.LowestPrice, marketData.ClosePrice, marketData.UpperLimitPrice, marketData.LowerLimitPrice)
}

// OnRspQryMarketData 行情查询回报
func (api *marketAPI) OnRspQryMarketData(marketData *GoCQFMatchMarketDataField, err *GoCQFMatchRspInfoField, requestID int, isLast bool) {
	if err != nil && err.ErrorID != 0 {
		log.Printf("Market data query failed: [%d]%s\n", err.ErrorID, err.ErrorMsg)
	} else {
		showMarketData(marketData)
	}
}

// OnRspQryInstrument 合约查询回报
func (api *marketAPI) OnRspQryInstrument(instrument *GoCQFMatchRspInstrumentField, err *GoCQFMatchRspInfoField, requestID int, isLast bool) {
	if err != nil && err.ErrorID != 0 {
		log.Printf("Instrument query failed: [%d]%s\n", err.ErrorID, err.ErrorMsg)
	} else {
		switch instrument.ProductClass {
		case PCFuture:
			log.Printf("[%s.%s] %s: PrdID[%s], Tick[%f], xVol[%d]\n", instrument.ProductGroupID, instrument.InstrumentID, instrument.InstrumentName, instrument.ProductID, instrument.PriceTick, instrument.VolumeMultiple)
		case PCOptions:
			log.Printf("[%s.%s] %s: PrdID[%s], Tick[%f], xVol[%d], Udr[%s]\n", instrument.ProductGroupID, instrument.InstrumentID, instrument.InstrumentName, instrument.ProductID, instrument.PriceTick, instrument.VolumeMultiple, instrument.UnderlyingInstrID)
		}

		if isLast {
			log.Println("Instruments query finished.")
		}
	}
}

func showInstrumentStatus(insStatus *GoCQFMatchInstrumentStatusField) {
	log.Printf("[%s]: %s enter %s by %s\n", insStatus.InstrumentID, insStatus.EnterTime.Format("2006-01-02 15:04:05"), insStatus.InstrumentStatus.GetName(), insStatus.EnterReason.GetName())
}

// OnRspQryInstrumentStatus 合约状态查询回报
func (api *marketAPI) OnRspQryInstrumentStatus(insStatus *GoCQFMatchInstrumentStatusField, err *GoCQFMatchRspInfoField, requestID int, isLast bool) {
	if err != nil && err.ErrorID != 0 {
		log.Printf("Instrument status query failed: [%d]%s\n", err.ErrorID, err.ErrorMsg)
	} else {
		showInstrumentStatus(insStatus)

		if isLast {
			log.Println("Instrument status query finished.")
		}
	}
}

// OnRspQryTopic 主题查询应答
func (api *marketAPI) OnRspQryTopic(dissemination *GoCQFMatchDisseminationField, err *GoCQFMatchRspInfoField, requestID int, isLast bool) {
	if err != nil && err.ErrorID != 0 {
		log.Printf("Topic query failed: [%d]%s\n", err.ErrorID, err.ErrorMsg)
	} else {
		log.Printf("Available topic: %d, %d\n", dissemination.SequenceNo, dissemination.SequenceSeries)

		if isLast {
			log.Println("Topic query finished.")
		}
	}
}

// OnRspQryMBLMarketData 合约价位查询应答
func (api *marketAPI) OnRspQryMBLMarketData(mblMarketData *GoCQFMatchMBLMarketDataField, err *GoCQFMatchRspInfoField, requestID int, isLast bool) {
	if err != nil && err.ErrorID != 0 {
		log.Printf("MBLMarketData query failed: [%d]%s\n", err.ErrorID, err.ErrorMsg)
	} else {
		log.Printf("[%s] %s: P[%f], V[%d]\n", mblMarketData.InstrumentID, mblMarketData.Direction.GetName(), mblMarketData.Price, mblMarketData.Volume)

		if isLast {
			log.Println("MBLMarketData query finished.")
		}
	}
}

// OnRtnInstrumentStatus 合约状态推送
func (api *marketAPI) OnRtnInstrumentStatus(insStatus *GoCQFMatchInstrumentStatusField) {
	showInstrumentStatus(insStatus)
}

// OnRtnMarketData 1档行情推送
func (api *marketAPI) OnRtnMarketData(marketData *GoCQFMatchMarketDataField) {
	showMarketData(marketData)
}

func (api *marketAPI) OnRtnDepthMarketData(depthMarketData *GoCQFMatchDepthMarketDataField) {
	var ask, bid, last string

	if len(depthMarketData.Asks) > 0 {
		ask = fmt.Sprintf("Ask[%d@%f]", depthMarketData.Asks[0].Volume, depthMarketData.Asks[0].Price)
	} else {
		ask = "Ask[]"
	}

	if len(depthMarketData.Bids) > 0 {
		bid = fmt.Sprintf("Bid[%d@%f]", depthMarketData.Bids[0].Volume, depthMarketData.Bids[0].Price)
	} else {
		bid = "Bid[]"
	}

	if depthMarketData.Volume > 0 {
		last = fmt.Sprintf("Last[%d@%f]", depthMarketData.Volume, depthMarketData.LastPrice)
	} else {
		last = "Last[]"
	}

	log.Printf("[%-6s] %s: %s, %s, %s\n", depthMarketData.InstrumentID, depthMarketData.UpdateTime.Format("15:04:05.000"), ask, bid, last)
}
