package fmtsuit

import (
	"encoding/json"
)

type TestStrcut struct {
	Name string
}

func (t *TestStrcut) String() string {
	marshal, err := json.Marshal(t)
	if err != nil {
		print(err.Error())
	}
	return string(marshal)
	//return t.Name
}
