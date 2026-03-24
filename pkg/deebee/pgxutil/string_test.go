package pgxutil_test

import (
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
)

func TestStrToPgtypeText(t *testing.T) {
	str := "Hello, World!"
	pgText := pgxutil.StrToPgtypeText(&str)

	assert.True(t, pgText.Valid)
	assert.Equal(t, str, pgText.String)

	var nilStr *string
	pgTextNil := pgxutil.StrToPgtypeText(nilStr)
	assert.False(t, pgTextNil.Valid)
}

func TestStringToPgtypeText(t *testing.T) {
	str := "Hello, World!"
	pgText := pgxutil.StringToPgtypeText(&str)

	assert.True(t, pgText.Valid)
	assert.Equal(t, str, pgText.String)

	var nilStr *string
	pgTextNil := pgxutil.StringToPgtypeText(nilStr)
	assert.False(t, pgTextNil.Valid)
	assert.Equal(t, "", pgTextNil.String)
}

func TestPgtypeTextToStringPtr(t *testing.T) {
	str := "Hello, World!"
	pgText := pgtype.Text{String: str, Valid: true}
	result := pgxutil.PgtypeTextToStringPtr(pgText)

	assert.NotNil(t, result)
	assert.Equal(t, str, *result)
}

func TestPgtypeTextToStringPtrNil(t *testing.T) {
	pgText := pgtype.Text{String: "ignored", Valid: false}
	result := pgxutil.PgtypeTextToStringPtr(pgText)

	assert.Nil(t, result)
}

func TestPgtypeTextToStringPtrEmpty(t *testing.T) {
	pgText := pgtype.Text{String: "", Valid: true}
	result := pgxutil.PgtypeTextToStringPtr(pgText)

	assert.NotNil(t, result)
	assert.Equal(t, "", *result)
}
