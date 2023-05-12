## Notes

### Menambah Dependency

- Untuk menambah dependency dapat dilakukan dengan cara :
```bash 
go get nama_module
```

### Upgrade Versi Dependency

- Untuk meng-upgrade dependency dapat dilakukan dengan cara mengubah version/tag pada `go.mod` 
```bash 
require github.com/yogameleniawan/go-module-hello v1.0.0 // sebelum

require github.com/yogameleniawan/go-module-hello v1.5.0 // sesudah
```
```bash 
go get
```


### Major Upgrade Dependency

- Untuk melakukan major upgrade dependency dapat dilakukan dengan cara menjalankan pada terminal dengan catatan module yang diupgrade tidak sama seperti module yang diinstall sebelumnya
```bash 
require github.com/yogameleniawan/go-module-hello // module sebelumnya
```
maka `github.com/yogameleniawan/go-module-hello` sudah tidak digunakan dan harus dihapus terlebih dahulu apabila ingin melakukan major upgrade. Kemudian bisa mendownload ulang module yang akan digunakan.
```bash 
go get github.com/yogameleniawan/go-module-hello/v2
```
sehingga `go.mod` akan ada tambahan module 
```bash 
require github.com/yogameleniawan/go-module-hello/v2
```
