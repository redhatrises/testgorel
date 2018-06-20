/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package diff_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDiff(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Diff Suite")
}
