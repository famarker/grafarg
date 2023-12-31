// Code generated by MockGen. DO NOT EDIT.
// Source: gcsifaces.go

// Package mock_gcsifaces is a generated GoMock package.
package mock_gcsifaces

import (
	context "context"
	reflect "reflect"

	storage "cloud.google.com/go/storage"
	gomock "github.com/golang/mock/gomock"
	gcsifaces "github.com/famarker/grafarg/pkg/ifaces/gcsifaces"
	google "golang.org/x/oauth2/google"
	jwt "golang.org/x/oauth2/jwt"
)

// MockStorageClient is a mock of StorageClient interface
type MockStorageClient struct {
	ctrl     *gomock.Controller
	recorder *MockStorageClientMockRecorder
}

// MockStorageClientMockRecorder is the mock recorder for MockStorageClient
type MockStorageClientMockRecorder struct {
	mock *MockStorageClient
}

// NewMockStorageClient creates a new mock instance
func NewMockStorageClient(ctrl *gomock.Controller) *MockStorageClient {
	mock := &MockStorageClient{ctrl: ctrl}
	mock.recorder = &MockStorageClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorageClient) EXPECT() *MockStorageClientMockRecorder {
	return m.recorder
}

// Bucket mocks base method
func (m *MockStorageClient) Bucket(name string) gcsifaces.StorageBucket {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bucket", name)
	ret0, _ := ret[0].(gcsifaces.StorageBucket)
	return ret0
}

// Bucket indicates an expected call of Bucket
func (mr *MockStorageClientMockRecorder) Bucket(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bucket", reflect.TypeOf((*MockStorageClient)(nil).Bucket), name)
}

// FindDefaultCredentials mocks base method
func (m *MockStorageClient) FindDefaultCredentials(ctx context.Context, scope string) (*google.Credentials, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindDefaultCredentials", ctx, scope)
	ret0, _ := ret[0].(*google.Credentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindDefaultCredentials indicates an expected call of FindDefaultCredentials
func (mr *MockStorageClientMockRecorder) FindDefaultCredentials(ctx, scope interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindDefaultCredentials", reflect.TypeOf((*MockStorageClient)(nil).FindDefaultCredentials), ctx, scope)
}

// JWTConfigFromJSON mocks base method
func (m *MockStorageClient) JWTConfigFromJSON(keyJSON []byte) (*jwt.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JWTConfigFromJSON", keyJSON)
	ret0, _ := ret[0].(*jwt.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// JWTConfigFromJSON indicates an expected call of JWTConfigFromJSON
func (mr *MockStorageClientMockRecorder) JWTConfigFromJSON(keyJSON interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JWTConfigFromJSON", reflect.TypeOf((*MockStorageClient)(nil).JWTConfigFromJSON), keyJSON)
}

// SignedURL mocks base method
func (m *MockStorageClient) SignedURL(bucket, name string, opts *storage.SignedURLOptions) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignedURL", bucket, name, opts)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignedURL indicates an expected call of SignedURL
func (mr *MockStorageClientMockRecorder) SignedURL(bucket, name, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignedURL", reflect.TypeOf((*MockStorageClient)(nil).SignedURL), bucket, name, opts)
}

// MockStorageBucket is a mock of StorageBucket interface
type MockStorageBucket struct {
	ctrl     *gomock.Controller
	recorder *MockStorageBucketMockRecorder
}

// MockStorageBucketMockRecorder is the mock recorder for MockStorageBucket
type MockStorageBucketMockRecorder struct {
	mock *MockStorageBucket
}

// NewMockStorageBucket creates a new mock instance
func NewMockStorageBucket(ctrl *gomock.Controller) *MockStorageBucket {
	mock := &MockStorageBucket{ctrl: ctrl}
	mock.recorder = &MockStorageBucketMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorageBucket) EXPECT() *MockStorageBucketMockRecorder {
	return m.recorder
}

// Object mocks base method
func (m *MockStorageBucket) Object(key string) gcsifaces.StorageObject {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Object", key)
	ret0, _ := ret[0].(gcsifaces.StorageObject)
	return ret0
}

// Object indicates an expected call of Object
func (mr *MockStorageBucketMockRecorder) Object(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Object", reflect.TypeOf((*MockStorageBucket)(nil).Object), key)
}

// MockStorageObject is a mock of StorageObject interface
type MockStorageObject struct {
	ctrl     *gomock.Controller
	recorder *MockStorageObjectMockRecorder
}

// MockStorageObjectMockRecorder is the mock recorder for MockStorageObject
type MockStorageObjectMockRecorder struct {
	mock *MockStorageObject
}

// NewMockStorageObject creates a new mock instance
func NewMockStorageObject(ctrl *gomock.Controller) *MockStorageObject {
	mock := &MockStorageObject{ctrl: ctrl}
	mock.recorder = &MockStorageObjectMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorageObject) EXPECT() *MockStorageObjectMockRecorder {
	return m.recorder
}

// NewWriter mocks base method
func (m *MockStorageObject) NewWriter(ctx context.Context) gcsifaces.StorageWriter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewWriter", ctx)
	ret0, _ := ret[0].(gcsifaces.StorageWriter)
	return ret0
}

// NewWriter indicates an expected call of NewWriter
func (mr *MockStorageObjectMockRecorder) NewWriter(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewWriter", reflect.TypeOf((*MockStorageObject)(nil).NewWriter), ctx)
}

// MockStorageWriter is a mock of StorageWriter interface
type MockStorageWriter struct {
	ctrl     *gomock.Controller
	recorder *MockStorageWriterMockRecorder
}

// MockStorageWriterMockRecorder is the mock recorder for MockStorageWriter
type MockStorageWriterMockRecorder struct {
	mock *MockStorageWriter
}

// NewMockStorageWriter creates a new mock instance
func NewMockStorageWriter(ctrl *gomock.Controller) *MockStorageWriter {
	mock := &MockStorageWriter{ctrl: ctrl}
	mock.recorder = &MockStorageWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorageWriter) EXPECT() *MockStorageWriterMockRecorder {
	return m.recorder
}

// Write mocks base method
func (m *MockStorageWriter) Write(p []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockStorageWriterMockRecorder) Write(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockStorageWriter)(nil).Write), p)
}

// Close mocks base method
func (m *MockStorageWriter) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockStorageWriterMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStorageWriter)(nil).Close))
}

// SetACL mocks base method
func (m *MockStorageWriter) SetACL(acl string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetACL", acl)
}

// SetACL indicates an expected call of SetACL
func (mr *MockStorageWriterMockRecorder) SetACL(acl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetACL", reflect.TypeOf((*MockStorageWriter)(nil).SetACL), acl)
}
