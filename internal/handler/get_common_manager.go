package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zdarovich/org-hierarchy/internal/lca"
	"net/http"
	"strings"
)

type (
	CommonManagerHandler struct {
		NodeIds    map[string]int
		Names      map[int]string
		Upward     *lca.UpwardTree
		RootNodeId int
	}

	GetCommonManagerResponse struct {
		Name string `json:"name"`
	}
)

func (h *CommonManagerHandler) GetCommonManager(c *gin.Context) {

	p := c.Query("employees")
	if len(p) == 0 {
		c.JSON(http.StatusBadRequest, GetErrorResponse("'employees' param is empty", nil))
		return
	}
	employees := strings.Split(p, ",")
	if len(employees) != 2 {
		c.JSON(http.StatusBadRequest, GetErrorResponse("'employees' param list length must be equal 2", nil))
		return
	}
	emp1NodeId := h.NodeIds[employees[0]]
	emp2NodeId := h.NodeIds[employees[1]]
	nodeId := h.Upward.FindCommonManager(h.RootNodeId, emp1NodeId, emp2NodeId)
	if nodeId == nil {
		c.JSON(http.StatusNotFound, GetErrorResponse("common manager was not found", nil))
		return
	}
	req := GetCommonManagerResponse{
		Name: h.Names[*nodeId],
	}
	c.JSON(http.StatusOK, GetResponse("result", req))
}
