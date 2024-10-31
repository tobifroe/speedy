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

		out, err := os.Create("testfile")
		if err != nil {
			log.Fatalf(err.Error())
		}
		defer out.Close()

		t := time.Now()
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf(err.Error())
		}
		defer resp.Body.Close()

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			log.Fatalf(err.Error())
		}
		c <- time.Now().Sub(t)
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
	return t
}
