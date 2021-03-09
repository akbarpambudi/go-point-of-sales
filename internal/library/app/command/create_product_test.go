package command_test

import (
	"context"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/app/command"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/app/command/mock/mockproduct"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/domain/product"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/multierr"
	"testing"
)

//go:generate mockgen -source=./../../domain/product/repository.go -destination=./mock/mockproduct/mock_gen.go -package=mockproduct
type CreateProductHandlerTestSuite struct {
	suite.Suite
	repository *mockproduct.MockRepository
	sut        *command.CreateProductHandler
}

func (s *CreateProductHandlerTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(s.T())
	s.repository = mockproduct.NewMockRepository(mockCtrl)
	s.sut = command.NewCreateProductHandler(s.repository)
}

func (s CreateProductHandlerTestSuite) TestCallHandleToHandleCreateProductCommandShouldBeSuccess() {
	ctx := context.Background()

	variant1, variantConstructionErr := product.NewVariant("74d6eafe-e547-47e0-9058-7621ff86faf2", "110", "Paket Komplit", 1000)
	s.Require().NoError(variantConstructionErr)

	cmd := command.CreateProduct{
		ID:          "7da82be2-139f-4b16-b083-231d7c30ffbf",
		Name:        "Pecel",
		CategoryRef: "6d3335fa-0bf7-4f3c-a439-14d7c2e31c06",
		Variants: []*product.Variant{
			variant1,
		},
	}

	s.repository.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)

	err := s.sut.Handle(ctx, cmd)

	s.Assert().NoError(err)
}

func (s CreateProductHandlerTestSuite) TestCallHandleToHandleCreateProductCommandShouldValidateInput() {
	ctx := context.Background()

	cmd := command.CreateProduct{
		ID:          "",
		Name:        "",
		CategoryRef: "",
		Variants:    []*product.Variant{},
	}

	s.repository.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)

	err := s.sut.Handle(ctx, cmd)
	expectedErr := multierr.Combine(
		product.ErrProductIDCantBeEmpty,
		product.ErrProductCategoryRefCantBeEmpty,
		product.ErrProductNameCantBeEmpty,
		product.ErrProductAtLeastHaveOneVariant,
	)
	s.Assert().Equal(expectedErr, err)
}

func TestRunCreateProductHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(CreateProductHandlerTestSuite))
}
