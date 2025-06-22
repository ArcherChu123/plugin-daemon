package config

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func LoadFromEnv() error {
	cfg := &AppConfig{}
	v := reflect.ValueOf(cfg).Elem()
	t := reflect.TypeOf(*cfg)

	for i := 0; i < t.NumField(); i++ {
		fieldType := t.Field(i)
		fieldValue := v.Field(i)

		if fieldType.Type.Kind() == reflect.Struct {
			if err := loadStructFromEnv(fieldValue, fieldType.Type); err != nil {
				return err
			}
		}
	}

	AppCfg = cfg
	return nil
}

func loadStructFromEnv(v reflect.Value, t reflect.Type) error {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		val := v.Field(i)

		envKey := field.Tag.Get("env")
		if envKey == "" {
			continue
		}

		raw := os.Getenv(envKey)
		if raw == "" {
			if def := field.Tag.Get("default"); def != "" {
				raw = def
			}
		}

		if raw == "" && field.Tag.Get("required") == "true" {
			return fmt.Errorf("env variable %s is required but not set", envKey)
		}

		if raw == "" {
			continue
		}

		switch val.Kind() {
		case reflect.String:
			val.SetString(raw)
		case reflect.Int:
			num, err := strconv.Atoi(raw)
			if err != nil {
				return fmt.Errorf("invalid int in %s: %v", envKey, err)
			}
			val.SetInt(int64(num))
		case reflect.Bool:
			b, err := strconv.ParseBool(raw)
			if err != nil {
				return fmt.Errorf("invalid bool in %s: %v", envKey, err)
			}
			val.SetBool(b)
		case reflect.Slice:
			val.Set(reflect.ValueOf(strings.Split(raw, ",")))
		default:
			// optionally extend
		}
	}

	return nil
}
