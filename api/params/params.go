package params

import (
	"fmt"
	"strings"

	cf "github.com/iostrovok/go-convert"
)

type Param struct {
	Key   string
	Value any
}

type Params struct {
	Param []Param
}

func NewParams() *Params {
	return &Params{
		Param: []Param{},
	}
}

func (p *Params) AddFilled(key string, value any) *Params {
	if cf.String(value) != "" && (cf.Int(value) != 0 || cf.String(value) != "") {
		p.Param = append(p.Param, Param{Key: key, Value: value})
	}

	return p
}

func (p *Params) Add(key string, value any) *Params {
	p.Param = append(p.Param, Param{Key: key, Value: value})
	return p
}

func AddFilledList[T string | int | int32](p *Params, key string, value []T) {
	for i := range value {
		p.AddFilled(key, value[i])
	}
}

func AddList[T string | int | int32](p *Params, key string, value []T) {
	for i := range value {
		p.Add(key, value[i])
	}
}

func (p *Params) QueryString() string {
	if p != nil && len(p.Param) > 0 {
		var s []string
		for i := range p.Param {
			s = append(s, fmt.Sprintf("%s=%s", p.Param[i].Key, cf.String(p.Param[i].Value)))
		}

		return strings.Join(s, "&")
	}

	return ""
}
