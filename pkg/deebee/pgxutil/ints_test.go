package pgxutil_test

import (
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
)

func TestIntToPgtypeInt4(t *testing.T) {
	val := 100
	pgInt := pgxutil.IntToPgtypeInt4(&val)

	assert.True(t, pgInt.Valid)
	assert.Equal(t, int32(val), pgInt.Int32)

	var nilInt *int
	pgIntNil := pgxutil.IntToPgtypeInt4(nilInt)
	assert.False(t, pgIntNil.Valid)
}

func TestInt32ToPgtypeInt4(t *testing.T) {
	val := int32(42)
	pgInt := pgxutil.Int32ToPgtypeInt4(&val)

	assert.True(t, pgInt.Valid)
	assert.Equal(t, int32(42), pgInt.Int32)

	var nilInt *int32
	pgIntNil := pgxutil.Int32ToPgtypeInt4(nilInt)
	assert.False(t, pgIntNil.Valid)
	assert.Equal(t, int32(0), pgIntNil.Int32)

	negVal := int32(-100)
	pgIntNeg := pgxutil.Int32ToPgtypeInt4(&negVal)
	assert.True(t, pgIntNeg.Valid)
	assert.Equal(t, int32(-100), pgIntNeg.Int32)
}

func TestInt16ToPgtypeInt2(t *testing.T) {
	val := int16(50)
	pgInt := pgxutil.Int16ToPgtypeInt2(&val)

	assert.True(t, pgInt.Valid)
	assert.Equal(t, int16(50), pgInt.Int16)

	var nilInt *int16
	pgIntNil := pgxutil.Int16ToPgtypeInt2(nilInt)
	assert.False(t, pgIntNil.Valid)
	assert.Equal(t, int16(0), pgIntNil.Int16)

	negVal := int16(-100)
	pgIntNeg := pgxutil.Int16ToPgtypeInt2(&negVal)
	assert.True(t, pgIntNeg.Valid)
	assert.Equal(t, int16(-100), pgIntNeg.Int16)

	maxVal := int16(32767)
	pgIntMax := pgxutil.Int16ToPgtypeInt2(&maxVal)
	assert.True(t, pgIntMax.Valid)
	assert.Equal(t, int16(32767), pgIntMax.Int16)
}

func TestPgtypeInt2ToInt16Ptr(t *testing.T) {
	val := int16(50)
	pgInt := pgtype.Int2{Int16: val, Valid: true}
	result := pgxutil.PgtypeInt2ToInt16Ptr(pgInt)

	assert.NotNil(t, result)
	assert.Equal(t, val, *result)

	pgIntNil := pgtype.Int2{Int16: 0, Valid: false}
	resultNil := pgxutil.PgtypeInt2ToInt16Ptr(pgIntNil)
	assert.Nil(t, resultNil)

	negVal := int16(-100)
	pgIntNeg := pgtype.Int2{Int16: negVal, Valid: true}
	resultNeg := pgxutil.PgtypeInt2ToInt16Ptr(pgIntNeg)
	assert.NotNil(t, resultNeg)
	assert.Equal(t, negVal, *resultNeg)

	maxVal := int16(32767)
	pgIntMax := pgtype.Int2{Int16: maxVal, Valid: true}
	resultMax := pgxutil.PgtypeInt2ToInt16Ptr(pgIntMax)
	assert.NotNil(t, resultMax)
	assert.Equal(t, maxVal, *resultMax)

	minVal := int16(-32768)
	pgIntMin := pgtype.Int2{Int16: minVal, Valid: true}
	resultMin := pgxutil.PgtypeInt2ToInt16Ptr(pgIntMin)
	assert.NotNil(t, resultMin)
	assert.Equal(t, minVal, *resultMin)
}

