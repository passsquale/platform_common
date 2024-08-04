// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/a1exCross/common/pkg/storage.Redis -o redis.go -n RedisMock -p mocks

import (
	"sync"
	mm_atomic "sync/atomic"
	"time"
	mm_time "time"

	"github.com/go-redis/redis"
	"github.com/gojuno/minimock/v3"
)

// RedisMock implements storage.Redis
type RedisMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcClose          func() (err error)
	inspectFuncClose   func()
	afterCloseCounter  uint64
	beforeCloseCounter uint64
	CloseMock          mRedisMockClose

	funcGet          func(s1 string) (sp1 *redis.StringCmd)
	inspectFuncGet   func(s1 string)
	afterGetCounter  uint64
	beforeGetCounter uint64
	GetMock          mRedisMockGet

	funcPing          func() (err error)
	inspectFuncPing   func()
	afterPingCounter  uint64
	beforePingCounter uint64
	PingMock          mRedisMockPing

	funcSet          func(s1 string, p1 interface{}, d1 time.Duration) (sp1 *redis.StatusCmd)
	inspectFuncSet   func(s1 string, p1 interface{}, d1 time.Duration)
	afterSetCounter  uint64
	beforeSetCounter uint64
	SetMock          mRedisMockSet
}

// NewRedisMock returns a mock for storage.Redis
func NewRedisMock(t minimock.Tester) *RedisMock {
	m := &RedisMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CloseMock = mRedisMockClose{mock: m}

	m.GetMock = mRedisMockGet{mock: m}
	m.GetMock.callArgs = []*RedisMockGetParams{}

	m.PingMock = mRedisMockPing{mock: m}

	m.SetMock = mRedisMockSet{mock: m}
	m.SetMock.callArgs = []*RedisMockSetParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mRedisMockClose struct {
	mock               *RedisMock
	defaultExpectation *RedisMockCloseExpectation
	expectations       []*RedisMockCloseExpectation
}

// RedisMockCloseExpectation specifies expectation struct of the Redis.Close
type RedisMockCloseExpectation struct {
	mock *RedisMock

	results *RedisMockCloseResults
	Counter uint64
}

// RedisMockCloseResults contains results of the Redis.Close
type RedisMockCloseResults struct {
	err error
}

// Expect sets up expected params for Redis.Close
func (mmClose *mRedisMockClose) Expect() *mRedisMockClose {
	if mmClose.mock.funcClose != nil {
		mmClose.mock.t.Fatalf("RedisMock.Close mock is already set by Set")
	}

	if mmClose.defaultExpectation == nil {
		mmClose.defaultExpectation = &RedisMockCloseExpectation{}
	}

	return mmClose
}

// Inspect accepts an inspector function that has same arguments as the Redis.Close
func (mmClose *mRedisMockClose) Inspect(f func()) *mRedisMockClose {
	if mmClose.mock.inspectFuncClose != nil {
		mmClose.mock.t.Fatalf("Inspect function is already set for RedisMock.Close")
	}

	mmClose.mock.inspectFuncClose = f

	return mmClose
}

// Return sets up results that will be returned by Redis.Close
func (mmClose *mRedisMockClose) Return(err error) *RedisMock {
	if mmClose.mock.funcClose != nil {
		mmClose.mock.t.Fatalf("RedisMock.Close mock is already set by Set")
	}

	if mmClose.defaultExpectation == nil {
		mmClose.defaultExpectation = &RedisMockCloseExpectation{mock: mmClose.mock}
	}
	mmClose.defaultExpectation.results = &RedisMockCloseResults{err}
	return mmClose.mock
}

// Set uses given function f to mock the Redis.Close method
func (mmClose *mRedisMockClose) Set(f func() (err error)) *RedisMock {
	if mmClose.defaultExpectation != nil {
		mmClose.mock.t.Fatalf("Default expectation is already set for the Redis.Close method")
	}

	if len(mmClose.expectations) > 0 {
		mmClose.mock.t.Fatalf("Some expectations are already set for the Redis.Close method")
	}

	mmClose.mock.funcClose = f
	return mmClose.mock
}

// Close implements storage.Redis
func (mmClose *RedisMock) Close() (err error) {
	mm_atomic.AddUint64(&mmClose.beforeCloseCounter, 1)
	defer mm_atomic.AddUint64(&mmClose.afterCloseCounter, 1)

	if mmClose.inspectFuncClose != nil {
		mmClose.inspectFuncClose()
	}

	if mmClose.CloseMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmClose.CloseMock.defaultExpectation.Counter, 1)

		mm_results := mmClose.CloseMock.defaultExpectation.results
		if mm_results == nil {
			mmClose.t.Fatal("No results are set for the RedisMock.Close")
		}
		return (*mm_results).err
	}
	if mmClose.funcClose != nil {
		return mmClose.funcClose()
	}
	mmClose.t.Fatalf("Unexpected call to RedisMock.Close.")
	return
}

