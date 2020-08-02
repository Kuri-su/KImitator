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
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	jsoniter "github.com/json-iterator/go"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Kuri-su/confSyncer/pkg/unit"
)

const (
	configName = ".confsyncer"
)

var (
	ConfigExists = true
)

// LoadConfig reads in config file and ENV variables if set.
func LoadConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".confSyncer" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(configName)
	}

	viper.AutomaticEnv() // read in environment variables that match
	viper.SetConfigType("yaml")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		ConfigExists = true
		color.White("Using config file: %s", viper.ConfigFileUsed())
	} else {
		ConfigExists = false
		return
	}

	// found config file failed
	configCheck(true)
}

func initConfig(force bool) error {
	err := createConfigFile(force)
	if err != nil {
		return err
	}
	return nil
}

func configCheck(output bool) error {
	if viper.Get("maps") == nil || len(viper.Get("maps").([]interface{})) == 0 {
		msg := "Warning: 'maps' fields is empty in configFile "
		if output {
			color.Red(msg)
		}
		return errors.New(msg)
	}
	return nil
}

func createConfigFile(force bool) error {
	if !unit.IsFile(cfgFile) || force {

		// make dir
		err := unit.MakeDirWithFilePath(cfgFile)
		if err != nil {
			log.Fatalln(err)
		}

		newFile, err := os.Create(cfgFile)
		if err != nil {
			log.Fatalln(err)
		}
		// Write config
		_, err = newFile.Write([]byte(DefaultConfigContext))
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("\n\x1b[32mSuccess!\x1b[0m \n\n"+
			"Create config file in \x1b[32m %s \n"+
			"You should modify it before use. \x1b[0m \n\n",
			cfgFile)
	}

	return nil
}

// ShowConfig
func ShowConfig(*cobra.Command, []string) {
	settingMap := viper.AllSettings()
	settingStr, err := jsoniter.MarshalIndent(settingMap, "", "    ")
	if err != nil {
		fmt.Printf("json marshal failed in ShowConfig! err: %s \n", err.Error())
	}
	color.Green("\nThis is your config: \n%s \n", settingStr)
}

// get files map
func GetFilesMap() ([]Path, error) {
	var maps []Path

	marshal, err := jsoniter.MarshalToString(viper.Get("maps"))
	if err != nil {
		return nil, err
	}
	err = jsoniter.UnmarshalFromString(marshal, &maps)
	if err != nil {
		return nil, err
	}

	return maps, nil
}
