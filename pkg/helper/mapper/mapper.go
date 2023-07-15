package mapper

import (
	"reflect"
)

func Map(source, destination any) {
	sourceValue := reflect.ValueOf(source).Elem()
	destValue := reflect.ValueOf(destination).Elem()

	sourceType := sourceValue.Type()
	destType := destValue.Type()

	if sourceType.Kind() != destType.Kind() {
		panic("Source and destination must be the same type (struct or slice)")
	}

	switch sourceType.Kind() {
	case reflect.Struct:
		mapStruct(sourceValue, destValue)
	case reflect.Slice:
		mapSlice(sourceValue, destValue)
	default:
		panic("Source and destination must be either struct or slice")
	}
}

func mapStruct(sourceValue, destValue reflect.Value) {
	destType := destValue.Type()

	var maxIndex = 0

	if destValue.NumField() > sourceValue.NumField() {
		maxIndex = sourceValue.NumField()
	} else {
		maxIndex = destValue.NumField()
	}

	for i := 0; i < maxIndex; i++ {
		destField := destValue.Field(i)
		destFieldName := destType.Field(i).Name

		sourceField := sourceValue.FieldByName(destFieldName)
		if !sourceField.IsValid() {
			continue
		}

		if destField.CanSet() && sourceField.IsValid() && destField.Type() == sourceField.Type() {
			destField.Set(sourceField)
		}
	}
}

func mapSlice(sourceValue, destValue reflect.Value) {
	destType := destValue.Type()
	destSlice := reflect.MakeSlice(destType, sourceValue.Len(), sourceValue.Len())

	for i := 0; i < sourceValue.Len(); i++ {
		sourceElem := sourceValue.Index(i)

		if destValue.CanSet() && sourceElem.IsValid() {
			destElem := reflect.New(destType.Elem()).Elem()
			mapValues(sourceElem, destElem)
			destSlice.Index(i).Set(destElem)
		}
	}
	destValue.Set(destSlice)
}

func mapValues(sourceElem, destElem reflect.Value) {
	sourceType := sourceElem.Type()
	destType := destElem.Type()

	for j := 0; j < sourceType.NumField(); j++ {
		if j < destType.NumField() {

			sourceField := sourceElem.Field(j)
			destField := destElem.Field(j)

			if sourceField.IsValid() && destField.CanSet() {
				if destField.Type() == sourceField.Type() {
					destField.Set(sourceField)
				}
			}
		}
	}
}
