package controllers

import (
	"github.com/kataras/iris/v12"
	TaskController "todoapp/controllers/task"
)

func WithRouter(app *iris.Application){
	mainRouter:= app.Party("/")

	TaskController.EquipRouter(mainRouter)

}