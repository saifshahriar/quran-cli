package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

func surahs() []string {
	surahs := []string{
		"Al-Fatiha",
		"Al-Baqarah",
		"Al-Imran",
		"An-Nisa",
		"Al-Maidah",
		"Al-Anam",
		"Al-Araf",
		"Al-Anfal",
		"At-Tawbah",
		"Yunus",
		"Hud",
		"Yusuf",
		"Ar-Rad",
		"Ibrahim",
		"Al-Hijr",
		"An-Nahl",
		"Al-Isra",
		"Al-Kahf",
		"Maryam",
		"Ta-Ha",
		"Al-Anbiya",
		"Al-Hajj",
		"Al-Muminun",
		"An-Nur",
		"Al-Furqan",
		"Ash-Shuara",
		"An-Naml",
		"Al-Qasas",
		"Al-Ankabut",
		"Ar-Rum",
		"Luqman",
		"As-Sajda",
		"Al-Ahzab",
		"Saba",
		"Fatir",
		"Ya-Sin",
		"As-Saffat",
		"Sad",
		"Az-Zumar",
		"Ghafir",
		"Fussilat",
		"Ash-Shura",
		"Az-Zukhruf",
		"Ad-Dukhan",
		"Al-Jathiya",
		"Al-Ahqaf",
		"Muhammad",
		"Al-Fath",
		"Al-Hujraat",
		"Qaf",
		"Adh-Dhariyat",
		"At-Tur",
		"An-Najm",
		"Al-Qamar",
		"Ar-Rahman",
		"Al-Waqia",
		"Al-Hadid",
		"Al-Mujadila",
		"Al-Hashr",
		"Al-Mumtahina",
		"As-Saff",
		"Al-Jumua",
		"Al-Munafiqoon",
		"At-Taghabun",
		"At-Talaq",
		"At-Tahrim",
		"Al-Mulk",
		"Al-Qalam",
		"Al-Haaqqa",
		"Al-Maarij",
		"Nuh",
		"Al-Jinn",
		"Al-Muzzammil",
		"Al-Muddathir",
		"Al-Qiyama",
		"Al-Insan",
		"Al-Mursalat",
		"An-Naba",
		"An-Nazi'at",
		"Abasa",
		"At-Takwir",
		"Al-Infitar",
		"Al-Mutaffifin",
		"Al-Inshiqaq",
		"Al-Buruj",
		"At-Tariq",
		"Al-Ala",
		"Al-Ghashiya",
		"Al-Fajr",
		"Al-Balad",
		"Ash-Shams",
		"Al-Lail",
		"Ad-Duhaa",
		"Ash-Sharh",
		"At-Tin",
		"Al-Alaq",
		"Al-Qadr",
		"Al-Bayyina",
		"Az-Zalzala",
		"Al-Adiyat",
		"Al-Qaria",
		"At-Takathur",
		"Al-Asr",
		"Al-Humaza",
		"Al-Fil",
		"Quraish",
		"Al-Maun",
		"Al-Kawthar",
		"Al-Kafiroon",
		"An-Nasr",
		"Al-Masad",
		"Al-Ikhlas",
		"Al-Falaq",
		"An-Nas",
	}
	return surahs
}

func checkdbHealth() string {
	textPath := "/usr/local/share/quran-cli/db/en.sahih.txt"
	f, err := os.Open(textPath)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	return string(content) // Return file content
}

func text() [][]string {
	text := checkdbHealth()
	readtext := csv.NewReader(strings.NewReader(text))
	readtext.Comma = '|'
	readtext.LazyQuotes = true
	lines, err := readtext.ReadAll()
	if err != nil {
		log.Fatalf("failed to read lines: %s", err)
	}
	return lines
}
