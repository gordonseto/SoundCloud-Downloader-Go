package main

import (
	//"github.com/AlexJuca/soundcloud-go"
	//"fmt"
	//"net/http"
	//"io"
	//"os"
	"fmt"
	"net/http"
	"io/ioutil"
	"io"
	"os"
)

const CLIENT_ID = "175c043157ffae2c6d5fed16c3d95a4c"	// taken from someone elses project lol
const SECRET_KEY = "99a51990bd81b6a82c901d4cc6828e46"
const AGGRESSIVE_CLIENT_ID = "fDoItMDbsbZz8dY16ZzARCZmzgHBPotA"
const APP_VERSION = "1481046241"

const hard_track_url = "https://cf-media.sndcdn.com/fMITEgMn41U8.128.mp3?Policy=eyJTdGF0ZW1lbnQiOlt7IlJlc291cmNlIjoiKjovL2NmLW1lZGlhLnNuZGNkbi5jb20vZk1JVEVnTW40MVU4LjEyOC5tcDMiLCJDb25kaXRpb24iOnsiRGF0ZUxlc3NUaGFuIjp7IkFXUzpFcG9jaFRpbWUiOjE0OTk0NDk4NDd9fX1dfQ__\u0026Signature=HxBiQUi1fIpsSivvc9F715v2a0tuKg1DFrOqQhGrCHKoD2Hnt1vMjVjWcPPsNXswPUePfEy1sz7jQ~WXFHtpbjUUK4TGZiVOq5LS6pGe7T7krEana0N6wK98Gf8yQ5VZh5oap89XSyDCnptk1moMo-wXsjZ~aXe~aUZg282jxiCMiI7~jI1Sgyqe39F6nB88VgvQNcoMjAsFHBHhtBOhFuFiZ30ar6pHsaFS1me4V0t1RUbniCY7voGILcfYAfkJHsoMrueVXLx3tpfa4hmsLh2rZ1e9BS5QvNEMqJsnxa-tmZdaik30DURwLLfODcb2tc576wMaC6akMhcbwkXiHw__\u0026Key-Pair-Id=APKAJAGZ7VMH2PFPW6UQ"

func main(){
	resp, err := http.Get("https://api.soundcloud.com/i1/tracks/278678093/streams/?client_id=175c043157ffae2c6d5fed16c3d95a4c&app_version=1481046241")
	handleError(err)

	responseData, err := ioutil.ReadAll(resp.Body)
	handleError(err)

	fmt.Println(string(responseData))

	//client := soundclient.SoundCloud{ClientId: CLIENT_ID, ClientSecret: SECRET_KEY}
	//song := client.Tracks("13158")
	//
	//streamUrl, _ := song.GetString("stream_url")
	//waveformUrl, _ := song.GetString("waveform_url")
	//
	out, err := os.Create("output.mp3")
	handleError(err)
	defer out.Close()

	resp, err = http.Get(getHardTrackUrl())
	handleError(err)
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	handleError(err)

	//fmt.Println(streamUrl)
	//fmt.Println(waveformUrl)

	getTrackId("")
}

func getTrackId(url string){
	resp, err := http.Get("https://api.soundcloud.com/resolve.json?url=https%3A%2F%2Fsoundcloud.com%2Fmanilakilla%2Flive-hard-summer-2016&client_id=175c043157ffae2c6d5fed16c3d95a4c")
	responseData, err := ioutil.ReadAll(resp.Body)
	handleError(err)

	fmt.Println(string(responseData))
}

func getHardTrackUrl() string {
	return "https://cf-media.sndcdn.com/j9JO1HF2aEgX.128.mp3?Policy=eyJTdGF0ZW1lbnQiOlt7IlJlc291cmNlIjoiKjovL2NmLW1lZGlhLnNuZGNkbi5jb20vajlKTzFIRjJhRWdYLjEyOC5tcDMiLCJDb25kaXRpb24iOnsiRGF0ZUxlc3NUaGFuIjp7IkFXUzpFcG9jaFRpbWUiOjE0OTk0NTY4MTl9fX1dfQ__\u0026Signature=k0RaCM7CS8OrTPheJyykTyHQJJKK~97G1erhFNwuAe6QK1H5-TVYqK7bVw44BGBNoUSBQE1Ci2TTti3RXjqQKuvDpJpbGBpMRV6BSiD9qEq2cErXaKAxPA0OsCwbQJsc3FGiEpxfcPL0IDz0QCcNJJZ5Au8vUD8YKCmgC20RU4-1Ky~7dSQoeSd3ZwdxaP846zQwm~cdrYUJRdrc5KkgY2vpl1xwG3qBIQBxILHV9f1bUwiVu5FQZIq1cuXd17ROrbDLfz5I7szGO4rPmP1hNghyOwIUaGs~OAkvLOpLxRuhcybRvtR8KUtJQ59skoDU5D0E~FGvsAq4fKGScqUFEQ__\u0026Key-Pair-Id=APKAJAGZ7VMH2PFPW6UQ"
}

func handleError(err error){
	if err != nil {
		fmt.Println(err)
	}
}