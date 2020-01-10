package main

import (
	"context"
	"encoding/json"
	"gopkg.in/olivere/elastic.v5"
	"iceberg/frame/icelog"
	"laoyuegou.pb/godgame/model"
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
	// searchService := es.UpdateByQuery().Index("god_game_qa").Type("peiwan_stats")

	// query := elastic.NewTermQuery("god_id", 1992252)
	searchService := es.Search().Index("god_game_qa").Type("peiwan_stats")

	query := elastic.NewBoolQuery()

	k := elastic.NewGeoDistanceSort("location2").Point(float64(31.1713), float64(121.412)).
		Order(true).
		Unit("km").
		SortMode("max").
		GeoDistance("plane")
	icelog.Info(k)

	q := elastic.NewGeoDistanceQuery("location2").
		GeoPoint(elastic.GeoPointFromLatLon(float64(31.1713), float64(121.412))).
		Distance("100km")
	query = query.Filter(q)

	searchService = searchService.Query(query).From(int(0)).
		Size(int(50)).
		// Sort("weight", false).
		// Sort("_score", false).
		// Sort("lts", false).
		SortBy(k)

		// Sort("location2", true).
		// Pretty(true)
	ee, _ := k.Source()
	bs, _ := json.Marshal(ee)
	icelog.Infof("%s", string(bs))

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
	icelog.Infof("%+v,&&&&&& %+v")

	// fmt.Printf("query cost %d millisecond.\n", resp.TookInMillis)
	// icelog.Infof("%+v,&&&&&& %+v", resp)
	// if resp.Hits.TotalHits == 0 {
	// 	return nil
	// }
	if resp != nil {

		// var into godgamepb.QueryQuickOrderResp_Data
		// err := json.Unmarshal(*resp.Hits.Hits[0].Source, &into)
		// icelog.Infof("%+v @@@@@@^^^^^^^@", into)
		// icelog.Info(err)

		for _, item := range resp.Hits.Hits {
			if seq := strings.Split(item.Id, "-"); len(seq) == 2 {
				// icelog.Info(item.Id, "###", resp.Hits.TotalHits, "\n")
				var into model.ESGodGameRedefine
				json.Unmarshal(*item.Source, &into)
				ee := len(item.Sort)
				icelog.Infof("%+v %% %+v   ^^^   %+v    @@@@@@@", item.Id, item.Sort[ee-1])

				// if bs, err := item.Source.MarshalJSON(); err == nil {
				// 	icelog.Infof("%s", string(bs))
				//
				// 	// json.Unmarshal(bs)
				//
				// }
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
