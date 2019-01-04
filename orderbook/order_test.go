package orderbook

import (
	"encoding/json"
	"testing"
)

func TestMarshalJson(t *testing.T) {
	ord := newOrder(100, 1, nil)

	data, err := json.Marshal(ord)

	if err != nil {
		t.Error(err)
	} else {
		t.Log(string(data))
	}
}
