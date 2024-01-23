package gettodo

import (
	"encoding/json"
	"fmt"
	"os"
	"todos/infra"
	"todos/ports/bizmodel"
)

const errUnexpectedJSON = "unexpected end of JSON input"

type adaptorFile struct {
	fileManager *infra.FileManager
}

func NewAdaptorFile(fileManager *infra.FileManager) Port {
	return &adaptorFile{
		fileManager: fileManager,
	}
}

func (a *adaptorFile) Execute(request Request) (*Response, error) {
	a.fileManager.Mu.Lock()
	defer a.fileManager.Mu.Unlock()

	return newGetInfo(a.fileManager.Path).
		readFile().
		byteToObject().
		getTodo(request)
}

type getInfo struct {
	path     string
	byteData []byte
	todos    map[string]bizmodel.Todo
	err      error
}

func newGetInfo(path string) *getInfo {
	return &getInfo{
		path:  path,
		todos: make(map[string]bizmodel.Todo),
	}
}

func (i *getInfo) readFile() *getInfo {
	b, err := os.ReadFile(i.path)
	if err != nil {
		i.err = fmt.Errorf("failed to read file: %v", err)
	}

	i.byteData = b
	return i
}

func (i *getInfo) byteToObject() *getInfo {
	if i.err != nil {
		return i
	}

	err := json.Unmarshal(i.byteData, &i.todos)
	if err != nil && err.Error() != errUnexpectedJSON {
		i.err = fmt.Errorf("failed to transfer byte to object: %v", err)
	}

	return i
}

func (i *getInfo) getTodo(request Request) (*Response, error) {
	if i.err != nil {
		return nil, i.err
	}

	todo, ok := i.todos[request.ID]
	if !ok {
		return nil, ErrNotFound
	}

	return &todo, nil
}
