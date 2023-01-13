package padding

import (
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnPaddingRight{})
}

type fnPaddingRight struct {
}

func (fnPaddingRight) Name() string {
	return "padding"
}

func (fnPaddingRight) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, true
}

func (fnPaddingRight) Eval(params ...interface{}) (interface{}, error) {

	//PaddingRight
	fmt.Printf("'%-4s'", params)

	return params, fmt.Errorf("PaddingRight")

}
