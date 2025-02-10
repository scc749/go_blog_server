package config

// Redis 缓存数据库配置
type Redis struct {
	Address  string `json:"address" yaml:"address"`   // Redis 服务器的地址，通常为 "localhost:6379" 或其他主机和端口
	Password string `json:"password" yaml:"password"` // 连接 Redis 时的密码，如果没有设置密码则留空
	DB       int    `json:"db" yaml:"db"`             // 指定使用的数据库索引，单实例模式下可选择的数据库，默认为 0
}
