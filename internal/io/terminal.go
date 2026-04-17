// Copyright © 2026 CloudODX Contributors
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

package io

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func GetUsernamePassword() (username string, password string) {
	reader := bufio.NewReader(os.Stdin)
	username = ""
	for len(username) == 0 {
		fmt.Print("Enter username: ")
		username, _ = reader.ReadString('\n')
		username = strings.TrimSpace(username)
	}

	password = ""
	for len(password) == 0 {
		fmt.Print("Enter password: ")
		bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
		password = string(bytePassword)
	}

	return username, password
}
