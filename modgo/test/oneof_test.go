package main

import (
	"testing"

	json "github.com/json-iterator/go"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type isGetAbtestRequest_Type interface {
	isGetAbtestRequest_Type()
}

type GetAbtestRequest_GroupId struct {
	GroupId int64 `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3,oneof"`
}

type GetAbtestRequest_DirectionTxt struct {
	DirectionTxt string `protobuf:"bytes,3,opt,name=direction_txt,json=directionTxt,proto3,oneof"`
}

func (*GetAbtestRequest_GroupId) isGetAbtestRequest_Type() {}

func (*GetAbtestRequest_DirectionTxt) isGetAbtestRequest_Type() {}

type GetAbtestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AbtestId       int64  `protobuf:"varint,1,opt,name=abtest_id,json=abtestId,proto3" json:"abtest_id"`
	DirectionTxtv2 string `protobuf:"bytes,4,opt,name=direction_txtv2,json=directionTxtv2,proto3" json:"direction_txtv2"`
	// Types that are assignable to Type:
	//
	//	*GetAbtestRequest_GroupId
	//	*GetAbtestRequest_DirectionTxt
	Type isGetAbtestRequest_Type `protobuf_oneof:"type"`
}

func initbs() {
	bs := &GetAbtestRequest{
		AbtestId: 0,
		Type:     &GetAbtestRequest_DirectionTxt{DirectionTxt: "1"},
	}
	testBs1, _ = json.Marshal(bs)

	bs2 := &GetAbtestRequest{
		AbtestId:       0,
		DirectionTxtv2: "1",
	}
	testBs2, _ = json.Marshal(bs2)

}

var testBs1 []byte
var testBs2 []byte

func BenchmarkOneOf(b *testing.B) {
	initbs()
	for n := 0; n < b.N; n++ {
		data := new(GetAbtestRequest)
		json.Unmarshal(testBs1, data)
		_ = data
	}
}

func BenchmarkString(b *testing.B) {
	initbs()
	for n := 0; n < b.N; n++ {
		data := new(GetAbtestRequest)
		json.Unmarshal(testBs2, data)
		_ = data
	}
}

/*

go test *.go -bench . -benchmem -benchtime 5s  -cpu=1
goos: darwin
goarch: arm64

BenchmarkOneOf  	 6130515	       973.1 ns/op	    1040 B/op	      20 allocs/op
BenchmarkString 	26336937	       228.4 ns/op	     128 B/op	       5 allocs/op
PASS
ok  	command-line-arguments	13.525s


goos 和 goarch：这两行表示Go语言运行时的操作系统和架构。在这个例子中，操作系统（goos）是darwin（即macOS），架构（goarch）是arm64。
278158：表示在5秒内（由-benchtime 5s决定）这个测试被运行了278,158次。测试次数越多，结果越可靠。
18855 ns/op：表示每个操作（每次测试）的平均耗时是18,855纳秒。
2152 B/op：表示每次操作平均分配的内存是2152字节。这是由-benchmem标志启用的。
25 allocs/op：表示每次操作平均进行了25次内存分配。

Benchmark：基准函数的名称。
N：为基准执行的迭代次数。
ns/op：运行基准测试的一次迭代所花费的平均纳秒数。
B/op：在基准测试的每次迭代期间分配的平均字节数。
allocs/op：在基准测试的每次迭代期间分配内存的平均次数。



*/
