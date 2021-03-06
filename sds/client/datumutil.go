package client

import (
	"errors"
	"fmt"
	"github.com/XiaoMi/galaxy-sdk-go/sds/table"
)

func BoolDatum(val bool) *table.Datum {
	var typeA1 table.DataType = table.DataType_BOOL
	return &table.Datum{
		TypeA1: &typeA1,
		Value: &table.Value{BoolValue: &val},
	}
}

func Int8Datum(val int8) *table.Datum {
	var typeA1 table.DataType = table.DataType_INT8
	return &table.Datum{
		TypeA1: &typeA1,
		Value: &table.Value{Int8Value: &val},
	}
}

func Int16Datum(val int16) *table.Datum {
	var typeA1 table.DataType = table.DataType_INT16
	return &table.Datum{
		TypeA1: &typeA1,
		Value: &table.Value{Int16Value: &val},
	}
}

func Int32Datum(val int32) *table.Datum {
	var typeA1 table.DataType = table.DataType_INT32
	return &table.Datum{
		TypeA1: &typeA1,
		Value: &table.Value{Int32Value: &val},
	}
}

func Int64Datum(val int64) *table.Datum {
	var typeA1 table.DataType = table.DataType_INT64
	return &table.Datum{
		TypeA1: &typeA1,
		Value: &table.Value{Int64Value: &val},
	}
}

func Float32Datum(val float32) *table.Datum {
	var casted float64 = float64(val)
	var typeA1 table.DataType = table.DataType_FLOAT
	return &table.Datum{
		TypeA1: &typeA1,
		Value: &table.Value{DoubleValue: &casted},
	}
}

func Float64Datum(val float64) *table.Datum {
	var typeA1 table.DataType = table.DataType_DOUBLE
	return &table.Datum{
		TypeA1: &typeA1,
		Value: &table.Value{DoubleValue: &val},
	}
}


func StringDatum(val string) *table.Datum {
	var typeA1 table.DataType = table.DataType_STRING
	return &table.Datum{
		TypeA1: &typeA1,
		Value: &table.Value{StringValue: &val},
	}
}

func BinaryDatum(val []byte) *table.Datum {
	var typeA1 table.DataType = table.DataType_BINARY
	return &table.Datum{
		TypeA1: &typeA1,
		Value: &table.Value{BinaryValue: val},
	}
}

func RawbinaryDatum(val []byte) *table.Datum {
	var typeA1 table.DataType = table.DataType_RAWBINARY
	return &table.Datum{
		TypeA1: &typeA1,
		Value: &table.Value{BinaryValue: val},
	}
}

func BoolValue(datum *table.Datum) bool {
	if datum.GetTypeA1() != table.DataType_BOOL {
		panic(errors.New(fmt.Sprintf("Unexpected datum type: %s", datum.GetTypeA1())))
	}
	return *datum.GetValue().BoolValue
}

func Int8Value(datum *table.Datum) int8 {
	if datum.GetTypeA1() != table.DataType_INT8 {
		panic(errors.New(fmt.Sprintf("Unexpected datum type: %s", datum.GetTypeA1())))
	}
	return *datum.GetValue().Int8Value
}

func Int16Value(datum *table.Datum) int16 {
	if datum.GetTypeA1() != table.DataType_INT16 {
		panic(errors.New(fmt.Sprintf("Unexpected datum type: %s", datum.GetTypeA1())))
	}
	return *datum.GetValue().Int16Value
}

func Int32Value(datum *table.Datum) int32 {
	if datum.GetTypeA1() != table.DataType_INT32 {
		panic(errors.New(fmt.Sprintf("Unexpected datum type: %s", datum.GetTypeA1())))
	}
	return *datum.GetValue().Int32Value
}

func Int64Value(datum *table.Datum) int64 {
	if datum.GetTypeA1() != table.DataType_INT64 {
		panic(errors.New(fmt.Sprintf("Unexpected datum type: %s", datum.GetTypeA1())))
	}
	return *datum.GetValue().Int64Value
}

func Float32Value(datum *table.Datum) float32 {
	if datum.GetTypeA1() != table.DataType_FLOAT {
		panic(errors.New(fmt.Sprintf("Unexpected datum type: %s", datum.GetTypeA1())))
	}
	return float32(*datum.GetValue().DoubleValue)
}

func Float64Value(datum *table.Datum) float64 {
	if datum.GetTypeA1() != table.DataType_DOUBLE {
		panic(errors.New(fmt.Sprintf("Unexpected datum type: %s", datum.GetTypeA1())))
	}
	return *datum.GetValue().DoubleValue
}

func StringValue(datum *table.Datum) string {
	if datum.GetTypeA1() != table.DataType_STRING {
		panic(errors.New(fmt.Sprintf("Unexpected datum type: %s", datum.GetTypeA1())))
	}
	return *datum.GetValue().StringValue
}

func BinaryValue(datum *table.Datum) []byte {
	if datum.GetTypeA1() != table.DataType_BINARY {
		panic(errors.New(fmt.Sprintf("Unexpected datum type: %s", datum.GetTypeA1())))
	}
	return datum.GetValue().BinaryValue
}

