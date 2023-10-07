package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/riadafridishibly/proto-marshal-example/protoexample"
	"google.golang.org/protobuf/proto"
)

func main() {
	rect := &protoexample.Rectangle{
		Lo: &protoexample.Point{
			Latitude:  1,
			Longitude: 2,
		},
		Hi: &protoexample.Point{
			Latitude:  12,
			Longitude: 24,
		},
	}

	data, err := proto.Marshal(rect)
	if err != nil {
		log.Fatal("Marshal failed:", err)
	}

	dump := hex.Dump(data)
	fmt.Println("Proto Dump:")
	fmt.Println(dump)

	var r protoexample.Rectangle
	err = proto.Unmarshal(data, &r)
	if err != nil {
		log.Fatal("Unmarshal failed:", err)
	}

	fmt.Println("Unmarshalled:", r.String())
}
