/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package vcs_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestVcs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Vcs Suite")
}
