package queryparams

import (
	netUrl "net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockUrl(query string) *netUrl.URL {
	return &netUrl.URL{RawQuery: query}
}

func TestQueryParams_Exists(t *testing.T) {
	params := NewQueryParams(mockUrl("foo=bar"))
	assert.False(t, params.Exists("unknown"))
	assert.True(t, params.Exists("foo"))
}

func TestQueryParams_GetString(t *testing.T) {
	params := NewQueryParams(mockUrl("foo=bar"))
	assert.Equal(t, params.GetString("foo"), "bar")
	assert.Empty(t, params.GetString("unknown"))
	assert.Equal(t, params.GetString("unknown", "default"), "default")
}

func TestQueryParams_GetStrings(t *testing.T) {
	params := NewQueryParams(mockUrl("foo=bar1&foo=bar2"))
	assert.Equal(t, params.GetStrings("foo"), []string{"bar1", "bar2"})
	assert.Empty(t, params.GetStrings("unknown"))
	assert.Equal(t, params.GetStrings("unknown", []string{"default"}), []string{"default"})
}

func TestNewQueryParams_GetBool(t *testing.T) {
	params := NewQueryParams(mockUrl("flag=true"))
	flag, _ := params.GetBool("flag")
	assert.True(t, flag)
	flag, err := params.GetBool("unknown")
	assert.Empty(t, flag)
	assert.NotEmpty(t, err)
	flag, err = params.GetBool("unknown", true)
	assert.Empty(t, err)
	assert.True(t, flag)
}

func TestQueryParams_GetFloat(t *testing.T) {
	params := NewQueryParams(mockUrl("var=1.23&foo=?"))
	v, _ := params.GetFloat("var")
	assert.Equal(t, v, 1.23)
	f, err := params.GetFloat("foo")
	assert.Empty(t, f)
	assert.NotEmpty(t, err)
	v, err = params.GetFloat("unknown")
	assert.Empty(t, v)
	assert.NotEmpty(t, err)
	v, err = params.GetFloat("unknown")
	assert.Empty(t, v)
	assert.NotEmpty(t, err)
	v, err = params.GetFloat("unknown", 32.1)
	assert.Empty(t, err)
	assert.Equal(t, v, 32.1)
}

func TestQueryParams_GetInt(t *testing.T) {
	params := NewQueryParams(mockUrl("var=125&foo=?"))
	v, _ := params.GetInt("var")
	assert.Equal(t, v, 125)
	f, err := params.GetInt("foo")
	assert.Empty(t, f)
	assert.NotEmpty(t, err)
	v, err = params.GetInt("unknown")
	assert.Empty(t, v)
	assert.NotEmpty(t, err)
	v, err = params.GetInt("unknown")
	assert.Empty(t, v)
	assert.NotEmpty(t, err)
	v, err = params.GetInt("unknown", 32)
	assert.Empty(t, err)
	assert.Equal(t, v, 32)
}

func TestQueryParams_GetInt8(t *testing.T) {
	params := NewQueryParams(mockUrl("var=125&foo=?"))
	v, _ := params.GetInt8("var")
	assert.Equal(t, v, int8(125))
	f, err := params.GetInt8("foo")
	assert.Empty(t, f)
	assert.NotEmpty(t, err)
	v, err = params.GetInt8("unknown")
	assert.Empty(t, v)
	assert.NotEmpty(t, err)
	v, err = params.GetInt8("unknown")
	assert.Empty(t, v)
	assert.NotEmpty(t, err)
	v, err = params.GetInt8("unknown", 32)
	assert.Empty(t, err)
	assert.Equal(t, v, int8(32))
}

func TestQueryParams_GetInt16(t *testing.T) {
	params := NewQueryParams(mockUrl("var=125&foo=?"))
	v, _ := params.GetInt16("var")
	assert.Equal(t, v, int16(125))
	f, err := params.GetInt16("foo")
	assert.Empty(t, f)
	assert.NotEmpty(t, err)
	v, err = params.GetInt16("unknown")
	assert.Empty(t, v)
	assert.NotEmpty(t, err)
	v, err = params.GetInt16("unknown")
	assert.Empty(t, v)
	assert.NotEmpty(t, err)
	v, err = params.GetInt16("unknown", 32)
	assert.Empty(t, err)
	assert.Equal(t, v, int16(32))
}

