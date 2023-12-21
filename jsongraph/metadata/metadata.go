package metadata

import "encoding/json"

// Metadata is a basic map that can be used by any graph object
type Metadata struct {
	Elements []MetadataElement
}
type MetadataElement struct {
	Name    string
	Value   string
	IsValue bool

	IntValue int32
	IsInt    bool

	BoolValue bool
	IsBool    bool
}

func (m *Metadata) UnmarshalJSON(data []byte) error {
	m.Elements = []MetadataElement{}
	raw := make(map[string]interface{})
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}
	for k, r := range raw {
		element := MetadataElement{Name: k}
		value, ok := r.(string)
		if ok {
			element.Value = value
			element.IsValue = true
			m.Elements = append(m.Elements, element)
			continue
		}
		intValue, ok := r.(int32)
		if ok {
			element.IntValue = intValue
			element.IsInt = true
			m.Elements = append(m.Elements, element)
			continue
		}
		boolValue, ok := r.(bool)
		if ok {
			element.BoolValue = boolValue
			element.IsBool = true
			m.Elements = append(m.Elements, element)
			continue
		}
	}
	return nil
}

func (m *Metadata) MarshalJSON() ([]byte, error) {

	parsed := map[string]any{}

	for _, element := range m.Elements {
		if element.IsValue {
			parsed[element.Name] = element.Value
		} else if element.IsInt {
			parsed[element.Name] = element.IntValue
		} else if element.IsBool {
			parsed[element.Name] = element.BoolValue
		}
	}
	// JSON encoding is done the same way as before
	// returnes bytes, err
	return json.Marshal(parsed)
}
