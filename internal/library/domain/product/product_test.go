// +build unit_test

package product_test

import (
	"fmt"
	"testing"

	"github.com/akbarpambudi/go-point-of-sales/internal/library/domain/product"
	"github.com/stretchr/testify/suite"
	"go.uber.org/multierr"
)

type ItemTestSuite struct {
	suite.Suite
}

func (s ItemTestSuite) TestCallNewProductToCreateItemShouldBeSuccess() {

	variant, err := product.NewVariant(
		"7da82be2-139f-4b16-b083-231d7c30ffbf",
		"",
		"paket lengkap",
		1000,
	)

	s.Require().NoError(err)

	sut, itemErr := product.NewProduct(
		"7007d79e-8dfa-49dc-b89d-3a87523b616b",
		"6d3335fa-0bf7-4f3c-a439-14d7c2e31c06",
		"Pecel",
		[]*product.Variant{
			variant,
		},
	)

	s.Run("ShouldNotError", func() {
		s.Assert().NoError(itemErr)
	})

	s.Run("ItemIDShouldBeCorrect", func() {
		s.Assert().Equal("7007d79e-8dfa-49dc-b89d-3a87523b616b", sut.ID())
	})

	s.Run("ItemCategoryRefShouldBeCorrect", func() {
		s.Assert().Equal("6d3335fa-0bf7-4f3c-a439-14d7c2e31c06", sut.CategoryRef())
	})

	s.Run("ItemNameShouldBeCorrect", func() {
		s.Assert().Equal("Pecel", sut.Name())
	})

	s.Run("ItemVariantsLengthShouldBeCorrect", func() {
		s.Assert().Len(sut.Variants(), 1)
	})
}

func (s ItemTestSuite) TestCallAddVariantToAddNewVariantToItem() {

	variant, err := product.NewVariant(
		"7da82be2-139f-4b16-b083-231d7c30ffbf",
		"",
		"paket lengkap",
		1000,
	)

	s.Require().NoError(err)

	sut, err := product.NewProduct(
		"7007d79e-8dfa-49dc-b89d-3a87523b616b",
		"6d3335fa-0bf7-4f3c-a439-14d7c2e31c06",
		"Pecel",
		[]*product.Variant{
			variant,
		},
	)

	additionalVariant, err := product.NewVariant(
		"95827ca7-2244-47ab-8eca-d0b04715cfa3",
		"",
		"paket ala carte",
		12000,
	)

	s.Require().NoError(err)

	sut.AddVariant(additionalVariant)

	s.Run("ItemVariantsLengthShouldBeCorrect", func() {
		s.Assert().Len(sut.Variants(), 2)
	})
}

func (s ItemTestSuite) TestCallAddManyVariantToAddMultipleVariantToItem() {

	variant, err := product.NewVariant(
		"7da82be2-139f-4b16-b083-231d7c30ffbf",
		"",
		"paket lengkap",
		1000,
	)

	s.Require().NoError(err)

	sut, err := product.NewProduct(
		"7007d79e-8dfa-49dc-b89d-3a87523b616b",
		"6d3335fa-0bf7-4f3c-a439-14d7c2e31c06",
		"Pecel",
		[]*product.Variant{
			variant,
		},
	)

	additionalVariant, err := product.NewVariant(
		"7da82be2-139f-4b16-b083-231d7c30ffbf",
		"",
		"paket ala carte",
		1200,
	)

	s.Require().NoError(err)

	additionalVariant2, err := product.NewVariant(
		"7da82be2-139f-4b16-b083-231d7c30ffbf",
		"",
		"paket spesial",
		1300,
	)

	s.Require().NoError(err)

	sut.AddManyVariant([]*product.Variant{additionalVariant, additionalVariant2})

	s.Run("ItemVariantsLengthShouldBeCorrect", func() {
		s.Assert().Len(sut.Variants(), 3)
	})
}

func (s ItemTestSuite) TestCallNewProductToCreateNewProductShouldValidateInput() {
	type (
		testTableTestArgsVariant struct {
			id    string
			name  string
			price float64
			code  string
		}

		testTableTestArgs struct {
			id          string
			name        string
			categoryRef string
			variants    []testTableTestArgsVariant
		}

		testTable struct {
			testArgs testTableTestArgs
			want     error
		}
	)

	tt := []testTable{
		{
			testArgs: testTableTestArgs{
				id:          "7da82be2-139f-4b16-b083-231d7c30ffbf",
				name:        "Pecel",
				categoryRef: "6d3335fa-0bf7-4f3c-a439-14d7c2e31c06",
				variants:    nil,
			},
			want: product.IllegalProductCreationInputErr(product.ErrProductAtLeastHaveOneVariant),
		},
		{
			testArgs: testTableTestArgs{
				name:        "Pecel",
				categoryRef: "6d3335fa-0bf7-4f3c-a439-14d7c2e31c06",
				variants: []testTableTestArgsVariant{
					{
						id:    "7da82be2-139f-4b16-b083-231d7c30ffbf",
						name:  "paket spesial",
						price: 1300,
						code:  "",
					},
				},
			},
			want: product.IllegalProductCreationInputErr(product.ErrProductIDCantBeEmpty),
		},
		{
			testArgs: testTableTestArgs{
				id:          "7da82be2-139f-4b16-b083-231d7c30ffbf",
				categoryRef: "6d3335fa-0bf7-4f3c-a439-14d7c2e31c06",
				variants: []testTableTestArgsVariant{
					{
						id:    "7da82be2-139f-4b16-b083-231d7c30ffbf",
						name:  "paket spesial",
						price: 1300,
						code:  "",
					},
				},
			},
			want: product.IllegalProductCreationInputErr(product.ErrProductNameCantBeEmpty),
		},
		{
			testArgs: testTableTestArgs{
				id:   "7da82be2-139f-4b16-b083-231d7c30ffbf",
				name: "Pecel",
				variants: []testTableTestArgsVariant{
					{
						id:    "7da82be2-139f-4b16-b083-231d7c30ffbf",
						name:  "paket spesial",
						price: 1300,
						code:  "",
					},
				},
			},
			want: product.IllegalProductCreationInputErr(product.ErrProductCategoryRefCantBeEmpty),
		},
		{
			testArgs: testTableTestArgs{},
			want:     product.IllegalProductCreationInputErr(multierr.Combine(product.ErrProductIDCantBeEmpty, product.ErrProductCategoryRefCantBeEmpty, product.ErrProductNameCantBeEmpty, product.ErrProductAtLeastHaveOneVariant)),
		},
	}

	for i, r := range tt {
		testCaseName := fmt.Sprintf("TestCase#%v", i)
		s.Run(testCaseName, func() {
			var variants []*product.Variant
			for _, v := range r.testArgs.variants {
				variant, err := product.NewVariant(v.id, v.code, v.name, v.price)
				s.Require().NoError(err)
				variants = append(variants, variant)
			}
			_, got := product.NewProduct(r.testArgs.id, r.testArgs.categoryRef, r.testArgs.name, variants)

			s.Assert().EqualError(got, r.want.Error())
		})
	}

}

func TestItemTestSuite(t *testing.T) {
	suite.Run(t, new(ItemTestSuite))
}
