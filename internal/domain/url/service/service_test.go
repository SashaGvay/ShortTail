package service

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"short_tail/internal/domain/url/models"
	"testing"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Put(ctx context.Context, dto *models.URL) error {
	args := m.Called(ctx, dto)
	return args.Error(0)
}

func (m *MockRepository) Get(ctx context.Context, alias string) (*models.URL, error) {
	args := m.Called(ctx, alias)
	return args.Get(0).(*models.URL), args.Error(1)
}

func TestShort_Success(t *testing.T) {
	mockRepo := new(MockRepository)
	svc := Service{Repository: mockRepo}

	ctx := context.Background()
	testURL := "https://example.com"
	expectedAlias := svc.generateAlias(testURL)

	mockRepo.On("Put", ctx, mock.MatchedBy(func(dto *models.URL) bool {
		return dto.Original == testURL && dto.Alias == expectedAlias
	})).Return(nil)

	result, err := svc.Short(ctx, testURL)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, testURL, result.Original)
	assert.Equal(t, expectedAlias, result.Alias)

	mockRepo.AssertExpectations(t)
}

func TestShort_ErrorOnPut(t *testing.T) {
	mockRepo := new(MockRepository)
	svc := Service{Repository: mockRepo}

	ctx := context.Background()
	testURL := "https://example.com"

	mockRepo.On("Put", ctx, mock.Anything).Return(errors.New("any error"))

	result, err := svc.Short(ctx, testURL)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "s.Repository.Put")
}

func TestUnShort_Success(t *testing.T) {
	mockRepo := new(MockRepository)
	svc := Service{Repository: mockRepo}

	ctx := context.Background()
	testAlias := "abc123"
	expectedURL := &models.URL{
		Original: "https://example.com",
		Alias:    testAlias,
	}

	mockRepo.On("Get", ctx, testAlias).Return(expectedURL, nil)

	result, err := svc.UnShort(ctx, testAlias)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedURL.Original, result.Original)
	assert.Equal(t, expectedURL.Alias, result.Alias)

	mockRepo.AssertExpectations(t)
}

func TestUnShort_AliasNotFound(t *testing.T) {
	mockRepo := new(MockRepository)
	svc := Service{Repository: mockRepo}

	ctx := context.Background()
	testAlias := "not_found"

	mockRepo.On("Get", ctx, testAlias).Return((*models.URL)(nil), errors.New("not found"))

	result, err := svc.UnShort(ctx, testAlias)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "s.Repository.Get")

	mockRepo.AssertExpectations(t)
}
