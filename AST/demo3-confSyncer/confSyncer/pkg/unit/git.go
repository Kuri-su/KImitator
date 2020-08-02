/*
 * 	ConfSyncer - a little sync config files tool in the Linux.
 *     Copyright (C) 2020  Amatist_kurisu<misaki.zhcy@gmail.com>
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package unit

import (
	"fmt"

	"github.com/fatih/color"
)

func GitPull(dir string) error {
	output, err := RunCommandInShellGetOutput("cd " + dir + " && git pull origin master")
	if err != nil {
		color.White(output)
	}
	return err
}

func GitCommitAndPush(dir string) error {
	output, err := RunCommandInShellGetOutput(fmt.Sprintf(`cd %s && git add -A && git commit -a -m "sync push config" && git push origin master`, dir))
	if err != nil {
		color.White(output)
	}
	return err
}

func GitClone(gitRepoPath string, dir string) error {
	output, err := RunCommandInShellGetOutput("git clone " + gitRepoPath + " " + dir)
	if err != nil {
		color.White(output)
	}
	return err
}
