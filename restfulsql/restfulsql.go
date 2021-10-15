package restfulsql

import (
	"encoding/json"
	"errors"
	"reflect"
)

// RSQL 基础结构
type RSQL struct {
	Mode   string
	Fields []interface{}
	Values []interface{}

	refFields       reflect.Value
	refValues       reflect.Value
	refFieldsLength int
	refValuesLength int

	nestedSQLCount int
}

// Parser 解析器
type Parser struct {
	Query *RSQL

	rawQuery string
}

// 错误
var (
	ErrInvalidRestfulSQL          = errors.New("invalid restful sql")
	ErrNumOfFieldAndValueMismatch = errors.New("number of fields and values mismatched")
)

// NewRestfulSQLParser 新建一个 RestfulSQLParser 解析器
func NewRestfulSQLParser(query string) *Parser {
	return &Parser{
		rawQuery: query,
	}
}

// Parse 解析 SQL
func (r *Parser) Parse() (*RSQL, error) {
	err := json.Unmarshal([]byte(r.rawQuery), &r.Query)
	if err != nil {
		return nil, err
	}
	if r.Query.refFieldsLength != r.Query.refValuesLength {
		return nil, ErrNumOfFieldAndValueMismatch
	}

	return r.Query, nil
}

// UnmarshalJSON 反序列化
func (r *RSQL) UnmarshalJSON(data []byte) error {
	var raw []interface{}
	err := json.Unmarshal([]byte(data), &raw)
	if err != nil {
		return err
	}
	if len(raw) != 3 {
		return ErrInvalidRestfulSQL
	}

	var ok bool
	r.Mode, ok = raw[0].(string)
	if !ok {
		return ErrInvalidRestfulSQL
	}

	r.Fields, ok = raw[1].([]interface{})
	if !ok {
		return ErrInvalidRestfulSQL
	}

	r.Values, ok = raw[2].([]interface{})
	if !ok {
		return ErrInvalidRestfulSQL
	}

	r.refFields = reflect.ValueOf(r.Fields)
	r.refValues = reflect.ValueOf(r.Values)
	r.refFieldsLength = r.refFields.Len()
	r.refValuesLength = r.refValues.Len()
	r.nestedSQLCount = nestedSQLCount(r.refFields, r.refValues)
	return nil
}

func nestedSQLCount(fields, values reflect.Value) int {
	nestedFieldCount := 0
	for i := 0; i < fields.Len(); i++ {
		_, ok := fields.Index(i).Interface().([]interface{})
		if ok {
			nestedFieldCount++
		}
	}

	return nestedFieldCount
}

func buildNestedSQLFromFields() {

}
