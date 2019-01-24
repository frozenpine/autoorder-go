package qfmatch4go

// ResumeType 流水恢复模式
type ResumeType uint8

const (
	// RESTART 从头开始
	RESTART ResumeType = iota
	// RESUME 从断点开始
	RESUME
	// QUICK 继续
	QUICK
)

type orderPriceType uint8

const (
	_ orderPriceType = iota
	// OPTAnyPrice 任意价
	OPTAnyPrice
	// OPTLimitPrice 限价
	OPTLimitPrice
	// OPTBestPrice 最优价
	OPTBestPrice
)

type direction uint8

const (
	// DBuy 买方向
	DBuy direction = iota
	// DSell 卖方向
	DSell
)

func (d direction) GetName() string {
	switch d {
	case DBuy:
		return "买"
	case DSell:
		return "卖"
	default:
		panic("invalid direction value")
	}
}

func (d direction) Opposite() direction {
	switch d {
	case DBuy:
		return DSell
	case DSell:
		return DBuy
	default:
		panic("invalid direction value")
	}
}

type timeCondition uint8

const (
	_ timeCondition = iota
	// TCIOC 立即完成，否则撤销
	TCIOC
	// TCGFS 本节有效
	TCGFS
	// TCGFD 当日有效
	TCGFD
	// TCGTD 指定日期前有效
	TCGTD
	// TCGTC 撤销前有效
	TCGTC
	// TCGFA 集合竞价有效
	TCGFA
)

type volumeCondition uint8

const (
	_ volumeCondition = iota
	// VCAV 任何数量
	VCAV
	// VCMV 最小数量
	VCMV
	// VCCV 全部数量
	VCCV
)

type contingentCondition uint8

const (
	_ contingentCondition = iota
	// CCImmediately 立即
	CCImmediately
	// CCTouch 止损
	CCTouch
)

type forceCloseReason uint8

const (
	// FCRNotForceClose 非强平
	FCRNotForceClose forceCloseReason = iota
	// FCRLackDeposit 资金不足
	FCRLackDeposit
	// FCRClientOverPositionLimit 客户超仓
	FCRClientOverPositionLimit
	// FCRMemberOverPositionLimit 会员超仓
	FCRMemberOverPositionLimit
	// FCRNotMultiple 持仓非整数倍
	FCRNotMultiple
	// FCRViolation 违规
	FCRViolation
	// FCROther 其它
	FCROther
	// FCRPersonDeliv 自然人临近交割
	FCRPersonDeliv
)

type actionFlag uint8

const (
	// AFDelete 删除
	AFDelete actionFlag = iota
	// AFSuspend 挂起
	AFSuspend
	// AFActive 激活
	AFActive
	// AFModify 修改
	AFModify
)

type orderSource uint8

const (
	// OSParticipant 来自参与者
	OSParticipant orderSource = iota
	// OSAdministrator 来自管理员
	OSAdministrator
)

type orderStatus uint8

const (
	// OSAllTraded 全部成交
	OSAllTraded orderStatus = iota
	// OSPartTradedQueueing 部分成交还在队列中
	OSPartTradedQueueing
	// OSPartTradedNotQueueing 部分成交不在队列中
	OSPartTradedNotQueueing
	// OSNoTradeQueueing 未成交还在队列中
	OSNoTradeQueueing
	// OSNoTradeNotQueueing 未成交不在队列中
	OSNoTradeNotQueueing
	// OSCanceled 撤单
	OSCanceled
)

type orderType uint8

const (
	// OTNormal 正常
	OTNormal orderType = iota
	// OTDeriveFromQuote 报价衍生
	OTDeriveFromQuote
	// OTDeriveFromCombination 组合衍生
	OTDeriveFromCombination
	// OTDeriveFromOTCTrade 场外成交衍生
	OTDeriveFromOTCTrade
)

type tradingRole uint8

const (
	_ tradingRole = iota
	// TRBroker 代理
	TRBroker
	// TRHost 自营
	TRHost
	// TRMarketMaker 做市商
	TRMarketMaker
)

type offsetFlag uint8

