package validations

import (
	"fmt"
	"regexp"

	"github.com/alfisar/jastip-import/helpers/consts"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

var (
	AlphanumericSimbols       = validation.Match(regexp.MustCompile(consts.RegexAlphanumericSimbols)).Error(consts.AlphanumericSimbols)
	AlphanumericPetik         = validation.Match(regexp.MustCompile(consts.RegexAlphanumericPetik)).Error(consts.AlphanumericPetik)
	Numeric                   = is.Digit.Error(consts.Digit)
	MaxMinChar17              = validation.Length(17, 17).Error(consts.MaxMinChar17)
	Required                  = validation.Required.Error(consts.RequiredField)
	Email                     = is.Email.Error(consts.IsEmail)
	ErrInvalidImageType error = fmt.Errorf("Format gambar tidak sesuai")

	ImageJPG  string = "image/jpg"
	ImageJPEG string = "image/jpeg"
	ImagePNG  string = "image/png"
)
