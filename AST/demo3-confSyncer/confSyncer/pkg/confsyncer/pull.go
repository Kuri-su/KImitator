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
	"time"

	"github.com/fatih/color"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Kuri-su/confSyncer/pkg/unit"
)

func ConfigPull(cmd *cobra.Command, args []string) {
	f := func(cmd *cobra.Command, args []string) error {
		if !unit.IsDir(TmpDirPath) {
			err := initTmpDir()
			if err != nil {
				return err
			}
		}

		err := unit.GitPull(TmpDirPath)
		if err != nil && err.Error() != "exit status 1" {
			return err
		}

		c := viper.Get("maps")
		if c == nil {
			return nil
		}

		str, err := jsoniter.MarshalToString(c)
		if err != nil {
			return err
		}

		var maps []Path
		err = jsoniter.UnmarshalFromString(str, &maps)
		if err != nil {
			return err
		}

		for _, copyMap := range maps {
			unit.MakeDirWithFilePath(copyMap.Local)
			copySrc := TmpDirPath + copyMap.GitRepoPath
			copyDist := copyMap.Local
			err = unit.Copy(TmpDirPath+copyMap.GitRepoPath, copyMap.Local)
			if err != nil {
				color.Red(fmt.Sprintf("copy '%s' to '%s' failed!\nErr: %s", copySrc, copyDist, err.Error()))
			} else {
				color.Green(fmt.Sprintf("copy '%s' to '%s' success", copySrc, copyDist))
			}
		}

		fmt.Println("Configs pull finish!")
		return nil
	}

	err := f(cmd, args)
	if err != nil {
		color.Red(err.Error())
	}
}

func DaemonPull(cmd *cobra.Command, args []string) {
	ticker := time.NewTicker(time.Duration(viper.GetInt("gitPullTimeInternal")) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			ConfigPull(cmd, args)
		}
	}
}
