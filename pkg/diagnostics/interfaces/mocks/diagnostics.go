// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/diagnostics/interfaces.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	v1alpha1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	cluster "github.com/aws/eks-anywhere/pkg/cluster"
	diagnostics "github.com/aws/eks-anywhere/pkg/diagnostics"
	executables "github.com/aws/eks-anywhere/pkg/executables"
	providers "github.com/aws/eks-anywhere/pkg/providers"
	gomock "github.com/golang/mock/gomock"
)

// MockBundleClient is a mock of BundleClient interface.
type MockBundleClient struct {
	ctrl     *gomock.Controller
	recorder *MockBundleClientMockRecorder
}

// MockBundleClientMockRecorder is the mock recorder for MockBundleClient.
type MockBundleClientMockRecorder struct {
	mock *MockBundleClient
}

// NewMockBundleClient creates a new mock instance.
func NewMockBundleClient(ctrl *gomock.Controller) *MockBundleClient {
	mock := &MockBundleClient{ctrl: ctrl}
	mock.recorder = &MockBundleClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBundleClient) EXPECT() *MockBundleClientMockRecorder {
	return m.recorder
}

// Analyze mocks base method.
func (m *MockBundleClient) Analyze(ctx context.Context, bundleSpecPath, archivePath string) ([]*executables.SupportBundleAnalysis, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Analyze", ctx, bundleSpecPath, archivePath)
	ret0, _ := ret[0].([]*executables.SupportBundleAnalysis)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Analyze indicates an expected call of Analyze.
func (mr *MockBundleClientMockRecorder) Analyze(ctx, bundleSpecPath, archivePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Analyze", reflect.TypeOf((*MockBundleClient)(nil).Analyze), ctx, bundleSpecPath, archivePath)
}

// Collect mocks base method.
func (m *MockBundleClient) Collect(ctx context.Context, bundlePath string, sinceTime *time.Time, kubeconfig string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Collect", ctx, bundlePath, sinceTime, kubeconfig)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Collect indicates an expected call of Collect.
func (mr *MockBundleClientMockRecorder) Collect(ctx, bundlePath, sinceTime, kubeconfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Collect", reflect.TypeOf((*MockBundleClient)(nil).Collect), ctx, bundlePath, sinceTime, kubeconfig)
}

// MockDiagnosticBundleFactory is a mock of DiagnosticBundleFactory interface.
type MockDiagnosticBundleFactory struct {
	ctrl     *gomock.Controller
	recorder *MockDiagnosticBundleFactoryMockRecorder
}

// MockDiagnosticBundleFactoryMockRecorder is the mock recorder for MockDiagnosticBundleFactory.
type MockDiagnosticBundleFactoryMockRecorder struct {
	mock *MockDiagnosticBundleFactory
}

// NewMockDiagnosticBundleFactory creates a new mock instance.
func NewMockDiagnosticBundleFactory(ctrl *gomock.Controller) *MockDiagnosticBundleFactory {
	mock := &MockDiagnosticBundleFactory{ctrl: ctrl}
	mock.recorder = &MockDiagnosticBundleFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDiagnosticBundleFactory) EXPECT() *MockDiagnosticBundleFactoryMockRecorder {
	return m.recorder
}

// DiagnosticBundle mocks base method.
func (m *MockDiagnosticBundleFactory) DiagnosticBundle(spec *cluster.Spec, provider providers.Provider, kubeconfig, bundlePath string) (diagnostics.DiagnosticBundle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiagnosticBundle", spec, provider, kubeconfig, bundlePath)
	ret0, _ := ret[0].(diagnostics.DiagnosticBundle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DiagnosticBundle indicates an expected call of DiagnosticBundle.
func (mr *MockDiagnosticBundleFactoryMockRecorder) DiagnosticBundle(spec, provider, kubeconfig, bundlePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiagnosticBundle", reflect.TypeOf((*MockDiagnosticBundleFactory)(nil).DiagnosticBundle), spec, provider, kubeconfig, bundlePath)
}

// DiagnosticBundleCustom mocks base method.
func (m *MockDiagnosticBundleFactory) DiagnosticBundleCustom(kubeconfig, bundlePath string) diagnostics.DiagnosticBundle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiagnosticBundleCustom", kubeconfig, bundlePath)
	ret0, _ := ret[0].(diagnostics.DiagnosticBundle)
	return ret0
}

// DiagnosticBundleCustom indicates an expected call of DiagnosticBundleCustom.
func (mr *MockDiagnosticBundleFactoryMockRecorder) DiagnosticBundleCustom(kubeconfig, bundlePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiagnosticBundleCustom", reflect.TypeOf((*MockDiagnosticBundleFactory)(nil).DiagnosticBundleCustom), kubeconfig, bundlePath)
}

// DiagnosticBundleDefault mocks base method.
func (m *MockDiagnosticBundleFactory) DiagnosticBundleDefault() diagnostics.DiagnosticBundle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiagnosticBundleDefault")
	ret0, _ := ret[0].(diagnostics.DiagnosticBundle)
	return ret0
}

// DiagnosticBundleDefault indicates an expected call of DiagnosticBundleDefault.
func (mr *MockDiagnosticBundleFactoryMockRecorder) DiagnosticBundleDefault() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiagnosticBundleDefault", reflect.TypeOf((*MockDiagnosticBundleFactory)(nil).DiagnosticBundleDefault))
}

// DiagnosticBundleManagementCluster mocks base method.
func (m *MockDiagnosticBundleFactory) DiagnosticBundleManagementCluster(spec *cluster.Spec, kubeconfig string) (diagnostics.DiagnosticBundle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiagnosticBundleManagementCluster", spec, kubeconfig)
	ret0, _ := ret[0].(diagnostics.DiagnosticBundle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DiagnosticBundleManagementCluster indicates an expected call of DiagnosticBundleManagementCluster.
func (mr *MockDiagnosticBundleFactoryMockRecorder) DiagnosticBundleManagementCluster(spec, kubeconfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiagnosticBundleManagementCluster", reflect.TypeOf((*MockDiagnosticBundleFactory)(nil).DiagnosticBundleManagementCluster), spec, kubeconfig)
}

// DiagnosticBundleWorkloadCluster mocks base method.
func (m *MockDiagnosticBundleFactory) DiagnosticBundleWorkloadCluster(spec *cluster.Spec, provider providers.Provider, kubeconfig string) (diagnostics.DiagnosticBundle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiagnosticBundleWorkloadCluster", spec, provider, kubeconfig)
	ret0, _ := ret[0].(diagnostics.DiagnosticBundle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DiagnosticBundleWorkloadCluster indicates an expected call of DiagnosticBundleWorkloadCluster.
func (mr *MockDiagnosticBundleFactoryMockRecorder) DiagnosticBundleWorkloadCluster(spec, provider, kubeconfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiagnosticBundleWorkloadCluster", reflect.TypeOf((*MockDiagnosticBundleFactory)(nil).DiagnosticBundleWorkloadCluster), spec, provider, kubeconfig)
}

// MockDiagnosticBundle is a mock of DiagnosticBundle interface.
type MockDiagnosticBundle struct {
	ctrl     *gomock.Controller
	recorder *MockDiagnosticBundleMockRecorder
}

// MockDiagnosticBundleMockRecorder is the mock recorder for MockDiagnosticBundle.
type MockDiagnosticBundleMockRecorder struct {
	mock *MockDiagnosticBundle
}

// NewMockDiagnosticBundle creates a new mock instance.
func NewMockDiagnosticBundle(ctrl *gomock.Controller) *MockDiagnosticBundle {
	mock := &MockDiagnosticBundle{ctrl: ctrl}
	mock.recorder = &MockDiagnosticBundleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDiagnosticBundle) EXPECT() *MockDiagnosticBundleMockRecorder {
	return m.recorder
}

// CollectAndAnalyze mocks base method.
func (m *MockDiagnosticBundle) CollectAndAnalyze(ctx context.Context, sinceTimeValue *time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CollectAndAnalyze", ctx, sinceTimeValue)
	ret0, _ := ret[0].(error)
	return ret0
}

// CollectAndAnalyze indicates an expected call of CollectAndAnalyze.
func (mr *MockDiagnosticBundleMockRecorder) CollectAndAnalyze(ctx, sinceTimeValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CollectAndAnalyze", reflect.TypeOf((*MockDiagnosticBundle)(nil).CollectAndAnalyze), ctx, sinceTimeValue)
}

// PrintAnalysis mocks base method.
func (m *MockDiagnosticBundle) PrintAnalysis() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrintAnalysis")
	ret0, _ := ret[0].(error)
	return ret0
}

// PrintAnalysis indicates an expected call of PrintAnalysis.
func (mr *MockDiagnosticBundleMockRecorder) PrintAnalysis() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrintAnalysis", reflect.TypeOf((*MockDiagnosticBundle)(nil).PrintAnalysis))
}

// PrintBundleConfig mocks base method.
func (m *MockDiagnosticBundle) PrintBundleConfig() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrintBundleConfig")
	ret0, _ := ret[0].(error)
	return ret0
}

// PrintBundleConfig indicates an expected call of PrintBundleConfig.
func (mr *MockDiagnosticBundleMockRecorder) PrintBundleConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrintBundleConfig", reflect.TypeOf((*MockDiagnosticBundle)(nil).PrintBundleConfig))
}

// WithDatacenterConfig mocks base method.
func (m *MockDiagnosticBundle) WithDatacenterConfig(config v1alpha1.Ref) *diagnostics.EksaDiagnosticBundle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithDatacenterConfig", config)
	ret0, _ := ret[0].(*diagnostics.EksaDiagnosticBundle)
	return ret0
}

