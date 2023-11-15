package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	// 读取请求的数据
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	time.Sleep(time.Millisecond * 1000)

	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	// 处理请求数据，这里只是简单地打印出来
	fmt.Println("Received POST data:", string(body), time.Now().UnixNano())

	// 返回响应
	w.WriteHeader(http.StatusOK)
	// w.Write([]byte("POST request received successfully\n"))

	bidResponse := &BidResponse{
		RequestId:        "sad22222222",
		StatusCode:       122,
		Reason:           122,
		Desc:             "sad22222222",
		ProcessingTimeMs: 12,
		Ads:              []BidResponse_Ad_MaterialMeta{},
	}
	bs, _ := json.Marshal(bidResponse)
	write, err := w.Write(bs)
	fmt.Println(err)
	fmt.Println(write)

}

func main() {

	m := http.NewServeMux()
	m.Handle("/debug", http.HandlerFunc(postHandler))

	fmt.Println("starting the server:")
	// 启动HTTP服务器，监听8080端口
	if err := http.ListenAndServe(":8081", m); err != nil {
		fmt.Println("Error starting the server:", err)
	}
}

type BidResponse struct {
	RequestId        string                        `protobuf:"bytes,1,req,name=request_id,json=requestId" json:"request_id,omitempty"`
	StatusCode       int64                         `protobuf:"varint,2,req,name=status_code,json=statusCode" json:"status_code,omitempty"`
	Reason           int32                         `protobuf:"varint,3,req,name=reason" json:"reason,omitempty"`
	Desc             string                        `protobuf:"bytes,4,opt,name=desc" json:"desc,omitempty"`
	ProcessingTimeMs uint32                        `protobuf:"varint,5,opt,name=processing_time_ms,json=processingTimeMs" json:"processing_time_ms,omitempty"`
	Ads              []BidResponse_Ad_MaterialMeta `protobuf:"bytes,6,rep,name=ads" json:"ads,omitempty"`
}

type BidResponse_Ad_MaterialMeta struct {
	CreativeId           string   `protobuf:"bytes,1,req,name=creative_id,json=creativeId" json:"creative_id,omitempty"`
	TargetUrl            string   `protobuf:"bytes,7,opt,name=target_url,json=targetUrl" json:"target_url,omitempty"`
	DownloadUrl          string   `protobuf:"bytes,8,opt,name=download_url,json=downloadUrl" json:"download_url,omitempty"`
	Title                string   `protobuf:"bytes,9,opt,name=title" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,10,opt,name=description" json:"description,omitempty"`
	AppName              string   `protobuf:"bytes,11,opt,name=app_name,json=appName" json:"app_name,omitempty"`
	PackageName          string   `protobuf:"bytes,12,opt,name=package_name,json=packageName" json:"package_name,omitempty"`
	WinNoticeUrl         []string `protobuf:"bytes,13,rep,name=win_notice_url,json=winNoticeUrl" json:"win_notice_url,omitempty"`
	Source               string   `protobuf:"bytes,14,opt,name=source" json:"source,omitempty"`
	Icon                 string   `protobuf:"bytes,15,opt,name=icon" json:"icon,omitempty"`
	Adm                  string   `protobuf:"bytes,16,req,name=adm" json:"adm,omitempty"`
	LossNoticeUrl        []string `protobuf:"bytes,17,rep,name=loss_notice_url,json=lossNoticeUrl" json:"loss_notice_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
