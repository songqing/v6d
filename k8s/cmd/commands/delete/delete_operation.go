/*
* Copyright 2020-2023 Alibaba Group Holding Limited.

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
package delete

import (
	"github.com/spf13/cobra"

	"k8s.io/apimachinery/pkg/types"

	vineyardV1alpha1 "github.com/v6d-io/v6d/k8s/apis/k8s/v1alpha1"
	"github.com/v6d-io/v6d/k8s/cmd/commands/flags"
	"github.com/v6d-io/v6d/k8s/cmd/commands/util"
	"github.com/v6d-io/v6d/k8s/pkg/log"
)

var deleteOperationExample = util.Examples(`
	# delete the operation named "assembly-test" in the "vineyard-system" namespace
	vineyardctl delete operation --name assembly-test`)

// deleteOperationCmd deletes the specific operation
var deleteOperationCmd = &cobra.Command{
	Use:     "operation",
	Short:   "Delete the operation from kubernetes",
	Example: deleteOperationExample,
	Run: func(cmd *cobra.Command, args []string) {
		util.AssertNoArgs(cmd, args)

		client := util.KubernetesClient()

		operation := &vineyardV1alpha1.Operation{}
		if err := util.Delete(client, types.NamespacedName{
			Name:      flags.OperationName,
			Namespace: flags.GetDefaultVineyardNamespace(),
		}, operation); err != nil {
			log.Fatal(err, "failed to delete operation")
		}

		log.Info("Operation is deleted.")
	},
}

func NewDeleteOperationCmd() *cobra.Command {
	return deleteOperationCmd
}

func init() {
	flags.ApplyOperationName(deleteOperationCmd)
}
