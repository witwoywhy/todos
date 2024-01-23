package infra

import (
	"fmt"
	"os"
	"sync"

	"github.com/spf13/viper"
)

const ErrUnexpectedJSON = "unexpected end of JSON input"

var TodosStorage string

type FileManager struct {
	Path string
	Mu   sync.Mutex
}

func NewFileManager(path string) *FileManager {
	return &FileManager{
		Path: path,
		Mu:   sync.Mutex{},
	}
}

func InitFileStorage() {
	TodosStorage = viper.GetString("storage.todos")
	_, err := os.Stat(TodosStorage)
	if err == nil {
		return
	}

	_, err = os.Create(TodosStorage)
	if err != nil {
		panic(fmt.Errorf("failed to initialize file storage"))
	}
}
