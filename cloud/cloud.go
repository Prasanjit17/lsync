package cloud

import (
	"os"
)

type AWS interface {
	UploadFileToS3(f *os.File) error
	SyncDirToS3(dirName string) error
	DeleteFileFromS3(file string) error
	SyncLogsToCW(fileName string) error
	DeleteCWLogs(cwLogName string) error
}
