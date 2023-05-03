package mocks

import mock "github.com/stretchr/testify/mock"

type MockBcrypt struct {
	mock.Mock
}

func (m *MockBcrypt) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	args := m.Called(password, cost)
	return args.Get(0).([]byte), args.Error(1)
}
