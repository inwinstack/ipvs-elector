/*
Copyright © 2018 inwinSTACK.inc.

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

package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestGetPodDetails(t *testing.T) {
	expected := &PodInfo{
		Name:      "test-pod1",
		Namespace: "default",
		NodeIP:    "",
		Labels:    map[string]string{"test": "123"},
	}

	assert.Nil(t, os.Setenv("POD_NAME", "test-pod1"))
	assert.Nil(t, os.Setenv("POD_NAMESPACE", "default"))

	client := fake.NewSimpleClientset()

	p := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:   "test-pod1",
			Labels: map[string]string{"test": "123"},
		},
	}
	_, err := client.Core().Pods("default").Create(p)
	assert.Nil(t, err)

	pod, err := GetPodDetails(client)
	assert.Nil(t, err)
	assert.Equal(t, expected, pod)
}

func TestGetNodeIPOrName(t *testing.T) {
	nodeName := "test-node1"
	nodeInternalIP := "192.168.99.100"
	nodeExternalIP := "140.99.100.22"
	client := fake.NewSimpleClientset()

	n := &v1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name: nodeName,
		},
		Status: v1.NodeStatus{
			Addresses: []v1.NodeAddress{
				v1.NodeAddress{
					Type:    v1.NodeInternalIP,
					Address: nodeInternalIP,
				},
				v1.NodeAddress{
					Type:    v1.NodeExternalIP,
					Address: nodeExternalIP,
				},
			},
		},
	}
	_, err := client.Core().Nodes().Create(n)
	assert.Nil(t, err)

	internalIP := GetNodeIPOrName(client, nodeName, true)
	assert.Equal(t, nodeInternalIP, internalIP)

	externalIP := GetNodeIPOrName(client, nodeName, false)
	assert.Equal(t, nodeExternalIP, externalIP)
}
