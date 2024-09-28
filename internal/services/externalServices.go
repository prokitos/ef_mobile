package services

import (
	"encoding/json"
	"io/ioutil"
	"mymod/internal/models/responses"
	"mymod/internal/models/tables"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

var SongExternalAddress string

// получение данных из внешнего сервера, и разделение текста на части.
// если проблема на сервере, то не будем брать новые данные вообще
func EnrichtSong(song tables.Song) (tables.Song, error) {

	tempSong, err := sendRequestToGet(song.Group, song.Song)
	if err != nil {
		return song, err
	}

	allText := strings.Split(tempSong.Text, "\n\n")
	for iter, item := range allText {
		var temp tables.Verse
		temp.VerseId = iter + 1
		temp.VerseText = item
		song.Text = append(song.Text, temp)
	}

	song.ReleaseDate = tempSong.ReleaseDate
	song.Link = tempSong.Link

	return song, nil
}

// получение данных из внешнего сервера по роуту /info
func sendRequestToGet(group string, song string) (tables.ExternalSong, error) {

	baseURL, _ := url.Parse(SongExternalAddress + "/info")
	params := url.Values{}
	params.Add("group", group)
	params.Add("song", song)
	baseURL.RawQuery = params.Encode()

	// типо контекст. возврат с ошибкой если не успел сделать запрос за 2 секунду.
	client := http.Client{
		Timeout: 2 * time.Second,
	}

	resp, err := client.Get(baseURL.String())
	if err != nil {

		if err, ok := err.(net.Error); ok && err.Timeout() {
			log.Error("timeout request !!")
			return tables.ExternalSong{}, responses.ResponseBase{}.BaseExternalError()
		}

		log.Debug("Error connecting to external api")
		log.Error("Error getting data from api")
		return tables.ExternalSong{}, responses.ResponseBase{}.BaseExternalError()
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var carExemp tables.ExternalSong
	json.Unmarshal(body, &carExemp)
	return carExemp, nil
}
