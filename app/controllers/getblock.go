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
	var result models.Response

	params := mux.Vars(r)

	blockHash := params["blockHash"]

	network := params["network"]

	url := "https://sochain.com/api/v2/get_block/"+network+"/"+blockHash
	log.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	
	// fmt.Print(resp.Body)

	defer resp.Body.Close()
	 body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	
	//Convert body to string
	
	sb := string(body)
	log.Printf(sb)

	// we read response body
	// if err := json.NewDecoder(resp.Body).Decode(&block); err != nil {
	// 	logFatal(err)
	//  }
	if err := json.Unmarshal(body, &result); err != nil {   // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	fmt.Println(result)
	 helpers.JSON(w, http.StatusOK, result)

	
}