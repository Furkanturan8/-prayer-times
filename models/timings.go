package models

// Timings modeli
type Timings struct {
	ID              int    `json:"id" gorm:"unique"`
	CityID          int    `json:"city_id" gorm:"unique"`
	GregorianDateID int    `json:"gregorian_date_id"` // miladi takvim
	HijriDateID     int    `json:"hijri_date_id"`     // hicri takvim
	Fajr            string `json:"Fajr"`              // imsak
	Dhuhr           string `json:"Dhuhr"`             // gün doğumu
	Asr             string `json:"Asr"`               // öğle
	Maghrib         string `json:"Maghrib"`           // ikindi
	Isha            string `json:"Isha"`              // aksam
	Imsak           string `json:"Imsak"`             // yatsı
}

// GregorianDate modeli
type GregorianDate struct {
	Date  string `json:"date"`
	Day   string `json:"day"`
	Month int    `json:"month"`
	Year  string `json:"year"`
}

// HijriDate modeli
type HijriDate struct {
	Date  string `json:"date"`
	Day   string `json:"day"`
	Month int    `json:"month"`
	Year  string `json:"year"`
}

// PrayerTimes modeli, hem namaz saatlerini hem de tarihleri içerir --> Yani yukarıdakilerin hepsini kapsar
type PrayerTimes struct {
	ID            int           `json:"id" gorm:"unique"`
	City          string        `json:"city" gorm:"unique"`
	Timings       Timings       `json:"timings"`
	GregorianDate GregorianDate `json:"gregorian_date"`
	HijriDate     HijriDate     `json:"hijri_date"`
}
