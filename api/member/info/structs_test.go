package info

import (
	"encoding/json"
	"testing"
)

// TestUnmarshalResponse tests that the response struct can be properly unmarshaled from JSON
func TestUnmarshalResponse(t *testing.T) {
	jsonData := `{}` // Empty JSON for basic structure test

	var response MemberInfoResponse
	err := json.Unmarshal([]byte(jsonData), &response)

	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
}

// TestMarshalResponse tests that the response struct can be properly marshaled to JSON
func TestMarshalResponse(t *testing.T) {
	response := MemberInfoResponse{}

	_, err := json.Marshal(response)

	if err != nil {
		t.Fatalf("Failed to marshal response: %v", err)
	}
}
