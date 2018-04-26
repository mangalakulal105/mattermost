// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package utils

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	l4g "github.com/alecthomas/log4go"
	s3 "github.com/minio/minio-go"
	"github.com/minio/minio-go/pkg/credentials"

	"github.com/mattermost/mattermost-server/model"
)

type S3FileBackend struct {
	endpoint  string
	accessKey string
	secretKey string
	secure    bool
	signV2    bool
	region    string
	bucket    string
	encrypt   bool
	trace     bool
}

// Similar to s3.New() but allows initialization of signature v2 or signature v4 client.
// If signV2 input is false, function always returns signature v4.
//
// Additionally this function also takes a user defined region, if set
// disables automatic region lookup.
func (b *S3FileBackend) s3New() (*s3.Client, error) {
	var creds *credentials.Credentials

	if b.accessKey == "" && b.secretKey == "" {
		creds = credentials.NewIAM("")
	} else if b.signV2 {
		creds = credentials.NewStatic(b.accessKey, b.secretKey, "", credentials.SignatureV2)
	} else {
		creds = credentials.NewStatic(b.accessKey, b.secretKey, "", credentials.SignatureV4)
	}

	s3Clnt, err := s3.NewWithCredentials(b.endpoint, creds, b.secure, b.region)
	if err != nil {
		return nil, err
	}

	if b.trace {
		s3Clnt.TraceOn(os.Stdout)
	}

	return s3Clnt, nil
}

func (b *S3FileBackend) TestConnection() *model.AppError {
	s3Clnt, err := b.s3New()
	if err != nil {
		return model.NewAppError("TestFileConnection", "Bad connection to S3 or minio.", nil, err.Error(), http.StatusInternalServerError)
	}

	exists, err := s3Clnt.BucketExists(b.bucket)
	if err != nil {
		return model.NewAppError("TestFileConnection", "Error checking if bucket exists.", nil, err.Error(), http.StatusInternalServerError)
	}

	if !exists {
		l4g.Warn("Bucket specified does not exist. Attempting to create...")
		err := s3Clnt.MakeBucket(b.bucket, b.region)
		if err != nil {
			l4g.Error("Unable to create bucket.")
			return model.NewAppError("TestFileConnection", "Unable to create bucket", nil, err.Error(), http.StatusInternalServerError)
		}
	}
	l4g.Info("Connection to S3 or minio is good. Bucket exists.")
	return nil
}

func (b *S3FileBackend) ReadFile(path string) ([]byte, *model.AppError) {
	s3Clnt, err := b.s3New()
	if err != nil {
		return nil, model.NewAppError("ReadFile", "api.file.read_file.s3.app_error", nil, err.Error(), http.StatusInternalServerError)
	}
	minioObject, err := s3Clnt.GetObject(b.bucket, path, s3.GetObjectOptions{})
	if err != nil {
		return nil, model.NewAppError("ReadFile", "api.file.read_file.s3.app_error", nil, err.Error(), http.StatusInternalServerError)
	}
	defer minioObject.Close()
	if f, err := ioutil.ReadAll(minioObject); err != nil {
		return nil, model.NewAppError("ReadFile", "api.file.read_file.s3.app_error", nil, err.Error(), http.StatusInternalServerError)
	} else {
		return f, nil
	}
}

func (b *S3FileBackend) CopyFile(oldPath, newPath string) *model.AppError {
	s3Clnt, err := b.s3New()
	if err != nil {
		return model.NewAppError("copyFile", "api.file.write_file.s3.app_error", nil, err.Error(), http.StatusInternalServerError)
	}

	source := s3.NewSourceInfo(b.bucket, oldPath, nil)
	destination, err := s3.NewDestinationInfo(b.bucket, newPath, nil, s3CopyMetadata(b.encrypt))
	if err != nil {
		return model.NewAppError("copyFile", "api.file.write_file.s3.app_error", nil, err.Error(), http.StatusInternalServerError)
	}
	if err = s3Clnt.CopyObject(destination, source); err != nil {
		return model.NewAppError("copyFile", "api.file.move_file.copy_within_s3.app_error", nil, err.Error(), http.StatusInternalServerError)
	}
	return nil
}

func (b *S3FileBackend) MoveFile(oldPath, newPath string) *model.AppError {
	s3Clnt, err := b.s3New()
	if err != nil {
		return model.NewAppError("moveFile", "api.file.write_file.s3.app_error", nil, err.Error(), http.StatusInternalServerError)
	}

	source := s3.NewSourceInfo(b.bucket, oldPath, nil)
	destination, err := s3.NewDestinationInfo(b.bucket, newPath, nil, s3CopyMetadata(b.encrypt))
	if err != nil {
		return model.NewAppError("moveFile", "api.file.write_file.s3.app_error", nil, err.Error(), http.StatusInternalServerError)
	}
	if err = s3Clnt.CopyObject(destination, source); err != nil {
		return model.NewAppError("moveFile", "api.file.move_file.copy_within_s3.app_error", nil, err.Error(), http.StatusInternalServerError)
	}
	if err = s3Clnt.RemoveObject(b.bucket, oldPath); err != nil {
		return model.NewAppError("moveFile", "api.file.move_file.delete_from_s3.app_error", nil, err.Error(), http.StatusInternalServerError)
	}
	return nil
}

