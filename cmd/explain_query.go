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
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/tkuchiki/isu/exec"

	"github.com/spf13/cobra"
)

func NewExplainQueryCmd() *cobra.Command {
	// explainQueryCmd represents the explain_query command
	var explainQueryCmd = &cobra.Command{
		Use:   "explain_query",
		Short: "Execute the EXPLAIN and EXPLAIN ANALYZE.",
		Long: `Execute the EXPLAIN and EXPLAIN ANALYZE.

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

			query, err := cmd.Flags().GetString("query")
			if err != nil {
				return err
			}
			if query == "" {
				b, err := ioutil.ReadAll(os.Stdin)
				if err != nil {
					return err
				}
				query = string(b)
			}

			query = strings.ReplaceAll(strings.TrimRight(query, ";"), `""`, `\"`)

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

			explain, err := exec.CommandOutput(ctx, fmt.Sprintf(command+" -e \"EXPLAIN %s\\G\"", query))
			if err != nil {
				return err
			}

			analyze, err := exec.CommandOutput(ctx, fmt.Sprintf(command+" -e \"EXPLAIN ANALYZE %s\"", query))
			if err != nil {
				return err
			}

			fmt.Println(query)
			fmt.Println(explain)
			fmt.Println(analyze)

			return err
		},
	}

	explainQueryCmd.Flags().StringP("dbuser", "", "root", "Database user")
	explainQueryCmd.Flags().StringP("dbpass", "", "", "Database password")
	explainQueryCmd.Flags().StringP("dbhost", "", "127.0.0.1", "Database host")
	explainQueryCmd.Flags().StringP("dbname", "", "", "Database name")
	explainQueryCmd.Flags().StringP("dbsock", "", "", "Database socket")
	explainQueryCmd.Flags().IntP("dbport", "", 3306, "Database port")
	explainQueryCmd.Flags().StringP("query", "", "", "SQL (Read from stdin when omitted)")
	explainQueryCmd.Flags().DurationP("timeout", "", time.Minute*10, "Timeout")

	return explainQueryCmd
}
