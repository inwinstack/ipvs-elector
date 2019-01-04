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
	"testing"

	"github.com/stretchr/testify/assert"
	fake "k8s.io/kubernetes/pkg/util/sysctl/testing"
)

func TestEnableArpRequest(t *testing.T) {
	s := fake.NewFake()
	assert.Nil(t, EnableArpRequest(s))

	ignore, ignoreErr := s.GetSysctl(sysctlArpIgnore)
	assert.Nil(t, ignoreErr)
	assert.Equal(t, ignore, 0)

	announce, announceErr := s.GetSysctl(sysctlArpAnnounce)
	assert.Nil(t, announceErr)
	assert.Equal(t, announce, 0)
}

func TestDisableArpRequest(t *testing.T) {
	s := fake.NewFake()
	assert.Nil(t, DisableArpRequest(s))

	ignore, ignoreErr := s.GetSysctl(sysctlArpIgnore)
	assert.Nil(t, ignoreErr)
	assert.Equal(t, ignore, 1)

	announce, announceErr := s.GetSysctl(sysctlArpAnnounce)
	assert.Nil(t, announceErr)
	assert.Equal(t, announce, 2)
}
