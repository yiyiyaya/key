package main

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"bytes"
	"github.com/btcsuite/btcd/btcec"	
	//"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/keystore"
//	"github.com/pborman/uuid"
	

)

var rander io.Reader
func main() {
	// priv, err := generateAndPriKey("foo")
	randBytes := make([]byte, 64)
	_, err := rand.Read(randBytes)
	if err != nil {
		panic("key generation: could not read from random source: " + err.Error())
	}
	reader := bytes.NewReader(randBytes)
	key,pubkeyvalstr, err := newKey(reader)
	pubkeyStr := hex.EncodeToString([]byte(pubkeyvalstr))
	fmt.Println("publickey-----0001===", pubkeyStr)
	//key, err := newKey(rand.Reader)
	if err != nil {
		fmt.Printf("newkey err", err)
	}
	// var p *btcec.PrivateKey
	//privatekey := (*btcec.PrivateKey)key
	priv := hex.EncodeToString((*btcec.PrivateKey)(key.PrivateKey).Serialize())
	//priv := hex.EncodeToString(key.PrivateKey.Serialize())
	//priv := hex.EncodeToString(privatekey.Serialize())
	fmt.Printf("PrivKey=====:%v", priv)
	 //key := newKeyFromECDSA(priv)
	address := key.Address.Bytes()
	derivedAddr := hex.EncodeToString(address)
	fmt.Printf("address=====:%v", derivedAddr)
	//pubkey, err := generatePubKey()
	//pubkey0 := crypto.FromECDSAPub(priv.PublicKey)
	pubkey0 := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	pubkey := hex.EncodeToString(pubkey0)
	fmt.Printf("==============PubKey: %v", pubkey)

}
func newKey(rand io.Reader) (*keystore.Key,string, error) {
	privateKeyECDSA, err := ecdsa.GenerateKey(crypto.S256(), rand)
	if err != nil {
		return nil,"", err
	}
	//pubKeyECDSA := crypto.FromECDSAPub(&privateKeyECDSA.PublicKey)
	pubKeyECDSA2 :=(*btcec.PublicKey)(&privateKeyECDSA.PublicKey).SerializeUncompressed()

	return  newKeyFromECDSA(privateKeyECDSA),string(pubKeyECDSA2), nil
}

// func generateAndPriKey(password string) (key *Key, err error) {
// 	preSaleKeyStruct := struct {
// 		EncSeed string
// 		EthAddr string
// 		Email   string
// 		BtcAddr string
// 	}{}
// 	encSeedBytes, err := hex.DecodeString(preSaleKeyStruct.EncSeed)
// 	if err != nil {
// 		return nil, errors.New("invalid hex in encSeed")
// 	}
// 	iv := encSeedBytes[:16]
// 	cipherText := encSeedBytes[16:]
// 	passBytes := []byte(password)
// 	derivedKey := pbkdf2.Key(passBytes, passBytes, 2000, 16, sha256.New)
// 	plainText, err := keystore.aesCBCDecrypt(derivedKey, cipherText, iv)
// 	if err != nil {
// 		return nil, err
// 	}
// 	ethPriv := crypto.Keccak256(plainText)
// 	ecKey := crypto.ToECDSAUnsafe(ethPriv)
// 	return ecKey, err
// }

/*
type Key struct {
	Id         uuid.UUID
	Address    common.Address
	PrivateKey *ecdsa.PrivateKey
}
*/

// randomBits completely fills slice b with random data.
func randomBits(b []byte) {
	if _, err := io.ReadFull(rander, b); err != nil {
		panic(err.Error()) // rand should never fail
	}
}
type UUID []byte
func NewRandom() UUID {
	uuid := make([]byte, 16)
	randomBits([]byte(uuid))
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant is 10
	return uuid
}
func newKeyFromECDSA(privateKeyECDSA *ecdsa.PrivateKey) *keystore.Key {
	//id := NewRandom()
	//id := uuid.NewRandom()
//fmt.Println("id",id)
	key := &keystore.Key{
		//Id:         id,
		Address:    crypto.PubkeyToAddress(privateKeyECDSA.PublicKey),
		PrivateKey: privateKeyECDSA,
	}
	return key
}
