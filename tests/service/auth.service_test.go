package service

import (
	"chapapp-backend-api/internal/dto"
	"chapapp-backend-api/internal/entity"

	"github.com/stretchr/testify/mock"
)

type MockAccountRepo struct {
	mock.Mock
}

func (m *MockAccountRepo) GetUserByEmail(email string) (entity.Account, error) {
	args := m.Called(email)
	return args.Get(0).(entity.Account), args.Error(1)
}
func (m *MockAccountRepo) GetUserByUsername(username string) (entity.Account, error) {
	args := m.Called(username)
	return args.Get(0).(entity.Account), args.Error(1)
}
func (m *MockAccountRepo) Create(acc entity.Account) (entity.Account, error) {
	args := m.Called(acc)
	return args.Get(0).(entity.Account), args.Error(1)
}
func (m *MockAccountRepo) Update(acc entity.Account) (entity.Account, error) {
	args := m.Called(acc)
	return args.Get(0).(entity.Account), args.Error(1)
}
func (m *MockAccountRepo) GetList(input dto.GetListAccountInputDTO) ([]entity.Account, error) {
	args := m.Called(input)
	return args.Get(0).([]entity.Account), args.Error(1)
}
func (m *MockAccountRepo) GetRandomFive(id string) ([]entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).([]entity.Account), args.Error(1)
}
func (m *MockAccountRepo) GetListBan(input dto.GetListBanInputDTO) ([]entity.Account, error) {
	args := m.Called(input)
	return args.Get(0).([]entity.Account), args.Error(1)
}
func (m *MockAccountRepo) GetUserByAccountId(accountID string) (entity.Account, error) {
	args := m.Called(accountID)
	return args.Get(0).(entity.Account), args.Error(1)
}

// type MockAuthRepo struct {
// 	mock.Mock
// }

// func (m *MockAuthRepo) GetOTP(key string) (string, int, error) {
// 	args := m.Called(key)
// 	return args.String(0), args.Int(1), args.Error(2)
// }
// func (m *MockAuthRepo) AddOTP(key string, otp int, ttl int64) error {
// 	args := m.Called(key, otp, ttl)
// 	return args.Error(0)
// }

// func (m *MockAuthRepo) RemoveOTP(key string) error {
// 	args := m.Called(key)
// 	return args.Error(0)
// }

// func (m *MockAuthRepo) CanSendOTP(key string) (bool, error) {
// 	args := m.Called(key)
// 	return args.Bool(0), args.Error(1)
// }

// func TestLogin_Success(t *testing.T) {
// 	accountRepo := new(MockAccountRepo)
// 	authRepo := new(MockAuthRepo)
// 	authService := service.NewAuthService(accountRepo, authRepo)

// 	mockAccount := entity.Account{
// 		BaseEntity: entity.BaseEntity{
// 			ID:        1,
// 			CreatedAt: time.Now(),
// 			UpdatedAt: time.Now(),
// 		},
// 		Email:    "test@example.com",
// 		Username: "tester",
// 		Password: "$2a$10$validHashedPassword", // giả lập bcrypt hash
// 		Roles:    []entity.Role{{Rolename: "USER"}},
// 	}

// 	accountRepo.On("GetUserByUsername", "tester").Return(mockAccount, nil)

// 	// giả lập utils.CheckPassword và utils.GenerateAccessToken/RefreshToken (có thể mock utils nếu muốn)
// 	dtoInput := dto.LoginInputDTO{
// 		Username: "tester",
// 		Password: "123456", // giả định đúng
// 	}

// 	// Override check password và token cho test nếu cần

// 	output, err := authService.Login(dtoInput)

// 	assert.NoError(t, err)
// 	assert.Equal(t, "1", output.Id)
// 	assert.NotEmpty(t, output.AccessToken)
// 	assert.NotEmpty(t, output.RefreshToken)

// 	accountRepo.AssertExpectations(t)
// }
