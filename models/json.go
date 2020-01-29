package models

import (
	"encoding/json"
	"reflect"
)

type (
	Catalogs struct {
		// contains the actual data of catalogs
		data map[string]interface{}

		// shared common fields of catalogs
		// used for data quick reference
		commonFields
	}

	commonFields struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}
)

func (c *Catalogs) MarshalJSON() ([]byte, error) {
	c.resetData()

	return json.Marshal(c.data)
}

// Unmarshal parse input bytes to Catalogs
func (c *Catalogs) UnmarshalJSON(input []byte) error {
	if c.data == nil {
		c.data = make(map[string]interface{})
	}

	if err := json.Unmarshal(input, &c.data); err != nil {
		return err
	}

	if err := json.Unmarshal(input, c.commonFields); err != nil {
		return err
	}

	return nil
}

// resetData reset catalog common fields
func (c *Catalogs) resetData() {
	if c.data == nil {
		c.data = make(map[string]interface{})
	}

	// iterate over catalog common fields and assign value to data
	commonFieldsValues := reflect.ValueOf(c.commonFields)
	commonFieldsTypes := commonFieldsValues.Type()
	for i := 0; i < commonFieldsValues.NumField(); i++ {
		tag := commonFieldsTypes.Field(i).Tag.Get("json")
		if tag == "-" || tag == "" {
			continue
		}

		field := commonFieldsValues.Field(i)
		if field.CanInterface() {
			c.data[tag] = field.Interface()
		}
	}
}