// CloseAfterCounter returns a count of finished RedisMock.Close invocations
func (mmClose *RedisMock) CloseAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmClose.afterCloseCounter)
}

// CloseBeforeCounter returns a count of RedisMock.Close invocations
func (mmClose *RedisMock) CloseBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmClose.beforeCloseCounter)
}

// MinimockCloseDone returns true if the count of the Close invocations corresponds
// the number of defined expectations
func (m *RedisMock) MinimockCloseDone() bool {
	for _, e := range m.CloseMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CloseMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCloseCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcClose != nil && mm_atomic.LoadUint64(&m.afterCloseCounter) < 1 {
		return false
	}
	return true
}

// MinimockCloseInspect logs each unmet expectation
func (m *RedisMock) MinimockCloseInspect() {
	for _, e := range m.CloseMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to RedisMock.Close")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CloseMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCloseCounter) < 1 {
		m.t.Error("Expected call to RedisMock.Close")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcClose != nil && mm_atomic.LoadUint64(&m.afterCloseCounter) < 1 {
		m.t.Error("Expected call to RedisMock.Close")
	}
}

type mRedisMockGet struct {
	mock               *RedisMock
	defaultExpectation *RedisMockGetExpectation
	expectations       []*RedisMockGetExpectation

	callArgs []*RedisMockGetParams
	mutex    sync.RWMutex
}

// RedisMockGetExpectation specifies expectation struct of the Redis.Get
type RedisMockGetExpectation struct {
	mock    *RedisMock
	params  *RedisMockGetParams
	results *RedisMockGetResults
	Counter uint64
}

// RedisMockGetParams contains parameters of the Redis.Get
type RedisMockGetParams struct {
	s1 string
}

// RedisMockGetResults contains results of the Redis.Get
type RedisMockGetResults struct {
	sp1 *redis.StringCmd
}

// Expect sets up expected params for Redis.Get
func (mmGet *mRedisMockGet) Expect(s1 string) *mRedisMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("RedisMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &RedisMockGetExpectation{}
	}

	mmGet.defaultExpectation.params = &RedisMockGetParams{s1}
	for _, e := range mmGet.expectations {
		if minimock.Equal(e.params, mmGet.defaultExpectation.params) {
			mmGet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGet.defaultExpectation.params)
		}
	}

	return mmGet
}

// Inspect accepts an inspector function that has same arguments as the Redis.Get
func (mmGet *mRedisMockGet) Inspect(f func(s1 string)) *mRedisMockGet {
	if mmGet.mock.inspectFuncGet != nil {
		mmGet.mock.t.Fatalf("Inspect function is already set for RedisMock.Get")
	}

	mmGet.mock.inspectFuncGet = f

	return mmGet
}

// Return sets up results that will be returned by Redis.Get
func (mmGet *mRedisMockGet) Return(sp1 *redis.StringCmd) *RedisMock {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("RedisMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &RedisMockGetExpectation{mock: mmGet.mock}
	}
	mmGet.defaultExpectation.results = &RedisMockGetResults{sp1}
	return mmGet.mock
}

// Set uses given function f to mock the Redis.Get method
func (mmGet *mRedisMockGet) Set(f func(s1 string) (sp1 *redis.StringCmd)) *RedisMock {
	if mmGet.defaultExpectation != nil {
		mmGet.mock.t.Fatalf("Default expectation is already set for the Redis.Get method")
	}

	if len(mmGet.expectations) > 0 {
		mmGet.mock.t.Fatalf("Some expectations are already set for the Redis.Get method")
	}

	mmGet.mock.funcGet = f
	return mmGet.mock
}

// When sets expectation for the Redis.Get which will trigger the result defined by the following
// Then helper
func (mmGet *mRedisMockGet) When(s1 string) *RedisMockGetExpectation {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("RedisMock.Get mock is already set by Set")
	}

	expectation := &RedisMockGetExpectation{
		mock:   mmGet.mock,
		params: &RedisMockGetParams{s1},
	}
	mmGet.expectations = append(mmGet.expectations, expectation)
	return expectation
}

// Then sets up Redis.Get return parameters for the expectation previously defined by the When method
func (e *RedisMockGetExpectation) Then(sp1 *redis.StringCmd) *RedisMock {
	e.results = &RedisMockGetResults{sp1}
	return e.mock
}

