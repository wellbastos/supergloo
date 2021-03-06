// Code generated by solo-kit. DO NOT EDIT.

// +build solokit

package v1alpha3

import (
	"context"
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/go-utils/kubeutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	kuberc "github.com/solo-io/solo-kit/pkg/api/v1/clients/kube"
	"github.com/solo-io/solo-kit/pkg/utils/log"
	"github.com/solo-io/solo-kit/test/helpers"
	"github.com/solo-io/solo-kit/test/setup"
	"k8s.io/client-go/rest"

	// Needed to run tests in GKE
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	// From https://github.com/kubernetes/client-go/blob/53c7adfd0294caa142d961e1f780f74081d5b15f/examples/out-of-cluster-client-configuration/main.go#L31
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

var _ = Describe("V1Alpha3Emitter", func() {
	if os.Getenv("RUN_KUBE_TESTS") != "1" {
		log.Printf("This test creates kubernetes resources and is disabled by default. To enable, set RUN_KUBE_TESTS=1 in your env.")
		return
	}
	var (
		namespace1           string
		namespace2           string
		name1, name2         = "angela" + helpers.RandString(3), "bob" + helpers.RandString(3)
		cfg                  *rest.Config
		emitter              RoutingEmitter
		virtualServiceClient VirtualServiceClient
	)

	BeforeEach(func() {
		namespace1 = helpers.RandString(8)
		namespace2 = helpers.RandString(8)
		var err error
		cfg, err = kubeutils.GetConfig("", "")
		Expect(err).NotTo(HaveOccurred())
		err = setup.SetupKubeForTest(namespace1)
		Expect(err).NotTo(HaveOccurred())
		err = setup.SetupKubeForTest(namespace2)
		Expect(err).NotTo(HaveOccurred())
		// VirtualService Constructor
		virtualServiceClientFactory := &factory.KubeResourceClientFactory{
			Crd:         VirtualServiceCrd,
			Cfg:         cfg,
			SharedCache: kuberc.NewKubeCache(),
		}
		virtualServiceClient, err = NewVirtualServiceClient(virtualServiceClientFactory)
		Expect(err).NotTo(HaveOccurred())
		emitter = NewRoutingEmitter(virtualServiceClient)
	})
	AfterEach(func() {
		setup.TeardownKube(namespace1)
		setup.TeardownKube(namespace2)
	})
	It("tracks snapshots on changes to any resource", func() {
		ctx := context.Background()
		err := emitter.Register()
		Expect(err).NotTo(HaveOccurred())

		snapshots, errs, err := emitter.Snapshots([]string{namespace1, namespace2}, clients.WatchOpts{
			Ctx:         ctx,
			RefreshRate: time.Second,
		})
		Expect(err).NotTo(HaveOccurred())

		var snap *RoutingSnapshot

		/*
			VirtualService
		*/

		assertSnapshotVirtualservices := func(expectVirtualservices VirtualServiceList, unexpectVirtualservices VirtualServiceList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectVirtualservices {
						if _, err := snap.Virtualservices.List().Find(expected.Metadata.Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectVirtualservices {
						if _, err := snap.Virtualservices.List().Find(unexpected.Metadata.Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := virtualServiceClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := virtualServiceClient.List(namespace2, clients.ListOpts{})
					combined := nsList1.ByNamespace()
					combined.Add(nsList2...)
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		virtualService1a, err := virtualServiceClient.Write(NewVirtualService(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		virtualService1b, err := virtualServiceClient.Write(NewVirtualService(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotVirtualservices(VirtualServiceList{virtualService1a, virtualService1b}, nil)
		virtualService2a, err := virtualServiceClient.Write(NewVirtualService(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		virtualService2b, err := virtualServiceClient.Write(NewVirtualService(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotVirtualservices(VirtualServiceList{virtualService1a, virtualService1b, virtualService2a, virtualService2b}, nil)

		err = virtualServiceClient.Delete(virtualService2a.Metadata.Namespace, virtualService2a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = virtualServiceClient.Delete(virtualService2b.Metadata.Namespace, virtualService2b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotVirtualservices(VirtualServiceList{virtualService1a, virtualService1b}, VirtualServiceList{virtualService2a, virtualService2b})

		err = virtualServiceClient.Delete(virtualService1a.Metadata.Namespace, virtualService1a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = virtualServiceClient.Delete(virtualService1b.Metadata.Namespace, virtualService1b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotVirtualservices(nil, VirtualServiceList{virtualService1a, virtualService1b, virtualService2a, virtualService2b})
	})
})
