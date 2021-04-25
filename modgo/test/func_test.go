package main

import (
	"fmt"
	"testing"
)

type PayWay struct {
	//    支付id
	Id  int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Ids int64 `protobuf:"varint,2,opt,name=id,proto3" json:"ids,omitempty"`
	// 支付名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func TestSliceRange(t *testing.T) {
	t.Helper()
	aa := []*PayWay{}

	aa = append(aa, &PayWay{
		Id:  123,
		Ids: 123,
	})
	aa = append(aa, &PayWay{
		Id:  222,
		Ids: 222,
	})
	aa = append(aa, &PayWay{
		Id:  333,
		Ids: 333,
	})

	for k, v := range aa {
		fmt.Println(v, k)
	}

}
