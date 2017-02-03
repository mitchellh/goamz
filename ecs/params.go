package ecs

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"launchpad.net/goamz/aws"
)

type parameters map[string]string

func makeParams(action string) parameters {
	return parameters{
		"Action":  action,
		"Version": "2014-11-13",
	}
}

// #Set uses reflection to assign parameters based on the
// form tag.
func (p parameters) Set(v interface{}) parameters {
	r := reflect.TypeOf(v)
	formStruct := reflect.ValueOf(v)

	// if we have a pointer, dereference it
	if r.Kind() == reflect.Ptr {
		return p.Set(formStruct.Elem().Interface())
	}

	// enumerate through each field, search for the form tag
	for i := 0; i < r.NumField(); i++ {
		typeField := r.Field(i)
		if key := typeField.Tag.Get("form"); key != "" {
			value := formStruct.Field(i).Interface()
			p.assign(key, value)
		}
	}

	return p
}

func (p parameters) assign(key string, value interface{}) {
	set := func(k, v string) {
		if v != "" {
			p[k] = v
		}
	}

	switch v := value.(type) {
	case []string:
		for i, s := range v {
			set(fmt.Sprintf("%s.%d", key, i+1), s)
		}
	case string:
		set(key, v)
	case bool:
		set(key, strconv.FormatBool(v))
	case int:
		if v != 0 {
			set(key, strconv.Itoa(v))
		}
	}
}

func (p parameters) encoded() string {
	// AWS specifies that the parameters in a signed request must
	// be provided in the natural order of the keys. This is distinct
	// from the natural order of the encoded value of key=value.
	// Percent and equals affect the sorting order.
	var keys, sarray []string
	for k, _ := range p {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		sarray = append(sarray, aws.Encode(k)+"="+aws.Encode(p[k]))
	}
	return strings.Join(sarray, "&")
}
