package main

import (
	"bytes"
	"fmt"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type LegacyTransaction struct {
	Nonce    uint64
	Price    *big.Int
	GasLimit uint64
	To       *common.Address
	Amount   *big.Int
	Data     []byte
	V        *big.Int
	R        *big.Int
	S        *big.Int
}

type EnvelopeTransaction struct {
	Ty      uint
	Payload []byte
}

type FlatTransaction struct {
	Ty       uint
	Nonce    uint64
	Price    *big.Int
	GasLimit uint64
	To       *common.Address
	Amount   *big.Int
	Data     []byte
	V        *big.Int
	R        *big.Int
	S        *big.Int
}

type LazyTransaction struct {
	Ty uint
	A  rlp.RawValue `rlp:"tail"`
}

func main() {
	tx := LegacyTransaction{
		Nonce:    123,
		Price:    big.NewInt(10000000000),
		GasLimit: 25000,
		To:       &common.Address{},
		Amount:   common.Big0,
		Data:     make([]byte, 10),
		V:        big.NewInt(37),
		R:        common.Big0,
		S:        common.Big0,
	}

	tx.R, _ = tx.R.SetString("28ef61340bd939bc2195fe537567866003e1a15d3c71ff63e1590620aa636276", 16)
	tx.S, _ = tx.S.SetString("67cbe9d8997f761aecb703304b3800ccf555c9f3dc64214b297fb1966a3b6d83", 16)

	raw, _ := rlp.EncodeToBytes(&tx)
	fmt.Printf("Legacy(%d): %X\n", len(raw), raw)

	flat(tx)
	lazy(tx)
	envelope(tx)
}

func flat(l LegacyTransaction) {
	tx := FlatTransaction{
		Ty:       0,
		Nonce:    l.Nonce,
		Price:    l.Price,
		GasLimit: l.GasLimit,
		To:       l.To,
		Amount:   l.Amount,
		Data:     l.Data,
		V:        l.V,
		R:        l.R,
		S:        l.S,
	}

	raw, _ := rlp.EncodeToBytes(&tx)
	fmt.Printf("flat(%d): %X\n", len(raw), raw)
}

func envelope(l LegacyTransaction) {
	payload, _ := rlp.EncodeToBytes(l)
	etx := EnvelopeTransaction{Ty: 0, Payload: payload}
	raw, _ := rlp.EncodeToBytes(&etx)

	var retx EnvelopeTransaction
	rlp.Decode(bytes.NewReader(raw), &retx)

	if retx.Ty == 0 {
		var ll LegacyTransaction
		rlp.Decode(bytes.NewReader(retx.Payload), &ll)
		if reflect.DeepEqual(l, ll) != true {
			panic("envelope roundtrip failed")
		}
	} else {
		panic("should always be 0")
	}

	fmt.Printf("envelope(%d): %X\n", len(raw), raw)
}

func lazy(l LegacyTransaction) {
	payload, _ := rlp.EncodeToBytes(l)
	raw := append([]byte{0xEE, 0x80}, payload[1:]...)

	var ltx LazyTransaction
	rlp.DecodeBytes(raw, &ltx)

	// var ll LegacyTransaction
	// rlp.DecodeBytes(append([]byte{0xED}, flattened...), &ll)

	// if reflect.DeepEqual(l, ll) != true {
	//         panic("lazy roundtrip failed")
	// }

	fmt.Printf("lazy(%d): %X\n", len(raw), raw)
}
