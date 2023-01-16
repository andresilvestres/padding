package padding

import (
	"fmt"
	"math"
	"strings"

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

	var output string

	value := in[0].(string)
	length := in[1].(int)
	padCharacter := in[2].(string)

	//inputLength := len(value)
	padStringLength := len(padCharacter)

	repeat := math.Ceil(float64(1) + (float64(length-padStringLength))/float64(padStringLength))

	//Padding Left
	output = strings.Repeat(padCharacter, int(repeat)) + value
	output = output[len(output)-length:]

	return output, fmt.Errorf("fnPaddingLeft function must have three arguments")
}
