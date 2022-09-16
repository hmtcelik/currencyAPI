package models

import "encoding/xml"

type TcmbVeri struct {
	XMLName     xml.Name `xml:"tcmbVeri"`
	Text        string   `xml:",chardata"`
	BaslikBilgi struct {
		Text         string `xml:",chardata"`
		Kod          string `xml:"kod"`
		VeriTipi     string `xml:"veri_tipi"`
		VeriTanim    string `xml:"veri_tanim"`
		Yayimlayan   string `xml:"yayimlayan"`
		Tel          string `xml:"tel"`
		Faks         string `xml:"faks"`
		Eposta       string `xml:"eposta"`
		ZamanEtiketi string `xml:"zaman_etiketi"`
	} `xml:"baslik_bilgi"`
	DovizKurListe struct {
		Text             string `xml:",chardata"`
		GecerlilikTarihi string `xml:"gecerlilik_tarihi,attr"`
		Saat             string `xml:"saat,attr"`
		Kur              []struct {
			Text             string `xml:",chardata"`
			DovizCinsiTabani string `xml:"doviz_cinsi_tabani"`
			DovizCinsi       string `xml:"doviz_cinsi"`
			Birim            string `xml:"birim"`
			Alis             string `xml:"alis"`
			SiraNo           string `xml:"sira_no"`
		} `xml:"kur"`
	} `xml:"doviz_kur_liste"`
	Aciklama string `xml:"aciklama"`
}
