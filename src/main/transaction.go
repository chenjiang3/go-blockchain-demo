package main

import "bytes"

const subsidy = 0

type Transaction struct {
	ID 		[]byte
	Vin 	[]TXInput
	Vout 	[]TXOutput
}

func (tx Transaction) isCoinbase() bool {
	return len(tx.Vin) == 1 && len(tx.Vin[0].Txid) == 0 && tx.Vin[0].Vout == -1
}

func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

}

type TXInput struct {
	Txid 		[]byte
	Vout 		int
	ScriptSig 	string
}

type TXOutput struct {
	Value 		int
	ScriptSig 	string
}

