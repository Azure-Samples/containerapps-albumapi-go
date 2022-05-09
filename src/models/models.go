package models

import (
	"fmt"
)

func init() {
	fmt.Println("Initialize - models package")
}

type Album struct {
	Id     		int  		`json:"id"`
	Title  		string  	`json:"title"`
	Artist 		string  	`json:"artist"`
	Price  		float64 	`json:"price"`
	Image_URL  	string 		`json:"image_url"`
}

func GetAlbums() []Album {

	var albums []Album

	album1 := Album{Id: 1, Title: "Sgt Peppers Lonely Hearts Club Band", Artist: "The Beatles", Price: 56.99, Image_URL: "https://www.listchallenges.com/f/items/f3b05a20-31ae-4fdf-aebd-d1515af54eea.jpg"}
	albums = append(albums, album1)

	album2 := Album{Id: 2, Title: "Pet Sounds", Artist: "The Beach Boys", Price: 17.99, Image_URL: "https://www.listchallenges.com/f/items/fdef1440-e979-455a-90a7-14e77fac79af.jpg"}
	albums = append(albums, album2)

	album3 := Album{Id: 3, Title: "The Beatles: Revolver", Artist: "The Beatles", Price: 39.99, Image_URL: "https://www.listchallenges.com/f/items/19ff786d-d7a4-4fdc-bee2-59db8b33370d.jpg"}
	albums = append(albums, album3)

	album4 := Album{Id: 4, Title: "Highway 61 Revisited", Artist: "Bob Dylan", Price: 39.99, Image_URL: "https://www.listchallenges.com/f/items/428cf087-6c24-45ad-bf8c-bd3c57ddf444.jpg"}
	albums = append(albums, album4)

	album5 := Album{Id: 5, Title: "Rubber Soul", Artist: "The Beatles", Price: 39.99, Image_URL: "https://www.listchallenges.com/f/items/ebc794ef-1491-4672-a007-0081edc1a8ae.jpg"}
	albums = append(albums, album5)

	album6 := Album{Id: 6, Title: "What's Going On", Artist: "Marvin Gaye", Price: 39.99, Image_URL: "https://www.listchallenges.com/f/items/e5250d6c-1c15-4617-a943-b27e87e21704.jpg"}
	albums = append(albums, album6)

	return albums

};