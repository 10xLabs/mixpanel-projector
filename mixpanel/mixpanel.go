package mixpanel

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Mixpanel ...
type Mixpanel interface {
	Track(properties interface{}) error
}

type mixpanel struct {
	ApiUrl string
	Client *http.Client
}

// NewMixpanel ...
func NewMixpanel(ApiUrl string) Mixpanel {
	return &mixpanel{
		ApiUrl: ApiUrl,
		Client: http.DefaultClient,
	}
}

func (m *mixpanel) Track(properties interface{}) error {
	d, _ := json.Marshal(properties)
	url := m.ApiUrl + "/track/" + "?ip=1&data=" + base64.StdEncoding.EncodeToString(d)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println("BODY: ", body)

	return nil
}
