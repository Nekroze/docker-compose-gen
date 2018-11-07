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
	compose "github.com/docker/libcompose/config"
	"github.com/spf13/cobra"
)

// decorateCmd represents the decorate command
var decorateCmd = &cobra.Command{
	Use:   "decorate [options] <SERVICE>...",
	Short: "Decorate the given services",
	Long:  `Decorate the given services.`,
	Run: func(cmd *cobra.Command, args []string) {
		outputConfig(generateDecoratedConfig(args...))
	},
}

func generateDecoratedConfig(services ...string) (out compose.Config) {
	out.Services = compose.RawServiceMap{}
	for _, name := range services {
		out.Services[name] = compose.RawService{}
		if decorateDNS != "" {
			out.Services[name]["dns"] = decorateDNS
		}
	}
	return out
}

var decorateDNS string

func init() {
	rootCmd.AddCommand(decorateCmd)
	decorateCmd.Flags().StringVar(&decorateDNS,
		"dns", "", "dns ip address to add to each service")
}
