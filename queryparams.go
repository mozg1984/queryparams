package queryparams

import (
	"net/url"
	"strconv"
)

// QueryParams describes methods that are available over parameters.
type QueryParams interface {
	Exists(key string) bool
	GetString(key string, def ...string) string
	GetStrings(key string, def ...[]string) []string
	GetInt(key string, def ...int) (int, error)
	GetInt8(key string, def ...int8) (int8, error)
	GetInt16(key string, def ...int16) (int16, error)
	GetInt32(key string, def ...int32) (int32, error)
	GetInt64(key string, def ...int64) (int64, error)
	GetUint(key string, def ...uint) (uint, error)
	GetUint8(key string, def ...uint8) (uint8, error)
	GetUint16(key string, def ...uint16) (uint16, error)
	GetUint32(key string, def ...uint32) (uint32, error)
	GetUint64(key string, def ...uint64) (uint64, error)
	GetBool(key string, def ...bool) (bool, error)
	GetFloat(key string, def ...float64) (float64, error)
}

// queryParams contains all url query parameters.
type queryParams struct {
	values url.Values
}

// NewQueryParams returns an instance
func NewQueryParams(url *url.URL) QueryParams {
	return &queryParams{values: url.Query()}
}

// Exists checks if the parameter exists by given key.
func (p *queryParams) Exists(key string) bool {
	return len(p.values[key]) > 0
}

// GetString returns parameter value by key string or the default value if it's given and the parameter value is empty.
func (p *queryParams) GetString(key string, def ...string) string {
	if p.Exists(key) {
		return p.values[key][0]
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}

// GetStrings returns the slice of parameter values by given key or the default value if it's given and the parameter values are empty
// to support working with multi-value parameters.
func (p *queryParams) GetStrings(key string, def ...[]string) []string {
	if !p.Exists(key) && len(def) > 0 {
		return def[0]
	}
	return p.values[key]
}

// GetInt returns parameter value as int or the default value if it's given and the parameter value is empty.
func (p *queryParams) GetInt(key string, def ...int) (int, error) {
	if !p.Exists(key) && len(def) > 0 {
		return def[0], nil
	}
	return strconv.Atoi(p.GetString(key))
}

// GetInt8 returns parameter value as int8 or the default value if it's given and the parameter value is empty.
func (p *queryParams) GetInt8(key string, def ...int8) (int8, error) {
	if !p.Exists(key) && len(def) > 0 {
		return def[0], nil
	}
	val, err := strconv.ParseInt(p.GetString(key), 10, 8)
	return int8(val), err
}

// GetInt16 returns parameter value as int16 or the default value if it's given and the parameter value is empty.
func (p *queryParams) GetInt16(key string, def ...int16) (int16, error) {
	if !p.Exists(key) && len(def) > 0 {
		return def[0], nil
	}
	val, err := strconv.ParseInt(p.GetString(key), 10, 16)
	return int16(val), err
}

// GetInt32 returns parameter value as int32 or the default value if it's given and the parameter value is empty.
func (p *queryParams) GetInt32(key string, def ...int32) (int32, error) {
	if !p.Exists(key) && len(def) > 0 {
		return def[0], nil
	}
	val, err := strconv.ParseInt(p.GetString(key), 10, 32)
	return int32(val), err
}

// GetInt64 returns parameter value as int64 or the default value if it's given and the parameter value is empty.
func (p *queryParams) GetInt64(key string, def ...int64) (int64, error) {
	if !p.Exists(key) && len(def) > 0 {
		return def[0], nil
	}
	return strconv.ParseInt(p.GetString(key), 10, 64)
}

// GetUint returns parameter value as uint or the default value if it's given and the parameter value is empty.
func (p *queryParams) GetUint(key string, def ...uint) (uint, error) {
	if !p.Exists(key) && len(def) > 0 {
		return def[0], nil
	}
	val, err := strconv.ParseUint(p.GetString(key), 10, 0)
	return uint(val), err
}

// GetUint8 returns parameter value as uint8 or the default value if it's given and the parameter value is empty.
func (p *queryParams) GetUint8(key string, def ...uint8) (uint8, error) {
	if !p.Exists(key) && len(def) > 0 {
		return def[0], nil
	}
	val, err := strconv.ParseUint(p.GetString(key), 10, 8)
	return uint8(val), err
}

// GetUint16 returns parameter value as uint16 or the default value if it's given and the parameter value is empty.
func (p *queryParams) GetUint16(key string, def ...uint16) (uint16, error) {
	if !p.Exists(key) && len(def) > 0 {
		return def[0], nil
	}
	val, err := strconv.ParseUint(p.GetString(key), 10, 16)
	return uint16(val), err
}

// GetUint32 returns parameter value as uint32 or the default value if it's given and the parameter value is empty.
func (p *queryParams) GetUint32(key string, def ...uint32) (uint32, error) {
	if !p.Exists(key) && len(def) > 0 {
		return def[0], nil
	}
	val, err := strconv.ParseUint(p.GetString(key), 10, 32)
	return uint32(val), err
}

// GetUint64 returns parameter value as uint64 or the default value if it's given and the parameter value is empty.
func (p *queryParams) GetUint64(key string, def ...uint64) (uint64, error) {
	if !p.Exists(key) && len(def) > 0 {
		return def[0], nil
	}
	return strconv.ParseUint(p.GetString(key), 10, 64)
}

// GetBool returns parameter value as bool or the default value if it's given and the parameter value is empty.
func (p *queryParams) GetBool(key string, def ...bool) (bool, error) {
	if !p.Exists(key) && len(def) > 0 {
		return def[0], nil
	}
	return strconv.ParseBool(p.GetString(key))
}

// GetFloat returns parameter value as float64 or the default value if it's given and the parameter value is empty.
func (p *queryParams) GetFloat(key string, def ...float64) (float64, error) {
	if !p.Exists(key) && len(def) > 0 {
		return def[0], nil
	}
	return strconv.ParseFloat(p.GetString(key), 64)
}
