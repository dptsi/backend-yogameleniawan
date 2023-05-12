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
