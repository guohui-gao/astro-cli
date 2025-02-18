package software

import (
	"testing"

	"github.com/astronomer/astro-cli/houston"
	mocks "github.com/astronomer/astro-cli/houston/mocks"
	testUtil "github.com/astronomer/astro-cli/pkg/testing"
	"github.com/stretchr/testify/assert"
)

func TestDeploymentUserAddCommand(t *testing.T) {
	testUtil.InitTestConfig(testUtil.SoftwarePlatform)
	expectedOut := ` DEPLOYMENT ID                 USER                       ROLE                  
 ckggvxkw112212kc9ebv8vu6p     somebody@astronomer.io     DEPLOYMENT_VIEWER     

 Successfully added somebody@astronomer.io as a DEPLOYMENT_VIEWER
`
	expectedAddUserRequest := houston.UpdateDeploymentUserRequest{
		Email:        mockDeploymentUserRole.User.Username,
		Role:         mockDeploymentUserRole.Role,
		DeploymentID: mockDeploymentUserRole.Deployment.ID,
	}

	api := new(mocks.ClientInterface)
	api.On("GetAppConfig").Return(mockAppConfig, nil)
	api.On("AddDeploymentUser", expectedAddUserRequest).Return(mockDeploymentUserRole, nil)

	houstonClient = api
	output, err := execDeploymentCmd(
		"user",
		"add",
		"--deployment-id="+mockDeploymentUserRole.Deployment.ID,
		"--email="+mockDeploymentUserRole.User.Username,
	)
	assert.NoError(t, err)
	assert.Equal(t, expectedOut, output)
}

func TestDeploymentUserDeleteCommand(t *testing.T) {
	testUtil.InitTestConfig(testUtil.SoftwarePlatform)
	expectedOut := ` DEPLOYMENT ID                 USER                       ROLE                  
 ckggvxkw112212kc9ebv8vu6p     somebody@astronomer.io     DEPLOYMENT_VIEWER     

 Successfully removed the DEPLOYMENT_VIEWER role for somebody@astronomer.io from deployment ckggvxkw112212kc9ebv8vu6p
`

	api := new(mocks.ClientInterface)
	api.On("GetAppConfig").Return(mockAppConfig, nil)
	api.On("DeleteDeploymentUser", mockDeploymentUserRole.Deployment.ID, mockDeploymentUserRole.User.Username).
		Return(mockDeploymentUserRole, nil)

	houstonClient = api
	output, err := execDeploymentCmd(
		"user",
		"remove",
		"--deployment-id="+mockDeploymentUserRole.Deployment.ID,
		mockDeploymentUserRole.User.Username,
	)
	assert.NoError(t, err)
	assert.Equal(t, expectedOut, output)
}

func TestDeploymentUserList(t *testing.T) {
	testUtil.InitTestConfig(testUtil.SoftwarePlatform)
	mockUser := []houston.DeploymentUser{
		{
			ID:           "test-id",
			Emails:       []houston.Email{{Address: "test-email"}},
			FullName:     "test-name",
			RoleBindings: []houston.RoleBinding{{Role: houston.DeploymentViewerRole, Deployment: houston.Deployment{ID: "test-id"}}},
		},
	}
	api := new(mocks.ClientInterface)
	api.On("ListDeploymentUsers", houston.ListDeploymentUsersRequest{UserID: "test-user-id", Email: "test-email", FullName: "test-name", DeploymentID: "test-id"}).Return(mockUser, nil).Once()

	houstonClient = api
	output, err := execDeploymentCmd("user", "list", "--deployment-id", "test-id", "-u", "test-user-id", "-e", "test-email", "-n", "test-name")
	assert.NoError(t, err)
	assert.Contains(t, output, "test-id")
	assert.Contains(t, output, "test-name")
	api.AssertExpectations(t)
}

func TestDeploymentUserUpdateCommand(t *testing.T) {
	testUtil.InitTestConfig(testUtil.SoftwarePlatform)
	expectedNewRole := houston.DeploymentAdminRole
	expectedOut := `Successfully updated somebody@astronomer.io to a ` + expectedNewRole
	mockResponseUserRole := *mockDeploymentUserRole
	mockResponseUserRole.Role = expectedNewRole

	expectedUpdateUserRequest := houston.UpdateDeploymentUserRequest{
		Email:        mockResponseUserRole.User.Username,
		Role:         expectedNewRole,
		DeploymentID: mockDeploymentUserRole.Deployment.ID,
	}

	api := new(mocks.ClientInterface)
	api.On("GetAppConfig").Return(mockAppConfig, nil)
	api.On("UpdateDeploymentUser", expectedUpdateUserRequest).Return(&mockResponseUserRole, nil)

	houstonClient = api
	output, err := execDeploymentCmd(
		"user",
		"update",
		"--deployment-id="+mockResponseUserRole.Deployment.ID,
		"--role="+expectedNewRole,
		mockResponseUserRole.User.Username,
	)
	assert.NoError(t, err)
	assert.Contains(t, output, expectedOut)
}
