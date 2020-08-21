package main

import (
	"github.com/zeromsi/go_with_nick/handlers"
	"log"
	"net/http"
	"os"
	"time"
)

func main(){
	l:=log.New(os.Stdout,"product-api",log.LstdFlags)
	hh:= handlers.NewHello(l)
	gh:=handlers.NewGoodBye(l)
	sm:=http.NewServeMux()
	sm.Handle("/hello",hh)
	sm.Handle("/goodbye",gh)

	s:= &http.Server{
		Addr: ":8080",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}
	 s.ListenAndServe()
}
