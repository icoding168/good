package main

/*
$ curl "http://localhost:9999"
Hello Geektutu
$ curl "http://localhost:9999/panic"
{"message":"Internal Server Error"}
$ curl "http://localhost:9999"
Hello Geektutu

>>> log
2020/01/09 01:00:10 Route  GET - /
2020/01/09 01:00:10 Route  GET - /panic
2020/01/09 01:00:22 [200] / in 25.364µs
2020/01/09 01:00:32 runtime error: index out of range
Traceback:
        /usr/local/Cellar/go/1.12.5/libexec/src/runtime/panic.go:523
        /usr/local/Cellar/go/1.12.5/libexec/src/runtime/panic.go:44
        /Users/7days-golang/day7-panic-recover/main.go:47
        /Users/7days-golang/day7-panic-recover/gee/context.go:41
        /Users/7days-golang/day7-panic-recover/gee/recovery.go:37
        /Users/7days-golang/day7-panic-recover/gee/context.go:41
        /Users/7days-golang/day7-panic-recover/gee/logger.go:15
        /Users/7days-golang/day7-panic-recover/gee/context.go:41
        /Users/7days-golang/day7-panic-recover/gee/router.go:99
        /Users/7days-golang/day7-panic-recover/gee/gee.go:130
        /usr/local/Cellar/go/1.12.5/libexec/src/net/http/server.go:2775
        /usr/local/Cellar/go/1.12.5/libexec/src/net/http/server.go:1879
        /usr/local/Cellar/go/1.12.5/libexec/src/runtime/asm_amd64.s:1338

2020/01/09 01:00:32 [500] /panic in 395.846µs
2020/01/09 01:00:38 [200] / in 6.985µs
*/

import (
	"fmt"
	"geeorm/session"
	"html/template"
	"net/http"
	"time"

	"gee"
	"geeorm"
	_ "github.com/mattn/go-sqlite3"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func main() {
	r := gee.New()
	r.Use(gee.Logger())
	r.Use(gee.Recovery())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "Geektutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *gee.Context) {
		engine, err := geeorm.NewEngine("sqlite3", "gee.db")
		defer engine.Close()
		s := engine.NewSession()
		_ = s.Model(&User{}).DropTable()
		_ = s.Model(&User{}).CreateTable()
		_, err = engine.Transaction(func(s *session.Session) (result interface{}, err error) {
			_, err = s.Model(&User{}).Insert(&User{"Tom", 18})
			return nil, err
		})

		if err == nil || s.HasTable() {
			c.HTML(http.StatusOK, "arr.tmpl", gee.H{
				"title":  "gee",
				"stuArr": [2]*student{stu1, stu2},
			})
		} else {
			c.HTML(http.StatusInternalServerError, "arr.tmpl", gee.H{
				"title":  "gee",
				"stuArr": [2]*student{stu1, stu2},
			})
		}

	})

	r.GET("/date", func(c *gee.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
			"title": "gee",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})

	//r := gee.Default()
	//r.GET("/", func(c *gee.Context) {
	//	c.String(http.StatusOK, "Hello Geektutu\n")
	//})
	//// index out of range for testing Recovery()
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
