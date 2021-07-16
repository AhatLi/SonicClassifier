package main

type StarredList struct {
	SubsonicResponse struct {
		Starred struct {
			Entry []Entry `json:"song"`
		}
	} `json:"subsonic-response"`
}
