package main

import (
	goflag "flag"
	"os"
	"time"

	flag "github.com/spf13/pflag"

	"github.com/golang/glog"
	"github.com/inwinstack/ipvs-elector/pkg/util"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"k8s.io/client-go/tools/record"
	utilsysctl "k8s.io/kubernetes/pkg/util/sysctl"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		glog.Fatalln(err)
	}

	electionid := flag.String("election", "ipvs-arp-election", "Leader election ID (name of configmap)")
	kubeconfig := flag.String("kubeconfig", "", "Absolute path to kubeconfig file")
	ttlseconds := flag.Int("ttl", 10, "TTL for leader election in seconds")
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.Parse()

	config, err := util.GetRestConfig(*kubeconfig)
	if err != nil {
		glog.Fatalln(err)
	}

	client := clientset.NewForConfigOrDie(config)
	broadcaster := record.NewBroadcaster()
	recorder := broadcaster.NewRecorder(scheme.Scheme, apiv1.EventSource{
		Component: *electionid,
		Host:      hostname,
	})

	pod, err := util.GetPodDetails(client)
	if err != nil {
		panic(err)
	}

	lock := resourcelock.ConfigMapLock{
		ConfigMapMeta: metav1.ObjectMeta{Namespace: pod.Namespace, Name: *electionid},
		Client:        client.CoreV1(),
		LockConfig: resourcelock.ResourceLockConfig{
			Identity:      hostname,
			EventRecorder: recorder,
		},
	}

	sysctl := utilsysctl.New()
	callbacks := leaderelection.LeaderCallbacks{
		OnStartedLeading: func(stop <-chan struct{}) {
			if err := util.EnableArpRequest(sysctl); err != nil {
				glog.Errorln(err)
			}
		},
		OnStoppedLeading: func() {
			glog.Info("Stopped leading")
		},
		OnNewLeader: func(identity string) {
			if identity != pod.Name {
				if err := util.DisableArpRequest(sysctl); err != nil {
					glog.Errorln(err)
				}
				return
			}
			glog.Infof("New leader elected: %v", identity)
		},
	}

	ttl := time.Duration(*ttlseconds) * time.Second
	le, err := leaderelection.NewLeaderElector(leaderelection.LeaderElectionConfig{
		Lock:          &lock,
		LeaseDuration: ttl,
		RenewDeadline: ttl / 2,
		RetryPeriod:   ttl / 4,
		Callbacks:     callbacks,
	})
	le.Run()
}
