//
// Copyright © 2017-2019 Solus Project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package cmd

import (
	"fmt"
	"github.com/getsolus/ferryd/client"
	"github.com/spf13/cobra"
	"os"
)

var createRepoCmd = &cobra.Command{
	Use:   "create-repo [repoName]",
	Short: "create a new repository",
	Long:  "Create a new repository in the ferryd instance, if it doesn't exist",
	Run:   createRepo,
}

func init() {
	RootCmd.AddCommand(createRepoCmd)
}

func createRepo(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "usage: create-repo [repoName]\n")
		return
	}

	client := client.NewClient(socketPath)
	defer client.Close()

	if err := client.CreateRepo(args[0]); err != nil {
		fmt.Fprintf(os.Stderr, "Error while creating repo: %v\n", err)
		return
	}
}
