package main

//gittest

import (
	"encoding/json"
	"fmt"
	"log"
)

type Proto struct {
	Name        string
	RepoURL     string
	ID          int64
	MessageList []*Message
	ServiceList []*Service
}

type Message struct {
	Name      string
	FieldList []*Field
}

type Field struct {
	Name     string
	Type     string
	ID       int64
	Repeated bool
}

type Service struct {
	Name string
	Rpcs []*Rpc
}

type Rpc struct {
	Name       string
	Req        string
	Res        string
	HttpOption bool
	Path       string
	HttpMethod string
}

func main() {
	data := `{
		"Name":"datatest",
		"RepoURL":"github.com/fish-repo",
		"ID":27,
		"MessageList":[
			{
				"Name":"fishmessageone",
				"FieldList":[
					{
						"Name":"fishfieldone",
						"Type":"string",
						"ID":1,
						"Repeated":false
					},
					{
						"Name":"fishfieldtwo",
						"Type":"int64",
						"ID":2,
						"Repeated":true
					}
				]
			},
			{
				"Name":"fishmessagetwo",
				"FieldList":[
					{
						"Name":"fishfieldthree",
						"Type":"bool",
						"ID":1,
						"Repeated":true
					},
					{
						"Name":"fishfieldfour",
						"Type":"fishtype",
						"ID":2,
						"Repeated":false
					}
				]
			}
		],
		"ServiceList":[
			{
				"Name":"Serviceone",
				"Rpcs":[
					{
						"Name":"fisheat",
						"Req":"fishmessageone",
						"Res":"fishmessagetwo",
						"HttpOption":true,
						"Path":"/eat",
						"HttpMethod":"post"
					},
					{
						"Name":"fishsleep",
						"Req":"fishmessagetwo",
						"Res":"fishmessageone",
						"HttpOption":false,
						"Path":"/eat",
						"HttpMethod":"post"
					}
				]
			},
			{
				"Name":"egg",
				"Rpcs":[
					{
						"Name":"eggeat",
						"Req":"fishmessageone",
						"Res":"fishmessageone",
						"HttpOption":false,
						"Path":"/eat",
						"HttpMethod":"post"
					},
					{
						"Name":"eggsleep",
						"Req":"fishmessagetwo",
						"Res":"fishmessagetwo",
						"HttpOption":true,
						"Path":"/sleep",
						"HttpMethod":"get"
					}
				]
			}
		]
	}`
	createProtoRequest := &Proto{}
	err := json.Unmarshal([]byte(data), createProtoRequest)
	if err != nil {
		log.Printf("[Create Proto]Request Json error: %#v", err)
	}
	fmt.Printf("%+v\n", createProtoRequest)
}
