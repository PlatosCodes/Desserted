// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: service_desserted.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Desserted_CreateUser_FullMethodName          = "/pb.Desserted/CreateUser"
	Desserted_LoginUser_FullMethodName           = "/pb.Desserted/LoginUser"
	Desserted_Logout_FullMethodName              = "/pb.Desserted/Logout"
	Desserted_CheckUserSession_FullMethodName    = "/pb.Desserted/CheckUserSession"
	Desserted_RenewAccess_FullMethodName         = "/pb.Desserted/RenewAccess"
	Desserted_UpdateUser_FullMethodName          = "/pb.Desserted/UpdateUser"
	Desserted_CreateFriendship_FullMethodName    = "/pb.Desserted/CreateFriendship"
	Desserted_ListUserFriends_FullMethodName     = "/pb.Desserted/ListUserFriends"
	Desserted_ListFriendRequests_FullMethodName  = "/pb.Desserted/ListFriendRequests"
	Desserted_AcceptFriendRequest_FullMethodName = "/pb.Desserted/AcceptFriendRequest"
	Desserted_CreateGame_FullMethodName          = "/pb.Desserted/CreateGame"
	Desserted_InvitePlayerToGame_FullMethodName  = "/pb.Desserted/InvitePlayerToGame"
	Desserted_ListGameInvites_FullMethodName     = "/pb.Desserted/ListGameInvites"
	Desserted_AcceptGameInvite_FullMethodName    = "/pb.Desserted/AcceptGameInvite"
	Desserted_ListGamePlayers_FullMethodName     = "/pb.Desserted/ListGamePlayers"
	Desserted_GetPlayerGame_FullMethodName       = "/pb.Desserted/GetPlayerGame"
	Desserted_StartGame_FullMethodName           = "/pb.Desserted/StartGame"
	Desserted_GetPlayerHand_FullMethodName       = "/pb.Desserted/GetPlayerHand"
	Desserted_PlayDessert_FullMethodName         = "/pb.Desserted/PlayDessert"
	Desserted_DrawCard_FullMethodName            = "/pb.Desserted/DrawCard"
)

