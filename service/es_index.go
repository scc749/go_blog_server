package service

import (
	"context"
	"server/global"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// EsService 提供了对 Elasticsearch 索引的操作方法
type EsService struct {
}

// IndexCreate 创建一个新的 Elasticsearch 索引，带有指定的映射
func (esService *EsService) IndexCreate(indexName string, mapping *types.TypeMapping) error {
	_, err := global.ESClient.Indices.Create(indexName).Mappings(mapping).Do(context.TODO())
	return err
}

// IndexDelete 删除指定的 Elasticsearch 索引
func (esService *EsService) IndexDelete(indexName string) error {
	_, err := global.ESClient.Indices.Delete(indexName).Do(context.TODO())
	return err
}

// IndexExists 检查指定的 Elasticsearch 索引是否存在
func (esService *EsService) IndexExists(indexName string) (bool, error) {
	return global.ESClient.Indices.Exists(indexName).Do(context.TODO())
}
