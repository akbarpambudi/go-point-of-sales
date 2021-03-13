package library_test

import (
	"context"
	"fmt"
	"github.com/akbarpambudi/go-point-of-sales/internal/common/errors"
	"github.com/akbarpambudi/go-point-of-sales/internal/common/testinghelper"
	"github.com/akbarpambudi/go-point-of-sales/internal/library"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/adapter/adapterent/ent"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/adapter/adapterent/ent/enttest"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/domain/product"
	"github.com/google/uuid"
	"github.com/steinfletcher/apitest"
	"github.com/steinfletcher/apitest-jsonpath"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type JSONDictionary = testinghelper.JSONDictionary

type ProductComponentTestSuite struct {
	suite.Suite
	stopService func()
	ent         *ent.Client
	service     http.Handler
}

func (s *ProductComponentTestSuite) SetupSuite() {
	s.ent = enttest.Open(s.T(), "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	service, stopService, err := library.NewWebService(context.TODO(), library.SetDataSourceClient(s.ent))
	if err != nil {
		s.T().Fatal(err)
	}
	s.service = service
	s.stopService = stopService
	s.setupDataSample()
}

func (s *ProductComponentTestSuite) TearDownSuite() {
	_ = s.ent.Close()
	s.stopService()
}

func (s ProductComponentTestSuite) TestCreateProductShouldBeSuccess() {
	testProductID := uuid.New().String()
	productAPITest := apitest.New("Test Create Product").
		Handler(s.service)
	testResource := JSONDictionary{
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
	}

	productAPITest.Post("/api/product").
		JSON(testResource).
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

func (s ProductComponentTestSuite) TestGetProductByIDShouldBeSuccess() {
	productAPITest := apitest.New("Test Get Product").
		Handler(s.service)

	productAPITest.Get("/api/product/a6d78307-1c49-4d82-9e77-bc2537f08935").
		Expect(s.T()).
		Status(http.StatusOK).
		Assert(jsonpath.Equal("$.name", "Peperoni Pizza")).
		Assert(jsonpath.Equal("$.id", "a6d78307-1c49-4d82-9e77-bc2537f08935")).
		Assert(jsonpath.Equal("$.categoryRef", "a4c54aba-2bb3-4f33-8b90-30aece918ec9")).
		Assert(jsonpath.Len("$.variants", 2)).
		End()
}

func (s ProductComponentTestSuite) TestGetProductByIDShouldReturnStatusNotFound() {
	productAPITest := apitest.New("Test Get Product").
		Handler(s.service)

	productAPITest.Get("/api/product/" + uuid.NewString()).
		Expect(s.T()).
		Status(http.StatusNotFound).
		End()
}

func (s *ProductComponentTestSuite) setupDataSample() {
	s.Require().NotPanics(func() {
		ctx := context.TODO()
		type variants []struct {
			id    string
			code  string
			name  string
			price float64
		}

		dataSamples := []struct {
			id          string
			name        string
			categoryRef string
			variants    variants
		}{
			{
				id:          "a6d78307-1c49-4d82-9e77-bc2537f08935",
				name:        "Peperoni Pizza",
				categoryRef: "a4c54aba-2bb3-4f33-8b90-30aece918ec9",
				variants: variants{
					{
						id:    "343e8d6e-0efd-405d-8215-a9adcf2968c0",
						code:  "PEPERONI_S_ALACARTE",
						name:  "Small (Ala carte)",
						price: 100000,
					},
					{
						id:    "3a36be50-1dd6-482e-a527-303939b22b25",
						code:  "PEPERONI_M_ALACARTE",
						name:  "MEDIUM (Ala carte)",
						price: 120000,
					},
				},
			},
			{
				id:          "6f34769d-ab04-45ba-b1f4-1f0c4a2d6a52",
				name:        "Gudeg",
				categoryRef: "76667b4f-ca73-4e32-bd12-06249149ab46",
				variants: variants{
					{
						id:    "2eff00d5-05e5-48af-b1f9-9ee9c751f803",
						code:  "GUDEG_K_LNKP",
						name:  "Gudeg Kering (Lengkap)",
						price: 11000,
					},
					{
						id:    "cba7a7e6-3f38-4d6c-a849-71efec0f1237",
						code:  "GUDEG_B_LNKP",
						name:  "Gudeg BASAH (Lengkap)",
						price: 12000,
					},
				},
			},
		}

		for _, r := range dataSamples {
			var vrs []*ent.Variant
			for _, v := range r.variants {
				vr := s.ent.Variant.Create().
					SetID(uuid.MustParse(v.id)).
					SetName(v.name).
					SetPrice(v.price).
					SetCode(v.code).
					SaveX(ctx)
				vrs = append(vrs, vr)
			}

			s.ent.Product.Create().
				SetID(uuid.MustParse(r.id)).
				SetName(r.name).
				SetCategoryRef(r.categoryRef).
				AddVariants(vrs...).
				SaveX(ctx)
		}
	})
}

func TestRunProductComponentTestSuite(t *testing.T) {
	suite.Run(t, new(ProductComponentTestSuite))
}
