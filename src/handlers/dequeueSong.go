package handlers

import (
	"encoding/json"
	"lalu-storage/helpers"

	"github.com/streadway/amqp"
)

type SongData struct {
	Name string `json:"name"`
	Data map[string] []byte `json:"data"`
}

func DequeueSongs (channel *amqp.Channel) {
	queue, err := channel.QueueDeclare(
		"songs", // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	if err != nil {
		panic(err.Error())
	}

	songs, err := channel.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic("Error consuming msgs")
	}

	forever := make(chan bool)

	go func (){
		for song := range songs {

			var songData SongData

			json.Unmarshal(song.Body, &songData)
			helpers.Uploader.UploadSongFromQueue(songData.Data["data"], songData.Name)
		}
	}()

	<-forever
}