// WithDatacenterConfig indicates an expected call of WithDatacenterConfig.
func (mr *MockDiagnosticBundleMockRecorder) WithDatacenterConfig(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithDatacenterConfig", reflect.TypeOf((*MockDiagnosticBundle)(nil).WithDatacenterConfig), config)
}

// WithDefaultAnalyzers mocks base method.
func (m *MockDiagnosticBundle) WithDefaultAnalyzers() *diagnostics.EksaDiagnosticBundle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithDefaultAnalyzers")
	ret0, _ := ret[0].(*diagnostics.EksaDiagnosticBundle)
	return ret0
}

// WithDefaultAnalyzers indicates an expected call of WithDefaultAnalyzers.
func (mr *MockDiagnosticBundleMockRecorder) WithDefaultAnalyzers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithDefaultAnalyzers", reflect.TypeOf((*MockDiagnosticBundle)(nil).WithDefaultAnalyzers))
}

// WithDefaultCollectors mocks base method.
func (m *MockDiagnosticBundle) WithDefaultCollectors() *diagnostics.EksaDiagnosticBundle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithDefaultCollectors")
	ret0, _ := ret[0].(*diagnostics.EksaDiagnosticBundle)
	return ret0
}

// WithDefaultCollectors indicates an expected call of WithDefaultCollectors.
func (mr *MockDiagnosticBundleMockRecorder) WithDefaultCollectors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithDefaultCollectors", reflect.TypeOf((*MockDiagnosticBundle)(nil).WithDefaultCollectors))
}

