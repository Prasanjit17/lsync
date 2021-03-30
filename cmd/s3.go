package cmd

import (
	"fmt"
	"lsync/cloud/aws"
	"lsync/types"
	"os"
	"path/filepath"

	"lsync/cloud"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	s3Cmd = &cobra.Command{
		Use:   "s3",
		Short: "upload file or a directory to AWS S3",
		Long:  `upload file or a directory to AWS S3`,
		RunE: func(cmd *cobra.Command, args []string) error {

			fileName, err := cmd.Flags().GetString("file")
			if err != nil {
				return err
			}

			dirName, err := cmd.Flags().GetString("dir")
			if err != nil {
				return err
			}

			svc, err := aws.New("ap-south-1", "cloudintern1")
			if err != nil {
				fmt.Printf("ERROR_INIT_AWS: %s", err)
				os.Exit(2)
			}

			if dirName != "" {
				return syncDir(svc, dirName)
			}

			if fileName != "" {
				return uploadFile(svc, fileName)
			}

			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(s3Cmd)

	s3Cmd.Flags().StringP("file", "f", "", "--file=filename")
	s3Cmd.Flags().StringP("dir", "d", "", "--dir=directory path")
}

func uploadFile(svc cloud.AWS, fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = CheckHash(fileName)
	if err != nil {
		fmt.Printf("ERROR_CHECK_HASH \n\n\n%s", err)
		return err
	}

	if err = svc.UploadFileToS3(f); err != nil {
		color.Red("ERROR_FROM_S3: %s", err)
		return err
	}

	return err
}

func syncDir(svc cloud.AWS, dirName string) error {

	walker := fileWalker(dirName)

	for path := range walker {
		f, err := os.Open(path)
		if err != nil {
			continue
		}

		if err = svc.UploadFileToS3(f); err != nil {
			color.Red("ERROR_S3_UPLOAD: %s", err)
			continue
		}

		color.Green("file  %s  uploaded successfully", path)
		f.Close()
	}

	return nil
}

func fileWalker(dir string) types.FileWalk {
	walker := make(types.FileWalk)

	go func() {
		if err := filepath.Walk(dir, walker.Walk); err != nil {
			color.Red("error while reading files: %s", err)
			os.Exit(5)
		}
		close(walker)
	}()

	return walker
}
