package flag

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"server/global"
	"server/model/elasticsearch"
	"server/model/other"
	"time"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// ElasticsearchExport 导出 ES 中的数据到 JSON 文件
func ElasticsearchExport() error {
	// 声明变量用于存储响应结果
	var response other.ESIndexResponse

	// 发起第一次搜索请求，设置查询条件：索引名称、滚动时间（1分钟）、每次返回的文档数量、查询条件（匹配所有文档）
	res, err := global.ESClient.Search().
		Index(elasticsearch.ArticleIndex()).                   // 设置查询的索引名称
		Scroll("1m").                                          // 设置滚动时间为 1 分钟
		Size(1000).                                            // 每次查询返回 1000 条数据
		Query(&types.Query{MatchAll: &types.MatchAllQuery{}}). // 查询条件：匹配所有文档
		Do(context.TODO())                                     // 执行请求，传入空的 context
	if err != nil {
		return err
	}

	// 遍历第一次查询结果的文档
	for _, hit := range res.Hits.Hits {
		// 为每个文档创建一个 Data 结构体，并将其 ID 和 Source（文档内容）存储
		data := other.Data{
			ID:  hit.Id_,
			Doc: hit.Source_,
		}
		// 将数据追加到 response 的 Data 字段中
		response.Data = append(response.Data, data)
	}

	// 使用 Scroll API 进行后续的滚动查询，直到没有更多数据
	for {
		// 发起新的滚动查询，传入上一个查询返回的 ScrollId
		res, err := global.ESClient.Scroll().ScrollId(*res.ScrollId_).Scroll("1m").Do(context.TODO())
		if err != nil {
			// 如果请求出错，返回错误
			return err
		}

		// 如果没有更多数据，结束滚动查询
		if len(res.Hits.Hits) == 0 {
			break
		}

		// 遍历滚动查询结果中的文档
		for _, hit := range res.Hits.Hits {
			// 为每个文档创建一个 Data 结构体，并将其 ID 和 Source 存储
			data := other.Data{
				ID:  hit.Id_,
				Doc: hit.Source_,
			}
			// 将数据追加到 response 的 Data 字段中
			response.Data = append(response.Data, data)
		}
	}

	// 清除滚动查询，释放 Elasticsearch 上的资源
	_, err = global.ESClient.ClearScroll().ScrollId(*res.ScrollId_).Do(context.TODO())
	if err != nil {
		return err
	}

	// 生成文件名，格式为 "es_yyyyMMdd.json"
	fileName := fmt.Sprintf("es_%s.json", time.Now().Format("20060102"))

	// 创建文件
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将 response 数据结构转换为 JSON 格式的字节数据
	byteData, err := json.Marshal(response)
	if err != nil {
		return err
	}

	// 将 JSON 数据写入文件
	_, err = file.Write(byteData)
	if err != nil {
		return err
	}

	return nil
}
