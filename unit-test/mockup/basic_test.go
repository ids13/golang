package mockup

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MyMockedObject adalah objek yang di-mock yang mengimplementasikan sebuah interface.
type MyMockedObject struct {
	mock.Mock
}

// DoSomething adalah metode pada MyMockedObject yang melakukan sesuatu.
// Di sini kita hanya mem-stub metode ini karena ini adalah objek yang di-mock.
func (m *MyMockedObject) DoSomething(number int) (bool, error) {
	//karena parameternya hanya satu "number",jadi di masukan di m.called hanya satu, jika ada banyak masukan semuanya dan pisah kan dengan koma 
	args := m.Called(number)
	//args.bool dan args.error ini mewakili parameter return ketika kita menetapkan 
	//expektasi untuk mock nnya testObj.On("DoSomething", 123).Return(true, nil)
	//nomor index nya mewakili posisi parameter ketika di passing.dari kiri ke kanan, di mulai dari index ke nol
	//contohnya returnnya kan (bool,error) . boolnya index ke 0 dan errornya index ke 1
	//contoh lainnya
	//args.Get(0) . ini akan mengembalikan interface kosong, jadi harus di casitng dulu ,salah satu kegunaanya jika data returnya ingin berbentuk struct. jadi perlu di casitng dulu args.Get(0).(StructAnda)
	//args.String(0) ini berguna jika data returnnya berupa string
	//args.Int(0) ini berguna jika data returnnya berupa int
	return args.Bool(0), args.Error(1)
}

// targetFuncThatDoesSomethingWithObj adalah fungsi yang menggunakan objek yang di-mock.
func targetFuncThatDoesSomethingWithObj(obj *MyMockedObject) bool {
	// Memanggil metode DoSomething pada objek yang diberikan
	result, _ := obj.DoSomething(123)

	// Mengembalikan hasil panggilan metode
	return result
}

// TestSomething adalah contoh penggunaan objek mock untuk menguji kode yang kita uji.
func TestSomething(t *testing.T) {
	// Membuat instance dari objek mock
	testObj := new(MyMockedObject)

	// Menetapkan ekspektasi
	
	testObj.On("DoSomething", 123).Return(true, nil)

	// Memanggil kode yang kita uji
	result := targetFuncThatDoesSomethingWithObj(testObj)

	// Memastikan bahwa ekspektasi terpenuhi, contohnya jika ekspekatasi tidak di tetapkan,maka ini akan memicu error
	testObj.AssertExpectations(t)

	// Memastikan bahwa hasil yang dikembalikan adalah yang diharapkan
	assert.True(t, result, "Expected result to be true")
}
