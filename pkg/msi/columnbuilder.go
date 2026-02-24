package msi

type ColumnBuilder struct {
	Name          string
	IsLocalizable bool
	IsNullable    bool
	IsPrimarykey  bool
	ValueRange    valueRange
	ForeignKey    foreignKey
	Category      Category
	EnumValues    []string
}

func NewColumnBuilder(name string) *ColumnBuilder {
	return &ColumnBuilder{
		Name:       name,
		EnumValues: make([]string, 0),
	}
}

func (b *ColumnBuilder) SetLocalizable() *ColumnBuilder {
	b.IsLocalizable = true
	return b
}

func (b *ColumnBuilder) SetNullable() *ColumnBuilder {
	b.IsNullable = true
	return b
}

func (b *ColumnBuilder) SetPrimaryKey() *ColumnBuilder {
	b.IsPrimarykey = true
	return b
}

func (b *ColumnBuilder) SetRange(min, max int32) *ColumnBuilder {
	b.ValueRange = valueRange{Min: min, Max: max}
	return b
}

func (b *ColumnBuilder) SetForeignKey(tableName string, colIndex int32) *ColumnBuilder {
	b.ForeignKey = foreignKey{
		TableName:   tableName,
		ColumnIndex: colIndex,
	}
	return b
}

func (b *ColumnBuilder) SetCategory(category Category) *ColumnBuilder {
	b.Category = category
	return b
}

func (b *ColumnBuilder) SetEnumValues(values ...string) *ColumnBuilder {
	b.EnumValues = values
	return b
}

func (b *ColumnBuilder) Int16() *Column {
	return b.withType(ColumnTypeInt16)
}

func (b *ColumnBuilder) Int32() *Column {
	return b.withType(ColumnTypeInt32)
}

func (b *ColumnBuilder) String(maxLen int) *Column {
	return b.withTypeSize(ColumnTypeStr, maxLen)
}

func (b *ColumnBuilder) IDString(maxLen int) *Column {
	return b.SetCategory(CategoryIdentifier).String(maxLen)
}

func (b *ColumnBuilder) TextString(maxLen int) *Column {
	return b.SetCategory(CategoryText).String(maxLen)
}

func (b *ColumnBuilder) FormattedString(maxLen int) *Column {
	return b.SetCategory(CategoryFormatted).String(maxLen)
}

func (b *ColumnBuilder) Binary() *Column {
	return b.SetCategory(CategoryBinary).String(0)
}

func (b *ColumnBuilder) withType(colType ColumnType) *Column {
	return b.withTypeSize(colType, 0)
}

func (b *ColumnBuilder) withTypeSize(colType ColumnType, size int) *Column {
	return &Column{
		Name:             b.Name,
		ColumnType:       colType,
		ColumnStringSize: size,
		IsLocalizable:    b.IsLocalizable,
		IsNullable:       b.IsNullable,
		IsPrimarykey:     b.IsPrimarykey,
		ValueRange:       b.ValueRange,
		ForeignKey:       b.ForeignKey,
		Category:         b.Category,
		EnumValues:       b.EnumValues,
	}
}

func (b *ColumnBuilder) withBitFields(typeBits int32) (*Column, error) {
	isNullable := typeBits&ColNullableBit != 0

	colType, size, err := ColumnTypeFromBitField(typeBits)
	if err != nil {
		return nil, err
	}

	return &Column{
		Name:             b.Name,
		ColumnType:       colType,
		ColumnStringSize: size,
		IsLocalizable:    typeBits&ColLocalizableBit != 0,
		IsNullable:       isNullable || b.IsNullable,
		IsPrimarykey:     typeBits&ColPrimaryKeyBit != 0,
		ValueRange:       b.ValueRange,
		ForeignKey:       b.ForeignKey,
		Category:         b.Category,
		EnumValues:       b.EnumValues,
	}, nil
}
