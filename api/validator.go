package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/lucianocorreia/simplebank/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSuppoertedCurrency(currency)
	}

	return false
}