func TestQueryParams_GetInt32(t *testing.T) {
	params := NewQueryParams(mockUrl("var=125&foo=?"))
	v, _ := params.GetInt32("var")
	assert.Equal(t, v, int32(125))
	f, err := params.GetInt32("foo")
	assert.Empty(t, f)
	assert.NotEmpty(t, err)
	v, err = params.GetInt32("unknown")
	assert.Empty(t, v)
	assert.NotEmpty(t, err)
	v, err = params.GetInt32("unknown")
	assert.Empty(t, v)
	assert.NotEmpty(t, err)
	v, err = params.GetInt32("unknown", 32)
	assert.Empty(t, err)
	assert.Equal(t, v, int32(32))
}

func TestQueryParams_GetInt64(t *testing.T) {
	params := NewQueryParams(mockUrl("var=125&foo=?"))
	v, _ := params.GetInt64("var")
	assert.Equal(t, v, int64(125))
	f, err := params.GetInt64("foo")
	assert.Empty(t, f)
	assert.NotEmpty(t, err)
	v, err = params.GetInt64("unknown")
	assert.Empty(t, v)
	assert.NotEmpty(t, err)
	v, err = params.GetInt64("unknown")
	assert.Empty(t, v)
	assert.NotEmpty(t, err)
	v, err = params.GetInt64("unknown", 32)
	assert.Empty(t, err)
	assert.Equal(t, v, int64(32))
}

func TestQueryParams_GetUint(t *testing.T) {
	params := NewQueryParams(mockUrl("var=125&foo=?"))
	v, _ := params.GetUint("var")
	assert.Equal(t, v, uint(125))
	f, err := params.GetUint("foo")
	assert.Empty(t, f)
	assert.NotEmpty(t, err)
	v, err = params.GetUint("unknown")
	assert.Empty(t, v)
	assert.NotEmpty(t, err)
	v, err = params.GetUint("unknown")
	assert.Empty(t, v)
	assert.NotEmpty(t, err)
	v, err = params.GetUint("unknown", 32)
	assert.Empty(t, err)
	assert.Equal(t, v, uint(32))
}

func TestQueryParams_GetUint8(t *testing.T) {
	params := NewQueryParams(mockUrl("var=125&foo=?"))
	v, _ := params.GetUint8("var")
	assert.Equal(t, v, uint8(125))
	f, err := params.GetUint8("foo")
	assert.Empty(t, f)
	assert.NotEmpty(t, err)
	v, err = params.GetUint8("unknown")
	assert.Empty(t, v)
	assert.NotEmpty(t, err)
	v, err = params.GetUint8("unknown")
	assert.Empty(t, v)
	assert.NotEmpty(t, err)
	v, err = params.GetUint8("unknown", 32)
	assert.Empty(t, err)
	assert.Equal(t, v, uint8(32))
}

func TestQueryParams_GetUint16(t *testing.T) {
	params := NewQueryParams(mockUrl("var=125&foo=?"))
	v, _ := params.GetUint16("var")
	assert.Equal(t, v, uint16(125))
	f, err := params.GetUint16("foo")
	assert.Empty(t, f)
	assert.NotEmpty(t, err)
	v, err = params.GetUint16("unknown")
	assert.Empty(t, v)
	assert.NotEmpty(t, err)
	v, err = params.GetUint16("unknown")
	assert.Empty(t, v)
	assert.NotEmpty(t, err)
	v, err = params.GetUint16("unknown", 32)
	assert.Empty(t, err)
	assert.Equal(t, v, uint16(32))
}

func TestQueryParams_GetUint32(t *testing.T) {
	params := NewQueryParams(mockUrl("var=12590&foo=?"))
	v, _ := params.GetUint32("var")
	assert.Equal(t, v, uint32(12590))
	f, err := params.GetUint32("foo")
	assert.Empty(t, f)
	assert.NotEmpty(t, err)
	v, err = params.GetUint32("unknown")
	assert.Empty(t, v)
	assert.NotEmpty(t, err)
	v, err = params.GetUint32("unknown")
	assert.Empty(t, v)
	assert.NotEmpty(t, err)
	v, err = params.GetUint32("unknown", 12590)
	assert.Empty(t, err)
	assert.Equal(t, v, uint32(12590))
}

func TestQueryParams_GetUint64(t *testing.T) {
	params := NewQueryParams(mockUrl("var=125900&foo=?"))
	v, _ := params.GetUint64("var")
	assert.Equal(t, v, uint64(125900))
	f, err := params.GetUint64("foo")
	assert.Empty(t, f)
	assert.NotEmpty(t, err)
	v, err = params.GetUint64("unknown")
	assert.Empty(t, v)
	assert.NotEmpty(t, err)
	v, err = params.GetUint64("unknown")
	assert.Empty(t, v)
	assert.NotEmpty(t, err)
	v, err = params.GetUint64("unknown", 125900)
	assert.Empty(t, err)
	assert.Equal(t, v, uint64(125900))
}
