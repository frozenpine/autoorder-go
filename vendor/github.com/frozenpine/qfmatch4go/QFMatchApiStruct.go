package qfmatch4go

import "time"

// GoCQFMatchRspInfoField 响应信息
type GoCQFMatchRspInfoField struct {
	ErrorID  int
	ErrorMsg string
}

// GoCQFMatchRspUserLoginField 用户登录应答
type GoCQFMatchRspUserLoginField struct {
	TradingDay        string
	LoginTime         time.Time
	MaxOrderLocalID   int
	UserID            string
	ParticipantID     string
	TradingSystemName string
	DataCenterID      int
	PrivateFlowSize   int
	UserFlowSize      int
}

// GoCQFMatchRspUserLogoutField 用户登出应答
type GoCQFMatchRspUserLogoutField struct {
	UserID        string
	ParticipantID string
}

// GoCQFMatchInputOrderField 报单录入请求
type GoCQFMatchInputOrderField struct {
	OrderSysID          int
	ParticipantID       string
	ClientID            string
	UserID              string
	InstrumentID        string
	OrderPriceType      orderPriceType
	Direction           direction
	CombOffsetFlag      [4]byte
	CombHedgeFlag       [4]byte
	LimitPrice          float64
	VolumeTotalOriginal int
	TimeCondition       timeCondition
	GTDDate             string
	VolumeCondition     volumeCondition
	MinVolume           int
	ContingentCondition contingentCondition
	StopPrice           float32
	ForceCloseReason    forceCloseReason
	OrderLocalID        int
	IsAutoSuspend       bool
	BusinessUnit        string
}

// GoCQFMatchOrderActionField 报单操作请求
type GoCQFMatchOrderActionField struct {
	OrderSysID    int
	OrderLocalID  int
	ActionFlag    actionFlag
	ParticipantID string
	ClientID      string
	UserID        string
	LimitPrice    float64
	VolumeChange  int
	ActionLocalID int
	BusinessUnit  string
}

// GoCQFMatchDisseminationField 查询信息
type GoCQFMatchDisseminationField struct {
	SequenceSeries int
	SequenceNo     int
}

// GoCQFMatchOrderField 委托回报
type GoCQFMatchOrderField struct {
	TradingDay          string
	SettlementGroupID   string
	SettlementID        int
	OrderSysID          int
	ParticipantID       string
	ClientID            string
	UserID              string
	InstrumentID        string
	OrderPriceType      orderPriceType
	Direction           direction
	CombOffsetFlag      [4]offsetFlag
	CombHedgeFlag       [4]hedgeFlag
	LimitPrice          float32
	VolumeTotalOriginal int
	TimeCondition       timeCondition
	GTDDate             string
	VolumeCondition     volumeCondition
	MinVolume           int
	ContingentCondition contingentCondition
	StopPrice           float64
	ForceCloseReason    forceCloseReason
	OrderLocalID        int
	IsAutoSuspend       bool
	OrderSource         orderSource
	OrderStatus         orderStatus
	OrderType           orderType
	VolumeTraded        int
	VolumeTotal         int
	InsertTime          time.Time
	ActiveTime          time.Time
	SuspendTime         time.Time
	UpdateTime          time.Time
	CancelTime          time.Time
	ActiveUserID        string
	Priority            int
	TimeSortID          int
	ClearingPartID      string
	BusinessUnit        string
	CalendarDate        string
}

// GoCQFMatchTradeField 成交回报
type GoCQFMatchTradeField struct {
	TradingDay        string
	SettlementGroupID string
	SettlementID      int
	TradeID           int
	Direction         direction
	OrderSysID        int
	ParticipantID     string
	ClientID          string
	TradingRole       tradingRole
	AccountID         string
	InstrumentID      string
	OffsetFlag        offsetFlag
	HedgeFlag         hedgeFlag
	Price             float64
	Volume            int
	TradeTime         time.Time
	TradeType         tradeType
	PriceSource       priceSource
	UserID            string
	OrderLocalID      int
	ClearingPartID    string
	BusinessUnit      string
	CalendarDate      string
	MatchRole         matchRole
	Fee               float64
	CloseProfit       float64
}

// GoCQFMatchReqUserLoginField 用户登录请求
type GoCQFMatchReqUserLoginField struct {
	TradingDay           string
	UserID               string
	ParticipantID        string
	Password             string
	UserProductInfo      string
	InterfaceProductInfo string
	ProtocolInfo         string
	DataCenterID         int
}

