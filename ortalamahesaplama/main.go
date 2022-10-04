package ortalamahesaplama

// Vize Notu 0 ile 100 arasında olabilir
// Final Notu 0 ile 100 arasında olabilir
// İnsiyatif 0 ile 1 arasında olabilir
// Ortalama 0 ile 100 arasında olabilir
// Ortalama 65'ten küçükse ve ortalama*(1+insiyatif) 65'ten büyükse ortalama 65 olmalı
// Ortalama 65'ten büyükse ortalama değişmemelidir
// Çıktı olarak ortalama dönmelidir
// OrtalamaHesapla fonksiyonu `go test ./...` komutu ile test edilmelidir
// Testlerin hepsi PASS olmalıdır

func OrtalamaHesapla(vizeNot, finalNot int, insiyatif float64) float64 {

	ortalamaH := (float64(vizeNot) * 0.4) + (float64(finalNot) * 0.6)

	insiyatifH := ortalamaH * (1 + insiyatif)

	/*if ortalamaH < 65 && ortalamaH * (1+insiyatifH) > 65 {
		ortalamaH = 65


	}else if ortalamaH > 65{
		return ortalamaH

	}
	return ortalamaH*/

	switch ort := ortalamaH; {
	case ort < 65 && insiyatifH > 65:
		ortalamaH = 65
		return ortalamaH
	case ort > 65:
		return ortalamaH

	}

	return ortalamaH

}
