package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"encoding/json"
	"github.com/ha-wk/Assmnt1/database"
	"github.com/ha-wk/Assmnt1/ledger"
	//"github.com/ha-wk/Assmnt1"
	
)

type DefaultTxn struct {
	SIM map[string] database.LocalTxnInfo`json:"SIM"`
}


type RoutesService interface{
	PutInitialEntry()
	GetAllBlocks()
	GetBlockById()
	PrintDB()
}

type RouteServiceImpl struct{

}

var db=database.Create_Database("db")


func(R *RouteServiceImpl) PutInitialEntry(c *gin.Context) {

	
	var Inp_txn []map[string]database.LocalTxnInfo
	if err := c.ShouldBindJSON(&Inp_txn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.PushValtxns(Inp_txn)

	c.JSON(http.StatusOK, gin.H{"message": "Insertion Succesfull"})
}

func (R *RouteServiceImpl)GetAllBlocks(c *gin.Context){
    
	fileService := &ledger.FileServiceImpl{}
	data := fileService.GetAllBlks()
	var jsonData interface{}
	err := json.Unmarshal([]byte(data), &jsonData)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error parsing JSON data")
		return
	}
	c.JSON(http.StatusOK, gin.H{ "message": jsonData})
}

func (R *RouteServiceImpl)GetBlockById(c *gin.Context){
	id_str := c.Param("id")
	id , _ := strconv.Atoi(id_str)

	fileService := &ledger.FileServiceImpl{}
	data := fileService.GetBlkById(id)

	var jsonData interface{}
	err := json.Unmarshal([]byte(data), &jsonData)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error parsing JSON data")
			return
		}
	c.JSON(http.StatusOK, jsonData)
}

func (R *RouteServiceImpl)PrintDB(c *gin.Context){
	db.GetallInCsv()
	c.JSON(http.StatusOK, gin.H{"message": "SUCCESSFULLY PRINTED IN CSV EXTERNAL FILE"})
}









// func resetDBHandler(c *gin.Context){
// 	db.PopulateDB()
// 	c.JSON(http.StatusOK, gin.H{"message": "BACK TO DEFAULT MODE"})
// }