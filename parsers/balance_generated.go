// Code generated by go generate; DO NOT EDIT.
/*
Copyright 2019 HAProxy Technologies

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package parsers

import (
	"github.com/haproxytech/config-parser/v2/common"
	"github.com/haproxytech/config-parser/v2/errors"
	"github.com/haproxytech/config-parser/v2/types"
)

func (p *Balance) Init() {
    p.data = nil
}

func (p *Balance) GetParserName() string {
	return "balance"
}

func (p *Balance) Get(createIfNotExist bool) (common.ParserData, error) {
	if p.data == nil {
		if createIfNotExist {
			p.data = &types.Balance{}
			return p.data, nil
		}
		return nil, errors.ErrFetch
	}
	return p.data, nil
}

func (p *Balance) GetOne(index int) (common.ParserData, error) {
	if index > 0 {
		return nil, errors.ErrFetch
	}
	if p.data == nil {
		return nil, errors.ErrFetch
	}
	return p.data, nil
}

func (p *Balance) Delete(index int) error {
	p.Init()
	return nil
}

func (p *Balance) Insert(data common.ParserData, index int) error {
	return p.Set(data, index)
}

func (p *Balance) Set(data common.ParserData, index int) error {
	if data == nil {
		p.Init()
		return nil
	}
	switch newValue := data.(type) {
	case *types.Balance:
		p.data = newValue
	case types.Balance:
		p.data = &newValue
	default:
		return errors.ErrInvalidData
	}
	return nil
}
