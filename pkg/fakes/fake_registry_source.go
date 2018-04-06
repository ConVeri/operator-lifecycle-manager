// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/coreos-inc/alm/pkg/api/apis/clusterserviceversion/v1alpha1"
	"github.com/coreos-inc/alm/pkg/controller/registry"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

type FakeSource struct {
	FindCSVForPackageNameUnderChannelStub        func(packageName string, channelName string) (*v1alpha1.ClusterServiceVersion, error)
	findCSVForPackageNameUnderChannelMutex       sync.RWMutex
	findCSVForPackageNameUnderChannelArgsForCall []struct {
		packageName string
		channelName string
	}
	findCSVForPackageNameUnderChannelReturns struct {
		result1 *v1alpha1.ClusterServiceVersion
		result2 error
	}
	findCSVForPackageNameUnderChannelReturnsOnCall map[int]struct {
		result1 *v1alpha1.ClusterServiceVersion
		result2 error
	}
	FindReplacementCSVForPackageNameUnderChannelStub        func(packageName string, channelName string, csvName string) (*v1alpha1.ClusterServiceVersion, error)
	findReplacementCSVForPackageNameUnderChannelMutex       sync.RWMutex
	findReplacementCSVForPackageNameUnderChannelArgsForCall []struct {
		packageName string
		channelName string
		csvName     string
	}
	findReplacementCSVForPackageNameUnderChannelReturns struct {
		result1 *v1alpha1.ClusterServiceVersion
		result2 error
	}
	findReplacementCSVForPackageNameUnderChannelReturnsOnCall map[int]struct {
		result1 *v1alpha1.ClusterServiceVersion
		result2 error
	}
	AllPackagesStub        func() map[string]registry.PackageManifest
	allPackagesMutex       sync.RWMutex
	allPackagesArgsForCall []struct{}
	allPackagesReturns     struct {
		result1 map[string]registry.PackageManifest
	}
	allPackagesReturnsOnCall map[int]struct {
		result1 map[string]registry.PackageManifest
	}
	FindReplacementCSVForNameStub        func(name string) (*v1alpha1.ClusterServiceVersion, error)
	findReplacementCSVForNameMutex       sync.RWMutex
	findReplacementCSVForNameArgsForCall []struct {
		name string
	}
	findReplacementCSVForNameReturns struct {
		result1 *v1alpha1.ClusterServiceVersion
		result2 error
	}
	findReplacementCSVForNameReturnsOnCall map[int]struct {
		result1 *v1alpha1.ClusterServiceVersion
		result2 error
	}
	FindCSVByNameStub        func(name string) (*v1alpha1.ClusterServiceVersion, error)
	findCSVByNameMutex       sync.RWMutex
	findCSVByNameArgsForCall []struct {
		name string
	}
	findCSVByNameReturns struct {
		result1 *v1alpha1.ClusterServiceVersion
		result2 error
	}
	findCSVByNameReturnsOnCall map[int]struct {
		result1 *v1alpha1.ClusterServiceVersion
		result2 error
	}
	ListServicesStub        func() ([]v1alpha1.ClusterServiceVersion, error)
	listServicesMutex       sync.RWMutex
	listServicesArgsForCall []struct{}
	listServicesReturns     struct {
		result1 []v1alpha1.ClusterServiceVersion
		result2 error
	}
	listServicesReturnsOnCall map[int]struct {
		result1 []v1alpha1.ClusterServiceVersion
		result2 error
	}
	FindCRDByKeyStub        func(key registry.CRDKey) (*v1beta1.CustomResourceDefinition, error)
	findCRDByKeyMutex       sync.RWMutex
	findCRDByKeyArgsForCall []struct {
		key registry.CRDKey
	}
	findCRDByKeyReturns struct {
		result1 *v1beta1.CustomResourceDefinition
		result2 error
	}
	findCRDByKeyReturnsOnCall map[int]struct {
		result1 *v1beta1.CustomResourceDefinition
		result2 error
	}
	ListLatestCSVsForCRDStub        func(key registry.CRDKey) ([]registry.CSVAndChannelInfo, error)
	listLatestCSVsForCRDMutex       sync.RWMutex
	listLatestCSVsForCRDArgsForCall []struct {
		key registry.CRDKey
	}
	listLatestCSVsForCRDReturns struct {
		result1 []registry.CSVAndChannelInfo
		result2 error
	}
	listLatestCSVsForCRDReturnsOnCall map[int]struct {
		result1 []registry.CSVAndChannelInfo
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSource) FindCSVForPackageNameUnderChannel(packageName string, channelName string) (*v1alpha1.ClusterServiceVersion, error) {
	fake.findCSVForPackageNameUnderChannelMutex.Lock()
	ret, specificReturn := fake.findCSVForPackageNameUnderChannelReturnsOnCall[len(fake.findCSVForPackageNameUnderChannelArgsForCall)]
	fake.findCSVForPackageNameUnderChannelArgsForCall = append(fake.findCSVForPackageNameUnderChannelArgsForCall, struct {
		packageName string
		channelName string
	}{packageName, channelName})
	fake.recordInvocation("FindCSVForPackageNameUnderChannel", []interface{}{packageName, channelName})
	fake.findCSVForPackageNameUnderChannelMutex.Unlock()
	if fake.FindCSVForPackageNameUnderChannelStub != nil {
		return fake.FindCSVForPackageNameUnderChannelStub(packageName, channelName)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findCSVForPackageNameUnderChannelReturns.result1, fake.findCSVForPackageNameUnderChannelReturns.result2
}

func (fake *FakeSource) FindCSVForPackageNameUnderChannelCallCount() int {
	fake.findCSVForPackageNameUnderChannelMutex.RLock()
	defer fake.findCSVForPackageNameUnderChannelMutex.RUnlock()
	return len(fake.findCSVForPackageNameUnderChannelArgsForCall)
}

func (fake *FakeSource) FindCSVForPackageNameUnderChannelArgsForCall(i int) (string, string) {
	fake.findCSVForPackageNameUnderChannelMutex.RLock()
	defer fake.findCSVForPackageNameUnderChannelMutex.RUnlock()
	return fake.findCSVForPackageNameUnderChannelArgsForCall[i].packageName, fake.findCSVForPackageNameUnderChannelArgsForCall[i].channelName
}

func (fake *FakeSource) FindCSVForPackageNameUnderChannelReturns(result1 *v1alpha1.ClusterServiceVersion, result2 error) {
	fake.FindCSVForPackageNameUnderChannelStub = nil
	fake.findCSVForPackageNameUnderChannelReturns = struct {
		result1 *v1alpha1.ClusterServiceVersion
		result2 error
	}{result1, result2}
}

func (fake *FakeSource) FindCSVForPackageNameUnderChannelReturnsOnCall(i int, result1 *v1alpha1.ClusterServiceVersion, result2 error) {
	fake.FindCSVForPackageNameUnderChannelStub = nil
	if fake.findCSVForPackageNameUnderChannelReturnsOnCall == nil {
		fake.findCSVForPackageNameUnderChannelReturnsOnCall = make(map[int]struct {
			result1 *v1alpha1.ClusterServiceVersion
			result2 error
		})
	}
	fake.findCSVForPackageNameUnderChannelReturnsOnCall[i] = struct {
		result1 *v1alpha1.ClusterServiceVersion
		result2 error
	}{result1, result2}
}

func (fake *FakeSource) FindReplacementCSVForPackageNameUnderChannel(packageName string, channelName string, csvName string) (*v1alpha1.ClusterServiceVersion, error) {
	fake.findReplacementCSVForPackageNameUnderChannelMutex.Lock()
	ret, specificReturn := fake.findReplacementCSVForPackageNameUnderChannelReturnsOnCall[len(fake.findReplacementCSVForPackageNameUnderChannelArgsForCall)]
	fake.findReplacementCSVForPackageNameUnderChannelArgsForCall = append(fake.findReplacementCSVForPackageNameUnderChannelArgsForCall, struct {
		packageName string
		channelName string
		csvName     string
	}{packageName, channelName, csvName})
	fake.recordInvocation("FindReplacementCSVForPackageNameUnderChannel", []interface{}{packageName, channelName, csvName})
	fake.findReplacementCSVForPackageNameUnderChannelMutex.Unlock()
	if fake.FindReplacementCSVForPackageNameUnderChannelStub != nil {
		return fake.FindReplacementCSVForPackageNameUnderChannelStub(packageName, channelName, csvName)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findReplacementCSVForPackageNameUnderChannelReturns.result1, fake.findReplacementCSVForPackageNameUnderChannelReturns.result2
}

func (fake *FakeSource) FindReplacementCSVForPackageNameUnderChannelCallCount() int {
	fake.findReplacementCSVForPackageNameUnderChannelMutex.RLock()
	defer fake.findReplacementCSVForPackageNameUnderChannelMutex.RUnlock()
	return len(fake.findReplacementCSVForPackageNameUnderChannelArgsForCall)
}

func (fake *FakeSource) FindReplacementCSVForPackageNameUnderChannelArgsForCall(i int) (string, string, string) {
	fake.findReplacementCSVForPackageNameUnderChannelMutex.RLock()
	defer fake.findReplacementCSVForPackageNameUnderChannelMutex.RUnlock()
	return fake.findReplacementCSVForPackageNameUnderChannelArgsForCall[i].packageName, fake.findReplacementCSVForPackageNameUnderChannelArgsForCall[i].channelName, fake.findReplacementCSVForPackageNameUnderChannelArgsForCall[i].csvName
}

func (fake *FakeSource) FindReplacementCSVForPackageNameUnderChannelReturns(result1 *v1alpha1.ClusterServiceVersion, result2 error) {
	fake.FindReplacementCSVForPackageNameUnderChannelStub = nil
	fake.findReplacementCSVForPackageNameUnderChannelReturns = struct {
		result1 *v1alpha1.ClusterServiceVersion
		result2 error
	}{result1, result2}
}

func (fake *FakeSource) FindReplacementCSVForPackageNameUnderChannelReturnsOnCall(i int, result1 *v1alpha1.ClusterServiceVersion, result2 error) {
	fake.FindReplacementCSVForPackageNameUnderChannelStub = nil
	if fake.findReplacementCSVForPackageNameUnderChannelReturnsOnCall == nil {
		fake.findReplacementCSVForPackageNameUnderChannelReturnsOnCall = make(map[int]struct {
			result1 *v1alpha1.ClusterServiceVersion
			result2 error
		})
	}
	fake.findReplacementCSVForPackageNameUnderChannelReturnsOnCall[i] = struct {
		result1 *v1alpha1.ClusterServiceVersion
		result2 error
	}{result1, result2}
}

func (fake *FakeSource) AllPackages() map[string]registry.PackageManifest {
	fake.allPackagesMutex.Lock()
	ret, specificReturn := fake.allPackagesReturnsOnCall[len(fake.allPackagesArgsForCall)]
	fake.allPackagesArgsForCall = append(fake.allPackagesArgsForCall, struct{}{})
	fake.recordInvocation("AllPackages", []interface{}{})
	fake.allPackagesMutex.Unlock()
	if fake.AllPackagesStub != nil {
		return fake.AllPackagesStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.allPackagesReturns.result1
}

func (fake *FakeSource) AllPackagesCallCount() int {
	fake.allPackagesMutex.RLock()
	defer fake.allPackagesMutex.RUnlock()
	return len(fake.allPackagesArgsForCall)
}

func (fake *FakeSource) AllPackagesReturns(result1 map[string]registry.PackageManifest) {
	fake.AllPackagesStub = nil
	fake.allPackagesReturns = struct {
		result1 map[string]registry.PackageManifest
	}{result1}
}

func (fake *FakeSource) AllPackagesReturnsOnCall(i int, result1 map[string]registry.PackageManifest) {
	fake.AllPackagesStub = nil
	if fake.allPackagesReturnsOnCall == nil {
		fake.allPackagesReturnsOnCall = make(map[int]struct {
			result1 map[string]registry.PackageManifest
		})
	}
	fake.allPackagesReturnsOnCall[i] = struct {
		result1 map[string]registry.PackageManifest
	}{result1}
}

func (fake *FakeSource) FindReplacementCSVForName(name string) (*v1alpha1.ClusterServiceVersion, error) {
	fake.findReplacementCSVForNameMutex.Lock()
	ret, specificReturn := fake.findReplacementCSVForNameReturnsOnCall[len(fake.findReplacementCSVForNameArgsForCall)]
	fake.findReplacementCSVForNameArgsForCall = append(fake.findReplacementCSVForNameArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("FindReplacementCSVForName", []interface{}{name})
	fake.findReplacementCSVForNameMutex.Unlock()
	if fake.FindReplacementCSVForNameStub != nil {
		return fake.FindReplacementCSVForNameStub(name)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findReplacementCSVForNameReturns.result1, fake.findReplacementCSVForNameReturns.result2
}

func (fake *FakeSource) FindReplacementCSVForNameCallCount() int {
	fake.findReplacementCSVForNameMutex.RLock()
	defer fake.findReplacementCSVForNameMutex.RUnlock()
	return len(fake.findReplacementCSVForNameArgsForCall)
}

func (fake *FakeSource) FindReplacementCSVForNameArgsForCall(i int) string {
	fake.findReplacementCSVForNameMutex.RLock()
	defer fake.findReplacementCSVForNameMutex.RUnlock()
	return fake.findReplacementCSVForNameArgsForCall[i].name
}

func (fake *FakeSource) FindReplacementCSVForNameReturns(result1 *v1alpha1.ClusterServiceVersion, result2 error) {
	fake.FindReplacementCSVForNameStub = nil
	fake.findReplacementCSVForNameReturns = struct {
		result1 *v1alpha1.ClusterServiceVersion
		result2 error
	}{result1, result2}
}

func (fake *FakeSource) FindReplacementCSVForNameReturnsOnCall(i int, result1 *v1alpha1.ClusterServiceVersion, result2 error) {
	fake.FindReplacementCSVForNameStub = nil
	if fake.findReplacementCSVForNameReturnsOnCall == nil {
		fake.findReplacementCSVForNameReturnsOnCall = make(map[int]struct {
			result1 *v1alpha1.ClusterServiceVersion
			result2 error
		})
	}
	fake.findReplacementCSVForNameReturnsOnCall[i] = struct {
		result1 *v1alpha1.ClusterServiceVersion
		result2 error
	}{result1, result2}
}

func (fake *FakeSource) FindCSVByName(name string) (*v1alpha1.ClusterServiceVersion, error) {
	fake.findCSVByNameMutex.Lock()
	ret, specificReturn := fake.findCSVByNameReturnsOnCall[len(fake.findCSVByNameArgsForCall)]
	fake.findCSVByNameArgsForCall = append(fake.findCSVByNameArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("FindCSVByName", []interface{}{name})
	fake.findCSVByNameMutex.Unlock()
	if fake.FindCSVByNameStub != nil {
		return fake.FindCSVByNameStub(name)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findCSVByNameReturns.result1, fake.findCSVByNameReturns.result2
}

func (fake *FakeSource) FindCSVByNameCallCount() int {
	fake.findCSVByNameMutex.RLock()
	defer fake.findCSVByNameMutex.RUnlock()
	return len(fake.findCSVByNameArgsForCall)
}

func (fake *FakeSource) FindCSVByNameArgsForCall(i int) string {
	fake.findCSVByNameMutex.RLock()
	defer fake.findCSVByNameMutex.RUnlock()
	return fake.findCSVByNameArgsForCall[i].name
}

func (fake *FakeSource) FindCSVByNameReturns(result1 *v1alpha1.ClusterServiceVersion, result2 error) {
	fake.FindCSVByNameStub = nil
	fake.findCSVByNameReturns = struct {
		result1 *v1alpha1.ClusterServiceVersion
		result2 error
	}{result1, result2}
}

func (fake *FakeSource) FindCSVByNameReturnsOnCall(i int, result1 *v1alpha1.ClusterServiceVersion, result2 error) {
	fake.FindCSVByNameStub = nil
	if fake.findCSVByNameReturnsOnCall == nil {
		fake.findCSVByNameReturnsOnCall = make(map[int]struct {
			result1 *v1alpha1.ClusterServiceVersion
			result2 error
		})
	}
	fake.findCSVByNameReturnsOnCall[i] = struct {
		result1 *v1alpha1.ClusterServiceVersion
		result2 error
	}{result1, result2}
}

func (fake *FakeSource) ListServices() ([]v1alpha1.ClusterServiceVersion, error) {
	fake.listServicesMutex.Lock()
	ret, specificReturn := fake.listServicesReturnsOnCall[len(fake.listServicesArgsForCall)]
	fake.listServicesArgsForCall = append(fake.listServicesArgsForCall, struct{}{})
	fake.recordInvocation("ListServices", []interface{}{})
	fake.listServicesMutex.Unlock()
	if fake.ListServicesStub != nil {
		return fake.ListServicesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listServicesReturns.result1, fake.listServicesReturns.result2
}

func (fake *FakeSource) ListServicesCallCount() int {
	fake.listServicesMutex.RLock()
	defer fake.listServicesMutex.RUnlock()
	return len(fake.listServicesArgsForCall)
}

func (fake *FakeSource) ListServicesReturns(result1 []v1alpha1.ClusterServiceVersion, result2 error) {
	fake.ListServicesStub = nil
	fake.listServicesReturns = struct {
		result1 []v1alpha1.ClusterServiceVersion
		result2 error
	}{result1, result2}
}

func (fake *FakeSource) ListServicesReturnsOnCall(i int, result1 []v1alpha1.ClusterServiceVersion, result2 error) {
	fake.ListServicesStub = nil
	if fake.listServicesReturnsOnCall == nil {
		fake.listServicesReturnsOnCall = make(map[int]struct {
			result1 []v1alpha1.ClusterServiceVersion
			result2 error
		})
	}
	fake.listServicesReturnsOnCall[i] = struct {
		result1 []v1alpha1.ClusterServiceVersion
		result2 error
	}{result1, result2}
}

func (fake *FakeSource) FindCRDByKey(key registry.CRDKey) (*v1beta1.CustomResourceDefinition, error) {
	fake.findCRDByKeyMutex.Lock()
	ret, specificReturn := fake.findCRDByKeyReturnsOnCall[len(fake.findCRDByKeyArgsForCall)]
	fake.findCRDByKeyArgsForCall = append(fake.findCRDByKeyArgsForCall, struct {
		key registry.CRDKey
	}{key})
	fake.recordInvocation("FindCRDByKey", []interface{}{key})
	fake.findCRDByKeyMutex.Unlock()
	if fake.FindCRDByKeyStub != nil {
		return fake.FindCRDByKeyStub(key)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findCRDByKeyReturns.result1, fake.findCRDByKeyReturns.result2
}

func (fake *FakeSource) FindCRDByKeyCallCount() int {
	fake.findCRDByKeyMutex.RLock()
	defer fake.findCRDByKeyMutex.RUnlock()
	return len(fake.findCRDByKeyArgsForCall)
}

func (fake *FakeSource) FindCRDByKeyArgsForCall(i int) registry.CRDKey {
	fake.findCRDByKeyMutex.RLock()
	defer fake.findCRDByKeyMutex.RUnlock()
	return fake.findCRDByKeyArgsForCall[i].key
}

func (fake *FakeSource) FindCRDByKeyReturns(result1 *v1beta1.CustomResourceDefinition, result2 error) {
	fake.FindCRDByKeyStub = nil
	fake.findCRDByKeyReturns = struct {
		result1 *v1beta1.CustomResourceDefinition
		result2 error
	}{result1, result2}
}

func (fake *FakeSource) FindCRDByKeyReturnsOnCall(i int, result1 *v1beta1.CustomResourceDefinition, result2 error) {
	fake.FindCRDByKeyStub = nil
	if fake.findCRDByKeyReturnsOnCall == nil {
		fake.findCRDByKeyReturnsOnCall = make(map[int]struct {
			result1 *v1beta1.CustomResourceDefinition
			result2 error
		})
	}
	fake.findCRDByKeyReturnsOnCall[i] = struct {
		result1 *v1beta1.CustomResourceDefinition
		result2 error
	}{result1, result2}
}

func (fake *FakeSource) ListLatestCSVsForCRD(key registry.CRDKey) ([]registry.CSVAndChannelInfo, error) {
	fake.listLatestCSVsForCRDMutex.Lock()
	ret, specificReturn := fake.listLatestCSVsForCRDReturnsOnCall[len(fake.listLatestCSVsForCRDArgsForCall)]
	fake.listLatestCSVsForCRDArgsForCall = append(fake.listLatestCSVsForCRDArgsForCall, struct {
		key registry.CRDKey
	}{key})
	fake.recordInvocation("ListLatestCSVsForCRD", []interface{}{key})
	fake.listLatestCSVsForCRDMutex.Unlock()
	if fake.ListLatestCSVsForCRDStub != nil {
		return fake.ListLatestCSVsForCRDStub(key)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listLatestCSVsForCRDReturns.result1, fake.listLatestCSVsForCRDReturns.result2
}

func (fake *FakeSource) ListLatestCSVsForCRDCallCount() int {
	fake.listLatestCSVsForCRDMutex.RLock()
	defer fake.listLatestCSVsForCRDMutex.RUnlock()
	return len(fake.listLatestCSVsForCRDArgsForCall)
}

func (fake *FakeSource) ListLatestCSVsForCRDArgsForCall(i int) registry.CRDKey {
	fake.listLatestCSVsForCRDMutex.RLock()
	defer fake.listLatestCSVsForCRDMutex.RUnlock()
	return fake.listLatestCSVsForCRDArgsForCall[i].key
}

func (fake *FakeSource) ListLatestCSVsForCRDReturns(result1 []registry.CSVAndChannelInfo, result2 error) {
	fake.ListLatestCSVsForCRDStub = nil
	fake.listLatestCSVsForCRDReturns = struct {
		result1 []registry.CSVAndChannelInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeSource) ListLatestCSVsForCRDReturnsOnCall(i int, result1 []registry.CSVAndChannelInfo, result2 error) {
	fake.ListLatestCSVsForCRDStub = nil
	if fake.listLatestCSVsForCRDReturnsOnCall == nil {
		fake.listLatestCSVsForCRDReturnsOnCall = make(map[int]struct {
			result1 []registry.CSVAndChannelInfo
			result2 error
		})
	}
	fake.listLatestCSVsForCRDReturnsOnCall[i] = struct {
		result1 []registry.CSVAndChannelInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeSource) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findCSVForPackageNameUnderChannelMutex.RLock()
	defer fake.findCSVForPackageNameUnderChannelMutex.RUnlock()
	fake.findReplacementCSVForPackageNameUnderChannelMutex.RLock()
	defer fake.findReplacementCSVForPackageNameUnderChannelMutex.RUnlock()
	fake.allPackagesMutex.RLock()
	defer fake.allPackagesMutex.RUnlock()
	fake.findReplacementCSVForNameMutex.RLock()
	defer fake.findReplacementCSVForNameMutex.RUnlock()
	fake.findCSVByNameMutex.RLock()
	defer fake.findCSVByNameMutex.RUnlock()
	fake.listServicesMutex.RLock()
	defer fake.listServicesMutex.RUnlock()
	fake.findCRDByKeyMutex.RLock()
	defer fake.findCRDByKeyMutex.RUnlock()
	fake.listLatestCSVsForCRDMutex.RLock()
	defer fake.listLatestCSVsForCRDMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSource) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ registry.Source = new(FakeSource)
