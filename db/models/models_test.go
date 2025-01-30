package models

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Overshadowing(t *testing.T) {
	reflect.TypeOf((*Models)(nil)).Elem()
	x := (*Models)(nil)
	asd := reflect.TypeOf(x).Elem()
	fmt.Println(asd)
	Models := NewModels(nil)
	fmt.Println(reflect.TypeOf(Models))
	if 1 != 2-1 {
		t.Error("chupa")
	}
}
