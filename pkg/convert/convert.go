package convert

import (
	"bytes"
	"fmt"
)

type ConversionType string

const (
	Hex     ConversionType = "hex"
	Decimal ConversionType = "decimal"
	Octal   ConversionType = "octal"
	Binary  ConversionType = "binary"
)

type Array interface {
	ConvertToArray() (string, error)
}

type ArrayImpl struct {
	DataType     ConversionType
	VariableType string
	Name         string
	Data         []byte
}

func (a *ArrayImpl) ConvertToArray() (string, error) {
	switch a.DataType {
	case Hex:
		return a.convertToHexArray(), nil
	case Decimal:
		return a.convertToDecimalArray(), nil
	case Octal:
		return a.convertToOctalArray(), nil
	case Binary:
		return a.convertToBinaryArray(), nil
	default:
		return "", fmt.Errorf("unsupported conversion type: %s", a.DataType)
	}
}

func (a *ArrayImpl) convertToHexArray() string {
	cStyleHexArray := a.getVariableDeclaration()
	cStyleHexArray += a.getVariableInitialization(func(b byte) string {
		return fmt.Sprintf("0x%02x", b)
	})
	cStyleHexArray += "};"
	return cStyleHexArray
}

func (a *ArrayImpl) convertToDecimalArray() string {
	cStyleDecimalArray := a.VariableType + " " + a.Name + "[" + fmt.Sprintf("%d", len(a.Data)) + "] = {\n"
	cStyleDecimalArray += a.getVariableInitialization(func(b byte) string {
		return fmt.Sprintf("%d", b)
	})
	cStyleDecimalArray += "};"
	return cStyleDecimalArray
}

func (a *ArrayImpl) convertToOctalArray() string {
	cStyleOctalArray := a.VariableType + " " + a.Name + "[" + fmt.Sprintf("%d", len(a.Data)) + "] = {\n"
	cStyleOctalArray += a.getVariableInitialization(func(b byte) string {
		return fmt.Sprintf("%#o", b)
	})
	cStyleOctalArray += "};"
	return cStyleOctalArray
}

func (a *ArrayImpl) convertToBinaryArray() string {
	cStyleBinaryArray := a.VariableType + " " + a.Name + "[" + fmt.Sprintf("%d", len(a.Data)) + "] = {\n"
	cStyleBinaryArray += a.getVariableInitialization(func(b byte) string {
		return fmt.Sprintf("%#b", b)
	})
	cStyleBinaryArray += "};"
	return cStyleBinaryArray
}

func (a *ArrayImpl) getVariableDeclaration() string {
	return a.VariableType + " " + a.Name + "[" + fmt.Sprintf("%d", len(a.Data)) + "] = {\n"
}

func (a *ArrayImpl) getVariableInitialization(formatString func(byte) string) string {
	var buffer bytes.Buffer
	for i, b := range a.Data {
		buffer.WriteString(formatString(b))
		if i < len(a.Data)-1 {
			buffer.WriteString(", ")
		}
		if i%16 == 15 {
			buffer.WriteString("\n")
		}
	}
	return buffer.String()
}
