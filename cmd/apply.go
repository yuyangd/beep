/*
Copyright © 2020 yuyangd

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
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"nrops"
	"os"

	"nrops/kms"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply the specification to NewRelic",
	Long: `For example:
		nrops apply -f spec.yml
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if cfgFile == "" {
			panic(errors.New("NewRelic config not specified"))
		}

		viper.SetConfigType("yaml")
		b, err := ioutil.ReadFile(cfgFile)
		if err != nil {
			log.Println(err)
		}
		viper.ReadConfig(bytes.NewBuffer(b))

		nr := viper.GetStringMap("new_relic")

		apiKey, err := (&kms.Handler{
			Service:   kms.Client(os.Getenv("AWS_DEFAULT_REGION")),
			Context:   nil,
			CipherKey: fmt.Sprintf("%v", nr["x-api-key"]),
		}).Decrypt()
		if err != nil {
			panic(err)
		}

		session = nrops.NewClient(apiKey)
		if nr["deployment_marker"].(string) == "enabled" {
			log.Println("deployment marker enabled")
			err := session.SetDMarker(fmt.Sprintf("%v", nr["application_id"]), &nrops.DMarker{
				Deployment: nrops.Deploy{
					User:        fmt.Sprintf("%v", viper.Get("BUILD_CREATOR_EMAIL")),
					Description: fmt.Sprintf("%v", viper.Get("BUILD_URL")),
					Changelog:   fmt.Sprintf("%v", viper.Get("MESSAGE")),
					Revision:    fmt.Sprintf("%v", viper.Get("COMMIT")),
				},
			})
			if err != nil {
				log.Panicln(err)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(applyCmd)

	applyCmd.PersistentFlags().StringVarP(&cfgFile, "file", "f", "", "NewRelic config required for application")
}
