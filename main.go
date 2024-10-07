package main
import (
"bufio"
"fmt"
"os"
"strings"
)
// Struct untuk menyimpan informasi pengguna
type User struct {
name string
age int
email string
}
// Tipe data khusus untuk operasi matematika
type Operation func(float64, float64) float64
// Fungsi untuk menampilkan menu utama
func showMenu() {
fmt.Println("\n===== Menu Utama =====")
fmt.Println("1. Tampilkan 'Hello, World!'")
fmt.Println("2. Operasi Matematika Sederhana")
fmt.Println("3. Simpan dan Tampilkan Data Pengguna")
fmt.Println("4. Keluar")
fmt.Print("Pilih opsi: ")
}
// Fungsi untuk menampilkan "Hello, World!"
func helloWorld() {
fmt.Println("\nHello, World!")
}
// Fungsi untuk operasi matematika sederhana
func kalkulator() {
var a, b float64
var operator string
fmt.Print("\nMasukkan angka pertama: ")
fmt.Scanln(&a)
fmt.Print("Masukkan operator (+, -, *, /): ")
fmt.Scanln(&operator)
fmt.Print("Masukkan angka kedua: ")
fmt.Scanln(&b)
// Deklarasi operasi menggunakan map
operations := map[string]Operation{
"+": func(x, y float64) float64 { return x + y },
"-": func(x, y float64) float64 { return x - y },
"*": func(x, y float64) float64 { return x * y },
"/": func(x, y float64) float64 {
if y != 0 {
return x / y
}
fmt.Println("Error: Pembagian dengan nol!")
return 0
},
}
// Melakukan operasi berdasarkan input operator
if op, found := operations[operator]; found {
result := op(a, b)
fmt.Printf("Hasil: %.2f\n", result)
} else {
fmt.Println("Operator tidak valid!")
}
}
// Fungsi untuk menambahkan data pengguna ke map
func addUser(users map[string]User) {
reader := bufio.NewReader(os.Stdin)
fmt.Print("\nMasukkan nama: ")
name, _ := reader.ReadString('\n')
name = strings.TrimSpace(name)
fmt.Print("Masukkan umur: ")
var age int
fmt.Scanln(&age)
fmt.Print("Masukkan email: ")
email, _ := reader.ReadString('\n')
email = strings.TrimSpace(email)
// Tambahkan user baru ke map
users[name] = User{name: name, age: age, email: email}
fmt.Println("Data pengguna berhasil ditambahkan!")
}
// Fungsi untuk menampilkan semua data pengguna
func showAllUsers(users map[string]User) {
if len(users) == 0 {
fmt.Println("\nBelum ada data pengguna.")
} else {
fmt.Println("\nData Pengguna yang Disimpan:")
for name, user := range users {
fmt.Printf("Nama: %s, Umur: %d, Email: %s\n", name, user.age, user.email)
}
}
}
// Fungsi anonim untuk memfilter pengguna berdasarkan umur
func filterUsers(users map[string]User, minAge int) {
fmt.Printf("\nPengguna dengan umur di atas %d:\n", minAge)
filter := func(user User) bool {
return user.age > minAge
}
for _, user := range users {
if filter(user) {
fmt.Printf("Nama: %s, Umur: %d, Email: %s\n", user.name, user.age, user.email)
}
}
}
func main() {
users := make(map[string]User) // Map untuk menyimpan data pengguna
for {
showMenu()
var pilihan int
fmt.Scanln(&pilihan)
switch pilihan {
case 1:
helloWorld() // Tampilkan "Hello, World!"
case 2:
kalkulator() // Kalkulator sederhana
case 3:
// Sub-menu untuk menambahkan atau menampilkan data pengguna
for {
fmt.Println("\n===== Menu Data Pengguna =====")
fmt.Println("1. Tambahkan Data Pengguna")
fmt.Println("2. Tampilkan Semua Data Pengguna")
fmt.Println("3. Filter Pengguna Berdasarkan Umur")
fmt.Println("4. Kembali ke Menu Utama")
fmt.Print("Pilih opsi: ")
var subPilihan int
fmt.Scanln(&subPilihan)
switch subPilihan {
case 1:
addUser(users) // Menambahkan pengguna baru
case 2:
showAllUsers(users) // Menampilkan semua pengguna
case 3:
fmt.Print("Masukkan umur minimum untuk filter: ")
var minAge int
fmt.Scanln(&minAge)
filterUsers(users, minAge) // Filter pengguna berdasarkan umur
case 4:
break // Kembali ke menu utama
default:
fmt.Println("Opsi tidak valid, coba lagi.")
}
if subPilihan == 4 {
break
}
}
case 4:
fmt.Println("Keluar dari aplikasi.")
return
default:
fmt.Println("Opsi tidak valid, coba lagi.")
}
}
}
