package models

import (
	"time"

)


// Network Code
// Block Number
// Date/Time in the following format: dd-MM-yyyy HH:mm (01-11-2020 14:02) Previous Blockhash
// Next BlockHash
// Size
// The first ten transactions in the block including the following information for each transaction
// ○ Transaction Id (txid
// ○ Date/Time in the following format: dd-MM-yyyy HH:mm (01-11-2020
// 14:02)
// ○ The transaction fee (fee)
// ○ The sent value (sent_value)

type Transactions struct {
	TransactionId string `json:"txid"`
	TimeStamp     Epoch `json:"time"`
	Fee     	  string `json:"fee"`
	SentValue     string `json:"sent_value"`
}

type Epoch int64

func (t Epoch) MarshalJSON() ([]byte, error) {
    strDate := time.Time(time.Unix(int64(t), 0)).Format(time.RFC3339)
    out := []byte(`"` + strDate + `"`)
    return out, nil
}
type Block struct {
	Network 	  string `json:"network"`
	BlockNumber   int 	 `json:"block_no"`
	TimeStamp     Epoch    `json:"time"`
	PreviousBlockHash string `json:"previous_blockhash"`
	NextBlockHash string `json:"next_blockhash"`
	Size		  int    `json:"size"`
	TransactionRefs []string `json:"txs"`
}

type Response struct {
	Status string `json:"status"`
	Data Block `json:"data"`
}

type TrxResponse struct {
	Status string `json:"status"`
	Data Transactions `json:"data"`
}


type Result struct {
	Network 	  string `json:"network"`
	BlockNumber   int 	 `json:"block_no"`
	TimeStamp     Epoch    `json:"time"`
	NextBlockHash string `json:"next_blockhash"`
	Size		  int    `json:"size"`
	Transactions []Transactions
}