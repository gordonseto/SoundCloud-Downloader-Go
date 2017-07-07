package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"io"
	"os"
	"encoding/json"
	url2 "net/url"
	//"github.com/AlexJuca/soundcloud-go"
	//"github.com/mikkyang/id3-go"
	"github.com/cavaliercoder/grab"
	"time"
)

const CLIENT_ID = "175c043157ffae2c6d5fed16c3d95a4c"
const SECRET_KEY = "99a51990bd81b6a82c901d4cc6828e46"
const AGGRESSIVE_CLIENT_ID = "fDoItMDbsbZz8dY16ZzARCZmzgHBPotA"
const APP_VERSION = "1481046241"

const TRACK_URL_KEY = "http_mp3_128_url"	// key to get the track url from api body

// based off of https://github.com/Miserlou/SoundScrape

func main(){
	args := os.Args
	if len(args) == 2 {
		url := args[1]
		if isValidURL(url) {
			downloadFromSoundCloud(url)
		} else {
			fmt.Println("Please enter a valid SoundCloud URL")
		}
	} else {
		fmt.Println("ERROR: Usage - 'go run scdownloader.go https://soundcloud.com/myurl'")
	}
}

func isValidURL(url string) bool {
	_, err := url2.ParseRequestURI(url)
	return err == nil
}

func downloadFromSoundCloud(url string){
	// get trackId from url
	trackID := getTrackID(url)
	// use trackId to get mp3_url
	trackURL := getTrackURL(trackID)

	// download track
	resp, err := downloadFileFrom(trackURL)
	handleError(err)

	// show download UI
	showDownloadProgress(resp)

	if err = resp.Err(); err != nil {
		fmt.Printf("Download failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Download finished!")
}

func showDownloadProgress(resp *grab.Response){
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("Transferred %v / %v bytes (%.2f%%)\n", resp.BytesComplete(), resp.Size, 100*resp.Progress())
		case <-resp.Done:
			break Loop
		}
	}
}

func getTrackID(url string) string {
	escapedURL := url2.QueryEscape(url)
	apiURL := fmt.Sprintf("https://api.soundcloud.com/resolve.json?url=%s&client_id=%s", escapedURL, CLIENT_ID)
	resp, err := http.Get(apiURL)
	handleError(err)
	body, err := ioutil.ReadAll(resp.Body)
	handleError(err)

	var response map[string]interface{}
	err = json.Unmarshal([]byte(body), &response)
	handleError(err)

	// convert id (float64) to string
	return fmt.Sprintf("%.f", response["id"].(float64))
}

func getTrackURL(trackId string) string {
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

func downloadFileFrom(url string) (*grab.Response, error) {
	client := grab.NewClient()
	req, err := grab.NewRequest("output.mp3", url)

	resp := client.Do(req)
	return resp, err
}

func saveFile(data io.ReadCloser, fileName string){
	out, err := os.Create(fileName)
	handleError(err)
	defer out.Close()

	_, err = io.Copy(out, data)
	handleError(err)
}

//func tagFile(trackID string, fileName string){
//	// get track info
//	client := soundclient.SoundCloud{ClientId: CLIENT_ID, ClientSecret: SECRET_KEY}
//	song := client.Tracks(trackID)
//
//	file, err := id3.Open(fileName)
//	handleError(err)
//	defer file.Close()
//
//	title, ok := song.GetString("title")
//	if ok == nil {
//		fmt.Println(title)
//	}
//}

func handleError(err error){
	if err != nil {
		panic(err)
	}
}