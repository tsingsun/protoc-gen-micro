package internal_test

import (
	"testing"
	"github.com/tsingsun/protoc-gen-goql/generator/internal"
	"fmt"
)

func TestNewTagItems(t *testing.T) {
	a := `@inject_tag: valid:"ip" gorm:"foreignkey:UserRefer" json:"name"`

	exp:= internal.NewTagItems(a)
	fmt.Println(exp)
}

func TestNewTagItemsNoinjectTag(t *testing.T) {
	a := `@inject_tag: gorm:"foreignkey:manager_id"`

	exp:= internal.NewTagItems(a)
	fmt.Println(exp)
}

