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
	"fmt"

	utilsysctl "k8s.io/kubernetes/pkg/util/sysctl"
)

const sysctlArpIgnore = "net/ipv4/conf/all/arp_ignore"
const sysctlArpAnnounce = "net/ipv4/conf/all/arp_announce"

func EnableArpRequest(sysctl utilsysctl.Interface) error {
	if val, _ := sysctl.GetSysctl(sysctlArpIgnore); val != 0 {
		if err := sysctl.SetSysctl(sysctlArpIgnore, 0); err != nil {
			return fmt.Errorf("can't set sysctl %s: %v", sysctlArpIgnore, err)
		}
	}

	if val, _ := sysctl.GetSysctl(sysctlArpAnnounce); val != 0 {
		if err := sysctl.SetSysctl(sysctlArpAnnounce, 0); err != nil {
			return fmt.Errorf("can't set sysctl %s: %v", sysctlArpAnnounce, err)
		}
	}
	return nil
}

func DisableArpRequest(sysctl utilsysctl.Interface) error {
	if val, _ := sysctl.GetSysctl(sysctlArpIgnore); val != 1 {
		if err := sysctl.SetSysctl(sysctlArpIgnore, 1); err != nil {
			return fmt.Errorf("can't set sysctl %s: %v", sysctlArpIgnore, err)
		}
	}

	if val, _ := sysctl.GetSysctl(sysctlArpAnnounce); val != 2 {
		if err := sysctl.SetSysctl(sysctlArpAnnounce, 2); err != nil {
			return fmt.Errorf("can't set sysctl %s: %v", sysctlArpAnnounce, err)
		}
	}
	return nil
}
