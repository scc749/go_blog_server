package other

type Calendar struct {
	Date         string `json:"date"`         // 日期
	LunarDate    string `json:"lunar_date"`   // 农历
	Ganzhi       string `json:"ganzhi"`       // 干支
	Zodiac       string `json:"zodiac"`       // 星座
	DayOfYear    string `json:"day_of_year"`  // 天次
	SolarTerm    string `json:"solar_term"`   // 节气
	Auspicious   string `json:"auspicious"`   // 宜
	Inauspicious string `json:"inauspicious"` // 忌
}
