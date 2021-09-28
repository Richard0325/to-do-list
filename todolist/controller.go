package todolist

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var dao Dao

func Init() {
	// dao = InitDao(MemoryDaoType, 50)
	dao = InitDao(MariaDaoType, 50)
}

func GetTasks(c *gin.Context) {
	data, err := dao.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("ID error: %s", err.Error()),
		})
		return
	}
	data, err := dao.GetTask(int(id))
	if err != nil {
		if err == ErrNotFound {
			c.JSON(http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}
		return
	}
	c.JSON(http.StatusOK, data)
}

func AddTask(c *gin.Context) {
	t := Task{}
	err := c.BindJSON(&t)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("JSON parse error: %s", err.Error()),
		})
		return
	}

	id, err := dao.AddTask(&t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}
	t.ID = id
	c.JSON(http.StatusOK, t)
}

func DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("ID error: %s", err.Error()),
		})
		return
	}
	err = dao.DeleteTask(int(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]string{})
}
