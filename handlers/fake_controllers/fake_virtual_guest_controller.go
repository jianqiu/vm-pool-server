// This file was generated by counterfeiter
package fake_controllers

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/jianqiu/vm-pool-server/handlers"
	"github.com/jianqiu/vm-pool-server/models"
)

type FakeVirtualGuestController struct {
	VirtualGuestsStub        func(logger lager.Logger, domain, cellId string) ([]*models.VirtualGuest, error)
	virtualGuestsMutex       sync.RWMutex
	virtualGuestsArgsForCall []struct {
		logger lager.Logger
		domain string
		cellId string
	}
	virtualGuestsReturns struct {
		result1 []*models.VirtualGuest
		result2 error
	}
	VirtualGuestByCidStub        func(logger lager.Logger, taskGuid string) (*models.VirtualGuest, error)
	virtualGuestByCidMutex       sync.RWMutex
	virtualGuestByCidArgsForCall []struct {
		logger   lager.Logger
		taskGuid string
	}
	virtualGuestByCidReturns struct {
		result1 *models.VirtualGuest
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVirtualGuestController) VirtualGuests(logger lager.Logger, domain string, cellId string) ([]*models.VirtualGuest, error) {
	fake.virtualGuestsMutex.Lock()
	fake.virtualGuestsArgsForCall = append(fake.virtualGuestsArgsForCall, struct {
		logger lager.Logger
		domain string
		cellId string
	}{logger, domain, cellId})
	fake.recordInvocation("VirtualGuests", []interface{}{logger, domain, cellId})
	fake.virtualGuestsMutex.Unlock()
	if fake.VirtualGuestsStub != nil {
		return fake.VirtualGuestsStub(logger, domain, cellId)
	} else {
		return fake.virtualGuestsReturns.result1, fake.virtualGuestsReturns.result2
	}
}

func (fake *FakeVirtualGuestController) VirtualGuestsCallCount() int {
	fake.virtualGuestsMutex.RLock()
	defer fake.virtualGuestsMutex.RUnlock()
	return len(fake.virtualGuestsArgsForCall)
}

func (fake *FakeVirtualGuestController) VirtualGuestsArgsForCall(i int) (lager.Logger, string, string) {
	fake.virtualGuestsMutex.RLock()
	defer fake.virtualGuestsMutex.RUnlock()
	return fake.virtualGuestsArgsForCall[i].logger, fake.virtualGuestsArgsForCall[i].domain, fake.virtualGuestsArgsForCall[i].cellId
}

func (fake *FakeVirtualGuestController) VirtualGuestsReturns(result1 []*models.VirtualGuest, result2 error) {
	fake.VirtualGuestsStub = nil
	fake.virtualGuestsReturns = struct {
		result1 []*models.VirtualGuest
		result2 error
	}{result1, result2}
}

func (fake *FakeVirtualGuestController) VirtualGuestByCid(logger lager.Logger, taskGuid string) (*models.VirtualGuest, error) {
	fake.virtualGuestByCidMutex.Lock()
	fake.virtualGuestByCidArgsForCall = append(fake.virtualGuestByCidArgsForCall, struct {
		logger   lager.Logger
		taskGuid string
	}{logger, taskGuid})
	fake.recordInvocation("VirtualGuestByCid", []interface{}{logger, taskGuid})
	fake.virtualGuestByCidMutex.Unlock()
	if fake.VirtualGuestByCidStub != nil {
		return fake.VirtualGuestByCidStub(logger, taskGuid)
	} else {
		return fake.virtualGuestByCidReturns.result1, fake.virtualGuestByCidReturns.result2
	}
}

func (fake *FakeVirtualGuestController) VirtualGuestByCidCallCount() int {
	fake.virtualGuestByCidMutex.RLock()
	defer fake.virtualGuestByCidMutex.RUnlock()
	return len(fake.virtualGuestByCidArgsForCall)
}

func (fake *FakeVirtualGuestController) VirtualGuestByCidArgsForCall(i int) (lager.Logger, string) {
	fake.virtualGuestByCidMutex.RLock()
	defer fake.virtualGuestByCidMutex.RUnlock()
	return fake.virtualGuestByCidArgsForCall[i].logger, fake.virtualGuestByCidArgsForCall[i].taskGuid
}

func (fake *FakeVirtualGuestController) VirtualGuestByCidReturns(result1 *models.VirtualGuest, result2 error) {
	fake.VirtualGuestByCidStub = nil
	fake.virtualGuestByCidReturns = struct {
		result1 *models.VirtualGuest
		result2 error
	}{result1, result2}
}

func (fake *FakeVirtualGuestController) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.virtualGuestsMutex.RLock()
	defer fake.virtualGuestsMutex.RUnlock()
	fake.virtualGuestByCidMutex.RLock()
	defer fake.virtualGuestByCidMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeVirtualGuestController) recordInvocation(key string, args []interface{}) {
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

var _ handlers.VirtualGuestController = new(FakeVirtualGuestController)
