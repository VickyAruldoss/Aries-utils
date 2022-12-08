package logging

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
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
	actual := 5
	suite.Equal(expected, actual)
}
