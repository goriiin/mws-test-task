package repo

import (
	"mws/domain"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func setupTestRepo(t *testing.T) (*ProfileYAMLRepo, string) {
	t.Helper()
	tempDir := t.TempDir()
	r, err := NewProfileYAMLRepo(tempDir)
	require.NoError(t, err, "Инициализация репозитория не должна вызывать ошибку")
	return r, tempDir
}

func TestNewProfileYAMLRepo(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		path        string
		expectError bool
	}{
		{
			name:        "ошибка при пустом пути",
			path:        "",
			expectError: true,
		},
		{
			name:        "ошибка при некорректном пути с нулевым байтом",
			path:        "some\x00dir",
			expectError: true,
		},
		{
			name:        "успешное создание с корректным путем",
			path:        t.TempDir(),
			expectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			_, err := NewProfileYAMLRepo(tc.path)
			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestProfileYAMLRepo_Create(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		setup       func(t *testing.T) (*ProfileYAMLRepo, string)
		profile     domain.Profile
		expectErr   bool
		errContains string
	}{
		{
			name:      "успешное создание профиля",
			setup:     setupTestRepo,
			profile:   domain.Profile{Name: "test-profile", User: "test-user", Project: "test-project"},
			expectErr: false,
		},
		{
			name:        "ошибка при недопустимом имени - выход за пределы директории",
			setup:       setupTestRepo,
			profile:     domain.Profile{Name: "../../etc/passwd"},
			expectErr:   true,
			errContains: "недопустимое имя",
		},
		{
			name:        "ошибка при недопустимом имени - пустое имя",
			setup:       setupTestRepo,
			profile:     domain.Profile{Name: ""},
			expectErr:   true,
			errContains: "недопустимое имя",
		},
		{
			name: "ошибка при записи в не-записываемую директорию",
			setup: func(t *testing.T) (*ProfileYAMLRepo, string) {
				t.Helper()
				tempDir := t.TempDir()
				unwritableDir := filepath.Join(tempDir, "unwritable")
				require.NoError(t, os.Mkdir(unwritableDir, 0555))
				r, err := NewProfileYAMLRepo(unwritableDir)
				require.NoError(t, err)
				return r, unwritableDir
			},
			profile:     domain.Profile{Name: "any-profile"},
			expectErr:   true,
			errContains: "permission denied",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			r, tempDir := tc.setup(t)
			err := r.Create(tc.profile)

			if tc.expectErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.errContains)
			} else {
				require.NoError(t, err)
				expectedPath := filepath.Join(tempDir, tc.profile.Name+".yaml")
				yamlData, err := os.ReadFile(expectedPath)
				require.NoError(t, err)

				var createdData dtoProfile
				err = yaml.Unmarshal(yamlData, &createdData)
				require.NoError(t, err)
				assert.Equal(t, tc.profile.User, createdData.User)
				assert.Equal(t, tc.profile.Project, createdData.Project)
			}
		})
	}
}

func TestProfileYAMLRepo_Get(t *testing.T) {
	t.Parallel()

	r, _ := setupTestRepo(t)
	testProfile := domain.Profile{Name: "existing-profile", User: "user1", Project: "proj1"}
	require.NoError(t, r.Create(testProfile))

	testCases := []struct {
		name        string
		profileName string
		expectErr   bool
		errContains string
	}{
		{
			name:        "успешное получение существующего профиля",
			profileName: "existing-profile",
			expectErr:   false,
		},
		{
			name:        "ошибка при получении несуществующего профиля",
			profileName: "non-existing-profile",
			expectErr:   true,
			errContains: "не найден",
		},
		{
			name:        "ошибка при получении с недопустимым именем",
			profileName: "../secret",
			expectErr:   true,
			errContains: "недопустимое имя",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			profile, err := r.Get(tc.profileName)

			if tc.expectErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.errContains)
			} else {
				require.NoError(t, err)
				assert.Equal(t, testProfile, profile)
			}
		})
	}
}

func TestProfileYAMLRepo_Delete(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name        string
		profileName string
		setup       func(r *ProfileYAMLRepo)
		expectErr   bool
		errContains string
	}{
		{
			name:        "успешное удаление существующего профиля",
			profileName: "to-be-deleted",
			setup: func(r *ProfileYAMLRepo) {
				err := r.Create(domain.Profile{Name: "to-be-deleted", User: "u", Project: "p"})
				require.NoError(t, err)
			},
			expectErr: false,
		},
		{
			name:        "ошибка при удалении несуществующего профиля",
			profileName: "non-existing-profile",
			setup:       func(r *ProfileYAMLRepo) {},
			expectErr:   true,
			errContains: "не существует",
		},
		{
			name:        "ошибка при удалении с пустым именем",
			profileName: "",
			setup:       func(r *ProfileYAMLRepo) {},
			expectErr:   true,
			errContains: "недопустимое имя",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			r, tempDir := setupTestRepo(t)
			tc.setup(r)

			err := r.Delete(tc.profileName)

			if tc.expectErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.errContains)
			} else {
				require.NoError(t, err)
				filePath := filepath.Join(tempDir, tc.profileName+".yaml")
				_, err = os.Stat(filePath)
				assert.True(t, os.IsNotExist(err), "Файл должен быть удален")
			}
		})
	}
}

func TestProfileYAMLRepo_List(t *testing.T) {
	t.Parallel()
	r, tempDir := setupTestRepo(t)

	p1 := domain.Profile{Name: "profile1", User: "u1", Project: "p1"}
	p2 := domain.Profile{Name: "profile2", User: "u2", Project: "p2"}
	require.NoError(t, r.Create(p1))
	require.NoError(t, r.Create(p2))
	require.NoError(t, os.WriteFile(filepath.Join(tempDir, "notes.txt"), []byte("text"), 0644))
	require.NoError(t, os.Mkdir(filepath.Join(tempDir, "some-dir"), 0755))

	profiles, err := r.List()

	assert.NoError(t, err)
	assert.Len(t, profiles, 2, "Должно быть найдено ровно 2 профиля")
	assert.ElementsMatch(t, []domain.Profile{p1, p2}, profiles)
}
