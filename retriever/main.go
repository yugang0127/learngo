package main

import (
	"fmt"
	"learngo/retriever/mock"
	real2 "learngo/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,
		form map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
}
const url = "http://www.imooc.com"

func session(r RetrieverPoster) string {
	r.Post(url, map[string]string{
		"contents":	"Another fake content",
	})
	return r.Get(url)
}
func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url, map[string]string{
		"name": "CuMian",
		"birthday": "20180830",
	})
}

func inspect(r Retriever)  {
	fmt.Println("Inspecting.", r)
	fmt.Printf(">%T %v\n", r, r)
	fmt.Print("> Type Switch ")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real2.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}

func main() {
	var r, r2 Retriever
	mockRetriever := mock.Retriever{"This is a fake content."}
	r = &mockRetriever
	inspect(r)
	//fmt.Println(download(r))
	r2 = &real2.Retriever{UserAgent:"Mozilla5.0", Timeout:time.Minute}
	inspect(r2)

	// Type assertion
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println("MockRetriever contents:	", mockRetriever.Contents)
	} else {
		fmt.Println("not a mock Retriever.")
	}
	//fmt.Println(download(r))
	fmt.Println("Try a session.")
	fmt.Println(session(&mockRetriever))
}
