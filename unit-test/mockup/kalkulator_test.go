package mockup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func RunTestKalkulator(t *testing.T, k Kalkulator) {
	rTambah := k.Tambah(5, 2)
	rKurang := k.Kurang(5, 1)
	rKali := k.Kali(5, 3)
	assert.Equal(t, 7, rTambah)
	assert.Equal(t, 4, rKurang)
	assert.Equal(t, 15, rKali)
}
func MockUpSetUp(m *MockOprasi){
	m.On("Tambah", 5, 2).Return(7)
	m.On("Kurang", 5, 1).Return(4)
	m.On("Kali", 5, 3).Return(15)
}
func TestKalkulator(t *testing.T) {
	mockUp := new(MockOprasi)
	MockUpSetUp(mockUp)
	//mengetes semua method di MockOprasi (mockup dari oprasi)
	RunTestKalkulator(t, mockUp)
	//mengecek apakah semua test mock ekspektasi nya terpenuhi
	mockUp.AssertExpectations(t)
	//mengetes semua method di Oprasi
	RunTestKalkulator(t, &Oprasi{})

}
