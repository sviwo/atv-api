package model

import (
	"strconv"
	"sviwo/internal/consts"
	"time"
)

// 数据类型列表

type DataTypeOutput struct {
	BaseType      []DataTypeValueBase      `json:"baseType" dc:"基础类型"`
	ExtensionType []DataTypeValueExtension `json:"extensionType" dc:"扩展类型"`
}

type DataTypeValueBase struct {
	Title string `json:"title" dc:"类型名称"`
	Type  string `json:"type" dc:"数据类型"`
	TSLParamBase
}

type DataTypeValueExtension struct {
	Title string `json:"title" dc:"类型名称"`
	Type  string `json:"type" dc:"数据类型"`
	TSLParamExtension
}

// 参数值（类型、类型参数）
type TSLValueType struct {
	Type     string `json:"type" dc:"数据类型" v:"required#请选择数据类型"` // 类型
	TSLParam        // 参数
}

func (t TSLValueType) ConvertValue(v interface{}) interface{} {
	var transfer Transfer
	switch t.Type {
	case consts.TypeInt:
		transfer = TInt(t.TSLParam)
	case consts.TypeLong:
		transfer = TLong(t.TSLParam)
	case consts.TypeFloat:
		transfer = TFloat(t.TSLParam)
	case consts.TypeDouble:
		transfer = TDouble(t.TSLParam)
	case consts.TypeText:
		transfer = TText(t.TSLParam)
	case consts.TypeBool:
		transfer = TBoolean(t.TSLParam)
	case consts.TypeDate:
		transfer = TDate(t.TSLParam)
	case consts.TypeEnum:
		transfer = TEnum(t.TSLParam)
	case consts.TypeArray:
		transfer = TArray(t.TSLParam)
	case consts.TypeObject:
		transfer = TObject(t.TSLParam)
	default:
		return nil
	}
	return transfer.Convert(v)
}

type Transfer interface {
	Convert(interface{}) interface{}
}

type TInt TSLParam

func (tInt TInt) Convert(v interface{}) interface{} {
	number, ok := v.(int)
	if !ok {
		floatNumber, floatNumberOk := v.(float64)
		if !floatNumberOk {
			return 0
		}
		number = int(floatNumber)
	}
	if tInt.TSLParamBase.Min != nil && *tInt.TSLParamBase.Min > number {
		return *tInt.TSLParamBase.Min
	}
	if tInt.TSLParamBase.Max != nil && *tInt.TSLParamBase.Max > number {
		return *tInt.TSLParamBase.Max
	}
	return number
}

type TLong TSLParam

func (tLong TLong) Convert(v interface{}) interface{} {
	number, ok := v.(float64)
	if !ok {
		return 0
	}
	if tLong.TSLParamBase.Min != nil && float64(*tLong.TSLParamBase.Min) > number {
		return *tLong.TSLParamBase.Min
	}
	if tLong.TSLParamBase.Max != nil && float64(*tLong.TSLParamBase.Max) > number {
		return *tLong.TSLParamBase.Max
	}
	return number
}

type TFloat TSLParam

func (tFloat TFloat) Convert(v interface{}) interface{} {
	number, ok := v.(float64)
	if !ok {
		return 0
	}
	if tFloat.TSLParamBase.Min != nil && float32(*tFloat.TSLParamBase.Min) > float32(number) {
		number = float64(*tFloat.TSLParamBase.Min)
	}
	if tFloat.TSLParamBase.Max != nil && float32(*tFloat.TSLParamBase.Max) > float32(number) {
		number = float64(*tFloat.TSLParamBase.Max)
	}
	defaultDecimal := 2
	if tFloat.TSLParamBase.Decimals != nil {
		defaultDecimal = *tFloat.TSLParamBase.Decimals
	}
	number32, _ := strconv.ParseFloat(strconv.FormatFloat(number, 'f', defaultDecimal, 64), 32)
	return float32(number32)
}

type TDouble TSLParam

func (tDouble TDouble) Convert(v interface{}) interface{} {
	number, ok := v.(float64)
	if !ok {
		return 0
	}
	if tDouble.TSLParamBase.Min != nil && float64(*tDouble.TSLParamBase.Min) > number {
		number = float64(*tDouble.TSLParamBase.Min)
	}
	if tDouble.TSLParamBase.Max != nil && float64(*tDouble.TSLParamBase.Max) > number {
		number = float64(*tDouble.TSLParamBase.Max)
	}
	defaultDecimal := 2
	if tDouble.TSLParamBase.Decimals != nil {
		defaultDecimal = *tDouble.TSLParamBase.Decimals
	}
	number64, _ := strconv.ParseFloat(strconv.FormatFloat(number, 'f', defaultDecimal, 64), 32)
	return number64
}

type TText TSLParam

func (tText TText) Convert(v interface{}) interface{} {
	text, ok := v.(string)
	if !ok {
		return ""
	}
	if tText.MaxLength != nil && *tText.MaxLength > 0 && len(text) > *tText.MaxLength {
		return text[:*tText.MaxLength-1]
	} else {
		return text
	}
}

type TBoolean TSLParam

func (tBoolean TBoolean) Convert(v interface{}) interface{} {
	b, ok := v.(bool)
	if !ok {
		return ""
	}
	//TODO 这里是返回文字还是类型,暂定是返回映射文字
	if !b && tBoolean.FalseText != nil {
		return *tBoolean.FalseText
	} else if b && tBoolean.TrueText != nil {
		return *tBoolean.TrueText
	}
	return ""
}

type TDate TSLParam

const layout = "2006-01-02 15:04:05"

func (TDate TDate) Convert(v interface{}) interface{} {
	str, ok := v.(string)
	if !ok {
		return time.Time{}
	}
	t, err := time.Parse(layout, str)
	if err != nil {
		return time.Time{}
	} else {
		return t
	}
}

type TEnum TSLParam

func (tEnum TEnum) Convert(v interface{}) interface{} {
	tE, ok := v.(string)
	if !ok {
		return ""
	}
	if tEnum.TSLParamExtension.Elements == nil {
		return ""
	}
	for _, node := range tEnum.TSLParamExtension.Elements {
		if node.Text == tE {
			return node.Value
		}
	}
	return ""
}

type TArray TSLParam

func (tArray TArray) Convert(v interface{}) interface{} {
	tA, ok := v.([]interface{})
	if !ok {
		return nil
	}
	if tArray.TSLParamExtension.ElementType == nil {
		return nil
	} else {
		result := make([]interface{}, 0)
		for _, node := range tA {
			result = append(result, (*tArray.TSLParamExtension.ElementType).ConvertValue(node))
		}
		return result
	}
}

type TObject TSLParam

func (tObject TObject) Convert(v interface{}) interface{} {
	m, ok := v.(map[string]interface{})
	if !ok {
		return nil
	}
	result := make(map[string]interface{})
	if tObject.TSLParamExtension.Properties != nil {
		for k, value := range m {
			for _, t := range tObject.TSLParamExtension.Properties {
				if t.Name == k {
					result[t.Name] = t.ValueType.ConvertValue(value)
				}
			}
		}
	}
	return result
}
