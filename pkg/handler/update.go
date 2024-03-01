package handler

import (
	hezzl "Hezzl"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) update(c *gin.Context) {
	var input hezzl.Good

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
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

	priority, created_at, err := h.services.Update.UpdateGood(input)
	if priority == -1 {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    3,
			"message": "errors.good.notFound",
			"details": "",
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"id":          input.Id,
			"projectId":   input.Project_id,
			"name":        input.Name,
			"description": input.Description,
			"priority":    priority,
			"removed":     false,
			"createdAt":   created_at,
		})
	}
	fmt.Println(priority, created_at, err)

}
