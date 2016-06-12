package bot

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ChimeraCoder/anaconda"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func postJson(url string, input interface{}, output interface{}) error {
	data, err := json.Marshal(input)
	if err != nil {
		return errors.Wrapf(err, "failed to decode json. data: %v", input)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return errors.Wrapf(err, "POST message: %v", data)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrapf(err, "failed to read response. response: %v", resp)
	}

	err = json.Unmarshal(body, &output)
	if err != nil {
		return errors.Wrapf(err, "failed to encode json. json: %v", &output)
	}

	return nil
}

func getTwitterApi() *anaconda.TwitterApi {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return nil
	}

	consumer_key := viper.GetString("twitter.consumer_key")
	consumer_sec := viper.GetString("twitter.consumer_key_secret")
	access_token := viper.GetString("twitter.access_token")
	access_secret := viper.GetString("twitter.access_token_secret")

	anaconda.SetConsumerKey(consumer_key)
	anaconda.SetConsumerSecret(consumer_sec)

	api := anaconda.NewTwitterApi(access_token, access_secret)
	if api == nil {
		return nil
	}
	return api
}
