package utils

import (
	"context"
	"server/global"
	"server/model/other"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// MySQLPagination 实现 MySQL 数据分页查询
func MySQLPagination[T any](model *T, option other.MySQLOption) (list []T, total int64, err error) {
	// 设置分页的默认值
	if option.Page < 1 {
		option.Page = 1 // 页码不能小于1，默认为1
	}
	if option.PageSize < 1 {
		option.PageSize = 10 // 每页记录数不能小于1，默认为10
	}
	if option.Order == "" {
		option.Order = "id desc" // 默认按id降序排列
	}

	// 创建查询
	query := global.DB.Model(model)

	// 如果传入了额外的 WHERE 条件，则应用这些条件
	if option.Where != nil {
		query = query.Where(option.Where)
	}

	// 计算符合条件的记录总数
	if err = query.Count(&total).Error; err != nil {
		return nil, 0, err // 如果查询总数失败，返回错误
	}

	// 预加载关联模型
	for _, preload := range option.Preload {
		query = query.Preload(preload) // 应用预加载的关联查询
	}

	// 应用分页查询
	err = query.Order(option.Order).
		Limit(option.PageSize).                      // 设置每页记录数
		Offset((option.Page - 1) * option.PageSize). // 设置偏移量，根据页码计算
		Find(&list).Error                            // 执行查询，并将结果存入 list 中

	return list, total, err // 返回分页结果和总记录数
}

// EsPagination 实现 Elasticsearch 数据分页查询
func EsPagination(ctx context.Context, option other.EsOption) (list []types.Hit, total int64, err error) {
	// 设置分页的默认值
	if option.Page < 1 {
		option.Page = 1 // 页码不能小于1，默认为1
	}
	if option.PageSize < 1 {
		option.PageSize = 10 // 每页记录数不能小于1，默认为10
	}

	// 设置 Elasticsearch 查询的分页值
	from := (option.Page - 1) * option.PageSize // 计算从哪一条记录开始
	option.Request.Size = &option.PageSize      // 设置每页的记录数
	option.Request.From = &from                 // 设置起始记录位置

	// 执行 Elasticsearch 搜索查询
	res, err := global.ESClient.Search().
		Index(option.Index).                       // 指定索引
		Request(option.Request).                   // 应用查询请求
		SourceIncludes_(option.SourceIncludes...). // 设置需要包含的字段
		Do(ctx)                                    // 执行查询
	if err != nil {
		return nil, 0, err // 如果查询失败，返回错误
	}

	// 提取查询结果
	list = res.Hits.Hits         // 获取查询结果中的文档
	total = res.Hits.Total.Value // 获取符合条件的文档总数
	return list, total, nil      // 返回查询结果和总文档数
}