// WithExternalEtcd mocks base method.
func (m *MockDiagnosticBundle) WithExternalEtcd(config *v1alpha1.ExternalEtcdConfiguration) *diagnostics.EksaDiagnosticBundle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithExternalEtcd", config)
	ret0, _ := ret[0].(*diagnostics.EksaDiagnosticBundle)
	return ret0
}

// WithExternalEtcd indicates an expected call of WithExternalEtcd.
func (mr *MockDiagnosticBundleMockRecorder) WithExternalEtcd(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithExternalEtcd", reflect.TypeOf((*MockDiagnosticBundle)(nil).WithExternalEtcd), config)
}

// WithGitOpsConfig mocks base method.
func (m *MockDiagnosticBundle) WithGitOpsConfig(config *v1alpha1.GitOpsConfig) *diagnostics.EksaDiagnosticBundle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithGitOpsConfig", config)
	ret0, _ := ret[0].(*diagnostics.EksaDiagnosticBundle)
	return ret0
}

// WithGitOpsConfig indicates an expected call of WithGitOpsConfig.
func (mr *MockDiagnosticBundleMockRecorder) WithGitOpsConfig(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithGitOpsConfig", reflect.TypeOf((*MockDiagnosticBundle)(nil).WithGitOpsConfig), config)
}

