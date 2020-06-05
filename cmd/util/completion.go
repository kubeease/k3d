/*
Copyright © 2020 The k3d Author(s)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package util

import (
	"context"
	"strings"

	k3dcluster "github.com/rancher/k3d/pkg/cluster"
	"github.com/rancher/k3d/pkg/runtimes"
	k3d "github.com/rancher/k3d/pkg/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ValidArgsAvailableClusters is used for shell completion: proposes the list of existing clusters
func ValidArgsAvailableClusters(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {

	var completions []string
	var clusters []*k3d.Cluster
	clusters, err := k3dcluster.GetClusters(context.Background(), runtimes.SelectedRuntime)
	if err != nil {
		log.Errorln("Failed to get list of clusters for shell completion")
		return nil, cobra.ShellCompDirectiveError
	}

clusterLoop:
	for _, cluster := range clusters {
		for _, arg := range args {
			if arg == cluster.Name { // only clusters, that are not in the args yet
				continue clusterLoop
			}
		}
		if strings.HasPrefix(cluster.Name, toComplete) {
			completions = append(completions, cluster.Name)
		}
	}
	return completions, cobra.ShellCompDirectiveDefault
}
