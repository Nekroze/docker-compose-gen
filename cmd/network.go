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

	compose "github.com/docker/libcompose/config"
	composeyml "github.com/docker/libcompose/yaml"
	"github.com/spf13/cobra"

	"gopkg.in/yaml.v2"
)

var networkExternalName string
var networkName string

// networkCmd represents the network command
var networkCmd = &cobra.Command{
	Use:   "network",
	Short: "Generate a network block",
	Long: `Generate a network block.
Can be used to generate networks, helpful for dynamically replacing the default network.`,
	Run: func(cmd *cobra.Command, args []string) {
		output, err := yaml.Marshal(generateConfig())
		if err != nil {
			panic(err)
		}
		fmt.Print(string(output))
	},
}

func generateConfig() (out compose.Config) {
	out.Version = composeVersion
	if (networkExternalName == "" && networkName != "default") || networkExternalName != "" {
		out.Networks = map[string]interface{}{
			networkName: generateNetworkConfig(),
		}
	}
	return out
}

func generateNetworkConfig() (out compose.NetworkConfig) {
	if networkExternalName != "" {
		out.External = composeyml.External{
			External: true,
			Name:     networkExternalName,
		}
	}
	return out
}

func init() {
	rootCmd.AddCommand(networkCmd)
	networkCmd.Flags().StringVar(&networkExternalName,
		"external", "", "Name of the externally defined network, if not given an internal network will be used")
	networkCmd.Flags().StringVar(&networkName,
		"name", "default", "Name of the network to generate, defaults to 'default'")
}
