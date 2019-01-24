#include "cgoTraderApi.h"

void cgoOnRspOrderInsert(QFMatchSuperApiInstance instance, struct CQFMatchInputOrderField *pInputOrder, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    goOnRspOrderInsert(instance, pInputOrder, pRspInfo, nRequestID, bIsLast);
}

void cgoOnRspOrderAction(QFMatchSuperApiInstance instance, struct CQFMatchOrderActionField *pOrderAction, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    goOnRspOrderAction(instance, pOrderAction, pRspInfo, nRequestID, bIsLast);
}

// void cgoOnRspQryPartAccount(QFMatchSuperApiInstance instance, struct CQFMatchRspPartAccountField *pRspPartAccount, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
//     goOnRspQryPartAccount(instance, pRspPartAccount, pRspInfo, nRequestID, bIsLast);
// }

void cgoOnRspQryOrder(QFMatchSuperApiInstance instance, struct CQFMatchOrderField *pOrder, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    goOnRspQryOrder(instance, pOrder, pRspInfo, nRequestID, bIsLast);
}

void cgoOnRspQryTrade(QFMatchSuperApiInstance instance, struct CQFMatchTradeField *pTrade, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    goOnRspQryTrade(instance, pTrade, pRspInfo, nRequestID, bIsLast);
}

// void cgoOnRspQryClient(QFMatchSuperApiInstance instance, struct CQFMatchRspClientField *pRspClient, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
//     goOnRspQryClient(instance, pRspClient, pRspInfo, nRequestID, bIsLast);
// }

// void cgoOnRspQryClientPosition(QFMatchSuperApiInstance instance, struct CQFMatchRspClientPositionField *pRspClientPosition, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
//     goOnRspQryClientPosition(instance, pRspClientPosition, pRspInfo, nRequestID, bIsLast);
// }

void cgoOnRtnTrade(QFMatchSuperApiInstance instance, struct CQFMatchTradeField *pTrade) {
    goOnRtnTrade(instance, pTrade);
}

void cgoOnRtnOrder(QFMatchSuperApiInstance instance, struct CQFMatchOrderField *pOrder) {
    goOnRtnOrder(instance, pOrder);
}

// void cgoOnRtnClientPosition(QFMatchSuperApiInstance instance, struct CQFMatchClientPositionField *pClientPosition) {
//     OnRtnClientPosition(instance, pClientPosition);
// }

// void cgoOnRtnFlowMessageCancel(QFMatchSuperApiInstance instance, struct CQFMatchFlowMessageCancelField *pFlowMessageCancel) {
//     goOnRtnFlowMessageCancel(instance, pFlowMessageCancel);
// }

// void cgoOnErrRtnOrderInsert(QFMatchSuperApiInstance instance, struct CQFMatchInputOrderField *pInputOrder, struct CQFMatchRspInfoField *pRspInfo) {
//     goOnErrRtnOrderInsert(instance, pInputOrder, pRspInfo);
// }

// void cgoOnErrRtnOrderAction(QFMatchSuperApiInstance instance, struct CQFMatchOrderActionField *pOrderAction, struct CQFMatchRspInfoField *pRspInfo) {
//     goOnErrRtnOrderAction(instance, pOrderAction, pRspInfo);
// }

// void cgoOnRspQryPartClientFee(QFMatchSuperApiInstance instance, struct CQFMatchPartClientFeeField *pPartClientFee, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
//     goOnRspQryPartClientFee(instance, pPartClientFee, pRspInfo, nRequestID, bIsLast);
// }