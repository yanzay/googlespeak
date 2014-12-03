package googlespeak

import (
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/user"
	"unicode/utf8"
)

var langs = []string{"af", "ar", "ca", "cs", "cy", "da", "de", "el", "en",
	"es", "fi", "fr", "hi", "hr", "ht", "hu", "hy", "id", "is", "it",
	"ja", "ko", "la", "lv", "mk", "nl", "no", "pl", "pt", "ro", "ru",
	"sk", "sq", "sr", "sv", "sw", "ta", "tr", "vi", "zh"}

func Say(s string, args ...string) error {
	lang := "en"
	if len(args) > 0 {
		lang = args[0]
	}

	err := validateParams(s, lang)
	if err != nil {
		return err
	}

	log.Printf("Lang: %s, say: %s", lang, s)
	err = speak(s, lang)
	return err
}

func validateParams(s, lang string) error {
	if utf8.RuneCountInString(s) > 100 {
		return errors.New("Text exceeds max char limit (100)!")
	}

	if !isValidLang(lang) {
		return errors.New("Invalid language code!")
	}
	return nil
}

func isValidLang(s string) bool {
	for _, l := range langs {
		if l == s {
			return true
		}
	}
	log.Printf("Invalid language: %s", s)
	return false
}

func getAudio(s, lang string) (io.ReadCloser, error) {
	resp, err := http.Get("http://translate.google.com/translate_tts" +
		"?ie=UTF-8&tl=" + lang + "&q=" + url.QueryEscape(s))
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func play(audio io.Reader) error {
	mplayer := exec.Command("mplayer", "-cache", "8092", "-")
	mplayer.Stdin = audio
	return mplayer.Run()
}

func getFromCache(s, lang string) (io.ReadCloser, error) {
	cached, err := os.Open(getCacheDir() + "/" + lang + "/" + s + ".mp3")
	return cached, err
}

func cacheAudio(stream io.Reader, s, lang string) (io.ReadCloser, error) {
	langCacheDir := getCacheDir() + "/" + lang
	dir, err := os.Open(langCacheDir)
	if os.IsNotExist(err) {
		_ = os.MkdirAll(langCacheDir, 0700)
	}
	defer dir.Close()

	filename := s + ".mp3"

	f, err := os.Open(langCacheDir + "/" + filename)
	if os.IsNotExist(err) {
		f, err = os.Create(langCacheDir + "/" + filename)
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(f, stream)
		return f, err
	}
	return f, err
}

func speak(s, lang string) error {
	log.Printf("Get from cache %s/%s", lang, s)
	audio, err := getFromCache(s, lang)
	if err != nil {
		log.Printf("Cache for %s/%s not found. Trying to get audio from Google", lang, s)
		stream, err := getAudio(s, lang)
		if err == nil {
			log.Printf("Caching stream for %s/%s", lang, s)
			audio, _ = cacheAudio(stream, s, lang)
		}
	}
	defer audio.Close()
	err = play(audio)
	if err != nil {
		return err
	}
	return nil
}

func getCacheDir() string {
	xdgCacheHome := os.Getenv("XDG_CACHE_HOME")
	if xdgCacheHome == "" {
		user, _ := user.Current()
		home := user.HomeDir
		xdgCacheHome = home + "/.cache/gspeak"
	}

	dir, err := os.Open(xdgCacheHome)
	if os.IsNotExist(err) {
		_ = os.MkdirAll(xdgCacheHome, 0700)
		dir, err = os.Open(xdgCacheHome)
	}
	return dir.Name()
}
