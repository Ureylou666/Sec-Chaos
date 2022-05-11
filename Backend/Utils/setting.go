package Utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

func init() {
	file, err := ini.Load("/Config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误", err)
	}
	LoadServer(file)
	LoadData(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("Server").Key("Appmode").MustString("debug")
	HttpPort = file.Section("Server").Key("HttpPort").MustString("8080")
}

func LoadData(file *ini.File) {
	Db = file.Section("Database").Key("Db").MustString("mysql")
	DbHost = file.Section("Database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("Database").Key("DbPort").MustString("3306")
	DbUser = file.Section("Database").Key("DbUser").MustString("SecAdmin")
	DbPassword = file.Section("Database").Key("DbPassword").MustString("Nike12#$")
	DbName = file.Section("Database").Key("DbName").MustString("Sec_Chaos")
}
