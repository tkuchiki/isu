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
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tkuchiki/isu/db"
	"github.com/tkuchiki/isu/printer"
)

func NewTableRowsCmd() *cobra.Command {
	// tableRowsCmd represents the tableRows command
	var tableRowsCmd = &cobra.Command{
		Use:   "table_rows",
		Short: "Outputs the number of records in all tables.",
		Long: `Outputs the number of records in all tables.

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

			dbname, err := cmd.Flags().GetString("dbname")
			if err != nil {
				return err
			}

			dbhost, err := cmd.Flags().GetString("dbhost")
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

			sort, err := cmd.Flags().GetString("sort")
			if err != nil {
				return err
			}
			if sort != "table_name" && sort != "table_rows" {
				return fmt.Errorf(`Invalid sort option: %s`, sort)
			}

			reverse, err := cmd.Flags().GetBool("reverse")
			if err != nil {
				return err
			}

			dbcli, err := db.New(dbuser, dbpass, dbhost, dbname, dbsock, dbport)
			data, err := dbcli.TableRows(dbname, sort, reverse)
			if err != nil {
				return err
			}

			p := printer.New(os.Stdout)
			p.PrintTable([]string{"table_name", "table_rows"}, data)

			return nil
		},
	}

	tableRowsCmd.Flags().StringP("dbuser", "", "root", "Database user")
	tableRowsCmd.Flags().StringP("dbpass", "", "", "Database password")
	tableRowsCmd.Flags().StringP("dbhost", "", "localhost", "Database host")
	tableRowsCmd.Flags().StringP("dbname", "", "", "Database name")
	tableRowsCmd.Flags().StringP("dbsock", "", "", "Database socket")
	tableRowsCmd.Flags().IntP("dbport", "", 3306, "Database port")
	tableRowsCmd.Flags().StringP("sort", "", "table_name", "Output the results in sorted order(table_name or table_rows)")
	tableRowsCmd.Flags().BoolP("reverse", "r", false, "Sort results in reverse order")

	return tableRowsCmd
}
