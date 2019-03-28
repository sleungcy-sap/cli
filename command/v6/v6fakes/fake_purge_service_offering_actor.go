// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v2action"
	v6 "code.cloudfoundry.org/cli/command/v6"
)

type FakePurgeServiceOfferingActor struct {
	GetServiceByNameAndBrokerNameStub        func(string, string) (v2action.Service, v2action.Warnings, error)
	getServiceByNameAndBrokerNameMutex       sync.RWMutex
	getServiceByNameAndBrokerNameArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getServiceByNameAndBrokerNameReturns struct {
		result1 v2action.Service
		result2 v2action.Warnings
		result3 error
	}
	getServiceByNameAndBrokerNameReturnsOnCall map[int]struct {
		result1 v2action.Service
		result2 v2action.Warnings
		result3 error
	}
	PurgeServiceOfferingStub        func(v2action.Service) (v2action.Warnings, error)
	purgeServiceOfferingMutex       sync.RWMutex
	purgeServiceOfferingArgsForCall []struct {
		arg1 v2action.Service
	}
	purgeServiceOfferingReturns struct {
		result1 v2action.Warnings
		result2 error
	}
	purgeServiceOfferingReturnsOnCall map[int]struct {
		result1 v2action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePurgeServiceOfferingActor) GetServiceByNameAndBrokerName(arg1 string, arg2 string) (v2action.Service, v2action.Warnings, error) {
	fake.getServiceByNameAndBrokerNameMutex.Lock()
	ret, specificReturn := fake.getServiceByNameAndBrokerNameReturnsOnCall[len(fake.getServiceByNameAndBrokerNameArgsForCall)]
	fake.getServiceByNameAndBrokerNameArgsForCall = append(fake.getServiceByNameAndBrokerNameArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetServiceByNameAndBrokerName", []interface{}{arg1, arg2})
	fake.getServiceByNameAndBrokerNameMutex.Unlock()
	if fake.GetServiceByNameAndBrokerNameStub != nil {
		return fake.GetServiceByNameAndBrokerNameStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getServiceByNameAndBrokerNameReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakePurgeServiceOfferingActor) GetServiceByNameAndBrokerNameCallCount() int {
	fake.getServiceByNameAndBrokerNameMutex.RLock()
	defer fake.getServiceByNameAndBrokerNameMutex.RUnlock()
	return len(fake.getServiceByNameAndBrokerNameArgsForCall)
}

func (fake *FakePurgeServiceOfferingActor) GetServiceByNameAndBrokerNameCalls(stub func(string, string) (v2action.Service, v2action.Warnings, error)) {
	fake.getServiceByNameAndBrokerNameMutex.Lock()
	defer fake.getServiceByNameAndBrokerNameMutex.Unlock()
	fake.GetServiceByNameAndBrokerNameStub = stub
}

func (fake *FakePurgeServiceOfferingActor) GetServiceByNameAndBrokerNameArgsForCall(i int) (string, string) {
	fake.getServiceByNameAndBrokerNameMutex.RLock()
	defer fake.getServiceByNameAndBrokerNameMutex.RUnlock()
	argsForCall := fake.getServiceByNameAndBrokerNameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePurgeServiceOfferingActor) GetServiceByNameAndBrokerNameReturns(result1 v2action.Service, result2 v2action.Warnings, result3 error) {
	fake.getServiceByNameAndBrokerNameMutex.Lock()
	defer fake.getServiceByNameAndBrokerNameMutex.Unlock()
	fake.GetServiceByNameAndBrokerNameStub = nil
	fake.getServiceByNameAndBrokerNameReturns = struct {
		result1 v2action.Service
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakePurgeServiceOfferingActor) GetServiceByNameAndBrokerNameReturnsOnCall(i int, result1 v2action.Service, result2 v2action.Warnings, result3 error) {
	fake.getServiceByNameAndBrokerNameMutex.Lock()
	defer fake.getServiceByNameAndBrokerNameMutex.Unlock()
	fake.GetServiceByNameAndBrokerNameStub = nil
	if fake.getServiceByNameAndBrokerNameReturnsOnCall == nil {
		fake.getServiceByNameAndBrokerNameReturnsOnCall = make(map[int]struct {
			result1 v2action.Service
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getServiceByNameAndBrokerNameReturnsOnCall[i] = struct {
		result1 v2action.Service
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakePurgeServiceOfferingActor) PurgeServiceOffering(arg1 v2action.Service) (v2action.Warnings, error) {
	fake.purgeServiceOfferingMutex.Lock()
	ret, specificReturn := fake.purgeServiceOfferingReturnsOnCall[len(fake.purgeServiceOfferingArgsForCall)]
	fake.purgeServiceOfferingArgsForCall = append(fake.purgeServiceOfferingArgsForCall, struct {
		arg1 v2action.Service
	}{arg1})
	fake.recordInvocation("PurgeServiceOffering", []interface{}{arg1})
	fake.purgeServiceOfferingMutex.Unlock()
	if fake.PurgeServiceOfferingStub != nil {
		return fake.PurgeServiceOfferingStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.purgeServiceOfferingReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePurgeServiceOfferingActor) PurgeServiceOfferingCallCount() int {
	fake.purgeServiceOfferingMutex.RLock()
	defer fake.purgeServiceOfferingMutex.RUnlock()
	return len(fake.purgeServiceOfferingArgsForCall)
}

func (fake *FakePurgeServiceOfferingActor) PurgeServiceOfferingCalls(stub func(v2action.Service) (v2action.Warnings, error)) {
	fake.purgeServiceOfferingMutex.Lock()
	defer fake.purgeServiceOfferingMutex.Unlock()
	fake.PurgeServiceOfferingStub = stub
}

func (fake *FakePurgeServiceOfferingActor) PurgeServiceOfferingArgsForCall(i int) v2action.Service {
	fake.purgeServiceOfferingMutex.RLock()
	defer fake.purgeServiceOfferingMutex.RUnlock()
	argsForCall := fake.purgeServiceOfferingArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePurgeServiceOfferingActor) PurgeServiceOfferingReturns(result1 v2action.Warnings, result2 error) {
	fake.purgeServiceOfferingMutex.Lock()
	defer fake.purgeServiceOfferingMutex.Unlock()
	fake.PurgeServiceOfferingStub = nil
	fake.purgeServiceOfferingReturns = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakePurgeServiceOfferingActor) PurgeServiceOfferingReturnsOnCall(i int, result1 v2action.Warnings, result2 error) {
	fake.purgeServiceOfferingMutex.Lock()
	defer fake.purgeServiceOfferingMutex.Unlock()
	fake.PurgeServiceOfferingStub = nil
	if fake.purgeServiceOfferingReturnsOnCall == nil {
		fake.purgeServiceOfferingReturnsOnCall = make(map[int]struct {
			result1 v2action.Warnings
			result2 error
		})
	}
	fake.purgeServiceOfferingReturnsOnCall[i] = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakePurgeServiceOfferingActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getServiceByNameAndBrokerNameMutex.RLock()
	defer fake.getServiceByNameAndBrokerNameMutex.RUnlock()
	fake.purgeServiceOfferingMutex.RLock()
	defer fake.purgeServiceOfferingMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePurgeServiceOfferingActor) recordInvocation(key string, args []interface{}) {
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

var _ v6.PurgeServiceOfferingActor = new(FakePurgeServiceOfferingActor)