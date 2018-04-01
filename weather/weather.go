package weather

import (
	"os/user"
	"path/filepath"
	"time"

	"github.com/soumya92/barista/bar"
	"github.com/soumya92/barista/modules/weather"
	"github.com/soumya92/barista/modules/weather/openweathermap"
	"github.com/soumya92/barista/outputs"
	"github.com/soumya92/barista/pango"
	"github.com/soumya92/barista/pango/icons/typicons"
)

var spacer = pango.Span(" ", pango.XXSmall)

func home(path string) string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return filepath.Join(usr.HomeDir, path)
}

func Get() weather.Module {
	typicons.Load(home(".fonts/typicons"))

	// Weather information comes from OpenWeatherMap.
	// https://openweathermap.org/api.
	wthr := weather.New(
		openweathermap.CityID("524901").Build(),
		//openweathermap.Zipcode("94043", "US").Build(),
	).OutputFunc(func(w weather.Weather) bar.Output {
		iconName := ""
		switch w.Condition {
		case weather.Thunderstorm,
			weather.TropicalStorm,
			weather.Hurricane:
			iconName = "stormy"
		case weather.Drizzle,
			weather.Hail:
			iconName = "shower"
		case weather.Rain:
			iconName = "downpour"
		case weather.Snow,
			weather.Sleet:
			iconName = "snow"
		case weather.Mist,
			weather.Smoke,
			weather.Whirls,
			weather.Haze,
			weather.Fog:
			iconName = "windy-cloudy"
		case weather.Clear:
			if !w.Sunset.IsZero() && time.Now().After(w.Sunset) {
				iconName = "night"
			} else {
				iconName = "sunny"
			}
		case weather.PartlyCloudy:
			iconName = "partly-sunny"
		case weather.Cloudy, weather.Overcast:
			iconName = "cloudy"
		case weather.Tornado,
			weather.Windy:
			iconName = "windy"
		}
		if iconName == "" {
			iconName = "warning-outline"
		} else {
			iconName = "weather-" + iconName
		}
		return outputs.Pango(
			typicons.Icon(iconName), spacer,
			pango.Textf("%d℃", w.Temperature.C()),
			//pango.Span(" (provided by ", w.Attribution, ")", pango.XSmall),
		)
	})
	return wthr
}
