// +build unit_test

package category_test

import (
	"fmt"
	"testing"

	"github.com/akbarpambudi/go-point-of-sales/internal/library/domain/category"
	"github.com/stretchr/testify/suite"
	"go.uber.org/multierr"
)

type CategoryTestSuite struct {
	suite.Suite
}

func (s CategoryTestSuite) TestCallNewCategoryToCreateNewCategoryShouldBeSuccess() {
	sut, err := category.NewCategory("7da82be2-139f-4b16-b083-231d7c30ffbf", "Main Course")

	s.Run("ShouldNotError", func() {
		s.Assert().NoError(err)
	})

	s.Run("IDShouldBeCorrect", func() {
		s.Assert().Equal("7da82be2-139f-4b16-b083-231d7c30ffbf", sut.ID())
	})

	s.Run("NameShouldBeCorrect", func() {
		s.Assert().Equal("Main Course", sut.Name())
	})
}

func (s CategoryTestSuite) TestCallNewCategoryShouldValidateInput() {
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
			args: testArgs{},
			want: category.IllegalCreationInputError(
				multierr.Combine(category.ErrCategoryIDCantBeEmpty, category.ErrCategoryNameCantBeEmpty),
			),
		},
		{
			args: testArgs{
				id: "7da82be2-139f-4b16-b083-231d7c30ffbf",
			},
			want: category.IllegalCreationInputError(category.ErrCategoryNameCantBeEmpty),
		},
		{
			args: testArgs{
				name: "Main Course",
			},
			want: category.IllegalCreationInputError(category.ErrCategoryIDCantBeEmpty),
		},
	}

	for i, r := range testTable {
		testCaseName := fmt.Sprintf("TestCase#%v", i)
		s.Run(testCaseName, func() {
			_, got := category.NewCategory(r.args.id, r.args.name)
			s.Assert().Equal(r.want, got)
		})
	}
}

func TestRunCategoryTestSuite(t *testing.T) {
	suite.Run(t, new(CategoryTestSuite))
}
