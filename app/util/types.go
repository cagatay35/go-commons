package util

import (
	"github.com/magiconair/properties"
	"strconv"
	"strings"
)

type Int *int
type Byte *int8
type Short *int16
type Long *int64
type Float *float64
type String *string
type Bool *bool

func ToInt(v interface{}) Int {
	if v == nil {
		return nil
	}
	if value, ok := v.(string); ok && len(value) > 0 {
		i, _ := strconv.Atoi(value)
		return &i
	}
	if value, ok := v.(int64); ok {
		i := int(value)
		return &i
	}
	if value, ok := v.(int32); ok {
		i := int(value)
		return &i
	}
	if value, ok := v.(int16); ok {
		i := int(value)
		return &i
	}
	if value, ok := v.(int8); ok {
		i := int(value)
		return &i
	}
	if value, ok := v.(bool); ok {
		i := 0
		if value {
			i = 1
			return &i
		}
		return &i
	}
	return nil
}

func ToIntOrDefault(value interface{}, d int) int {
	if v := ToInt(value); v != nil {
		return *v
	}
	return d
}

func ToLong(v interface{}) Long {
	if v == nil {
		return nil
	}
	if value, ok := v.(string); ok && len(value) > 0 {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil
		}
		return &i
	}
	if value, ok := v.(int); ok {
		i := int64(value)
		return &i
	}
	if value, ok := v.(int64); ok {
		return &value
	}
	if value, ok := v.(int32); ok {
		i := int64(value)
		return &i
	}
	if value, ok := v.(int16); ok {
		i := int64(value)
		return &i
	}
	if value, ok := v.(int8); ok {
		i := int64(value)
		return &i
	}
	if value, ok := v.(bool); ok {
		var i int64
		if value {
			i = 1
			return &i
		}
		return &i
	}
	return nil
}

func ToLongOrDefault(value interface{}, d int64) int64 {
	if v := ToLong(value); v != nil {
		return *v
	}
	return d
}

func ToFloat(v interface{}) Float {
	if v == nil {
		return nil
	}
	if value, ok := v.(string); ok && len(value) > 0 {
		i, err := strconv.ParseFloat(value, 64)

		if err != nil {
			rv := strings.ReplaceAll(value, ".", "")
			rv = strings.ReplaceAll(rv, ",", ".")

			i, err = strconv.ParseFloat(rv, 64)

			if err != nil {
				rv = strings.ReplaceAll(value, ",", "")
				rv = strings.ReplaceAll(rv, ".", ",")

				i, _ = strconv.ParseFloat(rv, 64)
			}
		}

		return &i
	}
	if value, ok := v.(int); ok {
		i := float64(value)
		return &i
	}
	if value, ok := v.(int32); ok {
		i := float64(value)
		return &i
	}
	if value, ok := v.(int16); ok {
		i := float64(value)
		return &i
	}
	if value, ok := v.(int8); ok {
		i := float64(value)
		return &i
	}
	if value, ok := v.(bool); ok {
		var i float64
		if value {
			i = 1
			return &i
		}
		return &i
	}
	return nil
}

func ToFloatOrDefault(value interface{}, d float64) float64 {
	if v := ToFloat(value); v != nil {
		return *v
	}
	return d
}

func ToStringOrDefault(value interface{}, d string) string {
	if v := ToString(value); v != nil {
		return *v
	}
	return d
}

func ToString(v interface{}) *string {
	if value, ok := v.(string); ok && len(value) > 0 {
		return &value
	}
	return nil
}

func ToBool(v interface{}) Bool {
	t := true
	f := false

	if v == nil {
		return nil
	}
	if value, ok := v.(string); ok && len(value) > 0 {
		i, _ := strconv.ParseBool(value)
		return &i
	}
	if value, ok := v.(int64); ok {
		if value > 0 {
			return &t
		}
		return &f
	}
	if value, ok := v.(int32); ok {
		if value > 0 {
			return &t
		}
		return &f
	}
	if value, ok := v.(int16); ok {
		if value > 0 {
			return &t
		}
		return &f
	}
	if value, ok := v.(int8); ok {
		if value > 0 {
			return &t
		}
		return &f
	}
	if value, ok := v.(float64); ok {
		if value > 0 {
			return &t
		}
		return &f
	}
	if value, ok := v.(float32); ok {
		if value > 0 {
			return &t
		}
		return &f
	}
	return nil
}

func NewBool(value bool) Bool {
	return &value
}

func NewByte(value int8) Byte {
	return &value
}

func NewShort(value int16) Short {
	return &value
}

func NewInt(value int) Int {
	return &value
}

func NewLong(value int64) Long {
	return &value
}

func NewFloat(value float64) Float {
	return &value
}

func NewString(value string) String {
	return &value
}

func ConcatStringMap(to map[string]string, from map[string]string) map[string]string {
	for n, v := range from {
		to[n] = v
	}
	return to
}

func GetMessageProperty(messageProperties *properties.Properties, tag string) string {

	res := messageProperties.GetString(tag, "")

	return res
}

func StringEquals(string String, other string) bool {
	if string == nil {
		return other == ""
	}

	return *string == other
}

func StringNotEmpty(string String) bool {
	return string != nil && *string != ""
}
