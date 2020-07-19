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
	"os"
	"path/filepath"
	"time"

	"github.com/tkuchiki/isu/exec"

	"github.com/spf13/cobra"
)

func NewDbBackupCmd() *cobra.Command {
	// dbBackupCmd represents the db_backup command
	var dbBackupCmd = &cobra.Command{
		Use:   "db_backup",
		Short: "Execute the mysqldump.",
		Long: `Execute the mysqldump.

`,
		RunE: func(cmd *cobra.Command, args []string) error {
			dbuser, err := cmd.Flags().GetString("dbuser")
			if err != nil {
				return err
			}

			dbpass, err := cmd.Flags().GetString("dbpass")
			if err != nil {
				return err
			}
			// TODO: MYSQL_PWD is deprecated as of MySQL 8.0 and will be removed in a future MySQL version
			if os.Getenv("MYSQL_PWD") != "" {
				dbpass = os.Getenv("MYSQL_PWD")
			}
			os.Setenv("MYSQL_PWD", dbpass)

			dbhost, err := cmd.Flags().GetString("dbhost")
			if err != nil {
				return err
			}

			dbname, err := cmd.Flags().GetString("dbname")
			if err != nil {
				return err
			}

			dbsock, err := cmd.Flags().GetString("dbsock")
			if err != nil {
				return err
			}

			dbport, err := cmd.Flags().GetInt("dbport")
			if err != nil {
				return err
			}

			dir, err := cmd.Flags().GetString("output-dir")
			if err != nil {
				return err
			}
			filename := filepath.Join(dir, fmt.Sprintf(`%s.sql`, dbname))

			baseCommand := fmt.Sprintf(`mysqldump --single-transaction -u "%s" --port %d`, dbuser, dbport)
			var command string
			if dbsock == "" {
				command = fmt.Sprintf(baseCommand+` -h "%s" %s`, dbhost, dbname)
			} else {
				command = fmt.Sprintf(baseCommand+` --socket "%s" %s`, dbsock, dbname)
			}

			timeout, err := cmd.Flags().GetDuration("timeout")
			if err != nil {
				return err
			}

			compress, err := cmd.Flags().GetBool("compress")
			if err != nil {
				return err
			}

			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()
			if compress {
				err = exec.CommandWithGzip(ctx, command, filename)
			} else {
				command = fmt.Sprintf(command+` > %s`, filename)
				err = exec.Command(ctx, command)
			}

			return err
		},
	}

	dbBackupCmd.Flags().StringP("dbuser", "", "root", "Database user")
	dbBackupCmd.Flags().StringP("dbpass", "", "", "Database password")
	dbBackupCmd.Flags().StringP("dbhost", "", "127.0.0.1", "Database host")
	dbBackupCmd.Flags().StringP("dbname", "", "", "Database name")
	dbBackupCmd.Flags().StringP("dbsock", "", "", "Database socket")
	dbBackupCmd.Flags().StringP("output-dir", "d", "./", "Output directory")
	dbBackupCmd.Flags().IntP("dbport", "", 3306, "Database port")
	dbBackupCmd.Flags().DurationP("timeout", "", time.Minute*10, "Timeout")
	dbBackupCmd.Flags().BoolP("compress", "c", false, "Compress file with gzip command")

	return dbBackupCmd
}
