//Operations upon ledger.txt file

package main

import (
	"bufio"
	"log"
	//"bufio"
	"os"
	//"fmt"
)

func getAllBlocks() (string){
	filePath := "MyLedger.txt"
	
	//opening file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}



//How a Line should start?
	line := "["
	defer file.Close()


	scanner := bufio.NewScanner(file)  //Fetching by Line


	if(scanner.Scan()){
		line += scanner.Text()
	}

	for scanner.Scan() {
		line += ","
		line +=  scanner.Text()
		
	}
	line += "]"
	//fmt.Println(line)

	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
	return line
}

func getBlockById(id int)(string){
	filePath := "MyLedger.txt"
	
	//opening file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	
	defer file.Close()

	scanner := bufio.NewScanner(file)  //To read each line

	
	for scanner.Scan() {
		id -= 1
		//fmt.Println(id)
		line := scanner.Text()
		if id == 0{
			return line
		}
		
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
	return "Data doesn't Exist"
}

	