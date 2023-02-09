package libs

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

type LoginPayload struct {
	AppId       int    `json:"app_id"`
	AppServerId int    `json:"app_server_id"`
	LoginId     string `json:"login_id"`
}

type LoginSSOPayload struct {
	SessionKey string `json:"session_key"`
}

func GetDatadome() (string, []map[string]interface{}) {
	l := launcher.New().Headless(true)
	url := l.MustLaunch()
	browser := rod.New().ControlURL(url).MustConnect()
	link := "https://kiosgamer.co.id/app"
	page := browser.MustPage(link).MustWaitLoad()
	defer page.Close()
	defer browser.Close()
	cookies, _ := json.Marshal(page.MustCookies())
	var msg []map[string]interface{}
	err := json.Unmarshal([]byte(cookies), &msg)
	if err != nil {
		fmt.Println(err)
	}
	var datadome string
	for _, result := range msg {
		name := fmt.Sprint(result["name"])
		if name == "datadome" {
			datadome = fmt.Sprint(result["value"])
			break
		}
	}
	return datadome, msg
}

func Login(game_id, datadome string, cookies []map[string]interface{}) map[string]interface{} {
	data := LoginPayload{
		AppId:       100082,
		AppServerId: 0,
		LoginId:     game_id,
	}

	jsonBody, err := json.Marshal(data)

	if err != nil {
		println(err.Error())
	}

	// datadomeclientid := GenSeed(14)+"-"+GenSeed(64)+"-"+GenSeed(48)
	kgUrl := "https://kiosgamer.co.id/api/auth/player_id_login"
	req, err := http.NewRequest("POST", kgUrl, strings.NewReader(string(jsonBody)))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Host", "kiosgamer.co.id")
	req.Header.Add("Origin", "kiosgamer.co.id")
	req.Header.Add("Referer", "https://kiosgamer.co.id/app")
	req.Header.Add("Sec-Ch-Ua", "trailers")
	req.Header.Add("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Add("Sec-Ch-Ua-Platform", "Windows")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
	req.Header.Add("x-datadome-clientid", "3oZGR0vURo4tCF-v23ucuKZot317qptPXtsIom5C47XGqykaQlSQr0uYgDhEqzLk4iw2LDE0DWBw1Bb-418OOff6t0EjilNK9heQPElxu7o5mgYLhKcf72r1SR3qPoO6")
	for _, result := range cookies {
		name := fmt.Sprint(result["name"])
		value := fmt.Sprint(result["value"])
		req.AddCookie(&http.Cookie{Name: name, Value: value})
	}
	if err != nil {
		println(err.Error())
	}
	client := http.Client{}
	resp, err := client.Do(req)
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		println("Log => Error " + err.Error())
	}
	bodyString := string(bodyBytes)
	var result map[string]interface{}
	err = json.Unmarshal([]byte(bodyString), &result)
	if err != nil {
		fmt.Println(err)
	}
	// session_key := ""
	// for _, result := range cookies {
	// 	name := fmt.Sprint(result["name"])
	// 	value := fmt.Sprint(result["value"])
	// 	if name == "session_key" {
	// 		session_key = value
	// 		break
	// 	}
	// 	req.AddCookie(&http.Cookie{Name: name, Value: value})
	// }
	return result

}

