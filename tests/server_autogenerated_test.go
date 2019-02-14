// Code generated by go generate; DO NOT EDIT.
package tests

import (
	"testing"

	"github.com/haproxytech/config-parser/parsers"
)


func TestServerNormal0(t *testing.T) {
	err := ProcessLine("server name 127.0.0.1:8080", &parsers.Server{})
	if err != nil {
		t.Errorf(err.Error())
	}
}
func TestServerNormal1(t *testing.T) {
	err := ProcessLine("server name 127.0.0.1", &parsers.Server{})
	if err != nil {
		t.Errorf(err.Error())
	}
}
func TestServerNormal2(t *testing.T) {
	err := ProcessLine("server name 127.0.0.1 backup", &parsers.Server{})
	if err != nil {
		t.Errorf(err.Error())
	}
}
func TestServerFail0(t *testing.T) {
	err := ProcessLine("server", &parsers.Server{})
	if err == nil {
		t.Errorf("no data")
	}
}
func TestServerFail1(t *testing.T) {
	err := ProcessLine("---", &parsers.Server{})
	if err == nil {
		t.Errorf("no data")
	}
}
func TestServerFail2(t *testing.T) {
	err := ProcessLine("--- ---", &parsers.Server{})
	if err == nil {
		t.Errorf("no data")
	}
}