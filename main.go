package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

// Step instalasi (jangan dilakukan lagi jika sudah!!!)

// 1. buka terminal di folder/directory yang mengandung file ini (mungkin namanya kokobot)

// 2. di terminal itu, ketik 2 hal berikut di terminal tanpa tanda $ (kalau belum). Hal ini hanya dilakukan sekali saja

// $ go mod init belajar/koko

// $ go get github.com/bwmarrin/discordgo

func main() {
	var Token string

	// Mengaktivasi discord bot. Tokennya ada di discord
	fmt.Print("Discord bot token: ")
	fmt.Scan(&Token)

	bot, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error membuat discord bot", err)
		return
	}

	// Menambah handler, yaitu fungsi `handleMessage` sebagai _callback_ kejadian `MessageCreate`.
	// fungsi handleMessage akan di-call saat bot menerima pesan baru
	bot.AddHandler(handleMessage)

	// Intents = Niat
	// Kita berniat untuk membaca message di guild, makanya ada `IntentsGuildMessages`
	// informasi lebih lanjut: https://discord.com/developers/docs/topics/gateway#gateway-intents
	bot.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages

	// Buka koneksi
	err = bot.Open()

	// jika ada error, print error itu, kemudian exit
	if err != nil {
		fmt.Println("error membuat koneksi", err)
		return
	}

	// Bagian ini tidak usah diperhatikan, copy paste aja gak-apa-apa untuk permulaian
	// Buat `channel` untuk mendengarkan sinyal seperti SIGINT, SIGTERM, Interrupt, dan Kill
	// Ketika diketik CTRL-C akan menghentikan proses
	fmt.Println("Bot sedang jalan... CTRL-C untuk exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Tutup koneksi
	bot.Close()
}

func handleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Kalo IDnya milik sendiri, biarin.
	// kita cuman mau dengerin orang lain
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "$random" {
		n := rand.Intn(100)
		_, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprint(n))
		if err != nil {
			fmt.Println(err)
		}
	}

}
