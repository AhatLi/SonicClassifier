package main

import (
	"fmt"
	"net/http"
)

func indexTandler(w http.ResponseWriter, r *http.Request) {
	//전달할 http 내용
	fmt.Fprintf(w, "123456")
}

func webMain() {
	//웹서버 URL 리퀘스트 받기
	http.HandleFunc("/test", indexTandler)

	//파일서버
	http.Handle("/", http.StripPrefix("/free", http.FileServer(http.Dir("public"))))

	//서버 실행
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}

}
