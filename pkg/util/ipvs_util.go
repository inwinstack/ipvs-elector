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
