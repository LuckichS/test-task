package handler

import (
	hezzl "Hezzl"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) create(c *gin.Context) {
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

	id, priority, created_at, err := h.services.Create.CreateGood(input)
	//fmt.Println(input, id, priority, created_at)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":          id,
		"projectId":   input.Project_id,
		"name":        input.Name,
		"description": input.Description,
		"priority":    priority,
		"removed":     false,
		"createdAt":   created_at,
	})
}
