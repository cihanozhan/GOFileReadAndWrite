package main 

import (
    "fmt"
    "io"
    //"io/ioutil"
    "os"
    "bufio"
    "math/rand"
    "time"
    "strconv"
)

func main(){

    // Biz klavyeden girilen karakterleri content olarak alacağız. 
    // Bu nedenle bu kısmı açıklama satırı olarak bıraktım. 
    //content := "Merhaba GO!"

    // Klavyeden girilen veriyi elde edebilmek için projeye eklediğimiz bufio kütüphanesini kullanıyoruz. 
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("************\n")
    fmt.Print("Metin: ")
    content, _ := reader.ReadString('\n')
    fmt.Print("************\n")

    // Dosya adında kullanmak üzere rastgele değerler üretiyoruz.
    r := random(99, 99999)
    // Üretilen rastgele değeri strconv olarak eklediğimiz kütüphanenin Itoa() metodu ile string'e dönüştürüyoruz.
    randNumber := strconv.Itoa(r)

    // Şu anın tarihini formatlayarak ekrana yaz.
    t := time.Now()
    fmt.Printf("İşlem zamanı : %d-%02d-%02d / %02d:%02d:%02d\n",
        t.Year(), t.Month(), t.Day(),
        t.Hour(), t.Minute(), t.Second())

    // Dosyaya verilecek isim için adını ve uzantısını belirle.
    ext := ".txt"
    fileName := "./dummy-" + randNumber + ext

    // Dosyayı oluştur. 
    file, err := os.Create(fileName)
    // Bir hata oluşursa err dolu gelecek, onu checkError() metoduna gönder.
    checkError(err)
    // Dosyayı serbest bırak.
    defer file.Close()

    // Oluşturulan dosyaya kullanıcıdan alınan metni yaz.
    ln, err := io.WriteString(file, content)
    checkError(err)

    fmt.Printf("Dosyada %v karakter var.", ln)

    //bytes := []byte(content)
    //ioutil.WriteFile(fileName, bytes, 0644)

    // Bir alt satıra in(newLine)
    fmt.Print("\n")
}

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func checkError(err error){
    if err != nil {
        panic(err)
    }
}