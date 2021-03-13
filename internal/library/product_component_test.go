package library_test

import (
	"context"
	"fmt"
	"github.com/akbarpambudi/go-point-of-sales/internal/common/errors"
	"github.com/akbarpambudi/go-point-of-sales/internal/library"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/domain/product"
	"github.com/google/uuid"
	"github.com/steinfletcher/apitest"
	"github.com/steinfletcher/apitest-jsonpath"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type JSONDictionary = map[string]interface{}

type ProductComponentTestSuite struct {
	suite.Suite
	stopService func()
	service     http.Handler
}

func (s *ProductComponentTestSuite) SetupSuite() {
	service, stopService, err := library.NewWebService(context.TODO())
	if err != nil {
		s.T().Fatal(err)
	}

	s.service = service
	s.stopService = stopService
}

func (s *ProductComponentTestSuite) TearDownSuite() {
	s.stopService()
}

func (s ProductComponentTestSuite) TestCreateProductShouldBeSuccess() {
	testProductID := uuid.New().String()
	apitest.New("Test Create Product").
		Handler(s.service).
		Post("/api/product").
		JSON(JSONDictionary{
			"id":          testProductID,
			"categoryRef": uuid.New().String(),
			"name":        "pecel",
			"variants": []JSONDictionary{
				{
					"id":    uuid.New().String(),
					"name":  "paket lengkap",
					"price": 25000,
				},
			},
		}).
		Expect(s.T()).
		Status(http.StatusCreated).
		End()
}

func (s ProductComponentTestSuite) TestCreateProductShouldBeValidateRequest() {

	type (
		ExpectedResponse struct {
			status              int
			responseBodyMatcher func(*http.Response, *http.Request) error
		}
	)

	testTable := []struct {
		requestBody JSONDictionary
		want        ExpectedResponse
	}{
		{
			requestBody: JSONDictionary{
				"categoryRef": uuid.New().String(),
				"name":        "pecel",
				"variants": []JSONDictionary{
					{
						"id":    uuid.New().String(),
						"name":  "paket lengkap",
						"price": 25000,
					},
				},
			},
			want: ExpectedResponse{
				status: http.StatusBadRequest,
				responseBodyMatcher: jsonpath.
					Chain().
					Equal("$.message", "invalid input data").
					Equal("$.messageKey", product.CreationErrKey).
					Equal("$.errorType", errors.ErrorTypeIllegalInputError).
					Contains("$.children[*]", JSONDictionary{
						"errorType":  product.ErrProductIDCantBeEmpty.(*errors.POSError).ErrType(),
						"messageKey": product.ErrProductIDCantBeEmpty.(*errors.POSError).Key(),
						"message":    product.ErrProductIDCantBeEmpty.(*errors.POSError).Message(),
					}).
					End(),
			},
		},

		{
			requestBody: JSONDictionary{
				"id":   uuid.New().String(),
				"name": "pecel",
				"variants": []JSONDictionary{
					{
						"id":    uuid.New().String(),
						"name":  "paket lengkap",
						"price": 25000,
					},
				},
			},
			want: ExpectedResponse{
				status: http.StatusBadRequest,
				responseBodyMatcher: jsonpath.
					Chain().
					Equal("$.message", "invalid input data").
					Equal("$.messageKey", product.CreationErrKey).
					Equal("$.errorType", errors.ErrorTypeIllegalInputError).
					Contains("$.children[*]", JSONDictionary{
						"errorType":  product.ErrProductCategoryRefCantBeEmpty.(*errors.POSError).ErrType(),
						"messageKey": product.ErrProductCategoryRefCantBeEmpty.(*errors.POSError).Key(),
						"message":    product.ErrProductCategoryRefCantBeEmpty.(*errors.POSError).Message(),
					}).
					End(),
			},
		},

		{
			requestBody: JSONDictionary{
				"id":          uuid.New().String(),
				"categoryRef": uuid.New().String(),
				"variants": []JSONDictionary{
					{
						"id":    uuid.New().String(),
						"name":  "paket lengkap",
						"price": 25000,
					},
				},
			},
			want: ExpectedResponse{
				status: http.StatusBadRequest,
				responseBodyMatcher: jsonpath.
					Chain().
					Equal("$.message", "invalid input data").
					Equal("$.messageKey", product.CreationErrKey).
					Equal("$.errorType", errors.ErrorTypeIllegalInputError).
					Contains("$.children[*]", JSONDictionary{
						"errorType":  product.ErrProductNameCantBeEmpty.(*errors.POSError).ErrType(),
						"messageKey": product.ErrProductNameCantBeEmpty.(*errors.POSError).Key(),
						"message":    product.ErrProductNameCantBeEmpty.(*errors.POSError).Message(),
					}).
					End(),
			},
		},

		{
			requestBody: JSONDictionary{
				"id":          uuid.New().String(),
				"categoryRef": uuid.New().String(),
				"name":        "pecel",
			},
			want: ExpectedResponse{
				status: http.StatusBadRequest,
				responseBodyMatcher: jsonpath.
					Chain().
					Equal("$.message", "invalid input data").
					Equal("$.messageKey", product.CreationErrKey).
					Equal("$.errorType", errors.ErrorTypeIllegalInputError).
					Contains("$.children[*]", JSONDictionary{
						"errorType":  product.ErrProductAtLeastHaveOneVariant.(*errors.POSError).ErrType(),
						"messageKey": product.ErrProductAtLeastHaveOneVariant.(*errors.POSError).Key(),
						"message":    product.ErrProductAtLeastHaveOneVariant.(*errors.POSError).Message(),
					}).
					End(),
			},
		},
	}

	for i, r := range testTable {
		testCaseName := fmt.Sprintf("TestCase#%v", i)
		s.Run(testCaseName, func() {
			apitest.New("Test Create Product").
				Debug().
				Handler(s.service).
				Post("/api/product").
				JSON(r.requestBody).
				Expect(s.T()).
				Status(r.want.status).
				Assert(r.want.responseBodyMatcher).
				End()
		})
	}

}

func TestRunProductComponentTestSuite(t *testing.T) {
	suite.Run(t, new(ProductComponentTestSuite))
}
