package httpcontroller

import (
	"net/http"
	"fmt"
	"log"
)


type mymux struct {
}

func(m *mymux) doGet() string{
	hc := NewHttpController()
	return hc.Run()
}

func(m *mymux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	fmt.Println("current method should be get , it is[", method, "]")
	
	resp := m.doGet()
	fmt.Fprintf(w, resp)
}

func SetupServer(portArg string) {
	mux := &mymux{}
	
	fmt.Println("customed server is running @localhost", portArg, "...")
	//err := http.ListenAndServe(":9090", mux)
	err := http.ListenAndServe(portArg, mux)
	if err != nil {
		log.Fatal("ListenAndServe failed ", err)
	}
}