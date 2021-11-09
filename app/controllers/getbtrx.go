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


//TODO Getblock route
func (controller *Controllers) GetTrx(w http.ResponseWriter, r *http.Request) {
	var transactions models.TrxResponse
	var result models.Transactions

	params := mux.Vars(r)

	trxRef := params["trxRef"]

	network := params["network"]

	url := "https://sochain.com/api/v2/tx/" + network + "/" + trxRef
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

	fmt.Println(string(body))


	if err := json.Unmarshal(body, &transactions); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	record,_:=json.Marshal(transactions.Data)
	json.Unmarshal([]byte(record), &result)

	fmt.Println(transactions.Data)

	helpers.JSON(w, http.StatusOK, result)

}
