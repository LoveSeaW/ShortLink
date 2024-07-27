package urltool

import (
	"errors"
	"net/url"
	"path"
)

func GetBasePath(targetUrl string) (string, error) {
	myUrl, err := url.Parse(targetUrl)
	if err != nil {
		return "", err
	}
	if len(myUrl.Host) == 0 {
		return "", errors.New("no host in targetUrl")
	}
	basePath := path.Base(myUrl.Path)
	return basePath, nil
}
