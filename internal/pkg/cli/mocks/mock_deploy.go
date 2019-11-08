// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/cli/deploy.go

// Package mocks is a generated GoMock package.
package mocks

import (
	archer "github.com/aws/amazon-ecs-cli-v2/internal/pkg/archer"
	deploy "github.com/aws/amazon-ecs-cli-v2/internal/pkg/deploy"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockenvironmentDeployer is a mock of environmentDeployer interface
type MockenvironmentDeployer struct {
	ctrl     *gomock.Controller
	recorder *MockenvironmentDeployerMockRecorder
}

// MockenvironmentDeployerMockRecorder is the mock recorder for MockenvironmentDeployer
type MockenvironmentDeployerMockRecorder struct {
	mock *MockenvironmentDeployer
}

// NewMockenvironmentDeployer creates a new mock instance
func NewMockenvironmentDeployer(ctrl *gomock.Controller) *MockenvironmentDeployer {
	mock := &MockenvironmentDeployer{ctrl: ctrl}
	mock.recorder = &MockenvironmentDeployerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockenvironmentDeployer) EXPECT() *MockenvironmentDeployerMockRecorder {
	return m.recorder
}

// DeployEnvironment mocks base method
func (m *MockenvironmentDeployer) DeployEnvironment(env *deploy.CreateEnvironmentInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeployEnvironment", env)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeployEnvironment indicates an expected call of DeployEnvironment
func (mr *MockenvironmentDeployerMockRecorder) DeployEnvironment(env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployEnvironment", reflect.TypeOf((*MockenvironmentDeployer)(nil).DeployEnvironment), env)
}

// StreamEnvironmentCreation mocks base method
func (m *MockenvironmentDeployer) StreamEnvironmentCreation(env *deploy.CreateEnvironmentInput) (<-chan []deploy.ResourceEvent, <-chan deploy.CreateEnvironmentResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StreamEnvironmentCreation", env)
	ret0, _ := ret[0].(<-chan []deploy.ResourceEvent)
	ret1, _ := ret[1].(<-chan deploy.CreateEnvironmentResponse)
	return ret0, ret1
}

// StreamEnvironmentCreation indicates an expected call of StreamEnvironmentCreation
func (mr *MockenvironmentDeployerMockRecorder) StreamEnvironmentCreation(env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamEnvironmentCreation", reflect.TypeOf((*MockenvironmentDeployer)(nil).StreamEnvironmentCreation), env)
}

// MockprojectDeployer is a mock of projectDeployer interface
type MockprojectDeployer struct {
	ctrl     *gomock.Controller
	recorder *MockprojectDeployerMockRecorder
}

// MockprojectDeployerMockRecorder is the mock recorder for MockprojectDeployer
type MockprojectDeployerMockRecorder struct {
	mock *MockprojectDeployer
}

// NewMockprojectDeployer creates a new mock instance
func NewMockprojectDeployer(ctrl *gomock.Controller) *MockprojectDeployer {
	mock := &MockprojectDeployer{ctrl: ctrl}
	mock.recorder = &MockprojectDeployerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockprojectDeployer) EXPECT() *MockprojectDeployerMockRecorder {
	return m.recorder
}

// DeployProject mocks base method
func (m *MockprojectDeployer) DeployProject(in *deploy.CreateProjectInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeployProject", in)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeployProject indicates an expected call of DeployProject
func (mr *MockprojectDeployerMockRecorder) DeployProject(in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployProject", reflect.TypeOf((*MockprojectDeployer)(nil).DeployProject), in)
}

// AddAppToProject mocks base method
func (m *MockprojectDeployer) AddAppToProject(project *archer.Project, app *archer.Application) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAppToProject", project, app)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAppToProject indicates an expected call of AddAppToProject
func (mr *MockprojectDeployerMockRecorder) AddAppToProject(project, app interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAppToProject", reflect.TypeOf((*MockprojectDeployer)(nil).AddAppToProject), project, app)
}

// AddEnvToProject mocks base method
func (m *MockprojectDeployer) AddEnvToProject(project *archer.Project, env *archer.Environment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEnvToProject", project, env)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEnvToProject indicates an expected call of AddEnvToProject
func (mr *MockprojectDeployerMockRecorder) AddEnvToProject(project, env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEnvToProject", reflect.TypeOf((*MockprojectDeployer)(nil).AddEnvToProject), project, env)
}

// Mockdeployer is a mock of deployer interface
type Mockdeployer struct {
	ctrl     *gomock.Controller
	recorder *MockdeployerMockRecorder
}

// MockdeployerMockRecorder is the mock recorder for Mockdeployer
type MockdeployerMockRecorder struct {
	mock *Mockdeployer
}

// NewMockdeployer creates a new mock instance
func NewMockdeployer(ctrl *gomock.Controller) *Mockdeployer {
	mock := &Mockdeployer{ctrl: ctrl}
	mock.recorder = &MockdeployerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mockdeployer) EXPECT() *MockdeployerMockRecorder {
	return m.recorder
}

// DeployEnvironment mocks base method
func (m *Mockdeployer) DeployEnvironment(env *deploy.CreateEnvironmentInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeployEnvironment", env)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeployEnvironment indicates an expected call of DeployEnvironment
func (mr *MockdeployerMockRecorder) DeployEnvironment(env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployEnvironment", reflect.TypeOf((*Mockdeployer)(nil).DeployEnvironment), env)
}

// StreamEnvironmentCreation mocks base method
func (m *Mockdeployer) StreamEnvironmentCreation(env *deploy.CreateEnvironmentInput) (<-chan []deploy.ResourceEvent, <-chan deploy.CreateEnvironmentResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StreamEnvironmentCreation", env)
	ret0, _ := ret[0].(<-chan []deploy.ResourceEvent)
	ret1, _ := ret[1].(<-chan deploy.CreateEnvironmentResponse)
	return ret0, ret1
}

// StreamEnvironmentCreation indicates an expected call of StreamEnvironmentCreation
func (mr *MockdeployerMockRecorder) StreamEnvironmentCreation(env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamEnvironmentCreation", reflect.TypeOf((*Mockdeployer)(nil).StreamEnvironmentCreation), env)
}

// DeployProject mocks base method
func (m *Mockdeployer) DeployProject(in *deploy.CreateProjectInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeployProject", in)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeployProject indicates an expected call of DeployProject
func (mr *MockdeployerMockRecorder) DeployProject(in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployProject", reflect.TypeOf((*Mockdeployer)(nil).DeployProject), in)
}

// AddAppToProject mocks base method
func (m *Mockdeployer) AddAppToProject(project *archer.Project, app *archer.Application) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAppToProject", project, app)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAppToProject indicates an expected call of AddAppToProject
func (mr *MockdeployerMockRecorder) AddAppToProject(project, app interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAppToProject", reflect.TypeOf((*Mockdeployer)(nil).AddAppToProject), project, app)
}

// AddEnvToProject mocks base method
func (m *Mockdeployer) AddEnvToProject(project *archer.Project, env *archer.Environment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEnvToProject", project, env)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEnvToProject indicates an expected call of AddEnvToProject
func (mr *MockdeployerMockRecorder) AddEnvToProject(project, env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEnvToProject", reflect.TypeOf((*Mockdeployer)(nil).AddEnvToProject), project, env)
}
