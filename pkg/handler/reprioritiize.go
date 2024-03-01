package handler

import (
	hezzl "Hezzl"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GoodRepriority struct {
	Id            int    `json:"-"`
	Project_id    int    `json:"project_id"`
	Description   string `json:"description"`
	NewPririority int    `json:"newPriority"`
}

type Priority struct {
	Id       int
	Priority int
}

type Priorities struct {
	Priorities []Priority
}

func (h *Handler) reprioritiize(c *gin.Context) {
	var input_ GoodRepriority

	if err := c.BindJSON(&input_); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	input := hezzl.Good{
		Id:         input_.Id,
		Project_id: input_.Project_id,
		Priority:   input_.NewPririority,
	}

	urlParams := c.Request.URL.Query()

	if len(urlParams["projectId"]) == 1 {
		project_id, err := strconv.Atoi(urlParams["projectId"][0])
		if err != nil {
			panic(err)
		}

		input.Project_id = project_id
	} else {
		NewErrorResponse(c, http.StatusBadRequest, `No url "projectId" parametr`)
	}

	if len(urlParams["id"]) == 1 {
		id, err := strconv.Atoi(urlParams["id"][0])
		if err != nil {
			panic(err)
		}

		input.Id = id
	} else {
		NewErrorResponse(c, http.StatusBadRequest, `No url "id" parametr`)
	}

	priority, err := h.services.Reprioritiize.ReprioritiizeGood(input)
	var priorities []Priority
	priorities = append(priorities, Priority{
		Id:       input.Id,
		Priority: input.Priority,
	})
	if priority == -1 {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    3,
			"message": "errors.good.notFound",
			"details": "",
		})
	} else {
		c.JSON(http.StatusOK, Priorities{
			Priorities: priorities,
		})
	}
	fmt.Println(priority, err)

}
