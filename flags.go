package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jessevdk/go-flags"
)

var VERSION string = "1.0"

type Options struct {
	Version  bool   `short:"v" long:"version"   description:"Print version and exit"`
	FullText bool   `short:"f" long:"full-text" description:"Print the full text of the Qur'an to stdout and exit"`
	Surah    uint8  `short:"s" long:"surah"     description:"Surah number"`
	Ayah     uint16 `short:"a" long:"ayah"      description:"Ayah number"`
}

func printVersion() {
	fmt.Printf("quran-cli %s\n", VERSION)
	os.Exit(0)
}

func invalidSurahAyah() {
	fmt.Println(
		"The Surah number should be between 1 and 114 and the Ayah number should be between 1 and 6236.",
	)
	os.Exit(1)
}

func printFullText() {
	lines := text()
	for _, line := range lines {
		fmt.Println(line[0] + ":" + line[1] + "\t" + line[2])
	}
}

func printSpecificAyah(opts Options) {
	lines := text()
	for _, line := range lines {
		surahnum, _ := strconv.Atoi(line[0])
		ayahnum, _ := strconv.Atoi(line[1])
		if uint8(surahnum) == opts.Surah && uint16(ayahnum) == opts.Ayah {
			fmt.Println(line[0] + ":" + line[1] + "\t" + line[2])
		}
	}
}

func printSpecificAyahFromEachSurah(opts Options) {
	lines := text()
	for _, line := range lines {
		ayahnum, _ := strconv.Atoi(line[1])
		if uint16(ayahnum) == opts.Ayah {
			fmt.Println(line[0] + ":" + line[1] + "\t" + line[2])
		}
	}
}

func printSurah(opts Options) {
	lines := text()
	for _, line := range lines {
		surahnum, _ := strconv.Atoi(line[0])
		if uint8(surahnum) == opts.Surah {
			fmt.Println(line[0] + ":" + line[1] + "\t" + line[2])
		}
	}
}

func parseFlags() {
	var opts Options

	parser := flags.NewParser(&opts, flags.HelpFlag)
	parser.Usage = "[FLAGS] [OPTIONS]"
	parser.ShortDescription = "A tiny Qur'an utility for cli written in Go under 100 LOC"

	_, err := flags.Parse(&opts)
	if err != nil {
		return
	}

	// If no arguments are passed, print the help message
	if len(os.Args) == 1 {
		parser.WriteHelp(os.Stdout)
	}

	/* Flags */
	if opts.Version {
		printVersion()
	}

	if opts.Surah < 0 || opts.Surah > 114 || opts.Ayah < 0 || opts.Ayah > 6236 {
		invalidSurahAyah()
	}

	if opts.FullText {
		printFullText()
	} else if opts.Surah > 0 && opts.Ayah > 0 {
		printSpecificAyah(opts)
	} else if opts.Ayah > 0 {
		printSpecificAyahFromEachSurah(opts)
	} else if opts.Surah > 0 {
		printSurah(opts)
	}
}
