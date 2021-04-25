package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	loopApp()
	// return
	ticker := time.NewTicker(time.Hour * 1)
	times := 0
	for {
		select {
		case <-ticker.C:
			if times >= 10 {
				continue
			}
			loopApp()
			times++
		}
	}

}
func loopApp() {

	url := "https://"
	// 翻倍
	url2 := "weather.reikigames.com/a/caishen_weather/award_video"

	urls := []string{
		// "pgd-beta.cootekservice.com/a/caishen_weather/award_countdown",
		// "weather.reikigames.com/a/caishen_weather/award_countdown",
		// "weather.reikigames.com/a/caishen_weather/award_countdown",
		// "weather.reikigames.com/a/caishen_weather/award_countdown",
		// "weather.reikigames.com/a/caishen_weather/award_countdown",
		"weather.reikigames.com/a/caishen_weather/award_countdown",
		"weather.reikigames.com/a/caishen_weather/award_countdown",
		"weather.reikigames.com/a/caishen_weather/award_countdown",
		"weather.reikigames.com/a/caishen_weather/award_countdown",
		"weather.reikigames.com/a/caishen_weather/award_countdown",
		"weather.reikigames.com/a/caishen_weather/award_offline",
		"weather.reikigames.com/a/caishen_weather/award_time",
		"weather.reikigames.com/a/caishen_weather/award_temp",
		"weather.reikigames.com/a/caishen_weather/sign_do",
		"weather.reikigames.com/a/caishen_weather/task_do",
		"weather.reikigames.com/a/caishen_weather/task_do",
		"weather.reikigames.com/a/caishen_weather/task_do",
		"weather.reikigames.com/a/caishen_weather/task_do",
		"weather.reikigames.com/a/caishen_weather/task_do",
		"weather.reikigames.com/a/caishen_weather/task_do",
		"weather.reikigames.com/a/caishen_weather/award_watch",
	}

	headers := []string{
		// 	`{"id":1}`,
		// 	`{"id":2}`,
		// 	`{"id":3}`,
		// 	`{"id":4}`,
		// 	`{"id":5}`,
		`{"id":6}`,
		`{"id":7}`,
		`{"id":8}`,
		`{"id":9}`,
		`{"id":10}"`,
		"",
		"",
		"",
		"",
		`{"task_id":"AWARD_COIN"}`,
		`{"task_id":"FIRST_WITHDRAWAL"}`,
		`{"task_id":"AWARD_WATCH"}`,
		`{"task_id":"AWARD_SIGN"}`,
		`{"task_id":"WEATHER_INFO"}`,
		`{"task_id":"AWARD_APP"}`,
		"",
		"",
		"",
		"",
	}

	// 1 温差 2 倒计时 3 任务 4签到 5新人红包
	headers2 := []string{
		// "6",
		// "6",
		// "6",
		// "6",
		// "6",
		"6",
		"6",
		"6",
		"6",
		"6",
		"2",
		"2",
		"1",
		"4",
		"3",
		"3",
		"3",
		"3",
		"3",
		"3",
		"4",
		"4",
		"",
		"",
		"",
	}

	for k, v := range urls {
		getCoin(url+v, headers[k])
		// 翻倍
		if headers2[k] != "" {
			getCoin(url+url2, `{
    "multiple":3,
    "award_type":`+headers2[k]+`}`)
		}

	}

}

func getCoin(url, header string) {

	method := "POST"

	payload := strings.NewReader(header)
	// payload := strings.NewReader(`{"id":1}`)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err, 111)
		return
	}
	// req.Header.Add("auth-token", "a82da074-73e6-48d7-b187-f9655dd62a81")
	req.Header.Add("auth-token", "cn01:3cb7b95f-e1b5-479a-b87a-b82252c6f203")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err, 222)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err, 333)
		return
	}
	fmt.Println(string(body))
}
