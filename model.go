package cjson

type TabularJson struct {
	Columns []string	`json:"columns"`
	Data	[][]interface{}	`json:"data"`
}