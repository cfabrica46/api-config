package apiconfig

// FlagConfigurator Interfaz que define configuración de flags
type FlagConfigurator interface {
	ConfigureFlag(ConfigEntry) error
}
