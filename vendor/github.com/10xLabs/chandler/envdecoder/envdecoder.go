package envdecoder

import (
	"context"
	"encoding"
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/10xLabs/chandler/tracer"
)

var tr = tracer.New()

// ErrInvalidTarget indicates that the target value passed to
// Decode is invalid.  Target must be a non-nil pointer to a struct.
var ErrInvalidTarget = errors.New("target must be non-nil pointer to struct that has at least one exported field with a valid env tag.")
var ErrNoTargetFieldsAreSet = errors.New("none of the target fields were set from environment variables")

// FailureFunc is called when an error is encountered during a MustDecode
// operation. It prints the error and terminates the process.
//
// This variable can be assigned to another function of the user-programmer's
// design, allowing for graceful recovery of the problem, such as loading
// from a backup configuration file.
var FailureFunc = func(err error) {
	log.Fatalf("envdecode: an error was encountered while decoding: %v\n", err)
}

// Decoder is the interface implemented by an object that can decode an
// environment variable string representation of itself.
type Decoder interface {
	Decode(string) error
}

// Decode environment variables into the provided target.  The target
// must be a non-nil pointer to a struct.  Fields in the struct must
// be exported, and tagged with an "env" struct tag with a value
// containing the name of the environment variable.  An error is
// returned if there are no exported members tagged.
//
// Default values may be provided by appending ",default=value" to the
// struct tag.  Required values may be marked by appending ",required"
// to the struct tag.  It is an error to provide both "default" and
// "required". Strict values may be marked by appending ",strict" which
// will return an error on Decode if there is an error while parsing.
// If everything must be strict, consider using StrictDecode instead.
//
// All primitive types are supported, including bool, floating point,
// signed and unsigned integers, and string.  Boolean and numeric
// types are decoded using the standard strconv Parse functions for
// those types.  Structs and pointers to structs are decoded
// recursively.  time.Duration is supported via the
// time.ParseDuration() function and *url.URL is supported via the
// url.Parse() function. Slices are supported for all above mentioned
// primitive types. Semicolon is used as delimiter in environment variables.
func Decode(ctx context.Context, target interface{}) (err error) {
	nFields, err := decode(target, false)
	if err != nil {
		return err
	}

	// if we didn't do anything - the user probably did something
	// wrong like leave all fields unexported.
	if nFields == 0 {
		return ErrNoTargetFieldsAreSet
	}

	return nil
}

func decode(target interface{}, strict bool) (int, error) {
	s := reflect.ValueOf(target)
	if s.Kind() != reflect.Ptr || s.IsNil() {
		return 0, ErrInvalidTarget
	}

	s = s.Elem()
	if s.Kind() != reflect.Struct {
		return 0, ErrInvalidTarget
	}

	t := s.Type()
	setFieldCount := 0
	for i := 0; i < s.NumField(); i++ {
		// Localize the umbrella `strict` value to the specific field.
		strict := strict

		f := s.Field(i)

		switch f.Kind() {
		case reflect.Ptr:
			if f.Elem().Kind() != reflect.Struct {
				break
			}

			f = f.Elem()
			fallthrough

		case reflect.Struct:
			if !f.Addr().CanInterface() {
				continue
			}

			ss := f.Addr().Interface()
			_, custom := ss.(Decoder)
			if custom {
				break
			}

			n, err := decode(ss, strict)
			if err != nil {
				return 0, err
			}
			setFieldCount += n
		}

		if !f.CanSet() {
			continue
		}

		tag := t.Field(i).Tag.Get("env")
		if tag == "" {
			continue
		}

		parts := strings.Split(tag, ",")
		env := os.Getenv(parts[0])
		if env == "" {
			s := variablesFromSecrets()
			env = s[parts[0]]
			os.Setenv(parts[0], env)
		}

		required := false
		hasDefault := false
		defaultValue := ""

		for _, o := range parts[1:] {
			if !required {
				required = strings.HasPrefix(o, "required")
			}
			if strings.HasPrefix(o, "default=") {
				hasDefault = true
				defaultValue = o[8:]
			}
			if !strict {
				strict = strings.HasPrefix(o, "strict")
			}
		}

		if required && hasDefault {
			panic(`envdecode: "default" and "required" may not be specified in the same annotation`)
		}
		if env == "" && required {
			return 0, fmt.Errorf("the environment variable \"%s\" is missing", parts[0])
		}
		if env == "" {
			env = defaultValue
		}
		if env == "" {
			continue
		}

		setFieldCount++

		unmarshaler, implementsUnmarshaler := f.Addr().Interface().(encoding.TextUnmarshaler)
		decoder, implmentsDecoder := f.Addr().Interface().(Decoder)
		if implmentsDecoder {
			if err := decoder.Decode(env); err != nil {
				return 0, err
			}
		} else if implementsUnmarshaler {
			if err := unmarshaler.UnmarshalText([]byte(env)); err != nil {
				return 0, err
			}
		} else if f.Kind() == reflect.Slice {
			decodeSlice(&f, env)
		} else {
			if err := decodePrimitiveType(&f, env); err != nil && strict {
				return 0, err
			}
		}
	}

	return setFieldCount, nil
}

func decodeSlice(f *reflect.Value, env string) {
	parts := strings.Split(env, ";")

	values := parts[:0]
	for _, x := range parts {
		if x != "" {
			values = append(values, strings.TrimSpace(x))
		}
	}

	valuesCount := len(values)
	slice := reflect.MakeSlice(f.Type(), valuesCount, valuesCount)
	if valuesCount > 0 {
		for i := 0; i < valuesCount; i++ {
			e := slice.Index(i)
			decodePrimitiveType(&e, values[i])
		}
	}

	f.Set(slice)
}

func decodePrimitiveType(f *reflect.Value, env string) error {
	switch f.Kind() {
	case reflect.Bool:
		v, err := strconv.ParseBool(env)
		if err != nil {
			return err
		}
		f.SetBool(v)

	case reflect.Float32, reflect.Float64:
		bits := f.Type().Bits()
		v, err := strconv.ParseFloat(env, bits)
		if err != nil {
			return err
		}
		f.SetFloat(v)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if t := f.Type(); t.PkgPath() == "time" && t.Name() == "Duration" {
			v, err := time.ParseDuration(env)
			if err != nil {
				return err
			}
			f.SetInt(int64(v))
		} else {
			bits := f.Type().Bits()
			v, err := strconv.ParseInt(env, 0, bits)
			if err != nil {
				return err
			}
			f.SetInt(v)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		bits := f.Type().Bits()
		v, err := strconv.ParseUint(env, 0, bits)
		if err != nil {
			return err
		}
		f.SetUint(v)

	case reflect.String:
		f.SetString(env)

	case reflect.Ptr:
		if t := f.Type().Elem(); t.Kind() == reflect.Struct && t.PkgPath() == "net/url" && t.Name() == "URL" {
			v, err := url.Parse(env)
			if err != nil {
				return err
			}
			f.Set(reflect.ValueOf(v))
		}
	}
	return nil
}
