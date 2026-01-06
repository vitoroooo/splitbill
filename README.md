# 💸 Split Bill

**Split Bill** adalah aplikasi berbasis **Command Line Interface (CLI)** yang digunakan untuk membantu pembagian biaya bersama secara adil dan transparan. Aplikasi ini cocok digunakan dalam aktivitas kelompok seperti makan bersama, perjalanan, atau kegiatan komunitas lainnya.

Proyek ini dikembangkan sebagai bagian dari pembelajaran **Berpikir Komputasional dan Pengenalan Pemrograman** menggunakan bahasa **Go (Golang)**.

---

## ✨ Fitur Utama

- Input data anggota secara fleksibel  
- Pencatatan transaksi pembayaran  
- Perhitungan total pengeluaran otomatis  
- Pembagian biaya secara merata ke setiap anggota  
- Penentuan status anggota (bayar, terima, atau lunas)  
- Instruksi pembayaran antar anggota hingga semua lunas  

---

## 🛠 Teknologi yang Digunakan

- **Bahasa Pemrograman**: Go (Golang)  
- **Interface**: Command Line Interface (CLI)  
- **Library**: Standard library `fmt` (tanpa library eksternal)

---

## 📂 Struktur Program

- `main()` – Fungsi utama program  
- `inputAnggota()` – Input data anggota  
- `inputPengeluaran()` – Input transaksi pembayaran  
- `hitungPembagian()` – Perhitungan biaya per orang  
- `tampilkanHasil()` – Menampilkan ringkasan hasil  
- `hitungTransfer()` – Menentukan siapa membayar ke siapa  

---

## ▶️ Cara Menjalankan Program

Pastikan Go sudah terinstal di komputer Anda.

```bash
go run splitbill.go
