package druid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//top n only support one dimension one metric
var top_n_query_str = `{
  "queryType" : "topN",
  "dataSource" : "dsp_fee_stat",
  "intervals" : ["2019-03-18/2019-03-19"],
  "threshold": 20,
  "granularity" : "all",
"filter": {
    "type": "and",
    "fields": [
      {
        "type": "regex",
        "dimension": "dsp",
        "pattern": "ccc\\..*"
      },
      {
        "type": "regex",
        "dimension": "tagid",
        "pattern": "^(?!1\\.24\\.).*"
      }
    ]
  },
  "dimension" : "tagid",
  "metric" : "fee",
 "aggregations": [
    {
      "type": "longSum",
      "name": "cost",
      "fieldName": "cost"
    },
    {
      "type": "doubleSum",
      "name": "fee",
      "fieldName": "fee"
    }
  ],
  "postAggregations": [
    {
      "type": "arithmetic",
      "name": "cost_fee_ratio",
      "fn": "/",
      "fields": [
        {
          "type": "fieldAccess",
          "name": "cost",
          "fieldName": "cost"
        },
        {
          "type": "fieldAccess",
          "name": "fee",
          "fieldName": "fee"
        }
      ]
    }
  ]
}`

var group_by_query_str = `{
  "queryType": "groupBy",
  "dataSource": "dsp_fee_stat",
  "granularity": "day",
  "dimensions": ["tagid", "dsp"],
  "filter": {
    "type": "and",
    "fields": [
      {
        "type": "regex",
        "dimension": "dsp",
        "pattern": "ccc\\.(?!schedule).*"
      },
      {
        "type": "regex",
        "dimension": "tagid",
        "pattern": "^(?!1\\.24\\.).*"
      }
    ]
  },
  "aggregations": [
    {
      "type": "longSum",
      "name": "cost",
      "fieldName": "cost"
    },
    {
      "type": "doubleSum",
      "name": "fee",
      "fieldName": "fee"
    }
  ],
  "postAggregations": [
    {
      "type": "arithmetic",
      "name": "cost_fee_ratio",
      "fn": "/",
      "fields": [
        {
          "type": "fieldAccess",
          "name": "cost",
          "fieldName": "cost"
        },
        {
          "type": "fieldAccess",
          "name": "fee",
          "fieldName": "fee"
        }
      ]
    }
   ],
  "intervals": [ "2019-07-09T00:00:00.000/2019-07-10T00:00:00.000" ]
}`

var url = "http://broker.druid.data.srv/druid/v2"

type DspTagIdFee struct {
	tagId	string	`json:"tag_id"`
	dsp   	string
	cost 	float64
	fee  	float64
	costFeeRatio	float64 `json:"cost_fee_ratio"`
}

type Wrapper struct {
	version   string
	timestamp string
	event     *DspTagIdFee
}

type TagIdDetail struct {
	totalCost	float64
	totalFee	float64
	perDspData	map[string]*DspTagIdFee
}

func GetDspFeeStat() {
	content :=[]byte(group_by_query_str)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(content))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	var data []*Wrapper
	json.Unmarshal(body, &data)
	fmt.Println(len(data))

	var tagIdToDetail = make(map[string]*TagIdDetail)

	for _, d := range data {
		id := d.event.tagId
		var res *TagIdDetail
		res, exist := tagIdToDetail[id]
		if !exist {
			res = &TagIdDetail{0,0,make(map[string]*DspTagIdFee)}
			tagIdToDetail[id] = res
		}
		res.totalCost += d.event.cost
		res.totalFee += d.event.fee
		res.perDspData[d.event.dsp] = d.event
	}

}

