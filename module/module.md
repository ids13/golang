module adalah sekumpulan package
package adalah sekumpulan source file
workspace adalah fitur di golang 1.18 ke atas untuk bekerja dengan beberapa modul

mymodule
|_ go.mod
|_ paket1
|   |_ func1.go
|   |_ func2.go
|_ paket2
    |_ func1.go
    |_ func2.go

membuat modul
`go mod init mymodule`
`go mod init example.com/mymodule`
(example.com hanya sebagai contoh)

nama package di dalam file func sesuaikan dengan nama foldernya,contohnya di func1.go di paket1
`package paket1`

import ketika mau di gunakan 
`import "example.com/mymodule/paket1"`

menambah depedensi
`go get .`
`go get example.com/theirmodule`
`go get example.com/theirmodule@v1.3.4`
`go get example.com/theirmodule@latest`
hapus
`go get example.com/theirmodule@none`
update
`go get -u`
`go get -u example.com/moduleA`

`go get` updates the requirements listed in your go.mod
`go mod download` does not add new requirements or update existing requirements. (At most, it will ensure that the existing requirements.

`go mod vendor` / `go work vendor` : download depedensi dan membuat folder vendor ke project kita

`go install example.com/cmd`
`go install example.com/cmd@v1.2.3`
`go install example.com/cmd@latest`
(harus ada paket main nya)

melihat daftar modul
`go list -m all`
`go list -m -u all`
`go list -m example.com/theirmodule`
`go mod graph`
`go mod why all`

melakukan syncronisasi file go.mod (jika ada perubahan di file penambahan/pengurangan modul di go.mod, maka depedensi akan di perbaharui)
`go mod tidy`

melihat dokkumentasi komen(komen yang letaknya di paling atas,lebih atas dari `package namaPaket`)
`go doc namaModul`
`go doc namaModul/namaPaket`

`go clean -i -n`
Opsi -i menandakan bahwa perintah hanya akan membersihkan objek tempat, dan -n digunakan untuk menampilkan perintah yang akan dijalankan tanpa benar-benar menjalankannya.
`go clean -modcache -cache`
Perintah ini akan membersihkan cache modul (-modcache) dan cache lainnya (-cache).

mendeklarasikan versi minimum yang diperlukan dari ketergantungan modul tertentu
require example.com/othermodule v1.2.3
-require=path@version
-droprequire=path

menggantikan konten versi modul tertentu, atau semua versi modul, dengan konten yang ditemukan di tempat lain
replace example.com/othermodule => example.com/othermodule v1.2.3
replace example.com/othermodule v1.2.5 => ./othermodule
-replace=old[@v]=new[@v]
-dropreplace=old[@v]

mencegah versi modul dimuat oleh go
exclude example.com/theirmodule v1.3.0
-exclude=path@version
-dropexclude=path@version

menunjukkan bahwa versi atau rentang versi modul yang ditentukan go.mod tidak boleh diandalkan
retract v0.3.0
-retract=version
-dropretract=version

retract: Menarik kembali modul dari go.mod, tetapi modul masih dapat diimpor oleh proyek lain.
(jika kita sebagai pengembang ingin ada versi yang di kecualikan,supaya tidak di gunakan user lain)
exclude: Mengabaikan versi tertentu dari suatu modul dalam penyelesaian dependensi proyek Anda.

versi

In development
v0.x.x modul masih dalam pengembangan dan tidak stabil
Major version
v1.x.x perubahan API publik yang tidak kompatibel dengan versi sebelumnya (contohnya ada penambahan parameter di penggunaan fungsinya)
Minor version
vx.4.x perubahan API publik yang kompatibel dengan versi sebelumnya (hanya perubahan kecil)
Patch version
vx.x.1 perubahan yang tidak memengaruhi API publik modul atau dependensinya.
Pre-release version
vx.x.x-beta.2 Memberi sinyal bahwa ini adalah pencapaian pra-rilis, seperti alfa atau beta . Rilis ini tidak memberikan jaminan stabilitas

ketika go get sertakan versinya contoh github.com/user/module@v1.0.0
biar langsung terdownload, hati hati dengan menentukan versi karena module akan tercache ke goproxy
go get akan menangani versi 0.x.x sampai 1.x.x secara otomatis,tapi jika sudah v2.x.x ke atas harus menggunakan penangan majorversion.
tehnik yang bisa di gunakan 
contoh : github.com/user/module

-membuat branch baru
buat branch baru `git checkout -b v2`
ganti bagian moodule di go.mod jadi `github.com/user/module/v2`
push dengan `git push origin v2`
tambah tag `git tag v2.0.0` saat masih di branch v2 kemudian push
kemudian di download dengan `go get github.com/user/module/v2@v2.0.0`

-membuat folder baru
tambahkan folder v2
salin semua file di modul utama (file dan dir) ke folder v2
ganti bagian moodule di go.mod jadi `github.com/user/module/v2`
push dengan `git push origin v2`
tambah tag `git tag v2.0.0` saat masih di branch v2 kemudian push
kemudian di download dengan `go get github.com/user/module/v2@v2.0.0`

-langsung merubah go.mod (menambahkan v2)
ganti bagian moodule di go.mod jadi `github.com/user/module/v2`
push dengan `git push origin v2`
tambah tag `git tag v2.0.0` saat masih di branch v2 kemudian push
kemudian di download dengan `go get github.com/user/module/v2@v2.0.0`

jika menggunakan no v2 atau versi tinggi, harus menyertakan nomor versi nya ke jalur import
ketika `go get github.com/user/module/v2` maka harus jadi `go get github.com/user/module/v2@v2.0.0`
untuk lihat versi gunakan `go list -m -u all`.untuk v0 dan v1 tidak perlu di sertakan
jika tidak di sertakan nomor versinya,maka v1/v0 akan ikut terdownload,dalam hal ini `github.com/user/module` akan ikut terdownload

inisialisasi workspace
`go work init `
`go work init ./modul1 ./modul2`
menambahkan modul ke workspace
`go work use ./modul`
merubah go.work
`go work edit `
singkronisasi depedensi dengan go.work
`go work sync`