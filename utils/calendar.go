package utils

import (
	"errors"
	"io"
	"net/http"
	"regexp"
	"server/model/other"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

var solarTerm = []string{
	"立春", "雨水", "惊蛰", "春分", "清明", "谷雨",
	"立夏", "小满", "芒种", "夏至", "小暑", "大暑",
	"立秋", "处暑", "白露", "秋分", "寒露", "霜降",
	"立冬", "小雪", "大雪", "冬至", "小寒", "大寒",
}

func GetCalendar(dateStr string) (other.Calendar, error) {
	resp, err := http.Get("https://www.rili.com.cn/rili/json/today/" + dateStr + ".js")
	if err != nil {
		return other.Calendar{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return other.Calendar{}, err
	}

	var jsonStr string
	reg := regexp.MustCompile(`(?s)\((.*?)\);`)
	result := reg.FindAllStringSubmatch(string(body), -1)
	if len(result) > 0 && len(result[0]) > 1 {
		jsonStr = result[0][1]
	} else {
		return other.Calendar{}, errors.New("failed to get data")
	}

	jieqi := gjson.Get(jsonStr, "jieqi.jieqi").Str

	calendar := other.Calendar{
		Date:         gjson.Get(jsonStr, "yangli.date").Str + " " + gjson.Get(jsonStr, "yangli.xingqi").Str,
		LunarDate:    gjson.Get(jsonStr, "nongli.yueri").Str,
		Ganzhi:       gjson.Get(jsonStr, "nongli.ganzhi").Str,
		Zodiac:       gjson.Get(jsonStr, "xingzuo.xingzuo").Str + "座",
		DayOfYear:    "今年第" + strconv.FormatInt(gjson.Get(jsonStr, "nian_index").Int(), 10) + "天",
		SolarTerm:    jieqi + "第" + strconv.FormatInt(gjson.Get(jsonStr, "jieqi.jieqi_index").Int(), 10) + "天　距离" + nextSolarTerm(jieqi) + "还有" + strconv.FormatInt(gjson.Get(jsonStr, "jieqi.jieqi_next").Int(), 10) + "天",
		Auspicious:   strings.ReplaceAll(gjson.Get(jsonStr, "yi").Str, ",", " "),
		Inauspicious: strings.ReplaceAll(gjson.Get(jsonStr, "ji").Str, ",", " "),
	}

	return calendar, nil
}

func nextSolarTerm(currentTerm string) string {
	for i, term := range solarTerm {
		if term == currentTerm {
			if i == len(solarTerm)-1 {
				return solarTerm[0]
			}
			return solarTerm[i+1]
		}
	}
	return ""
}
