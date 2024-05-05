package mockup

type Oprasi struct{}

func (o *Oprasi) Tambah(a,b int) int {
	return a + b
}
func (o *Oprasi) Kurang(a,b int) int {
	return a - b
}
func (o *Oprasi) Kali(a,b int) int {
	return a * b
}
