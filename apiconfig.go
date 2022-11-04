package apiconfig

import (
	"errors"
	"time"
)

// ConfigEntry entrada de configuración.
type ConfigEntry struct {
	DefaultValue any
	VariableName string
	Description  string
	Shortcut     string
}

// CfgBase configuración basica de la API.
type CfgBase struct {
	Port      string
	URIPrefix string
	Timeout   time.Duration
	// LoggingLevel string
	/* EnabledTracing bool
	EnabledMetrics bool */
}

// Configurator Interfaz configurador de apiconfig.
type Configurator interface {
	Configure(entries []ConfigEntry) (map[string]any, error)
}

type configurator struct {
	flagConfigurator FlagConfigurator
	typeResolver     VariableTypeResolver
}

const (
	// configError = "Error reading ConfigFile".

	// flagError = "Error with FlagConfigurator".

	noEntriesError = "no entries given"

	// valueError = "No value given to the variable through config file, environment variables or flag".

	// typeResolverError = "Error on ResolveType".
)

// NewConfigurator constructor.
func NewConfigurator(flagConfigurator FlagConfigurator, typeResolver VariableTypeResolver) Configurator {
	return &configurator{
		flagConfigurator: flagConfigurator,
		typeResolver:     typeResolver,
	}
}

func (c *configurator) Configure(entries []ConfigEntry) (map[string]any, error) {
	if len(entries) == 0 {
		return nil, errors.New(noEntriesError)
	}

	return nil, nil
}
