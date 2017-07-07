package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"io"
	"os"
	"encoding/json"
)

const CLIENT_ID = "175c043157ffae2c6d5fed16c3d95a4c"
const SECRET_KEY = "99a51990bd81b6a82c901d4cc6828e46"
const AGGRESSIVE_CLIENT_ID = "fDoItMDbsbZz8dY16ZzARCZmzgHBPotA"
const APP_VERSION = "1481046241"

const TRACK_URL_KEY = "http_mp3_128_url"	// key to get the track url from api body

// based off of https://github.com/Miserlou/SoundScrape (thanks for the API keys)

func main(){
	// get trackId from url
	_ = getTrackId("")
	// use trackId to get mp3_url
	trackUrl := getTrackUrl("278678093")

	// download track with mp3url
	track := downloadFileFrom(trackUrl)
	defer track.Body.Close()

	// save track
	saveFile(track.Body, "output.mp3")
}

func getTrackId(url string) string {
	resp, err := http.Get("https://api.soundcloud.com/resolve.json?url=https%3A%2F%2Fsoundcloud.com%2Fmanilakilla%2Flive-hard-summer-2016&client_id=175c043157ffae2c6d5fed16c3d95a4c")
	handleError(err)
	body, err := ioutil.ReadAll(resp.Body)
	handleError(err)

	fmt.Println(string(body))

	var response map[string]interface{}
	err = json.Unmarshal([]byte(body), &response)
	handleError(err)

	// convert id (float64) to string
	return fmt.Sprintf("%f", response["id"].(float64))
}

func getTrackUrl(trackId string) string {
	// create url
	url := fmt.Sprintf("https://api.soundcloud.com/i1/tracks/%s/streams/?client_id=%s&app_version=%s", trackId, CLIENT_ID, APP_VERSION)

	// make request
	resp, err := http.Get(url)
	handleError(err)
	body, err := ioutil.ReadAll(resp.Body)
	handleError(err)

	// convert response body to map
	var response map[string]interface{}
	err = json.Unmarshal([]byte(body), &response)
	handleError(err)

	return response[TRACK_URL_KEY].(string)
}

func downloadFileFrom(url string) *http.Response {
	resp, err := http.Get(url)
	handleError(err)
	return resp
}

func saveFile(data io.ReadCloser, fileName string){
	out, err := os.Create(fileName)
	handleError(err)
	defer out.Close()

	_, err = io.Copy(out, data)
	handleError(err)
}

func handleError(err error){
	if err != nil {
		panic(err)
	}
}