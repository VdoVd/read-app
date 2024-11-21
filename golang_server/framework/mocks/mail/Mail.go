// Code generated by mockery. DO NOT EDIT.

package mail

import (
	mail "github.com/goravel/framework/contracts/mail"
	mock "github.com/stretchr/testify/mock"
)

// Mail is an autogenerated mock type for the Mail type
type Mail struct {
	mock.Mock
}

type Mail_Expecter struct {
	mock *mock.Mock
}

func (_m *Mail) EXPECT() *Mail_Expecter {
	return &Mail_Expecter{mock: &_m.Mock}
}

// Attach provides a mock function with given fields: files
func (_m *Mail) Attach(files []string) mail.Mail {
	ret := _m.Called(files)

	if len(ret) == 0 {
		panic("no return value specified for Attach")
	}

	var r0 mail.Mail
	if rf, ok := ret.Get(0).(func([]string) mail.Mail); ok {
		r0 = rf(files)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mail.Mail)
		}
	}

	return r0
}

// Mail_Attach_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Attach'
type Mail_Attach_Call struct {
	*mock.Call
}

// Attach is a helper method to define mock.On call
//   - files []string
func (_e *Mail_Expecter) Attach(files interface{}) *Mail_Attach_Call {
	return &Mail_Attach_Call{Call: _e.mock.On("Attach", files)}
}

func (_c *Mail_Attach_Call) Run(run func(files []string)) *Mail_Attach_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *Mail_Attach_Call) Return(_a0 mail.Mail) *Mail_Attach_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Mail_Attach_Call) RunAndReturn(run func([]string) mail.Mail) *Mail_Attach_Call {
	_c.Call.Return(run)
	return _c
}

// Bcc provides a mock function with given fields: addresses
func (_m *Mail) Bcc(addresses []string) mail.Mail {
	ret := _m.Called(addresses)

	if len(ret) == 0 {
		panic("no return value specified for Bcc")
	}

	var r0 mail.Mail
	if rf, ok := ret.Get(0).(func([]string) mail.Mail); ok {
		r0 = rf(addresses)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mail.Mail)
		}
	}

	return r0
}

// Mail_Bcc_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Bcc'
type Mail_Bcc_Call struct {
	*mock.Call
}

// Bcc is a helper method to define mock.On call
//   - addresses []string
func (_e *Mail_Expecter) Bcc(addresses interface{}) *Mail_Bcc_Call {
	return &Mail_Bcc_Call{Call: _e.mock.On("Bcc", addresses)}
}

func (_c *Mail_Bcc_Call) Run(run func(addresses []string)) *Mail_Bcc_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *Mail_Bcc_Call) Return(_a0 mail.Mail) *Mail_Bcc_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Mail_Bcc_Call) RunAndReturn(run func([]string) mail.Mail) *Mail_Bcc_Call {
	_c.Call.Return(run)
	return _c
}

// Cc provides a mock function with given fields: addresses
func (_m *Mail) Cc(addresses []string) mail.Mail {
	ret := _m.Called(addresses)

	if len(ret) == 0 {
		panic("no return value specified for Cc")
	}

	var r0 mail.Mail
	if rf, ok := ret.Get(0).(func([]string) mail.Mail); ok {
		r0 = rf(addresses)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mail.Mail)
		}
	}

	return r0
}

// Mail_Cc_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Cc'
type Mail_Cc_Call struct {
	*mock.Call
}

// Cc is a helper method to define mock.On call
//   - addresses []string
func (_e *Mail_Expecter) Cc(addresses interface{}) *Mail_Cc_Call {
	return &Mail_Cc_Call{Call: _e.mock.On("Cc", addresses)}
}

func (_c *Mail_Cc_Call) Run(run func(addresses []string)) *Mail_Cc_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *Mail_Cc_Call) Return(_a0 mail.Mail) *Mail_Cc_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Mail_Cc_Call) RunAndReturn(run func([]string) mail.Mail) *Mail_Cc_Call {
	_c.Call.Return(run)
	return _c
}

// Content provides a mock function with given fields: content
func (_m *Mail) Content(content mail.Content) mail.Mail {
	ret := _m.Called(content)

	if len(ret) == 0 {
		panic("no return value specified for Content")
	}

	var r0 mail.Mail
	if rf, ok := ret.Get(0).(func(mail.Content) mail.Mail); ok {
		r0 = rf(content)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mail.Mail)
		}
	}

	return r0
}

// Mail_Content_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Content'
type Mail_Content_Call struct {
	*mock.Call
}

// Content is a helper method to define mock.On call
//   - content mail.Content
func (_e *Mail_Expecter) Content(content interface{}) *Mail_Content_Call {
	return &Mail_Content_Call{Call: _e.mock.On("Content", content)}
}

func (_c *Mail_Content_Call) Run(run func(content mail.Content)) *Mail_Content_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(mail.Content))
	})
	return _c
}

func (_c *Mail_Content_Call) Return(_a0 mail.Mail) *Mail_Content_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Mail_Content_Call) RunAndReturn(run func(mail.Content) mail.Mail) *Mail_Content_Call {
	_c.Call.Return(run)
	return _c
}

// From provides a mock function with given fields: address
func (_m *Mail) From(address mail.Address) mail.Mail {
	ret := _m.Called(address)

	if len(ret) == 0 {
		panic("no return value specified for From")
	}

	var r0 mail.Mail
	if rf, ok := ret.Get(0).(func(mail.Address) mail.Mail); ok {
		r0 = rf(address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mail.Mail)
		}
	}

	return r0
}

