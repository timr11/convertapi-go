package convertapi

import (
	"github.com/timr11/convertapi-go/config"
	"github.com/timr11/convertapi-go/param"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

func TestSetup(t *testing.T) {
	config.Default.Secret = os.Getenv("CONVERTAPI_SECRET")
	assert.Equal(t, config.Default.Secret, os.Getenv("CONVERTAPI_SECRET"))
}

func TestConvertPath(t *testing.T) {
	config.Default.Secret = os.Getenv("CONVERTAPI_SECRET")
	file, errs := ConvertPath("assets/test.docx", path.Join(os.TempDir(), "convertapi-test.pdf"))

	assert.Nil(t, errs)
	assert.NotEmpty(t, file.Name())
}

func TestTokenAuth(t *testing.T) {
	config.Default.Token = os.Getenv("CONVERTAPI_TOKEN")
	config.Default.ApiKey = os.Getenv("CONVERTAPI_APIKEY")
	file, errs := ConvertPath("assets/test.docx", path.Join(os.TempDir(), "convertapi-test.pdf"))

	assert.Nil(t, errs)
	assert.NotEmpty(t, file.Name())
}

func TestChained(t *testing.T) {
	config.Default.Secret = os.Getenv("CONVERTAPI_SECRET")
	jpgRes := Convert("docx", "jpg", []param.IParam{
		param.NewPath("file", "assets/test.docx", nil),
	}, nil)

	zipRes := Convert("any", "zip", []param.IParam{
		param.NewResult("files", jpgRes, nil),
	}, nil)

	zipCost, err := zipRes.Cost()

	assert.Nil(t, err)
	assert.NotEmpty(t, zipCost)

	files, err := zipRes.Files()

	assert.Nil(t, err)
	assert.Equal(t, "test.zip", files[0].FileName)
}

func TestUserInfo(t *testing.T) {
	config.Default.Secret = os.Getenv("CONVERTAPI_SECRET")
	user, err := UserInfo(nil)

	assert.Nil(t, err)
	assert.NotEmpty(t, user.SecondsLeft)
}
