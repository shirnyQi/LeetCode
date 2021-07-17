package generateparenthesis

import (
	"fmt"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	// want: ["((()))","(()())","(())()","()(())","()()()"]
	fmt.Println(generateParenthesis(3))

	// want:["()"]
	//fmt.Println(generateParenthesis(1))
}
