package models_test

import (
	"encoding/json"
	"testing"

	"github.com/isaactl/HelloWorld/models"
)

func TestCatalogs_Marshal(t *testing.T) {
	c := models.Catalogs{}
	c.ID = "id_test"
	c.Name = "name_test"
	c.Description = "des_test"

	b, err := json.Marshal(&c)
	if err != nil {
		t.Fatalf("marshal catalogs failed: %v", err)
	}

	t.Logf("%q", b)
}
