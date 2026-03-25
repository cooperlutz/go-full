// Package pgxutil provides utility functions for working with pgx types, such as converting between Go types and pgx types, handling nullability, and more. These functions help simplify the process of working with pgx when dealing with various PostgreSQL data types.
package pgxutil

import (
	"database/sql"
	"net"
	"net/netip"
	"time"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/cooperlutz/go-full/pkg/utilitee"
)

// CREATE TABLE IF NOT EXISTS all_types_example (
//
//	-- Numeric Types
//	uuid UUID PRIMARY KEY, -- Auto-incrementing integer
//	uuid_not_null UUID NOT NULL,
//	small_int SMALLINT,
//	small_int_not_null SMALLINT NOT NULL,
//	regular_int INTEGER,
//	regular_int_not_null INTEGER NOT NULL,
//	big_int BIGINT,
//	big_int_not_null BIGINT NOT NULL,
//	exact_numeric NUMERIC(10, 2), -- Exact precision for currency, etc.
//	exact_numeric_not_null NUMERIC(10, 2) NOT NULL,
//	approx_real REAL, -- Single precision float
//	approx_real_not_null REAL NOT NULL,
//	approx_double DOUBLE PRECISION, -- Double precision float
//	approx_double_not_null DOUBLE PRECISION NOT NULL,
//	monetary_amount MONEY,
//	monetary_amount_not_null MONEY NOT NULL,
//	-- Character Types
//	fixed_char CHAR(3), -- Fixed length, blank-padded
//	fixed_char_not_null CHAR(3) NOT NULL, -- Fixed length, blank-padded
//	var_char VARCHAR(255), -- Variable length with limit
//	var_char_not_null VARCHAR(255) NOT NULL, -- Variable length with limit
//	unlimited_text TEXT, -- Variable length, unlimited
//	unlimited_text_not_null TEXT NOT NULL, -- Variable length, unlimited
//	-- Boolean Type
//	bool BOOLEAN,
//	bool_not_null BOOLEAN NOT NULL,
//	-- Date/Time Types
//	just_date DATE,
//	just_date_not_null DATE NOT NULL,
//	just_time TIME WITHOUT TIME ZONE,
//	just_time_not_null TIME WITHOUT TIME ZONE NOT NULL,
//	time_with_tz TIME WITH TIME ZONE,
//	time_with_tz_not_null TIME WITH TIME ZONE NOT NULL,
//	timestamp_without_tz TIMESTAMP WITHOUT TIME ZONE,
//	timestamp_without_tz_not_null TIMESTAMP WITHOUT TIME ZONE NOT NULL,
//	timestamp_with_tz TIMESTAMP WITH TIME ZONE DEFAULT now(), -- Stores creation time with time zone
//	timestamp_with_tz_not_null TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL, -- Stores creation time with time zone
//	duration INTERVAL,
//	duration_not_null INTERVAL NOT NULL,
//	-- Binary Data
//	raw_bytes BYTEA,
//	raw_bytes_not_null BYTEA NOT NULL,
//	-- Network Types
//	ip_address INET, -- IPv4 or IPv6 host address
//	ip_address_not_null INET NOT NULL, -- IPv4 or IPv6 host address
//	network_cidr CIDR, -- IPv4 or IPv6 network block
//	network_cidr_not_null CIDR NOT NULL, -- IPv4 or IPv6 network block
//	mac_address MACADDR,
//	mac_address_not_null MACADDR NOT NULL,
//	-- Geometric Types
//	geo_point POINT, -- (x, y) point
//	geo_point_not_null POINT NOT NULL, -- (x, y) point
//	geo_line LSEG, -- Line segment
//	geo_line_not_null LSEG NOT NULL, -- Line segment
//	geo_box BOX, -- Rectangular box
//	geo_box_not_null BOX NOT NULL, -- Rectangular box
//	geo_circle CIRCLE, -- Center point and radius
//	geo_circle_not_null CIRCLE NOT NULL, -- Center point and radius
//	-- JSON Types
//	json_data JSON, -- Textual JSON data
//	json_data_not_null JSON NOT NULL, -- Textual JSON data
//	jsonb_data JSONB, -- Decomposed binary JSON data (more efficient)
//	jsonb_data_not_null JSONB NOT NULL, -- Decomposed binary JSON data (more efficient)
//	-- Array Types
//	int_array INTEGER[],
//	int_array_not_null INTEGER[] NOT NULL,
//	text_array TEXT[],
//	text_array_not_null TEXT[] NOT NULL,
//	-- Range Types
//	int4_range INT4RANGE, -- Range of integer
//	int4_range_not_null INT4RANGE NOT NULL, -- Range of integer
//	date_range DATERANGE, -- Range of date
//	date_range_not_null DATERANGE NOT NULL -- Range of date
//
// );
//
// CORRESPONDING GENERATED GO STRUCTS:.
type AllTypesExample struct {
	Uuid                      pgtype.UUID               `db:"uuid" json:"uuid"`
	UuidNotNull               pgtype.UUID               `db:"uuid_not_null" json:"uuid_not_null"`
	SmallInt                  pgtype.Int2               `db:"small_int" json:"small_int"`
	SmallIntNotNull           int16                     `db:"small_int_not_null" json:"small_int_not_null"`
	RegularInt                pgtype.Int4               `db:"regular_int" json:"regular_int"`
	RegularIntNotNull         int32                     `db:"regular_int_not_null" json:"regular_int_not_null"`
	BigInt                    pgtype.Int8               `db:"big_int" json:"big_int"`
	BigIntNotNull             int64                     `db:"big_int_not_null" json:"big_int_not_null"`
	ExactNumeric              pgtype.Numeric            `db:"exact_numeric" json:"exact_numeric"`
	ExactNumericNotNull       pgtype.Numeric            `db:"exact_numeric_not_null" json:"exact_numeric_not_null"`
	ApproxReal                pgtype.Float4             `db:"approx_real" json:"approx_real"`
	ApproxRealNotNull         float32                   `db:"approx_real_not_null" json:"approx_real_not_null"`
	ApproxDouble              pgtype.Float8             `db:"approx_double" json:"approx_double"`
	ApproxDoubleNotNull       float64                   `db:"approx_double_not_null" json:"approx_double_not_null"`
	MonetaryAmount            pgtype.Numeric            `db:"monetary_amount" json:"monetary_amount"`
	MonetaryAmountNotNull     pgtype.Numeric            `db:"monetary_amount_not_null" json:"monetary_amount_not_null"`
	FixedChar                 pgtype.Text               `db:"fixed_char" json:"fixed_char"`
	FixedCharNotNull          string                    `db:"fixed_char_not_null" json:"fixed_char_not_null"`
	VarChar                   pgtype.Text               `db:"var_char" json:"var_char"`
	VarCharNotNull            string                    `db:"var_char_not_null" json:"var_char_not_null"`
	UnlimitedText             pgtype.Text               `db:"unlimited_text" json:"unlimited_text"`
	UnlimitedTextNotNull      string                    `db:"unlimited_text_not_null" json:"unlimited_text_not_null"`
	Bool                      pgtype.Bool               `db:"bool" json:"bool"`
	BoolNotNull               bool                      `db:"bool_not_null" json:"bool_not_null"`
	JustDate                  pgtype.Date               `db:"just_date" json:"just_date"`
	JustDateNotNull           pgtype.Date               `db:"just_date_not_null" json:"just_date_not_null"`
	JustTime                  pgtype.Time               `db:"just_time" json:"just_time"`
	JustTimeNotNull           pgtype.Time               `db:"just_time_not_null" json:"just_time_not_null"`
	TimeWithTz                sql.NullTime              `db:"time_with_tz" json:"time_with_tz"`
	TimeWithTzNotNull         time.Time                 `db:"time_with_tz_not_null" json:"time_with_tz_not_null"`
	TimestampWithoutTz        pgtype.Timestamp          `db:"timestamp_without_tz" json:"timestamp_without_tz"`
	TimestampWithoutTzNotNull pgtype.Timestamp          `db:"timestamp_without_tz_not_null" json:"timestamp_without_tz_not_null"`
	TimestampWithTz           pgtype.Timestamptz        `db:"timestamp_with_tz" json:"timestamp_with_tz"`
	TimestampWithTzNotNull    pgtype.Timestamptz        `db:"timestamp_with_tz_not_null" json:"timestamp_with_tz_not_null"`
	Duration                  pgtype.Interval           `db:"duration" json:"duration"`
	DurationNotNull           pgtype.Interval           `db:"duration_not_null" json:"duration_not_null"`
	RawBytes                  []byte                    `db:"raw_bytes" json:"raw_bytes"`
	RawBytesNotNull           []byte                    `db:"raw_bytes_not_null" json:"raw_bytes_not_null"`
	IpAddress                 *netip.Addr               `db:"ip_address" json:"ip_address"`
	IpAddressNotNull          netip.Addr                `db:"ip_address_not_null" json:"ip_address_not_null"`
	NetworkCidr               *netip.Prefix             `db:"network_cidr" json:"network_cidr"`
	NetworkCidrNotNull        netip.Prefix              `db:"network_cidr_not_null" json:"network_cidr_not_null"`
	MacAddress                net.HardwareAddr          `db:"mac_address" json:"mac_address"`
	MacAddressNotNull         net.HardwareAddr          `db:"mac_address_not_null" json:"mac_address_not_null"`
	GeoPoint                  pgtype.Point              `db:"geo_point" json:"geo_point"`
	GeoPointNotNull           pgtype.Point              `db:"geo_point_not_null" json:"geo_point_not_null"`
	GeoLine                   pgtype.Lseg               `db:"geo_line" json:"geo_line"`
	GeoLineNotNull            pgtype.Lseg               `db:"geo_line_not_null" json:"geo_line_not_null"`
	GeoBox                    pgtype.Box                `db:"geo_box" json:"geo_box"`
	GeoBoxNotNull             pgtype.Box                `db:"geo_box_not_null" json:"geo_box_not_null"`
	GeoCircle                 pgtype.Circle             `db:"geo_circle" json:"geo_circle"`
	GeoCircleNotNull          pgtype.Circle             `db:"geo_circle_not_null" json:"geo_circle_not_null"`
	JsonData                  []byte                    `db:"json_data" json:"json_data"`
	JsonDataNotNull           []byte                    `db:"json_data_not_null" json:"json_data_not_null"`
	JsonbData                 []byte                    `db:"jsonb_data" json:"jsonb_data"`
	JsonbDataNotNull          []byte                    `db:"jsonb_data_not_null" json:"jsonb_data_not_null"`
	IntArray                  []int32                   `db:"int_array" json:"int_array"`
	IntArrayNotNull           []int32                   `db:"int_array_not_null" json:"int_array_not_null"`
	TextArray                 []string                  `db:"text_array" json:"text_array"`
	TextArrayNotNull          []string                  `db:"text_array_not_null" json:"text_array_not_null"`
	Int4Range                 pgtype.Range[pgtype.Int4] `db:"int4_range" json:"int4_range"`
	Int4RangeNotNull          pgtype.Range[pgtype.Int4] `db:"int4_range_not_null" json:"int4_range_not_null"`
	DateRange                 pgtype.Range[pgtype.Date] `db:"date_range" json:"date_range"`
	DateRangeNotNull          pgtype.Range[pgtype.Date] `db:"date_range_not_null" json:"date_range_not_null"`
}

// SliceOfPtrsToPgtype converts a slice of pointers to a slice of values, treating nil pointers as null values.
func SliceOfPtrsToPgtype[T any](slice []*T) []T {
	return utilitee.SliceOfPointersToSlice(slice)
}
