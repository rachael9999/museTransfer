package test

import (
	"bytes"
	"cloud-disk/core/models"
	"encoding/json"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func TestXorm(t *testing.T) {
	engine, err := xorm.NewEngine("mysql", "root:gloves4869@tcp(localhost:3306)/clouddisk?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		t.Fatal(err)
	}

	data := make([]*models.User, 0)
	err = engine.Find(&data)

	if err != nil {
		t.Fatal(err)
	}

	// Convert to JSON
	b, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}

	// Pretty print
	dst := new(bytes.Buffer)
	error := json.Indent(dst, b, "", "    ")
	if error != nil {
		t.Fatal(error)
	}
	fmt.Println(dst.String())
}