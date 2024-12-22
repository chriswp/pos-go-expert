package main

import (
	"net/http"
	"strconv"
	"sync/atomic"
)

var number uint64 = 0

func main() {
	//m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//m.Lock()
		//number++
		atomic.AddUint64(&number, 1)
		//m.Unlock()
		_, err := w.Write([]byte("Você é o visitante número: " + strconv.FormatUint(number, 10) + "\n"))
		if err != nil {
			return
		}
	})
	http.ListenAndServe(":3000", nil)
}
