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
package confsyncer

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/Kuri-su/confSyncer/pkg/unit"
)

func ConfigPush(cmd *cobra.Command, args []string) {
	f := func(cmd *cobra.Command, args []string) error {
		// check
		if !unit.IsDir(TmpDirPath) {
			err := initTmpDir()
			if err != nil {
				return err
			}
		}

		// TODO package this code section
		// get config
		maps, err := GetFilesMap()

		for _, pathStruct := range maps {
			copySrc := pathStruct.Local
			copyDist := TmpDirPath + pathStruct.GitRepoPath
			err := unit.Copy(copySrc, copyDist)
			if err != nil {
				color.Red(fmt.Sprintf("copy '%s' to '%s' failed! \nErr: %s", copySrc, copyDist, err.Error()))
			} else {
				color.Green(fmt.Sprintf("copy '%s' to '%s' success", copySrc, copyDist))
			}
		}

		err = unit.GitCommitAndPush(TmpDirPath)
		if err != nil && err.Error() != "exit status 1" {
			return err
		}

		color.Green("Configs push finish!")
		return nil
	}

	err := f(cmd, args)
	if err != nil {
		color.Red(err.Error())
		return
	}
}
