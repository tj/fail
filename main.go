package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/apex/httplog"
	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

func init() {
	log.SetHandler(text.Default)
}

func main() {
	addr := flag.String("addr", ":3000", "Bind address")
	flag.Parse()

	err := http.ListenAndServe(*addr, httplog.New(http.HandlerFunc(serve)))
	if err != nil {
		log.Fatalf("error: %s", err)
	}
}

func serve(w http.ResponseWriter, r *http.Request) {
	errorRate := 0
	maxDelay := 0

	// max delay
	if s := r.URL.Query().Get("max_delay"); s != "" {
		if n, err := strconv.Atoi(s); err == nil {
			maxDelay = n
		} else {
			http.Error(w, "invalid delay parameter", http.StatusBadRequest)
		}
	}

	// error rate
	if s := r.URL.Query().Get("error_rate"); s != "" {
		if n, err := strconv.Atoi(s); err == nil {
			errorRate = n
		} else {
			http.Error(w, "invalid errors parameter", http.StatusBadRequest)
		}
	}

	// delay
	if maxDelay > 0 {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(maxDelay)))
	}

	// errors
	if errorRate > 0 {
		if rand.Intn(errorRate) == 0 {
			switch rand.Intn(2) {
			case 0:
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			case 1:
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}
	}

	fmt.Fprintf(w, "Hello World")
}
