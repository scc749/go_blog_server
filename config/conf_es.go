package config

// ES ElasticSearch 配置
type ES struct {
	URL            string `json:"url" yaml:"url"`                           // Elasticsearch 服务的 URL，例如 http://localhost:9200
	Username       string `json:"username" yaml:"username"`                 // 用于连接 Elasticsearch 的用户名
	Password       string `json:"password" yaml:"password"`                 // 用于连接 Elasticsearch 的密码
	IsConsolePrint bool   `json:"is_console_print" yaml:"is_console_print"` // 是否在控制台打印 Elasticsearch 语句，true 表示打印，false 表示不打印
}
