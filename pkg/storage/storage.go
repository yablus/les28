package storage

import "github.com/yablus/les28/pkg/student"

/*
type Student struct {
	Name  string
	Age   int
	Grade int
}
*/

type Storage struct {
	Students map[string]*student.Student
}

func NewStorage() *Storage {
	return &Storage{
		Students: make(map[string]*student.Student),
	}
}

func (m *Storage) Put(std *student.Student) bool {
	if m.contains(std.Name) {
		return false
	}
	m.Students[std.Name] = std
	return true
}

func (m *Storage) contains(stdName string) bool {
	for _, value := range m.Students {
		if value.Name == stdName {
			return true
		}
	}
	return false
}

func (m *Storage) Get() map[string]*student.Student {
	return m.Students
}
