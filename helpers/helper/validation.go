package helper

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/alfisar/jastip-import/domain"
	validator "github.com/alfisar/jastip-import/helpers/validation"
	"github.com/disintegration/imaging"
	"github.com/minio/minio-go/v7"

	validation "github.com/go-ozzo/ozzo-validation"
)

func ValidationDataUser(data domain.User) (err error) {
	err = validation.ValidateStruct(
		&data,
		validation.Field(&data.Email, validator.Required, validator.Email),
		validation.Field(&data.FullName, validator.Required, validator.AlphanumericSimbols),
		validation.Field(&data.NoHP, validator.Required, validator.Numeric),
		validation.Field(&data.Username, validator.Required, validator.AlphanumericSimbols),
		validation.Field(&data.Password, validator.Required, validator.AlphanumericSimbols),
	)
	return
}

func ValidationDataUserVerifyOTP(data domain.UserVerifyOtpRequest) (err error) {
	err = validation.ValidateStruct(
		&data,
		validation.Field(&data.Email, validator.Required, validator.Email),
		validation.Field(&data.NoHP, validator.Required, validator.Numeric),
	)
	return
}

func ValidationDataUserResendOTP(data domain.UserResendOtpRequest) (err error) {
	err = validation.ValidateStruct(
		&data,
		validation.Field(&data.Email, validator.Required, validator.Email),
		validation.Field(&data.NoHP, validator.Required, validator.Numeric),
	)
	return
}

func ValidationLogin(data domain.UserLoginRequest) (err error) {

	_, errs := strconv.Atoi(data.Username)
	if errs != nil {
		err = validation.ValidateStruct(
			&data,
			validation.Field(&data.Username, validator.Email),
			validation.Field(&data.Password, validator.Required, validator.AlphanumericSimbols),
		)
	} else {
		err = validation.ValidateStruct(
			&data,
			validation.Field(&data.Username, validator.Numeric),
			validation.Field(&data.Password, validator.Required, validator.AlphanumericSimbols),
		)
	}

	return
}

func filterRequestBody(data map[string]any, allowedKeys []string) map[string]any {
	filtered := make(map[string]any)
	for _, key := range allowedKeys {
		if val, exists := data[key]; exists {
			filtered[key] = val
		}
	}
	return filtered
}

func ValidationUpdateProfile(data map[string]any) (err error) {
	var (
		rules validation.Rule
	)
	allowedKeys := []string{
		"full_name",
		"email",
		"nohp",
		"password",
	}

	mappingData := filterRequestBody(data, allowedKeys)
	if len(mappingData) == 0 {
		err = fmt.Errorf("data cannot empty")
		return
	}
	for key, v := range mappingData {

		if key == "nohp" {
			rules = validator.Numeric
		} else if key == "email" {
			rules = validator.Email
		} else {
			rules = validator.AlphanumericSimbols
		}
		err = ValidateMappingData(v, key, err, rules)
	}

	return
}

func ValidationAddress(data map[string]any) (err error) {
	var (
		rules validation.Rule
	)
	allowedKeys := []string{
		"street",
		"city",
		"district",
		"subdistrict",
		"postalcode",
	}

	mappingData := filterRequestBody(data, allowedKeys)
	if len(mappingData) == 0 {
		err = fmt.Errorf("data cannot empty")
		return
	}
	for key, v := range mappingData {

		if key == "postalcode" {
			rules = validator.Numeric
		} else {
			rules = validator.AlphanumericSimbols
		}
		err = ValidateMappingData(v, key, err, rules)
	}

	return
}

func ValidateMappingData(data any, key string, errs error, rules validation.Rule) (err error) {

	err = validation.Validate(data, rules)
	if err != nil {
		err = fmt.Errorf(key + ": " + err.Error())
		if errs != nil {

			err = fmt.Errorf(errs.Error()[:len(errs.Error())-1] + "; " + err.Error())
		}
	} else {
		err = errs
	}

	return
}

func ValidationPostSchedule(data domain.TravelSchRequest) (err error) {
	dataDummy := struct {
		Location    string
		PeriodStart string
		PeriodEnd   string
	}{
		Location:    strconv.Itoa(data.Location),
		PeriodStart: data.PeriodStart,
		PeriodEnd:   data.PeriodEnd,
	}
	err = validation.ValidateStruct(
		&dataDummy,
		validation.Field(&dataDummy.Location, validator.Required, validator.Numeric),
		validation.Field(&dataDummy.PeriodEnd, validator.Required),
		validation.Field(&dataDummy.PeriodStart, validator.Required),
	)

	return
}

