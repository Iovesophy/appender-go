package appender

import "errors"

func Push(a []interface{}, v interface{}) []interface{} {
	return append(a, v)
}

func Pop(a []interface{}) []interface{} {
	return a[:len(a)-1]
}

func Unshift(a []interface{}, v interface{}) []interface{} {
	a, a[0] = append(a[:1], a[0:]...), v
	return a
}

func Shift(a []interface{}) []interface{} {
	return a[1:]
}

func Insert(a []interface{}, v interface{}, p int) []interface{} {
	a = append(a, 0)
	a = append(a[:p+1], a[p:len(a)-1]...)
	a[p] = v
	return a
}

func Edit(a []interface{}, v interface{}, p int) []interface{} {
	a = append(a[:p], a[p:]...)
	a[p] = v
	return a
}

func Concat(a []interface{}, b []interface{}) []interface{} {
	a = append(a, b...)
	return a
}

func Remove(a []interface{}, p int) []interface{} {
	return append(a[:p], a[p+1])
}

func Picker(a []interface{}, src []interface{}) ([]interface{}, error) {
	var (
		ci int
		cb int
	)
	for i := range src {
		switch src[i].(type) {
		case int:
			if len(src)-1 == i && ci == i {
				return indexPicker(a, interfaceSliceConvInterfaceToInt(src))
			}
			ci++
		case bool:
			if len(src)-1 == i && cb == i {
				return booleanPicker(a, interfaceSliceConvInterfaceToBool(src))
			}
			cb++
		default:
			return nil, errors.New("picker: unknown type of src")
		}
	}
	return nil, errors.New("picker: mixed type of src")
}

func interfaceSliceConvInterfaceToInt(src []interface{}) []int {
	var r []int
	for _, a := range src {
		r = append(r, a.(int))
	}
	return r
}

func interfaceSliceConvInterfaceToBool(src []interface{}) []bool {
	var r []bool
	for _, a := range src {
		r = append(r, a.(bool))
	}
	return r
}

func booleanPicker(a []interface{}, src []bool) ([]interface{}, error) {
	var r []interface{}
	if len(a) == len(src) {
		for i, v := range src {
			if v {
				r = append(r, a[i])
			}
		}
		return r, nil
	}
	return nil, errors.New("booleanPicker: invalid flag")
}

func indexPicker(a []interface{}, src []int) ([]interface{}, error) {
	var r []interface{}
	for i := range a {
		for j, v := range src {
			if len(a)-1 < v || v < 0 {
				return nil, errors.New("indexPicker: range over cause invalid index slice")
			} else if i == src[j] {
				r = append(r, a[i])
			}
		}
	}
	return r, nil
}