// Get implements storage.Redis
func (mmGet *RedisMock) Get(s1 string) (sp1 *redis.StringCmd) {
	mm_atomic.AddUint64(&mmGet.beforeGetCounter, 1)
	defer mm_atomic.AddUint64(&mmGet.afterGetCounter, 1)

	if mmGet.inspectFuncGet != nil {
		mmGet.inspectFuncGet(s1)
	}

	mm_params := RedisMockGetParams{s1}

	// Record call args
	mmGet.GetMock.mutex.Lock()
	mmGet.GetMock.callArgs = append(mmGet.GetMock.callArgs, &mm_params)
	mmGet.GetMock.mutex.Unlock()

	for _, e := range mmGet.GetMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.sp1
		}
	}

	if mmGet.GetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGet.GetMock.defaultExpectation.Counter, 1)
		mm_want := mmGet.GetMock.defaultExpectation.params
		mm_got := RedisMockGetParams{s1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGet.t.Errorf("RedisMock.Get got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGet.GetMock.defaultExpectation.results
		if mm_results == nil {
			mmGet.t.Fatal("No results are set for the RedisMock.Get")
		}
		return (*mm_results).sp1
	}
	if mmGet.funcGet != nil {
		return mmGet.funcGet(s1)
	}
	mmGet.t.Fatalf("Unexpected call to RedisMock.Get. %v", s1)
	return
}

// GetAfterCounter returns a count of finished RedisMock.Get invocations
func (mmGet *RedisMock) GetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.afterGetCounter)
}

// GetBeforeCounter returns a count of RedisMock.Get invocations
func (mmGet *RedisMock) GetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.beforeGetCounter)
}

// Calls returns a list of arguments used in each call to RedisMock.Get.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGet *mRedisMockGet) Calls() []*RedisMockGetParams {
	mmGet.mutex.RLock()

	argCopy := make([]*RedisMockGetParams, len(mmGet.callArgs))
	copy(argCopy, mmGet.callArgs)

	mmGet.mutex.RUnlock()

	return argCopy
}

// MinimockGetDone returns true if the count of the Get invocations corresponds
// the number of defined expectations
func (m *RedisMock) MinimockGetDone() bool {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetInspect logs each unmet expectation
func (m *RedisMock) MinimockGetInspect() {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RedisMock.Get with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		if m.GetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RedisMock.Get")
		} else {
			m.t.Errorf("Expected call to RedisMock.Get with params: %#v", *m.GetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		m.t.Error("Expected call to RedisMock.Get")
	}
}

type mRedisMockPing struct {
	mock               *RedisMock
	defaultExpectation *RedisMockPingExpectation
	expectations       []*RedisMockPingExpectation
}

// RedisMockPingExpectation specifies expectation struct of the Redis.Ping
type RedisMockPingExpectation struct {
	mock *RedisMock

	results *RedisMockPingResults
	Counter uint64
}

// RedisMockPingResults contains results of the Redis.Ping
type RedisMockPingResults struct {
	err error
}

// Expect sets up expected params for Redis.Ping
func (mmPing *mRedisMockPing) Expect() *mRedisMockPing {
	if mmPing.mock.funcPing != nil {
		mmPing.mock.t.Fatalf("RedisMock.Ping mock is already set by Set")
	}

	if mmPing.defaultExpectation == nil {
		mmPing.defaultExpectation = &RedisMockPingExpectation{}
	}

	return mmPing
}

// Inspect accepts an inspector function that has same arguments as the Redis.Ping
func (mmPing *mRedisMockPing) Inspect(f func()) *mRedisMockPing {
	if mmPing.mock.inspectFuncPing != nil {
		mmPing.mock.t.Fatalf("Inspect function is already set for RedisMock.Ping")
	}

	mmPing.mock.inspectFuncPing = f

	return mmPing
}

// Return sets up results that will be returned by Redis.Ping
func (mmPing *mRedisMockPing) Return(err error) *RedisMock {
	if mmPing.mock.funcPing != nil {
		mmPing.mock.t.Fatalf("RedisMock.Ping mock is already set by Set")
	}

	if mmPing.defaultExpectation == nil {
		mmPing.defaultExpectation = &RedisMockPingExpectation{mock: mmPing.mock}
	}
	mmPing.defaultExpectation.results = &RedisMockPingResults{err}
	return mmPing.mock
}

// Set uses given function f to mock the Redis.Ping method
func (mmPing *mRedisMockPing) Set(f func() (err error)) *RedisMock {
	if mmPing.defaultExpectation != nil {
		mmPing.mock.t.Fatalf("Default expectation is already set for the Redis.Ping method")
	}

	if len(mmPing.expectations) > 0 {
		mmPing.mock.t.Fatalf("Some expectations are already set for the Redis.Ping method")
	}

	mmPing.mock.funcPing = f
	return mmPing.mock
}

// Ping implements storage.Redis
func (mmPing *RedisMock) Ping() (err error) {
	mm_atomic.AddUint64(&mmPing.beforePingCounter, 1)
	defer mm_atomic.AddUint64(&mmPing.afterPingCounter, 1)

	if mmPing.inspectFuncPing != nil {
		mmPing.inspectFuncPing()
	}

	if mmPing.PingMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmPing.PingMock.defaultExpectation.Counter, 1)

		mm_results := mmPing.PingMock.defaultExpectation.results
		if mm_results == nil {
			mmPing.t.Fatal("No results are set for the RedisMock.Ping")
		}
		return (*mm_results).err
	}
	if mmPing.funcPing != nil {
		return mmPing.funcPing()
	}
	mmPing.t.Fatalf("Unexpected call to RedisMock.Ping.")
	return
}

