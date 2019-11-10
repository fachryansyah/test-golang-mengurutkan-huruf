# Solusi 2 : Mengurutkan abjad dari suatu kalimat
Code jawaban untuk "Buat fungsi sederhana untuk mengurutkan abjad dari suatu kalimat, dengan memisahkan huruf hidup dan huruf mati"

## Instalasi
### Menginstal Module
```sh
$ go get github.com/fachryansyah/test-golang-mengurutkan-huruf
```
### Menggunakan Module
```go
package main

import "testing"
import "github.com/fachryansyah/test-golang-mengurutkan-huruf"

func TestModule(test *testing.T) {
	result := SortWord("osama")
	if result != "aaoms" {
		test.Errorf("SortWord function was error, want: %s got: %s", "aaoms", result)
	}
}

```
