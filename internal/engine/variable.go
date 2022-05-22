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
		break
	case BaseTypeString:
		if inputBaseType == BaseTypeString {
			return true
		}
		break
	case BaseTypeInteger:
		if inputBaseType == BaseTypeInteger {
			return true
		}
		break
	case BaseTypeDecimal:
		if inputBaseType == BaseTypeDecimal {
			return true
		}
		break
	case BaseTypeTime:
		if inputBaseType == BaseTypeTime {
			return true
		}
		break
	}

	return false
}
