/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, VersionCmd 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	"github.com/Kuri-su/confSyncer/pkg/confsyncer"
)

type DockerCompose struct {
	Version  string               `yaml:"version"`
	Services map[string]Container `yaml:"service"`
}

type Container struct {
	Image   string   `yaml:"image"`
	Restart string   `yaml:"restart"`
	Volumes []string `yaml:"volumes"`
}

const (
	ContainerName                 = "confsyncer"
	GenerateDockerComposeFileName = "docker-compose.yaml"
)

var rootCmd *cobra.Command

func init() {
	initCobra()
}

func main() {
	initCobra()

	// run
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initCobra() {
	rootCmd = &cobra.Command{
		Use:   "confsyncerGen",
		Short: `build docker-compose.yaml from confsyncer conf`,
		Run:   run,
	}
	commands := []*cobra.Command{
		&cobra.Command{
			Use:   "config",
			Short: "show config",
		},
		&cobra.Command{
			Use:   "composeyaml",
			Short: "show composeyaml",
		},
	}

	rootCmd.AddCommand(commands...)
}

func run(cmd *cobra.Command, args []string) {
	dc := new(DockerCompose)
	initDockerComposeStruct(dc)

	maps, err := confsyncer.GetFilesMap()
	if err != nil {
		panic(err)
	}

	var volumes []string
	for _, m := range maps {
		volumes = append(volumes, fmt.Sprintf("%s:%s", m.Local, m.Local))
	}

	dc.Services[ContainerName] = Container{
		Image:   "kurisux/conf-syncer:latest",
		Restart: "always",
		Volumes: volumes,
	}

	marshal, err := yaml.Marshal(dc)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(GenerateDockerComposeFileName, marshal, os.FileMode(0744))
	if err != nil {
		panic(err)
	}
}

func initDockerComposeStruct(dc *DockerCompose) {
	dc.Version = "3"
	dc.Services = make(map[string]Container)
}
