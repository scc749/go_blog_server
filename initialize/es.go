package initialize

import (
	"os"
	"server/global"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"go.uber.org/zap"
)

// ConnectEs 初始化并返回一个配置好的 Elasticsearch 客户端
func ConnectEs() *elasticsearch.TypedClient {
	esCfg := global.Config.ES
	cfg := elasticsearch.Config{
		Addresses: []string{esCfg.URL},
		Username:  esCfg.Username,
		Password:  esCfg.Password,
	}

	// 如果配置中指定了需要打印日志到控制台，则启用日志打印
	if esCfg.IsConsolePrint {
		cfg.Logger = &elastictransport.ColorLogger{
			Output:             os.Stdout, // 设置日志输出到标准输出（控制台）
			EnableRequestBody:  true,      // 启用请求体打印
			EnableResponseBody: true,      // 启用响应体打印
		}
	}

	// 创建一个新的 Elasticsearch 客户端
	client, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		global.Log.Error("Failed to connect to Elasticsearch", zap.Error(err))
		os.Exit(1)
	}

	return client
}
