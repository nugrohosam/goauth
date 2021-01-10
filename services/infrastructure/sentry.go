package infrastructure

import (
	"fmt"
	"time"

	"github.com/spf13/viper"

	sentry "github.com/getsentry/sentry-go"
)

// PrepareSentry ...
func PrepareSentry() {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: viper.GetString("sentry_dsn."),
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}
}

// CaptureMessage ...
func CaptureMessage(err error) {
	defer sentry.Flush(2 * time.Second)
	sentry.CaptureMessage(err.Error())
}
