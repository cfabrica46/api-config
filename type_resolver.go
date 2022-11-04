package apiconfig

// VariableType tipo de dato para tipo de variable
type VariableType int

// VariableTypeResolver interface
type VariableTypeResolver interface {
	ResolveType(any) (VariableType, error)
}
