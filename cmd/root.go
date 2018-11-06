// Copyright © 2018 Taylor Lawson <nekroze.lawson@gmail.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "docker-compose-gen",
	Short: "Generate docker compose config files",
	Long: `Generate docker compose config files.

For those times when it is handy to generate a more dynamic section of config.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var composeVersion string
var writeToStdout bool

func init() {
	rootCmd.PersistentFlags().StringVar(&composeVersion,
		"compose-version", "2.1", "Docker compose config version to generate (defaults to '2.1')")
	rootCmd.PersistentFlags().BoolVar(&writeToStdout,
		"stdout", false, "Write the yaml config directly to stdout instead of writing a temporary file name")
}
