package appenginesdk

import (
	"bytes"
	"net/http"
	"regexp"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"

	"golang.org/x/net/context"
)

func init() {
	http.HandleFunc("/", redirect)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	body := apiPage(c)
	u := versionURL(c, body)
	http.Redirect(w, r, u, 301)
}

func versionURL(c context.Context, body string) string {
	re := regexp.MustCompile(`https:\/\/storage\.googleapis\.com\/appengine-sdks\/featured\/go_appengine_sdk_linux_amd64-[1-2]\.[0-9]{1,2}\.[0-9]{1,2}\.zip`)
	return re.FindString(body)
}

func apiPage(c context.Context) string {
	client := urlfetch.Client(c)
	resp, _ := client.Get("https://cloud.google.com/appengine/downloads")

	defer resp.Body.Close()
	b := new(bytes.Buffer)
	b.ReadFrom(resp.Body)
	return string(b.Bytes())
}
