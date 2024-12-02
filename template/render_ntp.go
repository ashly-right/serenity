package template

import (
	"reflect"
	"time"

	"github.com/sagernet/sing-box/option"
	"github.com/sagernet/sing/common/json/badoption"
)

func (t *Template) RenderNTP(options *option.Options) error {
	if t.EnableNTP {
		options.NTP = &option.NTPOptions{
			Enabled:  true,
			Interval: badoption.Duration(time.Minute * 30),
		}

		options.NTP.ServerPort = 123
		options.NTP.Server = "time.apple.com"

		if t.CustomNTP != nil {
			if t.CustomNTP.Server != "" {
				options.NTP.Server = t.CustomNTP.Server
			}

			if t.CustomNTP.ServerPort != 0 {
				options.NTP.ServerPort = t.CustomNTP.ServerPort
			}

			if t.CustomNTP.Interval != 0 {
				options.NTP.Interval = t.CustomNTP.Interval
			}

			if !reflect.DeepEqual(t.CustomNTP.DialerOptions, option.DialerOptions{}) {
				options.NTP.DialerOptions = t.CustomNTP.DialerOptions
			}
		}

	}
	return nil
}
