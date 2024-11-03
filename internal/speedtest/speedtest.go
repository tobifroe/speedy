package speedtest

import (
	"io"
	"log"
	"net/http"
	"os"
	"speedy/internal/database"
	"time"
)

func Speedtest(url string) time.Duration {
	c := make(chan time.Duration)
	go func() {
		log.Print("Running speedtest...")
		out, err := os.Create("testfile")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer out.Close()

		t := time.Now()
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer resp.Body.Close()

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			log.Fatal(err.Error())
		}
		c <- time.Since(t)
	}()
	t := <-c
	downspeed := 100 / t.Seconds()
	db := database.New()
	result := database.TestResult{
		Duration:  t.Seconds(),
		DownSpeed: downspeed,
		Target:    url,
	}
	db.Create(&result)

	os.Remove("testfile")
	log.Printf("Speedtest finished. Time elapsed: %v", t.Seconds())
	return t
}
