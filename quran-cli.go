package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jessevdk/go-flags"
)

var (
	VERSION string
)

type Options struct {
	Version  bool   `short:"v" long:"version"   description:"Print version and exit"`
	FullText bool   `short:"f" long:"full-text" description:"Print the full text of the Qur'an to stdout and exit"`
	Surah    uint8  `short:"s" long:"surah"     description:"Surah number"`
	Ayah     uint16 `short:"a" long:"ayah"      description:"Ayah number"`
}

func main() {
	var opts Options

	parser := flags.NewParser(&opts, flags.HelpFlag)
	parser.Usage = "[FLAGS] [OPTIONS]"
	parser.ShortDescription = "A tiny Qur'an utility for cli written in Go under 100 LOC"

	_, err := flags.Parse(&opts)
	if err != nil {
		return
	}

	if len(os.Args) == 1 {
		parser.WriteHelp(os.Stdout)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	textPath := homeDir + "/.local/share/quran-cli/db/en.sahih.txt"
	text, err := os.Open(textPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer text.Close()

	readtext := csv.NewReader(text)
	readtext.Comma = '|'
	readtext.LazyQuotes = true
	lines, err := readtext.ReadAll()
	if err != nil {
		log.Fatalf("failed to read lines: %s", err)
	}

	/* Flags */
	if opts.Version {
		fmt.Printf("quran-cli %s\n", VERSION)
		os.Exit(0)
	}

	if opts.Surah < 0 || opts.Surah > 114 || opts.Ayah < 0 || opts.Ayah > 6236 {
		fmt.Println("There is only 114 Surah (chapters) and 6236 Ayah (verses) in the Qur'an.")
		os.Exit(1)
	}

	if opts.FullText {
		for _, line := range lines {
			fmt.Println(line[0] + ":" + line[1] + "\t" + line[2])
		}
	} else if opts.Surah > 0 && opts.Ayah > 0 {
		for _, line := range lines {
			surahnum, _ := strconv.Atoi(line[0])
			ayahnum, _ := strconv.Atoi(line[1])
			if uint8(surahnum) == opts.Surah && uint16(ayahnum) == opts.Ayah {
				fmt.Println(line[0] + ":" + line[1] + "\t" + line[2])
			}
		}
	} else if opts.Surah > 0 {
		for _, line := range lines {
			surahnum, _ := strconv.Atoi(line[0])
			if uint8(surahnum) == opts.Surah {
				fmt.Println(line[0] + ":" + line[1] + "\t" + line[2])
			}
		}
	}

}
