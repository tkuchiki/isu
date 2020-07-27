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
	"time"

	"github.com/tkuchiki/isu/db"
	"github.com/tkuchiki/isu/exec"

	"github.com/spf13/cobra"
)

func NewShowIndexesCmd() *cobra.Command {
	// showIndexesCmd represents the show_indexes command
	var showIndexesCmd = &cobra.Command{
		Use:   "show_indexes",
		Short: "Show indexes for all tables",
		Long: `Show indexes for all tables

Only support MySQL`,
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

			reverse, err := cmd.Flags().GetBool("reverse")
			if err != nil {
				return err
			}

			dbcli, err := db.New(dbuser, dbpass, dbhost, dbname, dbsock, dbport)
			if err != nil {
				return err
			}
			defer dbcli.Close()

			tables, err := dbcli.GetTables(dbname, reverse)
			if err != nil {
				return err
			}

			timeout, err := cmd.Flags().GetDuration("timeout")
			if err != nil {
				return err
			}

			baseCommand := fmt.Sprintf(`mysql -u "%s" --port %d --table`, dbuser, dbport)
			var command string
			if dbsock == "" {
				command = fmt.Sprintf(baseCommand+` -h "%s" %s`, dbhost, dbname)
			} else {
				command = fmt.Sprintf(baseCommand+` --socket "%s" %s`, dbsock, dbname)
			}

			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			for _, table := range tables {
				out, err := exec.CommandOutput(ctx, fmt.Sprintf(command+" -e 'SHOW INDEXES FROM `%s`'", table))
				if err != nil {
					return err
				}

				if out == "" {
					continue
				}

				fmt.Println("##", table)
				fmt.Println(out)
			}

			return err
		},
	}

	showIndexesCmd.Flags().StringP("dbuser", "", "root", "Database user")
	showIndexesCmd.Flags().StringP("dbpass", "", "", "Database password")
	showIndexesCmd.Flags().StringP("dbhost", "", "127.0.0.1", "Database host")
	showIndexesCmd.Flags().StringP("dbname", "", "", "Database name")
	showIndexesCmd.Flags().StringP("dbsock", "", "", "Database socket")
	showIndexesCmd.Flags().IntP("dbport", "", 3306, "Database port")
	showIndexesCmd.Flags().BoolP("reverse", "r", false, "Sort results in reverse order")
	showIndexesCmd.Flags().DurationP("timeout", "", time.Minute*10, "Timeout")

	return showIndexesCmd
}
