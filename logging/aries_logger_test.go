package logging

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

type AriesLoggerTestSuite struct {
	suite.Suite
	context  *gin.Context
	recorder *httptest.ResponseRecorder
}

func TestAriesLoggerTestSuite(t *testing.T) {
	suite.Run(t, new(AriesLoggerTestSuite))
}

func (suite *AriesLoggerTestSuite) SetupTest() {
	suite.recorder = httptest.NewRecorder()
	suite.context, _ = gin.CreateTestContext(suite.recorder)
}

func (suite *AriesLoggerTestSuite) TestForTestSuite() {
	expected := 5
	actual := 4
	suite.Equal(expected, actual)
}

func (suite *AriesLoggerTestSuite) TestShouldCreateLogggerEntry() {
	logursEntry := logrus.StandardLogger().WithContext(context.Background())
	loggerEntry := newAriesLoggerEntry(context.Background(), logursEntry)

	suite.Equal(context.Background(), loggerEntry.context)
	suite.Equal(logursEntry, loggerEntry.stdEntry)
	suite.Equal(0, len(loggerEntry.Data))
	suite.Equal("", loggerEntry.err)
}

func (suite *AriesLoggerTestSuite) TestShouldCreatePublicLoggerEntry() {
	logursEntry := logrus.StandardLogger().WithContext(context.Background())
	loggerEntry := NewloggerEntry()

	suite.Equal(context.Background(), loggerEntry.context)
	suite.Equal(logursEntry, loggerEntry.stdEntry)
	suite.Equal(0, len(loggerEntry.Data))
	suite.Equal("", loggerEntry.err)
}
