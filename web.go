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

	//파일서버 이거 어떻게 못하나 이렇게 해야 에러가 안나네
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("public"))))
	http.Handle("/main/default", http.StripPrefix("/main/default", http.FileServer(http.Dir("public"))))
	http.Handle("/static", http.StripPrefix("/static", http.FileServer(http.Dir("public"))))

	//서버 실행
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}

}
