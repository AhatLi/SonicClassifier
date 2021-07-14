package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	requester, err := NewSonicRequester()
	if err != nil {
		fmt.Println("ERROR : ", err)
		os.Exit(0)
	}

	result := requester.GetPlaylists()

	entry := make([]Entry, 0)
	pid := ""

	for _, item := range result.SubsonicResponse.Playlists.Playlist {
		if item.Name == requester.conf.playlist {
			pid = item.Id
			list := requester.GetPlaylist(item.Id)
			for _, pitem := range list.SubsonicResponse.Playlist.Entry {
				entry = append(entry, pitem)
			}
		}
	}

	if requester.conf.item == "path" {
		sort.Slice(entry, func(i, j int) bool {
			return entry[i].Path < entry[j].Path
		})
	} else if requester.conf.item == "album" {
		sort.Slice(entry, func(i, j int) bool {
			return entry[i].Path < entry[j].Path
		})
	}

	for _, e := range entry {
		requester.UpdatePlaylist(pid, e.Id)

		fmt.Println(e.Path)
	}
}
