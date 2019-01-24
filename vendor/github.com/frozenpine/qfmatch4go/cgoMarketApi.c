#include "cgoMarketApi.h"

void cgoOnRspSubscribeTopic(QFMatchSuperApiInstance instance, struct CQFMatchDisseminationField *pDissemination, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    goOnRspSubscribeTopic(instance, pDissemination, pRspInfo, nRequestID, bIsLast);
}

void cgoOnRspQryMarketData(QFMatchSuperApiInstance instance, struct CQFMatchMarketDataField *pMarketData, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    goOnRspQryMarketData(instance, pMarketData, pRspInfo, nRequestID, bIsLast);
}

void cgoOnRspQryInstrument(QFMatchSuperApiInstance instance, struct CQFMatchRspInstrumentField *pRspInstrument, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    goOnRspQryInstrument(instance, pRspInstrument, pRspInfo, nRequestID, bIsLast);
}

void cgoOnRspQryInstrumentStatus(QFMatchSuperApiInstance instance, struct CQFMatchInstrumentStatusField *pInstrumentStatus, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    goOnRspQryInstrumentStatus(instance, pInstrumentStatus, pRspInfo, nRequestID, bIsLast);
}

// void cgoOnRtnInsInstrument(QFMatchSuperApiInstance instance, struct CQFMatchInstrumentField *pInstrument) {
//     goOnRtnInsInstrument(instance, pInstrument);
// }

// void cgoOnRtnDelInstrument(QFMatchSuperApiInstance instance, struct CQFMatchInstrumentField *pInstrument) {
//     goOnRtnDelInstrument(instance, pInstrument);
// }

void cgoOnRspQryTopic(QFMatchSuperApiInstance instance, struct CQFMatchDisseminationField *pDissemination, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    goOnRspQryTopic(instance, pDissemination, pRspInfo, nRequestID, bIsLast);
}

void cgoOnRspQryMBLMarketData(QFMatchSuperApiInstance instance, struct CQFMatchMBLMarketDataField *pMBLMarketData, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    goOnRspQryMBLMarketData(instance, pMBLMarketData, pRspInfo, nRequestID, bIsLast);
}

void cgoOnRtnInstrumentStatus(QFMatchSuperApiInstance instance, struct CQFMatchInstrumentStatusField *pInstrumentStatus) {
    goOnRtnInstrumentStatus(instance, pInstrumentStatus);
}

void cgoOnRtnMarketData(QFMatchSuperApiInstance instance, struct CQFMatchMarketDataField *pMarketData) {
    goOnRtnMarketData(instance, pMarketData);
}

void cgoOnRtnDepthMarketData(QFMatchSuperApiInstance instance, struct CQFMatchDepthMarketDataField *pDepthMarketData) {
goOnRtnDepthMarketData(instance, pDepthMarketData);
}