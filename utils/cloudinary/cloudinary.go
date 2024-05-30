package cloudinary

import (
	"capstone-mikti/configs"
	"context"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/sirupsen/logrus"
)

type CloudinaryInterface interface {
	UploadImageHelper(input interface{}) (string, string, error)
	DeleteImageHelper(publicID string) (string, error)
}

type Cloudinary struct {
	cfg configs.ProgrammingConfig
}

func InitCloud(config configs.ProgrammingConfig) CloudinaryInterface {
	return &Cloudinary{
		cfg: config,
	}
}

func (cld *Cloudinary) UploadImageHelper(input interface{}) (string, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Second)

	defer cancel()

	cl, err := cloudinary.NewFromURL(cld.cfg.Cloud_URL)
	logrus.Info("Ini Input : ", input)

	if err != nil {
		logrus.Error("ERROR Cloudinary Connection : ", err)
		return "", "", err
	}

	uploadParam, err := cl.Upload.Upload(ctx, input, uploader.UploadParams{Folder: "capstone-mikti"})
	if err != nil {
		logrus.Error("ERROR Upload : ", err)
		return "", "", err
	}

	return uploadParam.SecureURL, uploadParam.PublicID, nil
}

func (cld *Cloudinary) DeleteImageHelper(publicID string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Second)

	defer cancel()

	cl, err := cloudinary.NewFromURL(cld.cfg.Cloud_URL)

	if err != nil {
		logrus.Error("ERROR Cloudinary Connection : ", err)
		return "", err
	}

	resp, err := cl.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: publicID,
	})

	if err != nil {
		logrus.Error("ERROR Delete : ", err)
		return "", err
	}

	return resp.Result, nil
}
