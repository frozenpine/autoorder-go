package qfmatch4go

/*
#cgo CFLAGS: -I${SRCDIR} -fPIC
#cgo LDFLAGS: -L${SRCDIR} -lcQFMatchSuperApi

#include <string.h>

#include "cgoCommonApi.h"
*/
import "C"
import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/djimenez/iconv-go"
)

var enumChars = [10]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

// ToCChar Convert user defined enum(uint8) to C.char
func ToCChar(enum interface{}) C.char {
	enumValue := reflect.ValueOf(enum).Uint()

	return C.char(enumChars[enumValue])
}

var apiMap map[int]interface{}

type commonAPICallbacks interface {
	OnFrontConnected()
	OnFrontDisconnected(reason int)
	OnHeartBeatWarning(time int)
	OnPackageStart(topicID, sequenceNO int)
	OnPackageEnd(topicID, sequenceNO int)
	OnRspError(err *GoCQFMatchRspInfoField, requestID int, isLast bool)
	OnRspUserLogin(rspUserLogin *GoCQFMatchRspUserLoginField, err *GoCQFMatchRspInfoField, requestID int, isLast bool)
	OnRspUserLogout(rspUserLogout *GoCQFMatchRspUserLogoutField, err *GoCQFMatchRspInfoField, requestID int, isLast bool)
	OnRspQryBulletin(bulletin *GoCQFMatchBulletinField, err *GoCQFMatchRspInfoField, requestID int, isLast bool)
	OnRtnBulletin(bulletin *GoCQFMatchBulletinField)
}

func getInstanceCallback(instance C.QFMatchSuperApiInstance) interface{} {
	callback, exist := apiMap[int(instance)]

	if !exist {
		return nil
	}

	return callback
}

func convertRspInfo(pRspInfo *C.struct_CQFMatchRspInfoField) *GoCQFMatchRspInfoField {
	rsp := GoCQFMatchRspInfoField{}

	if pRspInfo != nil {
		rsp.ErrorID = int(pRspInfo.ErrorID)
		msg, _ := iconv.ConvertString(C.GoString(&pRspInfo.ErrorMsg[0]), "gbk", "utf-8")
		rsp.ErrorMsg = msg
	}

	return &rsp
}

func convertRspUserLogin(pRspUserLogin *C.struct_CQFMatchRspUserLoginField) *GoCQFMatchRspUserLoginField {
	login := GoCQFMatchRspUserLoginField{}

	login.TradingDay = C.GoString(&pRspUserLogin.TradingDay[0])

	loginTime, _ := time.Parse("15:04:05", C.GoString(&pRspUserLogin.LoginTime[0]))
	login.LoginTime = loginTime

	maxLocalID, _ := strconv.Atoi(C.GoString(&pRspUserLogin.MaxOrderLocalID[0]))
	login.MaxOrderLocalID = maxLocalID

	login.UserID = C.GoString(&pRspUserLogin.UserID[0])
	login.ParticipantID = C.GoString(&pRspUserLogin.ParticipantID[0])
	login.TradingSystemName = C.GoString(&pRspUserLogin.TradingSystemName[0])
	login.DataCenterID = int(pRspUserLogin.DataCenterID)
	login.PrivateFlowSize = int(pRspUserLogin.PrivateFlowSize)
	login.UserFlowSize = int(pRspUserLogin.UserFlowSize)

	return &login
}

func convertRspUserLogout(pRspUserLogout *C.struct_CQFMatchRspUserLogoutField) *GoCQFMatchRspUserLogoutField {
	logout := GoCQFMatchRspUserLogoutField{}

	logout.UserID = C.GoString(&pRspUserLogout.UserID[0])
	logout.ParticipantID = C.GoString(&pRspUserLogout.ParticipantID[0])

	return &logout
}

func convertBulletinField(pBulletin *C.struct_CQFMatchBulletinField) *GoCQFMatchBulletinField {
	bulletin := GoCQFMatchBulletinField{}

	bulletin.TradingDay = C.GoString(&pBulletin.TradingDay[0])
	bulletin.BulletinID = int(pBulletin.BulletinID)
	bulletin.SequenceNo = int(pBulletin.SequenceNo)
	bulletin.NewsType = C.GoString(&pBulletin.NewsType[0])
	bulletin.NewsUrgency = byte(pBulletin.NewsUrgency)

	sendTime, _ := time.Parse("20060102 15:04:05", fmt.Sprintf("%s %s", C.GoString(&pBulletin.SendTime[0]), C.GoString(&pBulletin.CalendarDate[0])))
	bulletin.SendTime = sendTime

	bulletin.Abstract = C.GoString(&pBulletin.Abstract[0])
	bulletin.ComeFrom = C.GoString(&pBulletin.ComeFrom[0])
	bulletin.Content = C.GoString(&pBulletin.Content[0])
	bulletin.URLLink = C.GoString(&pBulletin.URLLink[0])
	bulletin.MarketID = C.GoString(&pBulletin.MarketID[0])

	return &bulletin
}

