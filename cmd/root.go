// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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

package cmd

import (
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"video-intelligence/pkg/version"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func persistentPreRun(cmd *cobra.Command, args []string) {
	// profile init: env args config
	// home dir init:
	// log init: loki
}

var root = &cobra.Command{
	Long:             filepath.Base(os.Args[0]) + " is video intelligence system",
	Version:          version.String(),
	PersistentPreRun: persistentPreRun,
}

// AddCommand AddCommand
func AddCommand(cmds ...*cobra.Command) {
	root.AddCommand(cmds...)
}

// Execute Execute
func Execute() {
	if err := root.Execute(); err != nil {
		log.Fatal().Err(err).Msg("startup err")
	}
}
