package helper

// Integer 定义泛型整型集合
type Integer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// InArray 函数用于检查给定的元素是否存在于切片中
func InArray[T Integer | string](val T, inp []T) bool {
	for _, v := range inp {
		if v == val {
			return true
		}
	}
	return false
}

// ArrUnique 函数用于对切片中的元素去重
func ArrUnique[T Integer | string](inp []T) []T {
	out := make([]T, 0, len(inp))
	mp := make(map[T]bool)
	for _, val := range inp {
		if ok := mp[val]; !ok {
			out = append(out, val)
		}
		mp[val] = true
	}
	return out
}

// ArrFilter 函数用于过滤切片，默认去除零值
func ArrFilter[T Integer | string](inp []T, fn ...func(val T) bool) []T {
	var zero T
	out := make([]T, 0, len(inp))
	for _, val := range inp {
		if len(fn) > 0 {
			if fn[0](val) {
				out = append(out, val)
			}
		} else {
			if val != zero {
				out = append(out, val)
			}
		}
	}
	return out
}
