/* Tentu, berikut adalah studi kasus lain menggunakan Golang.
Kali ini, kita akan membuat program sederhana untuk mengelola daftar tugas (to-do list): */

package main

import (
	"fmt"
)

// Task adalah struktur data untuk merepresentasikan tugas
type Task struct {
	ID     int
	Title  string
	Status bool
}

// TaskList adalah struktur data untuk menyimpan daftar tugas
type TaskList struct {
	Tasks []Task
}

// AddTask menambahkan tugas baru ke daftar
func (tl *TaskList) AddTask(title string) {
	newTask := Task{ID: len(tl.Tasks) + 1, Title: title, Status: false}
	tl.Tasks = append(tl.Tasks, newTask)
}

// ShowTasks menampilkan semua tugas dalam daftar
func (tl *TaskList) ShowTasks() {
	fmt.Println("Daftar Tugas:")
	for _, task := range tl.Tasks {
		status := "Belum Selesai"
		if task.Status {
			status = "Sudah Selesai"
		}
		fmt.Printf("%d. %s - %s\n", task.ID, task.Title, status)
	}
}

// CompleteTask menandai tugas sebagai selesai berdasarkan ID
func (tl *TaskList) CompleteTask(id int) {
	for i := range tl.Tasks {
		if tl.Tasks[i].ID == id {
			tl.Tasks[i].Status = true
			fmt.Printf("Tugas '%s' selesai.\n", tl.Tasks[i].Title)
			return
		}
	}
	fmt.Println("Tugas tidak ditemukan.")
}

func main() {
	// Inisialisasi daftar tugas
	tasks := TaskList{}

	// Menambahkan beberapa tugas
	tasks.AddTask("Mengerjakan Tugas Kuliah")
	tasks.AddTask("Belajar Bahasa Golang")
	tasks.AddTask("Berolahraga")

	// Menampilkan daftar tugas
	tasks.ShowTasks()

	// Menandai tugas sebagai selesai
	tasks.CompleteTask(2)

	// Menampilkan daftar tugas setelah pembaruan
	tasks.ShowTasks()
}
