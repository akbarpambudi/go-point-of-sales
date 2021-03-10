package command_test

import (
	"context"
	"fmt"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/app/command"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/app/command/mock/mockcategory"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/domain/category"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/multierr"
	"testing"
)

//go:generate mockgen -source=./../../domain/category/repository.go -destination=./mock/mockcategory/mock_gen.go -package=mockcategory

type CreateCategoryTestSuite struct {
	suite.Suite
	mockRepository *mockcategory.MockRepository
	sut            *command.CreateCategoryHandlerImpl
}

func (s *CreateCategoryTestSuite) SetupTest() {
	goMockCtrl := gomock.NewController(s.T())
	s.mockRepository = mockcategory.NewMockRepository(goMockCtrl)
	s.sut = command.NewCreateCategoryHandlerImpl(s.mockRepository)
}

func (s CreateCategoryTestSuite) TestCallHandleToHandleCreateCategoryCommandShouldBeSuccess() {
	ctx := context.Background()
	cmd := command.CreateCategory{
		ID:   "7da82be2-139f-4b16-b083-231d7c30ffbf",
		Name: "Main Course",
	}
	s.mockRepository.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)

	err := s.sut.Handle(ctx, cmd)

	s.NoError(err)
}

func (s CreateCategoryTestSuite) TestCallHandleToHandleCreateCategoryCommandShouldPassRightEntityToRepository() {
	ctx := context.Background()
	cmd := command.CreateCategory{
		ID:   "7da82be2-139f-4b16-b083-231d7c30ffbf",
		Name: "Main Course",
	}
	expectedEntity, err := category.NewCategory(cmd.ID, cmd.Name)
	s.Require().NoError(err)

	s.mockRepository.EXPECT().Create(gomock.Any(), gomock.Eq(expectedEntity)).Return(nil)

	_ = s.sut.Handle(ctx, cmd)
}

func (s CreateCategoryTestSuite) TestCallHandleToHandleCreateCategoryCommandShouldReturnValidationErrorWhenInputInvalid() {

	type (
		testArgs struct {
			id   string
			name string
		}

		testRecord struct {
			args testArgs
			want error
		}
	)

	testTable := []testRecord{
		{
			args: testArgs{
				id:   "",
				name: "",
			},
			want: multierr.Combine(category.ErrCategoryIDCantBeEmpty, category.ErrCategoryNameCantBeEmpty),
		},
		{
			args: testArgs{
				id:   "7da82be2-139f-4b16-b083-231d7c30ffbf",
				name: "",
			},
			want: category.ErrCategoryNameCantBeEmpty,
		},
		{
			args: testArgs{
				id:   "",
				name: "Main Course",
			},
			want: category.ErrCategoryIDCantBeEmpty,
		},
	}

	for i, r := range testTable {
		testCaseName := fmt.Sprintf("testCase#%v", i)
		s.Run(testCaseName, func() {
			ctx := context.Background()
			cmd := command.CreateCategory{
				ID:   r.args.id,
				Name: r.args.name,
			}
			s.mockRepository.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)

			got := s.sut.Handle(ctx, cmd)

			s.Assert().Equal(r.want, got)
		})
	}

}

func TestRunCreateCategoryTestSuite(t *testing.T) {
	suite.Run(t, new(CreateCategoryTestSuite))
}
