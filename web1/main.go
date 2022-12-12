package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { // 인스턴스 복사할 필요 없기 때문에 포인터 Receiver로 등록
	fmt.Fprint(w, "Hello foo!")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world!")
	})

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello bar!")
	})

	http.Handle("/foo", &fooHandler{}) // ServeHTTP 메서드를 가지는 인스턴스(Handler interface)

	http.ListenAndServe(":3000", nil)
}
