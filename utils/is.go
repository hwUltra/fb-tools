package utils

import (
	"reflect"
	"strconv"
)

// IsNumeric  判断是否是数字
func IsNumeric(item interface{}) bool {
	switch item.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true
	case float32, float64, complex64, complex128:
		return true

	case string:
		return false
	default:
		return false
	}
}

func isNum(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func IsEmpty(data interface{}) bool {
	rv := reflect.ValueOf(data)
	rvk := rv.Kind()
	switch rvk {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		return rv.IsZero()
	case reflect.Uintptr, reflect.Map, reflect.Slice, reflect.Pointer, reflect.UnsafePointer:
		return rv.IsNil() // 指针类型数据全部调用 IsNil() 进行判断,非nil的指针视为非空
	case reflect.String:
		return rv.Len() == 0
	case reflect.Struct:
		if rv.MethodByName("IsZero").IsValid() {
			// 动态调用结构体中的 IsZero 方法
			rt := rv.MethodByName("IsZero").Call([]reflect.Value{}) // call后面的参数即是要调用方法的参数 []reflect.Value{} 表示无参数
			return len(rt) > 0 && rt[0].Bool()                      // IsZero方法的返回值只有一个bool,
		} else {
			// 对值的零值类型进行深度比较
			return reflect.DeepEqual(rv.Interface(), reflect.Zero(rv.Type()).Interface())
		}
	case reflect.Invalid: // nil值的reflect类型就是 Invalid
		return true // nil也算是空值
	}

	return false // 其他情况默认非空
}
