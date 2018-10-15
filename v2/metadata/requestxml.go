package metadata

import (
	"context"
	"encoding/xml"
	"net/http"
	"net/url"
	"time"

	"github.com/huin/goupnp/v2/errkind"
	"golang.org/x/net/html/charset"
)

func requestXml(
	ctx context.Context,
	loc *url.URL,
	defaultSpace string,
	doc interface{},
) error {
	req, err := http.NewRequest("GET", loc.String(), nil)
	if err != nil {
		return err
	}
	ctx, _ = context.WithTimeout(ctx, 3*time.Second)
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errkind.NewUnexpectedHTTPStatus(resp.StatusCode, resp.Status)
	}

	decoder := xml.NewDecoder(resp.Body)
	decoder.DefaultSpace = defaultSpace
	decoder.CharsetReader = charset.NewReaderLabel

	return decoder.Decode(doc)
}