// Mail_From_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'From'
type Mail_From_Call struct {
	*mock.Call
}

// From is a helper method to define mock.On call
//   - address mail.Address
func (_e *Mail_Expecter) From(address interface{}) *Mail_From_Call {
	return &Mail_From_Call{Call: _e.mock.On("From", address)}
}

func (_c *Mail_From_Call) Run(run func(address mail.Address)) *Mail_From_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(mail.Address))
	})
	return _c
}

func (_c *Mail_From_Call) Return(_a0 mail.Mail) *Mail_From_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Mail_From_Call) RunAndReturn(run func(mail.Address) mail.Mail) *Mail_From_Call {
	_c.Call.Return(run)
	return _c
}

// Queue provides a mock function with given fields: mailable
func (_m *Mail) Queue(mailable ...mail.Mailable) error {
	_va := make([]interface{}, len(mailable))
	for _i := range mailable {
		_va[_i] = mailable[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Queue")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(...mail.Mailable) error); ok {
		r0 = rf(mailable...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Mail_Queue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Queue'
type Mail_Queue_Call struct {
	*mock.Call
}

// Queue is a helper method to define mock.On call
//   - mailable ...mail.Mailable
func (_e *Mail_Expecter) Queue(mailable ...interface{}) *Mail_Queue_Call {
	return &Mail_Queue_Call{Call: _e.mock.On("Queue",
		append([]interface{}{}, mailable...)...)}
}

func (_c *Mail_Queue_Call) Run(run func(mailable ...mail.Mailable)) *Mail_Queue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]mail.Mailable, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(mail.Mailable)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Mail_Queue_Call) Return(_a0 error) *Mail_Queue_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Mail_Queue_Call) RunAndReturn(run func(...mail.Mailable) error) *Mail_Queue_Call {
	_c.Call.Return(run)
	return _c
}

// Send provides a mock function with given fields: mailable
func (_m *Mail) Send(mailable ...mail.Mailable) error {
	_va := make([]interface{}, len(mailable))
	for _i := range mailable {
		_va[_i] = mailable[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Send")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(...mail.Mailable) error); ok {
		r0 = rf(mailable...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Mail_Send_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Send'
type Mail_Send_Call struct {
	*mock.Call
}

// Send is a helper method to define mock.On call
//   - mailable ...mail.Mailable
func (_e *Mail_Expecter) Send(mailable ...interface{}) *Mail_Send_Call {
	return &Mail_Send_Call{Call: _e.mock.On("Send",
		append([]interface{}{}, mailable...)...)}
}

func (_c *Mail_Send_Call) Run(run func(mailable ...mail.Mailable)) *Mail_Send_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]mail.Mailable, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(mail.Mailable)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Mail_Send_Call) Return(_a0 error) *Mail_Send_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Mail_Send_Call) RunAndReturn(run func(...mail.Mailable) error) *Mail_Send_Call {
	_c.Call.Return(run)
	return _c
}

// Subject provides a mock function with given fields: subject
func (_m *Mail) Subject(subject string) mail.Mail {
	ret := _m.Called(subject)

	if len(ret) == 0 {
		panic("no return value specified for Subject")
	}

	var r0 mail.Mail
	if rf, ok := ret.Get(0).(func(string) mail.Mail); ok {
		r0 = rf(subject)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mail.Mail)
		}
	}

	return r0
}

// Mail_Subject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Subject'
type Mail_Subject_Call struct {
	*mock.Call
}

// Subject is a helper method to define mock.On call
//   - subject string
func (_e *Mail_Expecter) Subject(subject interface{}) *Mail_Subject_Call {
	return &Mail_Subject_Call{Call: _e.mock.On("Subject", subject)}
}

func (_c *Mail_Subject_Call) Run(run func(subject string)) *Mail_Subject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Mail_Subject_Call) Return(_a0 mail.Mail) *Mail_Subject_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Mail_Subject_Call) RunAndReturn(run func(string) mail.Mail) *Mail_Subject_Call {
	_c.Call.Return(run)
	return _c
}

// To provides a mock function with given fields: addresses
func (_m *Mail) To(addresses []string) mail.Mail {
	ret := _m.Called(addresses)

	if len(ret) == 0 {
		panic("no return value specified for To")
	}

	var r0 mail.Mail
	if rf, ok := ret.Get(0).(func([]string) mail.Mail); ok {
		r0 = rf(addresses)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mail.Mail)
		}
	}

	return r0
}

// Mail_To_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'To'
type Mail_To_Call struct {
	*mock.Call
}

// To is a helper method to define mock.On call
//   - addresses []string
func (_e *Mail_Expecter) To(addresses interface{}) *Mail_To_Call {
	return &Mail_To_Call{Call: _e.mock.On("To", addresses)}
}

func (_c *Mail_To_Call) Run(run func(addresses []string)) *Mail_To_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *Mail_To_Call) Return(_a0 mail.Mail) *Mail_To_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Mail_To_Call) RunAndReturn(run func([]string) mail.Mail) *Mail_To_Call {
	_c.Call.Return(run)
	return _c
}

// NewMail creates a new instance of Mail. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMail(t interface {
	mock.TestingT
	Cleanup(func())
}) *Mail {
	mock := &Mail{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}