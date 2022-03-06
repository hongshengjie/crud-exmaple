// Code generated by bcurd. DO NOT EDIT.

package alltypetable

import (
	"time"
)

// AllTypeTable represents a row from 'all_type_table'.
type AllTypeTable struct {
	Id              int64     `json:"id"`               // 自增id
	TInt            int64     `json:"t_int"`            // 小整型
	SInt            int64     `json:"s_int"`            //
	MInt            int64     `json:"m_int"`            // 中整数
	BInt            int64     `json:"b_int"`            //
	F32             float32   `json:"f32"`              //
	F64             float64   `json:"f64"`              //
	DecimalMysql    float64   `json:"decimal_mysql"`    //
	CharM           string    `json:"char_m"`           //
	VarcharM        string    `json:"varchar_m"`        //
	JsonM           string    `json:"json_m"`           //
	NvarcharM       string    `json:"nvarchar_m"`       //
	NcharM          string    `json:"nchar_m"`          //
	TimeM           string    `json:"time_m"`           //
	DateM           time.Time `json:"date_m"`           //
	DataTimeM       time.Time `json:"data_time_m"`      //
	TimestampM      time.Time `json:"timestamp_m"`      // 创建时间
	TimestampUpdate time.Time `json:"timestamp_update"` // 更新时间
	YearM           string    `json:"year_m"`           // 年
	TText           string    `json:"t_text"`           //
	MText           string    `json:"m_text"`           //
	TextM           string    `json:"text_m"`           //
	LText           string    `json:"l_text"`           //
	BinaryM         []byte    `json:"binary_m"`         //
	BlobM           []byte    `json:"blob_m"`           //
	LBlob           []byte    `json:"l_blob"`           //
	MBlob           []byte    `json:"m_blob"`           //
	TBlob           []byte    `json:"t_blob"`           //
	BitM            []byte    `json:"bit_m"`            //
	EnumM           string    `json:"enum_m"`           //
	SetM            string    `json:"set_m"`            //
	BoolM           int64     `json:"bool_m"`           //
}

const (
	// table tableName is all_type_table
	table = "all_type_table"
	//Id 自增id
	Id = "id"
	//TInt 小整型
	TInt = "t_int"
	//SInt
	SInt = "s_int"
	//MInt 中整数
	MInt = "m_int"
	//BInt
	BInt = "b_int"
	//F32
	F32 = "f32"
	//F64
	F64 = "f64"
	//DecimalMysql
	DecimalMysql = "decimal_mysql"
	//CharM
	CharM = "char_m"
	//VarcharM
	VarcharM = "varchar_m"
	//JsonM
	JsonM = "json_m"
	//NvarcharM
	NvarcharM = "nvarchar_m"
	//NcharM
	NcharM = "nchar_m"
	//TimeM
	TimeM = "time_m"
	//DateM
	DateM = "date_m"
	//DataTimeM
	DataTimeM = "data_time_m"
	//TimestampM 创建时间
	TimestampM = "timestamp_m"
	//TimestampUpdate 更新时间
	TimestampUpdate = "timestamp_update"
	//YearM 年
	YearM = "year_m"
	//TText
	TText = "t_text"
	//MText
	MText = "m_text"
	//TextM
	TextM = "text_m"
	//LText
	LText = "l_text"
	//BinaryM
	BinaryM = "binary_m"
	//BlobM
	BlobM = "blob_m"
	//LBlob
	LBlob = "l_blob"
	//MBlob
	MBlob = "m_blob"
	//TBlob
	TBlob = "t_blob"
	//BitM
	BitM = "bit_m"
	//EnumM
	EnumM = "enum_m"
	//SetM
	SetM = "set_m"
	//BoolM
	BoolM = "bool_m"
)

// columns holds all SQL columns.
var columns = []string{
	Id,
	TInt,
	SInt,
	MInt,
	BInt,
	F32,
	F64,
	DecimalMysql,
	CharM,
	VarcharM,
	JsonM,
	NvarcharM,
	NcharM,
	TimeM,
	DateM,
	DataTimeM,
	TimestampM,
	TimestampUpdate,
	YearM,
	TText,
	MText,
	TextM,
	LText,
	BinaryM,
	BlobM,
	LBlob,
	MBlob,
	TBlob,
	BitM,
	EnumM,
	SetM,
	BoolM,
}

// Columns returns table all columns field name slice
func Columns() []string {
	return columns
}
