package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tangx/gitlab-release-helper/cmd/releaser/internal/releaser"
)

var cmdUpload = &cobra.Command{
	Use:  "upload",
	Long: `upload file to s3`,
	Run: func(cmd *cobra.Command, args []string) {
		releaser.CreateRelease(args...)
	},
}
