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

package node

import (
	"testing"

	"github.com/kubernetes-incubator/descheduler/test"
	"k8s.io/kubernetes/pkg/api/v1"
	"k8s.io/kubernetes/pkg/client/clientset_generated/clientset/fake"
)

func TestReadyNodes(t *testing.T) {
	node1 := test.BuildTestNode("node1", 1000, 2000, 9)
	node1.Status.Conditions = []v1.NodeCondition{{Type: v1.NodeOutOfDisk, Status: v1.ConditionTrue}}
	node2 := test.BuildTestNode("node2", 1000, 2000, 9)
	node3 := test.BuildTestNode("node3", 1000, 2000, 9)
	node3.Status.Conditions = []v1.NodeCondition{{Type: v1.NodeMemoryPressure, Status: v1.ConditionTrue}}
	node4 := test.BuildTestNode("node4", 1000, 2000, 9)
	node4.Status.Conditions = []v1.NodeCondition{{Type: v1.NodeNetworkUnavailable, Status: v1.ConditionTrue}}
	node5 := test.BuildTestNode("node5", 1000, 2000, 9)
	node5.Status.Conditions = []v1.NodeCondition{{Type: v1.NodeReady, Status: v1.ConditionFalse}}

	if !IsReady(node1) {
		t.Errorf("Expected %v to be ready", node1.Name)
	}
	if !IsReady(node2) {
		t.Errorf("Expected %v to be ready", node2.Name)
	}
	if !IsReady(node3) {
		t.Errorf("Expected %v to be ready", node3.Name)
	}
	if !IsReady(node4) {
		t.Errorf("Expected %v to be ready", node4.Name)
	}
	if IsReady(node5) {
		t.Errorf("Expected %v to be not ready", node5.Name)
	}

}

func TestReadyNodesWithNodeSelector(t *testing.T) {
	node1 := test.BuildTestNode("node1", 1000, 2000, 9)
	node1.Labels = map[string]string{"type": "compute"}
	node2 := test.BuildTestNode("node2", 1000, 2000, 9)
	node2.Labels = map[string]string{"type": "infra"}

	fakeClient := fake.NewSimpleClientset(node1, node2)
	nodeSelector := "type=compute"
	nodes, _ := ReadyNodes(fakeClient, nodeSelector, nil)

	if nodes[0].Name != "node1" {
		t.Errorf("Expected node1, got %s", nodes[0].Name)
	}
}
