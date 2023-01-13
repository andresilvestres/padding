package padding

import (
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnPaddingZeros{})
}

type fnPaddingZeros struct {
}

func (fnPaddingZeros) Name() string {
	return "padding"
}

func (fnPaddingZeros) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, true
}

func (fnPaddingZeros) Eval(params ...interface{}) (interface{}, error) {

	//Padding with zeroes
	fmt.Printf("'%04dkm'", params)

	return params, fmt.Errorf("PaddingWithZeroes")
}
