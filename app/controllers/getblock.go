package controllers

import (
	// "fmt"
	"fmt"
	"net/http"
	// "strconv"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gorilla/mux"

	"example.com/helpers"
	"example.com/models"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//TODO Getblock route
func (controller *Controllers) GetBlock(w http.ResponseWriter, r *http.Request) {
	var block models.Response
	var transactions models.TrxResponse
	var result models.Result

	params := mux.Vars(r)

	blockHash := params["blockHash"]

	network := params["network"]

	url := "https://sochain.com/api/v2/get_block/" + network + "/" + blockHash
	log.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}


	if err := json.Unmarshal(body, &block); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	record,_:=json.Marshal(block.Data)
	json.Unmarshal([]byte(record), &result)
	
	for _, rec := range block.Data.TransactionRefs[0:10] {
		trxUrl := "https://sochain.com/api/v2/tx/" + network + "/" + rec

		resp, err := http.Get(trxUrl)
		if err != nil {
			log.Fatalln(err)
		}


		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		if err := json.Unmarshal(body, &transactions); err != nil { // Parse []byte to go struct pointer
			fmt.Println("Can not unmarshal JSON2")
		}
		result.Transactions = append(result.Transactions,transactions.Data)
	}
	helpers.JSON(w, http.StatusOK, result)

}
