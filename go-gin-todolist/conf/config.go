package conf

import (
	"gopkg.in/ini.v1"
	"strings"
)

var (
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func LoadIntConfig() string {
	/* ====== Step-1: 加载mysql数据库配置 ====== */

	// 第一种： 固定方式
	// dsn := fmt.Sprintf("root:root@tcp(127.0.0.1:3306)/todolist_db?charset=utf8mb4&parseTime=True&loc=Local")

	// 第二种： 字符串变量拼接
	// dsn := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true&loc=Local"}, "")

	// 第三种：加载文件中的mysql配置
	file, err := ini.Load("conf/config.ini")
	if err != nil {
		panic("Failed to load ini file")
	}

	LoadServer(file)
	LoadMysqlData(file)
	dsn := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true&loc=Local"}, "")

	return dsn
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}