// DessertedClient is the client API for Desserted service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DessertedClient interface {
	// Creates a new user
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	// Logs in a user
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
	// Logs out a user
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CheckUserSession(ctx context.Context, in *CheckUserSessionRequest, opts ...grpc.CallOption) (*CheckUserSessionResponse, error)
	RenewAccess(ctx context.Context, in *RenewAccessRequest, opts ...grpc.CallOption) (*RenewAccessResponse, error)
	// Updates a user's password
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	CreateFriendship(ctx context.Context, in *CreateFriendshipRequest, opts ...grpc.CallOption) (*CreateFriendshipResponse, error)
	ListUserFriends(ctx context.Context, in *ListUserFriendsRequest, opts ...grpc.CallOption) (*ListUserFriendsResponse, error)
	ListFriendRequests(ctx context.Context, in *ListFriendRequestsRequest, opts ...grpc.CallOption) (*ListFriendRequestsResponse, error)
	AcceptFriendRequest(ctx context.Context, in *AcceptFriendRequestRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateGame(ctx context.Context, in *CreateGameRequest, opts ...grpc.CallOption) (*CreateGameResponse, error)
	InvitePlayerToGame(ctx context.Context, in *InvitePlayerToGameRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListGameInvites(ctx context.Context, in *ListGameInvitesRequest, opts ...grpc.CallOption) (*ListGameInvitesResponse, error)
	AcceptGameInvite(ctx context.Context, in *AcceptGameInviteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	//	rpc AddPlayerToGame (AddPlayerToGameRequest) returns (google.protobuf.Empty) {
	//	  option (google.api.http) = {
	//	      post: "/v1/add_player_to_game"
	//	      body: "*"
	//	  };
	//	  option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
	//	    description: "Use this API to add player to the game";
	//	    summary: "Add player to the game";
	//	  };
	//	};
	ListGamePlayers(ctx context.Context, in *ListGamePlayersRequest, opts ...grpc.CallOption) (*ListGamePlayersResponse, error)
	GetPlayerGame(ctx context.Context, in *GetPlayerGameRequest, opts ...grpc.CallOption) (*GetPlayerGameResponse, error)
	StartGame(ctx context.Context, in *StartGameRequest, opts ...grpc.CallOption) (*StartGameResponse, error)
	GetPlayerHand(ctx context.Context, in *GetPlayerHandRequest, opts ...grpc.CallOption) (*GetPlayerHandResponse, error)
	PlayDessert(ctx context.Context, in *PlayDessertRequest, opts ...grpc.CallOption) (*PlayDessertResponse, error)
	DrawCard(ctx context.Context, in *DrawCardRequest, opts ...grpc.CallOption) (*DrawCardResponse, error)
}

type dessertedClient struct {
	cc grpc.ClientConnInterface
}

func NewDessertedClient(cc grpc.ClientConnInterface) DessertedClient {
	return &dessertedClient{cc}
}

func (c *dessertedClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, Desserted_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dessertedClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, Desserted_LoginUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dessertedClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Desserted_Logout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dessertedClient) CheckUserSession(ctx context.Context, in *CheckUserSessionRequest, opts ...grpc.CallOption) (*CheckUserSessionResponse, error) {
	out := new(CheckUserSessionResponse)
	err := c.cc.Invoke(ctx, Desserted_CheckUserSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dessertedClient) RenewAccess(ctx context.Context, in *RenewAccessRequest, opts ...grpc.CallOption) (*RenewAccessResponse, error) {
	out := new(RenewAccessResponse)
	err := c.cc.Invoke(ctx, Desserted_RenewAccess_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dessertedClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, Desserted_UpdateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dessertedClient) CreateFriendship(ctx context.Context, in *CreateFriendshipRequest, opts ...grpc.CallOption) (*CreateFriendshipResponse, error) {
	out := new(CreateFriendshipResponse)
	err := c.cc.Invoke(ctx, Desserted_CreateFriendship_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dessertedClient) ListUserFriends(ctx context.Context, in *ListUserFriendsRequest, opts ...grpc.CallOption) (*ListUserFriendsResponse, error) {
	out := new(ListUserFriendsResponse)
	err := c.cc.Invoke(ctx, Desserted_ListUserFriends_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dessertedClient) ListFriendRequests(ctx context.Context, in *ListFriendRequestsRequest, opts ...grpc.CallOption) (*ListFriendRequestsResponse, error) {
	out := new(ListFriendRequestsResponse)
	err := c.cc.Invoke(ctx, Desserted_ListFriendRequests_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dessertedClient) AcceptFriendRequest(ctx context.Context, in *AcceptFriendRequestRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Desserted_AcceptFriendRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dessertedClient) CreateGame(ctx context.Context, in *CreateGameRequest, opts ...grpc.CallOption) (*CreateGameResponse, error) {
	out := new(CreateGameResponse)
	err := c.cc.Invoke(ctx, Desserted_CreateGame_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dessertedClient) InvitePlayerToGame(ctx context.Context, in *InvitePlayerToGameRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Desserted_InvitePlayerToGame_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dessertedClient) ListGameInvites(ctx context.Context, in *ListGameInvitesRequest, opts ...grpc.CallOption) (*ListGameInvitesResponse, error) {
	out := new(ListGameInvitesResponse)
	err := c.cc.Invoke(ctx, Desserted_ListGameInvites_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dessertedClient) AcceptGameInvite(ctx context.Context, in *AcceptGameInviteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Desserted_AcceptGameInvite_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dessertedClient) ListGamePlayers(ctx context.Context, in *ListGamePlayersRequest, opts ...grpc.CallOption) (*ListGamePlayersResponse, error) {
	out := new(ListGamePlayersResponse)
	err := c.cc.Invoke(ctx, Desserted_ListGamePlayers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dessertedClient) GetPlayerGame(ctx context.Context, in *GetPlayerGameRequest, opts ...grpc.CallOption) (*GetPlayerGameResponse, error) {
	out := new(GetPlayerGameResponse)
	err := c.cc.Invoke(ctx, Desserted_GetPlayerGame_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dessertedClient) StartGame(ctx context.Context, in *StartGameRequest, opts ...grpc.CallOption) (*StartGameResponse, error) {
	out := new(StartGameResponse)
	err := c.cc.Invoke(ctx, Desserted_StartGame_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dessertedClient) GetPlayerHand(ctx context.Context, in *GetPlayerHandRequest, opts ...grpc.CallOption) (*GetPlayerHandResponse, error) {
	out := new(GetPlayerHandResponse)
	err := c.cc.Invoke(ctx, Desserted_GetPlayerHand_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dessertedClient) PlayDessert(ctx context.Context, in *PlayDessertRequest, opts ...grpc.CallOption) (*PlayDessertResponse, error) {
	out := new(PlayDessertResponse)
	err := c.cc.Invoke(ctx, Desserted_PlayDessert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dessertedClient) DrawCard(ctx context.Context, in *DrawCardRequest, opts ...grpc.CallOption) (*DrawCardResponse, error) {
	out := new(DrawCardResponse)
	err := c.cc.Invoke(ctx, Desserted_DrawCard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DessertedServer is the server API for Desserted service.
// All implementations must embed UnimplementedDessertedServer
// for forward compatibility
type DessertedServer interface {
	// Creates a new user
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	// Logs in a user
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	// Logs out a user
	Logout(context.Context, *LogoutRequest) (*emptypb.Empty, error)
	CheckUserSession(context.Context, *CheckUserSessionRequest) (*CheckUserSessionResponse, error)
	RenewAccess(context.Context, *RenewAccessRequest) (*RenewAccessResponse, error)
	// Updates a user's password
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	CreateFriendship(context.Context, *CreateFriendshipRequest) (*CreateFriendshipResponse, error)
	ListUserFriends(context.Context, *ListUserFriendsRequest) (*ListUserFriendsResponse, error)
	ListFriendRequests(context.Context, *ListFriendRequestsRequest) (*ListFriendRequestsResponse, error)
	AcceptFriendRequest(context.Context, *AcceptFriendRequestRequest) (*emptypb.Empty, error)
	CreateGame(context.Context, *CreateGameRequest) (*CreateGameResponse, error)
	InvitePlayerToGame(context.Context, *InvitePlayerToGameRequest) (*emptypb.Empty, error)
	ListGameInvites(context.Context, *ListGameInvitesRequest) (*ListGameInvitesResponse, error)
	AcceptGameInvite(context.Context, *AcceptGameInviteRequest) (*emptypb.Empty, error)
	//	rpc AddPlayerToGame (AddPlayerToGameRequest) returns (google.protobuf.Empty) {
	//	  option (google.api.http) = {
	//	      post: "/v1/add_player_to_game"
	//	      body: "*"
	//	  };
	//	  option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
	//	    description: "Use this API to add player to the game";
	//	    summary: "Add player to the game";
	//	  };
	//	};
	ListGamePlayers(context.Context, *ListGamePlayersRequest) (*ListGamePlayersResponse, error)
	GetPlayerGame(context.Context, *GetPlayerGameRequest) (*GetPlayerGameResponse, error)
	StartGame(context.Context, *StartGameRequest) (*StartGameResponse, error)
	GetPlayerHand(context.Context, *GetPlayerHandRequest) (*GetPlayerHandResponse, error)
	PlayDessert(context.Context, *PlayDessertRequest) (*PlayDessertResponse, error)
	DrawCard(context.Context, *DrawCardRequest) (*DrawCardResponse, error)
	mustEmbedUnimplementedDessertedServer()
}

// UnimplementedDessertedServer must be embedded to have forward compatible implementations.
type UnimplementedDessertedServer struct {
}

func (UnimplementedDessertedServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedDessertedServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedDessertedServer) Logout(context.Context, *LogoutRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedDessertedServer) CheckUserSession(context.Context, *CheckUserSessionRequest) (*CheckUserSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserSession not implemented")
}
func (UnimplementedDessertedServer) RenewAccess(context.Context, *RenewAccessRequest) (*RenewAccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenewAccess not implemented")
}
func (UnimplementedDessertedServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedDessertedServer) CreateFriendship(context.Context, *CreateFriendshipRequest) (*CreateFriendshipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFriendship not implemented")
}
func (UnimplementedDessertedServer) ListUserFriends(context.Context, *ListUserFriendsRequest) (*ListUserFriendsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserFriends not implemented")
}
func (UnimplementedDessertedServer) ListFriendRequests(context.Context, *ListFriendRequestsRequest) (*ListFriendRequestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFriendRequests not implemented")
}
func (UnimplementedDessertedServer) AcceptFriendRequest(context.Context, *AcceptFriendRequestRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptFriendRequest not implemented")
}
func (UnimplementedDessertedServer) CreateGame(context.Context, *CreateGameRequest) (*CreateGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGame not implemented")
}
func (UnimplementedDessertedServer) InvitePlayerToGame(context.Context, *InvitePlayerToGameRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvitePlayerToGame not implemented")
}
func (UnimplementedDessertedServer) ListGameInvites(context.Context, *ListGameInvitesRequest) (*ListGameInvitesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGameInvites not implemented")
}
func (UnimplementedDessertedServer) AcceptGameInvite(context.Context, *AcceptGameInviteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptGameInvite not implemented")
}
func (UnimplementedDessertedServer) ListGamePlayers(context.Context, *ListGamePlayersRequest) (*ListGamePlayersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGamePlayers not implemented")
}
func (UnimplementedDessertedServer) GetPlayerGame(context.Context, *GetPlayerGameRequest) (*GetPlayerGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerGame not implemented")
}
func (UnimplementedDessertedServer) StartGame(context.Context, *StartGameRequest) (*StartGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartGame not implemented")
}
func (UnimplementedDessertedServer) GetPlayerHand(context.Context, *GetPlayerHandRequest) (*GetPlayerHandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerHand not implemented")
}
func (UnimplementedDessertedServer) PlayDessert(context.Context, *PlayDessertRequest) (*PlayDessertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayDessert not implemented")
}
func (UnimplementedDessertedServer) DrawCard(context.Context, *DrawCardRequest) (*DrawCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DrawCard not implemented")
}
func (UnimplementedDessertedServer) mustEmbedUnimplementedDessertedServer() {}

// UnsafeDessertedServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DessertedServer will
// result in compilation errors.
type UnsafeDessertedServer interface {
	mustEmbedUnimplementedDessertedServer()
}

func RegisterDessertedServer(s grpc.ServiceRegistrar, srv DessertedServer) {
	s.RegisterService(&Desserted_ServiceDesc, srv)
}

func _Desserted_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DessertedServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Desserted_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DessertedServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Desserted_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DessertedServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Desserted_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DessertedServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Desserted_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DessertedServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Desserted_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DessertedServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Desserted_CheckUserSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUserSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DessertedServer).CheckUserSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Desserted_CheckUserSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DessertedServer).CheckUserSession(ctx, req.(*CheckUserSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Desserted_RenewAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenewAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DessertedServer).RenewAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Desserted_RenewAccess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DessertedServer).RenewAccess(ctx, req.(*RenewAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Desserted_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DessertedServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Desserted_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DessertedServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Desserted_CreateFriendship_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFriendshipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DessertedServer).CreateFriendship(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Desserted_CreateFriendship_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DessertedServer).CreateFriendship(ctx, req.(*CreateFriendshipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Desserted_ListUserFriends_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserFriendsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DessertedServer).ListUserFriends(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Desserted_ListUserFriends_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DessertedServer).ListUserFriends(ctx, req.(*ListUserFriendsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Desserted_ListFriendRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFriendRequestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DessertedServer).ListFriendRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Desserted_ListFriendRequests_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DessertedServer).ListFriendRequests(ctx, req.(*ListFriendRequestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Desserted_AcceptFriendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptFriendRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DessertedServer).AcceptFriendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Desserted_AcceptFriendRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DessertedServer).AcceptFriendRequest(ctx, req.(*AcceptFriendRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Desserted_CreateGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DessertedServer).CreateGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Desserted_CreateGame_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DessertedServer).CreateGame(ctx, req.(*CreateGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Desserted_InvitePlayerToGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvitePlayerToGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DessertedServer).InvitePlayerToGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Desserted_InvitePlayerToGame_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DessertedServer).InvitePlayerToGame(ctx, req.(*InvitePlayerToGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Desserted_ListGameInvites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGameInvitesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DessertedServer).ListGameInvites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Desserted_ListGameInvites_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DessertedServer).ListGameInvites(ctx, req.(*ListGameInvitesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Desserted_AcceptGameInvite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptGameInviteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DessertedServer).AcceptGameInvite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Desserted_AcceptGameInvite_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DessertedServer).AcceptGameInvite(ctx, req.(*AcceptGameInviteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Desserted_ListGamePlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGamePlayersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DessertedServer).ListGamePlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Desserted_ListGamePlayers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DessertedServer).ListGamePlayers(ctx, req.(*ListGamePlayersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Desserted_GetPlayerGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayerGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DessertedServer).GetPlayerGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Desserted_GetPlayerGame_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DessertedServer).GetPlayerGame(ctx, req.(*GetPlayerGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Desserted_StartGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DessertedServer).StartGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Desserted_StartGame_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DessertedServer).StartGame(ctx, req.(*StartGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Desserted_GetPlayerHand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayerHandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DessertedServer).GetPlayerHand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Desserted_GetPlayerHand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DessertedServer).GetPlayerHand(ctx, req.(*GetPlayerHandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Desserted_PlayDessert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayDessertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DessertedServer).PlayDessert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Desserted_PlayDessert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DessertedServer).PlayDessert(ctx, req.(*PlayDessertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Desserted_DrawCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DrawCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DessertedServer).DrawCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Desserted_DrawCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DessertedServer).DrawCard(ctx, req.(*DrawCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Desserted_ServiceDesc is the grpc.ServiceDesc for Desserted service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Desserted_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Desserted",
	HandlerType: (*DessertedServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Desserted_CreateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _Desserted_LoginUser_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Desserted_Logout_Handler,
		},
		{
			MethodName: "CheckUserSession",
			Handler:    _Desserted_CheckUserSession_Handler,
		},
		{
			MethodName: "RenewAccess",
			Handler:    _Desserted_RenewAccess_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Desserted_UpdateUser_Handler,
		},
		{
			MethodName: "CreateFriendship",
			Handler:    _Desserted_CreateFriendship_Handler,
		},
		{
			MethodName: "ListUserFriends",
			Handler:    _Desserted_ListUserFriends_Handler,
		},
		{
			MethodName: "ListFriendRequests",
			Handler:    _Desserted_ListFriendRequests_Handler,
		},
		{
			MethodName: "AcceptFriendRequest",
			Handler:    _Desserted_AcceptFriendRequest_Handler,
		},
		{
			MethodName: "CreateGame",
			Handler:    _Desserted_CreateGame_Handler,
		},
		{
			MethodName: "InvitePlayerToGame",
			Handler:    _Desserted_InvitePlayerToGame_Handler,
		},
		{
			MethodName: "ListGameInvites",
			Handler:    _Desserted_ListGameInvites_Handler,
		},
		{
			MethodName: "AcceptGameInvite",
			Handler:    _Desserted_AcceptGameInvite_Handler,
		},
		{
			MethodName: "ListGamePlayers",
			Handler:    _Desserted_ListGamePlayers_Handler,
		},
		{
			MethodName: "GetPlayerGame",
			Handler:    _Desserted_GetPlayerGame_Handler,
		},
		{
			MethodName: "StartGame",
			Handler:    _Desserted_StartGame_Handler,
		},
		{
			MethodName: "GetPlayerHand",
			Handler:    _Desserted_GetPlayerHand_Handler,
		},
		{
			MethodName: "PlayDessert",
			Handler:    _Desserted_PlayDessert_Handler,
		},
		{
			MethodName: "DrawCard",
			Handler:    _Desserted_DrawCard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_desserted.proto",
}
