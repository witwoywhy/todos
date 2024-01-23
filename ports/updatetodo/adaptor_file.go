package updatetodo

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

func (a *adaptorFile) Execute(request Request) error {
	a.fileManager.Mu.Lock()
	defer a.fileManager.Mu.Unlock()

	return newUpdateInfo().
		readFile(a.fileManager.Path).
		byteToObject().
		updateTodo(request).
		objectToByte().
		writeFile(a.fileManager.Path)
}

type updateInfo struct {
	byteData []byte
	todos    map[string]bizmodel.Todo
	err      error
}

func newUpdateInfo() *updateInfo {
	return &updateInfo{}
}

func (i *updateInfo) readFile(path string) *updateInfo {
	b, err := os.ReadFile(path)
	if err != nil {
		i.err = fmt.Errorf("failed to read file: %v", err)
	}

	i.byteData = b
	return i
}

func (i *updateInfo) byteToObject() *updateInfo {
	if i.err != nil {
		return i
	}

	err := json.Unmarshal(i.byteData, &i.todos)
	if err != nil && err.Error() != errUnexpectedJSON {
		i.err = fmt.Errorf("failed to transfer byte to object: %v", err)
	}

	return i
}

func (i *updateInfo) updateTodo(todo bizmodel.Todo) *updateInfo {
	if i.err != nil {
		return i
	}

	i.todos[todo.ID] = todo

	return i
}

func (i *updateInfo) objectToByte() *updateInfo {
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

func (i *updateInfo) writeFile(path string) error {
	if i.err != nil {
		return i.err
	}

	err := os.WriteFile(path, i.byteData, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	return nil
}