// WithLogTextAnalyzers mocks base method.
func (m *MockDiagnosticBundle) WithLogTextAnalyzers() *diagnostics.EksaDiagnosticBundle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithLogTextAnalyzers")
	ret0, _ := ret[0].(*diagnostics.EksaDiagnosticBundle)
	return ret0
}

// WithLogTextAnalyzers indicates an expected call of WithLogTextAnalyzers.
func (mr *MockDiagnosticBundleMockRecorder) WithLogTextAnalyzers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithLogTextAnalyzers", reflect.TypeOf((*MockDiagnosticBundle)(nil).WithLogTextAnalyzers))
}

// WithMachineConfigs mocks base method.
func (m *MockDiagnosticBundle) WithMachineConfigs(configs []providers.MachineConfig) *diagnostics.EksaDiagnosticBundle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithMachineConfigs", configs)
	ret0, _ := ret[0].(*diagnostics.EksaDiagnosticBundle)
	return ret0
}

// WithMachineConfigs indicates an expected call of WithMachineConfigs.
func (mr *MockDiagnosticBundleMockRecorder) WithMachineConfigs(configs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithMachineConfigs", reflect.TypeOf((*MockDiagnosticBundle)(nil).WithMachineConfigs), configs)
}

// WithOidcConfig mocks base method.
func (m *MockDiagnosticBundle) WithOidcConfig(config *v1alpha1.OIDCConfig) *diagnostics.EksaDiagnosticBundle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithOidcConfig", config)
	ret0, _ := ret[0].(*diagnostics.EksaDiagnosticBundle)
	return ret0
}

// WithOidcConfig indicates an expected call of WithOidcConfig.
func (mr *MockDiagnosticBundleMockRecorder) WithOidcConfig(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithOidcConfig", reflect.TypeOf((*MockDiagnosticBundle)(nil).WithOidcConfig), config)
}

// WriteAnalysisToFile mocks base method.
func (m *MockDiagnosticBundle) WriteAnalysisToFile() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteAnalysisToFile")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteAnalysisToFile indicates an expected call of WriteAnalysisToFile.
func (mr *MockDiagnosticBundleMockRecorder) WriteAnalysisToFile() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteAnalysisToFile", reflect.TypeOf((*MockDiagnosticBundle)(nil).WriteAnalysisToFile))
}

// WriteBundleConfig mocks base method.
func (m *MockDiagnosticBundle) WriteBundleConfig() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteBundleConfig")
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteBundleConfig indicates an expected call of WriteBundleConfig.
func (mr *MockDiagnosticBundleMockRecorder) WriteBundleConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteBundleConfig", reflect.TypeOf((*MockDiagnosticBundle)(nil).WriteBundleConfig))
}

// MockAnalyzerFactory is a mock of AnalyzerFactory interface.
type MockAnalyzerFactory struct {
	ctrl     *gomock.Controller
	recorder *MockAnalyzerFactoryMockRecorder
}

// MockAnalyzerFactoryMockRecorder is the mock recorder for MockAnalyzerFactory.
type MockAnalyzerFactoryMockRecorder struct {
	mock *MockAnalyzerFactory
}

// NewMockAnalyzerFactory creates a new mock instance.
func NewMockAnalyzerFactory(ctrl *gomock.Controller) *MockAnalyzerFactory {
	mock := &MockAnalyzerFactory{ctrl: ctrl}
	mock.recorder = &MockAnalyzerFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAnalyzerFactory) EXPECT() *MockAnalyzerFactoryMockRecorder {
	return m.recorder
}

// DataCenterConfigAnalyzers mocks base method.
func (m *MockAnalyzerFactory) DataCenterConfigAnalyzers(datacenter v1alpha1.Ref) []*diagnostics.Analyze {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataCenterConfigAnalyzers", datacenter)
	ret0, _ := ret[0].([]*diagnostics.Analyze)
	return ret0
}

