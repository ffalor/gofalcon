package models

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestSensorUpdateSettingsReqV2_BuildFieldIncludedInJSON(t *testing.T) {
	// Test that Build field is always included in JSON output, even when empty
	settings := &SensorUpdateSettingsReqV2{
		Build:     "", // Empty string should still be included in JSON
		Scheduler: &PolicySensorUpdateScheduler{},
		Variants:  []*SensorUpdateBuildReqV1{},
	}

	jsonData, err := json.Marshal(settings)
	if err != nil {
		t.Fatalf("Failed to marshal SensorUpdateSettingsReqV2: %v", err)
	}

	jsonString := string(jsonData)
	
	// Check that "build" field is present in JSON output
	if !strings.Contains(jsonString, `"build":`) {
		t.Errorf("Expected 'build' field to be present in JSON output, but it was missing. JSON: %s", jsonString)
	}

	// Verify that empty string is marshaled as empty string (not omitted)
	if !strings.Contains(jsonString, `"build":""`) {
		t.Errorf("Expected 'build' field to be empty string in JSON output. JSON: %s", jsonString)
	}

	// Test with a non-empty build value
	settings.Build = "version-1.2.3"
	jsonData2, err := json.Marshal(settings)
	if err != nil {
		t.Fatalf("Failed to marshal SensorUpdateSettingsReqV2 with build value: %v", err)
	}

	jsonString2 := string(jsonData2)
	if !strings.Contains(jsonString2, `"build":"version-1.2.3"`) {
		t.Errorf("Expected 'build' field to contain 'version-1.2.3' in JSON output. JSON: %s", jsonString2)
	}
}