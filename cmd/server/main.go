package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/zdarovich/org-hierarchy/internal/handler"
	"github.com/zdarovich/org-hierarchy/internal/lca"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	jsonFile, err := os.Open("./resources/hierarchy.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var hierarchy map[string][]string

	err = json.Unmarshal(byteValue, &hierarchy)
	if err != nil {
		log.Fatal(err)
	}
	upwardTree := lca.New()
	nodeIdsMap := map[string]int{}
	namesMap := map[int]string{}
	for managerName, employeesNames := range hierarchy {
		managerNodeId, exists := nodeIdsMap[managerName]
		if !exists {
			managerNodeId = len(nodeIdsMap) + 1
			nodeIdsMap[managerName] = managerNodeId
		}
		_, exists = namesMap[managerNodeId]
		if !exists {
			namesMap[managerNodeId] = managerName
		}

		for _, employeeName := range employeesNames {
			employeeNodeId, exists := nodeIdsMap[employeeName]
			if !exists {
				employeeNodeId = len(nodeIdsMap) + 1
				nodeIdsMap[employeeName] = employeeNodeId
			}
			_, exists = namesMap[employeeNodeId]
			if !exists {
				namesMap[employeeNodeId] = employeeName
			}

			upwardTree.AddEdge(managerNodeId, employeeNodeId)
		}
	}

	rootNodeId := nodeIdsMap["Clair"]
	// Root node does not have any parent; so set it to -1
	upwardTree.GetParent(rootNodeId, -1)

	h := handler.CommonManagerHandler{
		Names:   namesMap,
		NodeIds: nodeIdsMap,
		Upward:  upwardTree,
	}
	router := gin.Default()
	router.GET("/common/manager", h.GetCommonManager)
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	})
	router.Run(":8080")

}
