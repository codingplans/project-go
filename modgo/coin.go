package main

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
)

type QuotedPrice struct {
	Name            string
	BuyInRate       string
	SellOutRate     string
	PublishDateTime string
	Symbol          string
}

// var observedList = map[string]string{
// 	"美元":  "USD",
// 	"港币":  "HKD",
// 	"日元":  "JPY",
// 	"比索":  "PHP",
// 	"林吉特": "MYR",
// }

func pings(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "pong")
}

func timeoutMiddleware1(next http.Handler, timeout time.Duration) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), timeout)
		defer cancel()

		r = r.WithContext(ctx)
		done := make(chan struct{})

		go func() {
			defer close(done)
			next.ServeHTTP(w, r)
		}()

		select {
		case <-ctx.Done():
			w.WriteHeader(http.StatusRequestTimeout)
			w.Write([]byte("request timeout"))
		case <-done:
		}
	})
}
func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
		w.Write([]byte("Pong"))
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: timeoutMiddleware1(mux, 2*time.Second),
	}

	log.Println(server.ListenAndServe())

	g := gin.New()
	m := sync.Map{}
	g.GET("/", func(c *gin.Context) {

		// fmt.Println(c.Request.URL.Query().Get("i"), 222222)
		m.Store(c.Request.URL.Query().Get("i"), true)
		time.Sleep(30 * time.Millisecond)
		c.String(200, "ok")
	})

	go func() {
		tt := time.NewTicker(50 * time.Second)
		for {
			select {
			case <-tt.C:
				for i := 0; i < 1000000; i++ {
					if _, ok := m.Load(i); !ok {
						fmt.Println(i)
					}
				}
				m = sync.Map{}
			}
		}

	}()

	g.Run(":8081")

	// aa, _ := crawl("https://www.boc.cn/sourcedb/whpj/enindex_1619.html")
	// fmt.Printf("%+v", aa)
}

func crawl(url string) (map[string]QuotedPrice, error) {
	request := gorequest.New().
		Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36").
		TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req := request.Get(url)
	resp, _, err1 := req.
		Retry(3, 5*time.Second, http.StatusBadRequest, http.StatusInternalServerError).
		End()
	if err1 != nil {
		return nil, errors.New("http request failed")
	}
	doc, err2 := goquery.NewDocumentFromReader(resp.Body)
	if err2 != nil {
		return nil, errors.New("failed to create goquery instance")
	}
	trs := doc.Find("tr[align=\"center\"]")
	result := make(map[string]QuotedPrice, 0)
	trs.Each(func(i int, s *goquery.Selection) {
		tds := s.Find("td")
		if tds.Length() == 0 {
			return
		}
		name := tds.Eq(0).Text()
		sellOutRate := tds.Eq(2).Text()
		buyInRate := tds.Eq(4).Text()
		publishDate := tds.Eq(6).Text()
		quotePrice := QuotedPrice{
			Name:            name,
			BuyInRate:       buyInRate,
			SellOutRate:     sellOutRate,
			PublishDateTime: publishDate,
		}
		fmt.Printf("今日汇率 %v \n", quotePrice)
		result[quotePrice.Name] = quotePrice
	})
	return result, nil
}
