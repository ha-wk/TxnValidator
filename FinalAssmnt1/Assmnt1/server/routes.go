package server

import (
	"github.com/gin-gonic/gin"
)

  

func SetupRoutes(router *gin.Engine) {

	// fileStruct := &FileStruct {
    //     Service: &FileServiceImpl{},// Initialize your FileService here
    // }
	RouteService := &RouteServiceImpl{}

	router.POST("/Add", RouteService.PutInitialEntry)  //Entries of txns that need to be validated with existing ones(max-lim-5)
	router.GET("/Blocks", RouteService.GetAllBlocks)  //to get all existing blocks
	router.GET("/Block:id", RouteService.GetBlockById) //to get particular block info
	router.GET("/AllEntries",RouteService. PrintDB)   //To get updated DB in .csv file after succesfull validation operation
}