package hotSearch

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"server/model/other"
	"strconv"
	"time"

	"github.com/tidwall/gjson"
)

type Zhihu struct {
}

func (*Zhihu) GetHotSearchData(maxNum int) (other.HotSearchData, error) {
	resp, err := http.Get("https://www.zhihu.com/billboard")
	if err != nil {
		return other.HotSearchData{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return other.HotSearchData{}, err
	}

	var jsonStr string
	reg := regexp.MustCompile(`(?s)<script id="js-initialData" type="text/json">(.*?)</script>`)
	result := reg.FindAllStringSubmatch(string(body), -1)
	if len(result) > 0 && len(result[0]) > 1 {
		jsonStr = result[0][1]
	} else {
		return other.HotSearchData{}, errors.New("failed to get data")
	}

	updateTime := time.Now().Format("2006-01-02 15:04:05")

	var hotList []other.HotItem
	for i := 0; i < maxNum; i++ {
		if index := gjson.Get(jsonStr, "initialState.topstory.hotList."+strconv.Itoa(i)+".id"); !index.Exists() {
			break
		}
		hotList = append(hotList, other.HotItem{
			Index:       i + 1,
			Title:       gjson.Get(jsonStr, "initialState.topstory.hotList."+strconv.Itoa(i)+".target.titleArea.text").Str,
			Description: gjson.Get(jsonStr, "initialState.topstory.hotList."+strconv.Itoa(i)+".target.excerptArea.text").Str,
			Image:       fmt.Sprintf(gjson.Get(jsonStr, "initialState.topstory.hotList."+strconv.Itoa(i)+".target.imageArea.url").Str),
			Popularity:  gjson.Get(jsonStr, "initialState.topstory.hotList."+strconv.Itoa(i)+".target.metricsArea.text").Str,
			URL:         fmt.Sprintf(gjson.Get(jsonStr, "initialState.topstory.hotList."+strconv.Itoa(i)+".target.link.url").Str),
		})
	}

	return other.HotSearchData{Source: "知乎热榜", UpdateTime: updateTime, HotList: hotList}, nil
}
