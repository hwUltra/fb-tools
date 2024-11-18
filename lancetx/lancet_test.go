package lancetx

import (
	"fmt"
	"github.com/duke-git/lancet/v2/validator"
	"testing"
)

func Test_lan(t *testing.T) {
	result1 := validator.ContainChinese("你好")
	result2 := validator.ContainChinese("你好hello")
	result3 := validator.ContainChinese("hello")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}
