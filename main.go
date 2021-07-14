package main

// Simple Telegram bot using Go

import (
	"os"

	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/PaulSonOfLars/gotgbot/handlers"
	"github.com/PaulSonOfLars/gotgbot/handlers/Filters"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	log := zap.NewProductionEncoderConfig()
	log.EncodeLevel = zapcore.CapitalLevelEncoder
	log.EncodeTime = zapcore.RFC3339TimeEncoder

	logger := zap.New(zapcore.NewCore(zapcore.NewConsoleEncoder(log), os.Stdout, zap.InfoLevel))

	updater, err := gotgbot.NewUpdater(logger, "Token")
	if err != nil {
		logger.Panic("UPDATER FAILED TO START")
		return
	}
	logger.Sugar().Info("UPDATER STARTED SUCCESSFULLY")
	updater.StartCleanPolling()
	updater.Dispatcher.AddHandler(handlers.NewMessage(Filters.Text, echo))
	updater.Dispatcher.AddHandler(handlers.NewCommand("ban", banUser))
	updater.Dispatcher.AddHandler(handlers.NewCommand("unban", unbanUser))

	updater.Idle()
}

// Simple Telegram Echo bot
func echo(b ext.Bot, u *gotgbot.Update) error {
	b.SendMessage(u.EffectiveChat.Id, u.EffectiveMessage.Text)
	return nil
}

// Simple Telegram Ban user
func banUser(b ext.Bot, u *gotgbot.Update) error {
	b.KickChatMember(u.EffectiveChat.Id, u.EffectiveMessage.ReplyToMessage.From.Id)
	return nil
}

// Simple Telegram unban user
func unbanUser(b ext.Bot, u *gotgbot.Update) error {
	b.UnbanChatMember(u.EffectiveChat.Id, u.EffectiveMessage.ReplyToMessage.From.Id)
	u.EffectiveMessage.ReplyText("Unbanned that user")
	return nil
}
