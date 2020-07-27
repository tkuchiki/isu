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
	"github.com/spf13/cobra"
	"github.com/tkuchiki/isu/db"
)

func NewSlowlogCmd() *cobra.Command {
	// slowlogCmd represents the slowlog command
	var slowlogCmd = &cobra.Command{
		Use:   "slowlog",
		Short: "Enable/Disable slowlog",
		Long: `Enable/Disable slowlog

Only support MySQL`,
	}

	slowlogCmd.PersistentFlags().StringP("dbuser", "", "root", "Database user")
	slowlogCmd.PersistentFlags().StringP("dbpass", "", "", "Database password")
	slowlogCmd.PersistentFlags().StringP("dbhost", "", "localhost", "Database host")
	slowlogCmd.PersistentFlags().StringP("dbsock", "", "", "Database socket")
	slowlogCmd.PersistentFlags().IntP("dbport", "", 3306, "Database port")
	slowlogCmd.PersistentFlags().BoolP("persist", "", false, "Use `SET PERSIST`")

	// slowlogOnCmd represents the slowlog on command
	var slowlogOnCmd = &cobra.Command{
		Use:   "on",
		Short: "Enable slowlog.",
		Long: `Enable slowlog.

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

			longQueryTime, err := cmd.Flags().GetFloat64("long-query-time")
			if err != nil {
				return err
			}

			persist, err := cmd.Flags().GetBool("persist")
			if err != nil {
				return err
			}

			dbcli, err := db.New(dbuser, dbpass, dbhost, "", dbsock, dbport)
			if err != nil {
				return err
			}
			defer dbcli.Close()

			datadir, err := dbcli.GetGlobalVariable("datadir")
			if err != nil {
				return err
			}

			err = dbcli.SlowlogOn(datadir, longQueryTime, persist)
			if err != nil {
				return err
			}

			return nil
		},
	}

	slowlogOnCmd.Flags().Float64P("long-query-time", "", 0.1, "Set long_query_time")

	slowlogCmd.AddCommand(slowlogOnCmd)

	// slowlogOffCmd represents the slowlog off command
	var slowlogOffCmd = &cobra.Command{
		Use:   "off",
		Short: "Disable slowlog.",
		Long: `Disable slowlog.

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

			persist, err := cmd.Flags().GetBool("persist")
			if err != nil {
				return err
			}

			dbcli, err := db.New(dbuser, dbpass, dbhost, "", dbsock, dbport)
			if err != nil {
				return err
			}
			defer dbcli.Close()

			err = dbcli.SlowlogOff(persist)
			if err != nil {
				return err
			}

			return nil
		},
	}

	slowlogCmd.AddCommand(slowlogOffCmd)

	return slowlogCmd
}
