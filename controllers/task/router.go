package task

import (
	"github.com/kataras/iris/v12/core/router"
	"todoapp/models"
)

func EquipRouter(app router.Party){
	party:= app.Party("/tasks")
	for _, route:= range routes {
		party.Handle(route.Method, route.Pattern, route.HandlerFunc)
	}
}

var routes = models.Routes{
	models.Route{
		"GET",
		"/",
		getAll,
	},
	models.Route{
		"GET",
		"/{taskId}",
		getTaskById,
	},
	models.Route{
		"POST",
		"/",
		createTask,
	},
	models.Route{
		"PATCH",
		"/{taskId}",
		updateTask,
	},
	models.Route{
		"DELETE",
		"/{taskId",
		deleteTask,
	},
}