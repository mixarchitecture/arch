package validator

import (
	"regexp"

	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func validateUserName(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(userNameRegexp, fl.Field().String())
	return matched
}

func validatePassword(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(passwordRegexp, fl.Field().String())
	return matched
}

func validateSlug(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(slugRegexp, fl.Field().String())
	return matched
}

func validateLocale(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(localeRegexp, fl.Field().String())
	return matched
}

func validateObjectId(fl validator.FieldLevel) bool {
	return primitive.IsValidObjectID(fl.Field().String())
}

func validateGender(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(genderRegexp, fl.Field().String())
	return matched
}
