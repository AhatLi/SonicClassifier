package main

import (
	"fmt"
	"net/http"
	"regexp"
)

func getPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	requester, err := NewSonicRequester()
	if err != nil {
		fmt.Println("ERROR : ", err)
	}

	result := requester.GetPlaylists()
	playlists := ""
	for _, item := range result.SubsonicResponse.Playlists.Playlist {
		playlists += "," + item.Name
	}
	if len(playlists) != 0 {
		playlists = playlists[1:]
	}

	fmt.Fprintf(w, playlists)
}

func sortHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "2")
}

func webMain() {
	http.HandleFunc("/getplaylist", getPlaylistHandler)
	http.HandleFunc("/sort", sortHandler)

	fileServer := http.FileServer(http.Dir("public"))
	fileMatcher := regexp.MustCompile(`\.[a-zA-Z]*$`)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !fileMatcher.MatchString(r.URL.Path) {
			http.ServeFile(w, r, "public/index.html")
		} else {
			fileServer.ServeHTTP(w, r)
		}
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
