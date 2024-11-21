// Code generated by mockery. DO NOT EDIT.

package validation

import (
	validation "github.com/goravel/framework/contracts/validation"
	mock "github.com/stretchr/testify/mock"
)

// Validation is an autogenerated mock type for the Validation type
type Validation struct {
	mock.Mock
}

type Validation_Expecter struct {
	mock *mock.Mock
}

func (_m *Validation) EXPECT() *Validation_Expecter {
	return &Validation_Expecter{mock: &_m.Mock}
}

// AddFilters provides a mock function with given fields: _a0
func (_m *Validation) AddFilters(_a0 []validation.Filter) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for AddFilters")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]validation.Filter) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Validation_AddFilters_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddFilters'
type Validation_AddFilters_Call struct {
	*mock.Call
}

// AddFilters is a helper method to define mock.On call
//   - _a0 []validation.Filter
func (_e *Validation_Expecter) AddFilters(_a0 interface{}) *Validation_AddFilters_Call {
	return &Validation_AddFilters_Call{Call: _e.mock.On("AddFilters", _a0)}
}

func (_c *Validation_AddFilters_Call) Run(run func(_a0 []validation.Filter)) *Validation_AddFilters_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]validation.Filter))
	})
	return _c
}

func (_c *Validation_AddFilters_Call) Return(_a0 error) *Validation_AddFilters_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Validation_AddFilters_Call) RunAndReturn(run func([]validation.Filter) error) *Validation_AddFilters_Call {
	_c.Call.Return(run)
	return _c
}

// AddRules provides a mock function with given fields: _a0
func (_m *Validation) AddRules(_a0 []validation.Rule) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for AddRules")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]validation.Rule) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Validation_AddRules_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddRules'
type Validation_AddRules_Call struct {
	*mock.Call
}

// AddRules is a helper method to define mock.On call
//   - _a0 []validation.Rule
func (_e *Validation_Expecter) AddRules(_a0 interface{}) *Validation_AddRules_Call {
	return &Validation_AddRules_Call{Call: _e.mock.On("AddRules", _a0)}
}

func (_c *Validation_AddRules_Call) Run(run func(_a0 []validation.Rule)) *Validation_AddRules_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]validation.Rule))
	})
	return _c
}

func (_c *Validation_AddRules_Call) Return(_a0 error) *Validation_AddRules_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Validation_AddRules_Call) RunAndReturn(run func([]validation.Rule) error) *Validation_AddRules_Call {
	_c.Call.Return(run)
	return _c
}

// Filters provides a mock function with given fields:
func (_m *Validation) Filters() []validation.Filter {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Filters")
	}

	var r0 []validation.Filter
	if rf, ok := ret.Get(0).(func() []validation.Filter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]validation.Filter)
		}
	}

	return r0
}

// Validation_Filters_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Filters'
type Validation_Filters_Call struct {
	*mock.Call
}

// Filters is a helper method to define mock.On call
func (_e *Validation_Expecter) Filters() *Validation_Filters_Call {
	return &Validation_Filters_Call{Call: _e.mock.On("Filters")}
}

func (_c *Validation_Filters_Call) Run(run func()) *Validation_Filters_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Validation_Filters_Call) Return(_a0 []validation.Filter) *Validation_Filters_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Validation_Filters_Call) RunAndReturn(run func() []validation.Filter) *Validation_Filters_Call {
	_c.Call.Return(run)
	return _c
}

// Make provides a mock function with given fields: data, rules, options
func (_m *Validation) Make(data any, rules map[string]string, options ...validation.Option) (validation.Validator, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, data, rules)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Make")
	}

	var r0 validation.Validator
	var r1 error
	if rf, ok := ret.Get(0).(func(any, map[string]string, ...validation.Option) (validation.Validator, error)); ok {
		return rf(data, rules, options...)
	}
	if rf, ok := ret.Get(0).(func(any, map[string]string, ...validation.Option) validation.Validator); ok {
		r0 = rf(data, rules, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(validation.Validator)
		}
	}

	if rf, ok := ret.Get(1).(func(any, map[string]string, ...validation.Option) error); ok {
		r1 = rf(data, rules, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Validation_Make_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Make'
type Validation_Make_Call struct {
	*mock.Call
}

// Make is a helper method to define mock.On call
//   - data any
//   - rules map[string]string
//   - options ...validation.Option
func (_e *Validation_Expecter) Make(data interface{}, rules interface{}, options ...interface{}) *Validation_Make_Call {
	return &Validation_Make_Call{Call: _e.mock.On("Make",
		append([]interface{}{data, rules}, options...)...)}
}

func (_c *Validation_Make_Call) Run(run func(data any, rules map[string]string, options ...validation.Option)) *Validation_Make_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]validation.Option, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(validation.Option)
			}
		}
		run(args[0].(any), args[1].(map[string]string), variadicArgs...)
	})
	return _c
}

func (_c *Validation_Make_Call) Return(_a0 validation.Validator, _a1 error) *Validation_Make_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Validation_Make_Call) RunAndReturn(run func(any, map[string]string, ...validation.Option) (validation.Validator, error)) *Validation_Make_Call {
	_c.Call.Return(run)
	return _c
}

// Rules provides a mock function with given fields:
func (_m *Validation) Rules() []validation.Rule {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Rules")
	}

	var r0 []validation.Rule
	if rf, ok := ret.Get(0).(func() []validation.Rule); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]validation.Rule)
		}
	}

	return r0
}

// Validation_Rules_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Rules'
type Validation_Rules_Call struct {
	*mock.Call
}

// Rules is a helper method to define mock.On call
func (_e *Validation_Expecter) Rules() *Validation_Rules_Call {
	return &Validation_Rules_Call{Call: _e.mock.On("Rules")}
}

func (_c *Validation_Rules_Call) Run(run func()) *Validation_Rules_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Validation_Rules_Call) Return(_a0 []validation.Rule) *Validation_Rules_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Validation_Rules_Call) RunAndReturn(run func() []validation.Rule) *Validation_Rules_Call {
	_c.Call.Return(run)
	return _c
}

// NewValidation creates a new instance of Validation. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewValidation(t interface {
	mock.TestingT
	Cleanup(func())
}) *Validation {
	mock := &Validation{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}