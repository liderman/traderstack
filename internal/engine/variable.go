package engine

type Argument struct {
	Id           string
	Name         string
	Desc         string
	BaseType     string
	ExtendedType string
	Required     bool
	Value        interface{}
}

type SetArgument struct {
	Id    string
	Value interface{}
}

const (
	BaseTypeString  = "string"
	BaseTypeInteger = "integer"
	BaseTypeBoolean = "boolean"
	BaseTypeDecimal = "decimal"
	BaseTypeTime    = "time"
	BaseTypeNumeric = "numeric"
)

func (a *Argument) CheckInputBaseType(inputBaseType string) bool {
	switch a.BaseType {
	case BaseTypeNumeric:
		if inputBaseType == BaseTypeNumeric || inputBaseType == BaseTypeInteger || inputBaseType == BaseTypeDecimal {
			return true
		}
	case BaseTypeString:
		if inputBaseType == BaseTypeString {
			return true
		}
	case BaseTypeInteger:
		if inputBaseType == BaseTypeInteger {
			return true
		}
	case BaseTypeDecimal:
		if inputBaseType == BaseTypeDecimal {
			return true
		}
	case BaseTypeTime:
		if inputBaseType == BaseTypeTime {
			return true
		}
	}

	return false
}
