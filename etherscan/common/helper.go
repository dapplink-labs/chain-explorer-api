package common

import (
	"math/big"
	"reflect"
	"strconv"
	"time"
)

func Compose(param map[string]interface{}, tag string, input interface{}) {
	if input == nil {
		return
	}

	v := reflect.ValueOf(input)
	switch v.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Interface:
		if v.IsNil() {
			return
		}
	}

	param[tag] = input
}

type M map[string]interface{}

type BigInt big.Int

func (b *BigInt) UnmarshalText(text []byte) (err error) {
	var bigInt = new(big.Int)
	err = bigInt.UnmarshalText(text)
	if err != nil {
		return
	}

	*b = BigInt(*bigInt)
	return nil
}

func (b *BigInt) MarshalText() (text []byte, err error) {
	return []byte(b.Int().String()), nil
}

func (b *BigInt) Int() *big.Int {
	return (*big.Int)(b)
}

type Time time.Time

func (t *Time) UnmarshalText(text []byte) (err error) {
	input, err := strconv.ParseInt(string(text), 10, 64)
	if err != nil {
		err = WrapErr(err, "strconv.ParseInt")
		return
	}

	var timestamp = time.Unix(input, 0)
	*t = Time(timestamp)

	return nil
}

func (t Time) Time() time.Time {
	return time.Time(t)
}

func (t Time) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatInt(t.Time().Unix(), 10)), nil
}