// PingAfterCounter returns a count of finished RedisMock.Ping invocations
func (mmPing *RedisMock) PingAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmPing.afterPingCounter)
}

// PingBeforeCounter returns a count of RedisMock.Ping invocations
func (mmPing *RedisMock) PingBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmPing.beforePingCounter)
}

// MinimockPingDone returns true if the count of the Ping invocations corresponds
// the number of defined expectations
func (m *RedisMock) MinimockPingDone() bool {
	for _, e := range m.PingMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.PingMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterPingCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcPing != nil && mm_atomic.LoadUint64(&m.afterPingCounter) < 1 {
		return false
	}
	return true
}

// MinimockPingInspect logs each unmet expectation
func (m *RedisMock) MinimockPingInspect() {
	for _, e := range m.PingMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to RedisMock.Ping")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.PingMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterPingCounter) < 1 {
		m.t.Error("Expected call to RedisMock.Ping")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcPing != nil && mm_atomic.LoadUint64(&m.afterPingCounter) < 1 {
		m.t.Error("Expected call to RedisMock.Ping")
	}
}

type mRedisMockSet struct {
	mock               *RedisMock
	defaultExpectation *RedisMockSetExpectation
	expectations       []*RedisMockSetExpectation

	callArgs []*RedisMockSetParams
	mutex    sync.RWMutex
}

// RedisMockSetExpectation specifies expectation struct of the Redis.Set
type RedisMockSetExpectation struct {
	mock    *RedisMock
	params  *RedisMockSetParams
	results *RedisMockSetResults
	Counter uint64
}

// RedisMockSetParams contains parameters of the Redis.Set
type RedisMockSetParams struct {
	s1 string
	p1 interface{}
	d1 time.Duration
}

// RedisMockSetResults contains results of the Redis.Set
type RedisMockSetResults struct {
	sp1 *redis.StatusCmd
}

// Expect sets up expected params for Redis.Set
func (mmSet *mRedisMockSet) Expect(s1 string, p1 interface{}, d1 time.Duration) *mRedisMockSet {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("RedisMock.Set mock is already set by Set")
	}

	if mmSet.defaultExpectation == nil {
		mmSet.defaultExpectation = &RedisMockSetExpectation{}
	}

	mmSet.defaultExpectation.params = &RedisMockSetParams{s1, p1, d1}
	for _, e := range mmSet.expectations {
		if minimock.Equal(e.params, mmSet.defaultExpectation.params) {
			mmSet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSet.defaultExpectation.params)
		}
	}

	return mmSet
}

// Inspect accepts an inspector function that has same arguments as the Redis.Set
func (mmSet *mRedisMockSet) Inspect(f func(s1 string, p1 interface{}, d1 time.Duration)) *mRedisMockSet {
	if mmSet.mock.inspectFuncSet != nil {
		mmSet.mock.t.Fatalf("Inspect function is already set for RedisMock.Set")
	}

	mmSet.mock.inspectFuncSet = f

	return mmSet
}

// Return sets up results that will be returned by Redis.Set
func (mmSet *mRedisMockSet) Return(sp1 *redis.StatusCmd) *RedisMock {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("RedisMock.Set mock is already set by Set")
	}

	if mmSet.defaultExpectation == nil {
		mmSet.defaultExpectation = &RedisMockSetExpectation{mock: mmSet.mock}
	}
	mmSet.defaultExpectation.results = &RedisMockSetResults{sp1}
	return mmSet.mock
}

