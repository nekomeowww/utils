package collection

import "reflect"

// ToMapWithSize transforms a slice of structs to a map based on a pivot field: []*Foo => Map<int, *Foo>
// Code from https://github.com/thoas/go-funk, with slight modification.
// n is the pre-allocated size of the map.
// NOTICE: This function should be used in the same package where the struct type of in is defined, to avoid typed coupling.
// 从 benchmark 的结果看，运行速度相比原生写法，在数量级上差不太多
func ToMapWithSize(in interface{}, pivot string, n int) interface{} {
	value := reflect.ValueOf(in)

	// input value must be a slice
	if value.Kind() != reflect.Slice {
		panic("input must be a slice")
	}

	sliceType := value.Type()

	// retrieve the struct in the slice to deduce key type
	structType := sliceType.Elem()
	if structType.Kind() == reflect.Ptr {
		structType = structType.Elem()
	}

	keyField, ok := structType.FieldByName(pivot)
	if !ok {
		return nil
	}

	// value of the map will be the input type
	collectionType := reflect.MapOf(keyField.Type, sliceType.Elem())

	// create a map from scratch
	collection := reflect.MakeMapWithSize(collectionType, n)

	for i := 0; i < value.Len(); i++ {
		element := value.Index(i)
		key := reflect.Indirect(element).FieldByName(pivot)
		collection.SetMapIndex(key, element)
	}

	return collection.Interface()
}

// ToMap transforms a slice of instances to a Map.
// []*Foo => Map<int, *Foo>
func ToMap(in interface{}, pivot string) interface{} {
	return ToMapWithSize(in, pivot, 0)
}
