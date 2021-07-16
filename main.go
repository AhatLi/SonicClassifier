package main

import (
	"fmt"
	"net/http"
	"regexp"
	"sort"
)

func getPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	requester, err := NewSonicRequester()
	if err != nil {
		fmt.Println("ERROR : ", err)
		return
	}

	result := requester.GetPlaylists()
	playlists := ""
	for _, item := range result.SubsonicResponse.Playlists.Playlist {
		playlists += "|" + item.Name
	}
	if len(playlists) != 0 {
		playlists = playlists[1:]
	}

	fmt.Fprintf(w, playlists)
}

func sortPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	pItem := q.Get("item")
	pPlaylist := q.Get("playlist")
	pOrder := q.Get("order")

	if pItem == "" {
		pItem = "path"
	}
	if pPlaylist == "" {
		fmt.Println("ERROR : playlist error")
		fmt.Fprintf(w, "Fail")
		return
	}

	requester, err := NewSonicRequester()
	if err != nil {
		fmt.Println("ERROR : ", err)
		fmt.Fprintf(w, "Fail")
		return
	}

	result := requester.GetPlaylists()

	entry := make([]Entry, 0)
	pid := ""

	for _, item := range result.SubsonicResponse.Playlists.Playlist {
		if item.Name == pPlaylist {
			pid = item.Id
			list := requester.GetPlaylist(item.Id)
			for _, pitem := range list.SubsonicResponse.Playlist.Entry {
				entry = append(entry, pitem)
			}
		}
	}

	if pOrder == "desc" {
		if pItem == "path" {
			sort.Slice(entry, func(i, j int) bool {
				return entry[i].Path > entry[j].Path
			})
		} else if pItem == "title" {
			sort.Slice(entry, func(i, j int) bool {
				return entry[i].Title > entry[j].Title
			})
		} else if pItem == "album" {
			sort.Slice(entry, func(i, j int) bool {
				return entry[i].Path > entry[j].Path
			})
		} else if pItem == "artist" {
			sort.Slice(entry, func(i, j int) bool {
				return entry[i].Artist > entry[j].Artist
			})
		} else if pItem == "year" {
			sort.Slice(entry, func(i, j int) bool {
				return entry[i].Year > entry[j].Year
			})
		}
	} else {
		if pItem == "path" {
			sort.Slice(entry, func(i, j int) bool {
				return entry[i].Path < entry[j].Path
			})
		} else if pItem == "title" {
			sort.Slice(entry, func(i, j int) bool {
				return entry[i].Title < entry[j].Title
			})
		} else if pItem == "album" {
			sort.Slice(entry, func(i, j int) bool {
				return entry[i].Path < entry[j].Path
			})
		} else if pItem == "artist" {
			sort.Slice(entry, func(i, j int) bool {
				return entry[i].Artist < entry[j].Artist
			})
		} else if pItem == "year" {
			sort.Slice(entry, func(i, j int) bool {
				return entry[i].Year < entry[j].Year
			})
		}
	}

	for _, e := range entry {
		requester.UpdatePlaylist(pid, e.Id)
	}

	fmt.Fprintf(w, "OK")
}

func sortStarHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	//	pType := q.Get("type")
	pItem := q.Get("item")
	pOrder := q.Get("order")

	if pItem == "" {
		pItem = "path"
	}

	requester, err := NewSonicRequester()
	if err != nil {
		fmt.Println("ERROR : ", err)
		fmt.Fprintf(w, "Fail")
		return
	}

	result := requester.GetStarred()
	entry := make([]Entry, 0)

	for _, pitem := range result.SubsonicResponse.Starred.Entry {
		entry = append(entry, pitem)
	}

	if pOrder == "desc" {
		if pItem == "path" {
			sort.Slice(entry, func(i, j int) bool {
				return entry[i].Path < entry[j].Path
			})
		} else if pItem == "title" {
			sort.Slice(entry, func(i, j int) bool {
				return entry[i].Title < entry[j].Title
			})
		} else if pItem == "album" {
			sort.Slice(entry, func(i, j int) bool {
				return entry[i].Path < entry[j].Path
			})
		} else if pItem == "artist" {
			sort.Slice(entry, func(i, j int) bool {
				return entry[i].Artist < entry[j].Artist
			})
		} else if pItem == "year" {
			sort.Slice(entry, func(i, j int) bool {
				return entry[i].Year < entry[j].Year
			})
		}
	} else {
		if pItem == "path" {
			sort.Slice(entry, func(i, j int) bool {
				return entry[i].Path > entry[j].Path
			})
		} else if pItem == "title" {
			sort.Slice(entry, func(i, j int) bool {
				return entry[i].Title > entry[j].Title
			})
		} else if pItem == "album" {
			sort.Slice(entry, func(i, j int) bool {
				return entry[i].Path > entry[j].Path
			})
		} else if pItem == "artist" {
			sort.Slice(entry, func(i, j int) bool {
				return entry[i].Artist > entry[j].Artist
			})
		} else if pItem == "year" {
			sort.Slice(entry, func(i, j int) bool {
				return entry[i].Year > entry[j].Year
			})
		}
	}

	for _, e := range entry {
		requester.UpdateStar(e.Id)
	}

	fmt.Fprintf(w, "OK")
}

func main() {
	http.HandleFunc("/getPlaylist", getPlaylistHandler)
	http.HandleFunc("/sortPlaylist", sortPlaylistHandler)
	http.HandleFunc("/sortStar", sortStarHandler)

	fileServer := http.FileServer(http.Dir("public"))
	fileMatcher := regexp.MustCompile(`\.[a-zA-Z]*$`)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !fileMatcher.MatchString(r.URL.Path) {
			http.ServeFile(w, r, "public/index.html")
		} else {
			fileServer.ServeHTTP(w, r)
		}
	})
	err := http.ListenAndServe(":9255", nil)
	if err != nil {
		fmt.Println(err)
	}
}