func ValidationUpdateTravelSch(data map[string]any) (err error) {
	var (
		rules validation.Rule
	)
	allowedKeys := []string{
		"locations",
	}

	mappingData := filterRequestBody(data, allowedKeys)
	if len(mappingData) == 0 {
		return
	}

	for key, v := range mappingData {

		rules = validator.AlphanumericSimbols

		err = ValidateMappingData(v, key, err, rules)
	}

	return
}

func ValidationPostProducts(data domain.ProductData) (err error) {
	dataDummy := struct {
		Name     string
		Price    string
		Quantity string
	}{
		Name:     data.Name,
		Price:    strconv.Itoa(int(data.Price)),
		Quantity: strconv.Itoa(data.Quantity),
	}
	err = validation.ValidateStruct(
		&dataDummy,
		validation.Field(&dataDummy.Name, validator.Required, validator.AlphanumericPetik),
		validation.Field(&dataDummy.Price, validator.Required, validator.Numeric),
		validation.Field(&dataDummy.Quantity, validator.Required, validator.Numeric),
	)

	return
}

func ValidationUpdateProduct(data map[string]any) (err error) {
	var (
		rules validation.Rule
	)
	allowedKeys := []string{
		"name",
		"price",
		"quantity",
		"desc",
		"image",
	}

	mappingData := filterRequestBody(data, allowedKeys)
	if len(mappingData) == 0 {
		return
	}

	for key, v := range mappingData {
		rules = &validation.MatchRule{}
		if key == "price" || key == "quantity" {
			rules = validator.Numeric
		} else {
			rules = validator.AlphanumericPetik
		}

		err = ValidateMappingData(v, key, err, rules)
	}

	return
}

func SaveImageMinio(ctx context.Context, config *domain.Config, PathImage string, fileHeader *multipart.FileHeader) (name string, err error) {

	var (
		pattern string
		// tempFile *os.File
		reader              *bytes.Reader
		file                multipart.File
		compressedImgBuffer bytes.Buffer

		slicedName []string
	)

	file, _ = fileHeader.Open()
	defer file.Close()

	fileSize := fileHeader.Size

	if fileSize > 10*1024*1024 {
		return name, fmt.Errorf("Invalid size image")
	} else if fileSize > 1*1024*1024 && fileSize < 5*1024*1024 {
		img, _, err := image.Decode(file)
		if err != nil {
			return name, err
		}

		for quality := 90; quality > 0; quality = quality - 5 {
			compressedImgBuffer.Reset()
			err = imaging.Encode(&compressedImgBuffer, img, imaging.JPEG, imaging.JPEGQuality(quality))
			if err != nil {
				return name, err
			}
			compressedSize := int(compressedImgBuffer.Len())
			if int64(compressedSize) <= 1024*1024 {
				fileSize = int64(compressedSize)
				break
			}
		}
	}

	switch fileHeader.Header.Values("Content-Type")[0] {

	case validator.ImageJPG, validator.ImageJPEG, validator.ImagePNG: // continue
	default:
		return name, validator.ErrInvalidImageType
	}

	allowedExtensions := []string{".jpg", ".jpeg", ".png"}
	ext := filepath.Ext(fileHeader.Filename)
	validExtension := false

	for _, allowedExt := range allowedExtensions {
		if strings.EqualFold(ext, allowedExt) {
			validExtension = true
			break
		}
	}

	if !validExtension {
		return name, validator.ErrInvalidImageType
	}

	isImage := isImageFile(fileHeader)

	if !isImage {
		return name, validator.ErrInvalidImageType
	}

	removeSpace := strings.ReplaceAll(fileHeader.Filename, " ", "")
	pattern = fmt.Sprintf("*_%s", removeSpace)

	if compressedImgBuffer.Len() > 0 {
		reader = bytes.NewReader(compressedImgBuffer.Bytes())
	} else {
		// Reset file reader karena sudah dibaca sebelumnya
		_, err := file.Seek(0, 0)
		if err != nil {
			return "", err
		}
		reader = bytes.NewReader(fileHeaderHeaderToBytes(fileHeader))
	}
	_, errData := config.Minio.Client.PutObject(ctx, config.Minio.BucketName, PathImage+pattern, reader, fileSize, minio.PutObjectOptions{})

	if errData != nil {
		fmt.Println(errData)
		return "", errData
	}
	// return tempFile.Name(), nil
	return slicedName[len(slicedName)-1], nil
}

func fileHeaderHeaderToBytes(fh *multipart.FileHeader) []byte {
	file, _ := fh.Open()
	defer file.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	return buf.Bytes()
}

func isImageFile(fileHeader *multipart.FileHeader) bool {
	file, err := fileHeader.Open()
	if err != nil {
		return false
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return false
	}

	fileType := http.DetectContentType(buffer)
	return strings.HasPrefix(fileType, "image/")
}
