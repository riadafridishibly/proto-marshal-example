package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/riadafridishibly/proto-marshal-example/protoexample"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func addressbookExample() {
	person1 := &protoexample.Person{
		Name:  "Jon Doe",
		Id:    32,
		Email: "jondoe@example.com",
		Phones: []*protoexample.Person_PhoneNumber{
			{
				Number: "0123456",
				Type:   protoexample.Person_HOME,
			},
			{
				Number: "8976543",
				Type:   protoexample.Person_WORK,
			},
		},
		LastUpdated: timestamppb.New(time.Now()),
	}
	person2 := &protoexample.Person{
		Name:  "Susan",
		Id:    33,
		Email: "susan@example.com",
		Phones: []*protoexample.Person_PhoneNumber{
			{
				Number: "012345678",
				Type:   protoexample.Person_HOME,
			},
			{
				Number: "897654301",
				Type:   protoexample.Person_WORK,
			},
		},
		LastUpdated: timestamppb.New(time.Now().Add(20 * 24 * time.Hour)),
	}
	addrs := &protoexample.AddressBook{
		People: []*protoexample.Person{
			person1,
			person2,
		},
	}
	data := marshallDump("AddressBook", addrs)
	var v protoexample.AddressBook
	unmarshal("AddressBook", data, &v)
}

func marshallDump(name string, m proto.Message) []byte {
	data, err := proto.Marshal(m)
	if err != nil {
		log.Fatal("Marshal failed:", err)
	}
	dump := hex.Dump(data)
	log.Println("Proto Dump:", name)
	fmt.Print(dump)
	return data
}

func unmarshal(name string, data []byte, target proto.Message) {
	log.Printf("Unmershalling %s...", name)
	err := proto.Unmarshal(data, target)
	if err != nil {
		log.Fatalf("Faile: %s: err: %v", name, err)
	}
	jsonData, err := json.MarshalIndent(target, "", "  ")
	if err != nil {
		log.Fatal("Failed to marshal:", err)
	}

	log.Println("Result:\n", string(jsonData))
}

func rectExample() {
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

	data := marshallDump("Rect", rect)

	var r protoexample.Rectangle
	unmarshal("Rect", data, &r)
}

func main() {
	rectExample()
	addressbookExample()
}
