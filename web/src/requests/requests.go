package requests

import (
	"io"
	"net/http"
	"web/src/cookies"
)

// RequestsWithAuthentication - Method to request API with authentication
func RequestsWithAuthentication(r *http.Request, method, url string, datas io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(method, url, datas)
	if err != nil {
		return nil, err
	}

	cookie, _ := cookies.Read(r)
	request.Header.Add("Authorization", "Bearer "+cookie["token"])

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
