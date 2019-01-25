package pattern1

import "fmt"

type Input struct {
	Value1 string
	Value2 string
}

type Output struct {
	FilePath string
	Model interface{}
}

type Handler interface {
	Execute(input Input) (*Output, error)
}

func Handle(input Input, handler Handler) error {
	output, err := handler.Execute(input)
	if err != nil {
		return err
	}
	fmt.Println(output)
	return nil
}
