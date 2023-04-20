package s3client

import (
	"bytes"
	"fmt"
	"net/url"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	log "github.com/sirupsen/logrus"
)

type S3Client struct {
	sess *session.Session
}

func NewS3Client() *S3Client {
	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession())

	return &S3Client{
		sess: sess,
	}
}

func parse(location string) (string, string, error) {
	parsedURL, err := url.Parse(location)
	if err != nil {
		return "", "", err
	}
	bkt := parsedURL.Host
	key := strings.TrimPrefix(strings.TrimSuffix(parsedURL.Path, "/"), "/")
	return bkt, key, nil
}

func (s *S3Client) Upload(data []byte, path, region string) error {
	bkt, key, err := parse(path)
	if err != nil {
		return err
	}

	// Create an uploader with the session and default options
	s.sess.Config.Region = aws.String(region)
	uploader := s3manager.NewUploader(s.sess)

	body := bytes.NewReader(data)

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bkt),
		Key:    aws.String(key),
		Body:   body,
	})
	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}
	log.Infof("file uploaded to, %s\n", result.Location)

	return nil
}

func (s *S3Client) Download(path, region string) ([]byte, error) {
	bkt, key, err := parse(path)
	if err != nil {
		return nil, err
	}

	// Create a downloader with the session and default options
	s.sess.Config.Region = aws.String(region)
	downloader := s3manager.NewDownloader(s.sess)

	buffer := aws.NewWriteAtBuffer([]byte{})

	// Write the contents of S3 Object to the file
	n, err := downloader.Download(buffer, &s3.GetObjectInput{
		Bucket: aws.String(bkt),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to download file, %v", err)
	}
	log.Infof("file downloaded, %d bytes\n", n)

	return buffer.Bytes(), nil
}
