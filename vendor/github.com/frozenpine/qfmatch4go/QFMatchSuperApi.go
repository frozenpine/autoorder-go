package qfmatch4go

/*
#cgo CFLAGS: -I${SRCDIR} -fPIC
#cgo LDFLAGS: -L${SRCDIR} -lcQFMatchSuperApi

#include "cgoQFMatchSuperApi.h"
*/
import "C"

// QFMatchSuperAPICallbacks Quantfair撮合引擎超级API回调接口
type QFMatchSuperAPICallbacks interface {
	commonAPICallbacks
	marketAPICallbacks
	tradeAPICallbacks
	adminAPICallbacks
}

// QFMatchSuperAPI Quantfair撮合引擎超级API
type QFMatchSuperAPI struct {
	commonAPI
	marketAPI
	traderAPI
	adminAPI
}

// RegisterCallbacks 注册回调接口
func (api *QFMatchSuperAPI) RegisterCallbacks(cb interface{}) {
	vtCallbacks := C.Callbacks{}

	switch cb.(type) {
	case QFMatchMarketAPICallbacks:
		api.registerMarketCallbacks(&vtCallbacks)
	case QFMatchTraderAPICallbacks:
		api.registerMarketCallbacks(&vtCallbacks)
		api.registerTraderCallbacks(&vtCallbacks)
	case QFMatchSuperAPICallbacks:
		api.registerMarketCallbacks(&vtCallbacks)
		api.registerTraderCallbacks(&vtCallbacks)
		api.registerAdminCallbacks(&vtCallbacks)
	default:
		panic("Invalid callback interface.")
	}

	api.registerCommonCallbacks(&vtCallbacks)

	C.RegisterCallbacks(C.QFMatchSuperApiInstance(api.apiInstance), &vtCallbacks)

	apiMap[api.apiInstance] = cb
}
