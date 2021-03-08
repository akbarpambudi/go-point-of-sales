package product_test

import (
	"fmt"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/domain/product"
	"github.com/stretchr/testify/suite"
	"go.uber.org/multierr"
	"testing"
)

type VariantTestSuite struct {
	suite.Suite
}

func (s VariantTestSuite) TestCallNewVariantToCreateVariantShouldBeSuccess() {
	sut, err := product.NewVariant("74d6eafe-e547-47e0-9058-7621ff86faf2", "11001100", "Paket Komplit", 1000)

	s.Run("ShouldNotError", func() {
		s.Assert().NoError(err)
	})

	s.Run("IDShouldBeCorrect", func() {
		s.Assert().Equal("74d6eafe-e547-47e0-9058-7621ff86faf2", sut.ID())
	})

	s.Run("NameShouldBeCorrect", func() {
		s.Assert().Equal("Paket Komplit", sut.Name())
	})

	s.Run("CodeShouldBeCorrect", func() {
		s.Assert().Equal("11001100", sut.Code())
	})

	s.Run("PriceShouldBeCorrect", func() {
		s.Assert().Equal(float64(1000), sut.Price())
	})

}

func (s VariantTestSuite) TestCallNewVariantToCreateVariantShouldValidateInput() {

	type (
		testRecordArgs struct {
			id    string
			name  string
			code  string
			price float64
		}

		testRecord struct {
			args testRecordArgs
			want error
		}
	)

	tt := []testRecord{
		{
			args: testRecordArgs{
				id:    "",
				name:  "",
				code:  "",
				price: 0,
			},
			want: multierr.Combine(product.ErrVariantIDCantBeEmpty, product.ErrVariantNameCantBeEmpty, product.ErrVariantPriceMustGreaterThanZero),
		},
		{
			args: testRecordArgs{
				name:  "Paket Komplit",
				code:  "1001001",
				price: 1000,
			},
			want: product.ErrVariantIDCantBeEmpty,
		},
		{
			args: testRecordArgs{
				id:    "74d6eafe-e547-47e0-9058-7621ff86faf2",
				code:  "1001001",
				price: 1000,
			},
			want: product.ErrVariantNameCantBeEmpty,
		},
		{
			args: testRecordArgs{
				id:   "74d6eafe-e547-47e0-9058-7621ff86faf2",
				name: "Paket Komplit",
			},
			want: product.ErrVariantPriceMustGreaterThanZero,
		},
		{
			args: testRecordArgs{
				id:    "74d6eafe-e547-47e0-9058-7621ff86faf2",
				name:  "Paket Komplit",
				price: 1000,
			},
			want: nil,
		},
	}

	for i, r := range tt {
		testCaseName := fmt.Sprintf("TestCase#%v", i)
		s.Run(testCaseName, func() {
			_, got := product.NewVariant(r.args.id, r.args.code, r.args.name, r.args.price)

			s.Assert().Equal(got, r.want)
		})
	}

}

func TestRunVariantTestSuite(t *testing.T) {
	suite.Run(t, new(VariantTestSuite))
}