const (
	// OFOpen 开仓
	OFOpen offsetFlag = iota
	// OFClose 平仓
	OFClose
	// OFForceClose 强平
	OFForceClose
	// OFCloseToday 平今
	OFCloseToday
	// OFCloseYesterday 平昨
	OFCloseYesterday
)

type hedgeFlag uint8

const (
	_ hedgeFlag = iota
	// HFSpeculation 投机
	HFSpeculation
	// HFArbitrage 套利
	HFArbitrage
	// HFHedge 套保
	HFHedge
	// HFMarketMaker 做市商
	HFMarketMaker
)

type tradeType uint8

const (
	// TPCommon 普通成交
	TPCommon tradeType = iota
	// TPOptionsExecution 期权执行
	TPOptionsExecution
	// TPOTC OTC成交
	TPOTC
	// TPEFPDerived 期转现衍生成交
	TPEFPDerived
	// TPCombinationDerived 组合衍生成交
	TPCombinationDerived
	// TPOTCTrade 大宗交易成交
	TPOTCTrade
	// TPEFPTrade 期转现成交
	TPEFPTrade
	// TPForceCloseTrade 强平成交
	TPForceCloseTrade
)

type priceSource uint8

const (
	// PSLastPrice 前成交价
	PSLastPrice priceSource = iota
	// PSBuy 买委托价
	PSBuy
	// PSSell 卖委托价
	PSSell
)

type matchRole uint8

const (
	// MRNone 不区分
	MRNone matchRole = iota
	// MRTaker 主动成交方
	MRTaker
	// MRMaker 被动成交方
	MRMaker
)

type productClass uint8

const (
	_ productClass = iota
	// PCFuture 期货
	PCFuture
	// PCOptions 期权
	PCOptions
	// PCCombination 组合
	PCCombination
	// PCSpot 即期
	PCSpot
	// PCEFP 期转现
	PCEFP
	// PCSwap 掉期
	PCSwap
)

type positionType uint8

const (
	_ positionType = iota
	// PTNet 净持仓
	PTNet
	// PTGross 综合持仓
	PTGross
)

type optionsType uint8

const (
	// OTNotOptions 非期权
	OTNotOptions optionsType = iota
	// OTCallOptions 看涨
	OTCallOptions
	// OTPutOptions 看跌
	OTPutOptions
)

func (opt optionsType) GetName() string {
	switch opt {
	case OTNotOptions:
		return "非期权"
	case OTCallOptions:
		return "看涨"
	case OTPutOptions:
		return "看跌"
	default:
		return "未知"
	}
}

type instrumentStatus uint8

const (
	// ISBeforeTrading 开盘前
	ISBeforeTrading instrumentStatus = iota
	// ISNoTrading 非交易
	ISNoTrading
	// ISContinous 连续交易
	ISContinous
	// ISAuctionOrdering 集合竞价报单
	ISAuctionOrdering
	// ISAuctionBalance 集合竞价价格平衡
	ISAuctionBalance
	// ISAuctionMatch 集合竞价撮合
	ISAuctionMatch
	// ISClosed 收盘
	ISClosed
)

func (status instrumentStatus) GetName() string {
	switch status {
	case ISBeforeTrading:
		return "开盘前"
	case ISNoTrading:
		return "非交易"
	case ISContinous:
		return "连续交易"
	case ISAuctionOrdering:
		return "集合竞价报单"
	case ISAuctionBalance:
		return "集合竞价平衡"
	case ISAuctionMatch:
		return "集合竞价撮合"
	case ISClosed:
		return "收盘"
	default:
		return "未知"
	}
}

type enterReason uint8

const (
	_ enterReason = iota
	// IERAutomatic 自动切换
	IERAutomatic
	// IERManual 手动切换
	IERManual
	// IERFuse 熔断
	IERFuse
	// IERFuseManual 熔断手动
	IERFuseManual
)

func (reason enterReason) GetName() string {
	switch reason {
	case IERAutomatic:
		return "自动切换"
	case IERManual:
		return "手动切换"
	case IERFuse:
		return "熔断"
	case IERFuseManual:
		return "手动熔断"
	default:
		return "未知"
	}
}
