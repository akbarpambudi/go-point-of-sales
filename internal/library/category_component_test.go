package library_test

import (
	"context"
	"github.com/akbarpambudi/go-point-of-sales/internal/common/testinghelper"
	"github.com/akbarpambudi/go-point-of-sales/internal/library"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/adapter/adapterent/ent"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/adapter/adapterent/ent/enttest"
	"github.com/google/uuid"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type CategoryComponentTestSuite struct {
	suite.Suite
	stopService func()
	ent         *ent.Client
	service     http.Handler
}

func (s *CategoryComponentTestSuite) SetupSuite() {
	s.ent = enttest.Open(s.T(), "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	service, stopService, err := library.NewWebService(context.TODO(), library.SetDataSourceClient(s.ent))
	if err != nil {
		s.T().Fatal(err)
	}
	s.service = service
	s.stopService = stopService
	s.setupDataSample()
}

func (s *CategoryComponentTestSuite) TearDownSuite() {
	_ = s.ent.Close()
	s.stopService()
}

func (s CategoryComponentTestSuite) TestCreateCategoryShouldBeSuccess() {
	apitest.New("Test Create Category").
		Handler(s.service).
		Post("/api/category").
		JSON(testinghelper.JSONDictionary{
			"id":   "f1e1a9fa-690f-4756-907a-1d85bc044391",
			"name": "Dessert",
		}).Expect(s.T()).
		Status(http.StatusCreated).
		End()
}

func (s *CategoryComponentTestSuite) setupDataSample() {
	ctx := context.TODO()
	dataSamples := []struct {
		id   string
		name string
	}{
		{
			id:   "e88d1f9b-d7b7-437a-9c6a-2a80291c2427",
			name: "Main Course",
		},
		{
			id:   "3965e166-4be0-46cb-ae6c-db4f467f5815",
			name: "Appetizer",
		},
	}

	for _, r := range dataSamples {
		s.ent.Category.Create().SetID(uuid.MustParse(r.id)).SetName(r.name).SaveX(ctx)
	}
}

func TestRunCategoryComponentTestSuite(t *testing.T) {
	suite.Run(t, new(CategoryComponentTestSuite))
}
