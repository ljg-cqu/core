package postgres_notice

import (
	"github.com/ljg-cqu/core/logger"
	"github.com/ljg-cqu/core/postgres"
	"github.com/stretchr/testify/suite"
	"testing"
)

type NoticerTestSuite struct {
	*suite.Suite
}

func (n *NoticerTestSuite) SetupTest() {

}

func (s *NoticerTestSuite) TearDownTest() {

}

func TestNoticerTestSuite(t *testing.T) {
	suite.Run(t, new(NoticerTestSuite))
}

func (s *NoticerTestSuite) TestNewNoticer() {
	log := logger.New()
	pool := postgres.PgxPool(postgres.TestDBAliConnStr)
	n := NewNoticer(log, pool, "chat")
	s.Require().NotNil(n)
}
