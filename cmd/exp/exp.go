package main

import (
	"fmt"
	"os"
)

type CustomError struct {
	msg  string
	code int
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("%s %d", c.msg, c.code)
}

func NewCustomError(msg string, code int) error {
	return &CustomError{msg: msg, code: code}

}

func process() (string, error) {
	f, err := os.Open("foo")
	if err != nil {
		//return "", errors.New("Não foi possível executar o arquivo!")
		//return "", fmt.ErroF("Não foi possível executar o arquivo!")
		return "", NewCustomError("Não foi possível executar o processo", 100)
	}
	return f.Name(), nil
}

func main() {
	r, err := process()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)
}
