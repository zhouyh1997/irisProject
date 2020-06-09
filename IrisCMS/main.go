package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/sessions"
	"time"
)

/*
 *程序主入口
 */

func main() {
	app := newApp()

	//应用app设置
	configuration(app)
	//路由设置
	//mvcHandle(app)
	//config := config.InitConfig()
	app.Run(iris.Addr(":9000"))
}

//构建APP
func newApp() *iris.Application {
	app := iris.New()
	//设置日志级别
	app.Logger().SetLevel("debug")
	//注册静态资源
	app.HandleDir("/static", "./static")
	app.HandleDir("/manage/static", "./static")
	app.HandleDir("/img", "./static/img")
	//注册视图文件
	app.RegisterView(iris.HTML("./static", "html"))
	app.Get("/", func(context context.Context) {
		context.View("index.html")
	})

	return app
}

//项目设置
func configuration(app *iris.Application) {
	//配置字符编码
	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))
	//错误配置
	//未发现错误
	app.OnErrorCode(iris.StatusNotFound, func(context context.Context) {
		context.JSON(iris.Map{
			"err_msg": iris.StatusNotFound,
			"msg":     "not found",
			"date":    iris.Map{},
		})
	})
	app.OnErrorCode(iris.StatusInternalServerError, func(context context.Context) {
		context.JSON(iris.Map{
			"err_msg": iris.StatusInternalServerError,
			"msg":     "internal error",
			"date":    iris.Map{},
		})
	})
}

//MVC架构模式处理
func mvcHandle(app *iris.Application) {
	//启用session
	sessionManager := sessions.New(sessions.Config{
		Cookie:                      "sessioncookie",
		CookieSecureTLS:             false,
		AllowReclaim:                false,
		Encode:                      nil,
		Decode:                      nil,
		Encoding:                    nil,
		Expires:                     24 * time.Hour,
		SessionIDGenerator:          nil,
		DisableSubdomainPersistence: false,
	})
	//engine := datasource.NewMySQLEngine()

	//管理员模块功能
	//adminService :=
}
