package main

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
	"iceberg/frame/icelog"
	godgamepb "laoyuegou.pb/godgame/pb"
	"strings"
)

type EsCfg struct {
	Hosts    []string `json:"hosts"`
	UserName string   `json:"username"`
	Password string   `json:"password"`
}

func main() {
	// TestRedis()
	testEs()
}

func testEs() []*elastic.SearchHit {
	es := connEs()
	searchService := es.Search().Index("quick_order_qa").Type("peiwan_stats")
	query := elastic.NewBoolQuery()

	query = query.Must(elastic.NewTermQuery("god_id", 1896))

	query = query.Must(elastic.NewTermQuery("game_id", 15))

	resp, err := searchService.Query(query).
		// From(0).
		Size(300).
		Sort("update_time", false). // 倒序
		// Pretty(true).
		Do(context.Background())

	if err != nil {
		icelog.Debug(err.Error())
		return nil
	}
	// fmt.Printf("query cost %d millisecond.\n", resp.TookInMillis)
	icelog.Infof("%+v,&&&&&& %+v", resp)
	if resp.Hits.TotalHits == 0 {
		return nil
	}
	if resp != nil {

		var into godgamepb.QueryQuickOrderResp_Data
		err := json.Unmarshal(*resp.Hits.Hits[0].Source, &into)
		icelog.Infof("%+v @@@@@@^^^^^^^@", into)
		icelog.Info(err)

		for _, item := range resp.Hits.Hits {
			if seq := strings.Split(item.Id, "-"); len(seq) == 2 {
				// icelog.Info(item.Id, "###", resp.Hits.TotalHits)
				// var into godgamepb.QueryQuickOrderResp_Data
				var into godgamepb.QueryQuickOrderResp_Data
				json.Unmarshal(*item.Source, &into)
				icelog.Infof("%+v @@@@@@@", into)

				if bs, err := item.Source.MarshalJSON(); err == nil {
					icelog.Infof("%s", string(bs))

					// json.Unmarshal(bs)

				}
				// data := item.Source.UnmarshalJSON(into)
			}
		}
		return resp.Hits.Hits

	}
	return nil
}

// func connEs() *CacheStore {
func connEs() *elastic.Client {

	ES := &EsCfg{
		Hosts:    []string{"http://es-cn-mp90zsqsu000m4d08.elasticsearch.aliyuncs.com:9200"},
		UserName: "elastic",
		Password: "TEeOVJHwRY1U",
	}

	// fmt.Printf("%+v", cfg)
	esOptions := []elastic.ClientOptionFunc{
		elastic.SetURL(ES.Hosts...),
		elastic.SetSniff(false),
		elastic.SetMaxRetries(10),
	}
	if ES.UserName != "" && ES.Password != "" {
		esOptions = append(esOptions, elastic.SetBasicAuth(ES.UserName, ES.Password))
	}
	esClient, err := elastic.NewClient(esOptions...)
	if err != nil {
		icelog.Errorf("Init esClient error:%s", err)
	} else {
		// es = esClient
	}
	return esClient
}
