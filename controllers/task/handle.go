package task

import (
	"context"
	"time"
	"todoapp/configs"
	"todoapp/models"
	"todoapp/responses"

	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// client := configs.ConnectDB()

// coll:= client.Database("todoapp").Collection("tasks")



func getAll(ctx iris.Context) {
	ctxx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	coll:= configs.GetCollection("tasks")

	filter := bson.M{}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
		ctx.JSON(responses.Response{Status: 400, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	var tasks []models.Task
	for cursor.Next(ctxx) {
		var task models.Task
		cursor.Decode(&task)
		tasks = append(tasks, task)
	}
	ctx.JSON(responses.Response{Status: 200, Message: "Success", Data: map[string]interface{}{"data": tasks}})
}


func getTaskById(ctx iris.Context){
	coll:= configs.GetCollection("tasks")
	taskId:= ctx.Params().Get("taskId")

	var task models.Task
	objId, _ := primitive.ObjectIDFromHex(taskId)
	err:= coll.FindOne(context.TODO(), bson.D{{"_id", objId}}).Decode(&task)
	if err!=nil{
		ctx.JSON(responses.Response{Status: 500, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	ctx.JSON(responses.Response{Status: 200, Message: "Success", Data: map[string]interface{}{"data": task}})
	
}

func createTask(ctx iris.Context){
	coll:= configs.GetCollection("tasks")
	var task models.Task
	err:= ctx.ReadJSON(&task)
	if err!=nil{
		ctx.JSON(responses.Response{Status: 500, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	
	// coll:= client.Database("todoapp").Collection("tasks")

	result, err:= coll.InsertOne(context.TODO(), task)
	if err!=nil{
		ctx.JSON(responses.Response{Status: 400, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	ctx.JSON(responses.Response{Status: 201, Message: "Success", Data: map[string]interface{}{"data": result}})

}

func deleteTask(ctx iris.Context){
	coll:= configs.GetCollection("tasks")
	taskId:= ctx.Params().Get("taskId")

	objId, _:= primitive.ObjectIDFromHex(taskId)
	result, err:= coll.DeleteOne(context.TODO(), bson.D{{"_id", objId}})
	if err!=nil{
		ctx.JSON(responses.Response{Status: 500, Message: "ERROR", Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	ctx.JSON(responses.Response{Status: 200, Message: "Delete success", Data: map[string]interface{}{"data": result}})
}

func updateTask(ctx iris.Context){
	coll:= configs.GetCollection("tasks")
	taskId:= ctx.Params().Get("taskId")
	objId, _:= primitive.ObjectIDFromHex(taskId)
	
	var task bson.M

	err:= ctx.ReadJSON(&task)
	if err!=nil{
		ctx.JSON(responses.Response{Status: 500, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	filter := bson.D{{"_id", objId}}
	update := bson.D{{"$set", task}}
	result, err:= coll.UpdateOne(context.TODO(), filter, update)
	if err!=nil{
		ctx.JSON(responses.Response{Status: 500, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	ctx.JSON(responses.Response{Status: 200, Message: "Update success", Data: map[string]interface{}{"data": result}})
	
}