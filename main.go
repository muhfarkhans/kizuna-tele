package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// videoID := "k2MWBy-Hb1M"
	// client := youtube.Client{}

	// video, err := client.GetVideo(videoID)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// audioRegex := regexp.MustCompile(`^(audio/.+)`)

	// for _, format := range video.Formats {
	// 	matched := audioRegex.FindStringSubmatch(format.MimeType)
	// 	if len(matched) > 1 {
	// 		audioFormat := matched[1]
	// 		fmt.Println(audioFormat, format.AudioQuality)
	// 	}

	// 	fmt.Println(format.MimeType)
	// 	fmt.Println(format.QualityLabel, format.URL)
	// }

	// fmt.Println(video.Formats)

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		fmt.Println(err)
	}

	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "help":
			msg.Text = `
			List of commands:
			1. /youtube <url> mp3|video
			`
		case "yt":
			args := strings.Split(update.Message.CommandArguments(), " ")
			re := regexp.MustCompile(`v=([^&]+)`)
			code := re.FindStringSubmatch(args[0])
			msg.Text = code[1]
		default:
			if update.Message.IsCommand() {
				msg.Text = "I don't know that command"
			} else {
				// audioURL := ""
				// fmt.Println("Downloading audio file...")
				// data, err := downloadFile(audioURL)
				// if err != nil {
				// 	log.Fatalf("Error downloading file: %v", err)
				// }
				// fmt.Println("Audio file downloaded")

				// audioFile := tgbotapi.FileBytes{Name: "audio.mp3", Bytes: data}
				// audio := tgbotapi.NewAudio(update.Message.Chat.ID, audioFile)
				// _, err = bot.Send(audio)
				// if err != nil {
				// 	log.Fatalf("Error sending audio: %v", err)
				// }

				// photoURL := "https://muhfarkhans.com/img/app-image.jpg"
				// fmt.Println("Downloading photo file...")
				// dataPhoto, err := downloadFile(photoURL)
				// if err != nil {
				// 	log.Fatalf("Error downloading photo: %v", err)
				// }
				// fmt.Println("Photo file downloaded")

				// photoFile := tgbotapi.FileBytes{Name: "photo.jpg", Bytes: dataPhoto}
				// photo := tgbotapi.NewPhoto(update.Message.Chat.ID, photoFile)
				// _, err = bot.Send(photo)
				// if err != nil {
				// 	log.Fatalf("Error sending photo: %v", err)
				// }

				msg.Text = update.Message.Text
			}
			msg.ReplyToMessageID = update.Message.MessageID
		}

		if _, err := bot.Send(msg); err != nil {
			fmt.Println(err)
		}
	}
}

func downloadFile(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
