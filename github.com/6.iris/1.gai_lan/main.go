package main

import (
	"github.com/kataras/iris"
	"net"
	"net/http"
)

func main() {
	app := iris.New()

	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		errMessage := ctx.Values().GetString("error")
		if errMessage != "" {
			_, _ = ctx.Writef("Internal server error: %s", errMessage)
			return
		}

		_, _ = ctx.Writef("(Unexpected) internal server error")
	})

	app.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Begin request for path: %s", ctx.Path())
		ctx.Next()
	})

	app.Done(func(ctx iris.Context){})

	app.Run(iris.Addr(":8080"))

	app.Run(iris.Server(&http.Server{Addr: ":8080"}))

	l, err := net.Listen("tcp4", ":8080")
	if err != nil {
		panic(err)
	}
	app.Run(iris.Listener(l))

}
