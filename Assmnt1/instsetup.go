package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"strconv"
)

//LevelDb data format
type LocalTxnInfo struct {
	Value int `json:"val"`
	Version float64 `json:"ver"`
}

//local txn data format
type LedgerFields struct {
	Key string
	Trnx LdgrTxn
}

//INitializing DB with local instance of 1000 entries
func (ldb *LevelDB)PopulateDB(){

	TempData := LocalTxnInfo{Value:1 , Version:1.0}
	for i := 1; i <= 1000; i++ {
		
		key := "SIM" + strconv.Itoa(i) 
		TempData.Value = TempData.Value +1
		TempData.Version=1.0

		err := ldb.Add(key, TempData)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Database populated successfully")
	ldb.GetallInCsv()
	return
}

//Pushing the transactions if it is Valid
func (ldb *LevelDB)PushValtxns(Alltxns []map[string]LocalTxnInfo)(){

	for _, txn := range Alltxns {
		for key, value := range txn {
			go ldb.PushInDb(key , value)
		
		}
	}
}

//updating in LocalDB
func(ldb *LevelDB)PushInDb(key string,trnx LocalTxnInfo){

	var NewTxn LdgrTxn
	NewTxn.Val , NewTxn.Ver = trnx.Value , trnx.Version	
	existingEntry, _ := ldb.Get(key) //modfcs in getentry params
		
  
			  if(trnx.Version == existingEntry.Version){
			   // if validateTransaction(string(key),newEntry) {
				trnx.Version += 1.0
				ldb.Add(key, trnx)
				NewTxn.Valid = true 
				   } else {
				NewTxn.Valid = false          //marking invalid transaction
			  }
		  
			  str := key + strconv.Itoa(trnx.Value) + strconv.FormatFloat(trnx.Version, 'E', -1, 32) 
			  NewTxn.Hash = DeriveHash(str)
		  
			  //defining the pair of key and transaction to push to the channel
			  ledPair  := LedgerFields{
				  Key : key,
				  Trnx: NewTxn,
			  }
		  
			  //checking if there is any deadlock
			  //     var txns []Transaction
//     err := json.Unmarshal(jsonData, &txns)
//     if err != nil {
//         return err
//     }

    // for _, newTxn := range txns {
    //   existingEntry, err := ldb.Get(string(newTxn.ID)) //modfcs in getentry params
    //   if err != nil {
    //     return err
	//   }
			  select {
			  case BlkDtls.TnxKV <- ledPair:
				  fmt.Println(key + "transaction sent to ledger channel successfully /n" )
			  default:
				  fmt.Println(key + "waiting ledger transaction channel  is full /n")
			  }






			}


 //FUNCTION FOR GENERATING HASH OF TXN
 func DeriveHash(data string) string{
	res:=sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x",res)
 }
	  





    //         if newTxn.Ver == existingEntry.Ver && newTxn.ID == existingEntry.ID {
    //          // if validateTransaction(string(key),newEntry) {
    //              existingEntry.Ver += 1.0
	// 			 existingEntry.Val=newTxn.Val
	// 			 existingEntry.Valid = true 
	// 			 } else {
	// 			existingEntry.Valid = false//marking invalid transaction
	// 		}
		
	// 			err = ldb.Put(string(newTxn.ID), existingEntry)
	// 			if err != nil {
	// 				return err
    //         }
    //     }
    
//     return nil
// }
