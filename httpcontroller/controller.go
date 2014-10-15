package httpcontroller

import (
	"net/http"
	//"net/url"
	digest "digestalgo"
	"io/ioutil"
	"fmt"
	"urls"
)

type HttpController struct {
	DigestMessage *digest.DigestMessage
}

func NewHttpController() *HttpController {
	digestMessage := digest.NewDigestMessage()
	digestMessage.GetDigestHeader()
	return &HttpController {
		DigestMessage: digestMessage,
	}
}

func (hc * HttpController) addHeader(req *http.Request, key string) {
	req.Header.Add(key, hc.DigestMessage.AuthHeader[key])
}



func (hc *HttpController) Run() string {

	client := &http.Client{}
	
	aepUrl := urls.WorldUrl
	//aepUrl = "http://www.crafet.com"
	request, _ := http.NewRequest("GET", aepUrl, nil)
	//hc.addHeader(request, "Authorization")
	//hc.addHeader(request, "X-WSSE")
	
	validAuthorization := hc.DigestMessage.AuthHeader["Authorization"]
	validXWSSE := hc.DigestMessage.AuthHeader["X-WSSE"]
	
	fmt.Println("validAuthorization", validAuthorization)
	fmt.Println("validWSSE", validXWSSE)
	
	//request.Header.Add("Authorization", validAuthorization)
	//request.Header.Add("X-WSSE", validXWSSE)
	request.Header["Authorization"] = []string {validAuthorization}
	request.Header["X-WSSE"] = []string{validXWSSE}
	//request.Header.Write(fmt.Sprintf("X-WSSE:%s", validXWSSE))
	
	//request.WriteHeader("X-WSSE", validXWSSE)
	//request("Content-Type", "text/xml;charset=UTF-8")
	
	response, err := client.Do(request)
	fmt.Println("response error:", err)
	defer response.Body.Close()	
	fmt.Println("statuscode: ", response.StatusCode)
	
	bodyStr := ""
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		bodyStr = string(body)
		fmt.Println(bodyStr)
	}
	
	fmt.Println("request: ", request)
	
	return bodyStr
}

func (hc *HttpController) Run2() {
	client := &http.Client{}
	
	aepUrl := urls.WorldUrl
	
	//validAuthorization := hc.DigestMessage.AuthHeader["Authorization"]
	//validXWSSE := hc.DigestMessage.AuthHeader["X-WSSE"]
	
	/*
	req := &http.Request {
		Method: "GET",
		Proto: "HTTP/1.1",
		Header: http.Header{
			"Authorization": []string {validAuthorization},
			"X-WSSE:": []string {validXWSSE},
		},
		URL: &url.URL {
			Scheme:"http",
			Host:"api.sdpaep.com/sinanews/world/v1",
			Path:"/",
		},
	}
	*/
	req, _ := http.NewRequest("GET", aepUrl, nil)

	response, err := client.Do(req)
	fmt.Println("statusCode:", response.StatusCode)
	fmt.Println("err: ", err)
}

