package msgpack

import (
	"reflect"
	"time"
)

var (
	timeType = reflect.TypeOf((*time.Time)(nil)).Elem()
)

func init() {
	Register(timeType, encodeTime, decodeTime)
}

func (e *Encoder) EncodeTime(tm time.Time) error {
	if err := e.EncodeInt64(tm.Unix()); err != nil {
		return err
	}
	return e.EncodeInt(tm.Nanosecond())
}

func (d *Decoder) DecodeTime() (time.Time, error) {
	sec, err := d.DecodeInt64()
	if err != nil {
		return time.Time{}, err
	}
	nsec, err := d.DecodeInt64()
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(sec, nsec), nil
}

func encodeTime(e *Encoder, v reflect.Value) error {
	tm := v.Interface().(time.Time)
	return e.EncodeTime(tm)
}

func decodeTime(d *Decoder, v reflect.Value) error {
	tm, err := d.DecodeTime()
	if err != nil {
		return err
	}
	v.Set(reflect.ValueOf(tm))
	return nil
}
