
package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"lsync/cloud"
	"lsync/cloud/aws"
	"os"
)

var (
 deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a file or directory from AWS S3",
	Long: `The file or directory that are present inside the AWS S3 can be deleted with this command and sub
command`,
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

		if fileName != "" {
			return deleteFile(svc, fileName)
		}
		if dirName != "" {
			return deleteFile(svc, dirName)
		}
		return nil
	},
}
)

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringP( "file", "f", "", "--file=filename")
	deleteCmd.Flags().StringP("dir", "d", "", "--dir=directory path")

}

func deleteFile(svc cloud.AWS, fileName string) error {
	err := svc.DeleteFileFromS3(fileName)
	if err != nil {
		fmt.Printf("Error while file Delete: %s", err)
		return err
	}

	if err != nil {
		color.Red("ERROR_FROM_S3: %s", err)
		return err
	}

	return nil
}
