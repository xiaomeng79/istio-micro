package utils

import (
	"errors"
	"reflect"
	"strconv"

	"github.com/xiaomeng79/istio-micro/cinit"

	"github.com/mitchellh/mapstructure"
	"github.com/xiaomeng79/go-utils/math"
)

func S2ID(idS string) (int64, error) {
	if idS == "" {
		return 0, errors.New("id不能为空")
	}
	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		return 0, err
	}
	if id <= 0 {
		return 0, errors.New("id不能小于0")
	}
	return id, nil
}

func S2N(idS string) (int64, error) {
	if idS == "" {
		return 0, nil
	}
	num, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func S2F64(idS string) (float64, error) {
	if idS == "" {
		return 0, nil
	}
	num, err := strconv.ParseFloat(idS, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func S2I32(idS string) (int32, error) {
	if idS == "" {
		return 0, nil
	}
	num, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		return 0, err
	}
	return int32(num), nil
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func OddsCompute(o1, o2 float64) bool {
	return math.Round(o1, cinit.FloatComputeBit) == math.Round(o2, cinit.FloatComputeBit)
}

func Map2Struct(input, result interface{}) {
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &result,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		panic(err)
	}

	err = decoder.Decode(input)
	if err != nil {
		panic(err)
	}
}
