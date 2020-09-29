/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"beep"
	"bytes"
	"errors"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// tryCmd represents the try command
var tryCmd = &cobra.Command{
	Use:   "try",
	Short: "A brief description of your command",
	Long: `For example:
		beep try -f spec.yml
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if cfgFile == "" {
			panic(errors.New("Beep config not specified"))
		}
		viper.SetConfigType("yaml")
		b, err := ioutil.ReadFile(cfgFile)
		if err != nil {
			log.Println(err)
		}
		viper.ReadConfig(bytes.NewBuffer(b))

		var p beep.Policy

		err = viper.Unmarshal(&p)
		if err != nil {
			log.Printf("unable to decode into struct, %v", err)
		}
		p.Rate()

	},
}

func init() {
	rootCmd.AddCommand(tryCmd)
	tryCmd.PersistentFlags().StringVarP(&cfgFile, "file", "f", "", "beep configration specified here")
}
