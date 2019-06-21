package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	fListen = flag.String("listen", "localhost:8000", "listen address")
)

func fatalIfError(err error, msg string) {
	if err != nil {
		log.Fatal("error ", msg, ": ", err)
	}
}

func Dump(w http.ResponseWriter, r *http.Request) {
	io.Copy(os.Stdout, r.Body)
	fmt.Fprintln(os.Stdout)

	w.WriteHeader(http.StatusOK)
}

func main() {
	flag.Parse()

	http.HandleFunc("/", Dump)
	log.Print("starting listener on ", *fListen)
	fatalIfError(http.ListenAndServe(*fListen, nil), "listening")
}
