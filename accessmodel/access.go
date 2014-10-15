package accessmodel

import (
	"math/rand"
	"time"
	"fmt"
)

const (
	charSets string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	appKey = "9e45190b41bf4b55b534ba4049b3bd39"
	appSecret = "79c242600fdf0e6c"
	nonceLength = 36
)

type AccessorModel struct {
	CharSets string
	AppKey string
	AppSecret string
}

func NewAccessorModel() *AccessorModel {
	return &AccessorModel {
		CharSets: charSets,
		AppKey: appKey,
		AppSecret: appSecret,
	}
}

func(am *AccessorModel) GetNonce() string{
	rand.Seed(time.Now().UnixNano())
	
	charSetsSize := len(am.CharSets)
	//fmt.Println("charsetsize: ", charSetsSize)
	iter := 0
	// 
	//nonce := make([]byte, nonceLength)
	var nonce []byte
	for {
		idx := rand.Intn(charSetsSize)
		//fmt.Println("idx:", idx)
		nonce = append(nonce, am.CharSets[idx])
		
		//fmt.Println(am.CharSets[idx])
		iter++
		if (iter == nonceLength) {
			break
		}
	}

	//fmt.Println(string(nonce[:nonceLength]))
	return string(nonce[:nonceLength])
}

func (am *AccessorModel) GetCreatedTime() string {
	//fmt.Println(time.Now().String())
	year, month, day := time.Now().UTC().Date()
	hour, min, sec := time.Now().UTC().Clock()
	
	utcTimeStr := fmt.Sprintf("%d-%d-%dT%02d:%02d:%02dZ", year, month, day, hour, min, sec)
	fmt.Println("utcTimeStr: ", utcTimeStr)
	return utcTimeStr
}