package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic"
)

// ES insert data demo

// Student ...
type Student struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func (s *Student) run() *Student {
	fmt.Printf("%s在跑...", s.Name)
	return s
}

func (s *Student) wang() {
	fmt.Printf("%s在汪汪汪的叫...\n", s.Name)
}

func main() {
	// luminghui := Student{
	// 	Name:    "卢明辉",
	// 	Age:     9000,
	// 	Married: false,
	// }
	// luminghui.run()
	// luminghui.wang()

	// luminghui.run().wang()

	client, err := elastic.NewClient(elastic.SetURL("http://192.168.34.112:9200"))
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Println("connect to es success")

	p1 := Student{Name: "Rion", Age: 22, Married: false}
	// 链式操作
	put1, err := client.Index().Index("student").Type("go").BodyJson(p1).Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed student %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}