// DataCenterConfigAnalyzers indicates an expected call of DataCenterConfigAnalyzers.
func (mr *MockAnalyzerFactoryMockRecorder) DataCenterConfigAnalyzers(datacenter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataCenterConfigAnalyzers", reflect.TypeOf((*MockAnalyzerFactory)(nil).DataCenterConfigAnalyzers), datacenter)
}

// DefaultAnalyzers mocks base method.
func (m *MockAnalyzerFactory) DefaultAnalyzers() []*diagnostics.Analyze {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultAnalyzers")
	ret0, _ := ret[0].([]*diagnostics.Analyze)
	return ret0
}

// DefaultAnalyzers indicates an expected call of DefaultAnalyzers.
func (mr *MockAnalyzerFactoryMockRecorder) DefaultAnalyzers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultAnalyzers", reflect.TypeOf((*MockAnalyzerFactory)(nil).DefaultAnalyzers))
}

// EksaExternalEtcdAnalyzers mocks base method.
func (m *MockAnalyzerFactory) EksaExternalEtcdAnalyzers() []*diagnostics.Analyze {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EksaExternalEtcdAnalyzers")
	ret0, _ := ret[0].([]*diagnostics.Analyze)
	return ret0
}

// EksaExternalEtcdAnalyzers indicates an expected call of EksaExternalEtcdAnalyzers.
func (mr *MockAnalyzerFactoryMockRecorder) EksaExternalEtcdAnalyzers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EksaExternalEtcdAnalyzers", reflect.TypeOf((*MockAnalyzerFactory)(nil).EksaExternalEtcdAnalyzers))
}

// EksaGitopsAnalyzers mocks base method.
func (m *MockAnalyzerFactory) EksaGitopsAnalyzers() []*diagnostics.Analyze {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EksaGitopsAnalyzers")
	ret0, _ := ret[0].([]*diagnostics.Analyze)
	return ret0
}

// EksaGitopsAnalyzers indicates an expected call of EksaGitopsAnalyzers.
func (mr *MockAnalyzerFactoryMockRecorder) EksaGitopsAnalyzers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EksaGitopsAnalyzers", reflect.TypeOf((*MockAnalyzerFactory)(nil).EksaGitopsAnalyzers))
}

// EksaLogTextAnalyzers mocks base method.
func (m *MockAnalyzerFactory) EksaLogTextAnalyzers(collectors []*diagnostics.Collect) []*diagnostics.Analyze {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EksaLogTextAnalyzers", collectors)
	ret0, _ := ret[0].([]*diagnostics.Analyze)
	return ret0
}

// EksaLogTextAnalyzers indicates an expected call of EksaLogTextAnalyzers.
func (mr *MockAnalyzerFactoryMockRecorder) EksaLogTextAnalyzers(collectors interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EksaLogTextAnalyzers", reflect.TypeOf((*MockAnalyzerFactory)(nil).EksaLogTextAnalyzers), collectors)
}

// EksaOidcAnalyzers mocks base method.
func (m *MockAnalyzerFactory) EksaOidcAnalyzers() []*diagnostics.Analyze {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EksaOidcAnalyzers")
	ret0, _ := ret[0].([]*diagnostics.Analyze)
	return ret0
}

// EksaOidcAnalyzers indicates an expected call of EksaOidcAnalyzers.
func (mr *MockAnalyzerFactoryMockRecorder) EksaOidcAnalyzers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EksaOidcAnalyzers", reflect.TypeOf((*MockAnalyzerFactory)(nil).EksaOidcAnalyzers))
}

// ManagementClusterAnalyzers mocks base method.
func (m *MockAnalyzerFactory) ManagementClusterAnalyzers() []*diagnostics.Analyze {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ManagementClusterAnalyzers")
	ret0, _ := ret[0].([]*diagnostics.Analyze)
	return ret0
}

// ManagementClusterAnalyzers indicates an expected call of ManagementClusterAnalyzers.
func (mr *MockAnalyzerFactoryMockRecorder) ManagementClusterAnalyzers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ManagementClusterAnalyzers", reflect.TypeOf((*MockAnalyzerFactory)(nil).ManagementClusterAnalyzers))
}

