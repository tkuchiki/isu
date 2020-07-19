/*
Copyright © 2020 tkuchiki

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/tkuchiki/isu/exec"

	"github.com/spf13/cobra"
)

func NewCodeBackupCmd() *cobra.Command {
	// codeBackupCmd represents the code_backup command
	var codeBackupCmd = &cobra.Command{
		Use:   "code_backup",
		Short: "Create a backup of directory",
		Long:  `Create a backup of directory`,
		RunE: func(cmd *cobra.Command, args []string) error {
			destDir, err := cmd.Flags().GetString("dest-dir")
			if err != nil {
				return err
			}

			srcDir, err := cmd.Flags().GetString("src-dir")
			if err != nil {
				return err
			}

			excludesStr, err := cmd.Flags().GetString("excludes")
			if err != nil {
				return err
			}

			abs, err := filepath.Abs(srcDir)
			if err != nil {
				return err
			}

			basedir := strings.TrimRight(filepath.Base(abs), "/")
			filename := filepath.Join(destDir, fmt.Sprintf(`%s.tar.gz`, basedir))
			wdir := filepath.Dir(srcDir)

			excludesSlice := make([]string, 0)
			for _, exc := range strings.Split(excludesStr, ",") {
				excludesSlice = append(excludesSlice, fmt.Sprintf(`--exclude '%s'`, exc))
			}

			command := fmt.Sprintf(`( cd %s; tar %s -czf %s %s )`, wdir, strings.Join(excludesSlice, " "), filename, basedir)

			timeout, err := cmd.Flags().GetDuration("timeout")
			if err != nil {
				return err
			}

			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()
			err = exec.Command(ctx, command)

			return err
		},
	}

	codeBackupCmd.Flags().StringP("excludes", "", ".git", "Do not process files or directories that match the specified pattern (comma separated)")
	codeBackupCmd.Flags().StringP("dest-dir", "d", "./", "Destination directory (required)")
	codeBackupCmd.Flags().StringP("src-dir", "s", "", "Source directory (required)")
	codeBackupCmd.Flags().DurationP("timeout", "", time.Minute*10, "Timeout")

	codeBackupCmd.MarkFlagRequired("dest-dir")
	codeBackupCmd.MarkFlagRequired("src-dir")

	return codeBackupCmd
}
