package events

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"sync"
	"testing"
	"time"
)
import "github.com/stretchr/testify/suite"

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event EventInterface, wg *sync.WaitGroup) {}

type EventDispatcherTestSuite struct {
	suite.Suite
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handler2        TestEventHandler
	handler3        TestEventHandler
	EventDispatcher *EventDispatcher
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.EventDispatcher = NewEventDispatcher()
	suite.handler = TestEventHandler{
		ID: 1,
	}
	suite.handler2 = TestEventHandler{
		ID: 2,
	}
	suite.handler3 = TestEventHandler{
		ID: 3,
	}
	suite.event = TestEvent{Name: "TestEvent", Payload: "TestPayload"}
	suite.event2 = TestEvent{Name: "TestEvent2", Payload: "TestPayload2"}
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	err = suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	err = suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(3, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	assert.Equal(suite.T(), &suite.handler, suite.EventDispatcher.handlers[suite.event.GetName()][0])
	assert.Equal(suite.T(), &suite.handler2, suite.EventDispatcher.handlers[suite.event.GetName()][1])
	assert.Equal(suite.T(), &suite.handler3, suite.EventDispatcher.handlers[suite.event.GetName()][2])
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_ErrHandlerAlreadyExists() {
	err := suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)

	err = suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Equal(ErrHandlerAlreadyExists, err)
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	//Event 1
	err := suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	err = suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	//Event 2
	err = suite.EventDispatcher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.EventDispatcher.handlers[suite.event2.GetName()]))

	suite.EventDispatcher.Clear()
	suite.Equal(0, len(suite.EventDispatcher.handlers))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	err := suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)

	err = suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)

	suite.True(suite.EventDispatcher.Has(suite.event.GetName(), &suite.handler))
	suite.True(suite.EventDispatcher.Has(suite.event.GetName(), &suite.handler2))
	suite.False(suite.EventDispatcher.Has(suite.event.GetName(), &suite.handler3))
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	m.Called(event)
	wg.Done()
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Dispatch() {
	eh := &MockHandler{}
	eh.On("Handle", &suite.event).Return(nil)

	eh2 := &MockHandler{}
	eh2.On("Handle", &suite.event).Return(nil)

	err := suite.EventDispatcher.Register(suite.event.GetName(), eh)
	suite.Nil(err)
	err = suite.EventDispatcher.Register(suite.event.GetName(), eh2)
	suite.Nil(err)

	suite.EventDispatcher.Dispatch(&suite.event)
	
	eh.AssertExpectations(suite.T())
	eh2.AssertExpectations(suite.T())

	eh.AssertNumberOfCalls(suite.T(), "Handle", 1)
	eh2.AssertNumberOfCalls(suite.T(), "Handle", 1)
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Remove() {
	err := suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	err = suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	suite.EventDispatcher.Remove(suite.event.GetName(), &suite.handler)
	suite.Equal(1, len(suite.EventDispatcher.handlers[suite.event.GetName()]))
	assert.Equal(suite.T(), &suite.handler2, suite.EventDispatcher.handlers[suite.event.GetName()][0])

	suite.EventDispatcher.Remove(suite.event.GetName(), &suite.handler2)
	suite.Equal(0, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	suite.EventDispatcher.Remove(suite.event.GetName(), &suite.handler3)
	suite.Equal(0, len(suite.EventDispatcher.handlers[suite.event.GetName()]))
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
