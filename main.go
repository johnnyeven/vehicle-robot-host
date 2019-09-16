// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bufio"
	"fmt"
	"github.com/johnnyeven/libtools/servicex"
	"github.com/johnnyeven/vehicle-robot-host/cmd"
	"github.com/johnnyeven/vehicle-robot-host/global"
	_ "github.com/johnnyeven/vehicle-robot-host/global"
	"github.com/johnnyeven/vehicle-robot-host/routes"
	"os"
	"runtime"
	"strings"
)

func main() {
	runtime.LockOSThread()
	servicex.Execute()
	routes.InitRouters()

	defer func() {
		global.Config.Window.Close()
	}()

	var command string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		command = scanner.Text()
		if command == "" {
			continue
		} else if command == "quit" || command == "exit" {
			fmt.Println("Goodbye")
			os.Exit(0)
		}
		commands := []string{os.Args[0]}
		commands = append(commands, strings.Split(command, " ")...)
		os.Args = commands

		go cmd.Execute()
	}
}
