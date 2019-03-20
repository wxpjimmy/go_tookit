package demo

import (
	"encoding/json"
)

func ParseProduct(data string) (int, string) {
	var dat map[string]interface{}

	var bytes = []byte(data)

	if err := json.Unmarshal(bytes, &dat); err != nil {
		panic(err)
	}

	appid := dat["app_id"].(float64)
	appname := dat["app_name"].(map[string]interface{})

	return int(appid), appname["zh_CN"].(string)
}
