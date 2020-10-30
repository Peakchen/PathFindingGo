package core

import "reflect"

/*

for go 1.8

func reverseAny(s interface{}) {
    n := reflect.ValueOf(s).Len()
    swap := reflect.Swapper(s)
    for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
        swap(i, j)
    }
}
*/

func Reverse(nodes interface{}) {
	data := reflect.ValueOf(nodes)
	vlen := reflect.ValueOf(nodes).Len()
	switch nodes.(type) {
	case ArrayInt8:
		is := (data.Interface().(ArrayInt8))
		for i, j := 0, vlen-1; i < j; i, j = i+1, j-1 {
			is[i], is[j] = is[j], is[i]
		}
	case ArrayInt16:
		is := (data.Interface().(ArrayInt16))
		for i, j := 0, vlen-1; i < j; i, j = i+1, j-1 {
			is[i], is[j] = is[j], is[i]
		}
	case ArrayInt32:
		is := (data.Interface().(ArrayInt32))
		for i, j := 0, vlen-1; i < j; i, j = i+1, j-1 {
			is[i], is[j] = is[j], is[i]
		}
	case ArrayInt64:
		is := (data.Interface().(ArrayInt64))
		for i, j := 0, vlen-1; i < j; i, j = i+1, j-1 {
			is[i], is[j] = is[j], is[i]
		}
	case ArrayUInt8:
		is := (data.Interface().(ArrayUInt8))
		for i, j := 0, vlen-1; i < j; i, j = i+1, j-1 {
			is[i], is[j] = is[j], is[i]
		}
	case ArrayUInt16:
		is := (data.Interface().(ArrayUInt16))
		for i, j := 0, vlen-1; i < j; i, j = i+1, j-1 {
			is[i], is[j] = is[j], is[i]
		}
	case ArrayUInt32:
		is := (data.Interface().(ArrayUInt32))
		for i, j := 0, vlen-1; i < j; i, j = i+1, j-1 {
			is[i], is[j] = is[j], is[i]
		}
	case ArrayUInt64:
		is := (data.Interface().(ArrayUInt64))
		for i, j := 0, vlen-1; i < j; i, j = i+1, j-1 {
			is[i], is[j] = is[j], is[i]
		}
	case DoubleInt32:
		is := (data.Interface().(DoubleInt32))
		for i, j := 0, vlen-1; i < j; i, j = i+1, j-1 {
			is[i], is[j] = is[j], is[i]
		}
	case DoubleInt64:
		is := (data.Interface().(DoubleInt64))
		for i, j := 0, vlen-1; i < j; i, j = i+1, j-1 {
			is[i], is[j] = is[j], is[i]
		}
	default:
		panic("invalid data type.")
	}

}
