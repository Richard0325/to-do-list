package todolist

import (
	"fmt"
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
)

var dao Dao

func GenResponse(data interface{}) map[string]interface{}{
	return map[string]interface{}{
		"msg": "",
		"data": data,
	}
}
type httpErrResponseType int
var httpErrNotFound httpErrResponseType = 1
var httpErrBadRequest httpErrResponseType = 2
var httpErrOthers httpErrResponseType = 3
func GenErrResponse(c *gin.Context, err error, errType httpErrResponseType){
	switch errType{
	case httpErrNotFound:
		c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	case httpErrBadRequest:
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("JSON parse error: %s", err.Error()),
		})
	case httpErrOthers:
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return
}

var ErrInvalidArg error = errors.New("Invalid Argument")

func Init(dbType int, size int) error{
	switch dbType{
	case 1:
		dao = InitDao(MemoryDaoType, size)
		fmt.Println("Memory is using now")
		return nil
	case 2:
		dao = InitDao(MariaDaoType, size)
		fmt.Println("MariaDB is using now")
		return nil
	case 3:
		dao = InitDao(MongoDaoType, size)
		fmt.Println("MongoDB is using now")
		return nil
	default:
		fmt.Println("invalid DB type")
		return ErrInvalidArg
	}
}

func GetTasks(c *gin.Context) {
	data, err := dao.GetTasks()
	if err != nil {
		GenErrResponse(c, err, httpErrOthers)
		return
	}
	c.JSON(http.StatusOK, GenResponse(data))
}

func GetTask(c *gin.Context) {
	idStr := c.Param("id")
	data, err := dao.GetTask(idStr)
	if err != nil {
		if err == ErrNotFound {
			GenErrResponse(c, err, httpErrNotFound)
		} else {
			GenErrResponse(c, err, httpErrOthers)
		}
		return
	}
	c.JSON(http.StatusOK, GenResponse(data))
}

func AddTask(c *gin.Context) {
	t := Task{}
	err := c.BindJSON(&t)
	if err != nil {
		GenErrResponse(c, err, httpErrBadRequest)
		return
	}

	id, err := dao.AddTask(&t)
	if err != nil {
		GenErrResponse(c, err, httpErrOthers)
		return
	}
	t.ID = id
	c.JSON(http.StatusOK, GenResponse(t))
}

func DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	err := dao.DeleteTask(idStr)
	if err == ErrInvalidInput {
		GenErrResponse(c, err, httpErrBadRequest)
	}else{
		GenErrResponse(c, err, httpErrOthers)
	}
	c.JSON(http.StatusOK, map[string]string{})
}