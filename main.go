package main

import (
	"github.com/kataras/iris/v12"	 
	"todoapp/controllers"
)


func main(){
	app:= iris.New()



	controllers.WithRouter(app)




	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

