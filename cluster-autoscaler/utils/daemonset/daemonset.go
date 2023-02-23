/*
Copyright 2017 The Kubernetes Authors.

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

package daemonset

import (
	"fmt"
	"math/rand"

	"k8s.io/autoscaler/cluster-autoscaler/simulator/clustersnapshot"
	"k8s.io/autoscaler/cluster-autoscaler/simulator/predicatechecker"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/kubernetes/pkg/controller/daemon"
	schedulerframework "k8s.io/kubernetes/pkg/scheduler/framework"
)

const (
	// EnableDsEvictionKey is the name of annotation controlling whether a
	// certain DaemonSet pod should be evicted.
	EnableDsEvictionKey = "cluster-autoscaler.kubernetes.io/enable-ds-eviction"
)

// GetDaemonSetPodsForNode returns daemonset nodes for the given pod.
func GetDaemonSetPodsForNode(nodeInfo *schedulerframework.NodeInfo, daemonsets []*appsv1.DaemonSet, predicateChecker predicatechecker.PredicateChecker) ([]*apiv1.Pod, error) {
	result := make([]*apiv1.Pod, 0)

	// here we can use empty snapshot
	clusterSnapshot := clustersnapshot.NewBasicClusterSnapshot()

	// add a node with pods - node info is created by cloud provider,
	// we don't know whether it'll have pods or not.
	var pods []*apiv1.Pod
	for _, podInfo := range nodeInfo.Pods {
		pods = append(pods, podInfo.Pod)
	}
	if err := clusterSnapshot.AddNodeWithPods(nodeInfo.Node(), pods); err != nil {
		return nil, err
	}

	for _, ds := range daemonsets {
		shouldRun, _ := daemon.NodeShouldRunDaemonPod(nodeInfo.Node(), ds)
		if shouldRun {
			pod := daemon.NewPod(ds, nodeInfo.Node().Name)
			pod.Name = fmt.Sprintf("%s-pod-%d", ds.Name, rand.Int63())
			result = append(result, pod)
		}
	}
	return result, nil
}

// PodsToEvict returns a list of DaemonSet pods that should be evicted during scale down.
func PodsToEvict(pods []*apiv1.Pod, evictByDefault bool) (evictable []*apiv1.Pod) {
	for _, pod := range pods {
		if a, ok := pod.Annotations[EnableDsEvictionKey]; ok {
			if a == "true" {
				evictable = append(evictable, pod)
			}
		} else if evictByDefault {
			evictable = append(evictable, pod)
		}
	}
	return
}
