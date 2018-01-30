	checkprikaddr("5KD3sSntucZFRzDNUJusRjEhVDADir1xfDQqFxqoi7djDX5k81b")
	

//sjg 0109 add for check privkey to address:
func checkprikaddr(serializedKey string) error{
		//var keys map[string]*btcutil.WIF
	if serializedKey != "" {
		key := serializedKey
			wif, err := btcutil.DecodeWIF(key)
			if err != nil {
				return err
			}
		/*
			addr, err := btcutil.NewAddressPubKey(wif.SerializePubKey(),
				activeNet.Params)
		*/
		//sgj --6
		btchashAddr, pubKey, err := keyToAddr(wif.PrivKey, activeNet.Params)
	fmt.Printf("address's ECPubKey(),get PubKeyHash' Address is :%v \n", btchashAddr)
	//输出：私钥；公钥；公钥hash（地址）
			//5KD3sSntucZFRzDNUJusRjEhVDADir1xfDQqFxqoi7djDX5k81b
			fmt.Printf("==checkprikaddr(),,param serializedKey is :%s,wif.PrivKey is %v, btchashAddr is :%v,pubKey is :%v\n",serializedKey,wif.PrivKey,btchashAddr,pubKey)
	
			//2) times:
				addr, err := btcutil.NewAddressPubKey(wif.SerializePubKey(),
				activeNet.Params)
		fmt.Printf("==sgjwatching==exec signRawTransaction()--6666, generted addr is :%v,cmd PrivKeys is:%v\n",addr.EncodeAddress(),serializedKey)
			wif.CompressPubKey = true
			//2) times:
				addr, err = btcutil.NewAddressPubKey(wif.SerializePubKey(),
				activeNet.Params)
		fmt.Printf("==sgjwatching==exec signRawTransaction()-7777, generted addr is :%v,cmd PrivKeys is:%v\n",addr.EncodeAddress(),serializedKey)
					
	}
	return nil
}

func keyToAddr(key *btcec.PrivateKey, net *chaincfg.Params) (pubhash btcutil.Address, getpubKeyAddr *btcutil.AddressPubKey, err error) {

	serializedKey := key.PubKey().SerializeCompressed()
	pubKeyAddr, err := btcutil.NewAddressPubKey(serializedKey, net)
	if err != nil {
		return nil, pubKeyAddr,nil
	}
	//generate pubKeyAddr
	//sgj add for pubkey 11.27:
	fmt.Printf("bef keyToAddr,get cur pubKey is :%s \n", pubKeyAddr)

	//return pubKeyAddr.AddressPubKeyHash(), nil
	return pubKeyAddr.AddressPubKeyHash(), pubKeyAddr, nil
}
