package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"reflect"
	"runtime"
	"slices"
	"strings"
)

const replaceStr = "**********"

var sensitives = []string{"email", "password"}

func init() {
	opts := &slog.HandlerOptions{
		AddSource:   true,
		ReplaceAttr: ReplaceAttr,
	}
	jsonHandler := slog.NewJSONHandler(os.Stdout, opts)
	slog.SetDefault(slog.New(jsonHandler))
}

func Debug(ctx context.Context, msg string, groups []any) {
	slog.Debug(
		msg,
		slog.Group("context",
			slog.Any("userID", ctx.Value("test")),
		),
		slog.Group("attr",
			groups...,
		),
	)
}

func Info(ctx context.Context, msg string, groups []any) {
	slog.Info(
		msg,
		slog.Group("context",
			slog.Any("userID", ctx.Value("test")),
		),
		slog.Group("attr",
			groups...,
		),
	)
}

func Warn(ctx context.Context, msg string, groups []any) {
	slog.Warn(
		msg,
		slog.Group("context",
			slog.Any("userID", ctx.Value("test")),
		),
		slog.Group("attr",
			groups...,
		),
	)
}

func Error(ctx context.Context, msg string, groups []any) {
	slog.Error(
		msg,
		slog.Group("context",
			slog.Any("userID", ctx.Value("test")),
		),
		slog.Group("attr",
			groups...,
		),
	)
}

func ReplaceAttr(groups []string, a slog.Attr) slog.Attr {
	// add source info
	if a.Key == slog.SourceKey {
		const skip = 7
		_, file, line, ok := runtime.Caller(skip)
		if !ok {
			return a
		}
		v := fmt.Sprintf("%s:%d", file, line)
		a.Value = slog.StringValue(v)
	}

	// if groups are empty, finish this func
	if len(groups) == 0 {
		return a
	}

	for _, group := range groups {
		// Replace sentences only "attr" group
		if group != "attr" {
			continue
		}

		key := strings.ToLower(a.Key)
		if slices.Contains(sensitives, key) {
			return slog.Attr{
				Key:   a.Key,
				Value: slog.StringValue(replaceStr),
			}
		}

		// get value type, mask it
		value := a.Value.Any()

		v := recursiveReplaceAttr(value)

		return slog.Attr{Key: a.Key, Value: slog.AnyValue(v)}
	}

	return a
}

func GetDiff[T any](before, after T) (map[string]any, map[string]any) {
	afterM, beforeM := recursiveGetDiff(after, before)

	return beforeM.(map[string]any), afterM.(map[string]any)
}

func recursiveReplaceAttr(value any) any {
	// value may be pointer, so reference the real value
	v := reflect.Indirect(reflect.ValueOf(value))

	// skip private field
	if v.Type().PkgPath() != "" {
		return value
	}

	switch v.Kind() {
	case reflect.Struct:
		newValue := reflect.New(v.Type()).Elem()
		newValue = recursiveReplaceAttrForStruct(v, newValue)

		return newValue.Interface()
	case reflect.Slice:
		newValue := reflect.MakeSlice(v.Type(), v.Len(), v.Cap())
		for i := 0; i < v.Len(); i++ {
			val := recursiveReplaceAttr(v.Index(i).Interface())
			// case of pointer
			if v.Index(i).Kind() == reflect.Ptr {
				if v.Index(i).IsNil() {
					continue
				}
				// copy pointer type
				typeVal := reflect.New(reflect.TypeOf(val))
				// set the real value of pointer
				typeVal.Elem().Set(reflect.ValueOf(val))
				// set pointer
				newValue.Index(i).Set(typeVal)
			} else {
				newValue.Index(i).Set(reflect.ValueOf(val))
			}
		}

		return newValue.Interface()
	case reflect.Map:
		newValue := reflect.MakeMap(v.Type())
		newValue = recursiveReplaceAttrForMap(v, newValue)

		return newValue
	default:

		return value
	}
}

func recursiveReplaceAttrForStruct(originValue, newValue reflect.Value) reflect.Value {
	for i := 0; i < originValue.NumField(); i++ {
		originField := originValue.Type().Field(i)
		originFieldValue := originValue.Field(i)
		// skip private field
		if originField.PkgPath != "" {
			continue
		}

		originFieldName := originField.Name
		lowerFieldName := strings.ToLower(originFieldName)

		// if sensitive fields
		if slices.Contains(sensitives, lowerFieldName) {
			switch originFieldValue.Kind() {
			case reflect.String:
				// replace "replaceStr" when field type is string
				newValue.Field(i).SetString(replaceStr)
			case reflect.Ptr:
				if originFieldValue.IsNil() {
					continue
				}
				newValue.Field(i).Set(reflect.New(originFieldValue.Type().Elem()))
				if originFieldValue.Elem().Kind() != reflect.String {
					val := recursiveReplaceAttr(originFieldValue.Elem().Interface())
					newValue.Field(i).Elem().Set(reflect.ValueOf(val))
				} else {
					newValue.Field(i).Elem().SetString(replaceStr)
				}
			}
		} else {
			if originFieldValue.Kind() == reflect.Ptr {
				if originFieldValue.IsNil() {
					continue
				}
				val := recursiveReplaceAttr(originFieldValue.Elem().Interface())
				// copy pointer type
				typeVal := reflect.New(reflect.TypeOf(val))
				// set the real value of pointer
				typeVal.Elem().Set(reflect.ValueOf(val))
				// set pointer
				newValue.Field(i).Set(typeVal)
			} else {
				val := recursiveReplaceAttr(originFieldValue.Interface())
				newValue.Field(i).Set(reflect.ValueOf(val))
			}
		}
	}

	return newValue
}

