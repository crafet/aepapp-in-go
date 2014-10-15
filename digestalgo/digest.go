package digestalgo

import (
	"accessmodel"
	"crypto/sha1"
	"encoding/base64"

	"fmt"
)

const (  
    base64Table = "123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"  
) 

type DigestHeader interface {
	GetDigestHeader() string
}

type DigestMessage struct {
	AuthHeader map[string]string
}

var am = accessmodel.NewAccessorModel()

func NewDigestMessage() *DigestMessage {
	authHeader := make(map[string]string)
	
	//一旦要换行，就需要每一行加上逗号
	return &DigestMessage{
		AuthHeader: authHeader,
		}
}

func (dm * DigestMessage) GetPasswordDigest(nonce string, createdTime string) string{
	//

	//nonce := am.GetNonce()
	//createdTime := am.GetCreatedTime()
	
	//nonce = "voXouXzOmYZ4PjFGj0g0DDfjL4fYPnqh1JjO"
	//createdTime = "2014-10-12T13:53:44Z"
	aggre := nonce + createdTime + am.AppSecret
	fmt.Println("aggre: ", aggre)
	
	shaAlgo := sha1.New()
	shaAlgo.Write([]byte(aggre))
	aggreBytes := shaAlgo.Sum(nil)
	
	fmt.Println(string(aggreBytes))
	
	//base64Str := base64.NewEncoding(base64Table).EncodeToString(aggreBytes)
	// we use default base64 encoding algorithm
	base64Str := base64.StdEncoding.EncodeToString(aggreBytes)
	fmt.Println("base64Str: ", base64Str)
	return base64Str
}

func(dm * DigestMessage) GetDigestHeader() {
	// Authorization:WSSE realm="SDP", profile="UsernameToken", type="Appkey"
	// X-WSSE:UsernameToken Username="9e45190b41bf4b55b534ba4049b3bd39", PasswordDigest="Qh3t30xxril8F8q8Vu0Xr9opp6M=", Nonce="voXouXzOmYZ4PjFGj0g0DDfjL4fYPnqh1JjO", Created="2014-10-12T13:53:44Z"
	
	dm.AuthHeader["Authorization"] = "WSSE realm=\"SDP\", profile=\"UsernameToken\", type=\"Appkey\""
	nonce := am.GetNonce()
	createdTime := am.GetCreatedTime()
	formatWSSEStr := fmt.Sprintf(
					"UsernameToken Username=\"%s\", PasswordDigest=\"%s\", Nonce=\"%s\", Created=\"%s\"",
					am.AppKey, 
					dm.GetPasswordDigest(nonce, createdTime),
					nonce,
					createdTime)
	dm.AuthHeader["X-WSSE"] = formatWSSEStr
	
	fmt.Println("Authorization: ", dm.AuthHeader["Authorization"])
	fmt.Println("X-WSSE: ", dm.AuthHeader["X-WSSE"])
}