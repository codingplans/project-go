package main

import (
	"testgo/modgo/gen/pkg"
)

type Clients struct {
	name string
}

func NewClient() (*Clients, *Clients2) {
	return &Clients{}, &Clients2{}
}
func NewClient2() *Clients {
	return &Clients{}
}

type Clients2 struct {
}

func NewA() *Clients2 {
	return &Clients2{}
}

func NewAdxHttpServer(
	flowSvc *FlowService,
	partnerSvc *PartnerService,
	mediaSvc *MediaService,
	dspService *DspService,
	priceService *PriceService,
	bigReportSvc *BigDataReportService,
	adxSvc *AdxService,
	splashDspSvc *SplashDspService,
	failReportSvc *FailReportService,
	contractSvc *ContractPdbService,
	wordPacketSvc *WordPacketGroupService) *pkg.Pkgs {
	return &pkg.Pkgs{}
}

type FlowService struct {
}
type PartnerService struct {
}
type MediaService struct {
}
type DspService struct {
}
type PriceService struct {
}
type BigDataReportService struct {
}
type AdxService struct {
}
type SplashDspService struct {
}
type FailReportService struct {
}
type ContractPdbService struct {
}
type WordPacketGroupService struct {
}
