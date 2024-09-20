package utils

import (
	"encoding/json"
)

func ToJson(a any) string {
	b, _ := json.Marshal(a)
	return string(b)
}
