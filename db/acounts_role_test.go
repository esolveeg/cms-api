package db

import (
	"context"
	"testing"

	"github.com/esolveeg/cms-api/pkg/random"
	"github.com/stretchr/testify/require"
)

func TestRoleCreateUpdate(t *testing.T) {
	// Define test input parameters.
	validName := random.RandomName()
	testcases := []struct {
		name             string
		params           *RoleCreateUpdateParams
		expectedErrorMsg string
		expectErr        bool
	}{
		{
			name: "ValidRole",
			params: &RoleCreateUpdateParams{
				RoleName:        validName,
				RoleDescription: random.RandomString(50),
				Permissions:     []int32{1, 2, 3},
			},
			expectErr: false,
		},
		{
			name: "ValidRoleUpdate",
			params: &RoleCreateUpdateParams{
				RoleID:          1,
				RoleName:        random.RandomName(),
				RoleDescription: random.RandomString(50),
				Permissions:     []int32{1, 2, 3},
			},
			expectErr: false,
		},
		{
			name: "DuplicatedRoleName",
			params: &RoleCreateUpdateParams{
				RoleName:        validName,
				RoleDescription: random.RandomString(50),
				Permissions:     []int32{1, 2, 3},
			},
			expectErr: true,
		},

		{
			name: "RoleNameTooLong",
			params: &RoleCreateUpdateParams{
				RoleName:        random.RandomString(220), // Exceeds max length
				RoleDescription: random.RandomString(50),
				Permissions:     []int32{1, 2, 3},
			},
			expectErr: true,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the RoleCreateUpdate function with the test parameters
			role, err := store.RoleCreateUpdate(context.Background(), *tc.params)

			if tc.expectErr {
				require.Error(t, err)
				require.Empty(t, role) // Expect no role to be returned on error
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, role)
				require.Equal(t, tc.params.RoleName, role.RoleName)
				require.Equal(t, tc.params.RoleDescription, role.RoleDescription.String)
			}
		})
	}

}
