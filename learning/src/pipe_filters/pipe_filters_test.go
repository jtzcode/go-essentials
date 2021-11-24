package pipefilters

import (
	"errors"
	"strings"
	"testing"
)

type Request interface{}
type Response interface{}

type Filter interface {
	Process(data Request) (Response, error)
}
type SplitFilter struct {
	delimiter string
}

var SplitFilterWrongFormatError = errors.New("Input data should be string")

func CreateSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter}
}

func (sf *SplitFilter) Process(data Request) (Response, error) {
	str, ok := data.(string)
	if !ok {
		return nil, SplitFilterWrongFormatError
	}
	splits := strings.Split(str, sf.delimiter)
	return splits, nil
}

func TestSplitFilters(t *testing.T) {
	splitFilter := CreateSplitFilter(",")
	//ret, err := splitFilter.Process("1,2,3,4,5")
	ret, err := splitFilter.Process(100)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}
