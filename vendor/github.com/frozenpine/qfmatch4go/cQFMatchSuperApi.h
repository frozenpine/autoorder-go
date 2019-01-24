/////////////////////////////////////////////////////////////////////////
///@file cQFMatchSuperApi.h
///@brief 定义了超级用户API的C接口
/////////////////////////////////////////////////////////////////////////
#ifdef __cplusplus
extern "C" {
#endif

#ifndef CQFMATCH_SUPERAPI_H
#define CQFMATCH_SUPERAPI_H

#include <stdbool.h>
#include "QFMatchApiStruct.h"

#if defined(WINDOWS) && !defined(WIN32)
#define WIN32
#endif

#if defined(CQFMATCH_SUPERAPI_IS_LIB) && defined(WIN32)
#ifdef LIB_CQFMATCH_SUPER_API_EXPORT
#define CQFMATCH_SUPER_API_EXPORT __declspec(dllexport)
#else
#define CQFMATCH_SUPER_API_EXPORT __declspec(dllimport)
#endif
#else
#define CQFMATCH_SUPER_API_EXPORT __attribute__ ((dllimport))
#endif

/////////////////////////////////////////////////
///
/// C封装C++ API类容器指针
///
////////////////////////////////////////////////

typedef unsigned int QFMatchSuperApiInstance;

///////////////////////////////////////////////////
///
/// 定义回调函数指针
///
///////////////////////////////////////////////////

typedef void (*FuncOnFrontConnected)(QFMatchSuperApiInstance instance);
typedef void (*FuncOnFrontDisconnected)(QFMatchSuperApiInstance instance, int nReason);
typedef void (*FuncOnHeartBeatWarning)(QFMatchSuperApiInstance instance, int nTimeLapse);
typedef void (*FuncOnPackageStart)(QFMatchSuperApiInstance instance, int nTopicID, int nSequenceNo);
typedef void (*FuncOnPackageEnd)(QFMatchSuperApiInstance instance, int nTopicID, int nSequenceNo);
///错误应答
typedef void (*FuncOnRspError)(QFMatchSuperApiInstance instance, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///用户登录应答
typedef void (*FuncOnRspUserLogin)(QFMatchSuperApiInstance instance, struct CQFMatchRspUserLoginField *pRspUserLogin, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///用户退出应答
typedef void (*FuncOnRspUserLogout)(QFMatchSuperApiInstance instance, struct CQFMatchRspUserLogoutField *pRspUserLogout, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///报单录入应答
typedef void (*FuncOnRspOrderInsert)(QFMatchSuperApiInstance instance, struct CQFMatchInputOrderField *pInputOrder, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///报单操作应答
typedef void (*FuncOnRspOrderAction)(QFMatchSuperApiInstance instance, struct CQFMatchOrderActionField *pOrderAction, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///报价录入应答
typedef void (*FuncOnRspQuoteInsert)(QFMatchSuperApiInstance instance, struct CQFMatchInputQuoteField *pInputQuote, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///报价操作应答
typedef void (*FuncOnRspQuoteAction)(QFMatchSuperApiInstance instance, struct CQFMatchQuoteActionField *pQuoteAction, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///OTC报单录入应答
typedef void (*FuncOnRspOTCOrderInsert)(QFMatchSuperApiInstance instance, struct CQFMatchOTCOrderField *pOTCOrder, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///用户密码修改应答
typedef void (*FuncOnRspUserPasswordUpdate)(QFMatchSuperApiInstance instance, struct CQFMatchUserPasswordUpdateField *pUserPasswordUpdate, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///执行宣告录入应答
typedef void (*FuncOnRspExecOrderInsert)(QFMatchSuperApiInstance instance, struct CQFMatchInputExecOrderField *pInputExecOrder, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///执行宣告操作应答
typedef void (*FuncOnRspExecOrderAction)(QFMatchSuperApiInstance instance, struct CQFMatchExecOrderActionField *pExecOrderAction, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///管理报单录入应答
typedef void (*FuncOnRspAdminOrderInsert)(QFMatchSuperApiInstance instance, struct CQFMatchInputAdminOrderField *pInputAdminOrder, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///非标组合报单录入应答
typedef void (*FuncOnRspCombOrderInsert)(QFMatchSuperApiInstance instance, struct CQFMatchInputCombOrderField *pInputCombOrder, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///订阅主题应答
typedef void (*FuncOnRspSubscribeTopic)(QFMatchSuperApiInstance instance, struct CQFMatchDisseminationField *pDissemination, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///交易所公告请求应答
typedef void (*FuncOnRspBulletin)(QFMatchSuperApiInstance instance, struct CQFMatchBulletinField *pBulletin, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///交易所公告信息应答
typedef void (*FuncOnRspInformation)(QFMatchSuperApiInstance instance, struct CQFMatchInformationField *pInformation, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///合约交易状态改变应答
typedef void (*FuncOnRspInstrumentStatusUpdate)(QFMatchSuperApiInstance instance, struct CQFMatchInstrumentStatusField *pInstrumentStatus, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///强制用户退出应答
typedef void (*FuncOnRspForceUserExit)(QFMatchSuperApiInstance instance, struct CQFMatchForceUserExitField *pForceUserExit, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///用户会话删除应答
typedef void (*FuncOnRspForceUserLogout)(QFMatchSuperApiInstance instance, struct CQFMatchForceUserExitField *pForceUserExit, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///交易所数据同步状态改变应答
typedef void (*FuncOnRspExchangeDataSyncStatusUpdate)(QFMatchSuperApiInstance instance, struct CQFMatchExchangeDataSyncStatusField *pExchangeDataSyncStatus, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///结算组数据同步状态改变应答
typedef void (*FuncOnRspSGDataSyncStatusUpdate)(QFMatchSuperApiInstance instance, struct CQFMatchSGDataSyncStatusField *pSGDataSyncStatus, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///实时上场应答
typedef void (*FuncOnRspTBCommand)(QFMatchSuperApiInstance instance, struct CQFMatchTBCommandField *pTBCommand, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///更新行情应答
typedef void (*FuncOnRspUpdateMarketData)(QFMatchSuperApiInstance instance, struct CQFMatchUpdateMarketDataField *pUpdateMarketData, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///仓位保证金操作应答
typedef void (*FuncOnRspPositionMarginAction)(QFMatchSuperApiInstance instance, struct CQFMatchPositionMarginActionField *pPositionMarginAction, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///会员资金查询应答
typedef void (*FuncOnRspQryPartAccount)(QFMatchSuperApiInstance instance, struct CQFMatchRspPartAccountField *pRspPartAccount, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///报单查询应答
typedef void (*FuncOnRspQryOrder)(QFMatchSuperApiInstance instance, struct CQFMatchOrderField *pOrder, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///报价查询应答
typedef void (*FuncOnRspQryQuote)(QFMatchSuperApiInstance instance, struct CQFMatchQuoteField *pQuote, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///成交单查询应答
typedef void (*FuncOnRspQryTrade)(QFMatchSuperApiInstance instance, struct CQFMatchTradeField *pTrade, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///会员客户查询应答
typedef void (*FuncOnRspQryClient)(QFMatchSuperApiInstance instance, struct CQFMatchRspClientField *pRspClient, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///会员持仓查询应答
typedef void (*FuncOnRspQryPartPosition)(QFMatchSuperApiInstance instance, struct CQFMatchRspPartPositionField *pRspPartPosition, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///客户持仓查询应答
typedef void (*FuncOnRspQryClientPosition)(QFMatchSuperApiInstance instance, struct CQFMatchRspClientPositionField *pRspClientPosition, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///合约查询应答
typedef void (*FuncOnRspQryInstrument)(QFMatchSuperApiInstance instance, struct CQFMatchRspInstrumentField *pRspInstrument, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///合约交易状态查询应答
typedef void (*FuncOnRspQryInstrumentStatus)(QFMatchSuperApiInstance instance, struct CQFMatchInstrumentStatusField *pInstrumentStatus, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///结算组状态查询应答
typedef void (*FuncOnRspQrySGDataSyncStatus)(QFMatchSuperApiInstance instance, struct CQFMatchSGDataSyncStatusField *pSGDataSyncStatus, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///保值额度应答
typedef void (*FuncOnRspQryHedgeVolume)(QFMatchSuperApiInstance instance, struct CQFMatchHedgeVolumeField *pHedgeVolume, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///信用限额查询应答
typedef void (*FuncOnRspQryCreditLimit)(QFMatchSuperApiInstance instance, struct CQFMatchCreditLimitField *pCreditLimit, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///非标组合报单查询应答
typedef void (*FuncOnRspQryCombOrder)(QFMatchSuperApiInstance instance, struct CQFMatchCombOrderField *pCombOrder, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///普通行情查询应答
typedef void (*FuncOnRspQryMarketData)(QFMatchSuperApiInstance instance, struct CQFMatchMarketDataField *pMarketData, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///交易所公告查询请求应答
typedef void (*FuncOnRspQryBulletin)(QFMatchSuperApiInstance instance, struct CQFMatchBulletinField *pBulletin, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///主题查询应答
typedef void (*FuncOnRspQryTopic)(QFMatchSuperApiInstance instance, struct CQFMatchDisseminationField *pDissemination, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///用户在线查询应答
typedef void (*FuncOnRspQryUserSession)(QFMatchSuperApiInstance instance, struct CQFMatchUserSessionField *pUserSession, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///用户查询应答
typedef void (*FuncOnRspQryUser)(QFMatchSuperApiInstance instance, struct CQFMatchUserField *pUser, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///会员查询应答
typedef void (*FuncOnRspQryParticipant)(QFMatchSuperApiInstance instance, struct CQFMatchParticipantField *pParticipant, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///合约价位查询应答
typedef void (*FuncOnRspQryMBLMarketData)(QFMatchSuperApiInstance instance, struct CQFMatchMBLMarketDataField *pMBLMarketData, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///信息查询应答
typedef void (*FuncOnRspQryInformation)(QFMatchSuperApiInstance instance, struct CQFMatchInformationField *pInformation, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///客户持仓明细查询应答
typedef void (*FuncOnRspQryClientPositionDetail)(QFMatchSuperApiInstance instance, struct CQFMatchClientPositionDetailField *pClientPositionDetail, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///成交回报
typedef void (*FuncOnRtnTrade)(QFMatchSuperApiInstance instance, struct CQFMatchTradeField *pTrade);
///报单回报
typedef void (*FuncOnRtnOrder)(QFMatchSuperApiInstance instance, struct CQFMatchOrderField *pOrder);
///执行宣告回报
typedef void (*FuncOnRtnExecOrder)(QFMatchSuperApiInstance instance, struct CQFMatchExecOrderField *pExecOrder);
///报价回报
typedef void (*FuncOnRtnQuote)(QFMatchSuperApiInstance instance, struct CQFMatchQuoteField *pQuote);
///非标组合报单回报
typedef void (*FuncOnRtnCombOrder)(QFMatchSuperApiInstance instance, struct CQFMatchCombOrderField *pCombOrder);
///合约交易状态通知
typedef void (*FuncOnRtnInstrumentStatus)(QFMatchSuperApiInstance instance, struct CQFMatchInstrumentStatusField *pInstrumentStatus);
///增加合约通知
typedef void (*FuncOnRtnInsInstrument)(QFMatchSuperApiInstance instance, struct CQFMatchInstrumentField *pInstrument);
///删除合约通知
typedef void (*FuncOnRtnDelInstrument)(QFMatchSuperApiInstance instance, struct CQFMatchInstrumentField *pInstrument);
///增加合约单腿通知
typedef void (*FuncOnRtnInsCombinationLeg)(QFMatchSuperApiInstance instance, struct CQFMatchCombinationLegField *pCombinationLeg);
///删除合约单腿通知
typedef void (*FuncOnRtnDelCombinationLeg)(QFMatchSuperApiInstance instance, struct CQFMatchCombinationLegField *pCombinationLeg);
///客户持仓通知
typedef void (*FuncOnRtnClientPosition)(QFMatchSuperApiInstance instance, struct CQFMatchClientPositionField *pClientPosition);
///会员资金通知
typedef void (*FuncOnRtnTradingAccount)(QFMatchSuperApiInstance instance, struct CQFMatchTradingAccountField *pTradingAccount);
///别名定义通知
typedef void (*FuncOnRtnAliasDefine)(QFMatchSuperApiInstance instance, struct CQFMatchAliasDefineField *pAliasDefine);
///数据流回退通知
typedef void (*FuncOnRtnFlowMessageCancel)(QFMatchSuperApiInstance instance, struct CQFMatchFlowMessageCancelField *pFlowMessageCancel);
///公告通知
typedef void (*FuncOnRtnBulletin)(QFMatchSuperApiInstance instance, struct CQFMatchBulletinField *pBulletin);
///普通行情通知
typedef void (*FuncOnRtnMarketData)(QFMatchSuperApiInstance instance, struct CQFMatchMarketDataField *pMarketData);
///深度行情通知
typedef void (*FuncOnRtnDepthMarketData)(QFMatchSuperApiInstance instance, struct CQFMatchDepthMarketDataField *pDepthMarketData);
///管理报单通知
typedef void (*FuncOnRtnAdminOrder)(QFMatchSuperApiInstance instance, struct CQFMatchAdminOrderField *pAdminOrder);
///报单录入错误回报
typedef void (*FuncOnErrRtnOrderInsert)(QFMatchSuperApiInstance instance, struct CQFMatchInputOrderField *pInputOrder, struct CQFMatchRspInfoField *pRspInfo);
///报单操作错误回报
typedef void (*FuncOnErrRtnOrderAction)(QFMatchSuperApiInstance instance, struct CQFMatchOrderActionField *pOrderAction, struct CQFMatchRspInfoField *pRspInfo);
///报价录入错误回报
typedef void (*FuncOnErrRtnQuoteInsert)(QFMatchSuperApiInstance instance, struct CQFMatchInputQuoteField *pInputQuote, struct CQFMatchRspInfoField *pRspInfo);
///报价操作错误回报
typedef void (*FuncOnErrRtnQuoteAction)(QFMatchSuperApiInstance instance, struct CQFMatchQuoteActionField *pQuoteAction, struct CQFMatchRspInfoField *pRspInfo);
///执行宣告录入错误回报
typedef void (*FuncOnErrRtnExecOrderInsert)(QFMatchSuperApiInstance instance, struct CQFMatchInputExecOrderField *pInputExecOrder, struct CQFMatchRspInfoField *pRspInfo);
///执行宣告操作错误回报
typedef void (*FuncOnErrRtnExecOrderAction)(QFMatchSuperApiInstance instance, struct CQFMatchExecOrderActionField *pExecOrderAction, struct CQFMatchRspInfoField *pRspInfo);
///非标组合报单录入错误回报
typedef void (*FuncOnErrRtnCombOrderInsert)(QFMatchSuperApiInstance instance, struct CQFMatchInputCombOrderField *pInputCombOrder, struct CQFMatchRspInfoField *pRspInfo);
///导出内存表请求应答
typedef void (*FuncOnRspDumpMemTable)(QFMatchSuperApiInstance instance, struct CQFMatchDumpMemTableField *pDumpMemTable, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///场外成交对申报应答
typedef void (*FuncOnRspOTCTradeInsert)(QFMatchSuperApiInstance instance, struct CQFMatchInputOTCTradeField *pInputOTCTrade, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///场外成交回报
typedef void (*FuncOnRtnOTCTrade)(QFMatchSuperApiInstance instance, struct CQFMatchOTCTradeField *pOTCTrade);
///场外成交错误回报
typedef void (*FuncOnErrRtnOTCTrade)(QFMatchSuperApiInstance instance, struct CQFMatchInputOTCTradeField *pInputOTCTrade, struct CQFMatchRspInfoField *pRspInfo);
///用户出入金应答
typedef void (*FuncOnRspAccountDeposit)(QFMatchSuperApiInstance instance, struct CQFMatchAccountDepositField *pAccountDeposit, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///出入金查询应答
typedef void (*FuncOnRspQryAccountDepositDetail)(QFMatchSuperApiInstance instance, struct CQFMatchAccountDepositDetailField *pAccountDepositDetail, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///实时上场查询应答
typedef void (*FuncOnRspQryTBCommandDetail)(QFMatchSuperApiInstance instance, struct CQFMatchTBCommandDetailField *pTBCommandDetail, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///客户手续费查询应答
typedef void (*FuncOnRspQryPartClientFee)(QFMatchSuperApiInstance instance, struct CQFMatchPartClientFeeField *pPartClientFee, struct CQFMatchRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
///多播行情心跳
typedef void (*FuncOnRtnMultiHeartbeat)(QFMatchSuperApiInstance instance, struct CQFMatchMultiHeartbeatField *pMultiHeartbeat);

/////////////////////////////////////////////////
///
/// 回调函数虚表
///
////////////////////////////////////////////////

typedef struct
{
	FuncOnFrontConnected ptrOnFrontConnected;
	FuncOnFrontDisconnected ptrOnFrontDisconnected;
	FuncOnHeartBeatWarning ptrOnHeartBeatWarning;
	FuncOnPackageStart ptrOnPackageStart;
	FuncOnPackageEnd ptrOnPackageEnd;
	FuncOnRspError ptrOnRspError;
	FuncOnRspUserLogin ptrOnRspUserLogin;
	FuncOnRspUserLogout ptrOnRspUserLogout;
	FuncOnRspOrderInsert ptrOnRspOrderInsert;
	FuncOnRspOrderAction ptrOnRspOrderAction;
	FuncOnRspQuoteInsert ptrOnRspQuoteInsert;
	FuncOnRspQuoteAction ptrOnRspQuoteAction;
	FuncOnRspOTCOrderInsert ptrOnRspOTCOrderInsert;
	FuncOnRspUserPasswordUpdate ptrOnRspUserPasswordUpdate;
	FuncOnRspExecOrderInsert ptrOnRspExecOrderInsert;
	FuncOnRspExecOrderAction ptrOnRspExecOrderAction;
	FuncOnRspAdminOrderInsert ptrOnRspAdminOrderInsert;
	FuncOnRspCombOrderInsert ptrOnRspCombOrderInsert;
	FuncOnRspSubscribeTopic ptrOnRspSubscribeTopic;
	FuncOnRspBulletin ptrOnRspBulletin;
	FuncOnRspInformation ptrOnRspInformation;
	FuncOnRspInstrumentStatusUpdate ptrOnRspInstrumentStatusUpdate;
	FuncOnRspForceUserExit ptrOnRspForceUserExit;
	FuncOnRspForceUserLogout ptrOnRspForceUserLogout;
	FuncOnRspExchangeDataSyncStatusUpdate ptrOnRspExchangeDataSyncStatusUpdate;
	FuncOnRspSGDataSyncStatusUpdate ptrOnRspSGDataSyncStatusUpdate;
	FuncOnRspTBCommand ptrOnRspTBCommand;
	FuncOnRspUpdateMarketData ptrOnRspUpdateMarketData;
	FuncOnRspPositionMarginAction ptrOnRspPositionMarginAction;
	FuncOnRspQryPartAccount ptrOnRspQryPartAccount;
	FuncOnRspQryOrder ptrOnRspQryOrder;
	FuncOnRspQryQuote ptrOnRspQryQuote;
	FuncOnRspQryTrade ptrOnRspQryTrade;
	FuncOnRspQryClient ptrOnRspQryClient;
	FuncOnRspQryPartPosition ptrOnRspQryPartPosition;
	FuncOnRspQryClientPosition ptrOnRspQryClientPosition;
	FuncOnRspQryInstrument ptrOnRspQryInstrument;
	FuncOnRspQryInstrumentStatus ptrOnRspQryInstrumentStatus;
	FuncOnRspQrySGDataSyncStatus ptrOnRspQrySGDataSyncStatus;
	FuncOnRspQryHedgeVolume ptrOnRspQryHedgeVolume;
	FuncOnRspQryCreditLimit ptrOnRspQryCreditLimit;
	FuncOnRspQryCombOrder ptrOnRspQryCombOrder;
	FuncOnRspQryMarketData ptrOnRspQryMarketData;
	FuncOnRspQryBulletin ptrOnRspQryBulletin;
	FuncOnRspQryTopic ptrOnRspQryTopic;
	FuncOnRspQryUserSession ptrOnRspQryUserSession;
	FuncOnRspQryUser ptrOnRspQryUser;
	FuncOnRspQryParticipant ptrOnRspQryParticipant;
	FuncOnRspQryMBLMarketData ptrOnRspQryMBLMarketData;
	FuncOnRspQryInformation ptrOnRspQryInformation;
	FuncOnRspQryClientPositionDetail ptrOnRspQryClientPositionDetail;
	FuncOnRtnTrade ptrOnRtnTrade;
	FuncOnRtnOrder ptrOnRtnOrder;
	FuncOnRtnExecOrder ptrOnRtnExecOrder;
	FuncOnRtnQuote ptrOnRtnQuote;
	FuncOnRtnCombOrder ptrOnRtnCombOrder;
	FuncOnRtnInstrumentStatus ptrOnRtnInstrumentStatus;
	FuncOnRtnInsInstrument ptrOnRtnInsInstrument;
	FuncOnRtnDelInstrument ptrOnRtnDelInstrument;
	FuncOnRtnInsCombinationLeg ptrOnRtnInsCombinationLeg;
	FuncOnRtnDelCombinationLeg ptrOnRtnDelCombinationLeg;
	FuncOnRtnClientPosition ptrOnRtnClientPosition;
	FuncOnRtnTradingAccount ptrOnRtnTradingAccount;
	FuncOnRtnAliasDefine ptrOnRtnAliasDefine;
	FuncOnRtnFlowMessageCancel ptrOnRtnFlowMessageCancel;
	FuncOnRtnBulletin ptrOnRtnBulletin;
	FuncOnRtnMarketData ptrOnRtnMarketData;
	FuncOnRtnDepthMarketData ptrOnRtnDepthMarketData;
	FuncOnRtnAdminOrder ptrOnRtnAdminOrder;
	FuncOnErrRtnOrderInsert ptrOnErrRtnOrderInsert;
	FuncOnErrRtnOrderAction ptrOnErrRtnOrderAction;
	FuncOnErrRtnQuoteInsert ptrOnErrRtnQuoteInsert;
	FuncOnErrRtnQuoteAction ptrOnErrRtnQuoteAction;
	FuncOnErrRtnExecOrderInsert ptrOnErrRtnExecOrderInsert;
	FuncOnErrRtnExecOrderAction ptrOnErrRtnExecOrderAction;
	FuncOnErrRtnCombOrderInsert ptrOnErrRtnCombOrderInsert;
	FuncOnRspDumpMemTable ptrOnRspDumpMemTable;
	FuncOnRspOTCTradeInsert ptrOnRspOTCTradeInsert;
	FuncOnRtnOTCTrade ptrOnRtnOTCTrade;
	FuncOnErrRtnOTCTrade ptrOnErrRtnOTCTrade;
	FuncOnRspAccountDeposit ptrOnRspAccountDeposit;
	FuncOnRspQryAccountDepositDetail ptrOnRspQryAccountDepositDetail;
	FuncOnRspQryTBCommandDetail ptrOnRspQryTBCommandDetail;
	FuncOnRspQryPartClientFee ptrOnRspQryPartClientFee;
	FuncOnRtnMultiHeartbeat ptrOnRtnMultiHeartbeat;
} Callbacks;


///////////////////////////////////////////////
///
/// 主动请求函数的C封装
///
///////////////////////////////////////////////

///创建UserApi
///@param pszFlowPath 存贮订阅信息文件的目录，默认为当前目录
///@param pszUserApiType UserAPI类型
///@return 创建出的UserApi
CQFMATCH_SUPER_API_EXPORT QFMatchSuperApiInstance InitApi(const char *pszFlowPath, const char *pszSuperApiTyp);

///注册回调函数虚表
///@param pInstance InitApi函数返回的API实例地址
///@param callbacks 回调函数虚表的地址
CQFMATCH_SUPER_API_EXPORT void RegisterCallbacks(QFMatchSuperApiInstance instance, Callbacks *callbacks);

///获取系统版本号
///@param pInstance InitApi函数返回的API实例地址
///@return 版本号
CQFMATCH_SUPER_API_EXPORT const char *GetVersion(QFMatchSuperApiInstance instance);

///删除接口对象本身
///@param pInstance InitApi函数返回的API实例地址
///@remark 不再使用本接口对象时,调用该函数删除接口对象
CQFMATCH_SUPER_API_EXPORT void Release(QFMatchSuperApiInstance instance);

///初始化
///@param pInstance InitApi函数返回的API实例地址
///@remark 初始化运行环境,只有调用后,接口才开始工作
CQFMATCH_SUPER_API_EXPORT void Init(QFMatchSuperApiInstance instance);

///等待接口线程结束运行
///@param pInstance InitApi函数返回的API实例地址
///@return 线程退出代码
CQFMATCH_SUPER_API_EXPORT int Join(QFMatchSuperApiInstance instance);

///获取当前交易日
///@param pInstance InitApi函数返回的API实例地址
///@retrun 获取到的交易日
///@remark 只有登录成功后,才能得到正确的交易日
CQFMATCH_SUPER_API_EXPORT const char *GetTradingDay(QFMatchSuperApiInstance instance);

///注册前置机网络地址
///@param pInstance InitApi函数返回的API实例地址
///@param pszFrontAddress：前置机网络地址。
///@remark 网络地址的格式为：“protocol://ipaddress:port”，如：”tcp://127.0.0.1:17001”。 
///@remark “tcp”代表传输协议，“127.0.0.1”代表服务器地址。”17001”代表服务器端口号。
CQFMATCH_SUPER_API_EXPORT void RegisterFront(QFMatchSuperApiInstance instance, char *pszFrontAddress);

///注册名字服务器网络地址
///@param pInstance InitApi函数返回的API实例地址
///@param pszNsAddress：名字服务器网络地址。
///@remark 网络地址的格式为：“protocol://ipaddress:port”，如：”tcp://127.0.0.1:12001”。 
///@remark “tcp”代表传输协议，“127.0.0.1”代表服务器地址。”12001”代表服务器端口号。
///@remark RegisterFront优先于RegisterNameServer
CQFMATCH_SUPER_API_EXPORT void RegisterNameServer(QFMatchSuperApiInstance instance, char *pszNsAddress);

///加载证书
///@param pInstance InitApi函数返回的API实例地址
///@param pszCertFileName 用户证书文件名
///@param pszKeyFileName 用户私钥文件名
///@param pszCaFileName 可信任CA证书文件名
///@param pszKeyFilePassword 用户私钥文件密码
///@return 0 操作成功
///@return -1 可信任CA证书载入失败
///@return -2 用户证书载入失败
///@return -3 用户私钥载入失败	
///@return -4 用户证书校验失败
CQFMATCH_SUPER_API_EXPORT int RegisterCertificateFile(QFMatchSuperApiInstance instance, const char *pszCertFileName, const char *pszKeyFileName, 
	const char *pszCaFileName, const char *pszKeyFilePassword);

///订阅市场行情。
///@param pInstance InitApi函数返回的API实例地址
///@param nTopicID 市场行情主题  
///@param nResumeType 市场行情重传方式  
///        QFMATCH_TERT_RESTART:从本交易日开始重传
///        QFMATCH_TERT_RESUME:从上次收到的续传
///        QFMATCH_TERT_QUICK:先传送当前行情快照,再传送登录后市场行情的内容
///@remark 该方法要在Init方法前调用。若不调用则不会收到私有流的数据。
CQFMATCH_SUPER_API_EXPORT void SubscribeMarketDataTopic(QFMatchSuperApiInstance instance, int nTopicID, enum QFMATCH_TE_RESUME_TYPE nResumeType);

///订阅私有流。
///@param pInstance InitApi函数返回的API实例地址
///@param nResumeType 私有流重传方式  
///        QFMATCH_TERT_RESTART:从本交易日开始重传
///        QFMATCH_TERT_RESUME:从上次收到的续传
///        QFMATCH_TERT_QUICK:只传送登录后私有流的内容
///@remark 该方法要在Init方法前调用。若不调用则不会收到私有流的数据。
CQFMATCH_SUPER_API_EXPORT void SubscribePrivateTopic(QFMatchSuperApiInstance instance, enum QFMATCH_TE_RESUME_TYPE nResumeType);

///订阅公共流。
///@param pInstance InitApi函数返回的API实例地址
///@param nResumeType 公共流重传方式  
///        QFMATCH_TERT_RESTART:从本交易日开始重传
///        QFMATCH_TERT_RESUME:从上次收到的续传
///        QFMATCH_TERT_QUICK:只传送登录后公共流的内容
///@remark 该方法要在Init方法前调用。若不调用则不会收到公共流的数据。
CQFMATCH_SUPER_API_EXPORT void SubscribePublicTopic(QFMatchSuperApiInstance instance, enum QFMATCH_TE_RESUME_TYPE nResumeType);

///订阅交易员流。
///@param pInstance InitApi函数返回的API实例地址
///@param nResumeType 交易员流重传方式
///        QFMATCH_TERT_RESTART:从本交易日开始重传
///        QFMATCH_TERT_RESUME:从上次收到的续传
///        QFMATCH_TERT_QUICK:只传送登录后交易员流的内容
///@remark 该方法要在Init方法前调用。若不调用则不会收到交易员流的数据。
CQFMATCH_SUPER_API_EXPORT void SubscribeUserTopic(QFMatchSuperApiInstance instance, enum QFMATCH_TE_RESUME_TYPE nResumeType);


///设置心跳超时时间。
///@param pInstance InitApi函数返回的API实例地址
///@param timeout 心跳超时时间(秒)
CQFMATCH_SUPER_API_EXPORT void SetHeartbeatTimeout(QFMatchSuperApiInstance instance, unsigned int timeout);

///打开请求日志文件
///@param pInstance InitApi函数返回的API实例地址
///@param pszReqLogFileName 请求日志文件名
///@return 0 操作成功
///@return -1 打开日志文件失败
CQFMATCH_SUPER_API_EXPORT int OpenRequestLog(QFMatchSuperApiInstance instance, const char *pszReqLogFileName);

///打开应答日志文件
///@param pInstance InitApi函数返回的API实例地址
///@param pszRspLogFileName 应答日志文件名  
///@return 0 操作成功
///@return -1 打开日志文件失败
CQFMATCH_SUPER_API_EXPORT int OpenResponseLog(QFMatchSuperApiInstance instance, const char *pszRspLogFileName);

///用户登录请求
///@param pInstance InitApi函数返回的API实例地址
CQFMATCH_SUPER_API_EXPORT int ReqUserLogin(QFMatchSuperApiInstance instance, struct CQFMatchReqUserLoginField *pReqUserLoginField, int nRequestID);

///用户退出请求
CQFMATCH_SUPER_API_EXPORT  int ReqUserLogout(QFMatchSuperApiInstance instance, struct CQFMatchReqUserLogoutField *pReqUserLogout, int nRequestID);

///报单录入请求
CQFMATCH_SUPER_API_EXPORT  int ReqOrderInsert(QFMatchSuperApiInstance instance, struct CQFMatchInputOrderField *pInputOrder, int nRequestID);

///报单操作请求
CQFMATCH_SUPER_API_EXPORT  int ReqOrderAction(QFMatchSuperApiInstance instance, struct CQFMatchOrderActionField *pOrderAction, int nRequestID);

///报价录入请求
CQFMATCH_SUPER_API_EXPORT  int ReqQuoteInsert(QFMatchSuperApiInstance instance, struct CQFMatchInputQuoteField *pInputQuote, int nRequestID);

///报价操作请求
CQFMATCH_SUPER_API_EXPORT  int ReqQuoteAction(QFMatchSuperApiInstance instance, struct CQFMatchQuoteActionField *pQuoteAction, int nRequestID);

///OTC报单录入请求
CQFMATCH_SUPER_API_EXPORT  int ReqOTCOrderInsert(QFMatchSuperApiInstance instance, struct CQFMatchOTCOrderField *pOTCOrder, int nRequestID);

///用户密码修改请求
CQFMATCH_SUPER_API_EXPORT  int ReqUserPasswordUpdate(QFMatchSuperApiInstance instance, struct CQFMatchUserPasswordUpdateField *pUserPasswordUpdate, int nRequestID);

///执行宣告录入请求
CQFMATCH_SUPER_API_EXPORT  int ReqExecOrderInsert(QFMatchSuperApiInstance instance, struct CQFMatchInputExecOrderField *pInputExecOrder, int nRequestID);

///执行宣告操作请求
CQFMATCH_SUPER_API_EXPORT  int ReqExecOrderAction(QFMatchSuperApiInstance instance, struct CQFMatchExecOrderActionField *pExecOrderAction, int nRequestID);

///管理报单录入请求
CQFMATCH_SUPER_API_EXPORT  int ReqAdminOrderInsert(QFMatchSuperApiInstance instance, struct CQFMatchInputAdminOrderField *pInputAdminOrder, int nRequestID);

///非标组合报单录入请求
CQFMATCH_SUPER_API_EXPORT  int ReqCombOrderInsert(QFMatchSuperApiInstance instance, struct CQFMatchInputCombOrderField *pInputCombOrder, int nRequestID);

///订阅主题请求
CQFMATCH_SUPER_API_EXPORT  int ReqSubscribeTopic(QFMatchSuperApiInstance instance, struct CQFMatchDisseminationField *pDissemination, int nRequestID);

///交易所公告请求
CQFMATCH_SUPER_API_EXPORT  int ReqBulletin(QFMatchSuperApiInstance instance, struct CQFMatchBulletinField *pBulletin, int nRequestID);

///交易所信息请求
CQFMATCH_SUPER_API_EXPORT  int ReqInformation(QFMatchSuperApiInstance instance, struct CQFMatchInformationField *pInformation, int nRequestID);

///合约交易状态改变请求
CQFMATCH_SUPER_API_EXPORT  int ReqInstrumentStatusUpdate(QFMatchSuperApiInstance instance, struct CQFMatchInstrumentStatusField *pInstrumentStatus, int nRequestID);

///强制用户退出请求
CQFMATCH_SUPER_API_EXPORT  int ReqForceUserExit(QFMatchSuperApiInstance instance, struct CQFMatchForceUserExitField *pForceUserExit, int nRequestID);

///用户会话删除请求
CQFMATCH_SUPER_API_EXPORT  int ReqForceUserLogout(QFMatchSuperApiInstance instance, struct CQFMatchForceUserExitField *pForceUserExit, int nRequestID);

///交易所数据同步状态改变请求
CQFMATCH_SUPER_API_EXPORT  int ReqExchangeDataSyncStatusUpdate(QFMatchSuperApiInstance instance, struct CQFMatchExchangeDataSyncStatusField *pExchangeDataSyncStatus, int nRequestID);

///结算组数据同步状态改变请求
CQFMATCH_SUPER_API_EXPORT  int ReqSGDataSyncStatusUpdate(QFMatchSuperApiInstance instance, struct CQFMatchSGDataSyncStatusField *pSGDataSyncStatus, int nRequestID);

///实时上场请求
CQFMATCH_SUPER_API_EXPORT  int ReqTBCommand(QFMatchSuperApiInstance instance, struct CQFMatchTBCommandField *pTBCommand, int nRequestID);

///更新行情请求
CQFMATCH_SUPER_API_EXPORT  int ReqUpdateMarketData(QFMatchSuperApiInstance instance, struct CQFMatchUpdateMarketDataField *pUpdateMarketData, int nRequestID);

///仓位保证金操作请求
CQFMATCH_SUPER_API_EXPORT  int ReqPositionMarginAction(QFMatchSuperApiInstance instance, struct CQFMatchPositionMarginActionField *pPositionMarginAction, int nRequestID);

///会员资金查询请求
CQFMATCH_SUPER_API_EXPORT  int ReqQryPartAccount(QFMatchSuperApiInstance instance, struct CQFMatchQryPartAccountField *pQryPartAccount, int nRequestID);

///报单查询请求
CQFMATCH_SUPER_API_EXPORT  int ReqQryOrder(QFMatchSuperApiInstance instance, struct CQFMatchQryOrderField *pQryOrder, int nRequestID);

///报价查询请求
CQFMATCH_SUPER_API_EXPORT  int ReqQryQuote(QFMatchSuperApiInstance instance, struct CQFMatchQryQuoteField *pQryQuote, int nRequestID);

///成交单查询请求
CQFMATCH_SUPER_API_EXPORT  int ReqQryTrade(QFMatchSuperApiInstance instance, struct CQFMatchQryTradeField *pQryTrade, int nRequestID);

///会员客户查询请求
CQFMATCH_SUPER_API_EXPORT  int ReqQryClient(QFMatchSuperApiInstance instance, struct CQFMatchQryClientField *pQryClient, int nRequestID);

///会员持仓查询请求
CQFMATCH_SUPER_API_EXPORT  int ReqQryPartPosition(QFMatchSuperApiInstance instance, struct CQFMatchQryPartPositionField *pQryPartPosition, int nRequestID);

///客户持仓查询请求
CQFMATCH_SUPER_API_EXPORT  int ReqQryClientPosition(QFMatchSuperApiInstance instance, struct CQFMatchQryClientPositionField *pQryClientPosition, int nRequestID);

///合约查询请求
CQFMATCH_SUPER_API_EXPORT  int ReqQryInstrument(QFMatchSuperApiInstance instance, struct CQFMatchQryInstrumentField *pQryInstrument, int nRequestID);

///合约交易状态查询请求
CQFMATCH_SUPER_API_EXPORT  int ReqQryInstrumentStatus(QFMatchSuperApiInstance instance, struct CQFMatchQryInstrumentStatusField *pQryInstrumentStatus, int nRequestID);

///结算组状态查询请求
CQFMATCH_SUPER_API_EXPORT  int ReqQrySGDataSyncStatus(QFMatchSuperApiInstance instance, struct CQFMatchQrySGDataSyncStatusField *pQrySGDataSyncStatus, int nRequestID);

///保值额度查询
CQFMATCH_SUPER_API_EXPORT  int ReqQryHedgeVolume(QFMatchSuperApiInstance instance, struct CQFMatchQryHedgeVolumeField *pQryHedgeVolume, int nRequestID);

///信用限额查询请求
CQFMATCH_SUPER_API_EXPORT  int ReqQryCreditLimit(QFMatchSuperApiInstance instance, struct CQFMatchQryCreditLimitField *pQryCreditLimit, int nRequestID);

///非标组合报单查询请求
CQFMATCH_SUPER_API_EXPORT  int ReqQryCombOrder(QFMatchSuperApiInstance instance, struct CQFMatchQryCombOrderField *pQryCombOrder, int nRequestID);

///普通行情查询请求
CQFMATCH_SUPER_API_EXPORT  int ReqQryMarketData(QFMatchSuperApiInstance instance, struct CQFMatchQryMarketDataField *pQryMarketData, int nRequestID);

///交易所公告查询请求
CQFMATCH_SUPER_API_EXPORT  int ReqQryBulletin(QFMatchSuperApiInstance instance, struct CQFMatchQryBulletinField *pQryBulletin, int nRequestID);

///主题查询请求
CQFMATCH_SUPER_API_EXPORT  int ReqQryTopic(QFMatchSuperApiInstance instance, struct CQFMatchDisseminationField *pDissemination, int nRequestID);

///用户在线查询请求
CQFMATCH_SUPER_API_EXPORT  int ReqQryUserSession(QFMatchSuperApiInstance instance, struct CQFMatchQryUserSessionField *pQryUserSession, int nRequestID);

///用户查询请求
CQFMATCH_SUPER_API_EXPORT  int ReqQryUser(QFMatchSuperApiInstance instance, struct CQFMatchQryUserField *pQryUser, int nRequestID);

///会员查询请求
CQFMATCH_SUPER_API_EXPORT  int ReqQryParticipant(QFMatchSuperApiInstance instance, struct CQFMatchQryParticipantField *pQryParticipant, int nRequestID);

///合约价位查询
CQFMATCH_SUPER_API_EXPORT  int ReqQryMBLMarketData(QFMatchSuperApiInstance instance, struct CQFMatchQryMBLMarketDataField *pQryMBLMarketData, int nRequestID);

///信息查询
CQFMATCH_SUPER_API_EXPORT  int ReqQryInformation(QFMatchSuperApiInstance instance, struct CQFMatchQryInformationField *pQryInformation, int nRequestID);

///客户持仓明细查询
CQFMATCH_SUPER_API_EXPORT  int ReqQryClientPositionDetail(QFMatchSuperApiInstance instance, struct CQFMatchQryClientPositionDetailField *pQryClientPositionDetail, int nRequestID);

///客户持仓查询请求
CQFMATCH_SUPER_API_EXPORT  int ReqQryClientPositionV1(QFMatchSuperApiInstance instance, struct CQFMatchQryClientPositionV1Field *pQryClientPositionV1, int nRequestID);

///导出内存表请求
CQFMATCH_SUPER_API_EXPORT  int ReqDumpMemTable(QFMatchSuperApiInstance instance, struct CQFMatchDumpMemTableField *pDumpMemTable, int nRequestID);

///场外成交对申报
CQFMATCH_SUPER_API_EXPORT  int ReqOTCTradeInsert(QFMatchSuperApiInstance instance, struct CQFMatchInputOTCTradeField *pInputOTCTrade, int nRequestID);

///用户出入金请求
CQFMATCH_SUPER_API_EXPORT  int ReqAccountDeposit(QFMatchSuperApiInstance instance, struct CQFMatchAccountDepositField *pAccountDeposit, int nRequestID);

///出入金查询请求
CQFMATCH_SUPER_API_EXPORT  int ReqQryAccountDepositDetail(QFMatchSuperApiInstance instance, struct CQFMatchQryAccountDepositDetailField *pQryAccountDepositDetail, int nRequestID);

///实时上场查询请求
CQFMATCH_SUPER_API_EXPORT  int ReqQryTBCommandDetail(QFMatchSuperApiInstance instance, struct CQFMatchQryTBCommandDetailField *pQryTBCommandDetail, int nRequestID);

///客户手续费查询
CQFMATCH_SUPER_API_EXPORT  int ReqQryPartClientFee(QFMatchSuperApiInstance instance, struct CQFMatchQryPartClientFeeField *pQryPartClientFee, int nRequestID);

#endif

#ifdef __cplusplus
}
#endif
