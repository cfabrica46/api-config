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
	noEntriesError = "no entries given"
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
