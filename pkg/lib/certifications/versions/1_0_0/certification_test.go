/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package certification_test

import (
	"github.com/opencontrol/compliance-masonry/pkg/lib/certifications/versions/1_0_0"
	"testing"
)

type standardOrderTest struct {
	certification certification.Certification
	expectedOrder string
}

var standardOrderTests = []standardOrderTest{
	{
		// Verify Natural sort order
		certification.Certification{Standards: map[string]map[string]interface{}{
			"A": {"3": nil, "2": nil, "1": nil},
			"B": {"12": nil, "2": nil, "1": nil},
			"C": {"2": nil, "11": nil, "101": nil, "1000": nil, "100": nil, "10": nil, "1": nil},
		}},
		"A1A2A3B1B2B12C1C2C10C11C100C101C1000",
	},
	{
		// Check that data is returned in order given letters and numbers
		certification.Certification{Standards: map[string]map[string]interface{}{
			"1":  {"3": nil, "2": nil, "1": nil},
			"B":  {"3": nil, "2": nil, "1": nil},
			"B2": {"3": nil, "2": nil, "1": nil},
		}},
		"111213B1B2B3B21B22B23",
	},
}

func TestStandardOrder(t *testing.T) {
	for _, example := range standardOrderTests {
		actualOrder := ""
		standardKeys := example.certification.GetSortedStandards()
		for _, standardKey := range standardKeys {
			controlKeys := example.certification.GetControlKeysFor(standardKey)
			for _, controlKey := range controlKeys {
				actualOrder += standardKey + controlKey
			}
		}
		// Verify that the actual order is the expected order
		if actualOrder != example.expectedOrder {
			t.Errorf("Expected %s, Actual: %s", example.expectedOrder, actualOrder)
		}
	}
}

func TestGetKey(t *testing.T) {
	cert := certification.Certification{Key: "test"}
	if cert.GetKey() != "test" {
		t.Errorf("GetKey expected test. Actual %s", cert.GetKey())
	}
}
