/*
Copyright 2020 The Kubernetes Authors.

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

package strategies

import (
	"context"
	v1 "k8s.io/api/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/klog"
	"sigs.k8s.io/descheduler/pkg/api"
	"sigs.k8s.io/descheduler/pkg/descheduler/evictions"
	nodeutil "sigs.k8s.io/descheduler/pkg/descheduler/node"
	podutil "sigs.k8s.io/descheduler/pkg/descheduler/pod"
)

// RemovePodsViolatingNodeSelector evicts pods from a node which no longer satisfies nodeSelector constraints.
func RemovePodsViolatingNodeSelector(ctx context.Context, client clientset.Interface, strategy api.DeschedulerStrategy, nodes []*v1.Node, evictLocalStoragePods bool, podEvictor *evictions.PodEvictor) {
	for _, node := range nodes {
		klog.V(1).Infof("Processing node: %#v\n", node.Name)

		if strategy.Params == nil || strategy.Params.NodeSelectionSettings == nil {
			// Params are optional, assign a safe default
			strategy.Params = &api.StrategyParameters{
				NodeSelectionSettings: &api.NodeSelectionSettings{
					DegradationAllowed: false,
				},
			}
		}

		pods, err := podutil.ListEvictablePodsOnNode(ctx, client, node, evictLocalStoragePods)
		if err != nil {
			klog.Errorf("failed to get pods from %v: %v", node.Name, err)
		}

		for _, pod := range pods {
			if pod.Spec.NodeSelector != nil {
				if !nodeutil.PodFitsCurrentNode(pod, node) && (nodeutil.PodFitsAnyNode(pod, nodes) || strategy.Params.NodeSelectionSettings.DegradationAllowed) {
					klog.V(1).Infof("Evicting pod: %v", pod.Name)
					if _, err := podEvictor.EvictPod(ctx, pod, node); err != nil {
						klog.Errorf("Error evicting pod: (%#v)", err)
						break
					}
				}
			}
		}
	}

	klog.V(1).Infof("Evicted %v pods", podEvictor.TotalEvicted())
}
