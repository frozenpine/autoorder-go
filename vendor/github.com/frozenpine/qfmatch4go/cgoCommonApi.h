#ifndef CGOCOMMON_API_H
#define CGOCOMMON_API_H

#include "cQFMatchSuperApi.h"

/*
 *  common api definitions
 */
void cgoOnFrontConnected(QFMatchSuperApiInstance instance);
extern void goOnFrontConnected(QFMatchSuperApiInstance instance);

void cgoOnFrontDisconnected(QFMatchSuperApiInstance instance, int nReason);
extern void goOnFrontDisconnected(QFMatchSuperApiInstance instance, int nReason);

void cgoOnHeartBeatWarning(QFMatchSuperApiInstance instance, int nTime);
extern void goOnHeartBeatWarning(QFMatchSuperApiInstance instance, int nTime);

void cgoOnPackageStart(QFMatchSuperApiInstance instance, int nTopicID, int nSequenceNO);
extern void goOnPackageStart(QFMatchSuperApiInstance instance, int nTopicID, int nSequenceNO);

void cgoOnPackageEnd(QFMatchSuperApiInstance instance, int nTopicID, int nSequenceNO);
extern void goOnPackageEnd(QFMatchSuperApiInstance instance, int nTopicID, int nSequenceNO);

void cgoOnRspError(QFMatchSuperApiInstance instance, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
extern void goOnRspError(QFMatchSuperApiInstance instance, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

void cgoOnRspUserLogin(QFMatchSuperApiInstance instance, struct CQFMatchRspUserLoginField *pRspUserLogin, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
extern void goOnRspUserLogin(QFMatchSuperApiInstance instance, struct CQFMatchRspUserLoginField *pRspUserLogin, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

void cgoOnRspUserLogout(QFMatchSuperApiInstance instance, struct CQFMatchRspUserLogoutField *pRspUserLogout, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
extern void goOnRspUserLogout(QFMatchSuperApiInstance instance, struct CQFMatchRspUserLogoutField *pRspUserLogout, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

void cgoOnRspQryBulletin(QFMatchSuperApiInstance instance, struct CQFMatchBulletinField *pBulletin, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
extern void goOnRspQryBulletin(QFMatchSuperApiInstance instance, struct CQFMatchBulletinField *pBulletin, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

void cgoOnRtnBulletin(QFMatchSuperApiInstance instance, struct CQFMatchBulletinField *pBulletin);
extern void goOnRtnBulletin(QFMatchSuperApiInstance instance, struct CQFMatchBulletinField *pBulletin);

#endif