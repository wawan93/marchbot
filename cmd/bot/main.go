package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	tgbot "github.com/wawan93/bot-framework"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {
	token := os.Getenv("TOKEN")
	log.Printf("token=%v", token)

	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	api.Debug = false

	log.Printf("logged in as %v", api.Self.UserName)

	bot := tgbot.NewBotFramework(api)
	updates := getUpdatesChannel(api, "")

	bot.RegisterCommand("Хочу еще комплимент", RandVoice, 0)

	bot.HandleUpdates(updates)
}

func getUpdatesChannel(api *tgbotapi.BotAPI, webhookAddress string) tgbotapi.UpdatesChannel {
	var updates tgbotapi.UpdatesChannel
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, _ = api.GetUpdatesChan(u)
	return updates
}

func RandVoice(bot *tgbot.BotFramework, update *tgbotapi.Update) error {
	chat := bot.GetChatID(update)
	voices := []string{
		"AwACAgIAAxkDAAIGXV5ii4RAiE84_cuprVFnyMfR_KaeAALUBgAC9HwRS20S5kR3x7zeGAQ",
		"AwACAgIAAxkDAAIGXl5ii4V_PyZTFeaQnKBmSLCx2fJGAALYBgAC-PIZSziPPuYTIan8GAQ",
		"AwACAgIAAxkDAAIGX15ii4joMqHzsUehHcEEhLgjEWEhAALVBgAC9HwRS4f56iiUazFJGAQ",
		"AwACAgIAAxkDAAIGYF5ii4oIZMkilFtAcSF2cjxrmPRhAALZBgAC-PIZSylSw6_FhFiWGAQ",
		"AwACAgIAAxkDAAIGYV5ii41rRyhbumO98vwWzZEnu4nBAALWBgAC9HwRS3mxHjF207mrGAQ",
		"AwACAgIAAxkDAAIGYl5ii5IFS5_ZB-v4x_52ZbF2pKcTAALXBgAC9HwRS7VkHT02-JGXGAQ",
		"AwACAgIAAxkDAAIGY15ii5S_SHvRy6cG2KRNqubjMII1AALYBgAC9HwRS0GJN3Q5PoOKGAQ",
		"AwACAgIAAxkDAAIGZF5ii5d6KkuQR_Y1xdCqF3vJkUFMAALZBgAC9HwRS5KnLUUeyeb9GAQ",
		"AwACAgIAAxkDAAIGZV5ii5cGXZAhpBYPyyRKc3GM9eBtAALaBgAC9HwRS_mHnNoP9Fl3GAQ",
		"AwACAgIAAxkDAAIGZl5ii5h2ogt1TmltVO44Yq-8NszkAALbBgAC9HwRSzzYYDQrloPmGAQ",
		"AwACAgIAAxkDAAIGZ15ii5lH8EGurxPEttNsA4UZQCONAALcBgAC9HwRS5_73WRAgMTsGAQ",
		"AwACAgIAAxkDAAIGaF5ii5mgj6IqfWqOsE7lWQdKtfAyAALaBgAC-PIZS3jRfYfL1uOoGAQ",
		"AwACAgIAAxkDAAIGaV5ii5p1Y-FFYCpNqzzkFgcShXgcAALbBgAC-PIZS7iXQEFbFtZDGAQ",
		"AwACAgIAAxkDAAIGal5ii5qK-eF6Dgw2uBpIOfbstC6kAALdBgAC9HwRS6LGrAvjXLxDGAQ",
		"AwACAgIAAxkDAAIGa15ii5uT3-mSJnczfkaTMmcM8SO8AALeBgAC9HwRSysXkzp-XJjQGAQ",
		"AwACAgIAAxkDAAIGbF5ii52bb_5oq4dCb0BGYtIh9Dg4AALfBgAC9HwRS5SOn7UbcV1mGAQ",
		"AwACAgIAAxkDAAIGbV5ii59epKnnyZCpg6G4DGktfdtcAALgBgAC9HwRSyPUngnp_wT3GAQ",
		"AwACAgIAAxkDAAIGbl5ii6N00-conrUr9ewpUveLIkBDAALcBgAC-PIZS79mpDKWA9ArGAQ",
		"AwACAgIAAxkDAAIGb15ii6YXZKmZ2b2APkHOp1zr2RhKAALdBgAC-PIZSzVRiIbd9362GAQ",
		"AwACAgIAAxkDAAIGcF5ii6nmDAlpkUfIvuYgF92dcXOzAALeBgAC-PIZS5uvllL_Qx5IGAQ",
		"AwACAgIAAxkDAAIGcV5ii69lpFWl83kfL3be5GMnUUIZAALfBgAC-PIZS7TmVhmYY-edGAQ",
		"AwACAgIAAxkDAAIGcl5ii7A6oEMccDUWMU6RPxk0nQL5AALhBgAC-PIZS3y9NQRyP0x2GAQ",
		"AwACAgIAAxkDAAIGc15ii7Hti1oWpN6Q5ntcflDb5KJ-AALhBgAC9HwRS8jQvaXse7B2GAQ",
		"AwACAgIAAxkDAAIGdF5ii7LjDrFajxUhZTxAJAWgW46hAALiBgAC9HwRS2V_yZZKVH8zGAQ",
		"AwACAgIAAxkDAAIGdV5ii7OPZnV22YJXGcg6-QSzJikeAALiBgAC-PIZS30QmxRWdK3QGAQ",
		"AwACAgIAAxkDAAIGdl5ii7uaj8ez6JaDWt49Ubs5cuw-AALjBgAC9HwRSwYXFzIIQ38gGAQ",
		"AwACAgIAAxkDAAIGd15ii70w7n16MseNoWvjpqIlQdWPAALjBgAC-PIZS-IszeXAXmRlGAQ",
		"AwACAgIAAxkDAAIGeF5ii79wmzCVCbpGU-X3yVOJlt_MAALkBgAC-PIZS_LdjznVygmwGAQ",
		"AwACAgIAAxkDAAIGeV5ii811U2t7QNgAASeqEYrjuEIHcAAC5QYAAvjyGUtAnG29jcLxXhgE",
		"AwACAgIAAxkDAAIGel5ii84vU86gqxQFf28sV6pY9s2QAALnBgAC-PIZS6a3tg3fg7YrGAQ",
		"AwACAgIAAxkDAAIGe15ii9FWL8t3f12ibTvlktXLA1-2AALoBgAC-PIZS0WZgTkVl14BGAQ",
		"AwACAgIAAxkDAAIGfF5ii9hVSu1ZGFBjhEZt8651UegUAALpBgAC-PIZSw1EBMmbE2RtGAQ",
		"AwACAgIAAxkDAAIGfl5ii-ZQaw9y2rlHxxv4ljida_RXAAL3CAACq1IYS6ZZjUQKtlOtGAQ",
		"AwACAgIAAxkDAAIGf15ii-eWn0IrU1YxweMo0OdqQAu_AALlBgAC9HwRSzWiQDhZL2_LGAQ",
		"AwACAgIAAxkDAAIGgF5ii-e4hl_P50VL4Bll1ItfJGIGAALrBgAC-PIZS5tZsa0-C55sGAQ",
		"AwACAgIAAxkDAAIGgV5ii-2QVDAXPw-G1q2uOjzDZdWnAALsBgAC-PIZS-53BxTfYkQOGAQ",
		"AwACAgIAAxkDAAIGgl5ii_sMvg_Cf88g9OlS958wF851AALtBgAC-PIZS0IgB6DoC6k1GAQ",
		"AwACAgIAAxkDAAIGg15ijAABY7QYZsT0-O7MIBkYsv4dCgAC7gYAAvjyGUv-Q2WSp1542xgE",
		"AwACAgIAAxkDAAIGhF5ijA2doA5hel52RhDI1gIrzmMZAALmBgAC9HwRS6iANWr4i0leGAQ",
		"AwACAgIAAxkDAAIGhV5ijBGkrXWHbFAiCpKFTKK8ySLGAALnBgAC9HwRS1RcL8nEynEUGAQ",
		"AwACAgIAAxkDAAIGhl5ijBM77TqI0pOrxUOwr74O1IqbAAKbBgACIm8QS1y9eAFmuRvLGAQ",
		"AwACAgIAAxkDAAIGh15ijBYi_QABCBrIx4LqxykV7MjQkAACnAYAAiJvEEtOSgSfLjbPaxgE",
		"AwACAgIAAxkDAAIGiF5ijBdup8t5N0dVIy2JLnzzwAB-QAC7wYAAvjyGUsNxhTwq28gBBgE",
		"AwACAgIAAxkDAAIGiV5ijBn90a3eJBDHVXMh9gv4uzXxAALoBgAC9HwRSyHNCpI6b6l9GAQ",
		"AwACAgIAAxkDAAIGil5ijBqcOdCMSc_3X1vnpk29zEjlAALpBgAC9HwRS_l-7NIcF5jdGAQ",
		"AwACAgIAAxkDAAIGi15ijBv2HwmSgJDIR6HTvUw7ozx9AALwBgAC-PIZS8KsNuyRgs0lGAQ",
		"AwACAgIAAxkDAAIGjF5ijBwjZFX5C0N5W5_p26Z4gRO7AALxBgAC-PIZS3X2l0bOd4_IGAQ",
		"AwACAgIAAxkDAAIGjV5ijB6Pcg4kk5fogDjQMofZ56p4AALqBgAC9HwRS2-Dw-6cdLbEGAQ",
		"AwACAgIAAxkDAAIGjl5ijB8-UeMpfRmsVZ4HKEq7FSLTAALrBgAC9HwRSzORwA2fOPXsGAQ",
		"AwACAgIAAxkDAAIGj15ijCKrqGo45eMfzY8y7_e06Np3AALyBgAC-PIZS0SYfZWaiEuLGAQ",
		"AwACAgIAAxkDAAIGkF5ijCNfFunhAAEBLlLjzWVKRTf75AAC8wYAAvjyGUsJCvr_46PJxBgE",
		"AwACAgIAAxkDAAIGkV5ijCUcj7nOWblySKB63yfKgKdeAAL0BgAC-PIZS9mvdY5xDTzKGAQ",
		"AwACAgIAAxkDAAIGkl5ijCYVgTHP4I0gD5a74uSjE2DeAAL1BgAC-PIZS5-zrp6iw6zZGAQ",
		"AwACAgIAAxkDAAIGk15ijCe01XBazfHKYH6zJUCLt6EgAALsBgAC9HwRS5jyShLUF-_nGAQ",
		"AwACAgIAAxkDAAIGlF5ijCllBsBuIuPUp8iiyt26brDDAAL2BgAC-PIZSzCFLjuWFGazGAQ",
		"AwACAgIAAxkDAAIGlV5ijCrMeu8_OAOjayAGLuy1d2btAAL3BgAC-PIZS6hAXYwVmkq7GAQ",
		"AwACAgIAAxkDAAIGll5ijCuYmzxMhubFSf_36cWA0IlCAALtBgAC9HwRSycUk-9Hw-cLGAQ",
		"AwACAgIAAxkDAAIGl15ijCttuGvkQo0uKyjXv3GvzEPmAAL4BgAC-PIZS6kZzdj385SyGAQ",
		"AwACAgIAAxkDAAIGmF5ijCxBUwELja-bctXjhN0ZuNyWAALuBgAC9HwRS50l1JOQxlToGAQ",
		"AwACAgIAAxkDAAIGmV5ijCzTLmcscjhtTl4ZiWQKqOsqAALvBgAC9HwRSwigggKBbQgUGAQ",
		"AwACAgIAAxkDAAIGml5ijC0By1zmzfxk6e1nbdD7gVGQAAL5BgAC-PIZS4zP0Bb53ieaGAQ",
		"AwACAgIAAxkDAAIGm15ijC0tUUt8Y-SG3PHJG5M9c5B4AAL6BgAC-PIZSzlQIWV-7NYQGAQ",
		"AwACAgIAAxkDAAIGnF5ijC0ny5U4o83mHcN46ysxRvBpAAKeBgACIm8QS3xMnC3u3s1vGAQ",
		"AwACAgIAAxkDAAIGnV5ijC31zL8aheTJBck57dQ143AyAAKfBgACIm8QS1fIj5dKkVxVGAQ",
		"AwACAgIAAxkDAAIGnl5ijC9lv1vDXjsqL0ed-MEFDqW7AALwBgAC9HwRS2KdymMQKZLGGAQ",
		"AwACAgIAAxkDAAIGn15ijDEytF5pcd0jSG6YPGWEQR9NAALxBgAC9HwRS6pjUJN8WqYVGAQ",
	}
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	voice := voices[r.Intn(len(voices))]

	msg := tgbotapi.NewVoiceShare(chat, voice)
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard([]tgbotapi.KeyboardButton{
		{Text: "Хочу еще комплимент"},
	})

	_, err := bot.Send(msg)
	return err
}
