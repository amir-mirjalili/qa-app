package userValidator

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"qa-app/pkg/errmsg"
	"qa-app/pkg/richerror"
	"qa-app/request/user"
	"regexp"
)

func UserValidateRegisterRequest(req user.CreateUserRequest) (map[string]string, error) {
	const op = "uservalidator.ValidateRegisterRequest"

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Name,
			validation.Required,
			validation.Length(3, 50)),

		validation.Field(&req.Password,
			validation.Required,
			validation.Match(regexp.MustCompile(`^[A-Za-z0-9!@#%^&*]{4,}$`))),

		validation.Field(&req.PhoneNumber,
			validation.Required,
			validation.Match(regexp.MustCompile(phoneNumberRegex)).Error(errmsg.ErrorMsgPhoneNumberIsNotValid),
		),
	); err != nil {
		fieldErrors := make(map[string]string)

		errV, ok := err.(validation.Errors)
		if ok {
			for key, value := range errV {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}

		return fieldErrors, richerror.New(op).WithMessage(errmsg.ErrorMsgInvalidInput).
			WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req}).WithErr(err)
	}

	return nil, nil
}

func UserValidateLoginRequest(req user.UserLoginReq) (map[string]string, error) {
	const op = "uservalidator.ValidateLoginRequest"

	if err := validation.ValidateStruct(&req,

		validation.Field(&req.PhoneNumber,
			validation.Required,
			validation.Match(regexp.MustCompile(phoneNumberRegex)).Error(errmsg.ErrorMsgPhoneNumberIsNotValid)),
		validation.Field(&req.Password, validation.Required),
	); err != nil {
		fieldErrors := make(map[string]string)

		var errV validation.Errors
		ok := errors.As(err, &errV)
		if ok {
			for key, value := range errV {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}

		return fieldErrors, richerror.New(op).WithMessage(errmsg.ErrorMsgInvalidInput).
			WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req}).WithErr(err)
	}

	return nil, nil
}
