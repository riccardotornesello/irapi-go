package get

import (
	"encoding/json"
	"testing"
)

// TestUnmarshalParameters tests that the parameters struct can be properly unmarshaled from JSON
func TestUnmarshalParameters(t *testing.T) {
	jsonData := `{}` // Empty JSON for basic structure test

	var params MemberGetParams
	err := json.Unmarshal([]byte(jsonData), &params)

	if err != nil {
		t.Fatalf("Failed to unmarshal parameters: %v", err)
	}
}

// TestMarshalParameters tests that the parameters struct can be properly marshaled to JSON
func TestMarshalParameters(t *testing.T) {
	params := MemberGetParams{}

	_, err := json.Marshal(params)

	if err != nil {
		t.Fatalf("Failed to marshal parameters: %v", err)
	}
}

// TestUnmarshalResponse tests that the response struct can be properly unmarshaled from JSON
func TestUnmarshalResponse(t *testing.T) {
	jsonData := `{}` // Empty JSON for basic structure test

	var response MemberGetResponse
	err := json.Unmarshal([]byte(jsonData), &response)

	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
}

// TestMarshalResponse tests that the response struct can be properly marshaled to JSON
func TestMarshalResponse(t *testing.T) {
	response := MemberGetResponse{}

	_, err := json.Marshal(response)

	if err != nil {
		t.Fatalf("Failed to marshal response: %v", err)
	}
}