//export goOnFrontConnected
func goOnFrontConnected(instance C.QFMatchSuperApiInstance) {
	callback := getInstanceCallback(instance)

	if callback != nil {
		callback.(commonAPICallbacks).OnFrontConnected()
	}
}

//export goOnFrontDisconnected
func goOnFrontDisconnected(instance C.QFMatchSuperApiInstance, nReason C.int) {
	callback := getInstanceCallback(instance)

	if callback != nil {
		callback.(commonAPICallbacks).OnFrontDisconnected(int(nReason))
	}
}

//export goOnHeartBeatWarning
func goOnHeartBeatWarning(instance C.QFMatchSuperApiInstance, nTime C.int) {
	callback := getInstanceCallback(instance)

	if callback != nil {
		callback.(commonAPICallbacks).OnHeartBeatWarning(int(nTime))
	}
}

//export goOnPackageStart
func goOnPackageStart(instance C.QFMatchSuperApiInstance, nTopicID, nSequenceNO C.int) {
	callback := getInstanceCallback(instance)

	if callback != nil {
		callback.(commonAPICallbacks).OnPackageStart(int(nTopicID), int(nSequenceNO))
	}
}

//export goOnPackageEnd
func goOnPackageEnd(instance C.QFMatchSuperApiInstance, nTopicID, nSequenceNO C.int) {
	callback := getInstanceCallback(instance)

	if callback != nil {
		callback.(commonAPICallbacks).OnPackageEnd(int(nTopicID), int(nSequenceNO))
	}
}

//export goOnRspError
func goOnRspError(instance C.QFMatchSuperApiInstance, pRspInfo *C.struct_CQFMatchRspInfoField, nRequestID C.int, bIsLast C.bool) {
	if pRspInfo == nil {
		return
	}

	callback := getInstanceCallback(instance)

	if callback != nil {
		callback.(commonAPICallbacks).OnRspError(convertRspInfo(pRspInfo), int(nRequestID), bool(bIsLast))
	}
}

//export goOnRspUserLogin
func goOnRspUserLogin(instance C.QFMatchSuperApiInstance, pRspUserLogin *C.struct_CQFMatchRspUserLoginField, pRspInfo *C.struct_CQFMatchRspInfoField, nRequestID C.int, bIsLast C.bool) {
	if pRspUserLogin == nil {
		return
	}

	callback := getInstanceCallback(instance)

	if callback != nil {
		callback.(commonAPICallbacks).OnRspUserLogin(convertRspUserLogin(pRspUserLogin), convertRspInfo(pRspInfo), int(nRequestID), bool(bIsLast))
	}
}

//export goOnRspUserLogout
func goOnRspUserLogout(instance C.QFMatchSuperApiInstance, pRspUserLogout *C.struct_CQFMatchRspUserLogoutField, pRspInfo *C.struct_CQFMatchRspInfoField, nRequestID C.int, bIsLast C.bool) {
	if pRspUserLogout == nil {
		return
	}

	callback := getInstanceCallback(instance)

	if callback != nil {
		callback.(commonAPICallbacks).OnRspUserLogout(convertRspUserLogout(pRspUserLogout), convertRspInfo(pRspInfo), int(nRequestID), bool(bIsLast))
	}
}

//export goOnRspQryBulletin
func goOnRspQryBulletin(instance C.QFMatchSuperApiInstance, pBulletin *C.struct_CQFMatchBulletinField, pRspInfo *C.struct_CQFMatchRspInfoField, nRequestID C.int, bIsLast C.bool) {
	if pBulletin == nil {
		return
	}

	callback := getInstanceCallback(instance)

	if callback != nil {
		callback.(commonAPICallbacks).OnRspQryBulletin(convertBulletinField(pBulletin), convertRspInfo(pRspInfo), int(nRequestID), bool(bIsLast))
	}
}

//export goOnRtnBulletin
func goOnRtnBulletin(instance C.QFMatchSuperApiInstance, pBulletin *C.struct_CQFMatchBulletinField) {
	if pBulletin == nil {
		return
	}

	callback := getInstanceCallback(instance)

	if callback != nil {
		callback.(commonAPICallbacks).OnRtnBulletin(convertBulletinField(pBulletin))
	}
}

