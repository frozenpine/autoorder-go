/////////////////////////////////////////////////////////////////////////
///@file QFMatchApiStruct.h
///@brief 定义了客户端接口使用的业务数据结构
/////////////////////////////////////////////////////////////////////////

#if !defined(QFMatch_STRUCT_H)
#define QFMatch_STRUCT_H

#if _MSC_VER > 1000
#pragma once
#endif // _MSC_VER > 1000

#include "QFMatchApiDataType.h"

///信息分发
struct CQFMatchDisseminationField
{
	///序列系列号
	TQFMatchSequenceSeriesType	SequenceSeries;
	///序列号
	TQFMatchSequenceNoType	SequenceNo;
};
///响应信息
struct CQFMatchRspInfoField
{
	///错误代码
	TQFMatchErrorIDType	ErrorID;
	///错误信息
	TQFMatchErrorMsgType	ErrorMsg;
};
///通讯阶段
struct CQFMatchCommPhaseField
{
	///交易日
	TQFMatchDateType	TradingDay;
	///通讯时段号
	TQFMatchCommPhaseNoType	CommPhaseNo;
};
///交易所交易日
struct CQFMatchExchangeTradingDayField
{
	///交易日
	TQFMatchDateType	TradingDay;
	///交易所代码
	TQFMatchExchangeIDType	ExchangeID;
};
///结算会话
struct CQFMatchSettlementSessionField
{
	///交易日
	TQFMatchDateType	TradingDay;
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///结算编号
	TQFMatchSettlementIDType	SettlementID;
};
///当前时间
struct CQFMatchCurrentTimeField
{
	///当前日期
	TQFMatchDateType	CurrDate;
	///当前时间
	TQFMatchTimeType	CurrTime;
	///当前时间（毫秒）
	TQFMatchMillisecType	CurrMillisec;
};
///用户登录请求
struct CQFMatchReqUserLoginField
{
	///交易日
	TQFMatchDateType	TradingDay;
	///交易用户代码
	TQFMatchUserIDType	UserID;
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///密码
	TQFMatchPasswordType	Password;
	///用户端产品信息
	TQFMatchProductInfoType	UserProductInfo;
	///接口端产品信息
	TQFMatchProductInfoType	InterfaceProductInfo;
	///协议信息
	TQFMatchProtocolInfoType	ProtocolInfo;
	///数据中心代码
	TQFMatchDataCenterIDType	DataCenterID;
};
///用户登录应答
struct CQFMatchRspUserLoginField
{
	///交易日
	TQFMatchDateType	TradingDay;
	///登录成功时间
	TQFMatchTimeType	LoginTime;
	///最大本地报单号
	TQFMatchOrderLocalIDType	MaxOrderLocalID;
	///交易用户代码
	TQFMatchUserIDType	UserID;
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///交易系统名称
	TQFMatchTradingSystemNameType	TradingSystemName;
	///数据中心代码
	TQFMatchDataCenterIDType	DataCenterID;
	///会员私有流当前长度
	TQFMatchSequenceNoType	PrivateFlowSize;
	///交易员私有流当前长度
	TQFMatchSequenceNoType	UserFlowSize;
};
///用户登出请求
struct CQFMatchReqUserLogoutField
{
	///交易用户代码
	TQFMatchUserIDType	UserID;
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
};
///用户登出应答
struct CQFMatchRspUserLogoutField
{
	///交易用户代码
	TQFMatchUserIDType	UserID;
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
};
///输入报单
struct CQFMatchInputOrderField
{
	///报单编号
	TQFMatchOrderSysIDType	OrderSysID;
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///客户代码
	TQFMatchClientIDType	ClientID;
	///交易用户代码
	TQFMatchUserIDType	UserID;
	///合约代码
	TQFMatchInstrumentIDType	InstrumentID;
	///报单价格条件
	TQFMatchOrderPriceTypeType	OrderPriceType;
	///买卖方向
	TQFMatchDirectionType	Direction;
	///组合开平标志
	TQFMatchCombOffsetFlagType	CombOffsetFlag;
	///组合投机套保标志
	TQFMatchCombHedgeFlagType	CombHedgeFlag;
	///价格
	TQFMatchPriceType	LimitPrice;
	///数量
	TQFMatchVolumeType	VolumeTotalOriginal;
	///有效期类型
	TQFMatchTimeConditionType	TimeCondition;
	///GTD日期
	TQFMatchDateType	GTDDate;
	///成交量类型
	TQFMatchVolumeConditionType	VolumeCondition;
	///最小成交量
	TQFMatchVolumeType	MinVolume;
	///触发条件
	TQFMatchContingentConditionType	ContingentCondition;
	///止损价
	TQFMatchPriceType	StopPrice;
	///强平原因
	TQFMatchForceCloseReasonType	ForceCloseReason;
	///本地报单编号
	TQFMatchOrderLocalIDType	OrderLocalID;
	///自动挂起标志
	TQFMatchBoolType	IsAutoSuspend;
	///业务单元
	TQFMatchBusinessUnitType	BusinessUnit;
};
///报单操作
struct CQFMatchOrderActionField
{
	///报单编号
	TQFMatchOrderSysIDType	OrderSysID;
	///本地报单编号
	TQFMatchOrderLocalIDType	OrderLocalID;
	///报单操作标志
	TQFMatchActionFlagType	ActionFlag;
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///客户代码
	TQFMatchClientIDType	ClientID;
	///交易用户代码
	TQFMatchUserIDType	UserID;
	///价格
	TQFMatchPriceType	LimitPrice;
	///数量变化
	TQFMatchVolumeType	VolumeChange;
	///操作本地编号
	TQFMatchOrderLocalIDType	ActionLocalID;
	///业务单元
	TQFMatchBusinessUnitType	BusinessUnit;
};
///OTC报单
struct CQFMatchOTCOrderField
{
	///交易日
	TQFMatchDateType	TradingDay;
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///结算编号
	TQFMatchSettlementIDType	SettlementID;
	///OTC报单编号
	TQFMatchOTCOrderSysIDType	OTCOrderSysID;
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///客户代码
	TQFMatchClientIDType	ClientID;
	///交易用户代码
	TQFMatchUserIDType	UserID;
	///合约代码
	TQFMatchInstrumentIDType	InstrumentID;
	///买卖方向
	TQFMatchDirectionType	Direction;
	///开平标志
	TQFMatchOffsetFlagType	OffsetFlag;
	///投机套保标志
	TQFMatchHedgeFlagType	HedgeFlag;
	///价格
	TQFMatchPriceType	Price;
	///数量
	TQFMatchVolumeType	Volume;
	///对手方会员代码
	TQFMatchParticipantIDType	OtherParticipantID;
	///对手方客户代码
	TQFMatchClientIDType	OtherClientID;
	///对手方交易用户代码
	TQFMatchUserIDType	OtherUserID;
	///对手方开平标志
	TQFMatchOffsetFlagType	OtherOffsetFlag;
	///对手方套保标志
	TQFMatchHedgeFlagType	OtherHedgeFlag;
	///本地OTC报单编号
	TQFMatchOrderLocalIDType	OTCOrderLocalID;
	///OTC报单状态
	TQFMatchOTCOrderStatusType	OTCOrderStatus;
	///插入时间
	TQFMatchTimeType	InsertTime;
	///撤销时间
	TQFMatchTimeType	CancelTime;
	///结算会员编号
	TQFMatchParticipantIDType	ClearingPartID;
	///对手方结算会员编号
	TQFMatchParticipantIDType	OtherClearingPartID;
	///自然日期
	TQFMatchDateType	CalendarDate;
};
///输入报价
struct CQFMatchInputQuoteField
{
	///报价编号
	TQFMatchQuoteSysIDType	QuoteSysID;
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///客户代码
	TQFMatchClientIDType	ClientID;
	///交易用户代码
	TQFMatchUserIDType	UserID;
	///数量
	TQFMatchVolumeType	Volume;
	///合约代码
	TQFMatchInstrumentIDType	InstrumentID;
	///本地报价编号
	TQFMatchOrderLocalIDType	QuoteLocalID;
	///业务单元
	TQFMatchBusinessUnitType	BusinessUnit;
	///买方组合开平标志
	TQFMatchCombOffsetFlagType	BidCombOffsetFlag;
	///买方组合套保标志
	TQFMatchCombHedgeFlagType	BidCombHedgeFlag;
	///买方价格
	TQFMatchPriceType	BidPrice;
	///卖方组合开平标志
	TQFMatchCombOffsetFlagType	AskCombOffsetFlag;
	///卖方组合套保标志
	TQFMatchCombHedgeFlagType	AskCombHedgeFlag;
	///卖方价格
	TQFMatchPriceType	AskPrice;
};
///报价操作
struct CQFMatchQuoteActionField
{
	///报价编号
	TQFMatchQuoteSysIDType	QuoteSysID;
	///本地报价编号
	TQFMatchOrderLocalIDType	QuoteLocalID;
	///报单操作标志
	TQFMatchActionFlagType	ActionFlag;
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///客户代码
	TQFMatchClientIDType	ClientID;
	///交易用户代码
	TQFMatchUserIDType	UserID;
	///操作本地编号
	TQFMatchOrderLocalIDType	ActionLocalID;
	///业务单元
	TQFMatchBusinessUnitType	BusinessUnit;
};
///输入执行宣告
struct CQFMatchInputExecOrderField
{
	///合约编号
	TQFMatchInstrumentIDType	InstrumentID;
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///客户代码
	TQFMatchClientIDType	ClientID;
	///交易用户代码
	TQFMatchUserIDType	UserID;
	///本地执行宣告编号
	TQFMatchOrderLocalIDType	ExecOrderLocalID;
	///数量
	TQFMatchVolumeType	Volume;
	///业务单元
	TQFMatchBusinessUnitType	BusinessUnit;
};
///执行宣告操作
struct CQFMatchExecOrderActionField
{
	///执行宣告编号
	TQFMatchExecOrderSysIDType	ExecOrderSysID;
	///本地执行宣告编号
	TQFMatchOrderLocalIDType	ExecOrderLocalID;
	///报单操作标志
	TQFMatchActionFlagType	ActionFlag;
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///客户代码
	TQFMatchClientIDType	ClientID;
	///交易用户代码
	TQFMatchUserIDType	UserID;
	///操作本地编号
	TQFMatchOrderLocalIDType	ActionLocalID;
	///业务单元
	TQFMatchBusinessUnitType	BusinessUnit;
};
///用户登录退出
struct CQFMatchUserLogoutField
{
	///交易用户代码
	TQFMatchUserIDType	UserID;
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
};
///用户口令修改
struct CQFMatchUserPasswordUpdateField
{
	///交易用户代码
	TQFMatchUserIDType	UserID;
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///旧密码
	TQFMatchPasswordType	OldPassword;
	///新密码
	TQFMatchPasswordType	NewPassword;
};
///输入非标组合报单
struct CQFMatchInputCombOrderField
{
	///组合报单编号
	TQFMatchOrderSysIDType	CombOrderSysID;
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///客户代码
	TQFMatchClientIDType	ClientID;
	///交易用户代码
	TQFMatchUserIDType	UserID;
	///价格
	TQFMatchPriceType	LimitPrice;
	///数量
	TQFMatchVolumeType	VolumeTotalOriginal;
	///本地报单编号
	TQFMatchOrderLocalIDType	CombOrderLocalID;
	///业务单元
	TQFMatchBusinessUnitType	BusinessUnit;
	///合约代码1
	TQFMatchInstrumentIDType	InstrumentID1;
	///买卖方向1
	TQFMatchDirectionType	Direction1;
	///分腿乘数1
	TQFMatchLegMultipleType	LegMultiple1;
	///开平标志1
	TQFMatchOffsetFlagType	OffsetFlag1;
	///投机套保标志1
	TQFMatchHedgeFlagType	HedgeFlag1;
	///合约代码2
	TQFMatchInstrumentIDType	InstrumentID2;
	///买卖方向2
	TQFMatchDirectionType	Direction2;
	///分腿乘数2
	TQFMatchLegMultipleType	LegMultiple2;
	///开平标志2
	TQFMatchOffsetFlagType	OffsetFlag2;
	///投机套保标志2
	TQFMatchHedgeFlagType	HedgeFlag2;
	///合约代码3
	TQFMatchInstrumentIDType	InstrumentID3;
	///买卖方向3
	TQFMatchDirectionType	Direction3;
	///分腿乘数3
	TQFMatchLegMultipleType	LegMultiple3;
	///开平标志3
	TQFMatchOffsetFlagType	OffsetFlag3;
	///投机套保标志3
	TQFMatchHedgeFlagType	HedgeFlag3;
	///合约代码4
	TQFMatchInstrumentIDType	InstrumentID4;
	///买卖方向4
	TQFMatchDirectionType	Direction4;
	///分腿乘数4
	TQFMatchLegMultipleType	LegMultiple4;
	///开平标志4
	TQFMatchOffsetFlagType	OffsetFlag4;
	///投机套保标志4
	TQFMatchHedgeFlagType	HedgeFlag4;
};
///强制用户退出
struct CQFMatchForceUserExitField
{
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///交易用户代码
	TQFMatchUserIDType	UserID;
};
///导出内存表
struct CQFMatchDumpMemTableField
{
	///内存表名称
	TQFMatchTableNameType	TableName;
};
///资金账户出入金
struct CQFMatchAccountDepositField
{
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///出入金本地号
	TQFMatchOrderLocalIDType	DepositLocalID;
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///交易角色
	TQFMatchTradingRoleType	TradingRole;
	///资金账号
	TQFMatchAccountIDType	AccountID;
	///出入金金额
	TQFMatchMoneyType	Deposit;
	///全局出入金流水序号
	TQFMatchSequenceNoType	DepositNo;
	///币种
	TQFMatchCurrencyType	Currency;
	///资金操作类型
	TQFMatchAccountActionTypeType	ActionType;
	///备注
	TQFMatchComeFromType	Comments;
};
///出入金
struct CQFMatchAccountDepositDetailField
{
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///出入金本地号
	TQFMatchOrderLocalIDType	DepositLocalID;
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///交易角色
	TQFMatchTradingRoleType	TradingRole;
	///资金账号
	TQFMatchAccountIDType	AccountID;
	///出入金金额
	TQFMatchMoneyType	Deposit;
	///全局出入金流水序号
	TQFMatchSequenceNoType	DepositNo;
	///币种
	TQFMatchCurrencyType	Currency;
	///资金操作类型
	TQFMatchAccountActionTypeType	ActionType;
	///备注
	TQFMatchComeFromType	Comments;
	///处理结果
	TQFMatchTransStatusType	Status;
	///插入日期
	TQFMatchDateType	CalendarDate;
	///插入时间
	TQFMatchTimeType	InsertTime;
};
///实时上场
struct CQFMatchTBCommandDetailField
{
	///交易日
	TQFMatchDateType	TradingDay;
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///结算编号
	TQFMatchSettlementIDType	SettlementID;
	///DB命令本地号
	TQFMatchOrderLocalIDType	CommandLocalID;
	///全局DB命令序号
	TQFMatchSequenceNoType	CommandNo;
	///DB命令类型
	TQFMatchCommandTypeType	CommandType;
	///字段名
	TQFMatchFieldNameType	FieldName;
	///字段内容
	TQFMatchFieldContentType	FieldContent;
	///备注
	TQFMatchComeFromType	Comments;
	///处理结果
	TQFMatchTransStatusType	Status;
	///插入日期
	TQFMatchDateType	CalendarDate;
	///插入时间
	TQFMatchTimeType	InsertTime;
};
///手续费
struct CQFMatchPartClientFeeField
{
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///客户号
	TQFMatchClientIDType	ClientID;
	///合约的编号
	TQFMatchInstrumentIDType	InstrumentID;
	///成交角色
	TQFMatchMatchRoleType	MatchRole;
	///开仓手续费按比例
	TQFMatchRatioType	OpenFeeRate;
	///开仓手续费按手数
	TQFMatchRatioType	OpenFeeAmt;
	///平仓手续费按比例
	TQFMatchRatioType	OffsetFeeRate;
	///平仓手续费按手数
	TQFMatchRatioType	OffsetFeeAmt;
	///平今仓手续费按比例
	TQFMatchRatioType	OTFeeRate;
	///平今仓手续费按手数
	TQFMatchRatioType	OTFeeAmt;
};
///客户持仓明细
struct CQFMatchClientPositionDetailField
{
	///交易日
	TQFMatchDateType	TradingDay;
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///结算编号
	TQFMatchSettlementIDType	SettlementID;
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///客户代码
	TQFMatchClientIDType	ClientID;
	///成交日期
	TQFMatchDateType	TradeDate;
	///合约代码
	TQFMatchInstrumentIDType	InstrumentID;
	///买卖方向
	TQFMatchDirectionType	Direction;
	///成交编号
	TQFMatchTradeIDType	TradeID;
	///成交类型
	TQFMatchTradeTypeType	TradeType;
	///当日可平仓量
	TQFMatchVolumeType	Position;
	///成交价格
	TQFMatchPriceType	Price;
	///持仓成本包括手续费等
	TQFMatchMoneyType	PositionCost;
	///持仓成本
	TQFMatchMoneyType	PositionMoney;
	///是否今日仓
	TQFMatchBoolType	IsToday;
	///占用的保证金
	TQFMatchMoneyType	UseMargin;
	///手续费
	TQFMatchMoneyType	Fee;
	///投机套保标志
	TQFMatchHedgeFlagType	HedgeFlag;
};
///会员
struct CQFMatchParticipantField
{
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///会员名称
	TQFMatchParticipantNameType	ParticipantName;
	///会员简称
	TQFMatchParticipantAbbrType	ParticipantAbbr;
	///会员类型
	TQFMatchMemberTypeType	MemberType;
	///是否活跃
	TQFMatchBoolType	IsActive;
};
///用户
struct CQFMatchUserField
{
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///交易用户代码
	TQFMatchUserIDType	UserID;
	///交易用户类型
	TQFMatchUserTypeType	UserType;
	///密码
	TQFMatchPasswordType	Password;
	///交易员权限
	TQFMatchUserActiveType	IsActive;
};
///客户
struct CQFMatchClientField
{
	///客户代码
	TQFMatchClientIDType	ClientID;
	///客户名称
	TQFMatchPartyNameType	ClientName;
	///证件类型
	TQFMatchIdCardTypeType	IdentifiedCardType;
	///证件号码
	TQFMatchIdentifiedCardNoType	IdentifiedCardNo;
	///交易角色
	TQFMatchTradingRoleType	TradingRole;
	///客户类型
	TQFMatchClientTypeType	ClientType;
	///是否活跃
	TQFMatchBoolType	IsActive;
	///客户交易类型
	TQFMatchHedgeFlagType	HedgeFlag;
};
///用户会话
struct CQFMatchUserSessionField
{
	///前置编号
	TQFMatchFrontIDType	FrontID;
	///交易用户代码
	TQFMatchUserIDType	UserID;
	///交易用户类型
	TQFMatchUserTypeType	UserType;
	///会话编号
	TQFMatchSessionIDType	SessionID;
	///登录时间
	TQFMatchTimeType	LoginTime;
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///IP地址
	TQFMatchIPAddressType	IPAddress;
	///用户端产品信息
	TQFMatchProductInfoType	UserProductInfo;
	///接口端产品信息
	TQFMatchProductInfoType	InterfaceProductInfo;
	///协议信息
	TQFMatchProtocolInfoType	ProtocolInfo;
};
///产品组
struct CQFMatchProductGroupField
{
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///产品组代码
	TQFMatchProductGroupIDType	ProductGroupID;
	///产品组名称
	TQFMatchProductGroupNameType	ProductGroupName;
	///商品代码
	TQFMatchCommodityIDType	CommodityID;
};
///产品
struct CQFMatchProductField
{
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///产品代码
	TQFMatchProductIDType	ProductID;
	///产品组代码
	TQFMatchProductGroupIDType	ProductGroupID;
	///产品名称
	TQFMatchProductNameType	ProductName;
	///产品类型
	TQFMatchProductClassType	ProductClass;
};
///合约
struct CQFMatchInstrumentField
{
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///产品代码
	TQFMatchProductIDType	ProductID;
	///产品组代码
	TQFMatchProductGroupIDType	ProductGroupID;
	///基础商品代码
	TQFMatchInstrumentIDType	UnderlyingInstrID;
	///产品类型
	TQFMatchProductClassType	ProductClass;
	///持仓类型
	TQFMatchPositionTypeType	PositionType;
	///执行价
	TQFMatchPriceType	StrikePrice;
	///期权类型
	TQFMatchOptionsTypeType	OptionsType;
	///合约数量乘数
	TQFMatchVolumeMultipleType	VolumeMultiple;
	///合约基础商品乘数
	TQFMatchUnderlyingMultipleType	UnderlyingMultiple;
	///合约代码
	TQFMatchInstrumentIDType	InstrumentID;
	///合约名称
	TQFMatchInstrumentNameType	InstrumentName;
	///交割年份
	TQFMatchYearType	DeliveryYear;
	///交割月
	TQFMatchMonthType	DeliveryMonth;
	///提前月份
	TQFMatchAdvanceMonthType	AdvanceMonth;
	///当前是否交易
	TQFMatchBoolType	IsTrading;
	///计价币种
	TQFMatchCurrencyType	PriceCurrency;
	///清算币种
	TQFMatchCurrencyType	ClearCurrency;
	///是否反向合约
	TQFMatchBoolType	IsInverse;
	///合约基础货币
	TQFMatchCurrencyType	BaseCurrency;
};
///组合交易合约的单腿
struct CQFMatchCombinationLegField
{
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///组合合约代码
	TQFMatchInstrumentIDType	CombInstrumentID;
	///单腿编号
	TQFMatchLegIDType	LegID;
	///单腿合约代码
	TQFMatchInstrumentIDType	LegInstrumentID;
	///买卖方向
	TQFMatchDirectionType	Direction;
	///单腿乘数
	TQFMatchLegMultipleType	LegMultiple;
	///推导层数
	TQFMatchImplyLevelType	ImplyLevel;
};
///会员合约持仓
struct CQFMatchPartPositionField
{
	///交易日
	TQFMatchDateType	TradingDay;
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///结算编号
	TQFMatchSettlementIDType	SettlementID;
	///投机套保标志
	TQFMatchHedgeFlagType	HedgeFlag;
	///持仓多空方向
	TQFMatchPosiDirectionType	PosiDirection;
	///上日持仓
	TQFMatchVolumeType	YdPosition;
	///今日持仓
	TQFMatchVolumeType	Position;
	///多头冻结
	TQFMatchVolumeType	LongFrozen;
	///空头冻结
	TQFMatchVolumeType	ShortFrozen;
	///昨日多头冻结
	TQFMatchVolumeType	YdLongFrozen;
	///昨日空头冻结
	TQFMatchVolumeType	YdShortFrozen;
	///合约代码
	TQFMatchInstrumentIDType	InstrumentID;
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///交易角色
	TQFMatchTradingRoleType	TradingRole;
};
///客户合约持仓
struct CQFMatchClientPositionField
{
	///交易日
	TQFMatchDateType	TradingDay;
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///结算编号
	TQFMatchSettlementIDType	SettlementID;
	///投机套保标志
	TQFMatchHedgeFlagType	HedgeFlag;
	///持仓多空方向
	TQFMatchPosiDirectionType	PosiDirection;
	///上日持仓
	TQFMatchVolumeType	YdPosition;
	///今日持仓
	TQFMatchVolumeType	Position;
	///多头冻结
	TQFMatchVolumeType	LongFrozen;
	///空头冻结
	TQFMatchVolumeType	ShortFrozen;
	///昨日多头冻结
	TQFMatchVolumeType	YdLongFrozen;
	///昨日空头冻结
	TQFMatchVolumeType	YdShortFrozen;
	///当日买成交量
	TQFMatchVolumeType	BuyTradeVolume;
	///当日卖成交量
	TQFMatchVolumeType	SellTradeVolume;
	///持仓成本
	TQFMatchMoneyType	PositionCost;
	///昨日持仓成本
	TQFMatchMoneyType	YdPositionCost;
	///占用的保证金
	TQFMatchMoneyType	UseMargin;
	///冻结的保证金
	TQFMatchMoneyType	FrozenMargin;
	///多头冻结的保证金
	TQFMatchMoneyType	LongFrozenMargin;
	///空头冻结的保证金
	TQFMatchMoneyType	ShortFrozenMargin;
	///冻结的权利金
	TQFMatchMoneyType	FrozenPremium;
	///合约代码
	TQFMatchInstrumentIDType	InstrumentID;
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///客户代码
	TQFMatchClientIDType	ClientID;
	///上次结算之后平仓盈亏
	TQFMatchMoneyType	CloseProfit;
	///持仓盈亏
	TQFMatchMoneyType	PositionProfit;
	///累计开仓成本
	TQFMatchMoneyType	TotalPositionCost;
	///累计平仓盈亏
	TQFMatchMoneyType	TotalCloseProfit;
	///累计持仓盈亏
	TQFMatchMoneyType	TotalPositionProfit;
	///杠杆倍数
	TQFMatchRatioType	Leverage;
	///接管价格
	TQFMatchRatioType	LiquidationPrice;
	///资金帐号
	TQFMatchAccountIDType	AccountID;
	///币种
	TQFMatchCurrencyType	ClearCurrency;
};
///保值额度量
struct CQFMatchHedgeVolumeField
{
	///交易日
	TQFMatchDateType	TradingDay;
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///结算编号
	TQFMatchSettlementIDType	SettlementID;
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///客户代码
	TQFMatchClientIDType	ClientID;
	///合约代码
	TQFMatchInstrumentIDType	InstrumentID;
	///多头保值额度最初申请量
	TQFMatchVolumeType	LongVolumeOriginal;
	///空头保值额度最初申请量
	TQFMatchVolumeType	ShortVolumeOriginal;
	///多头保值额度
	TQFMatchVolumeType	LongVolume;
	///空头保值额度
	TQFMatchVolumeType	ShortVolume;
};
///市场行情
struct CQFMatchMarketDataField
{
	///交易日
	TQFMatchDateType	TradingDay;
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///结算编号
	TQFMatchSettlementIDType	SettlementID;
	///最新价
	TQFMatchPriceType	LastPrice;
	///昨结算
	TQFMatchPriceType	PreSettlementPrice;
	///昨收盘
	TQFMatchPriceType	PreClosePrice;
	///昨持仓量
	TQFMatchLargeVolumeType	PreOpenInterest;
	///今开盘
	TQFMatchPriceType	OpenPrice;
	///最高价
	TQFMatchPriceType	HighestPrice;
	///最低价
	TQFMatchPriceType	LowestPrice;
	///数量
	TQFMatchVolumeType	Volume;
	///成交金额
	TQFMatchMoneyType	Turnover;
	///持仓量
	TQFMatchLargeVolumeType	OpenInterest;
	///今收盘
	TQFMatchPriceType	ClosePrice;
	///今结算
	TQFMatchPriceType	SettlementPrice;
	///涨停板价
	TQFMatchPriceType	UpperLimitPrice;
	///跌停板价
	TQFMatchPriceType	LowerLimitPrice;
	///昨虚实度
	TQFMatchRatioType	PreDelta;
	///今虚实度
	TQFMatchRatioType	CurrDelta;
	///最后修改时间
	TQFMatchTimeType	UpdateTime;
	///最后修改毫秒
	TQFMatchMillisecType	UpdateMillisec;
	///合约代码
	TQFMatchInstrumentIDType	InstrumentID;
	///资金费率
	TQFMatchPriceType	FundingRate;
	///合理价格
	TQFMatchPriceType	FairPrice;
};
///深度市场行情
struct CQFMatchDepthMarketDataField
{
	///交易日
	TQFMatchDateType	TradingDay;
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///结算编号
	TQFMatchSettlementIDType	SettlementID;
	///最新价
	TQFMatchPriceType	LastPrice;
	///昨结算
	TQFMatchPriceType	PreSettlementPrice;
	///昨收盘
	TQFMatchPriceType	PreClosePrice;
	///昨持仓量
	TQFMatchLargeVolumeType	PreOpenInterest;
	///今开盘
	TQFMatchPriceType	OpenPrice;
	///最高价
	TQFMatchPriceType	HighestPrice;
	///最低价
	TQFMatchPriceType	LowestPrice;
	///数量
	TQFMatchVolumeType	Volume;
	///成交金额
	TQFMatchMoneyType	Turnover;
	///持仓量
	TQFMatchLargeVolumeType	OpenInterest;
	///今收盘
	TQFMatchPriceType	ClosePrice;
	///今结算
	TQFMatchPriceType	SettlementPrice;
	///涨停板价
	TQFMatchPriceType	UpperLimitPrice;
	///跌停板价
	TQFMatchPriceType	LowerLimitPrice;
	///昨虚实度
	TQFMatchRatioType	PreDelta;
	///今虚实度
	TQFMatchRatioType	CurrDelta;
	///最后修改时间
	TQFMatchTimeType	UpdateTime;
	///最后修改毫秒
	TQFMatchMillisecType	UpdateMillisec;
	///合约代码
	TQFMatchInstrumentIDType	InstrumentID;
	///申买价一
	TQFMatchPriceType	BidPrice1;
	///申买量一
	TQFMatchVolumeType	BidVolume1;
	///申卖价一
	TQFMatchPriceType	AskPrice1;
	///申卖量一
	TQFMatchVolumeType	AskVolume1;
	///申买价二
	TQFMatchPriceType	BidPrice2;
	///申买量二
	TQFMatchVolumeType	BidVolume2;
	///申卖价二
	TQFMatchPriceType	AskPrice2;
	///申卖量二
	TQFMatchVolumeType	AskVolume2;
	///申买价三
	TQFMatchPriceType	BidPrice3;
	///申买量三
	TQFMatchVolumeType	BidVolume3;
	///申卖价三
	TQFMatchPriceType	AskPrice3;
	///申卖量三
	TQFMatchVolumeType	AskVolume3;
	///申买价四
	TQFMatchPriceType	BidPrice4;
	///申买量四
	TQFMatchVolumeType	BidVolume4;
	///申卖价四
	TQFMatchPriceType	AskPrice4;
	///申卖量四
	TQFMatchVolumeType	AskVolume4;
	///申买价五
	TQFMatchPriceType	BidPrice5;
	///申买量五
	TQFMatchVolumeType	BidVolume5;
	///申卖价五
	TQFMatchPriceType	AskPrice5;
	///申卖量五
	TQFMatchVolumeType	AskVolume5;
	///自然日期
	TQFMatchDateType	CalendarDate;
	///资金费率
	TQFMatchPriceType	FundingRate;
	///合理价格
	TQFMatchPriceType	FairPrice;
};
///分价表
struct CQFMatchMBLMarketDataField
{
	///合约代码
	TQFMatchInstrumentIDType	InstrumentID;
	///买卖方向
	TQFMatchDirectionType	Direction;
	///价格
	TQFMatchPriceType	Price;
	///数量
	TQFMatchVolumeType	Volume;
};
///别名定义
struct CQFMatchAliasDefineField
{
	///起始位置
	TQFMatchStartPosType	StartPos;
	///别名
	TQFMatchAliasType	Alias;
	///原文
	TQFMatchOriginalTextType	OriginalText;
};
///行情基础属性
struct CQFMatchMarketDataBaseField
{
	///交易日
	TQFMatchDateType	TradingDay;
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///结算编号
	TQFMatchSettlementIDType	SettlementID;
	///昨结算
	TQFMatchPriceType	PreSettlementPrice;
	///昨收盘
	TQFMatchPriceType	PreClosePrice;
	///昨持仓量
	TQFMatchLargeVolumeType	PreOpenInterest;
	///昨虚实度
	TQFMatchRatioType	PreDelta;
};
///行情静态属性
struct CQFMatchMarketDataStaticField
{
	///今开盘
	TQFMatchPriceType	OpenPrice;
	///最高价
	TQFMatchPriceType	HighestPrice;
	///最低价
	TQFMatchPriceType	LowestPrice;
	///今收盘
	TQFMatchPriceType	ClosePrice;
	///涨停板价
	TQFMatchPriceType	UpperLimitPrice;
	///跌停板价
	TQFMatchPriceType	LowerLimitPrice;
	///今结算
	TQFMatchPriceType	SettlementPrice;
	///今虚实度
	TQFMatchRatioType	CurrDelta;
};
///行情最新成交属性
struct CQFMatchMarketDataLastMatchField
{
	///最新价
	TQFMatchPriceType	LastPrice;
	///数量
	TQFMatchVolumeType	Volume;
	///成交金额
	TQFMatchMoneyType	Turnover;
	///持仓量
	TQFMatchLargeVolumeType	OpenInterest;
};
///行情最优价属性
struct CQFMatchMarketDataBestPriceField
{
	///申买价一
	TQFMatchPriceType	BidPrice1;
	///申买量一
	TQFMatchVolumeType	BidVolume1;
	///申卖价一
	TQFMatchPriceType	AskPrice1;
	///申卖量一
	TQFMatchVolumeType	AskVolume1;
};
///行情申买二、三属性
struct CQFMatchMarketDataBid23Field
{
	///申买价二
	TQFMatchPriceType	BidPrice2;
	///申买量二
	TQFMatchVolumeType	BidVolume2;
	///申买价三
	TQFMatchPriceType	BidPrice3;
	///申买量三
	TQFMatchVolumeType	BidVolume3;
};
///行情申卖二、三属性
struct CQFMatchMarketDataAsk23Field
{
	///申卖价二
	TQFMatchPriceType	AskPrice2;
	///申卖量二
	TQFMatchVolumeType	AskVolume2;
	///申卖价三
	TQFMatchPriceType	AskPrice3;
	///申卖量三
	TQFMatchVolumeType	AskVolume3;
};
///行情申买四、五属性
struct CQFMatchMarketDataBid45Field
{
	///申买价四
	TQFMatchPriceType	BidPrice4;
	///申买量四
	TQFMatchVolumeType	BidVolume4;
	///申买价五
	TQFMatchPriceType	BidPrice5;
	///申买量五
	TQFMatchVolumeType	BidVolume5;
};
///行情申卖四、五属性
struct CQFMatchMarketDataAsk45Field
{
	///申卖价四
	TQFMatchPriceType	AskPrice4;
	///申卖量四
	TQFMatchVolumeType	AskVolume4;
	///申卖价五
	TQFMatchPriceType	AskPrice5;
	///申卖量五
	TQFMatchVolumeType	AskVolume5;
};
///行情更新时间属性
struct CQFMatchMarketDataUpdateTimeField
{
	///合约代码
	TQFMatchInstrumentIDType	InstrumentID;
	///最后修改时间
	TQFMatchTimeType	UpdateTime;
	///最后修改毫秒
	TQFMatchMillisecType	UpdateMillisec;
	///Calendar Date
	TQFMatchDateType	CalendarDate;
};
///报价
struct CQFMatchQuoteField
{
	///交易日
	TQFMatchDateType	TradingDay;
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///结算编号
	TQFMatchSettlementIDType	SettlementID;
	///报价编号
	TQFMatchQuoteSysIDType	QuoteSysID;
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///客户代码
	TQFMatchClientIDType	ClientID;
	///交易用户代码
	TQFMatchUserIDType	UserID;
	///数量
	TQFMatchVolumeType	Volume;
	///合约代码
	TQFMatchInstrumentIDType	InstrumentID;
	///本地报价编号
	TQFMatchOrderLocalIDType	QuoteLocalID;
	///业务单元
	TQFMatchBusinessUnitType	BusinessUnit;
	///买方组合开平标志
	TQFMatchCombOffsetFlagType	BidCombOffsetFlag;
	///买方组合套保标志
	TQFMatchCombHedgeFlagType	BidCombHedgeFlag;
	///买方价格
	TQFMatchPriceType	BidPrice;
	///卖方组合开平标志
	TQFMatchCombOffsetFlagType	AskCombOffsetFlag;
	///卖方组合套保标志
	TQFMatchCombHedgeFlagType	AskCombHedgeFlag;
	///卖方价格
	TQFMatchPriceType	AskPrice;
	///插入时间
	TQFMatchTimeType	InsertTime;
	///撤销时间
	TQFMatchTimeType	CancelTime;
	///成交时间
	TQFMatchTimeType	TradeTime;
	///买方报单编号
	TQFMatchOrderSysIDType	BidOrderSysID;
	///卖方报单编号
	TQFMatchOrderSysIDType	AskOrderSysID;
	///结算会员编号
	TQFMatchParticipantIDType	ClearingPartID;
	///自然日期
	TQFMatchDateType	CalendarDate;
};
///成交
struct CQFMatchTradeField
{
	///交易日
	TQFMatchDateType	TradingDay;
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///结算编号
	TQFMatchSettlementIDType	SettlementID;
	///成交编号
	TQFMatchTradeIDType	TradeID;
	///买卖方向
	TQFMatchDirectionType	Direction;
	///报单编号
	TQFMatchOrderSysIDType	OrderSysID;
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///客户代码
	TQFMatchClientIDType	ClientID;
	///交易角色
	TQFMatchTradingRoleType	TradingRole;
	///资金帐号
	TQFMatchAccountIDType	AccountID;
	///合约代码
	TQFMatchInstrumentIDType	InstrumentID;
	///开平标志
	TQFMatchOffsetFlagType	OffsetFlag;
	///投机套保标志
	TQFMatchHedgeFlagType	HedgeFlag;
	///价格
	TQFMatchPriceType	Price;
	///数量
	TQFMatchVolumeType	Volume;
	///成交时间
	TQFMatchTimeType	TradeTime;
	///成交类型
	TQFMatchTradeTypeType	TradeType;
	///成交价来源
	TQFMatchPriceSourceType	PriceSource;
	///交易用户代码
	TQFMatchUserIDType	UserID;
	///本地报单编号
	TQFMatchOrderLocalIDType	OrderLocalID;
	///结算会员编号
	TQFMatchParticipantIDType	ClearingPartID;
	///业务单元
	TQFMatchBusinessUnitType	BusinessUnit;
	///自然日期
	TQFMatchDateType	CalendarDate;
	///成交角色
	TQFMatchMatchRoleType	MatchRole;
	///手续费
	TQFMatchMoneyType	Fee;
	///平仓盈亏
	TQFMatchMoneyType	CloseProfit;
};
///报单
struct CQFMatchOrderField
{
	///交易日
	TQFMatchDateType	TradingDay;
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///结算编号
	TQFMatchSettlementIDType	SettlementID;
	///报单编号
	TQFMatchOrderSysIDType	OrderSysID;
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///客户代码
	TQFMatchClientIDType	ClientID;
	///交易用户代码
	TQFMatchUserIDType	UserID;
	///合约代码
	TQFMatchInstrumentIDType	InstrumentID;
	///报单价格条件
	TQFMatchOrderPriceTypeType	OrderPriceType;
	///买卖方向
	TQFMatchDirectionType	Direction;
	///组合开平标志
	TQFMatchCombOffsetFlagType	CombOffsetFlag;
	///组合投机套保标志
	TQFMatchCombHedgeFlagType	CombHedgeFlag;
	///价格
	TQFMatchPriceType	LimitPrice;
	///数量
	TQFMatchVolumeType	VolumeTotalOriginal;
	///有效期类型
	TQFMatchTimeConditionType	TimeCondition;
	///GTD日期
	TQFMatchDateType	GTDDate;
	///成交量类型
	TQFMatchVolumeConditionType	VolumeCondition;
	///最小成交量
	TQFMatchVolumeType	MinVolume;
	///触发条件
	TQFMatchContingentConditionType	ContingentCondition;
	///止损价
	TQFMatchPriceType	StopPrice;
	///强平原因
	TQFMatchForceCloseReasonType	ForceCloseReason;
	///本地报单编号
	TQFMatchOrderLocalIDType	OrderLocalID;
	///自动挂起标志
	TQFMatchBoolType	IsAutoSuspend;
	///报单来源
	TQFMatchOrderSourceType	OrderSource;
	///报单状态
	TQFMatchOrderStatusType	OrderStatus;
	///报单类型
	TQFMatchOrderTypeType	OrderType;
	///今成交数量
	TQFMatchVolumeType	VolumeTraded;
	///剩余数量
	TQFMatchVolumeType	VolumeTotal;
	///报单日期
	TQFMatchDateType	InsertDate;
	///插入时间
	TQFMatchTimeType	InsertTime;
	///激活时间
	TQFMatchTimeType	ActiveTime;
	///挂起时间
	TQFMatchTimeType	SuspendTime;
	///最后修改时间
	TQFMatchTimeType	UpdateTime;
	///撤销时间
	TQFMatchTimeType	CancelTime;
	///最后修改交易用户代码
	TQFMatchUserIDType	ActiveUserID;
	///优先权
	TQFMatchPriorityType	Priority;
	///按时间排队的序号
	TQFMatchTimeSortIDType	TimeSortID;
	///结算会员编号
	TQFMatchParticipantIDType	ClearingPartID;
	///业务单元
	TQFMatchBusinessUnitType	BusinessUnit;
	///自然日期
	TQFMatchDateType	CalendarDate;
};
///执行宣告
struct CQFMatchExecOrderField
{
	///交易日
	TQFMatchDateType	TradingDay;
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///结算编号
	TQFMatchSettlementIDType	SettlementID;
	///合约编号
	TQFMatchInstrumentIDType	InstrumentID;
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///客户代码
	TQFMatchClientIDType	ClientID;
	///交易用户代码
	TQFMatchUserIDType	UserID;
	///本地执行宣告编号
	TQFMatchOrderLocalIDType	ExecOrderLocalID;
	///数量
	TQFMatchVolumeType	Volume;
	///业务单元
	TQFMatchBusinessUnitType	BusinessUnit;
	///执行宣告编号
	TQFMatchExecOrderSysIDType	ExecOrderSysID;
	///报单日期
	TQFMatchDateType	InsertDate;
	///插入时间
	TQFMatchTimeType	InsertTime;
	///撤销时间
	TQFMatchTimeType	CancelTime;
	///执行结果
	TQFMatchExecResultType	ExecResult;
	///结算会员编号
	TQFMatchParticipantIDType	ClearingPartID;
	///自然日期
	TQFMatchDateType	CalendarDate;
};
///非标组合报单
struct CQFMatchCombOrderField
{
	///交易日
	TQFMatchDateType	TradingDay;
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///结算编号
	TQFMatchSettlementIDType	SettlementID;
	///组合报单编号
	TQFMatchOrderSysIDType	CombOrderSysID;
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///客户代码
	TQFMatchClientIDType	ClientID;
	///交易用户代码
	TQFMatchUserIDType	UserID;
	///价格
	TQFMatchPriceType	LimitPrice;
	///数量
	TQFMatchVolumeType	VolumeTotalOriginal;
	///本地报单编号
	TQFMatchOrderLocalIDType	CombOrderLocalID;
	///业务单元
	TQFMatchBusinessUnitType	BusinessUnit;
	///合约代码1
	TQFMatchInstrumentIDType	InstrumentID1;
	///买卖方向1
	TQFMatchDirectionType	Direction1;
	///分腿乘数1
	TQFMatchLegMultipleType	LegMultiple1;
	///开平标志1
	TQFMatchOffsetFlagType	OffsetFlag1;
	///投机套保标志1
	TQFMatchHedgeFlagType	HedgeFlag1;
	///合约代码2
	TQFMatchInstrumentIDType	InstrumentID2;
	///买卖方向2
	TQFMatchDirectionType	Direction2;
	///分腿乘数2
	TQFMatchLegMultipleType	LegMultiple2;
	///开平标志2
	TQFMatchOffsetFlagType	OffsetFlag2;
	///投机套保标志2
	TQFMatchHedgeFlagType	HedgeFlag2;
	///合约代码3
	TQFMatchInstrumentIDType	InstrumentID3;
	///买卖方向3
	TQFMatchDirectionType	Direction3;
	///分腿乘数3
	TQFMatchLegMultipleType	LegMultiple3;
	///开平标志3
	TQFMatchOffsetFlagType	OffsetFlag3;
	///投机套保标志3
	TQFMatchHedgeFlagType	HedgeFlag3;
	///合约代码4
	TQFMatchInstrumentIDType	InstrumentID4;
	///买卖方向4
	TQFMatchDirectionType	Direction4;
	///分腿乘数4
	TQFMatchLegMultipleType	LegMultiple4;
	///开平标志4
	TQFMatchOffsetFlagType	OffsetFlag4;
	///投机套保标志4
	TQFMatchHedgeFlagType	HedgeFlag4;
	///报单来源
	TQFMatchOrderSourceType	OrderSource;
	///今成交数量
	TQFMatchVolumeType	VolumeTraded;
	///剩余数量
	TQFMatchVolumeType	VolumeTotal;
	///报单日期
	TQFMatchDateType	InsertDate;
	///插入时间
	TQFMatchTimeType	InsertTime;
	///结算会员编号
	TQFMatchParticipantIDType	ClearingPartID;
	///自然日期
	TQFMatchDateType	CalendarDate;
};
///资金
struct CQFMatchTradingAccountField
{
	///交易日
	TQFMatchDateType	TradingDay;
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///结算编号
	TQFMatchSettlementIDType	SettlementID;
	///上次结算准备金
	TQFMatchMoneyType	PreBalance;
	///当前保证金总额
	TQFMatchMoneyType	CurrMargin;
	///平仓盈亏
	TQFMatchMoneyType	CloseProfit;
	///期权权利金收支
	TQFMatchMoneyType	Premium;
	///入金金额
	TQFMatchMoneyType	Deposit;
	///出金金额
	TQFMatchMoneyType	Withdraw;
	///期货结算准备金
	TQFMatchMoneyType	Balance;
	///可用资金
	TQFMatchMoneyType	Available;
	///资金帐号
	TQFMatchAccountIDType	AccountID;
	///冻结的保证金
	TQFMatchMoneyType	FrozenMargin;
	///冻结的权利金
	TQFMatchMoneyType	FrozenPremium;
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///币种
	TQFMatchCurrencyType	Currency;
	///可取资金
	TQFMatchMoneyType	Useable;
	///冻结资金
	TQFMatchMoneyType	FrozenMoney;
	///手续费支出
	TQFMatchMoneyType	Fee;
};
///管理报单
struct CQFMatchAdminOrderField
{
	///合约代码
	TQFMatchInstrumentIDType	InstrumentID;
	///管理报单命令
	TQFMatchAdminOrderCommandFlagType	AdminOrderCommand;
	///结算会员编号
	TQFMatchParticipantIDType	ClearingPartID;
	///交易会员编号
	TQFMatchParticipantIDType	ParticipantID;
	///金额
	TQFMatchMoneyType	Amount;
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
};
///管理报单输入
struct CQFMatchInputAdminOrderField
{
	///合约代码
	TQFMatchInstrumentIDType	InstrumentID;
	///管理报单命令
	TQFMatchAdminOrderCommandFlagType	AdminOrderCommand;
	///结算会员编号
	TQFMatchParticipantIDType	ClearingPartID;
	///交易会员编号
	TQFMatchParticipantIDType	ParticipantID;
	///金额
	TQFMatchMoneyType	Amount;
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
};
///公告
struct CQFMatchBulletinField
{
	///交易日
	TQFMatchDateType	TradingDay;
	///公告编号
	TQFMatchBulletinIDType	BulletinID;
	///序列号
	TQFMatchSequenceNoType	SequenceNo;
	///公告类型
	TQFMatchNewsTypeType	NewsType;
	///紧急程度
	TQFMatchNewsUrgencyType	NewsUrgency;
	///发送时间
	TQFMatchTimeType	SendTime;
	///消息摘要
	TQFMatchAbstractType	Abstract;
	///消息来源
	TQFMatchComeFromType	ComeFrom;
	///消息正文
	TQFMatchContentType	Content;
	///WEB地址
	TQFMatchURLLinkType	URLLink;
	///市场代码
	TQFMatchMarketIDType	MarketID;
	///自然日期
	TQFMatchDateType	CalendarDate;
};
///交易所数据同步状态
struct CQFMatchExchangeDataSyncStatusField
{
	///交易日
	TQFMatchDateType	TradingDay;
	///交易所代码
	TQFMatchExchangeIDType	ExchangeID;
	///交易所数据同步状态
	TQFMatchExchangeDataSyncStatusType	ExchangeDataSyncStatus;
};
///结算组数据同步状态
struct CQFMatchSGDataSyncStatusField
{
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///交易日
	TQFMatchDateType	TradingDay;
	///结算编号
	TQFMatchSettlementIDType	SettlementID;
	///结算组数据同步状态
	TQFMatchSGDataSyncStatusType	SGDataSyncStatus;
};
///合约状态
struct CQFMatchInstrumentStatusField
{
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///合约代码
	TQFMatchInstrumentIDType	InstrumentID;
	///合约交易状态
	TQFMatchInstrumentStatusType	InstrumentStatus;
	///交易阶段编号
	TQFMatchTradingSegmentSNType	TradingSegmentSN;
	///进入本状态时间
	TQFMatchTimeType	EnterTime;
	///进入本状态原因
	TQFMatchInstStatusEnterReasonType	EnterReason;
	///自然日期
	TQFMatchDateType	CalendarDate;
};
///实时上场
struct CQFMatchTBCommandField
{
	///DB命令本地号
	TQFMatchOrderLocalIDType	CommandLocalID;
	///全局DB命令序号
	TQFMatchSequenceNoType	CommandNo;
	///DB命令类型
	TQFMatchCommandTypeType	CommandType;
	///字段名
	TQFMatchFieldNameType	FieldName;
	///字段内容
	TQFMatchFieldContentType	FieldContent;
	///备注
	TQFMatchComeFromType	Comments;
};
///报单查询
struct CQFMatchQryOrderField
{
	///起始会员代码
	TQFMatchParticipantIDType	PartIDStart;
	///结束会员代码
	TQFMatchParticipantIDType	PartIDEnd;
	///报单编号
	TQFMatchOrderSysIDType	OrderSysID;
	///合约代码
	TQFMatchInstrumentIDType	InstrumentID;
	///客户代码
	TQFMatchClientIDType	ClientID;
	///交易用户代码
	TQFMatchUserIDType	UserID;
	///开始时间
	TQFMatchTimeType	TimeStart;
	///结束时间
	TQFMatchTimeType	TimeEnd;
};
///报价查询
struct CQFMatchQryQuoteField
{
	///起始会员代码
	TQFMatchParticipantIDType	PartIDStart;
	///结束会员代码
	TQFMatchParticipantIDType	PartIDEnd;
	///报价编号
	TQFMatchQuoteSysIDType	QuoteSysID;
	///客户代码
	TQFMatchClientIDType	ClientID;
	///合约代码
	TQFMatchInstrumentIDType	InstrumentID;
	///交易用户代码
	TQFMatchUserIDType	UserID;
};
///成交查询
struct CQFMatchQryTradeField
{
	///起始会员代码
	TQFMatchParticipantIDType	PartIDStart;
	///结束会员代码
	TQFMatchParticipantIDType	PartIDEnd;
	///起始合约代码
	TQFMatchInstrumentIDType	InstIDStart;
	///结束合约代码
	TQFMatchInstrumentIDType	InstIDEnd;
	///成交编号
	TQFMatchTradeIDType	TradeID;
	///客户代码
	TQFMatchClientIDType	ClientID;
	///交易用户代码
	TQFMatchUserIDType	UserID;
	///开始时间
	TQFMatchTimeType	TimeStart;
	///结束时间
	TQFMatchTimeType	TimeEnd;
};
///行情查询
struct CQFMatchQryMarketDataField
{
	///产品代码
	TQFMatchProductIDType	ProductID;
	///合约代码
	TQFMatchInstrumentIDType	InstrumentID;
};
///客户查询
struct CQFMatchQryClientField
{
	///起始会员代码
	TQFMatchParticipantIDType	PartIDStart;
	///结束会员代码
	TQFMatchParticipantIDType	PartIDEnd;
	///起始客户代码
	TQFMatchClientIDType	ClientIDStart;
	///结束客户代码
	TQFMatchClientIDType	ClientIDEnd;
};
///会员持仓查询
struct CQFMatchQryPartPositionField
{
	///起始会员代码
	TQFMatchParticipantIDType	PartIDStart;
	///结束会员代码
	TQFMatchParticipantIDType	PartIDEnd;
	///起始合约代码
	TQFMatchInstrumentIDType	InstIDStart;
	///结束合约代码
	TQFMatchInstrumentIDType	InstIDEnd;
};
///客户持仓查询
struct CQFMatchQryClientPositionField
{
	///起始会员代码
	TQFMatchParticipantIDType	PartIDStart;
	///结束会员代码
	TQFMatchParticipantIDType	PartIDEnd;
	///起始客户代码
	TQFMatchClientIDType	ClientIDStart;
	///结束客户代码
	TQFMatchClientIDType	ClientIDEnd;
	///起始合约代码
	TQFMatchInstrumentIDType	InstIDStart;
	///结束合约代码
	TQFMatchInstrumentIDType	InstIDEnd;
	///客户类型
	TQFMatchClientTypeType	ClientType;
};
///客户持仓查询
struct CQFMatchQryClientPositionV1Field
{
	///起始会员代码
	TQFMatchParticipantIDType	PartIDStart;
	///结束会员代码
	TQFMatchParticipantIDType	PartIDEnd;
	///起始客户代码
	TQFMatchClientIDType	ClientIDStart;
	///结束客户代码
	TQFMatchClientIDType	ClientIDEnd;
	///起始合约代码
	TQFMatchInstrumentIDType	InstIDStart;
	///结束合约代码
	TQFMatchInstrumentIDType	InstIDEnd;
};
///交易资金查询
struct CQFMatchQryPartAccountField
{
	///起始会员代码
	TQFMatchParticipantIDType	PartIDStart;
	///结束会员代码
	TQFMatchParticipantIDType	PartIDEnd;
	///资金帐号
	TQFMatchAccountIDType	AccountID;
	///币种
	TQFMatchCurrencyType	Currency;
};
///合约查询
struct CQFMatchQryInstrumentField
{
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///产品组代码
	TQFMatchProductGroupIDType	ProductGroupID;
	///产品代码
	TQFMatchProductIDType	ProductID;
	///合约代码
	TQFMatchInstrumentIDType	InstrumentID;
};
///合约状态查询
struct CQFMatchQryInstrumentStatusField
{
	///起始合约代码
	TQFMatchInstrumentIDType	InstIDStart;
	///结束合约代码
	TQFMatchInstrumentIDType	InstIDEnd;
};
///结算组状态查询
struct CQFMatchQrySGDataSyncStatusField
{
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
};
///用户在线查询
struct CQFMatchQryUserSessionField
{
	///起始交易用户代码
	TQFMatchUserIDType	UserIDStart;
	///结束交易用户代码
	TQFMatchUserIDType	UserIDEnd;
};
///用户查询
struct CQFMatchQryUserField
{
	///起始交易用户代码
	TQFMatchUserIDType	UserIDStart;
	///结束交易用户代码
	TQFMatchUserIDType	UserIDEnd;
};
///公告查询
struct CQFMatchQryBulletinField
{
	///交易日
	TQFMatchDateType	TradingDay;
	///市场代码
	TQFMatchMarketIDType	MarketID;
	///公告编号
	TQFMatchBulletinIDType	BulletinID;
	///公告类型
	TQFMatchNewsTypeType	NewsType;
	///紧急程度
	TQFMatchNewsUrgencyType	NewsUrgency;
};
///会员查询
struct CQFMatchQryParticipantField
{
	///起始会员代码
	TQFMatchParticipantIDType	PartIDStart;
	///结束会员代码
	TQFMatchParticipantIDType	PartIDEnd;
};
///保值额度查询
struct CQFMatchQryHedgeVolumeField
{
	///起始会员代码
	TQFMatchParticipantIDType	PartIDStart;
	///结束会员代码
	TQFMatchParticipantIDType	PartIDEnd;
	///起始客户代码
	TQFMatchClientIDType	ClientIDStart;
	///结束客户代码
	TQFMatchClientIDType	ClientIDEnd;
	///起始合约代码
	TQFMatchInstrumentIDType	InstIDStart;
	///结束合约代码
	TQFMatchInstrumentIDType	InstIDEnd;
};
///合约价位查询
struct CQFMatchQryMBLMarketDataField
{
	///起始合约代码
	TQFMatchInstrumentIDType	InstIDStart;
	///结束合约代码
	TQFMatchInstrumentIDType	InstIDEnd;
	///买卖方向
	TQFMatchDirectionType	Direction;
};
///信用限额查询
struct CQFMatchQryCreditLimitField
{
	///交易会员编号
	TQFMatchParticipantIDType	ParticipantID;
	///结算会员编号
	TQFMatchParticipantIDType	ClearingPartID;
};
///非标组合报单查询
struct CQFMatchQryCombOrderField
{
	///起始会员代码
	TQFMatchParticipantIDType	PartIDStart;
	///结束会员代码
	TQFMatchParticipantIDType	PartIDEnd;
	///组合报单编号
	TQFMatchOrderSysIDType	CombOrderSysID;
	///客户代码
	TQFMatchClientIDType	ClientID;
	///交易用户代码
	TQFMatchUserIDType	UserID;
};
///会员资金应答
struct CQFMatchRspPartAccountField
{
	///交易日
	TQFMatchDateType	TradingDay;
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///结算编号
	TQFMatchSettlementIDType	SettlementID;
	///上次结算准备金
	TQFMatchMoneyType	PreBalance;
	///当前保证金总额
	TQFMatchMoneyType	CurrMargin;
	///平仓盈亏
	TQFMatchMoneyType	CloseProfit;
	///期权权利金收支
	TQFMatchMoneyType	Premium;
	///入金金额
	TQFMatchMoneyType	Deposit;
	///出金金额
	TQFMatchMoneyType	Withdraw;
	///期货结算准备金
	TQFMatchMoneyType	Balance;
	///可用资金
	TQFMatchMoneyType	Available;
	///资金帐号
	TQFMatchAccountIDType	AccountID;
	///冻结的保证金
	TQFMatchMoneyType	FrozenMargin;
	///冻结的权利金
	TQFMatchMoneyType	FrozenPremium;
	///基本准备金
	TQFMatchMoneyType	BaseReserve;
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///可取资金
	TQFMatchMoneyType	Useable;
	///冻结资金
	TQFMatchMoneyType	FrozenMoney;
	///币种
	TQFMatchCurrencyType	Currency;
	///手续费支出
	TQFMatchMoneyType	Fee;
};
///会员持仓应答
struct CQFMatchRspPartPositionField
{
	///交易日
	TQFMatchDateType	TradingDay;
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///结算编号
	TQFMatchSettlementIDType	SettlementID;
	///投机套保标志
	TQFMatchHedgeFlagType	HedgeFlag;
	///持仓多空方向
	TQFMatchPosiDirectionType	PosiDirection;
	///上日持仓
	TQFMatchVolumeType	YdPosition;
	///今日持仓
	TQFMatchVolumeType	Position;
	///多头冻结
	TQFMatchVolumeType	LongFrozen;
	///空头冻结
	TQFMatchVolumeType	ShortFrozen;
	///昨日多头冻结
	TQFMatchVolumeType	YdLongFrozen;
	///昨日空头冻结
	TQFMatchVolumeType	YdShortFrozen;
	///合约代码
	TQFMatchInstrumentIDType	InstrumentID;
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///交易角色
	TQFMatchTradingRoleType	TradingRole;
};
///客户持仓应答
struct CQFMatchRspClientPositionField
{
	///交易日
	TQFMatchDateType	TradingDay;
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///结算编号
	TQFMatchSettlementIDType	SettlementID;
	///投机套保标志
	TQFMatchHedgeFlagType	HedgeFlag;
	///持仓多空方向
	TQFMatchPosiDirectionType	PosiDirection;
	///上日持仓
	TQFMatchVolumeType	YdPosition;
	///今日持仓
	TQFMatchVolumeType	Position;
	///多头冻结
	TQFMatchVolumeType	LongFrozen;
	///空头冻结
	TQFMatchVolumeType	ShortFrozen;
	///昨日多头冻结
	TQFMatchVolumeType	YdLongFrozen;
	///昨日空头冻结
	TQFMatchVolumeType	YdShortFrozen;
	///当日买成交量
	TQFMatchVolumeType	BuyTradeVolume;
	///当日卖成交量
	TQFMatchVolumeType	SellTradeVolume;
	///持仓成本
	TQFMatchMoneyType	PositionCost;
	///昨日持仓成本
	TQFMatchMoneyType	YdPositionCost;
	///占用的保证金
	TQFMatchMoneyType	UseMargin;
	///冻结的保证金
	TQFMatchMoneyType	FrozenMargin;
	///多头冻结的保证金
	TQFMatchMoneyType	LongFrozenMargin;
	///空头冻结的保证金
	TQFMatchMoneyType	ShortFrozenMargin;
	///冻结的权利金
	TQFMatchMoneyType	FrozenPremium;
	///合约代码
	TQFMatchInstrumentIDType	InstrumentID;
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///客户代码
	TQFMatchClientIDType	ClientID;
	///上次结算之后平仓盈亏
	TQFMatchMoneyType	CloseProfit;
	///持仓盈亏
	TQFMatchMoneyType	PositionProfit;
	///累计开仓成本
	TQFMatchMoneyType	TotalPositionCost;
	///累计平仓盈亏
	TQFMatchMoneyType	TotalCloseProfit;
	///累计持仓盈亏
	TQFMatchMoneyType	TotalPositionProfit;
	///杠杆倍数
	TQFMatchRatioType	Leverage;
	///接管价格
	TQFMatchRatioType	LiquidationPrice;
	///资金帐号
	TQFMatchAccountIDType	AccountID;
	///币种
	TQFMatchCurrencyType	ClearCurrency;
};
///合约查询应答
struct CQFMatchRspInstrumentField
{
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///产品代码
	TQFMatchProductIDType	ProductID;
	///产品组代码
	TQFMatchProductGroupIDType	ProductGroupID;
	///基础商品代码
	TQFMatchInstrumentIDType	UnderlyingInstrID;
	///产品类型
	TQFMatchProductClassType	ProductClass;
	///持仓类型
	TQFMatchPositionTypeType	PositionType;
	///执行价
	TQFMatchPriceType	StrikePrice;
	///期权类型
	TQFMatchOptionsTypeType	OptionsType;
	///合约数量乘数
	TQFMatchVolumeMultipleType	VolumeMultiple;
	///合约基础商品乘数
	TQFMatchUnderlyingMultipleType	UnderlyingMultiple;
	///合约代码
	TQFMatchInstrumentIDType	InstrumentID;
	///合约名称
	TQFMatchInstrumentNameType	InstrumentName;
	///交割年份
	TQFMatchYearType	DeliveryYear;
	///交割月
	TQFMatchMonthType	DeliveryMonth;
	///提前月份
	TQFMatchAdvanceMonthType	AdvanceMonth;
	///当前是否交易
	TQFMatchBoolType	IsTrading;
	///创建日
	TQFMatchDateType	CreateDate;
	///上市日
	TQFMatchDateType	OpenDate;
	///到期日
	TQFMatchDateType	ExpireDate;
	///开始交割日
	TQFMatchDateType	StartDelivDate;
	///最后交割日
	TQFMatchDateType	EndDelivDate;
	///挂牌基准价
	TQFMatchPriceType	BasisPrice;
	///市价单最大下单量
	TQFMatchVolumeType	MaxMarketOrderVolume;
	///市价单最小下单量
	TQFMatchVolumeType	MinMarketOrderVolume;
	///限价单最大下单量
	TQFMatchVolumeType	MaxLimitOrderVolume;
	///限价单最小下单量
	TQFMatchVolumeType	MinLimitOrderVolume;
	///最小变动价位
	TQFMatchPriceType	PriceTick;
	///交割月自然人开仓
	TQFMatchMonthCountType	AllowDelivPersonOpen;
	///币种
	TQFMatchCurrencyType	Currency;
};
///信息查询
struct CQFMatchQryInformationField
{
	///起始信息代码
	TQFMatchInformationIDType	InformationIDStart;
	///结束信息代码
	TQFMatchInformationIDType	InformationIDEnd;
};
///信息查询
struct CQFMatchInformationField
{
	///信息编号
	TQFMatchInformationIDType	InformationID;
	///序列号
	TQFMatchSequenceNoType	SequenceNo;
	///消息正文
	TQFMatchContentType	Content;
	///正文长度
	TQFMatchContentLengthType	ContentLength;
	///是否完成
	TQFMatchBoolType	IsAccomplished;
};
///信用限额
struct CQFMatchCreditLimitField
{
	///交易日
	TQFMatchDateType	TradingDay;
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///结算编号
	TQFMatchSettlementIDType	SettlementID;
	///上次结算准备金
	TQFMatchMoneyType	PreBalance;
	///当前保证金总额
	TQFMatchMoneyType	CurrMargin;
	///平仓盈亏
	TQFMatchMoneyType	CloseProfit;
	///期权权利金收支
	TQFMatchMoneyType	Premium;
	///入金金额
	TQFMatchMoneyType	Deposit;
	///出金金额
	TQFMatchMoneyType	Withdraw;
	///期货结算准备金
	TQFMatchMoneyType	Balance;
	///可用资金
	TQFMatchMoneyType	Available;
	///交易会员编号
	TQFMatchParticipantIDType	ParticipantID;
	///结算会员编号
	TQFMatchParticipantIDType	ClearingPartID;
	///冻结的保证金
	TQFMatchMoneyType	FrozenMargin;
	///冻结的权利金
	TQFMatchMoneyType	FrozenPremium;
	///手续费支出
	TQFMatchMoneyType	Fee;
};
///客户查询应答
struct CQFMatchRspClientField
{
	///客户代码
	TQFMatchClientIDType	ClientID;
	///客户名称
	TQFMatchPartyNameType	ClientName;
	///证件类型
	TQFMatchIdCardTypeType	IdentifiedCardType;
	///原证件号码
	TQFMatchIdentifiedCardNoV1Type	UseLess;
	///交易角色
	TQFMatchTradingRoleType	TradingRole;
	///客户类型
	TQFMatchClientTypeType	ClientType;
	///是否活跃
	TQFMatchBoolType	IsActive;
	///会员号
	TQFMatchParticipantIDType	ParticipantID;
	///证件号码
	TQFMatchIdentifiedCardNoType	IdentifiedCardNo;
};
///数据流回退
struct CQFMatchFlowMessageCancelField
{
	///序列系列号
	TQFMatchSequenceSeriesType	SequenceSeries;
	///交易日
	TQFMatchDateType	TradingDay;
	///数据中心代码
	TQFMatchDataCenterIDType	DataCenterID;
	///回退起始序列号
	TQFMatchSequenceNoType	StartSequenceNo;
	///回退结束序列号
	TQFMatchSequenceNoType	EndSequenceNo;
};
///心跳
struct CQFMatchMultiHeartbeatField
{
	///Current Time
	TQFMatchTimeType	CurrTime;
	///MultiCastIP
	TQFMatchMultiIPAddressType	MultiCastIP;
};
///出入金查询
struct CQFMatchQryAccountDepositDetailField
{
	///出入金流水序号
	TQFMatchSequenceNoType	DepositNo;
	///起始会员代码
	TQFMatchAccountIDType	AccountIDStart;
	///结束会员代码
	TQFMatchAccountIDType	AccountIDEnd;
};
///实时上场查询
struct CQFMatchQryTBCommandDetailField
{
	///全局DB命令序号
	TQFMatchSequenceNoType	CommandNo;
	///DB命令类型
	TQFMatchCommandTypeType	CommandType;
};
///查询手续费
struct CQFMatchQryPartClientFeeField
{
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///客户编码
	TQFMatchClientIDType	ClientID;
	///产品合约代码
	TQFMatchInstrumentIDType	InstrumentID;
};
///查询客户持仓明细
struct CQFMatchQryClientPositionDetailField
{
	///起始会员代码
	TQFMatchParticipantIDType	PartIDStart;
	///结束会员代码
	TQFMatchParticipantIDType	PartIDEnd;
	///起始客户代码
	TQFMatchClientIDType	ClientIDStart;
	///结束客户代码
	TQFMatchClientIDType	ClientIDEnd;
	///起始合约代码
	TQFMatchInstrumentIDType	InstIDStart;
	///结束合约代码
	TQFMatchInstrumentIDType	InstIDEnd;
};
///输入场外成交对
struct CQFMatchInputOTCTradeField
{
	///成交编号
	TQFMatchTradeIDType	TradeID;
	///交易用户代码
	TQFMatchUserIDType	UserID;
	///成交对本地编号
	TQFMatchOrderLocalIDType	TradeLocalID;
	///合约代码
	TQFMatchInstrumentIDType	InstrumentID;
	///价格
	TQFMatchPriceType	Price;
	///数量
	TQFMatchVolumeType	Volume;
	///买方会员代码
	TQFMatchParticipantIDType	BuyParticipantID;
	///买方客户代码
	TQFMatchClientIDType	BuyClientID;
	///买方交易用户代码
	TQFMatchUserIDType	BuyUserID;
	///买方本地报单编号
	TQFMatchOrderLocalIDType	BuyOrderLocalID;
	///买方组合开平标志
	TQFMatchCombOffsetFlagType	BuyCombOffsetFlag;
	///买方组合投机套保标志
	TQFMatchCombHedgeFlagType	BuyCombHedgeFlag;
	///卖方方会员代码
	TQFMatchParticipantIDType	SellParticipantID;
	///卖方客户代码
	TQFMatchClientIDType	SellClientID;
	///卖方交易用户代码
	TQFMatchUserIDType	SellUserID;
	///卖方组合开平标志
	TQFMatchCombOffsetFlagType	SellCombOffsetFlag;
	///卖方组合套保标志
	TQFMatchCombHedgeFlagType	SellCombHedgeFlag;
	///卖方本地报单编号
	TQFMatchOrderLocalIDType	SellOrderLocalID;
	///OTC类型
	TQFMatchOTCTypeType	OTCType;
	///状态
	TQFMatchOTCStatusType	Status;
	///错误代码
	TQFMatchErrorIDType	OTCErrorID;
	///错误信息
	TQFMatchErrorMsgType	OTCErrorMsg;
	///备注
	TQFMatchComeFromType	Comments;
};
///场外成交对
struct CQFMatchOTCTradeField
{
	///交易日
	TQFMatchDateType	TradingDay;
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///结算编号
	TQFMatchSettlementIDType	SettlementID;
	///成交编号
	TQFMatchTradeIDType	TradeID;
	///交易用户代码
	TQFMatchUserIDType	UserID;
	///成交对本地编号
	TQFMatchOrderLocalIDType	TradeLocalID;
	///合约代码
	TQFMatchInstrumentIDType	InstrumentID;
	///价格
	TQFMatchPriceType	Price;
	///数量
	TQFMatchVolumeType	Volume;
	///买方会员代码
	TQFMatchParticipantIDType	BuyParticipantID;
	///买方客户代码
	TQFMatchClientIDType	BuyClientID;
	///买方交易用户代码
	TQFMatchUserIDType	BuyUserID;
	///买方本地报单编号
	TQFMatchOrderLocalIDType	BuyOrderLocalID;
	///买方组合开平标志
	TQFMatchCombOffsetFlagType	BuyCombOffsetFlag;
	///买方组合投机套保标志
	TQFMatchCombHedgeFlagType	BuyCombHedgeFlag;
	///卖方方会员代码
	TQFMatchParticipantIDType	SellParticipantID;
	///卖方客户代码
	TQFMatchClientIDType	SellClientID;
	///卖方交易用户代码
	TQFMatchUserIDType	SellUserID;
	///卖方组合开平标志
	TQFMatchCombOffsetFlagType	SellCombOffsetFlag;
	///卖方组合套保标志
	TQFMatchCombHedgeFlagType	SellCombHedgeFlag;
	///卖方本地报单编号
	TQFMatchOrderLocalIDType	SellOrderLocalID;
	///OTC类型
	TQFMatchOTCTypeType	OTCType;
	///状态
	TQFMatchOTCStatusType	Status;
	///错误代码
	TQFMatchErrorIDType	OTCErrorID;
	///错误信息
	TQFMatchErrorMsgType	OTCErrorMsg;
	///备注
	TQFMatchComeFromType	Comments;
};
///更新行情
struct CQFMatchUpdateMarketDataField
{
	///合约代码
	TQFMatchInstrumentIDType	InstrumentID;
	///资金费率
	TQFMatchRatioType	FundingRate;
	///合理价格
	TQFMatchPriceType	FairPrice;
	///今结算
	TQFMatchPriceType	SettlementPrice;
};
///仓位保证金操作
struct CQFMatchPositionMarginActionField
{
	///结算组代码
	TQFMatchSettlementGroupIDType	SettlementGroupID;
	///合约代码
	TQFMatchInstrumentIDType	InstrumentID;
	///会员代码
	TQFMatchParticipantIDType	ParticipantID;
	///客户代码
	TQFMatchClientIDType	ClientID;
	///投机套保标志
	TQFMatchHedgeFlagType	HedgeFlag;
	///持仓多空方向
	TQFMatchPosiDirectionType	PosiDirection;
	///值
	TQFMatchRatioType	Value;
	///操作类型
	TQFMatchPositionActionTypeType	ActionType;
	///备注
	TQFMatchComeFromType	Comments;
};

#endif
