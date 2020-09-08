package main

import (
	"flag"
	"fmt"
	"json-diff/handler"
	"net/http"
	_ "net/http/pprof"
	"strconv"

	"github.com/gorilla/mux"
)

var port int

func init() {
	flag.IntVar(&port, "port", 8083, "指定服务监控的端口，例如 :8083 ，注意，需要带冒号")
}

func main() {
	r := mux.NewRouter()

	r.Handle("/req", handler.NewBaseHandler(handler.NewReqHandler(), nil, nil))

	fmt.Println("server start at :", port)
	http.Handle("/", r)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
