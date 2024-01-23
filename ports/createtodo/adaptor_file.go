package createtodo

import (
	"encoding/json"
	"fmt"
	"os"
	"todos/infra"
	"todos/ports/bizmodel"
)

type adaptorFile struct {
	fileManager *infra.FileManager
}

func NewAdaptorFile(fileManager *infra.FileManager) Port {
	return &adaptorFile{
		fileManager: fileManager,
	}
}

func (a *adaptorFile) Execute(request Request) error {
	a.fileManager.Mu.Lock()
	defer a.fileManager.Mu.Unlock()

	return newCreateInfo().
		readFile(a.fileManager.Path).
		byteToObject().
		addTodo(request).
		objectToByte().
		writeFile(a.fileManager.Path)
}

type createInfo struct {
	byteData []byte
	todos    map[string]bizmodel.Todo
	err      error
}

func newCreateInfo() *createInfo {
	return &createInfo{
		todos: make(map[string]bizmodel.Todo),
	}
}

func (i *createInfo) readFile(path string) *createInfo {
	b, err := os.ReadFile(path)
	if err != nil {
		i.err = fmt.Errorf("failed to read file: %v", err)
	}

	i.byteData = b
	return i
}

func (i *createInfo) byteToObject() *createInfo {
	if i.err != nil {
		return i
	}

	err := json.Unmarshal(i.byteData, &i.todos)
	if err != nil && err.Error() != infra.ErrUnexpectedJSON {
		i.err = fmt.Errorf("failed to transfer byte to object: %v", err)
	}

	return i
}

func (i *createInfo) addTodo(todo bizmodel.Todo) *createInfo {
	if i.err != nil {
		return i
	}

	i.todos[todo.ID] = todo

	return i
}

func (i *createInfo) objectToByte() *createInfo {
	if i.err != nil {
		return i
	}

	b, err := json.Marshal(i.todos)
	if err != nil {
		i.err = fmt.Errorf("failed to object to byte: %v", err)
	} else {
		i.byteData = b
	}

	return i
}

func (i *createInfo) writeFile(path string) error {
	if i.err != nil {
		return i.err
	}

	err := os.WriteFile(path, i.byteData, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	return nil
}
