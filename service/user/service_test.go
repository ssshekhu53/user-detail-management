package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/ssshekhu53/user-detail-management/errors"
	"github.com/ssshekhu53/user-detail-management/models"
	"github.com/ssshekhu53/user-detail-management/store"
	"github.com/ssshekhu53/user-detail-management/utils"
)

func Test_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockUser(ctrl)
	service := New(mockStore)

	sampleUserReq := &models.UserRequest{
		Fname:   utils.StrPtr("John"),
		City:    utils.StrPtr("New York"),
		Phone:   utils.StrPtr("1234567890"),
		Height:  utils.Float64Ptr(180),
		Married: utils.BoolPtr(false),
	}

	sampleUser := &models.User{
		Fname:   *sampleUserReq.Fname,
		City:    *sampleUserReq.City,
		Phone:   *sampleUserReq.Phone,
		Height:  *sampleUserReq.Height,
		Married: *sampleUserReq.Married,
	}

	sampleFilter := &models.Filters{
		Fname:  sampleUserReq.Fname,
		City:   sampleUserReq.City,
		Phone:  sampleUserReq.Phone,
		Height: sampleUserReq.Height,
	}

	tests := []struct {
		name        string
		userRequest *models.UserRequest
		mockSetup   func()
		expectedUsr *models.User
		expectedErr error
	}{
		{
			"Successful creation", sampleUserReq,
			func() {
				mockStore.EXPECT().Get(sampleFilter).Return(nil)
				mockStore.EXPECT().Create(sampleUser).Return(1)
				mockStore.EXPECT().GetByID(1).Return(&models.User{
					ID:      1,
					Fname:   "John",
					City:    "New York",
					Phone:   "1234567890",
					Height:  180,
					Married: false,
				}, nil)
			},
			&models.User{
				ID:      1,
				Fname:   "John",
				City:    "New York",
				Phone:   "1234567890",
				Height:  180,
				Married: false,
			}, nil,
		},
		{
			"User already exists", sampleUserReq,
			func() {
				mockStore.EXPECT().Get(sampleFilter).Return([]models.User{
					{
						ID:      1,
						Fname:   "John",
						City:    "New York",
						Phone:   "1234567890",
						Height:  180,
						Married: false,
					},
				})
			},
			nil, errors.UserAlreadyExists{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			user, err := service.Create(tt.userRequest)

			assert.Equal(t, tt.expectedUsr, user)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func Test_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockUser(ctrl)
	service := New(mockStore)

	tests := []struct {
		name          string
		mockSetup     func()
		expectedUsers []models.User
	}{
		{
			"Get all users",
			func() {
				mockStore.EXPECT().Get(nil).Return([]models.User{
					{ID: 1, Fname: "John", City: "New York", Phone: "1234567890", Height: 180, Married: false},
					{ID: 2, Fname: "Jane", City: "Los Angeles", Phone: "0987654321", Height: 160, Married: true},
				})
			},
			[]models.User{
				{ID: 1, Fname: "John", City: "New York", Phone: "1234567890", Height: 180, Married: false},
				{ID: 2, Fname: "Jane", City: "Los Angeles", Phone: "0987654321", Height: 160, Married: true},
			},
		},
		{
			"No users found",
			func() {
				mockStore.EXPECT().Get(nil).Return([]models.User{})
			},
			[]models.User{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			users := service.Get()
			assert.Equal(t, tt.expectedUsers, users)
		})
	}
}

func Test_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockUser(ctrl)
	service := New(mockStore)

	tests := []struct {
		name        string
		userID      int
		mockSetup   func()
		expectedUsr *models.User
		expectedErr error
	}{
		{
			"User found", 1,
			func() {
				mockStore.EXPECT().GetByID(1).Return(&models.User{
					ID:      1,
					Fname:   "John",
					City:    "New York",
					Phone:   "1234567890",
					Height:  180,
					Married: false,
				}, nil)
			},
			&models.User{
				ID:      1,
				Fname:   "John",
				City:    "New York",
				Phone:   "1234567890",
				Height:  180,
				Married: false,
			}, nil,
		},
		{
			"User not found", 2,
			func() {
				mockStore.EXPECT().GetByID(2).Return(nil, errors.UserNotFound{ID: 2})
			},
			nil, errors.UserNotFound{ID: 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			user, err := service.GetByID(tt.userID)

			assert.Equal(t, tt.expectedUsr, user)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func Test_GetByIDs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockUser(ctrl)
	service := New(mockStore)

	ids := []int{1, 2}

	tests := []struct {
		name          string
		mockSetup     func()
		expectedUsers []models.User
	}{
		{
			"Get all users",
			func() {
				mockStore.EXPECT().GetByIDs(ids).Return([]models.User{
					{ID: 1, Fname: "John", City: "New York", Phone: "1234567890", Height: 180, Married: false},
					{ID: 2, Fname: "Jane", City: "Los Angeles", Phone: "0987654321", Height: 160, Married: true},
				})
			},
			[]models.User{
				{ID: 1, Fname: "John", City: "New York", Phone: "1234567890", Height: 180, Married: false},
				{ID: 2, Fname: "Jane", City: "Los Angeles", Phone: "0987654321", Height: 160, Married: true},
			},
		},
		{
			"No users found",
			func() {
				mockStore.EXPECT().GetByIDs(ids).Return([]models.User{})
			},
			[]models.User{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			users := service.GetByIDs(ids)
			assert.Equal(t, tt.expectedUsers, users)
		})
	}
}

func Test_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockUser(ctrl)
	service := New(mockStore)

	sampleUserReq := &models.UserUpdateRequest{
		ID:      utils.IntPtr(1),
		Fname:   utils.StrPtr("Johnny"),
		City:    utils.StrPtr("San Francisco"),
		Phone:   utils.StrPtr("0987654321"),
		Height:  utils.Float64Ptr(185),
		Married: utils.BoolPtr(true),
	}

	tests := []struct {
		name        string
		userRequest *models.UserUpdateRequest
		mockSetup   func()
		expectedUsr *models.User
		expectedErr error
	}{
		{
			"Successful update", sampleUserReq,
			func() {
				mockStore.EXPECT().GetByID(1).Return(&models.User{
					ID:      1,
					Fname:   "John",
					City:    "New York",
					Phone:   "1234567890",
					Height:  180,
					Married: false,
				}, nil)
				mockStore.EXPECT().Update(gomock.Any()).Times(1)
				mockStore.EXPECT().GetByID(1).Return(&models.User{
					ID:      1,
					Fname:   "Johnny",
					City:    "San Francisco",
					Phone:   "0987654321",
					Height:  185,
					Married: true,
				}, nil).Times(1)
			},
			&models.User{
				ID:      1,
				Fname:   "Johnny",
				City:    "San Francisco",
				Phone:   "0987654321",
				Height:  185,
				Married: true,
			},
			nil,
		},
		{
			"User not found", sampleUserReq,
			func() {
				mockStore.EXPECT().GetByID(1).Return(nil, errors.UserNotFound{ID: 1}).Times(1)
			},
			nil, errors.UserNotFound{ID: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			updatedUser, err := service.Update(tt.userRequest)

			assert.Equal(t, tt.expectedUsr, updatedUser)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func Test_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockUser(ctrl)
	service := New(mockStore)

	tests := []struct {
		name        string
		userID      int
		mockSetup   func()
		expectedErr error
	}{
		{
			"Successful deletion", 1,
			func() {
				mockStore.EXPECT().GetByID(1).Return(&models.User{
					ID: 1,
				}, nil).Times(1)
				mockStore.EXPECT().Delete(1).Times(1)
			},
			nil,
		},
		{
			"User not found", 2,
			func() {
				mockStore.EXPECT().GetByID(2).Return(nil, errors.UserNotFound{ID: 2}).Times(1)
			},
			errors.UserNotFound{ID: 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			err := service.Delete(tt.userID)

			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func Test_Search(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockUser(ctrl)
	service := New(mockStore)

	tests := []struct {
		name          string
		filters       *models.Filters
		mockSetup     func()
		expectedUsers []models.User
	}{
		{
			"Search by Fname and City",
			&models.Filters{
				Fname: utils.StrPtr("John"),
				City:  utils.StrPtr("New York"),
			},
			func() {
				mockStore.EXPECT().Get(&models.Filters{
					Fname: utils.StrPtr("John"),
					City:  utils.StrPtr("New York"),
				}).Return([]models.User{
					{ID: 1, Fname: "John", City: "New York", Phone: "1234567890", Height: 180, Married: false},
				})
			},
			[]models.User{
				{ID: 1, Fname: "John", City: "New York", Phone: "1234567890", Height: 180, Married: false},
			},
		},
		{
			"No users found",
			&models.Filters{
				Fname: utils.StrPtr("Jane"),
				City:  utils.StrPtr("Los Angeles"),
			},
			func() {
				mockStore.EXPECT().Get(&models.Filters{
					Fname: utils.StrPtr("Jane"),
					City:  utils.StrPtr("Los Angeles"),
				}).Return([]models.User{})
			},
			[]models.User{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			users := service.Search(tt.filters)

			assert.Equal(t, tt.expectedUsers, users)
		})
	}
}
