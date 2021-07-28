package main

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"

	"gopkg.in/ini.v1"
)

type Conf struct {
	username  string
	passwd    string
	sonic_url string
}

func (conf *Conf) initConf() error {
	cfg, err := ini.Load("SonicClassifier.conf")
	if err != nil {
		return err
	}

	conf.username = cfg.Section("account").Key("username").String()
	conf.passwd = cfg.Section("account").Key("passwd").String()
	conf.sonic_url = cfg.Section("network").Key("sonic_url").String()

	if conf.username == "" || conf.passwd == "" || conf.sonic_url == "" {
		return errors.New("check config")
	}

	return nil
}

type SonicRequester struct {
	conf   Conf
	client string
	salt   string
	token  string
}

func NewSonicRequester() *SonicRequester {
	requester := &SonicRequester{}
	err := requester.conf.initConf()
	if err != nil {
		return nil
	}

	requester.client = "CLI"
	requester.salt = RandomString(20)
	requester.token = fmt.Sprintf("%x", md5.Sum([]byte(requester.conf.passwd+requester.salt)))

	return requester
}

func (r *SonicRequester) CheckConnection() bool {
	url := r.conf.sonic_url + "/rest/ping?u=" + r.conf.username + "&t=" + string(r.token[:]) + "&s=" + r.salt + "&v=1.15.0&c=" + r.client + "&f=json"
	res, _ := getMessage(url)

	return strings.Contains(string(res), "\"ok\"")
}

func (r *SonicRequester) GetPlaylists() Playlists {
	url := r.conf.sonic_url + "/rest/getPlaylists?u=" + r.conf.username + "&t=" + string(r.token[:]) + "&s=" + r.salt + "&v=1.15.0&c=" + r.client + "&f=json"
	body, _ := getMessage(url)
	var result Playlists
	json.Unmarshal(body, &result)

	return result
}

func (r *SonicRequester) GetPlaylist(pid string) Playlist {
	url := r.conf.sonic_url + "/rest/getPlaylist?u=" + r.conf.username + "&t=" + string(r.token[:]) + "&s=" + r.salt + "&v=1.15.0&c=" + r.client + "&f=json&id=" + pid
	body, _ := getMessage(url)
	var result Playlist
	json.Unmarshal(body, &result)

	return result
}

func (r *SonicRequester) GetStarred() StarredList {
	url := r.conf.sonic_url + "/rest/getStarred?u=" + r.conf.username + "&t=" + string(r.token[:]) + "&s=" + r.salt + "&v=1.15.0&c=" + r.client + "&f=json"
	body, _ := getMessage(url)
	var result StarredList
	fmt.Println(string(body))
	json.Unmarshal(body, &result)

	return result
}

func (r *SonicRequester) UpdateStar(sid string) {
	urlUnStar := r.conf.sonic_url + "/rest/unstar?u=" + r.conf.username + "&t=" + string(r.token[:]) + "&s=" + r.salt + "&v=1.15.0&c=" + r.client +
		"&f=json&id=" + sid
	getMessage(urlUnStar)

	urlStar := r.conf.sonic_url + "/rest/star?u=" + r.conf.username + "&t=" + string(r.token[:]) + "&s=" + r.salt + "&v=1.15.0&c=" + r.client +
		"&f=json&id=" + sid
	getMessage(urlStar)
}

func (r *SonicRequester) UpdatePlaylist(pid string, sid string) {
	url := r.conf.sonic_url + "/rest/updatePlaylist?u=" + r.conf.username + "&t=" + string(r.token[:]) + "&s=" + r.salt + "&v=1.15.0&c=" + r.client +
		"&f=json&playlistId=" + pid + "&songIdToAdd=" + sid + "&songIndexToRemove=0"
	getMessage(url)
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func getMessage(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return body, nil
}
