package string

import (
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnPaddingLeft{})
}

type fnPaddingLeft struct {
}

func (fnPaddingLeft) Name() string {
	return "padding"
}

func (fnPaddingLeft) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, true
}

func (fnPaddingLeft) Eval(params ...interface{}) (interface{}, error) {

	//Padding Left
	//fmt.Println("'%4dkm'", params)
	//fmt.Printf("'%4dkm'", params)
	//return params, fmt.Errorf("Paddingleft")

	if params == nil {
		//Do nothing
		return 0, nil
	}
	return "...ERROR...", fmt.Errorf("fnPaddingLeft function must have three arguments")
}