func recursiveReplaceAttrForMap(originValue, newValue reflect.Value) reflect.Value {
	for _, key := range originValue.MapKeys() {
		lowerKeyStr := strings.ToLower(key.String())
		if slices.Contains(sensitives, lowerKeyStr) {
			if originValue.MapIndex(key).Kind() == reflect.Ptr {
				if originValue.MapIndex(key).IsNil() {
					continue
				}
				// type is only *string. if needs, you can fix
				if originValue.MapIndex(key).Elem().Kind() == reflect.String {
					// copy pointer type
					typeVal := reflect.New(reflect.TypeOf(replaceStr))
					// set the real value of pointer
					typeVal.Elem().Set(reflect.ValueOf(replaceStr))
					// set pointer
					newValue.SetMapIndex(key, typeVal)
				}
			} else {
				newValue.SetMapIndex(key, reflect.ValueOf(replaceStr))
			}
			continue
		} else {
			var valRef any
			mapVal := originValue.MapIndex(key)
			if mapVal.Kind() == reflect.Ptr {
				if mapVal.IsNil() {
					continue
				}
				valRef = mapVal.Elem().Interface()
				val := recursiveReplaceAttr(valRef)
				// copy pointer type
				typeVal := reflect.New(reflect.TypeOf(val))
				// set the real value of pointer
				typeVal.Elem().Set(reflect.ValueOf(val))
				// set pointer
				newValue.SetMapIndex(key, typeVal)
			} else {
				valRef = mapVal.Interface()
				val := recursiveReplaceAttr(valRef)
				newValue.SetMapIndex(key, reflect.ValueOf(val))
			}
		}
	}

	return newValue
}

func recursiveGetDiff(beforeValue, afterValue any) (any, any) {
	// create initial map
	beforeM := make(map[string]any)
	afterM := make(map[string]any)

	beforeV := reflect.ValueOf(beforeValue)
	afterV := reflect.ValueOf(afterValue)

	switch beforeV.Kind() {
	case reflect.Struct:
		for i := 0; i < beforeV.NumField(); i++ {
			// skip private field
			if beforeV.Type().Field(i).PkgPath != "" {
				continue
			}

			if beforeV.Field(i).Kind() == reflect.Ptr {
				if beforeV.Field(i).IsNil() && afterV.Field(i).IsNil() {
					continue
				}
				// the one is nil, but the other is not nil
				if beforeV.Field(i).IsNil() || afterV.Field(i).IsNil() {
					if beforeV.Field(i).IsNil() {
						beforeM[beforeV.Type().Field(i).Name] = "nil"
						afterM[beforeV.Type().Field(i).Name] = afterV.Field(i).Elem().Interface()
					} else {
						beforeM[beforeV.Type().Field(i).Name] = beforeV.Field(i).Elem().Interface()
						afterM[beforeV.Type().Field(i).Name] = "nil"
					}
					continue
				}
				beforeDiff, afterDiff := recursiveGetDiff(beforeV.Field(i).Elem().Interface(), afterV.Field(i).Elem().Interface())
				if !reflect.DeepEqual(beforeDiff, afterDiff) {
					beforeM[beforeV.Type().Field(i).Name] = beforeDiff
					afterM[beforeV.Type().Field(i).Name] = afterDiff
				}
			}
		}
	case reflect.Map:
		for _, key := range afterV.MapKeys() {
			// key: after has, but before not
			if !beforeV.MapIndex(key).IsValid() {
				beforeM[key.String()] = "nil"
				afterM[key.String()] = afterV.MapIndex(key).Interface()
			}
		}

		for _, key := range beforeV.MapKeys() {
			// key: before has, but after not
			if !afterV.MapIndex(key).IsValid() {
				beforeM[key.String()] = beforeV.MapIndex(key).Interface()
				afterM[key.String()] = "nil"
				continue
			}
			// key: before and after has
			beforeDiff, afterDiff := recursiveGetDiff(beforeV.MapIndex(key).Interface(), afterV.MapIndex(key).Interface())
			if !reflect.DeepEqual(beforeDiff, afterDiff) {
				beforeM[key.String()] = beforeDiff
				afterM[key.String()] = afterDiff
			}
		}
	case reflect.Slice:
		// if the length is different, return
		if beforeV.Len() != afterV.Len() {
			return beforeValue, afterValue
		}
		for i := 0; i < beforeV.Len(); i++ {
			beforeDiff, afterDiff := recursiveGetDiff(beforeV.Index(i).Interface(), afterV.Index(i).Interface())
			if !reflect.DeepEqual(beforeDiff, afterDiff) {
				beforeM[fmt.Sprintf("%d", i)] = beforeDiff
				afterM[fmt.Sprintf("%d", i)] = afterDiff
			}
		}
	case reflect.Ptr:
		// check the value of pointer
		if beforeV.IsNil() && afterV.IsNil() {
			return nil, nil
		} else if beforeV.IsNil() || afterV.IsNil() {
			if beforeV.IsNil() {
				return "nil", afterV.Elem().Interface()
			} else {
				return beforeV.Elem().Interface(), "nil"
			}
		} else {
			beforeDiff, afterDiff := recursiveGetDiff(beforeV.Elem().Interface(), afterV.Elem().Interface())
			if !reflect.DeepEqual(beforeDiff, afterDiff) {
				return beforeDiff, afterDiff
			}
		}
	default:
		// only primitive value. if you need to deal channel, unsafe.Pointer, fix
		if beforeValue == afterValue {
			return nil, nil
		}
		return beforeValue, afterValue
	}

	return beforeM, afterM
}
