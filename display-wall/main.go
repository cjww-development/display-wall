package main

import (
	"display-wall/web/routes/chat"
	"log"
	"net/http"
)

type Result struct {
	Id   string `bson:"id"`
	Name string `bson:"name"`
	Code string `bson:"code"`
}

func main() {
	hub := chat.NewHub()
	go hub.Run()
	//http.HandleFunc("/ws", test.WsEndpoint)
	http.HandleFunc("/ws/chat", func(writer http.ResponseWriter, request *http.Request) {
		chat.ChatEndpoint(hub, writer, request)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
	//collection := mongo.GetCollection("display-wall", "rooms")
	//ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	//resultDoc := &Result{}
	//err := collection.FindOne(ctx, bson.M{"id" : "testRoomId"}).Decode(resultDoc)
	//if err != nil { fmt.Println(err) }
	//fmt.Println(resultDoc)
}
