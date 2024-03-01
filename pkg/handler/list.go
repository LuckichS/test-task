package handler

import (
	hezzl "Hezzl"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type metaData struct {
	Total   int
	Removed int
	Limit   int
	Offset  int
}
type getAllListsResponse struct {
	Meta  metaData     `json:data`
	Goods []hezzl.Good `json:data`
}

func (h *Handler) list(c *gin.Context) {
	var input hezzl.Good

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	urlParams := c.Request.URL.Query()

	var limit int
	if len(urlParams["limit"]) == 1 {
		limit_, err := strconv.Atoi(urlParams["limit"][0])
		if err != nil {
			panic(err)
		}

		limit = limit_
	} else {
		NewErrorResponse(c, http.StatusBadRequest, `No url "limit" parametr`)
	}

	var offset int
	if len(urlParams["offset"]) == 1 {
		offset_, err := strconv.Atoi(urlParams["offset"][0])
		if err != nil {
			panic(err)
		}

		offset = offset_
	} else {
		NewErrorResponse(c, http.StatusBadRequest, `No url "offset" parametr`)
	}

	goods, removed, err := h.services.List.ListGood(input, limit, offset)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	meta := metaData{
		Removed: removed,
		Total:   len(goods),
		Offset:  offset,
		Limit:   limit,
	}
	c.JSON(http.StatusOK, getAllListsResponse{
		Meta:  meta,
		Goods: goods,
	})
	//fmt.Println(goods, removed, err)

}
