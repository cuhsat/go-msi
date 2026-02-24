package msi

import (
	"encoding/binary"
	"fmt"

	"github.com/cuhsat/go-mscfb/pkg/mscfb"
)

const (
	ColFieldSizeMask  int32 = 0xff
	ColLocalizableBit int32 = 0x200
	ColStringBit      int32 = 0x800
	ColNullableBit    int32 = 0x1000
	ColPrimaryKeyBit  int32 = 0x2000
)

type ColumnType int

const (
	ColumnTypeInt16 ColumnType = iota
	ColumnTypeInt32
	ColumnTypeStr
)

func ColumnTypeFromBitField(typeBits int32) (ColumnType, int, error) {
	fieldSize := int(typeBits & ColFieldSizeMask)
	if typeBits&ColStringBit != 0 {
		return ColumnTypeStr, fieldSize, nil
	} else if fieldSize == 4 {
		return ColumnTypeInt32, fieldSize, nil
	} else if fieldSize == 2 {
		return ColumnTypeInt16, fieldSize, nil
	} else if fieldSize == 1 {
		return ColumnTypeInt16, fieldSize, nil
	}

	return -1, 0, fmt.Errorf("invalid column type bits: %x", typeBits)
}

func (c ColumnType) Width(longStringRefs bool) uint64 {
	switch c {
	case ColumnTypeInt16:
		return 2
	case ColumnTypeInt32:
		return 4
	case ColumnTypeStr:
		if longStringRefs {
			return 3
		}
		return 2
	}

	return 0
}

func (c ColumnType) ReadValue(stream *mscfb.Stream, longStringRefs bool) (*ValueRef, error) {
	switch c {
	case ColumnTypeInt16:
		var value int16
		err := binary.Read(stream, binary.LittleEndian, &value)
		if err != nil {
			return nil, err
		}
		if value == 0 {
			return &ValueRef{IsNull: true}, nil
		}
		return &ValueRef{
			IsInt: true,
			Value: int(value ^ -0x8000),
		}, nil
	case ColumnTypeInt32:
		var value int32
		err := binary.Read(stream, binary.LittleEndian, &value)
		if err != nil {
			return nil, err
		}
		if value == 0 {
			return &ValueRef{IsNull: true}, nil
		}

		return &ValueRef{
			IsInt: true,
			Value: int(value ^ -0x8000_0000),
		}, nil
	case ColumnTypeStr:
		var value StringRef
		rd, err := value.Read(stream, longStringRefs)
		if err != nil {
			return nil, err
		}
		if rd.Num == 0 {
			return &ValueRef{
				IsNull: true,
			}, nil
		}

		return &ValueRef{
			IsStr: true,
			Value: rd,
		}, nil
	}

	return nil, nil
}

type valueRange struct {
	Min int32
	Max int32
}

type foreignKey struct {
	TableName   string
	ColumnIndex int32
}

type Column struct {
	Name             string
	ColumnType       ColumnType
	ColumnStringSize int
	IsLocalizable    bool
	IsNullable       bool
	IsPrimarykey     bool
	ValueRange       valueRange
	ForeignKey       foreignKey
	Category         Category
	EnumValues       []string
}
