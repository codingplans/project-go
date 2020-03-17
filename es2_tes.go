package main

import (
	"context"
	"encoding/json"
	"gopkg.in/olivere/elastic.v5"
	"iceberg/frame/icelog"
	"laoyuegou.pb/godgame/model"
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
	// searchService := es.UpdateByQuery().Index("god_game_qa").Type("peiwan_stats")

	// query := elastic.NewTermQuery("god_id", 1992252)
	searchService := es.Search().Index("god_game_qa").Type("peiwan_stats")

	query := elastic.NewBoolQuery()

	// query = query.Must(elastic.NewTermsQuery("god_level", []int{1, 3}))

	arr := []int64{7, 9, 11, 13, 15, 17, 19, 149}
	aa := make([]interface{}, 0)
	for _, v := range arr {
		aa = append(aa, v)
	}

	icelog.Info(aa)
	query = query.Must(elastic.NewTermsQuery("highest_level_id", aa...))

	// ss, _ := json.Marshal(query)
	// icelog.Infof("%s", string(ss))

	searchService = searchService.Query(query)

	resp, err := searchService.
		Size(50).
		// Sort("highest_level_id_score", true).
		// Sort("emails", true).
		Do(context.Background())

	// 检查是否正常返回
	if err != nil {
		icelog.Debug(err.Error())
		return nil
	}
	// 打印结果
	if resp != nil {
		var pwObj model.ESGodGameRedefine
		icelog.Infof("%+v ****** %+v   ", len(resp.Hits.Hits))

		for _, item := range resp.Hits.Hits {
			if err = json.Unmarshal(*item.Source, &pwObj); err != nil {
				icelog.Infof("%+v %######", pwObj)
			}

			icelog.Infof("%+v ****** %+v   ^^^   %+v    @@@@@@@", pwObj.GodID, pwObj.HighestLevelID, pwObj.HighestLevelIdScore)

			// icelog.Infof("%+v %######", pwObj.)

		}

		// aa, _ := json.Marshal(resp.Hits.Hits)
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
