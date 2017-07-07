package main

import (
	"github.com/AlexJuca/soundcloud-go"
	"fmt"
)

const CLIENT_ID = "175c043157ffae2c6d5fed16c3d95a4c"	// taken from someone elses project lol
const SECRET_KEY = "99a51990bd81b6a82c901d4cc6828e46"

func main(){
	client := soundclient.SoundCloud{ClientId: CLIENT_ID, ClientSecret: SECRET_KEY}
	song := client.Tracks("13158")

	title, _ := song.GetString("title")
	description, _ := song.GetString("description")

	fmt.Println("Title ->", title)
	fmt.Println("Description ->", description)
}