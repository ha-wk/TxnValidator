package main

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/Add", PutInitialEntry)  //Entries of txns that need to be validated with existing ones(max-lim-5)
	router.GET("/Blocks", GetAllBlocks)  //to get all existing blocks
	router.GET("/Block:id", GetBlockById) //to get particular block info
	router.GET("/AllEntries", PrintDB)   //To get updated DB in .csv file after succesfull validation operation
}