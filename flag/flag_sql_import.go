package flag

import (
	"os"
	"server/global"
	"strings"
)

// SQLImport 导入 MySQL 数据
func SQLImport(sqlPath string) (errs []error) {
	byteData, err := os.ReadFile(sqlPath)
	if err != nil {
		return append(errs, err)
	}
	// 分割数据
	sqlList := strings.Split(string(byteData), ";")
	for _, sql := range sqlList {
		// 去除字符串开头和结尾的空白符
		sql = strings.TrimSpace(sql)
		if sql == "" {
			continue
		}
		// 执行sql语句
		err = global.DB.Exec(sql).Error
		if err != nil {
			errs = append(errs, err)
			continue
		}
	}
	return nil
}