func TestPgtypeInt4ToInt32Ptr(t *testing.T) {
	val := int32(100)
	pgInt := pgtype.Int4{Int32: val, Valid: true}
	result := pgxutil.PgtypeInt4ToInt32Ptr(pgInt)

	assert.NotNil(t, result)
	assert.Equal(t, val, *result)

	pgIntNil := pgtype.Int4{Int32: 0, Valid: false}
	resultNil := pgxutil.PgtypeInt4ToInt32Ptr(pgIntNil)
	assert.Nil(t, resultNil)

	negVal := int32(-100)
	pgIntNeg := pgtype.Int4{Int32: negVal, Valid: true}
	resultNeg := pgxutil.PgtypeInt4ToInt32Ptr(pgIntNeg)
	assert.NotNil(t, resultNeg)
	assert.Equal(t, negVal, *resultNeg)

	maxVal := int32(2147483647)
	pgIntMax := pgtype.Int4{Int32: maxVal, Valid: true}
	resultMax := pgxutil.PgtypeInt4ToInt32Ptr(pgIntMax)
	assert.NotNil(t, resultMax)
	assert.Equal(t, maxVal, *resultMax)

	minVal := int32(-2147483648)
	pgIntMin := pgtype.Int4{Int32: minVal, Valid: true}
	resultMin := pgxutil.PgtypeInt4ToInt32Ptr(pgIntMin)
	assert.NotNil(t, resultMin)
	assert.Equal(t, minVal, *resultMin)
}

func TestPgtypeInt8ToInt64Ptr(t *testing.T) {
	val := int64(100)
	pgInt := pgtype.Int8{Int64: val, Valid: true}
	result := pgxutil.PgtypeInt8ToInt64Ptr(pgInt)

	assert.NotNil(t, result)
	assert.Equal(t, val, *result)

	pgIntNil := pgtype.Int8{Int64: 0, Valid: false}
	resultNil := pgxutil.PgtypeInt8ToInt64Ptr(pgIntNil)
	assert.Nil(t, resultNil)

	negVal := int64(-100)
	pgIntNeg := pgtype.Int8{Int64: negVal, Valid: true}
	resultNeg := pgxutil.PgtypeInt8ToInt64Ptr(pgIntNeg)
	assert.NotNil(t, resultNeg)
	assert.Equal(t, negVal, *resultNeg)

	maxVal := int64(9223372036854775807)
	pgIntMax := pgtype.Int8{Int64: maxVal, Valid: true}
	resultMax := pgxutil.PgtypeInt8ToInt64Ptr(pgIntMax)
	assert.NotNil(t, resultMax)
	assert.Equal(t, maxVal, *resultMax)

	minVal := int64(-9223372036854775808)
	pgIntMin := pgtype.Int8{Int64: minVal, Valid: true}
	resultMin := pgxutil.PgtypeInt8ToInt64Ptr(pgIntMin)
	assert.NotNil(t, resultMin)
	assert.Equal(t, minVal, *resultMin)
}

func TestInt64ToPgtypeInt8(t *testing.T) {
	val := int64(100)
	pgInt := pgxutil.Int64ToPgtypeInt8(&val)

	assert.True(t, pgInt.Valid)
	assert.Equal(t, int64(100), pgInt.Int64)

	var nilInt *int64
	pgIntNil := pgxutil.Int64ToPgtypeInt8(nilInt)
	assert.False(t, pgIntNil.Valid)
	assert.Equal(t, int64(0), pgIntNil.Int64)

	negVal := int64(-100)
	pgIntNeg := pgxutil.Int64ToPgtypeInt8(&negVal)
	assert.True(t, pgIntNeg.Valid)
	assert.Equal(t, int64(-100), pgIntNeg.Int64)

	maxVal := int64(9223372036854775807)
	pgIntMax := pgxutil.Int64ToPgtypeInt8(&maxVal)
	assert.True(t, pgIntMax.Valid)
	assert.Equal(t, int64(9223372036854775807), pgIntMax.Int64)

	minVal := int64(-9223372036854775808)
	pgIntMin := pgxutil.Int64ToPgtypeInt8(&minVal)
	assert.True(t, pgIntMin.Valid)
	assert.Equal(t, int64(-9223372036854775808), pgIntMin.Int64)
}
