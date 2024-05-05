package mockup

import "github.com/stretchr/testify/mock"

type MockOprasi struct{
	mock.Mock
}

func (o *MockOprasi) Tambah(a,b int) int {
	args := o.Called(a,b)
	return args.Int(0)
}
func (o *MockOprasi) Kurang(a,b int) int {
	args := o.Called(a,b)
	return args.Int(0)
}
func (o *MockOprasi) Kali(a,b int) int {
	args := o.Called(a,b)
	return args.Int(0)
}
