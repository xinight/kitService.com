package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	khttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	endp "kittest.com/endpoint"
	pb "kittest.com/pbs"
	"kittest.com/services"
	mtsp "kittest.com/transport"
	"kittest.com/util"
)

var c = make(chan os.Signal)
var errChan = make(chan error)

func main() {
	t := *util.ServiceType
	if t == "rpc" {
		//rpc
		rpcRun()
	} else {
		//http
		httpRun()
	}
	go func() {
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	fmt.Println("start")
	e := <-errChan
	util.DeregistService()
	fmt.Println(e)
}

func httpRun() {
	r := mux.NewRouter()

	r.Handle("/activity/doublegift/info", khttp.NewServer(endp.GenGetInfoEndpoint(), mtsp.GetInfoDecodeRequest, mtsp.EncodeResponse))
	r.Handle("/activity/doublegift/exchange", khttp.NewServer(endp.GenExchangeEndpoint(), mtsp.ExchangeDecodeRequest, mtsp.EncodeResponse))

	r.HandleFunc("/activity/doublegift/health", func(w http.ResponseWriter, r *http.Request) {
		s := services.DoubleGIftService{}
		w.Header().Set("Content-type", "application/json")
		if s.CheckHealth() {
			w.Write([]byte(`{"status":true}`))
		} else {
			w.Write([]byte(`{"status":falae}`))
		}
	})
	go func() {
		util.RegistService("http")
		err := http.ListenAndServe(":"+strconv.Itoa(*util.ServicePort), r)
		if err != nil {
			fmt.Println(err)
			util.DeregistService()
			os.Exit(0)
		}
	}()

}

func rpcRun() {
	s := grpc.NewServer()
	server := mtsp.NewDoubleGiftRpc()
	pb.RegisterGetInfoServer(s, server)
	pb.RegisterExchangeServer(s, server)
	go func() {
		util.RegistService("rpc")
		l, err := net.Listen("tcp", ":"+strconv.Itoa(*util.ServicePort))
		if err != nil {
			fmt.Println(err)
			util.DeregistService()
			os.Exit(0)
		}
		s.Serve(l)
	}()

}
