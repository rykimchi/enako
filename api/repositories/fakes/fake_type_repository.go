// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	sync "sync"

	models "github.com/BrandonWade/enako/api/models"
	repositories "github.com/BrandonWade/enako/api/repositories"
)

type FakeTypeRepository struct {
	GetTypesStub        func() ([]models.ExpenseType, error)
	getTypesMutex       sync.RWMutex
	getTypesArgsForCall []struct {
	}
	getTypesReturns struct {
		result1 []models.ExpenseType
		result2 error
	}
	getTypesReturnsOnCall map[int]struct {
		result1 []models.ExpenseType
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTypeRepository) GetTypes() ([]models.ExpenseType, error) {
	fake.getTypesMutex.Lock()
	ret, specificReturn := fake.getTypesReturnsOnCall[len(fake.getTypesArgsForCall)]
	fake.getTypesArgsForCall = append(fake.getTypesArgsForCall, struct {
	}{})
	fake.recordInvocation("GetTypes", []interface{}{})
	fake.getTypesMutex.Unlock()
	if fake.GetTypesStub != nil {
		return fake.GetTypesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getTypesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTypeRepository) GetTypesCallCount() int {
	fake.getTypesMutex.RLock()
	defer fake.getTypesMutex.RUnlock()
	return len(fake.getTypesArgsForCall)
}

func (fake *FakeTypeRepository) GetTypesCalls(stub func() ([]models.ExpenseType, error)) {
	fake.getTypesMutex.Lock()
	defer fake.getTypesMutex.Unlock()
	fake.GetTypesStub = stub
}

func (fake *FakeTypeRepository) GetTypesReturns(result1 []models.ExpenseType, result2 error) {
	fake.getTypesMutex.Lock()
	defer fake.getTypesMutex.Unlock()
	fake.GetTypesStub = nil
	fake.getTypesReturns = struct {
		result1 []models.ExpenseType
		result2 error
	}{result1, result2}
}

func (fake *FakeTypeRepository) GetTypesReturnsOnCall(i int, result1 []models.ExpenseType, result2 error) {
	fake.getTypesMutex.Lock()
	defer fake.getTypesMutex.Unlock()
	fake.GetTypesStub = nil
	if fake.getTypesReturnsOnCall == nil {
		fake.getTypesReturnsOnCall = make(map[int]struct {
			result1 []models.ExpenseType
			result2 error
		})
	}
	fake.getTypesReturnsOnCall[i] = struct {
		result1 []models.ExpenseType
		result2 error
	}{result1, result2}
}

func (fake *FakeTypeRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getTypesMutex.RLock()
	defer fake.getTypesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTypeRepository) recordInvocation(key string, args []interface{}) {
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

var _ repositories.TypeRepository = new(FakeTypeRepository)
