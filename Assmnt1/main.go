package main

import (
	//"encoding/json"
	"fmt"
	"log"
	//"log"
	//	"io"
	//"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/syndtr/goleveldb/leveldb"
)

var db=Create_Database("db")

var BlkDtls TnxFields


func main() {
    
	 BlkDtls.Initialization()
	 fmt.Println("Printing the initial Default Block")
	 fmt.Println(BlkDtls)


	 fmt.Println("Also calling TransactionAdd Function")
	 BlkDtls.AddTransaction()

	 fmt.Println("Performing Operations on File")
	 BlkDtls.WriteToFile()
	 BlkDtls.AutoWrite()

	 fmt.Println(BlkDtls)


	 router := gin.Default()
	 
	 SetupRoutes(router) // Use the SetupRoutes function
 
	 log.Println("Server listening on port 8080...")
	 router.Run(":8080")


}










//go router.GET("/admin/reset" , resetDBHandler)
