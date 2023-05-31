package main

// 仅在linux下有效 默认注释掉这段
// func TestMemInfo(t *testing.T) {
// 	p, err := process.NewProcess(int32(os.Getpid()))
// 	if err != nil {
// 		t.Error(err.Error())
// 		return
// 	}
// 	cpuPercent, err := p.Percent(time.Second)
// 	if err != nil {
// 		t.Error(err.Error())
// 		return
// 	}
//
// 	mem, err := p.MemoryInfo()
// 	if err != nil {
// 		t.Error(err.Error())
// 		return
// 	}
//
// 	rss := mem.RSS
// 	gNum := runtime.NumGoroutine()
// 	tNum := getThreadNum()
// 	t.Log(cpuPercent, rss, gNum, tNum)
// }
//
// func getThreadNum() int {
// 	return pprof.Lookup("threadcreate").Count()
// }
