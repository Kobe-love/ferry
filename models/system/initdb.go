package system

import (
	"ferry/global/orm"
	"fmt"
	"os"
	"strings"
)

/*
  @Author : lanyulei
*/

func InitDb() error {
	filePath := "config/db.sql"
	sql, err := Ioutil(filePath)
	if err != nil {
		fmt.Println("数据库基础数据初始化脚本读取失败！原因:", err.Error())
		return err
	}
	sqlList := strings.Split(sql, ";")
	for _, sql := range sqlList {
		if strings.Contains(sql, "--") {
			fmt.Println(sql)
			continue
		}
		sqlValue := strings.Replace(sql+";", "\n", "", 1)
		if err = orm.Eloquent.Exec(sqlValue).Error; err != nil {
			if !strings.Contains(err.Error(), "Query was empty") {
				return err
			}
		}
	}
	return nil
}

func Ioutil(name string) (string, error) {
	if contents, err := os.ReadFile(name); err == nil {
		//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
		result := strings.Replace(string(contents), "\n", "", 1)
		return result, nil
	} else {
		return "", err
	}
}
