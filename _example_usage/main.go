package main

import (
	"net/http"

	"github.com/rogierlommers/go-varnish-purger"
	"github.com/sirupsen/logrus"
)

func main() {
	varnishAPI, err := varnish.NewInvalidator("http://localhost", 80, true)
	if err != nil {
		logrus.Fatal(err)
	}

	varnishAPI.BeforeRequest(func(request *http.Request) {
		// for example, set auth header for varnish: request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", "my-special-token"))
		logrus.Info("about to make request...")
	})

	targetPath := "static-content/myfile.jpg"
	if err := varnishAPI.InvalidatePath(targetPath); err != nil {
		logrus.Fatalf("error invalidating url: %s [url: %s]", err, targetPath)
	}

	logrus.Info("succesfully purged URL...")
}
