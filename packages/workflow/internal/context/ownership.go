package workercontext

import "reflect"

// Captured public document data consists of exported structs, scalar values,
// maps and slices. Preserve their concrete types and nil/empty distinction while
// giving each returned context its own storage, including nested metadata.
func owned[T any](value T) T { return ownValue(reflect.ValueOf(value)).Interface().(T) }
func ownValue(value reflect.Value) reflect.Value {
	switch value.Kind() {
	case reflect.Interface, reflect.Pointer:
		if value.IsNil() {
			return reflect.Zero(value.Type())
		}
		result := reflect.New(value.Type()).Elem()
		if value.Kind() == reflect.Pointer {
			result = reflect.New(value.Type().Elem())
			result.Elem().Set(ownValue(value.Elem()))
		} else {
			result.Set(ownValue(value.Elem()))
		}
		return result
	case reflect.Map:
		if value.IsNil() {
			return reflect.Zero(value.Type())
		}
		result := reflect.MakeMapWithSize(value.Type(), value.Len())
		it := value.MapRange()
		for it.Next() {
			result.SetMapIndex(it.Key(), ownValue(it.Value()))
		}
		return result
	case reflect.Slice:
		if value.IsNil() {
			return reflect.Zero(value.Type())
		}
		result := reflect.MakeSlice(value.Type(), value.Len(), value.Len())
		for i := 0; i < value.Len(); i++ {
			result.Index(i).Set(ownValue(value.Index(i)))
		}
		return result
	case reflect.Array, reflect.Struct:
		result := reflect.New(value.Type()).Elem()
		if value.Kind() == reflect.Struct {
			for i := 0; i < value.NumField(); i++ {
				result.Field(i).Set(ownValue(value.Field(i)))
			}
		} else {
			for i := 0; i < value.Len(); i++ {
				result.Index(i).Set(ownValue(value.Index(i)))
			}
		}
		return result
	default:
		return value
	}
}
