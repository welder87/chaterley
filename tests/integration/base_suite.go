// go:build integration

package integration_test

import (
	"database/sql"

	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/suite"
	_ "modernc.org/sqlite"
)

var migrationsDir = "../../internal/infrastructure/persistence/db/migrations"

type BaseIntegrationSuite struct {
	suite.Suite
	db *sql.DB
}

func (s *BaseIntegrationSuite) SetupSuite() {
	s.T().Log("Setup Suite")
	db, err := sql.Open("sqlite", ":memory:")
	s.Require().NoError(err)
	s.db = db
	s.applyMigrations()
}

func (s *BaseIntegrationSuite) TearDownSuite() {
	s.T().Log("Teardown Suite")
	if s.db == nil {
		return
	}
	s.unApplyMigrations()
	err := s.db.Close()
	s.Require().NoError(err)
}

func (s *BaseIntegrationSuite) applyMigrations() {
	s.T().Helper()
	goose.SetDialect("sqlite")
	err := goose.Up(s.db, migrationsDir)
	s.Require().NoError(err)
}

func (s *BaseIntegrationSuite) unApplyMigrations() {
	s.T().Helper()
	goose.SetDialect("sqlite")
	err := goose.Down(s.db, migrationsDir)
	s.Require().NoError(err)
}

// // Выполняется перед КАЖДЫМ тестом
func (s *BaseIntegrationSuite) SetupTest() {
	s.T().Log("Setup Test")
}

// // Выполняется после КАЖДОГО теста
func (s *BaseIntegrationSuite) TearDownTest() {
	s.T().Log("Teardown Test")
}

// // Сам тест
// func (s *IntegrationTestSuite) TestSomething() {
// 	resp, err := s.app.Get("/api/something")
// 	s.Require().NoError(err)
// 	s.Equal(200, resp.StatusCode)
// }

// // Функция запуска suite
// func TestIntegrationSuite(t *testing.T) {
// 	suite.Run(t, new(IntegrationTestSuite))
// }