func LoginSSO(session_key string, cookies []map[string]interface{}) {
	url := "https://kiosgamer.co.id/api/auth/sso"
	data := LoginSSOPayload{
		SessionKey: session_key,
	}
	jsonBody, _ := json.Marshal(data)
	println(string(jsonBody))

	req, _ := http.NewRequest("POST", url, strings.NewReader(string(jsonBody)))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Host", "kiosgamer.co.id")
	req.Header.Add("Origin", "kiosgamer.co.id")
	req.Header.Add("Referer", "https://kiosgamer.co.id/app")
	req.Header.Add("Sec-Ch-Ua", "trailers")
	req.Header.Add("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Add("Sec-Ch-Ua-Platform", "Windows")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
	req.Header.Add("x-datadome-clientid", "3oZGR0vURo4tCF-v23ucuKZot317qptPXtsIom5C47XGqykaQlSQr0uYgDhEqzLk4iw2LDE0DWBw1Bb-418OOff6t0EjilNK9heQPElxu7o5mgYLhKcf72r1SR3qPoO6")
	for _, result := range cookies {
		name := fmt.Sprint(result["name"])
		value := fmt.Sprint(result["value"])
		if name == "session_key" {
			session_key = value
		}
		req.AddCookie(&http.Cookie{Name: name, Value: value})
	}
	// req.AddCookie(&http.Cookie{Name: "session_key", Value: session_key})
	client := http.Client{}
	resp, err := client.Do(req)
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		println("Log => Error " + err.Error())
	}
	bodyString := string(bodyBytes)
	println(bodyString)
}

func GetCsrf(session_key string) string {
	kgUrl := "https://kiosgamer.co.id/api/preflight"
	req, err := http.NewRequest("POST", kgUrl, strings.NewReader(""))
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Sec-Ch-Ua", "trailers")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
	req.Header.Add("Sec-Ch-Ua-Platform", "Windows")
	req.Header.Add("Origin", "https://kiosgamer.co.id")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Referer", "https://kiosgamer.co.id/app")
	req.AddCookie(&http.Cookie{Name: "_ga", Value: "GA1.3.277950594.1638345978"})
	req.AddCookie(&http.Cookie{Name: "source", Value: "pc"})
	req.AddCookie(&http.Cookie{Name: "b.vnpopup.1", Value: "1"})
	req.AddCookie(&http.Cookie{Name: "__csrf__", Value: ""})
	req.AddCookie(&http.Cookie{Name: "_gat", Value: "1"})
	req.AddCookie(&http.Cookie{Name: "session_key", Value: session_key})
	if err != nil {
		println(err.Error())
	}
	client := http.Client{}
	resp, err := client.Do(req)
	csrf := resp.Header.Get("Set-Cookie")
	startCsrf := strings.Index(csrf, "=")
	csrf = csrf[startCsrf+1:]
	endCsrf := strings.Index(csrf, ";")
	csrf = csrf[:endCsrf]
	return csrf
}

func GenSeed(length int) string {
	var alphabet []rune = []rune("abcdefghijklmnopqrstuvwxyz1234567890")
	alphabetSize := len(alphabet)
	var sb strings.Builder

	for i := 0; i < length; i++ {
		ch := alphabet[rand.Intn(alphabetSize)]
		sb.WriteRune(ch)
	}

	s := sb.String()
	return s
}

func Recaptcha(session_key string) (string, string) {
	genToken := GenSeed(10) + "-" + GenSeed(10) + "-" + GenSeed(10) + "-" + GenSeed(10) + "-" + GenSeed(10)
	captchaLink := "https://gop.captcha.garena.com/image?key=" + genToken
	println(captchaLink)

	res, err := http.Get(captchaLink)
	if err != nil {
		println("http.Get -> %v", err)
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	ioutil.WriteFile("captcha.png", data, 0666)
	cap := Normal{
		Numberic:      4,
		MinLen:        4,
		MaxLen:        20,
		Phrase:        true,
		CaseSensitive: true,
		Lang:          "en",
		File:          "captcha.png",
	}
	client := NewClient("33e6be9ebc115b61681c1d00c6340db5")
	code, err := client.Solve(cap.ToRequest())
	if err != nil {
		if err == ErrTimeout {
			println("Timeout")
		} else if err == ErrApi {
			println("API error")
		} else if err == ErrNetwork {
			println("Network error")
		} else {
			log.Fatal(err)
		}
		return "error", "error"
	} else {
		return code, genToken
	}
}
