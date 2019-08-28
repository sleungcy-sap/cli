// Code generated by counterfeiter. DO NOT EDIT.
package applicationfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/cf/commandregistry"
	"code.cloudfoundry.org/cli/cf/commands/application"
	"code.cloudfoundry.org/cli/cf/flags"
	"code.cloudfoundry.org/cli/cf/models"
	"code.cloudfoundry.org/cli/cf/requirements"
)

type FakeStopper struct {
	ApplicationStopStub        func(models.Application, string, string) (models.Application, error)
	applicationStopMutex       sync.RWMutex
	applicationStopArgsForCall []struct {
		arg1 models.Application
		arg2 string
		arg3 string
	}
	applicationStopReturns struct {
		result1 models.Application
		result2 error
	}
	applicationStopReturnsOnCall map[int]struct {
		result1 models.Application
		result2 error
	}
	ExecuteStub        func(flags.FlagContext) error
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
		arg1 flags.FlagContext
	}
	executeReturns struct {
		result1 error
	}
	executeReturnsOnCall map[int]struct {
		result1 error
	}
	MetaDataStub        func() commandregistry.CommandMetadata
	metaDataMutex       sync.RWMutex
	metaDataArgsForCall []struct {
	}
	metaDataReturns struct {
		result1 commandregistry.CommandMetadata
	}
	metaDataReturnsOnCall map[int]struct {
		result1 commandregistry.CommandMetadata
	}
	RequirementsStub        func(requirements.Factory, flags.FlagContext) ([]requirements.Requirement, error)
	requirementsMutex       sync.RWMutex
	requirementsArgsForCall []struct {
		arg1 requirements.Factory
		arg2 flags.FlagContext
	}
	requirementsReturns struct {
		result1 []requirements.Requirement
		result2 error
	}
	requirementsReturnsOnCall map[int]struct {
		result1 []requirements.Requirement
		result2 error
	}
	SetDependencyStub        func(commandregistry.Dependency, bool) commandregistry.Command
	setDependencyMutex       sync.RWMutex
	setDependencyArgsForCall []struct {
		arg1 commandregistry.Dependency
		arg2 bool
	}
	setDependencyReturns struct {
		result1 commandregistry.Command
	}
	setDependencyReturnsOnCall map[int]struct {
		result1 commandregistry.Command
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStopper) ApplicationStop(arg1 models.Application, arg2 string, arg3 string) (models.Application, error) {
	fake.applicationStopMutex.Lock()
	ret, specificReturn := fake.applicationStopReturnsOnCall[len(fake.applicationStopArgsForCall)]
	fake.applicationStopArgsForCall = append(fake.applicationStopArgsForCall, struct {
		arg1 models.Application
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("ApplicationStop", []interface{}{arg1, arg2, arg3})
	fake.applicationStopMutex.Unlock()
	if fake.ApplicationStopStub != nil {
		return fake.ApplicationStopStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.applicationStopReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStopper) ApplicationStopCallCount() int {
	fake.applicationStopMutex.RLock()
	defer fake.applicationStopMutex.RUnlock()
	return len(fake.applicationStopArgsForCall)
}

func (fake *FakeStopper) ApplicationStopCalls(stub func(models.Application, string, string) (models.Application, error)) {
	fake.applicationStopMutex.Lock()
	defer fake.applicationStopMutex.Unlock()
	fake.ApplicationStopStub = stub
}

func (fake *FakeStopper) ApplicationStopArgsForCall(i int) (models.Application, string, string) {
	fake.applicationStopMutex.RLock()
	defer fake.applicationStopMutex.RUnlock()
	argsForCall := fake.applicationStopArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeStopper) ApplicationStopReturns(result1 models.Application, result2 error) {
	fake.applicationStopMutex.Lock()
	defer fake.applicationStopMutex.Unlock()
	fake.ApplicationStopStub = nil
	fake.applicationStopReturns = struct {
		result1 models.Application
		result2 error
	}{result1, result2}
}

func (fake *FakeStopper) ApplicationStopReturnsOnCall(i int, result1 models.Application, result2 error) {
	fake.applicationStopMutex.Lock()
	defer fake.applicationStopMutex.Unlock()
	fake.ApplicationStopStub = nil
	if fake.applicationStopReturnsOnCall == nil {
		fake.applicationStopReturnsOnCall = make(map[int]struct {
			result1 models.Application
			result2 error
		})
	}
	fake.applicationStopReturnsOnCall[i] = struct {
		result1 models.Application
		result2 error
	}{result1, result2}
}

func (fake *FakeStopper) Execute(arg1 flags.FlagContext) error {
	fake.executeMutex.Lock()
	ret, specificReturn := fake.executeReturnsOnCall[len(fake.executeArgsForCall)]
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
		arg1 flags.FlagContext
	}{arg1})
	fake.recordInvocation("Execute", []interface{}{arg1})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.executeReturns
	return fakeReturns.result1
}

func (fake *FakeStopper) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *FakeStopper) ExecuteCalls(stub func(flags.FlagContext) error) {
	fake.executeMutex.Lock()
	defer fake.executeMutex.Unlock()
	fake.ExecuteStub = stub
}

