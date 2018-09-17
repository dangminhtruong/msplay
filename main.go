package main

import (
	"io/ioutil"
	"flag"
	"github.com/hajimehoshi/oto"
	"github.com/tosone/minimp3"
	"os"
	"github.com/dimiro1/banner"
	"github.com/mattn/go-colorable"
	"fmt"
	"github.com/janeczku/go-spinner"
	"os/user"
	"path/filepath"
)

func main() {
	var song string
	var shuffle string
	user, _ := user.Current()
	defaultPath := user.HomeDir + "/Music/"

	flag.StringVar(&song, "song", "", "Song name")
	flag.StringVar(&shuffle, "shuffle", "off", "Turn on shuffle mode")
	flag.Parse()
	songPath :=  defaultPath + song + ".mp3"
	//-----------------------------------------------
	displayBanner()
	if song == "" && shuffle != "off"{
		shuffleMode(defaultPath)
	} else {
		playSong(song, songPath)
	}
}

func playSong(song, songPath string) {
	if _, err := os.Stat(songPath); err == nil {
		s := spinner.StartNew(song + " now playing... ♪ ♫ ♬ ")
		var file, _ = ioutil.ReadFile(songPath)
		dec, data, _ := minimp3.DecodeFull(file)
		player, _ := oto.NewPlayer(dec.SampleRate, dec.Channels, 2, 2048)
		player.Write(data)
		s.Stop()
	} else {
		fmt.Println(song + "\033[0;0H song could not be found !")
	}
}

func displayBanner() {
	if _, fileErr := os.Stat("/tmp/banner.txt"); os.IsNotExist(fileErr) {
		bannerTxtBytes := MustAsset("banner.txt")
		errRewrite := ioutil.WriteFile("/tmp/banner.txt", bannerTxtBytes, 0777)
		if errRewrite != nil {
			fmt.Println(errRewrite)
		}
	}
	in, _ := os.Open("/tmp/banner.txt")
	fmt.Println()
	banner.Init(colorable.NewColorableStdout(), true, true, in)
	fmt.Println()
}

func shuffleMode(defaultPath string) {
	files, err := filepath.Glob(defaultPath + "/" + "*.mp3")
	if err != nil {
		fmt.Println(err)
    }
	if len(files) >= 1 {
		for _,songPath := range(files) {
			_, songName := filepath.Split(songPath)
			playSong(songName, songPath)
		}
	}
}