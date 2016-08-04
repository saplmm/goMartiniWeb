package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"log"
	"goMartiniWeb/service/syService"
)

func init(){
	syService.SysInit()	//做系统初始化
}

func main() {
	m := martini.Classic()

	//m.RunOnAddr("ip:por")		//设置ip地址
	//测试中间件
	m.Use(testMid)
	//中间件
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	//404的页面
	m.NotFound(func() string{
		return "傻逼,没找到东西吧?哈哈哈哈"
	})


	m.Get("/render", func(r render.Render) {
		r.HTML(200, "hello", "小知乐")
	})



	m.Get("/", func() string {
		return "Hello world!"
	})

	m.Get("/index", func(r render.Render) {
		r.HTML(200, "hello", "小知乐")
	})

	m.Run()
}


func testMid(c martini.Context, l *log.Logger){
	l.Println("这是测试中间件,哈哈哈哈")
	c.Next()
	l.Println("after a request")
}