/*
 * Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License").
 * You may not use this file except in compliance with the License.
 * A copy of the License is located at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 * or in the "license" file accompanying this file. This file is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package commands

import (
	handler "es-cli/odfe-cli/handler/ad"
	"fmt"

	"github.com/spf13/cobra"
)

const (
	deleteDetectorsCommandName    = "delete"
	deleteDetectorIDFlagName      = "id"
	detectorForceDeletionFlagName = "force"
)

//deleteDetectorsCmd deletes detectors based on id, name or name regex pattern.
//default input is name pattern, one can change this format to be id by passing --id flag
var deleteDetectorsCmd = &cobra.Command{
	Use:   deleteDetectorsCommandName + " detector_name ..." + " [flags] ",
	Short: "Delete detectors based on list of id, name or name regex pattern ",
	Long: fmt.Sprintf("Description:\n  " +
		`Delete detectors based on list of user input. Use "" to make sure the name does not match with pwd lists'.
  The default input is detector name, use --id flag if input is detector id instead of name`),
	Run: func(cmd *cobra.Command, args []string) {
		//If no args, display usage
		if len(args) < 1 {
			fmt.Println(cmd.Usage())
			return
		}
		force, _ := cmd.Flags().GetBool(detectorForceDeletionFlagName)
		detectorID, _ := cmd.Flags().GetBool(deleteDetectorIDFlagName)
		action := handler.DeleteAnomalyDetectorByNamePattern
		if detectorID {
			action = handler.DeleteAnomalyDetectorByID
		}
		err := deleteDetectors(args, force, action)
		DisplayError(err, deleteDetectorsCommandName)
	},
}

func init() {
	GetADCommand().AddCommand(deleteDetectorsCmd)
	deleteDetectorsCmd.Flags().BoolP(detectorForceDeletionFlagName, "f", false, "deletes even if detector is running")
	deleteDetectorsCmd.Flags().BoolP(deleteDetectorIDFlagName, "", false, "input is detector's id")
}

func deleteDetectors(detectors []string, force bool, f func(*handler.Handler, string, bool) error) error {
	commandHandler, err := GetADHandler()
	if err != nil {
		return err
	}
	for _, detector := range detectors {
		err = f(commandHandler, detector, force)
		if err != nil {
			return err
		}
	}
	return nil
}
