// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"policy-server/store"
	"sync"
)

type EgressPolicyStore struct {
	CreateStub        func([]store.EgressPolicy) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 []store.EgressPolicy
	}
	createReturns struct {
		result1 error
	}
	createReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStub        func([]store.EgressPolicy) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 []store.EgressPolicy
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	AllStub        func() ([]store.EgressPolicy, error)
	allMutex       sync.RWMutex
	allArgsForCall []struct{}
	allReturns     struct {
		result1 []store.EgressPolicy
		result2 error
	}
	allReturnsOnCall map[int]struct {
		result1 []store.EgressPolicy
		result2 error
	}
	ByGuidsStub        func(srcGuids []string) ([]store.EgressPolicy, error)
	byGuidsMutex       sync.RWMutex
	byGuidsArgsForCall []struct {
		srcGuids []string
	}
	byGuidsReturns struct {
		result1 []store.EgressPolicy
		result2 error
	}
	byGuidsReturnsOnCall map[int]struct {
		result1 []store.EgressPolicy
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *EgressPolicyStore) Create(arg1 []store.EgressPolicy) error {
	var arg1Copy []store.EgressPolicy
	if arg1 != nil {
		arg1Copy = make([]store.EgressPolicy, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 []store.EgressPolicy
	}{arg1Copy})
	fake.recordInvocation("Create", []interface{}{arg1Copy})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createReturns.result1
}

func (fake *EgressPolicyStore) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *EgressPolicyStore) CreateArgsForCall(i int) []store.EgressPolicy {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].arg1
}

func (fake *EgressPolicyStore) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyStore) CreateReturnsOnCall(i int, result1 error) {
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyStore) Delete(arg1 []store.EgressPolicy) error {
	var arg1Copy []store.EgressPolicy
	if arg1 != nil {
		arg1Copy = make([]store.EgressPolicy, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 []store.EgressPolicy
	}{arg1Copy})
	fake.recordInvocation("Delete", []interface{}{arg1Copy})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteReturns.result1
}

func (fake *EgressPolicyStore) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *EgressPolicyStore) DeleteArgsForCall(i int) []store.EgressPolicy {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].arg1
}

func (fake *EgressPolicyStore) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyStore) DeleteReturnsOnCall(i int, result1 error) {
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyStore) All() ([]store.EgressPolicy, error) {
	fake.allMutex.Lock()
	ret, specificReturn := fake.allReturnsOnCall[len(fake.allArgsForCall)]
	fake.allArgsForCall = append(fake.allArgsForCall, struct{}{})
	fake.recordInvocation("All", []interface{}{})
	fake.allMutex.Unlock()
	if fake.AllStub != nil {
		return fake.AllStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.allReturns.result1, fake.allReturns.result2
}

func (fake *EgressPolicyStore) AllCallCount() int {
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	return len(fake.allArgsForCall)
}

func (fake *EgressPolicyStore) AllReturns(result1 []store.EgressPolicy, result2 error) {
	fake.AllStub = nil
	fake.allReturns = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyStore) AllReturnsOnCall(i int, result1 []store.EgressPolicy, result2 error) {
	fake.AllStub = nil
	if fake.allReturnsOnCall == nil {
		fake.allReturnsOnCall = make(map[int]struct {
			result1 []store.EgressPolicy
			result2 error
		})
	}
	fake.allReturnsOnCall[i] = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyStore) ByGuids(srcGuids []string) ([]store.EgressPolicy, error) {
	var srcGuidsCopy []string
	if srcGuids != nil {
		srcGuidsCopy = make([]string, len(srcGuids))
		copy(srcGuidsCopy, srcGuids)
	}
	fake.byGuidsMutex.Lock()
	ret, specificReturn := fake.byGuidsReturnsOnCall[len(fake.byGuidsArgsForCall)]
	fake.byGuidsArgsForCall = append(fake.byGuidsArgsForCall, struct {
		srcGuids []string
	}{srcGuidsCopy})
	fake.recordInvocation("ByGuids", []interface{}{srcGuidsCopy})
	fake.byGuidsMutex.Unlock()
	if fake.ByGuidsStub != nil {
		return fake.ByGuidsStub(srcGuids)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.byGuidsReturns.result1, fake.byGuidsReturns.result2
}

func (fake *EgressPolicyStore) ByGuidsCallCount() int {
	fake.byGuidsMutex.RLock()
	defer fake.byGuidsMutex.RUnlock()
	return len(fake.byGuidsArgsForCall)
}

func (fake *EgressPolicyStore) ByGuidsArgsForCall(i int) []string {
	fake.byGuidsMutex.RLock()
	defer fake.byGuidsMutex.RUnlock()
	return fake.byGuidsArgsForCall[i].srcGuids
}

func (fake *EgressPolicyStore) ByGuidsReturns(result1 []store.EgressPolicy, result2 error) {
	fake.ByGuidsStub = nil
	fake.byGuidsReturns = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyStore) ByGuidsReturnsOnCall(i int, result1 []store.EgressPolicy, result2 error) {
	fake.ByGuidsStub = nil
	if fake.byGuidsReturnsOnCall == nil {
		fake.byGuidsReturnsOnCall = make(map[int]struct {
			result1 []store.EgressPolicy
			result2 error
		})
	}
	fake.byGuidsReturnsOnCall[i] = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	fake.byGuidsMutex.RLock()
	defer fake.byGuidsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *EgressPolicyStore) recordInvocation(key string, args []interface{}) {
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
