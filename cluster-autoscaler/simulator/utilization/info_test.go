/*
Copyright 2016 The Kubernetes Authors.

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

package utilization

import (
	"testing"
	"time"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	. "k8s.io/autoscaler/cluster-autoscaler/utils/test"
	"k8s.io/kubernetes/pkg/kubelet/types"
	schedulerframework "k8s.io/kubernetes/pkg/scheduler/framework"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	testTime := time.Date(2020, time.December, 18, 17, 0, 0, 0, time.UTC)
	gpuLabel := GetGPULabel()
	pod := BuildTestPod("p1", 100, 200000)
	pod2 := BuildTestPod("p2", -1, -1)

	nodeInfo := schedulerframework.NewNodeInfo(pod, pod, pod2)
	node := BuildTestNode("node1", 2000, 2000000)
	SetNodeReadyState(node, true, time.Time{})

	utilInfo, err := Calculate(node, nodeInfo, false, false, gpuLabel, testTime)
	assert.NoError(t, err)
	assert.InEpsilon(t, 2.0/10, utilInfo.Utilization, 0.01)

	node2 := BuildTestNode("node1", 2000, -1)

	_, err = Calculate(node2, nodeInfo, false, false, gpuLabel, testTime)
	assert.Error(t, err)

	daemonSetPod3 := BuildTestPod("p3", 100, 200000)
	daemonSetPod3.OwnerReferences = GenerateOwnerReferences("ds", "DaemonSet", "apps/v1", "")

	daemonSetPod4 := BuildTestPod("p4", 100, 200000)
	daemonSetPod4.OwnerReferences = GenerateOwnerReferences("ds", "CustomDaemonSet", "crd/v1", "")
	daemonSetPod4.Annotations = map[string]string{"cluster-autoscaler.kubernetes.io/daemonset-pod": "true"}

	nodeInfo = schedulerframework.NewNodeInfo(pod, pod, pod2, daemonSetPod3, daemonSetPod4)
	utilInfo, err = Calculate(node, nodeInfo, true, false, gpuLabel, testTime)
	assert.NoError(t, err)
	assert.InEpsilon(t, 2.5/10, utilInfo.Utilization, 0.01)

	nodeInfo = schedulerframework.NewNodeInfo(pod, pod2, daemonSetPod3)
	utilInfo, err = Calculate(node, nodeInfo, false, false, gpuLabel, testTime)
	assert.NoError(t, err)
	assert.InEpsilon(t, 2.0/10, utilInfo.Utilization, 0.01)

	terminatedPod := BuildTestPod("podTerminated", 100, 200000)
	terminatedPod.DeletionTimestamp = &metav1.Time{Time: testTime.Add(-10 * time.Minute)}
	nodeInfo = schedulerframework.NewNodeInfo(pod, pod, pod2, terminatedPod)
	utilInfo, err = Calculate(node, nodeInfo, false, false, gpuLabel, testTime)
	assert.NoError(t, err)
	assert.InEpsilon(t, 2.0/10, utilInfo.Utilization, 0.01)

	mirrorPod := BuildTestPod("p4", 100, 200000)
	mirrorPod.Annotations = map[string]string{
		types.ConfigMirrorAnnotationKey: "",
	}

	nodeInfo = schedulerframework.NewNodeInfo(pod, pod, pod2, mirrorPod)
	utilInfo, err = Calculate(node, nodeInfo, false, true, gpuLabel, testTime)
	assert.NoError(t, err)
	assert.InEpsilon(t, 2.0/9.0, utilInfo.Utilization, 0.01)

	nodeInfo = schedulerframework.NewNodeInfo(pod, pod2, mirrorPod)
	utilInfo, err = Calculate(node, nodeInfo, false, false, gpuLabel, testTime)
	assert.NoError(t, err)
	assert.InEpsilon(t, 2.0/10, utilInfo.Utilization, 0.01)

	nodeInfo = schedulerframework.NewNodeInfo(pod, mirrorPod, daemonSetPod3)
	utilInfo, err = Calculate(node, nodeInfo, true, true, gpuLabel, testTime)
	assert.NoError(t, err)
	assert.InEpsilon(t, 1.0/8.0, utilInfo.Utilization, 0.01)

	gpuNode := BuildTestNode("gpu_node", 2000, 2000000)
	AddGpusToNode(gpuNode, 1)
	gpuPod := BuildTestPod("gpu_pod", 100, 200000)
	RequestGpuForPod(gpuPod, 1)
	TolerateGpuForPod(gpuPod)
	nodeInfo = schedulerframework.NewNodeInfo(pod, pod, gpuPod)
	utilInfo, err = Calculate(gpuNode, nodeInfo, false, false, gpuLabel, testTime)
	assert.NoError(t, err)
	assert.InEpsilon(t, 1/1, utilInfo.Utilization, 0.01)

	// Node with Unready GPU
	gpuNode = BuildTestNode("gpu_node", 2000, 2000000)
	AddGpuLabelToNode(gpuNode)
	nodeInfo = schedulerframework.NewNodeInfo(pod, pod)
	utilInfo, err = Calculate(gpuNode, nodeInfo, false, false, gpuLabel, testTime)
	assert.NoError(t, err)
	assert.Zero(t, utilInfo.Utilization)
}

func nodeInfos(nodes []*apiv1.Node) []*schedulerframework.NodeInfo {
	result := make([]*schedulerframework.NodeInfo, len(nodes))
	for i, node := range nodes {
		ni := schedulerframework.NewNodeInfo()
		ni.SetNode(node)
		result[i] = ni
	}
	return result
}
