## Notes

### Menjalankan Unit Test

- Untuk menjalankan unit test secara keseluruhan dalam satu package
```bash 
go test
```
- Untuk menjalankan unit test dengan menampilkan function apa saja yang sudah di jalankan
```bash
go test -v
```
- Untuk menjalankan unit test dengan menjalankan function yang ingin dijalankan
```bash
go test -v -run=TestNamaFunction
```

### Menggagalkan Unit Test

- `t.Fail()` menggagalkan unit test namun akan melanjutkan eksekusi unit test selanjutnya. Namun diakhir ketika selesai, unit test tersebut dianggap gagal.
- `t.FailNow()` menggagalkan unit test dan langsung menghentikan eksekusi unit test tanpa melanjutkan ke unit test berikutnya
- `t.Error()` menggagalkan unit test dengan menuliskan log (print) error kemudian dilanjutkan dengan memanggil function `Fail()`.
- `t.Fatal()` menggagalkan unit test dengan menuliskan log (print) error kemudian memanggil function `FailNow()` sehingga ekseksui unit test berhenti

### Menjalankan Unit Test Benchmark

- Untuk menjalankan unit test Benchmark dengan menampilkan function apa saja yang sudah di jalankan
```bash
go test -v -bench=.
```
- Untuk menjalankan unit test Benchmark tanpa unit test
```bash
go test -v -run=NotMathUnitTest -bench=.
```
- Untuk menjalankan unit test Benchmark tertentu
```bash
go test -v -run=NotMathUnitTest -bench=BenchmarkTest
```
- Untuk menjalankan bencmark di root module dan ingin semua module dijalankan
```bash
go test -v -bench=../..
```