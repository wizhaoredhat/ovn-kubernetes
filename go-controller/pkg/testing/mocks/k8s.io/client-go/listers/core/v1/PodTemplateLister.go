// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	labels "k8s.io/apimachinery/pkg/labels"
	corev1 "k8s.io/client-go/listers/core/v1"

	mock "github.com/stretchr/testify/mock"

	v1 "k8s.io/api/core/v1"
)

// PodTemplateLister is an autogenerated mock type for the PodTemplateLister type
type PodTemplateLister struct {
	mock.Mock
}

// List provides a mock function with given fields: selector
func (_m *PodTemplateLister) List(selector labels.Selector) ([]*v1.PodTemplate, error) {
	ret := _m.Called(selector)

	var r0 []*v1.PodTemplate
	if rf, ok := ret.Get(0).(func(labels.Selector) []*v1.PodTemplate); ok {
		r0 = rf(selector)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.PodTemplate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(labels.Selector) error); ok {
		r1 = rf(selector)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PodTemplates provides a mock function with given fields: namespace
func (_m *PodTemplateLister) PodTemplates(namespace string) corev1.PodTemplateNamespaceLister {
	ret := _m.Called(namespace)

	var r0 corev1.PodTemplateNamespaceLister
	if rf, ok := ret.Get(0).(func(string) corev1.PodTemplateNamespaceLister); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(corev1.PodTemplateNamespaceLister)
		}
	}

	return r0
}

type mockConstructorTestingTNewPodTemplateLister interface {
	mock.TestingT
	Cleanup(func())
}

// NewPodTemplateLister creates a new instance of PodTemplateLister. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPodTemplateLister(t mockConstructorTestingTNewPodTemplateLister) *PodTemplateLister {
	mock := &PodTemplateLister{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
