package main

import (
	"fmt"
	"github.com/prometheus/common/log"
)


func main() {

	aa := fmt.Sprintf("%g", 123313323.312312)
	aa1 := fmt.Sprintf("%f", 123313323.312312)
	aa2 := fmt.Sprintf("%v", 123313323.312312)


	log.Infof("%+v ", aa)
	log.Infof("%+v ", aa1)
	log.Infof("%+v ", aa2)
}
