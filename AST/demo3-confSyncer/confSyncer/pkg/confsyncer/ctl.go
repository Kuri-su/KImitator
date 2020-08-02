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
	"os"
	"os/user"
	"strings"
	"time"

	"github.com/spf13/viper"

	"github.com/Kuri-su/confSyncer/pkg/unit"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	// Root
	dirPath = "$HOME/.confsyncer"
	cfgFile = "$HOME/.confsyncer/config.yaml"
	// TmpDir(TmpDirPath) is a tmp git clone dir
	TmpDirPath = "/tmp/confsyncer-" + fmt.Sprint(time.Now().Format("20060102"))
	// version
	version = "0.0.3"
	// use it when need a default config
	DefaultConfigContext = `---
gitRepo: ""
gitPullTimeInternal: 600 # second
maps:
  - gitRepoPath: /.confsyncer/config.yaml
    local: ~/.confsyncer/config.yaml
`

	// ========================================================================
	// ========================================================================
	// ========================================================================

	// CMD
	rootCmd = &cobra.Command{
		Use:   "confsyncer",
		Short: "confsyncer",
		Long:  `confsyncer`,
	}
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "initialization config",
		RunE:  initConfigCmd,
	}
	configCmd = &cobra.Command{
		Use:     "config",
		Short:   "show config",
		PreRunE: configFileExistsCheckCmd,
		Run:     ShowConfig,
	}
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "show version",
		Run:   VersionCmd,
	}
	pushCmd = &cobra.Command{
		Use:     "push",
		Short:   "push",
		PreRunE: configFileExistsCheckCmd,
		Run:     ConfigPush,
	}
	pullCmd = &cobra.Command{
		Use:     "pull",
		Short:   "pull",
		PreRunE: configFileExistsCheckCmd,
		Run:     ConfigPull,
	}
	deamonPullCmd = &cobra.Command{
		Use:     "daemon",
		Short:   "daemon",
		PreRunE: configFileExistsCheckCmd,
		Run:     DaemonPull,
	}

	// flags
	initForce bool
)

func init() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	cfgFile = strings.Replace(cfgFile, "$HOME", u.HomeDir, 1)
	dirPath = strings.Replace(dirPath, "$HOME", u.HomeDir, 1)

	cobra.OnInitialize(LoadConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", cfgFile, "config file")
	initCmd.Flags().BoolVarP(&initForce, "force", "f", initForce, "force init")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	commands := []*cobra.Command{
		initCmd,
		configCmd,
		versionCmd,
		pushCmd,
		pullCmd,
		deamonPullCmd,
	}

	// register commands
	rootCmd.AddCommand(commands...)

	// add postRun in all command
	for _, command := range commands {
		command.PostRun = removeTmpDirCmd
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// initConfigCmd
func initConfigCmd(cmd *cobra.Command, args []string) error {
	err := initConfig(initForce)
	if err != nil {
		return err
	}
	LoadConfig()
	ShowConfig(cmd, args)
	return nil
}

// configFileExistsCheckCmd
func configFileExistsCheckCmd(cmd *cobra.Command, args []string) error {
	if !ConfigExists {
		msg := color.RedString(fmt.Sprintf(`config file not found in "%s", please run "confsyncer init" first! `, cfgFile))
		return errors.New(msg)
	}
	return nil
}

// VersionCmd
func VersionCmd(cmd *cobra.Command, args []string) {
	color.Set(color.Bold)
	color.HiWhite("confSyncer version: %s", version)
}

func initTmpDir() error {
	err := unit.GitClone(viper.GetString("gitRepo"), TmpDirPath)
	if err != nil {
		// TODO Replace to log
		color.Red(fmt.Sprintf("git clone error! err: %s", err.Error()))
	}
	return err
}

func removeTmpDirCmd(cmd *cobra.Command, args []string) {
	unit.RemoveFiles(TmpDirPath)
}
