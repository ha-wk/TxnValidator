package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	//"strconv"

	//"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

type Leveldbinterface interface {
	NewLevelDB(dbPath string)(*LevelDB,error)
	Put(key string ,newEntry LdgrTxn)(error)
	Get(key string)(LdgrTxn,error)
}

type LevelDB struct {
	db *leveldb.DB
}

func Create_Database(dbPath string)(*LevelDB) {
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil{
		fmt.Println("Error in Creating Database")
	}
	return &LevelDB{db:db}

}
/*func PopulateDB(db *leveldb.DB) {
	for i := 1; i <= 1000; i++ {
		key := fmt.Sprintf("SIM%d", i)
		value := fmt.Sprintf(`{"val": %d, "ver": %f}`, i, 1.0)

		err := db.Put([]byte(key), []byte(value), nil)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Database populated successfully")
}
func (ldb *LevelDB) Close() {
	ldb.db.Close()
}*/


//var db *leveldb.DB


func (ldb *LevelDB) Get(key string) (LocalTxnInfo, error) {
    data, err := ldb.db.Get([]byte(key), nil)
    if err != nil {
        return LocalTxnInfo{}, err
    }

    var TxnById LocalTxnInfo
    err = json.Unmarshal(data, &TxnById)

    return TxnById,err
}

func(ldb *LevelDB)Add(key string, newEntry LocalTxnInfo) error {
    newData, err := json.Marshal(newEntry)
    if err != nil {
        return err
    }

    return ldb.db.Put([]byte(key), newData, nil)

}


// data,err := ldb.db.Get([]byte(key), nil)
// 	var txndata Txn_inp
// 	err = json.Unmarshal([]byte(data), &txndata)
// 	return txndata, nil
// }

// //getall() left


func (ldb *LevelDB)GetallInCsv()error{

	 Result, err := os.Create("Result.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer Result.Close()

	Update := csv.NewWriter(Result)
	defer Update.Flush()

	header := []string{"key", "value"}
	Update.Write(header)

	iter := ldb.db.NewIterator(nil , nil)

	for iter.Next() {
			record := []string{
				string(iter.Key()),
				string(iter.Value()),
			}
			Update.Write(record)
	}
	iter.Release()
	return iter.Error()
}














/*

/*func (ldb *LevelDB)NewLevelDB(dbPath string) (*LevelDB, error) {
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		return nil, err
	}

	for i := 1; i <= 1000; i++ {
		key := fmt.Sprintf("SIM%d", strconv.Itoa(i))
		
		value := fmt.Sprintf(`{"val": %d, "ver": %f,"valid":%v}`, i, 1.0, false)


		err := db.Put([]byte(key), []byte(value), nil)

		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Database populated successfully")



	return &LevelDB{db: db}, nil
}*/




/*func (ldb *LevelDB) Get(key string) (Txn_inp, error) {
	data,err := ldb.db.Get([]byte(key), nil)
	var txndata Txn_inp
	err = json.Unmarshal([]byte(data), &txndata)
	return txndata, nil
}

//getall() left


func (ldb *LevelDB) Put(key, value Txn_inp) error {

	jsonstr, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return ldb.db.Put([]byte(key), []byte(jsonstr), nil)

}



func validateTransaction(key string, newTransaction Transaction) bool {
    existingEntry, err := GetEntry(key)
    if err != nil {
        return err
    }

    // Perform your validation logic here
    if newTransaction.Val < existingEntry.Ver {
        return fmt.Errorf("Validation failed: New value is less than existing value")
    }

    return nil
}


*/