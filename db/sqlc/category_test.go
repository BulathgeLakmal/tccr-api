package db

import (
	"context"
	"testing"

	"github.com/it21152832/Learning-Backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomCategory(t *testing.T) Category {
	arg := CreateCategoryParams{
		CategoryName:  util.RandomString(10),
		CategoryDesc:  util.RandomString(10),
	}

	category, err := testQueries.CreateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category.CategoryName)
	require.NotEmpty(t, category.CategoryDesc)
	require.NotZero(t, category.CategoryID)

	return category
}

func TestCreateCategory(t *testing.T) {
	createRandomCategory(t)
}
