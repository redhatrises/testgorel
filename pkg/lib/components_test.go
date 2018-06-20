/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package lib

import (
	"path/filepath"
	"testing"

	"github.com/opencontrol/compliance-masonry/pkg/lib/common/mocks"
	"github.com/opencontrol/compliance-masonry/pkg/lib/result"
	"github.com/stretchr/testify/assert"
)

func TestAddComponent(t *testing.T) {
	// Setup map
	m := newComponents()
	// Get nil component.
	component, found := m.get("test")
	assert.False(t, found)
	assert.Nil(t, component)
	// Create mock component
	newComponent := new(mocks.Component)
	newComponent.On("GetKey").Return("test")
	// Test add method
	m.add(newComponent)
	// Try to retrieve the component again.
	component, found = m.get("test")
	assert.True(t, found)
	assert.Equal(t, component.GetKey(), "test")
}

func TestCompareAndAddComponent(t *testing.T) {
	m := newComponents()
	// Get nil component.
	component, found := m.get("test")
	assert.False(t, found)
	assert.Nil(t, component)
	// Create mock component
	newComponent := new(mocks.Component)
	newComponent.On("GetKey").Return("test")
	// Use compare and add initially.
	added := m.compareAndAdd(newComponent)
	assert.True(t, added)
	// Use compare and add again to show failure.
	added = m.compareAndAdd(newComponent)
	assert.False(t, added)
}

func TestLoadSameComponentTwice(t *testing.T) {
	ws := localWorkspace{components: newComponents(), justifications: result.NewJustifications()}
	componentPath := filepath.Join("..", "..", "test", "fixtures", "component_fixtures", "v3_1_0", "EC2")
	err := ws.LoadComponent(componentPath)
	// Should load the component without a problem.
	assert.Nil(t, err)
	actualComponent, found := ws.components.get("EC2")
	assert.True(t, found)
	assert.NotNil(t, actualComponent)
	// Try to load component again.
	err = ws.LoadComponent(componentPath)
	// Should return an error that this component was already loaded.
	assert.NotNil(t, err)
	assert.Equal(t, "component: EC2 exists", err.Error())
}

func TestBadLoadComponent(t *testing.T) {
	ws := localWorkspace{}
	err := ws.LoadComponent("fake.file")
	// Should return an error because it can't load the file.
	assert.Equal(t, "Component files does not exist", err.Error())
}
