package gitlab

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/outillage/release-notary/internal/release"
	"github.com/stretchr/testify/assert"
)

func TestPublishExistingRelease(t *testing.T) {
	projectID := 1
	tagName := "v1.0.0"

	newReleaseNotes := "test"
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		url := fmt.Sprintf("/projects/%v/repository/tags/%v/release", projectID, tagName)
		assert.Equal(t, "PUT", req.Method)
		assert.Equal(t, url, req.URL.String())
		assert.Equal(t, "testtoken", req.Header["Private-Token"][0])

		body, err := ioutil.ReadAll(req.Body)

		assert.NoError(t, err)

		expectedBody := fmt.Sprintf("{\"name\":\"\",\"description\":\"%v\"}", newReleaseNotes)

		assert.Equal(t, expectedBody, string(body))

		_, err = rw.Write([]byte("ok"))

		assert.NoError(t, err)
	}))

	defer server.Close()

	service := &Gitlab{
		apiURL:    server.URL,
		client:    createClient("testtoken"),
		projectID: projectID,
		tagName:   tagName,
	}

	err := service.Publish(&release.Release{ID: projectID, Tag: tagName, Name: "v1.0.0", Message: "Description of the release", ReleaseNotes: newReleaseNotes})

	assert.NoError(t, err)
}

func TestPublishNewRelease(t *testing.T) {
	projectID := 1
	tagName := "v1.0.0"

	newReleaseNotes := "test"
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		url := fmt.Sprintf("/projects/%v/repository/tags/%v/release", projectID, tagName)
		assert.Equal(t, "POST", req.Method)
		assert.Equal(t, url, req.URL.String())
		assert.Equal(t, "testtoken", req.Header["Private-Token"][0])

		body, err := ioutil.ReadAll(req.Body)

		assert.NoError(t, err)

		expectedBody := fmt.Sprintf("{\"name\":\"\",\"description\":\"%v\"}", newReleaseNotes)

		assert.Equal(t, expectedBody, string(body))

		_, err = rw.Write([]byte("ok"))

		assert.NoError(t, err)
	}))

	defer server.Close()

	service := &Gitlab{
		apiURL:    server.URL,
		client:    createClient("testtoken"),
		projectID: projectID,
		tagName:   tagName,
	}

	err := service.Publish(&release.Release{ID: projectID, Tag: tagName, Name: "v1.0.0", Message: "", ReleaseNotes: newReleaseNotes})

	assert.NoError(t, err)
}
