package mongodb_handler

import (
	"fmt"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
)

type Handler struct {
	MongodbUrl string
}

func New(mongodbUrl string) *Handler {
	return &Handler{
		MongodbUrl: mongodbUrl,
	}
}

func (h *Handler) TestRun() {
	docs := bson.M{
		"name": "mark",
		"age":  10,
	}
	fmt.Printf("[docs][%v][%v]\n", reflect.TypeOf(docs), docs)

	data, err := bson.Marshal(docs)
	if err != nil {
		panic(err)
	}

	fmt.Printf("[data][%v]\n", data)

	var originDocs bson.M
	_ = bson.Unmarshal(data, &originDocs)

	fmt.Printf("[originDocs][%v]\n", originDocs)
}
