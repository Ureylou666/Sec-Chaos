package main

import (
	"Backend/Model"
	"Backend/Routers"
)

func main() {
	// 引用数据库
	Model.InitDb()
	Routers.InitRouter()
}
