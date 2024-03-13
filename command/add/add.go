//
// Copyright 2024-2024 The Haora Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package add

import (
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/command/internal/parsing"
	"github.com/spf13/cobra"
	"strings"
)

var (
	startFlag  string
	tagsFlag   string
	noTagsFlag bool
)

func init() {
	Command.Flags().StringVarP(&startFlag, "start", "s", "", "starting timestamp, like 10:00, of the task")
	Command.Flags().StringVarP(&tagsFlag, "tags", "t", "", "comma separated tags of the task")
	Command.Flags().BoolVar(&noTagsFlag, "no-tags", false, "set if the new task shall have no tags")
}

var Command = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to a day",
	Long: `Adds a new task to a day. 

The default and simplest to use format for the add command is 

$ haora add [time] [single tag] [text...]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		time, args, err := parsing.Time(startFlag, args)
		if err != nil {
			return err
		}
		var tags []string
		if !noTagsFlag {
			tags, args = parseTags(args)
		}
		text := strings.Join(args, " ")
		d := data.State.WorkingDay()
		return d.AddNewTask(time, text, tags)
	},
	PostRun: func(cmd *cobra.Command, args []string) { // reset flag so tests can rerun!
		startFlag = ""
		tagsFlag = ""
		noTagsFlag = false
	},
}

func parseTags(args []string) ([]string, []string) {
	if tagsFlag != "" {
		tags := strings.Split(tagsFlag, ",")
		return tags, args
	}
	if len(args) > 0 {
		tags := strings.Split(args[0], ",")
		return tags, args[1:]
	}
	return []string{}, args
}
