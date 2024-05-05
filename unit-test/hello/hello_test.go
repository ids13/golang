package hello

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSay(t *testing.T) {
	result := Say("naga")
	if result != "hallo naga" {
		panic("hasil tidak sesuai ekspektasi")
	}
}
func TestSay_Gagal(t *testing.T) {
	result := Say("naga")
	if result != "hallo naga" {
		// panic("hasil tidak sesuai ekspektasi")
		// t.Fail()
		// t.FailNow()
		// t.Error("hasil tidak sama")
		t.Fatal("fatal hasil tidak sama")
	}
	println("tes berakhir")
}

func TestSay_Skip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip()
	}
	
	println("ini sekipt test")
}

func TestSay_Assert(t *testing.T) {
	result := Say("naga")
	assert.Equal(t, "hallo nana", result, "hasil tidak sama")
	println("assert selesai")
}
func TestSay_Require(t *testing.T) {
	result := Say("naga")
	require.Equal(t, "hallo nana", result, "hasil tidak sama")
	//ini tidak akan di eksekusi karena require memanggil failnow
	println("require selesai")
}

func TestSay_Subtest(t *testing.T)  {
	t.Run("naga",func (t *testing.T){
		result := Say("naga")
		if result != "hallo naga" {
			panic("hasil tidak sesuai ekspektasi")
		}	
	})
	t.Run("nana",func (t *testing.T){
		result := Say("nana")
		if result != "hallo nana" {
			panic("hasil tidak sesuai ekspektasi")
		}	
	})
}



func TestSay_table(t *testing.T)  {
	data := []struct{
		nama string
		request string
		expect string
	}{
		{
			nama : "naga",
			request : "naga",
			expect : "hallo naga",
		},
		{
			nama : "nana",
			request : "nana",
			expect : "hallo nana",
		},
	}

	for _,val := range data{
		t.Run(val.nama,func (t *testing.T){
		result := Say(val.request)
		assert.Equal(t,val.expect,result,"tidak sesuai")
		})
	}
}

func BenchmarkSay_bench(b *testing.B){
	for i := 0;i<b.N;i++{
		Say("naga")
	}
}