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
	mockRepository *mockproduct.MockRepository
	sut            *command.CreateProductHandlerImpl
}

func (s *CreateProductHandlerTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(s.T())
	s.mockRepository = mockproduct.NewMockRepository(mockCtrl)
	s.sut = command.NewCreateProductHandlerImpl(s.mockRepository)
}

func (s CreateProductHandlerTestSuite) TestCallHandleToHandleCreateProductCommandShouldBeSuccess() {
	ctx := context.Background()

	cmd := command.CreateProduct{
		ID:          "7da82be2-139f-4b16-b083-231d7c30ffbf",
		Name:        "Pecel",
		CategoryRef: "6d3335fa-0bf7-4f3c-a439-14d7c2e31c06",
		Variants: []command.ProductVariantDTO{
			{
				ID:    "7da82be2-139f-4b16-b083-231d7c30ffbf",
				Code:  "",
				Name:  "paket komplit",
				Price: 10,
			},
		},
	}

	s.mockRepository.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)

	err := s.sut.Handle(ctx, cmd)

	s.Assert().NoError(err)
}

func (s CreateProductHandlerTestSuite) TestCallHandleToHandleCreateProductCommandShouldValidateInput() {
	ctx := context.Background()

	cmd := command.CreateProduct{
		ID:          "",
		Name:        "",
		CategoryRef: "",
		Variants:    []command.ProductVariantDTO{},
	}

	s.mockRepository.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)

	err := s.sut.Handle(ctx, cmd)
	expectedErr := product.IllegalProductCreationInputErr(multierr.Combine(
		product.ErrProductIDCantBeEmpty,
		product.ErrProductCategoryRefCantBeEmpty,
		product.ErrProductNameCantBeEmpty,
		product.ErrProductAtLeastHaveOneVariant,
	))
	s.Assert().Equal(expectedErr, err)
}

func TestRunCreateProductHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(CreateProductHandlerTestSuite))
}
