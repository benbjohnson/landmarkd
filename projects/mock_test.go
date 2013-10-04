package projects

import (
    "net/http"
    "github.com/skydb/sky.go"
    "github.com/stretchr/testify/mock"
    "time"
)

//--------------------------------------
// Redis Connection
//--------------------------------------

type mockRedisConn struct {
    mock.Mock
}

func (c *mockRedisConn) Close() error {
    args := c.Mock.Called()
    return args.Error(0)
}

func (c *mockRedisConn) Err() error {
    args := c.Mock.Called()
    return args.Error(0)
}

func (f *mockRedisConn) Do(commandName string, arg ...interface{}) (interface{}, error) {
    args := f.Mock.Called(commandName, arg)
    return args.Get(0), args.Error(1)
}

func (f *mockRedisConn) Send(commandName string, arg ...interface{}) error {
    args := f.Mock.Called(commandName, arg)
    return args.Error(0)
}

func (f *mockRedisConn) Flush() error {
    args := f.Mock.Called()
    return args.Error(0)
}

func (f *mockRedisConn) Receive() (interface{}, error) {
    args := f.Mock.Called()
    return args.Get(0), args.Error(1)
}


//--------------------------------------
// Sky Client
//--------------------------------------

type mockSkyClient struct {
    mock.Mock
}

func (c *mockSkyClient) GetTable(name string) (sky.Table, error) {
    args := c.Mock.Called(name)
    return args.Get(0).(sky.Table), args.Error(1)
}

func (c *mockSkyClient) GetTables() ([]sky.Table, error) {
    args := c.Mock.Called()
    return args.Get(0).([]sky.Table), args.Error(1)
}

func (c *mockSkyClient) CreateTable(table sky.Table) error {
    args := c.Mock.Called(table)
    return args.Error(0)
}

func (c *mockSkyClient) DeleteTable(table sky.Table) error {
    args := c.Mock.Called(table)
    return args.Error(0)
}

func (c *mockSkyClient) Ping() bool {
    args := c.Mock.Called()
    return args.Bool(0)
}

func (c *mockSkyClient) Send(method string, path string, data interface{}, ret interface{}) error {
    args := c.Mock.Called(method, path, data, ret)
    return args.Error(0)
}

func (c *mockSkyClient) URL(path string) string {
    args := c.Mock.Called(path)
    return args.String(0)
}

func (c *mockSkyClient) HTTPClient() *http.Client {
    args := c.Mock.Called()
    return args.Get(0).(*http.Client)
}


//--------------------------------------
// Sky Table
//--------------------------------------

type mockSkyTable struct {
    mock.Mock
}

func (t *mockSkyTable) Name() string {
    args := t.Mock.Called()
    return args.String(0)
}

func (t *mockSkyTable) Client() sky.Client {
    args := t.Mock.Called()
    return args.Get(0).(sky.Client)
}

func (t *mockSkyTable) SetClient(client sky.Client) {
    t.Mock.Called(client)
}

func (t *mockSkyTable) GetProperty(name string) (*sky.Property, error) {
    args := t.Mock.Called(name)
    return args.Get(0).(*sky.Property), args.Error(1)
}

func (t *mockSkyTable) GetProperties() ([]*sky.Property, error) {
    args := t.Mock.Called()
    return args.Get(0).([]*sky.Property), args.Error(1)
}

func (t *mockSkyTable) CreateProperty(property *sky.Property) error {
    args := t.Mock.Called(property)
    return args.Error(0)
}

func (t *mockSkyTable) UpdateProperty(name string, property *sky.Property) error {
    args := t.Mock.Called(name, property)
    return args.Error(0)
}

func (t *mockSkyTable) DeleteProperty(property *sky.Property) error {
    args := t.Mock.Called(property)
    return args.Error(0)
}

func (t *mockSkyTable) GetEvent(objectId string, timestamp time.Time) (*sky.Event, error) {
    args := t.Mock.Called(objectId, timestamp)
    return args.Get(0).(*sky.Event), args.Error(1)
}

func (t *mockSkyTable) GetEvents(objectId string) ([]*sky.Event, error) {
    args := t.Mock.Called(objectId)
    return args.Get(0).([]*sky.Event), args.Error(1)
}

func (t *mockSkyTable) AddEvent(objectId string, event *sky.Event, method string) error {
    args := t.Mock.Called(objectId, event, method)
    return args.Error(0)
}

func (t *mockSkyTable) DeleteEvent(objectId string, event *sky.Event) error {
    args := t.Mock.Called(objectId, event)
    return args.Error(0)
}

func (t *mockSkyTable) Stream(f func(*sky.EventStream)) error {
    args := t.Mock.Called(f)
    return args.Error(0)
}

func (t *mockSkyTable) Stats() (*sky.Stats, error) {
    args := t.Mock.Called()
    return args.Get(0).(*sky.Stats), args.Error(1)
}

func (t *mockSkyTable) RawQuery(q map[string]interface{}) (map[string]interface{}, error) {
    args := t.Mock.Called(q)
    return args.Get(0).(map[string]interface{}), args.Error(1)
}
