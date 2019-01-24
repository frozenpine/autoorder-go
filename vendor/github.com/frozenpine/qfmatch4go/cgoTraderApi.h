#ifndef CGOTRADER_API_H
#define CGOTRADER_API_H

#include "cQFMatchSuperApi.h"

/*
 *  trader api definitions
 */

void cgoOnRspOrderInsert(QFMatchSuperApiInstance instance, struct CQFMatchInputOrderField *pInputOrder, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
extern void goOnRspOrderInsert(QFMatchSuperApiInstance instance, struct CQFMatchInputOrderField *pInputOrder, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

void cgoOnRspOrderAction(QFMatchSuperApiInstance instance, struct CQFMatchOrderActionField *pOrderAction, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
extern void goOnRspOrderAction(QFMatchSuperApiInstance instance, struct CQFMatchOrderActionField *pOrderAction, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

// void cgoOnRspQryPartAccount(QFMatchSuperApiInstance instance, struct CQFMatchRspPartAccountField *pRspPartAccount, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
// extern void goOnRspQryPartAccount(QFMatchSuperApiInstance instance, struct CQFMatchRspPartAccountField *pRspPartAccount, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

void cgoOnRspQryOrder(QFMatchSuperApiInstance instance, struct CQFMatchOrderField *pOrder, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
extern void goOnRspQryOrder(QFMatchSuperApiInstance instance, struct CQFMatchOrderField *pOrder, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

void cgoOnRspQryTrade(QFMatchSuperApiInstance instance, struct CQFMatchTradeField *pTrade, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
extern void goOnRspQryTrade(QFMatchSuperApiInstance instance, struct CQFMatchTradeField *pTrade, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

// void cgoOnRspQryClient(QFMatchSuperApiInstance instance, struct CQFMatchRspClientField *pRspClient, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
// extern void goOnRspQryClient(QFMatchSuperApiInstance instance, struct CQFMatchRspClientField *pRspClient, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

// void cgoOnRspQryClientPosition(QFMatchSuperApiInstance instance, struct CQFMatchRspClientPositionField *pRspClientPosition, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
// extern void goOnRspQryClientPosition(QFMatchSuperApiInstance instance, struct CQFMatchRspClientPositionField *pRspClientPosition, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

void cgoOnRtnTrade(QFMatchSuperApiInstance instance, struct CQFMatchTradeField *pTrade);
extern void goOnRtnTrade(QFMatchSuperApiInstance instance, struct CQFMatchTradeField *pTrade);

void cgoOnRtnOrder(QFMatchSuperApiInstance instance, struct CQFMatchOrderField *pOrder);
extern void goOnRtnOrder(QFMatchSuperApiInstance instance, struct CQFMatchOrderField *pOrder);

// void cgoOnRtnClientPosition(QFMatchSuperApiInstance instance, struct CQFMatchClientPositionField *pClientPosition);
// extern void OnRtnClientPosition(QFMatchSuperApiInstance instance, struct CQFMatchClientPositionField *pClientPosition);

// void cgoOnRtnFlowMessageCancel(QFMatchSuperApiInstance instance, struct CQFMatchFlowMessageCancelField *pFlowMessageCancel);
// extern void goOnRtnFlowMessageCancel(QFMatchSuperApiInstance instance, struct CQFMatchFlowMessageCancelField *pFlowMessageCancel);

// void cgoOnErrRtnOrderInsert(QFMatchSuperApiInstance instance, struct CQFMatchInputOrderField *pInputOrder, struct CQFMatchRspInfoField *pRspInfo);
// extern void goOnErrRtnOrderInsert(QFMatchSuperApiInstance instance, struct CQFMatchInputOrderField *pInputOrder, struct CQFMatchRspInfoField *pRspInfo);

// void cgoOnErrRtnOrderAction(QFMatchSuperApiInstance instance, struct CQFMatchOrderActionField *pOrderAction, struct CQFMatchRspInfoField *pRspInfo);
// extern void goOnErrRtnOrderAction(QFMatchSuperApiInstance instance, struct CQFMatchOrderActionField *pOrderAction, struct CQFMatchRspInfoField *pRspInfo);

// void cgoOnRspQryPartClientFee(QFMatchSuperApiInstance instance, struct CQFMatchPartClientFeeField *pPartClientFee, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
// extern void goOnRspQryPartClientFee(QFMatchSuperApiInstance instance, struct CQFMatchPartClientFeeField *pPartClientFee, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

#endif