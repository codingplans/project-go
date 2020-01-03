package main

import (
	"context"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"iceberg/frame/icelog"
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
	searchService := es.UpdateByQuery().Index("god_game_qa").Type("peiwan_stats")
	query := elastic.NewTermQuery("god_id", 1992252)

	// searchService = searchService.Query(query).Script(elastic.NewScriptInline(fmt.Sprintf("ctx._source.%s=%v", "gender", 33)))

	searchService = searchService.Query(query).Script(elastic.NewScriptInline(fmt.Sprintf("ctx._source.%s=%v", "location2['lon']",
		float64(121.408171)))).Script(elastic.NewScriptInline(fmt.Sprintf("ctx._source.%s=%v", "location2['lat']", float64(31.171419))))

	// searchService = searchService.Query(query).XSource()
	// st, _ := elastic.NewScriptInline(fmt.Sprintf("ctx._source.%s=%v", "location2.lat", 111)).Source()
	// icelog.Infof("%s", elastic.NewScriptInline(fmt.Sprintf("ctx._source.%s=%v", "gender", 12)))
	// ss, _ := json.Marshal(st)
	// icelog.Infof("%s", string(ss))
	resp, err := searchService.
		Do(context.Background())

	if err != nil {
		icelog.Debug(err.Error())
		return nil
	}
	// fmt.Printf("query cost %d millisecond.\n", resp.TookInMillis)
	icelog.Infof("%+v,&&&&&& %+v", resp)
	// if resp.Hits.TotalHits == 0 {
	// 	return nil
	// }
	// if resp != nil {
	//
	// 	var into godgamepb.QueryQuickOrderResp_Data
	// 	err := json.Unmarshal(*resp.Hits.Hits[0].Source, &into)
	// 	icelog.Infof("%+v @@@@@@^^^^^^^@", into)
	// 	icelog.Info(err)
	//
	// 	for _, item := range resp.Hits.Hits {
	// 		if seq := strings.Split(item.Id, "-"); len(seq) == 2 {
	// 			// icelog.Info(item.Id, "###", resp.Hits.TotalHits)
	// 			// var into godgamepb.QueryQuickOrderResp_Data
	// 			var into godgamepb.QueryQuickOrderResp_Data
	// 			json.Unmarshal(*item.Source, &into)
	// 			icelog.Infof("%+v @@@@@@@", into)
	//
	// 			if bs, err := item.Source.MarshalJSON(); err == nil {
	// 				icelog.Infof("%s", string(bs))
	//
	// 				// json.Unmarshal(bs)
	//
	// 			}
	// 			// data := item.Source.UnmarshalJSON(into)
	// 		}
	// 	}
	// 	return resp.Hits.Hits
	//
	// }
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
