package util

import (
	"fmt"
	"net/http"

	stdjson "encoding/json"

	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
)

type JSONIterSerializer struct{}

func (j JSONIterSerializer) Serialize(c echo.Context, i interface{}, indent string) error {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	enc := json.NewEncoder(c.Response())
	if indent != "" {
		enc.SetIndent("", indent)
	}

	return enc.Encode(i)
}

func (j JSONIterSerializer) Deserialize(c echo.Context, i interface{}) error {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.NewDecoder(c.Request().Body).Decode(i)

	if ute, ok := err.(*stdjson.UnmarshalTypeError); ok {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unmarshal type error: expected=%v, got=%v, field=%v, offset=%v", ute.Type, ute.Value, ute.Field, ute.Offset)).SetInternal(err)
	} else if se, ok := err.(*stdjson.SyntaxError); ok {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Syntax error: offset=%v, error=%v", se.Offset, se.Error())).SetInternal(err)
	}

	return err
}
