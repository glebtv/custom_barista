package weather

import (
	"time"

	"github.com/glebtv/custom_barista/utils"
	"barista.run/bar"
	"barista.run/modules/weather"
	"barista.run/modules/weather/openweathermap"
	"barista.run/outputs"
	"barista.run/pango"
)

func Get(cityId string) *weather.Module {
	// Weather information comes from OpenWeatherMap.
	// https://openweathermap.org/api.
	wthr := weather.New(
		openweathermap.CityID(cityId).Build(),
	).Output(func(w weather.Weather) bar.Output {
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
			pango.Icon("typecn-"+iconName), utils.Spacer,
			pango.Textf("%.1f℃", w.Temperature.Celsius()),
		)
	})
	return wthr
}
