#ifndef CGOMARKET_API_H
#define CGOMARKET_API_H

#include "cQFMatchSuperApi.h"

/*
 *  market api definitions
 */

void cgoOnRspSubscribeTopic(QFMatchSuperApiInstance instance, struct CQFMatchDisseminationField *pDissemination, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
extern void goOnRspSubscribeTopic(QFMatchSuperApiInstance instance, struct CQFMatchDisseminationField *pDissemination, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

void cgoOnRspQryMarketData(QFMatchSuperApiInstance instance, struct CQFMatchMarketDataField *pMarketData, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
extern void goOnRspQryMarketData(QFMatchSuperApiInstance instance, struct CQFMatchMarketDataField *pMarketData, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

void cgoOnRspQryInstrument(QFMatchSuperApiInstance instance, struct CQFMatchRspInstrumentField *pRspInstrument, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
extern void goOnRspQryInstrument(QFMatchSuperApiInstance instance, struct CQFMatchRspInstrumentField *pRspInstrument, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

void cgoOnRspQryInstrumentStatus(QFMatchSuperApiInstance instance, struct CQFMatchInstrumentStatusField *pInstrumentStatus, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
extern void goOnRspQryInstrumentStatus(QFMatchSuperApiInstance instance, struct CQFMatchInstrumentStatusField *pInstrumentStatus, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

// void cgoOnRtnInsInstrument(QFMatchSuperApiInstance instance, struct CQFMatchInstrumentField *pInstrument);
// extern void goOnRtnInsInstrument(QFMatchSuperApiInstance instance, struct CQFMatchInstrumentField *pInstrument);

// void cgoOnRtnDelInstrument(QFMatchSuperApiInstance instance, struct CQFMatchInstrumentField *pInstrument);
// extern void goOnRtnDelInstrument(QFMatchSuperApiInstance instance, struct CQFMatchInstrumentField *pInstrument);

void cgoOnRspQryTopic(QFMatchSuperApiInstance instance, struct CQFMatchDisseminationField *pDissemination, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
extern void goOnRspQryTopic(QFMatchSuperApiInstance instance, struct CQFMatchDisseminationField *pDissemination, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

void cgoOnRspQryMBLMarketData(QFMatchSuperApiInstance instance, struct CQFMatchMBLMarketDataField *pMBLMarketData, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
extern void goOnRspQryMBLMarketData(QFMatchSuperApiInstance instance, struct CQFMatchMBLMarketDataField *pMBLMarketData, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

void cgoOnRtnInstrumentStatus(QFMatchSuperApiInstance instance, struct CQFMatchInstrumentStatusField *pInstrumentStatus);
extern void goOnRtnInstrumentStatus(QFMatchSuperApiInstance instance, struct CQFMatchInstrumentStatusField *pInstrumentStatus);

void cgoOnRtnMarketData(QFMatchSuperApiInstance instance, struct CQFMatchMarketDataField *pMarketData);
extern void goOnRtnMarketData(QFMatchSuperApiInstance instance, struct CQFMatchMarketDataField *pMarketData);

void cgoOnRtnDepthMarketData(QFMatchSuperApiInstance instance, struct CQFMatchDepthMarketDataField *pDepthMarketData);
extern void goOnRtnDepthMarketData(QFMatchSuperApiInstance instance, struct CQFMatchDepthMarketDataField *pDepthMarketData);

#endif