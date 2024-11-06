package api

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/esolveeg/cms-api/db"
	mockdb "github.com/esolveeg/cms-api/db/mock"
	"github.com/esolveeg/cms-api/pkg/random"
	"github.com/esolveeg/cms-api/proto_gen/devkit/v1"
	"github.com/golang/mock/gomock"
)

type roleCreateUpdateTest struct {
	name       string
	params     *devkitv1.RoleCreateUpdateRequest
	buildStubs func(store *mockdb.MockStore)
	expectErr  bool
}

func getValidRole() *devkitv1.RoleCreateUpdateRequest {
	return &devkitv1.RoleCreateUpdateRequest{
		RoleName:        random.RandomName(),
		RoleDescription: random.RandomString(50),
		Permissions:     []int32{1, 2, 3},
	}
}
func TestRoleCreateUpdate(t *testing.T) {
	validRole := getValidRole()
	// Define a slice of test cases
	testcases := []roleCreateUpdateTest{
		// Test for a valid role creation.
		{
			name: "ValidRole",
			params: &devkitv1.RoleCreateUpdateRequest{
				RoleName:        validRole.RoleName,
				RoleDescription: validRole.RoleDescription,
				Permissions:     validRole.Permissions,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.RoleCreateUpdateParams{
					RoleName:        validRole.RoleName,
					RoleDescription: validRole.RoleDescription,
					Permissions:     validRole.Permissions,
				}
				store.EXPECT().
					RoleCreateUpdate(gomock.Any(), arg).
					Times(1).
					Return(db.AccountsSchemaRole{
						RoleID:          1,
						RoleName:        validRole.RoleName,
						RoleDescription: db.StringToPgtext(validRole.RoleDescription),
					}, nil)
			},
			expectErr: false,
		},
		{
			name: "ValidRoleUpdate",
			params: &devkitv1.RoleCreateUpdateRequest{
				RoleId:          1,
				RoleName:        "updated role name",
				RoleDescription: validRole.RoleDescription,
				Permissions:     validRole.Permissions,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.RoleCreateUpdateParams{
					RoleID:          1,
					RoleName:        "updated role name",
					RoleDescription: validRole.RoleDescription,
					Permissions:     validRole.Permissions,
				}
				store.EXPECT().
					RoleCreateUpdate(gomock.Any(), arg).
					Times(1).
					Return(db.AccountsSchemaRole{
						RoleID:          1,
						RoleName:        "updated role name",
						RoleDescription: db.StringToPgtext(validRole.RoleDescription),
					}, nil)
			},
			expectErr: false,
		},

		{
			name: "InValidNameToShort",
			params: &devkitv1.RoleCreateUpdateRequest{
				RoleName:        random.RandomString(1),
				RoleDescription: validRole.RoleDescription,
				Permissions:     validRole.Permissions,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					RoleCreateUpdate(gomock.Any(), gomock.Any()).
					Times(0)

			},
			expectErr: true,
		},
		{
			name: "InValidNameToLong",
			params: &devkitv1.RoleCreateUpdateRequest{
				RoleName:        random.RandomString(220),
				RoleDescription: validRole.RoleDescription,
				Permissions:     validRole.Permissions,
			},
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					RoleCreateUpdate(gomock.Any(), gomock.Any()).
					Times(0)

			},
			expectErr: true,
		},
		{
			name: "InValidDescriptionToLong",
			params: &devkitv1.RoleCreateUpdateRequest{
				RoleName:        random.RandomString(120),
				RoleDescription: random.RandomString(220),
				Permissions:     validRole.Permissions,
			},
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					RoleCreateUpdate(gomock.Any(), gomock.Any()).
					Times(0)

			},
			expectErr: true,
		},

		{
			name: "InvalideDuplicatedPermissions",
			params: &devkitv1.RoleCreateUpdateRequest{
				RoleName:        random.RandomString(120),
				RoleDescription: random.RandomString(22),
				Permissions:     []int32{1, 1},
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					RoleCreateUpdate(gomock.Any(), gomock.Any()).
					Times(0)

			},
			expectErr: true,
		},
	}

	// Loop through the test cases and test each one
	// ctx := context.Background()
	storeCtrl := gomock.NewController(t)
	defer storeCtrl.Finish()
	store := mockdb.NewMockStore(storeCtrl)
	api := newTestApi(store)
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			tc.buildStubs(store)
			createdRole, err :=
				api.RoleCreateUpdate(context.Background(), connect.NewRequest(tc.params))
			// If the current test case expects an error and no error occurred, fail the test
			if tc.expectErr && err == nil {
				t.Errorf("Expected an error but got none %s", tc.name)
			}

			// If the current test case does not expect an error and an error occurred, fail the test
			if !tc.expectErr && err != nil {
				t.Errorf("Expected no error but got %v", err)
			}
			if !tc.expectErr {
				if createdRole.Msg.Role.RoleName != tc.params.RoleName {
					t.Errorf("un expected name wanted %s got %s", createdRole.Msg.Role.RoleName, tc.params.RoleName)
				}

			}

		})
	}
}
