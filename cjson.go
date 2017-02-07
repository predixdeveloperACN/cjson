package cjson

import (
	"reflect"
	"github.com/jackmanlabs/errors"
)

func CompressJSON(src interface{}) (retVal TabularJson, err error) {
	v := reflect.ValueOf(src)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Slice {
		err = errors.New("Unable to compress object. Object should be of type slice.")
		return
	}

	// get columns
	typ := v.Type().Elem()
	var columns []string
	for i := 0; i < v.Type().Elem().NumField(); i++ {
		jsonName, ok := parseJsonField(typ.Field(i))
		if !ok {
			continue
		}
		retVal.Columns = append(retVal.Columns, jsonName)
		columns = append(columns, typ.Field(i).Name)
	}

	// get values
	var dataRows [][]interface{}
	for i := 0; i < v.Len(); i++ {
		r := reflect.ValueOf(v.Index(i).Interface())
		var d []interface{}
		for _, c := range columns {
			d = append(d, r.FieldByName(c).Interface())
		}
		dataRows = append(dataRows, d)
	}
	retVal.Data = dataRows

	return
}