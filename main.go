package main

import (
	vault "github.com/hashicorp/vault/api"
	"github.com/jsiebens/blinkt-vault-controller/pgk/blinkt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config := vault.DefaultConfig()

	client, _ := vault.NewClient(config)

	brightness := 0.5
	bl := blinkt.NewBlinkt(brightness)

	bl.Setup()

	Delay(100)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)

	for {
		select {
		case <-signalChan:
			bl.Cleanup()
			return
		default:

			a, err := client.Sys().Health()

			if err == nil {

				if a.Sealed {
					status, _ := client.Sys().SealStatus()

					if status.T == 0 {
						bl.SetAllHex("FF0000")
					} else {

						percentage := float32(status.Progress) / float32(status.T)

						x := int(percentage * 8)

						for i := 0; i < 8; i++ {
							if (x - 1) >= i {
								bl.SetPixelBrightness(i, brightness)
								bl.SetPixelHex(i, "00FF00")
							} else {
								bl.SetPixelBrightness(i, brightness)
								bl.SetPixelHex(i, "FF0000")
							}
						}
					}

				} else {
					bl.SetAllHex("00FF00")
				}

				bl.Show()

			} else {
				bl.FlashAll(2, "FF0000")
			}

			Delay(5000)
		}
	}

}

// Delay maps to time.Sleep, for ms milliseconds
func Delay(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}
