package student

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name  string
	Age   int
	Grade int
}

type Data interface {
	// put
	// добавляет студента в карту.
	// возвращает:
	// true, если добавлен;
	// false, если элемент уже присутствует или был совершен некорректный ввод.
	Put(std *Student) bool

	// get
	// возвращает:
	// карту студентов.
	Get() map[string]Student
}

type DataTest struct{}

func (test *DataTest) Put(std Student) bool {
	return true
}
func (test *DataTest) Get() map[string]*Student {
	return map[string]*Student{
		"Vasya": {"Vasya", 23, 3},
		"Petya": {"Petya", 22, 2},
	}
}

type App struct {
	repository Data
}

func (a *App) Run() {
	for {
		if std, ok := a.inputStudent(); ok {
			a.saveStudent(std)
		} else {
			a.printStudents()
			break
		}
	}
}

func (a *App) printStudents() {
	fmt.Println("Студенты из хранилища:")
	for _, v := range a.repository.Get() {
		fmt.Printf("%s %d %d\n", v.Name, v.Age, v.Grade)
	}
}

func (a *App) inputStudent() (*Student, bool) {
	for {
		fmt.Println("Введите через пробел имя студента, его возраст и класс")
		var strName, strAge, strGrade string
		empty := Student{
			Name:  "",
			Age:   0,
			Grade: 0,
		}
		str, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err == io.EOF {
			return &empty, false
		}
		param := strings.Split(strings.TrimSpace(str), " ")
		fmt.Println(param)
		if len(param) == 3 || len(param) == 4 {
			strGrade = param[len(param)-1]
			strAge = param[len(param)-2]
			if len(param) == 4 {
				strName = param[0] + " " + param[1]
			} else {
				strName = param[0]
			}
		}
		intAge, errAge := strconv.Atoi(strAge)
		intGrade, errGrade := strconv.Atoi(strGrade)
		if intGrade != 0 && intAge != 0 && strName != "" && errAge == nil && errGrade == nil {
			std := Student{
				Name:  strName,
				Age:   intAge,
				Grade: intGrade,
			}
			return &std, true
		}
		fmt.Println("Некорректный ввод")
	}
}

func (a *App) saveStudent(std *Student) {
	var msg string
	if ok := a.repository.Put(std); ok {
		msg = "Студент %s успешно добавлен\n"
	} else {
		msg = "Студент %s уже есть в базе\n"
	}
	fmt.Printf(msg, std.Name)
}
