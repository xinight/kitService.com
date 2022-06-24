package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	khttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	endp "kittest.com/endpoint"
	"kittest.com/services"
	mtsp "kittest.com/transport"
	"kittest.com/util"
)

var c = make(chan os.Signal)

func main() {
	r := mux.NewRouter()
	r.Handle("/activity/doublegift/info", khttp.NewServer(endp.GenGetInfoEndpoint(), mtsp.GetInfoDecodeRequest, mtsp.GetInfoEncodeResponse))
	r.HandleFunc("/activity/doublegift/health", func(w http.ResponseWriter, r *http.Request) {
		s := services.DoubleGIftService{}
		w.Header().Set("Content-type", "application/json")
		if s.CheckHealth() {
			w.Write([]byte(`{"status":true}`))
		} else {
			w.Write([]byte(`{"status":falae}`))
		}

	})
	errChan := make(chan error)
	go func() {
		util.RegistService()
		err := http.ListenAndServe(":"+strconv.Itoa(*util.ServicePort), r)
		if err != nil {
			log.Println(err)
			errChan <- err
		}
	}()

	go func() {
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	e := <-errChan
	util.DeregistService()
	fmt.Println(e)
}
