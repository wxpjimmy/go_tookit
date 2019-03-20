package demo

import (
	"testing"
)

func TestParseProduct(t *testing.T) {
	m, n := ParseProduct("{\"app_id\":332,\"click_monitor_url\":\"\",\"expose_monitor_url\":\"\",\"legoSiteId\":0,\"app_name\":{\"zh_CN\":\"京东\"},\"package_name\":\"com.jingdong.app.mall\",\"app_download_url\":\"http://app.mi.com/download/332\",\"app_detail_url\":\"http://app.mi.com/detail/332\",\"icon_url\":\"AppStore/09bb3410dc4b444f1bc21dda3c4603b55b5416ccb\",\"rating_score\":2.5,\"level_1_category\":9,\"level_2_category\":160,\"brief\":\"新人送188元购物礼包\",\"singleAppInfo\":{\"type\":1}}")
	if m != 332 {
		t.Error("Parse app_id failed!")
	}

	if n != "京东" {
		t.Error("Parse app_name failed!")
	}
}
