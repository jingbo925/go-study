package main

import (
	"fmt"
	"retriever/mock"
	"retriever/real"
)

type Retriever interface {
	Get(url string) string
}
type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "http://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post("http://www.imooc.com", map[string]string{
		"name":   "ccmouse",
		"course": "golang",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another fakeed imooc.com",
	})
	return s.Get(url)
}

// duck typing
func main() {
	var r RetrieverPoster
	r = mock.Retriever{"this is a imooc.com"}
	// r = &real.Retriever{
	// 	UserAgent: "Mozilla/5.0",
	// 	TimeOUt:   time.Minute,
	// }
	fmt.Printf("%T %v\n", r, r)

	fmt.Println("Try a session")
	fmt.Println(session(r))

}

func inspect(r Retriever) {
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contects:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
