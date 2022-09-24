package Handle

import (
	"main/app/Business/TodoList"
	"strconv"
	"time"
	"errors"

	"github.com/gin-gonic/gin"
)

func GetTodo(c *gin.Context) (interface{}, error) {
	uidRaw := c.Query("uid")
	statusRaw := c.Query("status")

	if uidRaw == "" || statusRaw ==""{
		return nil, errors.New("參數為空")
	}

	uid, err := strconv.ParseInt(uidRaw, 10, 64)
	if err != nil {
		return nil, err
	} 

	status, err := strconv.ParseInt(statusRaw, 10, 64)
	if err != nil {
		return nil, err
	}

	result := TodoList.ListByStatus(uid ,status)
	return  result, nil 
}

func AddTodo(c *gin.Context) (interface{}, error) {
	uidRaw := c.Query("uid")
	if uidRaw == "" {
		return nil, errors.New("參數為空")
	}

	uid, err := strconv.ParseInt(uidRaw, 10, 64)
	if err != nil {
		return nil, err
	} 

	var todo TodoList.Todo
	err = c.ShouldBind(&todo)
	if err != nil {
		return  nil, err
	}

	// 加入建立時間
	todo.Status = 1
	todo.CreatAt = time.Now().Unix()

	TodoList.AddTodo(uid ,todo)
	return  "success", nil 
}

func DeleteTodo(c *gin.Context) (interface{}, error) {
	uidRaw := c.Query("uid")
	creatRaw := c.Query("creat_at")

	if uidRaw == "" || creatRaw == "" {
		return nil, errors.New("參數為空")
	}

	uid, err := strconv.ParseInt(uidRaw, 10, 64)
	if err != nil {
		return nil, err
	} 

	creatAt, err := strconv.ParseInt(creatRaw, 10, 64)
	if err != nil {
		return nil, err
	} 

	isSuccess := TodoList.DeleteTodo(uid ,creatAt)
	
	if !isSuccess {
		return nil, errors.New("刪除工作項目失敗")
	}

	return  "success", nil 
}
