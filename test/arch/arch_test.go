package arch

import (
	"testing"

	"github.com/mstrYoda/go-arctest/pkg/arctest"
)

const (
	_projectDir = "../../."
)

func TestArchitecture(t *testing.T) {
	architecture, err := arctest.New(_projectDir)

	if err != nil {
		t.Fatalf("Failed to create architecture: %v", err)
	}

	err = architecture.ParsePackages("internal/domain", "internal/adapters", "internal/ports", "internal/application", "cmd")
	if err != nil {
		t.Fatalf("Failed to parse packages: %v", err)
	}

	// Define layers
	cmdLayer, err := arctest.NewLayer("CMD", "^cmd%")
	domainLayer, err := arctest.NewLayer("Domain", "^internal/domain$")
	applicationLayer, err := arctest.NewLayer("Application", "^internal/application$")
	adaptersLayer, err := arctest.NewLayer("Adapters", "^internal/adapters$")
	portsLayer, err := arctest.NewLayer("Ports", "^internal/ports$")

	layeredArchitecture := architecture.NewLayeredArchitecture(
		domainLayer,
		applicationLayer,
		adaptersLayer,
		portsLayer,
		cmdLayer,
	)

	applicationLayer.DependsOnLayer(domainLayer)
	applicationLayer.DependsOnLayer(portsLayer)

	adaptersLayer.DependsOnLayer(portsLayer)
	adaptersLayer.DependsOnLayer(domainLayer)

	cmdLayer.DependsOnLayer(applicationLayer)
	cmdLayer.DependsOnLayer(portsLayer)
	cmdLayer.DependsOnLayer(adaptersLayer)

	// ...

	// Check layered architecture
	violations, err := layeredArchitecture.Check()
	if err != nil {
		t.Fatalf("Failed to check layered architecture: %v", err)
	}

	for _, violation := range violations {
		t.Errorf("Architecture violation: %s", violation)
	}

}
