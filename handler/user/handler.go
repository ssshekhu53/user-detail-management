package user

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/ssshekhu53/user-detail-management/errors"
	"github.com/ssshekhu53/user-detail-management/grpc"
	"github.com/ssshekhu53/user-detail-management/models"
	"github.com/ssshekhu53/user-detail-management/service"
	"github.com/ssshekhu53/user-detail-management/utils"
)

type user struct {
	grpc.UnimplementedUserServiceServer

	userService service.User
}

func New(userService service.User) *user {
	return &user{userService: userService}
}

func (u *user) Create(_ context.Context, req *grpc.UserRequest) (*grpc.User, error) {
	userReq := u.grpcUserRequestToUserRequest(req)

	err := userReq.ValidateMissingParam()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = userReq.ValidateInvalidParam()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	usr, err := u.userService.Create(userReq)
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}

	grpcUser := u.userToGRPCUser(usr)

	return grpcUser, nil
}

func (u *user) Get(context.Context, *emptypb.Empty) (*grpc.Users, error) {
	users := u.userService.Get()

	grpcUsers := u.userToGRPCUsers(users)

	return grpcUsers, nil
}

func (u *user) GetByID(_ context.Context, userID *grpc.UserID) (*grpc.User, error) {
	id := int(userID.GetId())

	if id <= 0 {
		err := errors.InvalidParams{Params: []string{"id"}}

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	usr, err := u.userService.GetByID(id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &grpc.User{
		Id:      int32(usr.ID),
		Fname:   usr.Fname,
		City:    usr.City,
		Phone:   usr.Phone,
		Height:  usr.Height,
		Married: usr.Married,
	}, nil
}

func (u *user) GetByIDs(_ context.Context, userIDs *grpc.UserIDs) (*grpc.Users, error) {
	ids := userIDs.GetIds()

	idsInt, err := u.validateIDs(ids)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	users := u.userService.GetByIDs(idsInt)

	grpcUsers := u.userToGRPCUsers(users)

	return grpcUsers, nil
}

func (u *user) Update(_ context.Context, req *grpc.UserUpdateRequest) (*grpc.User, error) {
	userReq := u.grpcUserUpdateRequestToUserUpdateRequest(req)

	err := userReq.ValidateMissingParam()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = userReq.ValidateInvalidParam()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	usr, err := u.userService.Update(userReq)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	grpcUser := u.userToGRPCUser(usr)

	return grpcUser, nil
}

func (u *user) Delete(_ context.Context, userID *grpc.UserID) (*emptypb.Empty, error) {
	id := int(userID.GetId())

	if id <= 0 {
		err := errors.InvalidParams{Params: []string{"id"}}

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := u.userService.Delete(id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return nil, nil
}

func (u *user) Search(_ context.Context, filters *grpc.Filters) (*grpc.Users, error) {
	users := u.userService.Search(u.grpcFiltersToFilters(filters))
	grpcUsers := u.userToGRPCUsers(users)

	return grpcUsers, nil
}

func (u *user) userToGRPCUsers(users []models.User) *grpc.Users {
	grpcUsers := make([]*grpc.User, 0)

	for i := range users {
		grpcUsers = append(grpcUsers, u.userToGRPCUser(&users[i]))
	}

	return &grpc.Users{Users: grpcUsers}
}

func (u *user) userToGRPCUser(usr *models.User) *grpc.User {
	return &grpc.User{
		Id:      int32(usr.ID),
		Fname:   usr.Fname,
		City:    usr.City,
		Phone:   usr.Phone,
		Height:  usr.Height,
		Married: usr.Married,
	}
}

func (u *user) validateIDs(ids []int32) ([]int, error) {
	idsInt := make([]int, 0)

	for i := range ids {
		if ids[i] <= 0 {
			return nil, errors.InvalidParams{Params: []string{"ids"}}
		}

		idsInt = append(idsInt, int(ids[i]))
	}

	return idsInt, nil
}

func (u *user) grpcFiltersToFilters(filters *grpc.Filters) *models.Filters {
	return &models.Filters{
		Fname:  utils.StrPtr(filters.Fname),
		City:   utils.StrPtr(filters.City),
		Phone:  utils.StrPtr(filters.Phone),
		Height: utils.Float64Ptr(filters.Height),
	}
}

func (u *user) grpcUserRequestToUserRequest(userReq *grpc.UserRequest) *models.UserRequest {
	return &models.UserRequest{
		Fname:   utils.StrPtr(userReq.Fname),
		City:    utils.StrPtr(userReq.City),
		Phone:   utils.StrPtr(userReq.Phone),
		Height:  utils.Float64Ptr(userReq.Height),
		Married: utils.BoolPtr(userReq.Married),
	}
}

func (u *user) grpcUserUpdateRequestToUserUpdateRequest(userReq *grpc.UserUpdateRequest) *models.UserUpdateRequest {
	return &models.UserUpdateRequest{
		ID:      utils.IntPtr(int(userReq.Id)),
		Fname:   utils.StrPtr(userReq.Fname),
		City:    utils.StrPtr(userReq.City),
		Phone:   utils.StrPtr(userReq.Phone),
		Height:  utils.Float64Ptr(userReq.Height),
		Married: utils.BoolPtr(userReq.Married),
	}
}
