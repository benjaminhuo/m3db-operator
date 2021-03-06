// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3db-operator/pkg/k8sops/types.go

// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package k8sops is a generated GoMock package.
package k8sops

import (
	"reflect"

	"github.com/m3db/m3db-operator/pkg/apis/m3dboperator/v1alpha1"

	"github.com/golang/mock/gomock"
	v1 "k8s.io/api/apps/v1"
	v10 "k8s.io/api/core/v1"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	v11 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v12 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
)

// MockK8sops is a mock of K8sops interface
type MockK8sops struct {
	ctrl     *gomock.Controller
	recorder *MockK8sopsMockRecorder
}

// MockK8sopsMockRecorder is the mock recorder for MockK8sops
type MockK8sopsMockRecorder struct {
	mock *MockK8sops
}

// NewMockK8sops creates a new mock instance
func NewMockK8sops(ctrl *gomock.Controller) *MockK8sops {
	mock := &MockK8sops{ctrl: ctrl}
	mock.recorder = &MockK8sopsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockK8sops) EXPECT() *MockK8sopsMockRecorder {
	return m.recorder
}

// ListM3DBCluster mocks base method
func (m *MockK8sops) ListM3DBCluster() (*v1alpha1.M3DBClusterList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListM3DBCluster")
	ret0, _ := ret[0].(*v1alpha1.M3DBClusterList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListM3DBCluster indicates an expected call of ListM3DBCluster
func (mr *MockK8sopsMockRecorder) ListM3DBCluster() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListM3DBCluster", reflect.TypeOf((*MockK8sops)(nil).ListM3DBCluster))
}

// GetM3DBCluster mocks base method
func (m *MockK8sops) GetM3DBCluster(namespace, name string) (*v1alpha1.M3DBCluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetM3DBCluster", namespace, name)
	ret0, _ := ret[0].(*v1alpha1.M3DBCluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetM3DBCluster indicates an expected call of GetM3DBCluster
func (mr *MockK8sopsMockRecorder) GetM3DBCluster(namespace, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetM3DBCluster", reflect.TypeOf((*MockK8sops)(nil).GetM3DBCluster), namespace, name)
}

// GetCRD mocks base method
func (m *MockK8sops) GetCRD(name string) (*v1beta1.CustomResourceDefinition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCRD", name)
	ret0, _ := ret[0].(*v1beta1.CustomResourceDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCRD indicates an expected call of GetCRD
func (mr *MockK8sopsMockRecorder) GetCRD(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCRD", reflect.TypeOf((*MockK8sops)(nil).GetCRD), name)
}

// CreateCRD mocks base method
func (m *MockK8sops) CreateCRD(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCRD", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCRD indicates an expected call of CreateCRD
func (mr *MockK8sopsMockRecorder) CreateCRD(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCRD", reflect.TypeOf((*MockK8sops)(nil).CreateCRD), name)
}

// GenerateCRD mocks base method
func (m *MockK8sops) GenerateCRD() *v1beta1.CustomResourceDefinition {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateCRD")
	ret0, _ := ret[0].(*v1beta1.CustomResourceDefinition)
	return ret0
}

// GenerateCRD indicates an expected call of GenerateCRD
func (mr *MockK8sopsMockRecorder) GenerateCRD() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateCRD", reflect.TypeOf((*MockK8sops)(nil).GenerateCRD))
}

// UpdateCRD mocks base method
func (m *MockK8sops) UpdateCRD(cluster *v1alpha1.M3DBCluster) (*v1alpha1.M3DBCluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCRD", cluster)
	ret0, _ := ret[0].(*v1alpha1.M3DBCluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCRD indicates an expected call of UpdateCRD
func (mr *MockK8sopsMockRecorder) UpdateCRD(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCRD", reflect.TypeOf((*MockK8sops)(nil).UpdateCRD), cluster)
}

// NewListWatcher mocks base method
func (m *MockK8sops) NewListWatcher() *cache.ListWatch {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListWatcher")
	ret0, _ := ret[0].(*cache.ListWatch)
	return ret0
}

// NewListWatcher indicates an expected call of NewListWatcher
func (mr *MockK8sopsMockRecorder) NewListWatcher() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListWatcher", reflect.TypeOf((*MockK8sops)(nil).NewListWatcher))
}

// GetService mocks base method
func (m *MockK8sops) GetService(cluster *v1alpha1.M3DBCluster, name string) (*v10.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService", cluster, name)
	ret0, _ := ret[0].(*v10.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService
func (mr *MockK8sopsMockRecorder) GetService(cluster, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockK8sops)(nil).GetService), cluster, name)
}

// DeleteService mocks base method
func (m *MockK8sops) DeleteService(cluster *v1alpha1.M3DBCluster, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteService", cluster, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteService indicates an expected call of DeleteService
func (mr *MockK8sopsMockRecorder) DeleteService(cluster, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteService", reflect.TypeOf((*MockK8sops)(nil).DeleteService), cluster, name)
}

// EnsureService mocks base method
func (m *MockK8sops) EnsureService(cluster *v1alpha1.M3DBCluster, svc *v10.Service) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureService", cluster, svc)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureService indicates an expected call of EnsureService
func (mr *MockK8sopsMockRecorder) EnsureService(cluster, svc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureService", reflect.TypeOf((*MockK8sops)(nil).EnsureService), cluster, svc)
}

// MultiLabelSelector mocks base method
func (m *MockK8sops) MultiLabelSelector(kvs map[string]string) v11.ListOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiLabelSelector", kvs)
	ret0, _ := ret[0].(v11.ListOptions)
	return ret0
}

// MultiLabelSelector indicates an expected call of MultiLabelSelector
func (mr *MockK8sopsMockRecorder) MultiLabelSelector(kvs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiLabelSelector", reflect.TypeOf((*MockK8sops)(nil).MultiLabelSelector), kvs)
}

// LabelSelector mocks base method
func (m *MockK8sops) LabelSelector(key, value string) v11.ListOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LabelSelector", key, value)
	ret0, _ := ret[0].(v11.ListOptions)
	return ret0
}

// LabelSelector indicates an expected call of LabelSelector
func (mr *MockK8sopsMockRecorder) LabelSelector(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LabelSelector", reflect.TypeOf((*MockK8sops)(nil).LabelSelector), key, value)
}

// DeleteStatefulSets mocks base method
func (m *MockK8sops) DeleteStatefulSets(cluster *v1alpha1.M3DBCluster, listOpts v11.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStatefulSets", cluster, listOpts)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStatefulSets indicates an expected call of DeleteStatefulSets
func (mr *MockK8sopsMockRecorder) DeleteStatefulSets(cluster, listOpts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStatefulSets", reflect.TypeOf((*MockK8sops)(nil).DeleteStatefulSets), cluster, listOpts)
}

// GetStatefulSets mocks base method
func (m *MockK8sops) GetStatefulSets(cluster *v1alpha1.M3DBCluster, listOpts v11.ListOptions) (*v1.StatefulSetList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatefulSets", cluster, listOpts)
	ret0, _ := ret[0].(*v1.StatefulSetList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStatefulSets indicates an expected call of GetStatefulSets
func (mr *MockK8sopsMockRecorder) GetStatefulSets(cluster, listOpts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatefulSets", reflect.TypeOf((*MockK8sops)(nil).GetStatefulSets), cluster, listOpts)
}

// GetPlacementDetails mocks base method
func (m *MockK8sops) GetPlacementDetails(cluster *v1alpha1.M3DBCluster) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlacementDetails", cluster)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlacementDetails indicates an expected call of GetPlacementDetails
func (mr *MockK8sopsMockRecorder) GetPlacementDetails(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlacementDetails", reflect.TypeOf((*MockK8sops)(nil).GetPlacementDetails), cluster)
}

// GetPodsByLabel mocks base method
func (m *MockK8sops) GetPodsByLabel(cluster *v1alpha1.M3DBCluster, listOpts v11.ListOptions) (*v10.PodList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodsByLabel", cluster, listOpts)
	ret0, _ := ret[0].(*v10.PodList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPodsByLabel indicates an expected call of GetPodsByLabel
func (mr *MockK8sopsMockRecorder) GetPodsByLabel(cluster, listOpts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodsByLabel", reflect.TypeOf((*MockK8sops)(nil).GetPodsByLabel), cluster, listOpts)
}

// CreateStatefulSet mocks base method
func (m *MockK8sops) CreateStatefulSet(cluster *v1alpha1.M3DBCluster, statefulSet *v1.StatefulSet) (*v1.StatefulSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStatefulSet", cluster, statefulSet)
	ret0, _ := ret[0].(*v1.StatefulSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStatefulSet indicates an expected call of CreateStatefulSet
func (mr *MockK8sopsMockRecorder) CreateStatefulSet(cluster, statefulSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStatefulSet", reflect.TypeOf((*MockK8sops)(nil).CreateStatefulSet), cluster, statefulSet)
}

// GetStatefulSet mocks base method
func (m *MockK8sops) GetStatefulSet(cluster *v1alpha1.M3DBCluster, name string) (*v1.StatefulSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatefulSet", cluster, name)
	ret0, _ := ret[0].(*v1.StatefulSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStatefulSet indicates an expected call of GetStatefulSet
func (mr *MockK8sopsMockRecorder) GetStatefulSet(cluster, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatefulSet", reflect.TypeOf((*MockK8sops)(nil).GetStatefulSet), cluster, name)
}

// UpdateStatefulSet mocks base method
func (m *MockK8sops) UpdateStatefulSet(cluster *v1alpha1.M3DBCluster, statefulSet *v1.StatefulSet) (*v1.StatefulSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatefulSet", cluster, statefulSet)
	ret0, _ := ret[0].(*v1.StatefulSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatefulSet indicates an expected call of UpdateStatefulSet
func (mr *MockK8sopsMockRecorder) UpdateStatefulSet(cluster, statefulSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatefulSet", reflect.TypeOf((*MockK8sops)(nil).UpdateStatefulSet), cluster, statefulSet)
}

// CheckStatefulStatus mocks base method
func (m *MockK8sops) CheckStatefulStatus(cluster *v1alpha1.M3DBCluster, statefulSet *v1.StatefulSet) (*v1.StatefulSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckStatefulStatus", cluster, statefulSet)
	ret0, _ := ret[0].(*v1.StatefulSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckStatefulStatus indicates an expected call of CheckStatefulStatus
func (mr *MockK8sopsMockRecorder) CheckStatefulStatus(cluster, statefulSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckStatefulStatus", reflect.TypeOf((*MockK8sops)(nil).CheckStatefulStatus), cluster, statefulSet)
}

// Events mocks base method
func (m *MockK8sops) Events(namespace string) v12.EventInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Events", namespace)
	ret0, _ := ret[0].(v12.EventInterface)
	return ret0
}

// Events indicates an expected call of Events
func (mr *MockK8sopsMockRecorder) Events(namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Events", reflect.TypeOf((*MockK8sops)(nil).Events), namespace)
}
