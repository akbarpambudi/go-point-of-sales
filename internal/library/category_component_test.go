//+build component_test

package library_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/akbarpambudi/go-point-of-sales/internal/common/errors"
	"github.com/akbarpambudi/go-point-of-sales/internal/common/testinghelper"
	"github.com/akbarpambudi/go-point-of-sales/internal/library"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/adapter/adapterent/ent"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/adapter/adapterent/ent/enttest"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/domain/category"
	"github.com/google/uuid"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/stretchr/testify/suite"
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
			"name": "Beverage",
		}).Expect(s.T()).
		Status(http.StatusCreated).
		End()
}

func (s CategoryComponentTestSuite) TestCreateCategoryShouldResponseWithBadRequest() {
	type expectedResponse struct {
		responseBodyMatcher func(*http.Response, *http.Request) error
	}
	testTable := []struct {
		requestBody JSONDictionary
		want        expectedResponse
	}{
		{
			requestBody: JSONDictionary{
				"name": "Dessert",
			},
			want: expectedResponse{
				responseBodyMatcher: jsonpath.
					Chain().
					Equal("$.message", "invalid input data").
					Equal("$.messageKey", category.CreationErrKey).
					Equal("$.errorType", errors.ErrorTypeIllegalInputError).
					Contains("$.children[*]", JSONDictionary{
						"message":    category.ErrCategoryIDCantBeEmpty.(*errors.POSError).Message(),
						"messageKey": category.ErrCategoryIDCantBeEmpty.(*errors.POSError).Key(),
						"errorType":  category.ErrCategoryIDCantBeEmpty.(*errors.POSError).ErrType(),
					}).
					End(),
			},
		},
		{
			requestBody: JSONDictionary{
				"id": "f1e1a9fa-690f-4756-907a-1d85bc044391",
			},
			want: expectedResponse{
				responseBodyMatcher: jsonpath.
					Chain().
					Equal("$.message", "invalid input data").
					Equal("$.messageKey", category.CreationErrKey).
					Equal("$.errorType", errors.ErrorTypeIllegalInputError).
					Contains("$.children[*]", JSONDictionary{
						"message":    category.ErrCategoryNameCantBeEmpty.(*errors.POSError).Message(),
						"messageKey": category.ErrCategoryNameCantBeEmpty.(*errors.POSError).Key(),
						"errorType":  category.ErrCategoryNameCantBeEmpty.(*errors.POSError).ErrType(),
					}).
					End(),
			},
		},
		{
			requestBody: JSONDictionary{},
			want: expectedResponse{
				responseBodyMatcher: jsonpath.
					Chain().
					Equal("$.message", "invalid input data").
					Equal("$.messageKey", category.CreationErrKey).
					Equal("$.errorType", errors.ErrorTypeIllegalInputError).
					Contains("$.children[*]", JSONDictionary{
						"message":    category.ErrCategoryNameCantBeEmpty.(*errors.POSError).Message(),
						"messageKey": category.ErrCategoryNameCantBeEmpty.(*errors.POSError).Key(),
						"errorType":  category.ErrCategoryNameCantBeEmpty.(*errors.POSError).ErrType(),
					}).
					Contains("$.children[*]", JSONDictionary{
						"message":    category.ErrCategoryIDCantBeEmpty.(*errors.POSError).Message(),
						"messageKey": category.ErrCategoryIDCantBeEmpty.(*errors.POSError).Key(),
						"errorType":  category.ErrCategoryIDCantBeEmpty.(*errors.POSError).ErrType(),
					}).
					End(),
			},
		},
	}

	for i, r := range testTable {
		testCaseName := fmt.Sprintf("TestCase#%d", i)
		s.Run(testCaseName, func() {
			apitest.New("Test Create Category").
				Handler(s.service).
				Post("/api/category").
				JSON(r.requestBody).Expect(s.T()).
				Status(http.StatusBadRequest).
				Assert(r.want.responseBodyMatcher).
				End()
		})
	}
}

func (s CategoryComponentTestSuite) TestGetCategoryByIDShouldBeSuccess() {
	apitest.New("Test Get Category by ID").
		Handler(s.service).
		Get("/api/category/e88d1f9b-d7b7-437a-9c6a-2a80291c2427").
		Expect(s.T()).
		Status(http.StatusOK).
		Assert(jsonpath.Equal("$.id", "e88d1f9b-d7b7-437a-9c6a-2a80291c2427")).
		Assert(jsonpath.Equal("$.name", "Main Course")).
		End()
}

func (s CategoryComponentTestSuite) TestGetCategoryByIDShouldReturnStatusNotFound() {
	apitest.New("Test Get Category by ID").
		Handler(s.service).
		Get("/api/category/bd3fcd27-40b7-493d-98b1-22d031c97960").
		Expect(s.T()).
		Status(http.StatusNotFound).
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
