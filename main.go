package main

import (
	"fmt"
	"log"
	"time"

	"github.com/micmonay/keybd_event"
)

func main() {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Copy the text you want to be pasted and then get your page ready by making it the focus of the screen/desktop - you have 5 seconds.")
	time.Sleep(5 * time.Second) //time to get page ready
	i := 1
	for i == 1 {
		time.Sleep(200 * time.Millisecond)
		//set keys
		kb.HasCTRL(false)
		kb.SetKeys(keybd_event.VK_DOWN)

		//launch
		err = kb.Launching()
		if err != nil {
			log.Fatal(err)
		}
		kb.HasCTRL(true)
		kb.SetKeys(keybd_event.VK_V)
		err = kb.Launching()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("done")
	}
}
