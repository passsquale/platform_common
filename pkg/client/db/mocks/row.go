// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/a1exCross/common/pkg/client/db.Row -o row.go -n RowMock -p mocks

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// RowMock implements db.Row
type RowMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcScan          func(dest ...interface{}) (err error)
	inspectFuncScan   func(dest ...interface{})
	afterScanCounter  uint64
	beforeScanCounter uint64
	ScanMock          mRowMockScan
}

// NewRowMock returns a mock for db.Row
func NewRowMock(t minimock.Tester) *RowMock {
	m := &RowMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ScanMock = mRowMockScan{mock: m}
	m.ScanMock.callArgs = []*RowMockScanParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mRowMockScan struct {
	mock               *RowMock
	defaultExpectation *RowMockScanExpectation
	expectations       []*RowMockScanExpectation

	callArgs []*RowMockScanParams
	mutex    sync.RWMutex
}

// RowMockScanExpectation specifies expectation struct of the Row.Scan
type RowMockScanExpectation struct {
	mock    *RowMock
	params  *RowMockScanParams
	results *RowMockScanResults
	Counter uint64
}

// RowMockScanParams contains parameters of the Row.Scan
type RowMockScanParams struct {
	dest []interface{}
}

// RowMockScanResults contains results of the Row.Scan
type RowMockScanResults struct {
	err error
}

// Expect sets up expected params for Row.Scan
func (mmScan *mRowMockScan) Expect(dest ...interface{}) *mRowMockScan {
	if mmScan.mock.funcScan != nil {
		mmScan.mock.t.Fatalf("RowMock.Scan mock is already set by Set")
	}

	if mmScan.defaultExpectation == nil {
		mmScan.defaultExpectation = &RowMockScanExpectation{}
	}

	mmScan.defaultExpectation.params = &RowMockScanParams{dest}
	for _, e := range mmScan.expectations {
		if minimock.Equal(e.params, mmScan.defaultExpectation.params) {
			mmScan.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmScan.defaultExpectation.params)
		}
	}

	return mmScan
}

// Inspect accepts an inspector function that has same arguments as the Row.Scan
func (mmScan *mRowMockScan) Inspect(f func(dest ...interface{})) *mRowMockScan {
	if mmScan.mock.inspectFuncScan != nil {
		mmScan.mock.t.Fatalf("Inspect function is already set for RowMock.Scan")
	}

	mmScan.mock.inspectFuncScan = f

	return mmScan
}

// Return sets up results that will be returned by Row.Scan
func (mmScan *mRowMockScan) Return(err error) *RowMock {
	if mmScan.mock.funcScan != nil {
		mmScan.mock.t.Fatalf("RowMock.Scan mock is already set by Set")
	}

	if mmScan.defaultExpectation == nil {
		mmScan.defaultExpectation = &RowMockScanExpectation{mock: mmScan.mock}
	}
	mmScan.defaultExpectation.results = &RowMockScanResults{err}
	return mmScan.mock
}

// Set uses given function f to mock the Row.Scan method
func (mmScan *mRowMockScan) Set(f func(dest ...interface{}) (err error)) *RowMock {
	if mmScan.defaultExpectation != nil {
		mmScan.mock.t.Fatalf("Default expectation is already set for the Row.Scan method")
	}

	if len(mmScan.expectations) > 0 {
		mmScan.mock.t.Fatalf("Some expectations are already set for the Row.Scan method")
	}

	mmScan.mock.funcScan = f
	return mmScan.mock
}

// When sets expectation for the Row.Scan which will trigger the result defined by the following
// Then helper
func (mmScan *mRowMockScan) When(dest ...interface{}) *RowMockScanExpectation {
	if mmScan.mock.funcScan != nil {
		mmScan.mock.t.Fatalf("RowMock.Scan mock is already set by Set")
	}

	expectation := &RowMockScanExpectation{
		mock:   mmScan.mock,
		params: &RowMockScanParams{dest},
	}
	mmScan.expectations = append(mmScan.expectations, expectation)
	return expectation
}

// Then sets up Row.Scan return parameters for the expectation previously defined by the When method
func (e *RowMockScanExpectation) Then(err error) *RowMock {
	e.results = &RowMockScanResults{err}
	return e.mock
}

// Scan implements db.Row
func (mmScan *RowMock) Scan(dest ...interface{}) (err error) {
	mm_atomic.AddUint64(&mmScan.beforeScanCounter, 1)
	defer mm_atomic.AddUint64(&mmScan.afterScanCounter, 1)

	if mmScan.inspectFuncScan != nil {
		mmScan.inspectFuncScan(dest...)
	}

	mm_params := RowMockScanParams{dest}

	// Record call args
	mmScan.ScanMock.mutex.Lock()
	mmScan.ScanMock.callArgs = append(mmScan.ScanMock.callArgs, &mm_params)
	mmScan.ScanMock.mutex.Unlock()

	for _, e := range mmScan.ScanMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmScan.ScanMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmScan.ScanMock.defaultExpectation.Counter, 1)
		mm_want := mmScan.ScanMock.defaultExpectation.params
		mm_got := RowMockScanParams{dest}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmScan.t.Errorf("RowMock.Scan got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmScan.ScanMock.defaultExpectation.results
		if mm_results == nil {
			mmScan.t.Fatal("No results are set for the RowMock.Scan")
		}
		return (*mm_results).err
	}
	if mmScan.funcScan != nil {
		return mmScan.funcScan(dest...)
	}
	mmScan.t.Fatalf("Unexpected call to RowMock.Scan. %v", dest)
	return
}

// ScanAfterCounter returns a count of finished RowMock.Scan invocations
func (mmScan *RowMock) ScanAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmScan.afterScanCounter)
}

// ScanBeforeCounter returns a count of RowMock.Scan invocations
func (mmScan *RowMock) ScanBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmScan.beforeScanCounter)
}

// Calls returns a list of arguments used in each call to RowMock.Scan.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmScan *mRowMockScan) Calls() []*RowMockScanParams {
	mmScan.mutex.RLock()

	argCopy := make([]*RowMockScanParams, len(mmScan.callArgs))
	copy(argCopy, mmScan.callArgs)

	mmScan.mutex.RUnlock()

	return argCopy
}

// MinimockScanDone returns true if the count of the Scan invocations corresponds
// the number of defined expectations
func (m *RowMock) MinimockScanDone() bool {
	for _, e := range m.ScanMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ScanMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterScanCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcScan != nil && mm_atomic.LoadUint64(&m.afterScanCounter) < 1 {
		return false
	}
	return true
}

// MinimockScanInspect logs each unmet expectation
func (m *RowMock) MinimockScanInspect() {
	for _, e := range m.ScanMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RowMock.Scan with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ScanMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterScanCounter) < 1 {
		if m.ScanMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RowMock.Scan")
		} else {
			m.t.Errorf("Expected call to RowMock.Scan with params: %#v", *m.ScanMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcScan != nil && mm_atomic.LoadUint64(&m.afterScanCounter) < 1 {
		m.t.Error("Expected call to RowMock.Scan")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *RowMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockScanInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *RowMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *RowMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockScanDone()
}