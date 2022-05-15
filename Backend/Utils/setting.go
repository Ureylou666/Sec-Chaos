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
	JwtKey     string
)

func init() {
	file, err := ini.Load("/Users/ureylou/Downloads/center/Backend/Config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误", err)
	}
	LoadServer(file)
	LoadData(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("Server").Key("Appmode").MustString("debug")
	HttpPort = file.Section("Server").Key("HttpPort").MustString("8080")
	JwtKey = file.Section("Server").Key("JWTkey").String()
}

func LoadData(file *ini.File) {
	Db = file.Section("Database").Key("Db").String()
	DbHost = file.Section("Database").Key("DbHost").String()
	DbPort = file.Section("Database").Key("DbPort").String()
	DbUser = file.Section("Database").Key("DbUser").String()
	DbPassword = file.Section("Database").Key("DbPassword").String()
	DbName = file.Section("Database").Key("DbName").String()
}