// GoCQFMatchReqUserLogoutField 用户登出请求
type GoCQFMatchReqUserLogoutField struct {
	UserID        string
	ParticipantID string
}

// GoCQFMatchBulletinField 公告
type GoCQFMatchBulletinField struct {
	TradingDay  string
	BulletinID  int
	SequenceNo  int
	NewsType    string
	NewsUrgency byte
	SendTime    time.Time
	Abstract    string
	ComeFrom    string
	Content     string
	URLLink     string
	MarketID    string
}

// GoCQFMatchMarketDataField 市场行情
type GoCQFMatchMarketDataField struct {
	TradingDay         string
	SettlementGroupID  string
	SettlementID       int
	LastPrice          float64
	PreSettlementPrice float64
	PreClosePrice      float64
	PreOpenInterest    float64
	OpenPrice          float64
	HighestPrice       float64
	LowestPrice        float64
	Volume             int
	Turnover           float64
	OpenInterest       float64
	ClosePrice         float64
	SettlementPrice    float64
	UpperLimitPrice    float64
	LowerLimitPrice    float64
	PreDelta           float64
	CurrDelta          float64
	UpdateTime         time.Time
	InstrumentID       string
	FundingRate        float64
	FairPrice          float64
}

// GoCQFMatchRspInstrumentField 合约查询应答
type GoCQFMatchRspInstrumentField struct {
	SettlementGroupID    string
	ProductID            string
	ProductGroupID       string
	UnderlyingInstrID    string
	ProductClass         productClass
	PositionType         positionType
	StrikePrice          float64
	OptionsType          optionsType
	VolumeMultiple       int
	UnderlyingMultiple   float64
	InstrumentID         string
	InstrumentName       string
	DeliveryYear         int
	DeliveryMonth        int
	AdvanceMonth         int
	IsTrading            bool
	CreateDate           string
	OpenDate             string
	ExpireDate           string
	StartDelivDate       string
	EndDelivDate         string
	BasisPrice           float64
	MaxMarketOrderVolume int
	MinMarketOrderVolume int
	MaxLimitOrderVolume  int
	MinLimitOrderVolume  int
	PriceTick            float64
	AllowDelivPersonOpen bool
	Currency             string
}

// GoCQFMatchInstrumentStatusField 合约状态
type GoCQFMatchInstrumentStatusField struct {
	SettlementGroupID string
	InstrumentID      string
	InstrumentStatus  instrumentStatus
	TradingSegmentSN  int
	EnterTime         time.Time
	EnterReason       enterReason
}

// priceItem 价格实体
type priceItem struct {
	Price  float64
	Volume int
}

// GoCQFMatchDepthMarketDataField 深度市场行情
type GoCQFMatchDepthMarketDataField struct {
	TradingDay         string
	SettlementGroupID  string
	SettlementID       int
	LastPrice          float64
	PreSettlementPrice float64
	PreClosePrice      float64
	PreOpenInterest    float64
	OpenPrice          float64
	HighestPrice       float64
	LowestPrice        float64
	Volume             int
	Turnover           float64
	OpenInterest       float64
	ClosePrice         float64
	SettlementPrice    float64
	UpperLimitPrice    float64
	LowerLimitPrice    float64
	PreDelta           float64
	CurrDelta          float64
	UpdateTime         time.Time
	InstrumentID       string
	Bids               []priceItem
	Asks               []priceItem
	FundingRate        float64
	FairPrice          float64
}

// GoCQFMatchQryInstrumentField 合约查询
type GoCQFMatchQryInstrumentField struct {
	SettlementGroupID string
	ProductGroupID    string
	ProductID         string
	InstrumentID      string
}

// GoCQFMatchQryInstrumentStatusField 合约状态查询
type GoCQFMatchQryInstrumentStatusField struct {
	InstIDStart string
	InstIDEnd   string
}

// GoCQFMatchQryMarketDataField 普通行情查询请求
type GoCQFMatchQryMarketDataField struct {
	ProductID    string
	InstrumentID string
}

// GoCQFMatchQryMBLMarketDataField 合约价位查询
type GoCQFMatchQryMBLMarketDataField struct {
	InstIDStart string
	InstIDEnd   string
	Direction   direction
}

// GoCQFMatchMBLMarketDataField 分价表
type GoCQFMatchMBLMarketDataField struct {
	InstrumentID string
	Direction    direction
	priceItem
}
