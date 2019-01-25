package pattern2

import "fmt"

type Temp1OutputModel struct {
	Concourse string
}

func Temp1(filepath string) Handler {
	return func(input Input) (*Output, error) {
		to := Temp1OutputModel{Concourse: fmt.Sprintf("concourse-model-%s-%s", input.Value1, input.Value2)}
		output := &Output{FilePath:filepath, Model:to}
		return output, nil
	}
}
