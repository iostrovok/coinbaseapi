package viper

import (
	"log"
	"os"
	"strings"
	"time"

	cf "github.com/iostrovok/go-convert"
)

type Viper struct {
	log map[string]any
}

func New() *Viper {
	return &Viper{
		log: map[string]any{},
	}
}

func (v Viper) Log() map[string]any {
	return v.log
}

func (v Viper) GetInt(name string) int {
	v.log[name] = os.Getenv(name)
	return cf.Int(os.Getenv(name))
}

func (v Viper) GetInt64(name string) int64 {
	v.log[name] = os.Getenv(name)
	return cf.Int64(os.Getenv(name))
}

func (v Viper) GetMaxInt(baseValue int, name string) int {
	v.log[name] = os.Getenv(name)

	if baseValue >= cf.Int(os.Getenv(name)) {
		return baseValue
	}

	return cf.Int(os.Getenv(name))
}

func (v Viper) GetUint32(name string) uint32 {
	v.log[name] = os.Getenv(name)
	return cf.Uint32(os.Getenv(name))
}

func (v Viper) GetBool(name string) bool {
	v.log[name] = os.Getenv(name)
	return cf.Bool(os.Getenv(name))
}

func (v Viper) GetString(name string) string {
	if strings.Contains(strings.ToLower(name), "pass") {
		v.log[name] = printPassword(os.Getenv(name))
	} else {
		v.log[name] = os.Getenv(name)
	}

	return cf.String(os.Getenv(name))
}

func (v Viper) GetFloat64(name string) float64 {
	v.log[name] = os.Getenv(name)
	return cf.Float64(os.Getenv(name))
}

func (v Viper) Duration(name string, defaultValue time.Duration) time.Duration {
	if timeOutRawValue := v.GetString(name); timeOutRawValue != "" {
		durationValue, errParseDuration := time.ParseDuration(timeOutRawValue)
		if errParseDuration != nil {
			return defaultValue
		}
		if durationValue > 0 {
			return durationValue
		}
	}

	return defaultValue
}

func (v Viper) Time(format, name string, defaultValue time.Time) time.Time {
	if timeOutRawValue := v.GetString(name); timeOutRawValue != "" {
		value, err := time.Parse(format, timeOutRawValue)
		if err == nil {
			return value
		}
		log.Printf("Time parsing error for '%s'", name)
	}

	return defaultValue
}

func printPassword(str string) string {
	if len(str) < 10 {
		return "..."
	}

	return str[:3] + "...." + str[len(str)-2:]
}
