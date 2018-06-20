/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package gitbook

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/opencontrol/compliance-masonry/pkg/lib"
	"github.com/opencontrol/compliance-masonry/tools/fs"
)

type exportControlTest struct {
	opencontrolDir    string
	certificationPath string
	standardKey       string
	controlKey        string
	expectedPath      string
	expectedText      string
}

var exportControlTests = []exportControlTest{
	// Check that a control is exported correctly
	{
		filepath.Join("..", "..", "..", "..", "test", "fixtures", "opencontrol_fixtures"),
		filepath.Join("..", "..", "..", "..", "test", "fixtures", "opencontrol_fixtures", "certifications", "LATO.yaml"),
		"NIST-800-53",
		"CM-2",
		"NIST-800-53-CM-2.md",
		`#NIST-800-53-CM-2
##Baseline Configuration
#### Description
'The organization develops, documents, and maintains under configuration
control, a current baseline configuration of the information system.'

#### Amazon Elastic Compute Cloud

##### Responsible Role: AWS Staff

##### Control Origin: shared

##### a
Justification in narrative form A for CM-2

##### b
Justification in narrative form B for CM-2
Covered By:
* [Amazon Elastic Compute Cloud - EC2 Verification 1](../components/EC2.md)
`,
	},
}

func TestExportControl(t *testing.T) {
	for _, example := range exportControlTests {
		dir, err := ioutil.TempDir("", "example")
		if err != nil {
			log.Fatal(err)
		}
		defer os.RemoveAll(dir)
		openControlData, errs := lib.LoadData(example.opencontrolDir, example.certificationPath)
		if len(errs) > 0 {
			log.Fatal("Should have loaded the opencontrol data.")
		}
		openControl := OpenControlGitBook{
			openControlData,
			"",
			dir,
			fs.OSUtil{},
		}
		standard, _ := openControl.GetStandard(example.standardKey)
		control := standard.GetControl(example.controlKey)
		actualPath, actualText := openControl.exportControl(&ControlGitbook{control, dir, example.standardKey, example.controlKey})
		expectedPath := filepath.Join(dir, example.expectedPath)
		// Verify the expected export path is the same as the actual export path
		if expectedPath != actualPath {
			t.Errorf("Expected %s, Actual: %s", example.expectedPath, actualPath)
		}
		// Verify the expected text is the same as the actual text
		if example.expectedText != strings.Replace(actualText, "\\", "/", -1) {
			t.Errorf("Expected %s, Actual: %s", example.expectedText, actualText)
		}

	}
}