func (b *S3FileBackend) WriteFile(fr io.Reader, path string) (int64, *model.AppError) {
	s3Clnt, err := b.s3New()
	if err != nil {
		return 0, model.NewAppError("WriteFile", "api.file.write_file.s3.app_error", nil, err.Error(), http.StatusInternalServerError)
	}

	var contentType string
	if ext := filepath.Ext(path); model.IsFileExtImage(ext) {
		contentType = model.GetImageMimeType(ext)
	} else {
		contentType = "binary/octet-stream"
	}

	options := s3PutOptions(b.encrypt, contentType)
	written, err := s3Clnt.PutObject(b.bucket, path, fr, -1, options)
	if err != nil {
		return 0, model.NewAppError("WriteFile", "api.file.write_file.s3.app_error", nil, err.Error(), http.StatusInternalServerError)
	}

	return written, nil
}

func (b *S3FileBackend) RemoveFile(path string) *model.AppError {
	s3Clnt, err := b.s3New()
	if err != nil {
		return model.NewAppError("RemoveFile", "utils.file.remove_file.s3.app_error", nil, err.Error(), http.StatusInternalServerError)
	}

	if err := s3Clnt.RemoveObject(b.bucket, path); err != nil {
		return model.NewAppError("RemoveFile", "utils.file.remove_file.s3.app_error", nil, err.Error(), http.StatusInternalServerError)
	}

	return nil
}

func getPathsFromObjectInfos(in <-chan s3.ObjectInfo) <-chan string {
	out := make(chan string, 1)

	go func() {
		defer close(out)

		for {
			info, done := <-in

			if !done {
				break
			}

			out <- info.Key
		}
	}()

	return out
}

func (b *S3FileBackend) ListDirectory(path string) (*[]string, *model.AppError) {
	var paths []string

	s3Clnt, err := b.s3New()
	if err != nil {
		return nil, model.NewAppError("ListDirectory", "utils.file.list_directory.s3.app_error", nil, err.Error(), http.StatusInternalServerError)
	}

	doneCh := make(chan struct{})

	defer close(doneCh)

	for object := range s3Clnt.ListObjects(b.bucket, path, false, doneCh) {
		if object.Err != nil {
			return nil, model.NewAppError("ListDirectory", "utils.file.list_directory.s3.app_error", nil, object.Err.Error(), http.StatusInternalServerError)
		}
		paths = append(paths, strings.Trim(object.Key, "/"))
	}

	return &paths, nil
}

func (b *S3FileBackend) RemoveDirectory(path string) *model.AppError {
	s3Clnt, err := b.s3New()
	if err != nil {
		return model.NewAppError("RemoveDirectory", "utils.file.remove_directory.s3.app_error", nil, err.Error(), http.StatusInternalServerError)
	}

	doneCh := make(chan struct{})

	for err := range s3Clnt.RemoveObjects(b.bucket, getPathsFromObjectInfos(s3Clnt.ListObjects(b.bucket, path, true, doneCh))) {
		if err.Err != nil {
			doneCh <- struct{}{}
			return model.NewAppError("RemoveDirectory", "utils.file.remove_directory.s3.app_error", nil, err.Err.Error(), http.StatusInternalServerError)
		}
	}

	close(doneCh)
	return nil
}

func s3PutOptions(encrypt bool, contentType string) s3.PutObjectOptions {
	options := s3.PutObjectOptions{}
	if encrypt {
		options.UserMetadata = make(map[string]string)
		options.UserMetadata["x-amz-server-side-encryption"] = "AES256"
	}
	options.ContentType = contentType

	return options
}

func s3CopyMetadata(encrypt bool) map[string]string {
	metaData := make(map[string]string)
	metaData["x-amz-server-side-encryption"] = "AES256"
	return metaData
}

func CheckMandatoryS3Fields(settings *model.FileSettings) *model.AppError {
	if len(settings.AmazonS3Bucket) == 0 {
		return model.NewAppError("S3File", "api.admin.test_s3.missing_s3_bucket", nil, "", http.StatusBadRequest)
	}

	// if S3 endpoint is not set call the set defaults to set that
	if len(settings.AmazonS3Endpoint) == 0 {
		settings.SetDefaults()
	}

	return nil
}
