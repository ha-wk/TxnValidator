package main

import (
	//"encoding/json"
	//"fmt"
	"log"
	
	"github.com/gin-gonic/gin"
	//"github.com/syndtr/goleveldb/leveldb"
	"github.com/ha-wk/Assmnt1/server"
  //  "github.com/ha-wk/Assmnt1/database"
	"github.com/ha-wk/Assmnt1/models"


)




func main() {
    
	models.DefaultInit()


	 router := gin.Default()
	 
	 server.SetupRoutes(router) // Use the SetupRoutes function
 
	 log.Println("Server listening on port 8080...")
	 router.Run(":8080")


}










//go router.GET("/admin/reset" , resetDBHandler)
