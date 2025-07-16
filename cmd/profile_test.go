package cmd

import (
	"bytes"
	"errors"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"mws/cmd/mock"
	"mws/domain"
	"testing"
)

func executeCommand(root *cobra.Command, args ...string) (string, error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	err := root.Execute()

	return buf.String(), err
}

func TestProfileCommand(t *testing.T) {
	// Эта функция-хелпер остается без изменений
	executeCommand := func(root *cobra.Command, args ...string) (string, error) {
		buf := new(bytes.Buffer)
		root.SetOut(buf)
		root.SetErr(buf)
		root.SetArgs(args)
		err := root.Execute()
		return buf.String(), err
	}

	t.Run("NoSubcommand", func(t *testing.T) {
		output, err := executeCommand(rootCmd, "profile")
		require.NoError(t, err)
		assert.Contains(t, output, "Необходимо указать подкоманду")
	})
}

func TestListCommand(t *testing.T) {
	testCases := []struct {
		name        string
		setupMock   func(mockRepo *mock.MockProfileRepo)
		expectedOut string
		expectedErr error
		expectErr   bool
		args        []string
	}{
		{
			name: "Success",
			setupMock: func(mockRepo *mock.MockProfileRepo) {
				mockRepo.EXPECT().List().Return([]domain.Profile{
					{Name: "profile1", User: "user1", Project: "project1"},
					{Name: "profile2", User: "user2", Project: "project2"},
				}, nil)
			},
			expectedOut: "name: profile1\n\tuser: user1\n\tproject: project1\nname: profile2\n\tuser: user2\n\tproject: project2\n",
			expectErr:   false,
			args:        []string{"profile", "list"},
		},
		{
			name: "Error",
			setupMock: func(mockRepo *mock.MockProfileRepo) {
				mockRepo.EXPECT().List().Return(nil, errors.New("не удалось прочитать директорию"))
			},
			expectedErr: errors.New("не удалось прочитать директорию"),
			expectErr:   true,
			args:        []string{"profile", "list"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mock.NewMockProfileRepo(ctrl)
			tc.setupMock(mockRepo)
			NewProfileRepo(mockRepo)

			output, err := executeCommand(rootCmd, tc.args...)

			if tc.expectErr {
				require.Error(t, err)
				assert.Contains(t, output, tc.expectedErr.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.expectedOut, output)
			}
		})
	}
}

func TestGetCommand(t *testing.T) {
	testCases := []struct {
		name        string
		setupMock   func(mockRepo *mock.MockProfileRepo)
		args        []string
		expectedOut string
		expectedErr error
		expectErr   bool
	}{
		{
			name: "Success",
			setupMock: func(mockRepo *mock.MockProfileRepo) {
				expectedProfile := domain.Profile{Name: "my-profile", User: "test-user", Project: "test-proj"}
				mockRepo.EXPECT().Get("my-profile").Return(expectedProfile, nil)
			},
			args:        []string{"profile", "get", "--name=my-profile"},
			expectedOut: "name: my-profile\n\tuser: test-user\n\tproject: test-proj\n",
			expectErr:   false,
		},
		{
			name: "NotFound",
			setupMock: func(mockRepo *mock.MockProfileRepo) {
				mockRepo.EXPECT().Get("not-found").Return(domain.Profile{}, errors.New("профиль 'not-found' не найден"))
			},
			args:        []string{"profile", "get", "--name=not-found"},
			expectedErr: errors.New("профиль 'not-found' не найден"),
			expectErr:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockRepo := mock.NewMockProfileRepo(ctrl)
			tc.setupMock(mockRepo)
			NewProfileRepo(mockRepo)

			output, err := executeCommand(rootCmd, tc.args...)

			if tc.expectErr {
				require.Error(t, err)
				assert.Contains(t, output, tc.expectedErr.Error())
			} else {
				require.NoError(t, err)
				assert.Contains(t, output, tc.expectedOut)
			}
		})
	}
}

func TestCreateCommand(t *testing.T) {
	profileToCreate := domain.Profile{
		Name:    "new-one",
		User:    "new-user",
		Project: "new-proj",
	}

	testCases := []struct {
		name        string
		setupMock   func(mockRepo *mock.MockProfileRepo)
		args        []string
		expectedErr error
		expectErr   bool
	}{
		{
			name: "Success",
			setupMock: func(mockRepo *mock.MockProfileRepo) {
				mockRepo.EXPECT().Create(profileToCreate).Return(nil)
			},
			args:      []string{"profile", "create", "--name=new-one", "--user=new-user", "--project=new-proj"},
			expectErr: false,
		},
		{
			name: "Error",
			setupMock: func(mockRepo *mock.MockProfileRepo) {
				mockRepo.EXPECT().Create(gomock.Any()).Return(errors.New("профиль уже существует"))
			},
			args:        []string{"profile", "create", "--name=new-one", "--user=new-user", "--project=new-proj"},
			expectedErr: errors.New("профиль уже существует"),
			expectErr:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockRepo := mock.NewMockProfileRepo(ctrl)
			tc.setupMock(mockRepo)
			NewProfileRepo(mockRepo)

			output, err := executeCommand(rootCmd, tc.args...)

			if tc.expectErr {
				require.Error(t, err)
				assert.Contains(t, output, tc.expectedErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestDeleteCommand(t *testing.T) {
	testCases := []struct {
		name        string
		setupMock   func(mockRepo *mock.MockProfileRepo)
		args        []string
		expectedErr error
		expectErr   bool
	}{
		{
			name: "Success",
			setupMock: func(mockRepo *mock.MockProfileRepo) {
				mockRepo.EXPECT().Delete("my-profile").Return(nil)
			},
			args:      []string{"profile", "delete", "--name=my-profile"},
			expectErr: false,
		},
		{
			name: "Error",
			setupMock: func(mockRepo *mock.MockProfileRepo) {
				mockRepo.EXPECT().Delete("my-profile").Return(errors.New("не удалось удалить"))
			},
			args:        []string{"profile", "delete", "--name=my-profile"},
			expectedErr: errors.New("не удалось удалить"),
			expectErr:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockRepo := mock.NewMockProfileRepo(ctrl)
			tc.setupMock(mockRepo)
			NewProfileRepo(mockRepo)

			output, err := executeCommand(rootCmd, tc.args...)

			if tc.expectErr {
				require.Error(t, err)
				assert.Contains(t, output, tc.expectedErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMarkFlagRequiredError(t *testing.T) {
	testCases := []struct {
		name     string
		flagName string
	}{
		{
			name:     "несуществующий флаг",
			flagName: "non-existent-flag",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testCmd := &cobra.Command{}
			err := testCmd.MarkFlagRequired(tc.flagName)
			require.Error(t, err, "Должна быть ошибка при попытке пометить несуществующий флаг как обязательный")
		})
	}
}
