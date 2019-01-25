package pattern1

import "fmt"

type Temp1OutputModel struct {
	Concourse string
}

type Temp1 struct {
	FilePath string
}

func (t Temp1) Execute(input Input) (*Output, error) {
	to, err := t.generateOutput(input)
	if err != nil {
		return nil, err
	}
	output := &Output{FilePath:t.FilePath, Model:to}
	return output, nil
}

func (t Temp1) generateOutput(input Input) (Temp1OutputModel, error) {
	return Temp1OutputModel{Concourse: fmt.Sprintf("concourse-model-%s-%s", input.Value1, input.Value2)}, nil
}
