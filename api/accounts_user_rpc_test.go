package api

import (
	"context"
	"fmt"
	"testing"

	"connectrpc.com/connect"
	"github.com/esolveeg/cms-api/pkg/random"
	"github.com/esolveeg/cms-api/proto_gen/devkit/v1"
	"github.com/stretchr/testify/require"
)

var (
	userEmail       = random.RandomEmail()
	userPassword    = random.RandomString(8)
	userNewPassword = random.RandomString(10)
	userPhone       = random.RandomPhone()
	userNewPhone    = random.RandomPhone()
	loginRequest    = connect.NewRequest(&devkitv1.UserLoginRequest{
		LoginCode:    userEmail,
		UserPassword: userPassword,
	})

	emptyRequest      = connect.NewRequest(&devkitv1.RolesListRequest{})
	userCreateRequest = connect.NewRequest(&devkitv1.UserCreateUpdateRequest{
		UserId:            0,
		UserEmail:         userEmail,
		UserPassword:      userPassword,
		UserPhone:         userPhone,
		UserTypeId:        1,
		UserSecurityLevel: 1,
		UserName:          random.RandomName(),
	})
)

func TestCycle(t *testing.T) {

	ctx := context.Background()
	t.Run("Login not existed user", func(t *testing.T) {
		_, err := realDbApi.UserLogin(ctx, loginRequest)
		require.NotEmpty(t, err)
		require.Contains(t, err.Error(), "400")
		require.Contains(t, err.Error(), "Invalid login credentials")
	})
	// // create new user
	t.Run("create new user with no roles", func(t *testing.T) {
		resp, err := realDbApi.UserCreateUpdate(ctx, userCreateRequest)
		require.NoError(t, err)
		userCreateRequest.Msg.UserId = resp.Msg.User.UserId

		require.Equal(t, resp.Msg.User.UserName, userCreateRequest.Msg.UserName)
		require.Equal(t, resp.Msg.User.UserEmail, userCreateRequest.Msg.UserEmail)
		require.Equal(t, resp.Msg.User.UserPhone, userCreateRequest.Msg.UserPhone)
	})
	// // login with wrong password
	t.Run("Login with wrong password", func(t *testing.T) {
		wrongLoginRequest := connect.NewRequest(&devkitv1.UserLoginRequest{
			LoginCode:    userEmail,
			UserPassword: "wrongPassword",
		})

		_, err := realDbApi.UserLogin(ctx, wrongLoginRequest)
		require.NotEmpty(t, err)
		require.Contains(t, err.Error(), "400")
		require.Contains(t, err.Error(), "Invalid login credentials")
	})
	// // login with wrong password
	t.Run("Login with correct password", func(t *testing.T) {
		resp, err := realDbApi.UserLogin(ctx, loginRequest)
		require.NoError(t, err)
		require.NotEmpty(t, resp.Msg.LoginInfo.AccessToken)
		require.Equal(t, resp.Msg.User.UserName, userCreateRequest.Msg.UserName)
		require.Equal(t, resp.Msg.User.UserEmail, userCreateRequest.Msg.UserEmail)
		require.Equal(t, resp.Msg.User.UserPhone, userCreateRequest.Msg.UserPhone)
	})
	// update the created user password
	t.Run("update the password of created user", func(t *testing.T) {
		userCreateRequest.Msg.UserPassword = userNewPassword
		resp, err := realDbApi.UserCreateUpdate(ctx, userCreateRequest)
		require.NoError(t, err)
		require.Equal(t, resp.Msg.User.UserName, userCreateRequest.Msg.UserName)
		require.Equal(t, resp.Msg.User.UserEmail, userCreateRequest.Msg.UserEmail)
		require.Equal(t, resp.Msg.User.UserPhone, userCreateRequest.Msg.UserPhone)
	})
	// login with the old password
	t.Run("Login with old password", func(t *testing.T) {
		_, err := realDbApi.UserLogin(ctx, loginRequest)
		require.NotEmpty(t, err)
		require.Contains(t, err.Error(), "400")
		require.Contains(t, err.Error(), "Invalid login credentials")
	})
	// login with the updated password
	t.Run("Login with updated password", func(t *testing.T) {
		loginRequest.Msg.UserPassword = userNewPassword
		resp, err := realDbApi.UserLogin(ctx, loginRequest)
		require.NoError(t, err)
		emptyRequest.Header().Add("Authorization", fmt.Sprintf("bearer %s", resp.Msg.LoginInfo.AccessToken))
		require.NotEmpty(t, resp.Msg.LoginInfo.AccessToken)
		require.Equal(t, resp.Msg.User.UserName, userCreateRequest.Msg.UserName)
		require.Equal(t, resp.Msg.User.UserEmail, userCreateRequest.Msg.UserEmail)
		require.Equal(t, resp.Msg.User.UserPhone, userCreateRequest.Msg.UserPhone)
	})
	// try to access ednpoint that needs access
	t.Run("forbidden access", func(t *testing.T) {
		_, err := realDbApi.RolesList(ctx, emptyRequest)
		require.NotEmpty(t, err)
		require.Contains(t, err.Error(), "don't have permission")

	})
	// update the created user  roles
	t.Run("update the roles of created user", func(t *testing.T) {
		userCreateRequest.Msg.Roles = []int32{1}
		resp, err := realDbApi.UserCreateUpdate(ctx, userCreateRequest)
		require.NoError(t, err)
		require.Equal(t, resp.Msg.User.UserName, userCreateRequest.Msg.UserName)
		require.Equal(t, resp.Msg.User.UserEmail, userCreateRequest.Msg.UserEmail)
		require.Equal(t, resp.Msg.User.UserPhone, userCreateRequest.Msg.UserPhone)
	})
	// // login with the updated roles
	// t.Run("Login with updated roles", func(t *testing.T) {
	// 	resp, err := realDbApi.UserLogin(ctx, loginRequest)
	// 	require.NoError(t, err)
	// 	emptyRequest.Header().Add("Authorization", fmt.Sprintf("bearer %s", resp.Msg.LoginInfo.AccessToken))
	// 	require.NotEmpty(t, resp.Msg.LoginInfo.AccessToken)
	// 	require.Equal(t, resp.Msg.User.UserName, userCreateRequest.Msg.UserName)
	// 	require.Equal(t, resp.Msg.User.UserEmail, userCreateRequest.Msg.UserEmail)
	// 	require.Equal(t, resp.Msg.User.UserPhone, userCreateRequest.Msg.UserPhone)
	// })
	// // try to recall the endpoint after permissions added
	// t.Run("granted access", func(t *testing.T) {
	// 	_, err := realDbApi.RolesList(ctx, emptyRequest)
	// 	require.NoError(t, err)
	//
	// })
	//
}