func (fake *FakeStopper) ExecuteArgsForCall(i int) flags.FlagContext {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	argsForCall := fake.executeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStopper) ExecuteReturns(result1 error) {
	fake.executeMutex.Lock()
	defer fake.executeMutex.Unlock()
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStopper) ExecuteReturnsOnCall(i int, result1 error) {
	fake.executeMutex.Lock()
	defer fake.executeMutex.Unlock()
	fake.ExecuteStub = nil
	if fake.executeReturnsOnCall == nil {
		fake.executeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.executeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStopper) MetaData() commandregistry.CommandMetadata {
	fake.metaDataMutex.Lock()
	ret, specificReturn := fake.metaDataReturnsOnCall[len(fake.metaDataArgsForCall)]
	fake.metaDataArgsForCall = append(fake.metaDataArgsForCall, struct {
	}{})
	fake.recordInvocation("MetaData", []interface{}{})
	fake.metaDataMutex.Unlock()
	if fake.MetaDataStub != nil {
		return fake.MetaDataStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.metaDataReturns
	return fakeReturns.result1
}

func (fake *FakeStopper) MetaDataCallCount() int {
	fake.metaDataMutex.RLock()
	defer fake.metaDataMutex.RUnlock()
	return len(fake.metaDataArgsForCall)
}

func (fake *FakeStopper) MetaDataCalls(stub func() commandregistry.CommandMetadata) {
	fake.metaDataMutex.Lock()
	defer fake.metaDataMutex.Unlock()
	fake.MetaDataStub = stub
}

func (fake *FakeStopper) MetaDataReturns(result1 commandregistry.CommandMetadata) {
	fake.metaDataMutex.Lock()
	defer fake.metaDataMutex.Unlock()
	fake.MetaDataStub = nil
	fake.metaDataReturns = struct {
		result1 commandregistry.CommandMetadata
	}{result1}
}

func (fake *FakeStopper) MetaDataReturnsOnCall(i int, result1 commandregistry.CommandMetadata) {
	fake.metaDataMutex.Lock()
	defer fake.metaDataMutex.Unlock()
	fake.MetaDataStub = nil
	if fake.metaDataReturnsOnCall == nil {
		fake.metaDataReturnsOnCall = make(map[int]struct {
			result1 commandregistry.CommandMetadata
		})
	}
	fake.metaDataReturnsOnCall[i] = struct {
		result1 commandregistry.CommandMetadata
	}{result1}
}

func (fake *FakeStopper) Requirements(arg1 requirements.Factory, arg2 flags.FlagContext) ([]requirements.Requirement, error) {
	fake.requirementsMutex.Lock()
	ret, specificReturn := fake.requirementsReturnsOnCall[len(fake.requirementsArgsForCall)]
	fake.requirementsArgsForCall = append(fake.requirementsArgsForCall, struct {
		arg1 requirements.Factory
		arg2 flags.FlagContext
	}{arg1, arg2})
	fake.recordInvocation("Requirements", []interface{}{arg1, arg2})
	fake.requirementsMutex.Unlock()
	if fake.RequirementsStub != nil {
		return fake.RequirementsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.requirementsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStopper) RequirementsCallCount() int {
	fake.requirementsMutex.RLock()
	defer fake.requirementsMutex.RUnlock()
	return len(fake.requirementsArgsForCall)
}

func (fake *FakeStopper) RequirementsCalls(stub func(requirements.Factory, flags.FlagContext) ([]requirements.Requirement, error)) {
	fake.requirementsMutex.Lock()
	defer fake.requirementsMutex.Unlock()
	fake.RequirementsStub = stub
}

func (fake *FakeStopper) RequirementsArgsForCall(i int) (requirements.Factory, flags.FlagContext) {
	fake.requirementsMutex.RLock()
	defer fake.requirementsMutex.RUnlock()
	argsForCall := fake.requirementsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStopper) RequirementsReturns(result1 []requirements.Requirement, result2 error) {
	fake.requirementsMutex.Lock()
	defer fake.requirementsMutex.Unlock()
	fake.RequirementsStub = nil
	fake.requirementsReturns = struct {
		result1 []requirements.Requirement
		result2 error
	}{result1, result2}
}

func (fake *FakeStopper) RequirementsReturnsOnCall(i int, result1 []requirements.Requirement, result2 error) {
	fake.requirementsMutex.Lock()
	defer fake.requirementsMutex.Unlock()
	fake.RequirementsStub = nil
	if fake.requirementsReturnsOnCall == nil {
		fake.requirementsReturnsOnCall = make(map[int]struct {
			result1 []requirements.Requirement
			result2 error
		})
	}
	fake.requirementsReturnsOnCall[i] = struct {
		result1 []requirements.Requirement
		result2 error
	}{result1, result2}
}

func (fake *FakeStopper) SetDependency(arg1 commandregistry.Dependency, arg2 bool) commandregistry.Command {
	fake.setDependencyMutex.Lock()
	ret, specificReturn := fake.setDependencyReturnsOnCall[len(fake.setDependencyArgsForCall)]
	fake.setDependencyArgsForCall = append(fake.setDependencyArgsForCall, struct {
		arg1 commandregistry.Dependency
		arg2 bool
	}{arg1, arg2})
	fake.recordInvocation("SetDependency", []interface{}{arg1, arg2})
	fake.setDependencyMutex.Unlock()
	if fake.SetDependencyStub != nil {
		return fake.SetDependencyStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setDependencyReturns
	return fakeReturns.result1
}

func (fake *FakeStopper) SetDependencyCallCount() int {
	fake.setDependencyMutex.RLock()
	defer fake.setDependencyMutex.RUnlock()
	return len(fake.setDependencyArgsForCall)
}

func (fake *FakeStopper) SetDependencyCalls(stub func(commandregistry.Dependency, bool) commandregistry.Command) {
	fake.setDependencyMutex.Lock()
	defer fake.setDependencyMutex.Unlock()
	fake.SetDependencyStub = stub
}

func (fake *FakeStopper) SetDependencyArgsForCall(i int) (commandregistry.Dependency, bool) {
	fake.setDependencyMutex.RLock()
	defer fake.setDependencyMutex.RUnlock()
	argsForCall := fake.setDependencyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStopper) SetDependencyReturns(result1 commandregistry.Command) {
	fake.setDependencyMutex.Lock()
	defer fake.setDependencyMutex.Unlock()
	fake.SetDependencyStub = nil
	fake.setDependencyReturns = struct {
		result1 commandregistry.Command
	}{result1}
}

func (fake *FakeStopper) SetDependencyReturnsOnCall(i int, result1 commandregistry.Command) {
	fake.setDependencyMutex.Lock()
	defer fake.setDependencyMutex.Unlock()
	fake.SetDependencyStub = nil
	if fake.setDependencyReturnsOnCall == nil {
		fake.setDependencyReturnsOnCall = make(map[int]struct {
			result1 commandregistry.Command
		})
	}
	fake.setDependencyReturnsOnCall[i] = struct {
		result1 commandregistry.Command
	}{result1}
}

func (fake *FakeStopper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.applicationStopMutex.RLock()
	defer fake.applicationStopMutex.RUnlock()
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	fake.metaDataMutex.RLock()
	defer fake.metaDataMutex.RUnlock()
	fake.requirementsMutex.RLock()
	defer fake.requirementsMutex.RUnlock()
	fake.setDependencyMutex.RLock()
	defer fake.setDependencyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStopper) recordInvocation(key string, args []interface{}) {
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

var _ application.Stopper = new(FakeStopper)