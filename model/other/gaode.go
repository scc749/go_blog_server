package other

// IPResponse 用于表示 IP 定位查询的响应结果
type IPResponse struct {
	Status    string `json:"status"`    // 返回结果状态值：0表示失败，1表示成功
	Info      string `json:"info"`      // 返回状态说明
	InfoCode  string `json:"infocode"`  // 状态码：10000代表正确
	Province  string `json:"province"`  // 省份名称
	City      string `json:"city"`      // 城市名称
	Adcode    string `json:"adcode"`    // 城市的 adcode 编码
	Rectangle string `json:"rectangle"` // 所在城市矩形区域范围
}

// Cast 表示天气预报中的每日数据
type Cast struct {
	Date         string `json:"date"`         // 日期
	Week         string `json:"week"`         // 星期几
	DayWeather   string `json:"dayweather"`   // 白天天气现象
	NightWeather string `json:"nightweather"` // 晚上天气现象
	DayTemp      string `json:"daytemp"`      // 白天温度
	NightTemp    string `json:"nighttemp"`    // 晚上温度
	DayWind      string `json:"daywind"`      // 白天风向
	NightWind    string `json:"nightwind"`    // 晚上风向
	DayPower     string `json:"daypower"`     // 白天风力
	NightPower   string `json:"nightpower"`   // 晚上风力
}

// Live 表示实况天气数据
type Live struct {
	Province         string `json:"province"`          // 省份名
	City             string `json:"city"`              // 城市名
	Adcode           string `json:"adcode"`            // 区域编码
	Weather          string `json:"weather"`           // 天气现象（汉字描述）
	Temperature      string `json:"temperature"`       // 实时气温，单位：摄氏度
	WindDirection    string `json:"winddirection"`     // 风向描述
	WindPower        string `json:"windpower"`         // 风力级别，单位：级
	Humidity         string `json:"humidity"`          // 空气湿度
	ReportTime       string `json:"reporttime"`        // 数据发布的时间
	TemperatureFloat string `json:"temperature_float"` // 浮点型气温
	HumidityFloat    string `json:"humidity_float"`    // 浮点型湿度
}

// Forecast 表示天气预报信息
type Forecast struct {
	City       string `json:"city"`       // 城市名称
	Adcode     string `json:"adcode"`     // 城市编码
	Province   string `json:"province"`   // 省份名称
	ReportTime string `json:"reporttime"` // 预报发布时间
	Casts      []Cast `json:"casts"`      // 预报数据 list
}

// WeatherResponse 用于表示天气查询的响应结果
type WeatherResponse struct {
	Status   string   `json:"status"`   // 返回状态
	Count    string   `json:"count"`    // 返回结果总数目
	Info     string   `json:"info"`     // 返回的状态信息
	InfoCode string   `json:"infocode"` // 返回状态说明, 10000代表正确
	Lives    []Live   `json:"lives"`    // 实况天气数据信息
	Forecast Forecast `json:"forecast"` // 预报天气信息数据
}
