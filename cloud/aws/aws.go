package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/fatih/color"
	"lsync/cloud"
	"os"
)


type LogSVC struct {
	s3     *s3.S3
	cw     *cloudwatch.CloudWatch
	s3Mgr  *s3manager.Uploader
	bucket string
}


func New(region, bucket string) (cloud.AWS, error) {

	cfg := &aws.Config{
		Credentials:                    credentials.NewSharedCredentials("", "default"),
		Region:                         aws.String(region),
		DisableRestProtocolURICleaning: aws.Bool(true),
	}

	sess, err := session.NewSession(cfg)
	if err != nil {
		color.Red("SESSION_ERROR: %s\n", err)
		os.Exit(1)
	}

	s3Mgr := s3manager.NewUploader(sess)
	s3Client := s3.New(sess)
	cwClient := cloudwatch.New(sess)

	return LogSVC{s3Client, cwClient, s3Mgr, bucket}, nil
}

func (svc LogSVC) UploadFileToS3(f *os.File) error {


	key := fmt.Sprintf(f.Name())
	uploadInput := &s3manager.UploadInput{
		Body:   f,
		Bucket: aws.String(svc.bucket),
		//ContentType:               nil,
		Key: aws.String(key),
	}

	resp, err := svc.s3Mgr.Upload(uploadInput)
	if err != nil {
		return err
	}

	color.Green("file %s upload to: %s", f.Name(), resp.Location)


	return nil
}

func (svc LogSVC) SyncDirToS3(dirName string) error {

	return nil
}

func (svc LogSVC) SyncLogsToCW(fileName string) error {
	return nil
}

func (svc LogSVC) DeleteFileFromS3(f string) error {

	//a := &s3.DeleteObjectInput{
	//	Bucket:                    aws.String(svc.bucket),
	//	Key:                       aws.String(f),
	//}
	//svc.DeleteFileFromS3(s3.DeleteObjectInput{
	//	Bucket:                    aws.String(svc.bucket),
	//	//Key:                       aws.String(),
	//})
		a := s3.DeleteObjectInput{
			Bucket: aws.String(svc.bucket),
			Key:    aws.String(f),
		}

return nil

}


func (svc LogSVC) DeleteCWLogs(cwLogName string) error {

	return nil
}
