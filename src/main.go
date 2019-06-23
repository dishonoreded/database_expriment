package main

import (
	"./controllers"
	"./models"
	"fmt"
	"net/http"
)

func main() {

	models.Initdb()

	mux := http.NewServeMux()

	//登录
	mux.HandleFunc("/api/login",controllers.Login)
	mux.HandleFunc("/api/logout",controllers.Logout)

	//权限检测
	mux.HandleFunc("/api/check",controllers.Check)

	//管理员
	mux.HandleFunc("/api/manager",controllers.Manager_controllers)

	//教师
	mux.HandleFunc("/api/teacher",controllers.Teacher_controllers)

	//学生
	mux.HandleFunc("/api/student",controllers.Student_controllers)

	fmt.Println("Web服务器启动...端口:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}