package flag

import (
	"fmt"
	"os"
	"os/exec"
	"server/global"
	"time"
)

// SQLExport 导出 MySQL 数据
func SQLExport() error {
	mysql := global.Config.Mysql

	timer := time.Now().Format("20060102")
	sqlPath := fmt.Sprintf("mysql_%s.sql", timer)
	cmd := exec.Command("docker", "exec", "mysql", "mysqldump", "-u"+mysql.Username, "-p"+mysql.Password, mysql.DBName)

	outFile, err := os.Create(sqlPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	cmd.Stdout = outFile
	return cmd.Run()
}
