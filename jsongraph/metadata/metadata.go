package metadata

import (
	"encoding/json"
	"reflect"
)

// Metadata is a basic map that can be used by any graph object
type Metadata struct {
	Elements []MetadataElement
}
type MetadataElement struct {
	Name    string
	Value   string
	IsValue bool

	// Handle generic interface
	InterfaceValue any
	IsInterface    bool

	IntValue int32
	IsInt    bool

	BoolValue bool
	IsBool    bool
}

// AddElement adds an element to the metadata elements list
// This can be used in the API or in the json Unmarshall function
func (m *Metadata) AddElement(name string, raw any) {
	element := MetadataElement{Name: name}
	value, ok := raw.(string)
	if ok {
		element.Value = value
		element.IsValue = true
		m.Elements = append(m.Elements, element)
		return
	}
	intValue, ok := raw.(int32)
	if ok {
		element.IntValue = intValue
		element.IsInt = true
		m.Elements = append(m.Elements, element)
		return
	}
	boolValue, ok := raw.(bool)
	if ok {
		element.BoolValue = boolValue
		element.IsBool = true
		m.Elements = append(m.Elements, element)
		return
	}
	element.InterfaceValue = raw
	element.IsInterface = true
	m.Elements = append(m.Elements, element)
}

func (m *Metadata) UnmarshalJSON(data []byte) error {
	m.Elements = []MetadataElement{}
	raw := make(map[string]interface{})
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}
	for k, r := range raw {
		m.AddElement(k, r)
	}
	return nil
}

func (m Metadata) MarshalJSON() ([]byte, error) {

	parsed := map[string]any{}

	for _, element := range m.Elements {
		if element.IsValue {
			parsed[element.Name] = element.Value
		} else if element.IsInt {
			parsed[element.Name] = element.IntValue
		} else if element.IsBool {
			parsed[element.Name] = element.BoolValue
		} else {
			parsed[element.Name] = unwrap(element.InterfaceValue)
		}
	}
	// JSON encoding is done the same way as before
	// returnes bytes, err
	return json.Marshal(parsed)
}

// Unwrap an interface into its proper data
func unwrap(data interface{}) interface{} {
	d := reflect.ValueOf(data)
	if reflect.ValueOf(data).Kind() == reflect.Slice {
		returnSlice := make([]interface{}, d.Len())
		for i := 0; i < d.Len(); i++ {
			returnSlice[i] = unwrap(d.Index(i).Interface())
		}
		return returnSlice
	} else if reflect.ValueOf(data).Kind() == reflect.Map {
		tmpData := make(map[string]interface{})
		for _, k := range d.MapKeys() {
			tmpData[k.String()] = unwrap(d.MapIndex(k).Interface())
		}
		return tmpData
	} else {
		return data
	}
}