type commonAPI struct {
	apiInstance int
	FrontAddr   string
	FensAddr    string
	Connected   bool
	Logged      bool
	UserInfo    *GoCQFMatchReqUserLoginField
	reqID       int64
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// InitAPI API初始化函数
func (api *commonAPI) InitAPI(flowPath, superAPIType string) {
	exist, _ := pathExists(flowPath)
	if !exist {
		os.Mkdir(flowPath, os.ModePerm)
	}

	api.apiInstance = int(C.InitApi(C.CString(flowPath), C.CString(superAPIType)))
}

func (api *commonAPI) registerCommonCallbacks(ptrVtCallbacks *C.Callbacks) {
	if apiMap == nil {
		apiMap = make(map[int]interface{})
	}

	ptrVtCallbacks.ptrOnFrontConnected = C.FuncOnFrontConnected(C.cgoOnFrontConnected)
	ptrVtCallbacks.ptrOnFrontDisconnected = C.FuncOnFrontDisconnected(C.cgoOnFrontDisconnected)
	ptrVtCallbacks.ptrOnHeartBeatWarning = C.FuncOnHeartBeatWarning(C.cgoOnHeartBeatWarning)
	ptrVtCallbacks.ptrOnPackageStart = C.FuncOnPackageStart(C.cgoOnPackageStart)
	ptrVtCallbacks.ptrOnPackageEnd = C.FuncOnPackageEnd(C.cgoOnPackageEnd)
	ptrVtCallbacks.ptrOnRspError = C.FuncOnRspError(C.cgoOnRspError)
	ptrVtCallbacks.ptrOnRspUserLogin = C.FuncOnRspUserLogin(C.cgoOnRspUserLogin)
	ptrVtCallbacks.ptrOnRspUserLogout = C.FuncOnRspUserLogout(C.cgoOnRspUserLogout)
	ptrVtCallbacks.ptrOnRspQryBulletin = C.FuncOnRspQryBulletin(C.cgoOnRspQryBulletin)
	ptrVtCallbacks.ptrOnRtnBulletin = C.FuncOnRtnBulletin(C.cgoOnRtnBulletin)
}

// GetVersion 获取API版本
func (api *commonAPI) GetVersion() string {
	return C.GoString(C.GetVersion(C.QFMatchSuperApiInstance(api.apiInstance)))
}

// Release 删除接口对象
func (api *commonAPI) Release() {
	C.Release(C.QFMatchSuperApiInstance(api.apiInstance))
	C.Join(C.QFMatchSuperApiInstance(api.apiInstance))
}

// Init 初始化，调用后接口才开始工作
func (api *commonAPI) Init() {
	C.Init(C.QFMatchSuperApiInstance(api.apiInstance))
}

// GetTradingDay 获取当前交易日
func (api *commonAPI) GetTradingDay() string {
	return C.GoString(C.GetTradingDay(C.QFMatchSuperApiInstance(api.apiInstance)))
}

// RegisterFront 注册前置地址
func (api *commonAPI) RegisterFront(frontAddr string) {
	C.RegisterFront(C.QFMatchSuperApiInstance(api.apiInstance), C.CString(frontAddr))

	api.FrontAddr = frontAddr
}

// RegisterNameServer 注册名字服务器网络地址
func (api *commonAPI) RegisterNameServer(nsAddr string) {
	C.RegisterNameServer(C.QFMatchSuperApiInstance(api.apiInstance), C.CString(nsAddr))

	api.FensAddr = nsAddr
}

// RegisterCertificateFile 加载证书
func (api *commonAPI) RegisterCertificateFile(certFileName, keyFileName, caFileName, keyFilePassword string) int {
	return int(C.RegisterCertificateFile(C.QFMatchSuperApiInstance(api.apiInstance), C.CString(certFileName), C.CString(keyFileName), C.CString(caFileName), C.CString(keyFilePassword)))
}

// SubscribePrivateTopic 订阅私有流
func (api *commonAPI) SubscribePrivateTopic(resumeType ResumeType) {
	C.SubscribePublicTopic(C.QFMatchSuperApiInstance(api.apiInstance), C.enum_QFMATCH_TE_RESUME_TYPE(resumeType))
}

// SubscribePublicTopic 订阅公共流
func (api *commonAPI) SubscribePublicTopic(resumeType ResumeType) {
	C.SubscribePublicTopic(C.QFMatchSuperApiInstance(api.apiInstance), C.enum_QFMATCH_TE_RESUME_TYPE(resumeType))
}

// SubscribeUserTopic 订阅交易员流
func (api *commonAPI) SubscribeUserTopic(resumeType ResumeType) {
	C.SubscribeUserTopic(C.QFMatchSuperApiInstance(api.apiInstance), C.enum_QFMATCH_TE_RESUME_TYPE(resumeType))
}

// SetHeartbeatTimeout 设置心跳超时时间
func (api *commonAPI) SetHeartbeatTimeout(timeout uint32) {
	C.SetHeartbeatTimeout(C.QFMatchSuperApiInstance(api.apiInstance), C.uint(timeout))
}

// OpenRequestLog 打开请求日志文件
func (api *commonAPI) OpenRequestLog(reqLogFileName string) int {
	return int(C.OpenRequestLog(C.QFMatchSuperApiInstance(api.apiInstance), C.CString(reqLogFileName)))
}

// OpenResponseLog 打开应答日志文件
func (api *commonAPI) OpenResponseLog(rspLogFileName string) int {
	return int(C.OpenResponseLog(C.QFMatchSuperApiInstance(api.apiInstance), C.CString(rspLogFileName)))
}

func (api *commonAPI) getRequestID() int64 {
	return atomic.AddInt64(&api.reqID, 1)
}

// ReqUserLogin 用户登录请求
func (api *commonAPI) ReqUserLogin(login *GoCQFMatchReqUserLoginField) int {
	api.UserInfo = login

	loginReq := C.struct_CQFMatchReqUserLoginField{}

	cUserID := C.CString(login.UserID)
	C.strncpy(&loginReq.UserID[0], cUserID, C.sizeof_TQFMatchUserIDType-1)

	cParticipantID := C.CString(login.ParticipantID)
	C.strncpy(&loginReq.ParticipantID[0], cParticipantID, C.sizeof_TQFMatchParticipantIDType-1)

	cPassword := C.CString(login.Password)
	C.strncpy(&loginReq.Password[0], cPassword, C.sizeof_TQFMatchPasswordType-1)

	return int(C.ReqUserLogin(C.QFMatchSuperApiInstance(api.apiInstance), &loginReq, C.int(api.getRequestID())))
}

// ReqUserLogout 用户退出请求
func (api *commonAPI) ReqUserLogout(logout *GoCQFMatchReqUserLogoutField) int {
	logoutReq := C.struct_CQFMatchReqUserLogoutField{}

	cUserID := C.CString(logout.UserID)
	C.strncpy(&logoutReq.UserID[0], cUserID, C.sizeof_TQFMatchUserIDType-1)

	cParticipantID := C.CString(logout.ParticipantID)
	C.strncpy(&logoutReq.ParticipantID[0], cParticipantID, C.sizeof_TQFMatchParticipantIDType-1)

	return int(C.ReqUserLogout(C.QFMatchSuperApiInstance(api.apiInstance), &logoutReq, C.int(api.getRequestID())))
}

// OnFrontConnected 前置连接消息
func (api *commonAPI) OnFrontConnected() {
	api.Connected = true

	log.Printf("Front[%s] connected.\n", api.FrontAddr)
}

// OnFrontDisconnected 前置断连消息
func (api *commonAPI) OnFrontDisconnected(reason int) {
	api.Connected = false

	log.Printf("Front[%s] disconnected\n", api.FrontAddr)
}

// OnHeartBeatWarning 心跳失败消息
func (api *commonAPI) OnHeartBeatWarning(time int) {
	log.Printf("Missing heartbeat in %d seconds\n", time)
}

// OnPackageStart 系列数据包开始消息
func (api *commonAPI) OnPackageStart(topicID, sequenceNO int) {}

// OnPackageEnd 系列数据包结束消息
func (api *commonAPI) OnPackageEnd(topicID, sequenceNO int) {}

// OnRspError 错误应答消息
func (api *commonAPI) OnRspError(err *GoCQFMatchRspInfoField, requestID int, isLast bool) {}

// OnRspUserLogin 用户登录消息
func (api *commonAPI) OnRspUserLogin(rspUserLogin *GoCQFMatchRspUserLoginField, err *GoCQFMatchRspInfoField, requestID int, isLast bool) {
	if err != nil && err.ErrorID != 0 {
		log.Printf("User[%s] login failed: [%d]%s\n", rspUserLogin.UserID, err.ErrorID, err.ErrorMsg)
	} else {
		api.Logged = true
		log.Printf("User[%s] login success\n", rspUserLogin.UserID)
	}
}

// OnRspUserLogout 用户登出消息
func (api *commonAPI) OnRspUserLogout(rspUserLogout *GoCQFMatchRspUserLogoutField, err *GoCQFMatchRspInfoField, requestID int, isLast bool) {
	if err != nil && err.ErrorID != 0 {
		log.Printf("User[%s] logout failed: [%d]%s\n", rspUserLogout.UserID, err.ErrorID, err.ErrorMsg)
	} else {
		log.Printf("User[%s] logout success\n", rspUserLogout.UserID)
	}
}

// OnRspQryBulletin 通告消息查询应答
func (api *commonAPI) OnRspQryBulletin(bulletin *GoCQFMatchBulletinField, err *GoCQFMatchRspInfoField, requestID int, isLast bool) {
}

// OnRtnBulletin 通告消息推送
func (api *commonAPI) OnRtnBulletin(bulletin *GoCQFMatchBulletinField) {}
