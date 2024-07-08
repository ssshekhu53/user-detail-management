package user

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/ssshekhu53/user-detail-management/errors"
	"github.com/ssshekhu53/user-detail-management/grpc"
	"github.com/ssshekhu53/user-detail-management/models"
	"github.com/ssshekhu53/user-detail-management/service"
)

func Test_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockUser(ctrl)
	handler := New(mockService)

	sampleReq := &grpc.UserRequest{
		Fname:   "John",
		City:    "New York",
		Phone:   "1234567890",
		Height:  180,
		Married: false,
	}

	sampleUser := &models.User{
		ID:      1,
		Fname:   "John",
		City:    "New York",
		Phone:   "1234567890",
		Height:  180,
		Married: false,
	}

	tests := []struct {
		name         string
		req          *grpc.UserRequest
		mockSetup    func()
		expectedResp *grpc.User
		expectedErr  error
	}{
		{
			"Success", sampleReq,
			func() {
				mockService.EXPECT().Create(gomock.Any()).Return(sampleUser, nil)
			},
			&grpc.User{
				Id:      1,
				Fname:   "John",
				City:    "New York",
				Phone:   "1234567890",
				Height:  180,
				Married: false,
			}, nil,
		},
		{
			"Missing params", &grpc.UserRequest{
				City:    "New York",
				Phone:   "1234567890",
				Height:  180,
				Married: false,
			},
			func() {
				// No mock expected as it should fail before calling the service
			},
			nil, status.Error(codes.InvalidArgument, "missing param: fname"),
		},
		{
			"Invalid params", &grpc.UserRequest{
				Fname:   "John",
				City:    "New York",
				Phone:   "1234567890",
				Height:  -180,
				Married: false,
			},
			func() {
				// No mock expected as it should fail before calling the service
			},
			nil, status.Error(codes.InvalidArgument, "invalid param: height"),
		},
		{
			"User Already Exists", sampleReq,
			func() {
				mockService.EXPECT().Create(gomock.Any()).Return(nil, errors.UserAlreadyExists{})
			},
			nil, status.Error(codes.AlreadyExists, "user already exists with given combination"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			resp, err := handler.Create(context.Background(), tt.req)

			assert.Equal(t, tt.expectedResp, resp)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func Test_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockUser(ctrl)
	handler := New(mockService)

	sampleUsers := []models.User{
		{
			ID:      1,
			Fname:   "John",
			City:    "New York",
			Phone:   "1234567890",
			Height:  180,
			Married: false,
		},
	}

	tests := []struct {
		name         string
		mockSetup    func()
		expectedResp *grpc.Users
		expectedErr  error
	}{
		{
			"Success",
			func() {
				mockService.EXPECT().Get().Return(sampleUsers)
			},
			&grpc.Users{
				Users: []*grpc.User{
					{
						Id:      1,
						Fname:   "John",
						City:    "New York",
						Phone:   "1234567890",
						Height:  180,
						Married: false,
					},
				},
			}, nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			resp, err := handler.Get(context.Background(), &emptypb.Empty{})

			assert.Equal(t, tt.expectedResp, resp)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func Test_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockUser(ctrl)
	handler := New(mockService)

	tests := []struct {
		name         string
		id           int32
		mockSetup    func()
		expectedResp *grpc.User
		expectedErr  error
	}{
		{
			"Success", 1,
			func() {
				mockService.EXPECT().GetByID(1).Return(&models.User{
					ID:      1,
					Fname:   "John",
					City:    "New York",
					Phone:   "1234567890",
					Height:  180,
					Married: false,
				}, nil)
			},
			&grpc.User{
				Id:      1,
				Fname:   "John",
				City:    "New York",
				Phone:   "1234567890",
				Height:  180,
				Married: false,
			}, nil,
		},
		{
			"Invalid id", -2,
			func() {
				// No mock expected as it should fail before calling the service
			},
			nil, status.Error(codes.InvalidArgument, "invalid param: id"),
		},
		{
			"User Not Found", 2,
			func() {
				mockService.EXPECT().GetByID(2).Return(nil, errors.UserNotFound{ID: 2})
			},
			nil, status.Error(codes.NotFound, "user with ID 2 not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			resp, err := handler.GetByID(context.Background(), &grpc.UserID{Id: tt.id})

			assert.Equal(t, tt.expectedResp, resp)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func Test_GetByIDs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockUser(ctrl)
	handler := New(mockService)

	sampleUsers := []models.User{
		{
			ID:      1,
			Fname:   "John",
			City:    "New York",
			Phone:   "1234567890",
			Height:  180,
			Married: false,
		},
		{
			ID:      2,
			Fname:   "Jane",
			City:    "Los Angeles",
			Phone:   "0987654321",
			Height:  170,
			Married: true,
		},
	}

	tests := []struct {
		name         string
		ids          []int32
		mockSetup    func()
		expectedResp *grpc.Users
		expectedErr  error
	}{
		{
			"Success", []int32{1, 2},
			func() {
				mockService.EXPECT().GetByIDs([]int{1, 2}).Return(sampleUsers)
			},
			&grpc.Users{
				Users: []*grpc.User{
					{
						Id:      1,
						Fname:   "John",
						City:    "New York",
						Phone:   "1234567890",
						Height:  180,
						Married: false,
					},
					{
						Id:      2,
						Fname:   "Jane",
						City:    "Los Angeles",
						Phone:   "0987654321",
						Height:  170,
						Married: true,
					},
				},
			}, nil,
		},
		{
			"Invalid ID", []int32{-1, 2},
			func() {
				// No mock expected as it should fail before calling the service
			},
			nil, status.Error(codes.InvalidArgument, "invalid param: ids"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			resp, err := handler.GetByIDs(context.Background(), &grpc.UserIDs{Ids: tt.ids})

			assert.Equal(t, tt.expectedResp, resp)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func Test_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockUser(ctrl)
	handler := New(mockService)

	sampleUpdateReq := &grpc.UserUpdateRequest{
		Id:      1,
		Fname:   "John",
		City:    "New York",
		Phone:   "1234567890",
		Height:  180,
		Married: false,
	}

	sampleUser := &models.User{
		ID:      1,
		Fname:   "John",
		City:    "New York",
		Phone:   "1234567890",
		Height:  180,
		Married: false,
	}

	tests := []struct {
		name         string
		req          *grpc.UserUpdateRequest
		mockSetup    func()
		expectedResp *grpc.User
		expectedErr  error
	}{
		{
			"Success", sampleUpdateReq,
			func() {
				mockService.EXPECT().Update(gomock.Any()).Return(sampleUser, nil)
			},
			&grpc.User{
				Id:      1,
				Fname:   "John",
				City:    "New York",
				Phone:   "1234567890",
				Height:  180,
				Married: false,
			}, nil,
		},
		{
			"Missing params", &grpc.UserUpdateRequest{
				Id:      1,
				City:    "New York",
				Phone:   "1234567890",
				Height:  180,
				Married: false,
			},
			func() {
				// No mock expected as it should fail before calling the service
			},
			nil, status.Error(codes.InvalidArgument, "missing param: fname"),
		},
		{
			"Invalid params", &grpc.UserUpdateRequest{
				Id:      1,
				Fname:   "John",
				City:    "New York",
				Phone:   "1234567890",
				Height:  -180,
				Married: false,
			},
			func() {
				// No mock expected as it should fail before calling the service
			},
			nil, status.Error(codes.InvalidArgument, "invalid param: height"),
		},
		{
			"User Not Found", sampleUpdateReq,
			func() {
				mockService.EXPECT().Update(gomock.Any()).Return(nil, errors.UserNotFound{ID: 1})
			},
			nil, status.Error(codes.NotFound, "user with ID 1 not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			resp, err := handler.Update(context.Background(), tt.req)

			assert.Equal(t, tt.expectedResp, resp)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func Test_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockUser(ctrl)
	handler := New(mockService)

	tests := []struct {
		name        string
		id          int32
		mockSetup   func()
		expectedErr error
	}{
		{
			"Success", 1,
			func() {
				mockService.EXPECT().Delete(1).Return(nil)
			}, nil,
		},
		{
			"Invalid ID", -1,
			func() {
				// No mock expected as it should fail before calling the service
			}, status.Error(codes.InvalidArgument, "invalid param: id"),
		},
		{
			"not found", 1,
			func() {
				mockService.EXPECT().Delete(1).Return(errors.UserNotFound{ID: 1})
			}, status.Error(codes.NotFound, "user with ID 1 not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			_, err := handler.Delete(context.Background(), &grpc.UserID{Id: tt.id})

			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func Test_Search(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockUser(ctrl)
	handler := New(mockService)

	sampleFilters := &grpc.Filters{
		Fname:   "John",
		City:    "New York",
		Phone:   "1234567890",
		Height:  180,
		Married: false,
	}

	sampleUsers := []models.User{
		{
			ID:      1,
			Fname:   "John",
			City:    "New York",
			Phone:   "1234567890",
			Height:  180,
			Married: false,
		},
	}

	tests := []struct {
		name         string
		filters      *grpc.Filters
		mockSetup    func()
		expectedResp *grpc.Users
		expectedErr  error
	}{
		{
			"Success",
			sampleFilters,
			func() {
				mockService.EXPECT().Search(gomock.Any()).Return(sampleUsers)
			},
			&grpc.Users{
				Users: []*grpc.User{
					{
						Id:      1,
						Fname:   "John",
						City:    "New York",
						Phone:   "1234567890",
						Height:  180,
						Married: false,
					},
				},
			}, nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			resp, err := handler.Search(context.Background(), tt.filters)

			assert.Equal(t, tt.expectedResp, resp)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
