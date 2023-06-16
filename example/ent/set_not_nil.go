// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	uuid "github.com/gofrs/uuid/v5"
)

// set field if value's pointer is not nil.
func (s *StudentUpdate) SetNotNilUpdatedAt(value *time.Time) *StudentUpdate {
	if value != nil {
		return s.SetUpdatedAt(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdateOne) SetNotNilUpdatedAt(value *time.Time) *StudentUpdateOne {
	if value != nil {
		return s.SetUpdatedAt(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentCreate) SetNotNilUpdatedAt(value *time.Time) *StudentCreate {
	if value != nil {
		return s.SetUpdatedAt(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdate) SetNotNilName(value *string) *StudentUpdate {
	if value != nil {
		return s.SetName(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdateOne) SetNotNilName(value *string) *StudentUpdateOne {
	if value != nil {
		return s.SetName(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentCreate) SetNotNilName(value *string) *StudentCreate {
	if value != nil {
		return s.SetName(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdate) SetNotNilAge(value *int) *StudentUpdate {
	if value != nil {
		return s.SetAge(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdateOne) SetNotNilAge(value *int) *StudentUpdateOne {
	if value != nil {
		return s.SetAge(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentCreate) SetNotNilAge(value *int) *StudentCreate {
	if value != nil {
		return s.SetAge(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdate) SetNotNilAgeInt8(value *int8) *StudentUpdate {
	if value != nil {
		return s.SetAgeInt8(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdateOne) SetNotNilAgeInt8(value *int8) *StudentUpdateOne {
	if value != nil {
		return s.SetAgeInt8(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentCreate) SetNotNilAgeInt8(value *int8) *StudentCreate {
	if value != nil {
		return s.SetAgeInt8(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdate) SetNotNilAgeUint8(value *uint8) *StudentUpdate {
	if value != nil {
		return s.SetAgeUint8(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdateOne) SetNotNilAgeUint8(value *uint8) *StudentUpdateOne {
	if value != nil {
		return s.SetAgeUint8(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentCreate) SetNotNilAgeUint8(value *uint8) *StudentCreate {
	if value != nil {
		return s.SetAgeUint8(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdate) SetNotNilAgeInt16(value *int16) *StudentUpdate {
	if value != nil {
		return s.SetAgeInt16(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdateOne) SetNotNilAgeInt16(value *int16) *StudentUpdateOne {
	if value != nil {
		return s.SetAgeInt16(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentCreate) SetNotNilAgeInt16(value *int16) *StudentCreate {
	if value != nil {
		return s.SetAgeInt16(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdate) SetNotNilAgeUint16(value *uint16) *StudentUpdate {
	if value != nil {
		return s.SetAgeUint16(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdateOne) SetNotNilAgeUint16(value *uint16) *StudentUpdateOne {
	if value != nil {
		return s.SetAgeUint16(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentCreate) SetNotNilAgeUint16(value *uint16) *StudentCreate {
	if value != nil {
		return s.SetAgeUint16(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdate) SetNotNilAgeInt32(value *int32) *StudentUpdate {
	if value != nil {
		return s.SetAgeInt32(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdateOne) SetNotNilAgeInt32(value *int32) *StudentUpdateOne {
	if value != nil {
		return s.SetAgeInt32(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentCreate) SetNotNilAgeInt32(value *int32) *StudentCreate {
	if value != nil {
		return s.SetAgeInt32(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdate) SetNotNilAgeUint32(value *uint32) *StudentUpdate {
	if value != nil {
		return s.SetAgeUint32(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdateOne) SetNotNilAgeUint32(value *uint32) *StudentUpdateOne {
	if value != nil {
		return s.SetAgeUint32(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentCreate) SetNotNilAgeUint32(value *uint32) *StudentCreate {
	if value != nil {
		return s.SetAgeUint32(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdate) SetNotNilAgeInt64(value *int64) *StudentUpdate {
	if value != nil {
		return s.SetAgeInt64(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdateOne) SetNotNilAgeInt64(value *int64) *StudentUpdateOne {
	if value != nil {
		return s.SetAgeInt64(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentCreate) SetNotNilAgeInt64(value *int64) *StudentCreate {
	if value != nil {
		return s.SetAgeInt64(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdate) SetNotNilAgeUint64(value *uint64) *StudentUpdate {
	if value != nil {
		return s.SetAgeUint64(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdateOne) SetNotNilAgeUint64(value *uint64) *StudentUpdateOne {
	if value != nil {
		return s.SetAgeUint64(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentCreate) SetNotNilAgeUint64(value *uint64) *StudentCreate {
	if value != nil {
		return s.SetAgeUint64(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdate) SetNotNilAgeInt(value *int) *StudentUpdate {
	if value != nil {
		return s.SetAgeInt(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdateOne) SetNotNilAgeInt(value *int) *StudentUpdateOne {
	if value != nil {
		return s.SetAgeInt(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentCreate) SetNotNilAgeInt(value *int) *StudentCreate {
	if value != nil {
		return s.SetAgeInt(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdate) SetNotNilAgeUint(value *uint) *StudentUpdate {
	if value != nil {
		return s.SetAgeUint(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdateOne) SetNotNilAgeUint(value *uint) *StudentUpdateOne {
	if value != nil {
		return s.SetAgeUint(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentCreate) SetNotNilAgeUint(value *uint) *StudentCreate {
	if value != nil {
		return s.SetAgeUint(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdate) SetNotNilWeightFloat(value *float64) *StudentUpdate {
	if value != nil {
		return s.SetWeightFloat(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdateOne) SetNotNilWeightFloat(value *float64) *StudentUpdateOne {
	if value != nil {
		return s.SetWeightFloat(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentCreate) SetNotNilWeightFloat(value *float64) *StudentCreate {
	if value != nil {
		return s.SetWeightFloat(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdate) SetNotNilWeightFloat32(value *float32) *StudentUpdate {
	if value != nil {
		return s.SetWeightFloat32(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdateOne) SetNotNilWeightFloat32(value *float32) *StudentUpdateOne {
	if value != nil {
		return s.SetWeightFloat32(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentCreate) SetNotNilWeightFloat32(value *float32) *StudentCreate {
	if value != nil {
		return s.SetWeightFloat32(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdate) SetNotNilClassID(value *uuid.UUID) *StudentUpdate {
	if value != nil {
		return s.SetClassID(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdateOne) SetNotNilClassID(value *uuid.UUID) *StudentUpdateOne {
	if value != nil {
		return s.SetClassID(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentCreate) SetNotNilClassID(value *uuid.UUID) *StudentCreate {
	if value != nil {
		return s.SetClassID(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdate) SetNotNilTeacherID(value *uint64) *StudentUpdate {
	if value != nil {
		return s.SetTeacherID(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdateOne) SetNotNilTeacherID(value *uint64) *StudentUpdateOne {
	if value != nil {
		return s.SetTeacherID(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentCreate) SetNotNilTeacherID(value *uint64) *StudentCreate {
	if value != nil {
		return s.SetTeacherID(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdate) SetNotNilEnrollAt(value *time.Time) *StudentUpdate {
	if value != nil {
		return s.SetEnrollAt(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdateOne) SetNotNilEnrollAt(value *time.Time) *StudentUpdateOne {
	if value != nil {
		return s.SetEnrollAt(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentCreate) SetNotNilEnrollAt(value *time.Time) *StudentCreate {
	if value != nil {
		return s.SetEnrollAt(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdate) SetNotNilStatusBool(value *bool) *StudentUpdate {
	if value != nil {
		return s.SetStatusBool(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentUpdateOne) SetNotNilStatusBool(value *bool) *StudentUpdateOne {
	if value != nil {
		return s.SetStatusBool(*value)
	}
	return s
}

// set field if value's pointer is not nil.
func (s *StudentCreate) SetNotNilStatusBool(value *bool) *StudentCreate {
	if value != nil {
		return s.SetStatusBool(*value)
	}
	return s
}
