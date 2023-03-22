package main

import (
	"net/http"
	//导入echo包
	"github.com/labstack/echo/v4"
)

func main() {
	//实例化echo对象。
	e := echo.New()

	//注册一个Get请求, 路由地址为: /tizi365  并且绑定一个控制器函数, 这里使用的是闭包函数。
	e.GET("/tizi365", func(c echo.Context) error {
		//控制器函数直接返回一个字符串，http响应状态为http.StatusOK，就是200状态。
		return c.String(http.StatusOK, "欢迎访问tizi365.com")
	})

	//启动http server, 并监听8080端口，冒号（:）前面为空的意思就是绑定网卡所有Ip地址，本机支持的所有ip地址都可以访问。
	e.Start(":8080")
}
