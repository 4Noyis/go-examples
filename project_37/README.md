# Stateful Goroutines

## Özet

Program eşzamanlı olarak 100 okuma ve 10 yazma işlemi gerçekleştirir.
Okuma ve yazma işlemleri için atomik sayaçlar kullanılarak her işlemin kaç kez yapıldığı güvenli bir şekilde izlenir.
Her okuma işlemi, eşzamanlı olarak çalışan bir goroutine ile rastgele bir anahtar üzerinden yapılır. Aynı şekilde yazma işlemleri de rastgele anahtar ve değerler üzerinden yapılır.
Durum yönetimi (state haritası) ve işlemler goroutine'ler ve kanallar aracılığıyla güvenli bir şekilde kontrol edilir.

### Genel Yapı:

readOp ve writeOp structları:
readOp: Okuma işlemlerini temsil eder. İçinde bir anahtar (key) ve okuma sonucunu döndürmek için kullanılan bir kanal (resp) bulunur.
writeOp: Yazma işlemlerini temsil eder. İçinde bir anahtar (key), yazılacak değer (val) ve yazma işleminin sonucunu bildirmek için bir kanal (resp) bulunur.

### Okuma ve Yazma Sayaçlarının Tanımlanması:

```Go
var readOps uint64
var writeOps uint64
```

readOps: Toplam okuma işlemlerini saymak için kullanılan bir atomik sayaç.
writeOps: Toplam yazma işlemlerini saymak için kullanılan bir atomik sayaç.

### Channelların Tanımlanması:

```Go
reads := make(chan readOp)
writes := make(chan writeOp)
```

reads kanalı, readOp (okuma işlemi) türündeki istekleri yönetir.
writes kanalı, writeOp (yazma işlemi) türündeki istekleri yönetir.

### Durum Yönetimi Goroutine'i

```Go
go func() {
    var state = make(map[int]int)
    for {
        select {
        case read := <-reads:
            read.resp <- state[read.key]
        case write := <-writes:
            state[write.key] = write.val
            write.resp <- true
        }
    }
}()
```

Bu goroutine bir durum haritası (state) yönetir.
Sürekli çalışan bir select bloğu, gelen okuma ve yazma isteklerini bekler.
reads kanalından gelen bir okuma isteğinde, state haritasındaki ilgili key değeri bulunur ve resp kanalına geri döner.
writes kanalından gelen bir yazma isteğinde, state haritasında belirtilen key'e yeni val atanır ve yazma işleminin başarılı olduğu resp kanalına gönderilir.

### Okuma Goroutine'i

```Go
for r := 0; r < 100; r++ {
    go func() {
        for {
            read := readOp{
                key:  rand.Intn(5),
                resp: make(chan int)}
            reads <- read
            <-read.resp
            atomic.AddUint64(&readOps, 1)
            time.Sleep(time.Millisecond)
        }
    }()
}

```

100 tane okuma goroutine'i başlatılır. Her goroutine şu adımları tekrarlar:

- 0-4 arasında rastgele bir anahtar (key) oluşturur ve bir readOp talebi hazırlar.
- Bu talep reads kanalına gönderilir.
- resp kanalından okuma sonucunu bekler ve alır.
- readOps atomik sayacı 1 artırılır.
- Kısa bir süre bekler (time.Sleep(time.Millisecond)) ve bu döngüyü tekrar eder.

### Yazma Goroutine'leri

```Go
for w := 0; w < 10; w++ {
    go func() {
        for {
            write := writeOp{
                key:  rand.Intn(5),
                val:  rand.Intn(100),
                resp: make(chan bool)}
            writes <- write
            <-write.resp
            atomic.AddUint64(&writeOps, 1)
            time.Sleep(time.Millisecond)
        }
    }()
}
```

10 tane yazma goroutine'i başlatılır. Her goroutine şu adımları tekrarlar:

- 0-4 arasında rastgele bir anahtar (key) ve 0-99 arasında rastgele bir değer (val) oluşturur.
- Bir writeOp talebi oluşturur ve bu talebi writes kanalına gönderir.
- resp kanalından yazma işleminin başarılı olduğunu öğrenir.
- writeOps atomik sayacı 1 artırılır.
- Kısa bir süre bekler ve döngüyü tekrar eder.

### Bir Saniye Bekleme

Program 1 saniye boyunca çalışır ve bu süre zarfında okuma ve yazma işlemleri sürekli devam eder.

```Go
time.Sleep(time.Second)
```

### Sonuçların Yazıdırılması

```Go
readOpsFinal := atomic.LoadUint64(&readOps)
fmt.Println("readOps:", readOpsFinal)
writeOpsFinal := atomic.LoadUint64(&writeOps)
fmt.Println("writeOps:", writeOpsFinal)
```

1 saniyenin sonunda, readOps ve writeOps sayaçlarının son değerleri atomik olarak yüklenir ve ekrana yazdırılır.
