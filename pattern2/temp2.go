package pattern2

import "fmt"

type Temp2OutputModel struct {
	Makefile string
}

func Temp2(filepath string) Handler {
	return func(input Input) (*Output, error) {
		to := Temp2OutputModel{Makefile: fmt.Sprintf("makefile-model-%s", input.Value1)}
		output := &Output{FilePath:filepath, Model:to}
		return output, nil
	}
}
