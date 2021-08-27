package main

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
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
	json.Unmarshal(body, &result)

	return result
}

//크롬 기준으로 3000개의 음악에 대한 정렬이 실패했다...
//그래도 500개 까지는 문제 없을것 같아 500개씩 모아서 처리하도록 수정한다.
func (r *SonicRequester) UpdateStar(entry []Entry) {
	sid := ""
	for i, e := range entry {
		sid += "&id=" + e.Id

		if i%500 == 0 {
			urlUnStar := r.conf.sonic_url + "/rest/unstar?u=" + r.conf.username + "&t=" + string(r.token[:]) + "&s=" + r.salt + "&v=1.15.0&c=" + r.client +
				"&f=json" + sid
			getMessage(urlUnStar)

			urlStar := r.conf.sonic_url + "/rest/star?u=" + r.conf.username + "&t=" + string(r.token[:]) + "&s=" + r.salt + "&v=1.15.0&c=" + r.client +
				"&f=json" + sid
			getMessage(urlStar)

			sid = ""
		}
	}
	if len(sid) != 0 {
		urlUnStar := r.conf.sonic_url + "/rest/unstar?u=" + r.conf.username + "&t=" + string(r.token[:]) + "&s=" + r.salt + "&v=1.15.0&c=" + r.client +
			"&f=json" + sid
		getMessage(urlUnStar)

		urlStar := r.conf.sonic_url + "/rest/star?u=" + r.conf.username + "&t=" + string(r.token[:]) + "&s=" + r.salt + "&v=1.15.0&c=" + r.client +
			"&f=json" + sid
		getMessage(urlStar)
	}
}

func (r *SonicRequester) UpdatePlaylist(pid string, entry []Entry) {
	remove := ""
	add := ""
	//remove
	for i, _ := range entry {
		remove += "&songIndexToRemove=" + strconv.Itoa(i)
		if i%500 == 0 {
			removeURL := r.conf.sonic_url + "/rest/updatePlaylist?u=" + r.conf.username + "&t=" + string(r.token[:]) + "&s=" + r.salt + "&v=1.15.0&c=" + r.client +
				"&f=json&playlistId=" + pid + remove
			getMessage(removeURL)
			remove = ""
		}
	}

	if len(remove) != 0 {
		removeURL := r.conf.sonic_url + "/rest/updatePlaylist?u=" + r.conf.username + "&t=" + string(r.token[:]) + "&s=" + r.salt + "&v=1.15.0&c=" + r.client +
			"&f=json&playlistId=" + pid + remove
		getMessage(removeURL)
	}

	//add
	for i, e := range entry {
		add += "&songIdToAdd=" + e.Id
		if i%500 == 0 {
			addURL := r.conf.sonic_url + "/rest/updatePlaylist?u=" + r.conf.username + "&t=" + string(r.token[:]) + "&s=" + r.salt + "&v=1.15.0&c=" + r.client +
				"&f=json&playlistId=" + pid + add
			getMessage(addURL)
			add = ""
		}
	}

	if len(remove) != 0 && len(add) != 0 {
		addURL := r.conf.sonic_url + "/rest/updatePlaylist?u=" + r.conf.username + "&t=" + string(r.token[:]) + "&s=" + r.salt + "&v=1.15.0&c=" + r.client +
			"&f=json&playlistId=" + pid + add
		getMessage(addURL)
	}
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
