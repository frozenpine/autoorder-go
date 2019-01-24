#include "cgoCommonApi.h"

void cgoOnFrontConnected(QFMatchSuperApiInstance instance) {
    goOnFrontConnected(instance);
}

void cgoOnFrontDisconnected(QFMatchSuperApiInstance instance, int nReason) {
    goOnFrontDisconnected(instance, nReason);
}

void cgoOnHeartBeatWarning(QFMatchSuperApiInstance instance, int nTime) {
    goOnHeartBeatWarning(instance, nTime);
}

void cgoOnPackageStart(QFMatchSuperApiInstance instance, int nTopicID, int nSequenceNO) {
    goOnPackageStart(instance, nTopicID, nSequenceNO);
}

void cgoOnPackageEnd(QFMatchSuperApiInstance instance, int nTopicID, int nSequenceNO) {
    goOnPackageEnd(instance, nTopicID, nSequenceNO);
}

void cgoOnRspError(QFMatchSuperApiInstance instance, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    goOnRspError(instance, pRspInfo, nRequestID, bIsLast);
}

void cgoOnRspUserLogin(QFMatchSuperApiInstance instance, struct CQFMatchRspUserLoginField *pRspUserLogin, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    goOnRspUserLogin(instance, pRspUserLogin, pRspInfo, nRequestID, bIsLast);
}

void cgoOnRspUserLogout(QFMatchSuperApiInstance instance, struct CQFMatchRspUserLogoutField *pRspUserLogout, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    goOnRspUserLogout(instance, pRspUserLogout, pRspInfo, nRequestID, bIsLast);
}

void cgoOnRspQryBulletin(QFMatchSuperApiInstance instance, struct CQFMatchBulletinField *pBulletin, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
    goOnRspQryBulletin(instance, pBulletin, pRspInfo, nRequestID, bIsLast);
}

void cgoOnRtnBulletin(QFMatchSuperApiInstance instance, struct CQFMatchBulletinField *pBulletin){
    goOnRtnBulletin(instance, pBulletin);
}