// PackageAnalyzers mocks base method.
func (m *MockAnalyzerFactory) PackageAnalyzers() []*diagnostics.Analyze {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PackageAnalyzers")
	ret0, _ := ret[0].([]*diagnostics.Analyze)
	return ret0
}

// PackageAnalyzers indicates an expected call of PackageAnalyzers.
func (mr *MockAnalyzerFactoryMockRecorder) PackageAnalyzers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PackageAnalyzers", reflect.TypeOf((*MockAnalyzerFactory)(nil).PackageAnalyzers))
}

// MockCollectorFactory is a mock of CollectorFactory interface.
type MockCollectorFactory struct {
	ctrl     *gomock.Controller
	recorder *MockCollectorFactoryMockRecorder
}

// MockCollectorFactoryMockRecorder is the mock recorder for MockCollectorFactory.
type MockCollectorFactoryMockRecorder struct {
	mock *MockCollectorFactory
}

// NewMockCollectorFactory creates a new mock instance.
func NewMockCollectorFactory(ctrl *gomock.Controller) *MockCollectorFactory {
	mock := &MockCollectorFactory{ctrl: ctrl}
	mock.recorder = &MockCollectorFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCollectorFactory) EXPECT() *MockCollectorFactoryMockRecorder {
	return m.recorder
}

// DataCenterConfigCollectors mocks base method.
func (m *MockCollectorFactory) DataCenterConfigCollectors(datacenter v1alpha1.Ref) []*diagnostics.Collect {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataCenterConfigCollectors", datacenter)
	ret0, _ := ret[0].([]*diagnostics.Collect)
	return ret0
}

// DataCenterConfigCollectors indicates an expected call of DataCenterConfigCollectors.
func (mr *MockCollectorFactoryMockRecorder) DataCenterConfigCollectors(datacenter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataCenterConfigCollectors", reflect.TypeOf((*MockCollectorFactory)(nil).DataCenterConfigCollectors), datacenter)
}

// DefaultCollectors mocks base method.
func (m *MockCollectorFactory) DefaultCollectors() []*diagnostics.Collect {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultCollectors")
	ret0, _ := ret[0].([]*diagnostics.Collect)
	return ret0
}

// DefaultCollectors indicates an expected call of DefaultCollectors.
func (mr *MockCollectorFactoryMockRecorder) DefaultCollectors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultCollectors", reflect.TypeOf((*MockCollectorFactory)(nil).DefaultCollectors))
}

// EksaHostCollectors mocks base method.
func (m *MockCollectorFactory) EksaHostCollectors(configs []providers.MachineConfig) []*diagnostics.Collect {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EksaHostCollectors", configs)
	ret0, _ := ret[0].([]*diagnostics.Collect)
	return ret0
}

// EksaHostCollectors indicates an expected call of EksaHostCollectors.
func (mr *MockCollectorFactoryMockRecorder) EksaHostCollectors(configs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EksaHostCollectors", reflect.TypeOf((*MockCollectorFactory)(nil).EksaHostCollectors), configs)
}

// ManagementClusterCollectors mocks base method.
func (m *MockCollectorFactory) ManagementClusterCollectors() []*diagnostics.Collect {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ManagementClusterCollectors")
	ret0, _ := ret[0].([]*diagnostics.Collect)
	return ret0
}

// ManagementClusterCollectors indicates an expected call of ManagementClusterCollectors.
func (mr *MockCollectorFactoryMockRecorder) ManagementClusterCollectors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ManagementClusterCollectors", reflect.TypeOf((*MockCollectorFactory)(nil).ManagementClusterCollectors))
}

// PackagesCollectors mocks base method.
func (m *MockCollectorFactory) PackagesCollectors() []*diagnostics.Collect {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PackagesCollectors")
	ret0, _ := ret[0].([]*diagnostics.Collect)
	return ret0
}

// PackagesCollectors indicates an expected call of PackagesCollectors.
func (mr *MockCollectorFactoryMockRecorder) PackagesCollectors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PackagesCollectors", reflect.TypeOf((*MockCollectorFactory)(nil).PackagesCollectors))
}
