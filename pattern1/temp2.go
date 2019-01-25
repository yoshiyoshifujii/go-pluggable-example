package pattern1

import "fmt"

type Temp2OutputModel struct {
	Makefile string
}

type Temp2 struct {
	FilePath string
}

func (t Temp2) Execute(input Input) (*Output, error) {
	to, err := t.generateOutput(input)
	if err != nil {
		return nil, err
	}
	output := &Output{FilePath:t.FilePath, Model:to}
	return output, nil
}

func (t Temp2) generateOutput(input Input) (Temp2OutputModel, error) {
	return Temp2OutputModel{Makefile: fmt.Sprintf("makefile-model-%s", input.Value1)}, nil
}
