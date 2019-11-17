package main

import (
	"encoding/json"
	"testing"
)

func TestCatalogs_Marshal(t *testing.T) {
	c := Catalogs{}
	c.ID = "id_test"
	c.Name = "name_test"
	c.Description = "des_test"

	b, err := json.Marshal(&c)
	if err != nil {
		t.Fatalf("marshal catalogs failed: %v", err)
	}

	t.Logf("%+v", c.data)
	t.Logf("%q", b)
}
