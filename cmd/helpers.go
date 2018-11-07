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
	"io/ioutil"

	compose "github.com/docker/libcompose/config"
	yaml "gopkg.in/yaml.v2"
)

func outputConfig(config compose.Config) {
	config.Version = composeVersion
	output, err := yaml.Marshal(config)
	if err != nil {
		panic(err)
	}
	if writeToStdout {
		fmt.Print(string(output))
		return
	}
	tmpfile, err := ioutil.TempFile("", "docker-compose.gen.*.yaml")
	if err != nil {
		panic(err)
	}
	if _, e := tmpfile.Write(output); e != nil {
		panic(err)
	}
	if e := tmpfile.Chmod(0664); e != nil {
		panic(err)
	}
	fmt.Print(tmpfile.Name())
	if e := tmpfile.Close(); e != nil {
		panic(err)
	}
}