// Set uses given function f to mock the Redis.Set method
func (mmSet *mRedisMockSet) Set(f func(s1 string, p1 interface{}, d1 time.Duration) (sp1 *redis.StatusCmd)) *RedisMock {
	if mmSet.defaultExpectation != nil {
		mmSet.mock.t.Fatalf("Default expectation is already set for the Redis.Set method")
	}

	if len(mmSet.expectations) > 0 {
		mmSet.mock.t.Fatalf("Some expectations are already set for the Redis.Set method")
	}

	mmSet.mock.funcSet = f
	return mmSet.mock
}

// When sets expectation for the Redis.Set which will trigger the result defined by the following
// Then helper
func (mmSet *mRedisMockSet) When(s1 string, p1 interface{}, d1 time.Duration) *RedisMockSetExpectation {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("RedisMock.Set mock is already set by Set")
	}

	expectation := &RedisMockSetExpectation{
		mock:   mmSet.mock,
		params: &RedisMockSetParams{s1, p1, d1},
	}
	mmSet.expectations = append(mmSet.expectations, expectation)
	return expectation
}

// Then sets up Redis.Set return parameters for the expectation previously defined by the When method
func (e *RedisMockSetExpectation) Then(sp1 *redis.StatusCmd) *RedisMock {
	e.results = &RedisMockSetResults{sp1}
	return e.mock
}

// Set implements storage.Redis
func (mmSet *RedisMock) Set(s1 string, p1 interface{}, d1 time.Duration) (sp1 *redis.StatusCmd) {
	mm_atomic.AddUint64(&mmSet.beforeSetCounter, 1)
	defer mm_atomic.AddUint64(&mmSet.afterSetCounter, 1)

	if mmSet.inspectFuncSet != nil {
		mmSet.inspectFuncSet(s1, p1, d1)
	}

	mm_params := RedisMockSetParams{s1, p1, d1}

	// Record call args
	mmSet.SetMock.mutex.Lock()
	mmSet.SetMock.callArgs = append(mmSet.SetMock.callArgs, &mm_params)
	mmSet.SetMock.mutex.Unlock()

	for _, e := range mmSet.SetMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.sp1
		}
	}

	if mmSet.SetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSet.SetMock.defaultExpectation.Counter, 1)
		mm_want := mmSet.SetMock.defaultExpectation.params
		mm_got := RedisMockSetParams{s1, p1, d1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSet.t.Errorf("RedisMock.Set got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSet.SetMock.defaultExpectation.results
		if mm_results == nil {
			mmSet.t.Fatal("No results are set for the RedisMock.Set")
		}
		return (*mm_results).sp1
	}
	if mmSet.funcSet != nil {
		return mmSet.funcSet(s1, p1, d1)
	}
	mmSet.t.Fatalf("Unexpected call to RedisMock.Set. %v %v %v", s1, p1, d1)
	return
}

// SetAfterCounter returns a count of finished RedisMock.Set invocations
func (mmSet *RedisMock) SetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSet.afterSetCounter)
}

// SetBeforeCounter returns a count of RedisMock.Set invocations
func (mmSet *RedisMock) SetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSet.beforeSetCounter)
}

// Calls returns a list of arguments used in each call to RedisMock.Set.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSet *mRedisMockSet) Calls() []*RedisMockSetParams {
	mmSet.mutex.RLock()

	argCopy := make([]*RedisMockSetParams, len(mmSet.callArgs))
	copy(argCopy, mmSet.callArgs)

	mmSet.mutex.RUnlock()

	return argCopy
}

// MinimockSetDone returns true if the count of the Set invocations corresponds
// the number of defined expectations
func (m *RedisMock) MinimockSetDone() bool {
	for _, e := range m.SetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSet != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		return false
	}
	return true
}

// MinimockSetInspect logs each unmet expectation
func (m *RedisMock) MinimockSetInspect() {
	for _, e := range m.SetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RedisMock.Set with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		if m.SetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RedisMock.Set")
		} else {
			m.t.Errorf("Expected call to RedisMock.Set with params: %#v", *m.SetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSet != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		m.t.Error("Expected call to RedisMock.Set")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *RedisMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCloseInspect()

			m.MinimockGetInspect()

			m.MinimockPingInspect()

			m.MinimockSetInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *RedisMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *RedisMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCloseDone() &&
		m.MinimockGetDone() &&
		m.MinimockPingDone() &&
		m.MinimockSetDone()
}
