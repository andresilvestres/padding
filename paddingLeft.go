package padding

import (
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

type fnPaddingLeft struct {
}

func init() {
	function.Register(&fnPaddingLeft{})
}

func (s *fnPaddingLeft) Name() string {
	return "paddingLeft"
}

func (s *fnPaddingLeft) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeInt, data.TypeString}, false
}

func (s *fnPaddingLeft) Eval(in ...interface{}) (interface{}, error) {

	//Padding Left
	fmt.Println("testing...................")
	//fmt.Printf("'%4dkm'", params)
	//return in, fmt.Errorf("Paddingleft")

	//if params == nil {
		//Do nothing
		//return 0, nil
	//}
	return "...ERROR...", fmt.Errorf("fnPaddingLeft function must have three arguments")
}
