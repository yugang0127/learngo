package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/scheduler"
	"learngo/crawler/zhengai/parser"
)

func main() {

	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url: "http://www.zhenai.com/zhenghun/",
	//	ParserFunc: parser.ParseCityList,
	//})

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}

	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun/",
		ParserFunc: parser.ParseCityList,
	})

	//contents, err := fetcher.Fetch("http://album.zhenai.com/u/110538321")
	//if err != nil {
	//	panic(err)
	//}
	//
	//parser.ParseProfile(contents)
}

