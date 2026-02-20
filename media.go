package whatsapp_business

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"path"
	"strings"
)

type Media struct {
	MessagingProduct string `json:"messaging_product"`
	URL              string `json:"url"`
	MimeType         string `json:"mime_type"`
	Sha256           string `json:"sha256"`
	FileSize         uint64 `json:"file_size"`
	ID               string `json:"id"`
	*ErrorResponse
}

func (s *Client) GetMedia(ctx context.Context, mediaID string) (*Media, error) {
	resp, err := s.metaRequestWithToken(ctx, nil, http.MethodGet, fmt.Sprintf("%s?phone_number_id=%s", mediaID, s.phoneNumberID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var toReturn Media
	if err = json.NewDecoder(resp.Body).Decode(&toReturn); err != nil {
		return nil, err
	}

	if toReturn.ErrorResponse != nil {
		return nil, errors.New(toReturn.ErrorResponse.Error.Message)
	}

	return &toReturn, nil
}

type DownloadMediaResponse struct {
	*http.Response
	GuessFileName string `json:"guess_file_name"`
}

func (s *Client) DownloadMedia(ctx context.Context, mediaURL string) (*DownloadMediaResponse, error) {
	resp, err := s.metaDownloadFile(ctx, mediaURL)
	if err != nil {
		return nil, err
	}

	fileName := guessFilename(resp, mediaURL, "download")

	return &DownloadMediaResponse{
		resp,
		fileName,
	}, nil
}

type UploadFromURL struct {
	URL string `json:"url"`
}

type UploadFromURLResponse struct {
	ID string `json:"id"`
	*ErrorResponse
}

func (s *Client) UploadFromURL(ctx context.Context, body UploadFromURL) (*UploadFromURLResponse, error) {
	reqIn, err := http.NewRequestWithContext(ctx, http.MethodGet, body.URL, nil)
	if err != nil {
		return nil, fmt.Errorf("build GET src: %w", err)
	}

	respIn, err := s.httpClient.Do(reqIn)
	if err != nil {
		return nil, fmt.Errorf("GET src: %w", err)
	}

	if respIn.StatusCode < 200 || respIn.StatusCode >= 300 {
		defer respIn.Body.Close()
		return nil, fmt.Errorf("GET src bad status: %s", respIn.Status)
	}

	mimeType := respIn.Header.Get("Content-Type")
	if strings.Contains(mimeType, "ogg") {
		mimeType = "audio/ogg; codecs=opus"
	}

	parsedURL, err := url.Parse(body.URL)
	filename := "media_file"
	if err == nil {
		filename = path.Base(parsedURL.Path)
	}
	if filename == "." || filename == "/" {
		filename = "upload.ogg"
	}

	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)

	go func() {
		defer respIn.Body.Close()
		defer func() {
			_ = mw.Close()
			_ = pw.Close()
		}()

		if err := mw.WriteField("messaging_product", "whatsapp"); err != nil {
			_ = pw.CloseWithError(fmt.Errorf("write field messaging_product: %w", err))
			return
		}

		h := make(textproto.MIMEHeader)
		h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "file", filename))
		h.Set("Content-Type", mimeType)

		part, err := mw.CreatePart(h)
		if err != nil {
			_ = pw.CloseWithError(fmt.Errorf("create part: %w", err))
			return
		}

		if _, err := io.Copy(part, respIn.Body); err != nil {
			_ = pw.CloseWithError(fmt.Errorf("copy body: %w", err))
			return
		}
	}()

	resp, err := s.metaMultipartRequestWithToken(ctx, pr, mw.FormDataContentType(), http.MethodPost, fmt.Sprintf("%s/media", s.phoneNumberID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var toReturn UploadFromURLResponse
	if err = json.NewDecoder(resp.Body).Decode(&toReturn); err != nil {
		return nil, err
	}

	if toReturn.ErrorResponse != nil {
		return nil, errors.New(toReturn.ErrorResponse.Error.Message)
	}

	return &toReturn, nil
}

type UploadFileResponse struct {
	ID string `json:"id"`
	*ErrorResponse
}

func (s *Client) UploadFile(ctx context.Context, reader io.Reader, filename string, mimeType string) (*UploadFileResponse, error) {
	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)

	go func() {
		defer func() {
			_ = mw.Close()
			_ = pw.Close()
		}()

		if err := mw.WriteField("messaging_product", "whatsapp"); err != nil {
			_ = pw.CloseWithError(fmt.Errorf("write field messaging_product: %w", err))
			return
		}

		h := make(textproto.MIMEHeader)
		h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "file", filename))
		h.Set("Content-Type", mimeType)

		part, err := mw.CreatePart(h)
		if err != nil {
			_ = pw.CloseWithError(fmt.Errorf("create part: %w", err))
			return
		}

		if _, err := io.Copy(part, reader); err != nil {
			_ = pw.CloseWithError(fmt.Errorf("copy body: %w", err))
			return
		}
	}()

	resp, err := s.metaMultipartRequestWithToken(ctx, pr, mw.FormDataContentType(), http.MethodPost, fmt.Sprintf("%s/media", s.phoneNumberID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var toReturn UploadFileResponse
	if err = json.NewDecoder(resp.Body).Decode(&toReturn); err != nil {
		return nil, err
	}

	if toReturn.ErrorResponse != nil {
		return nil, errors.New(toReturn.ErrorResponse.Error.Message)
	}

	return &toReturn, nil
}

func guessFilename(resp *http.Response, urlStr, fallbackBase string) string {
	if cd := resp.Header.Get("Content-Disposition"); cd != "" {
		if _, params, err := mime.ParseMediaType(cd); err == nil {
			if fn := params["filename"]; fn != "" {
				return fn
			}
			if fn := params["filename*"]; fn != "" { // RFC 5987
				return fn
			}
		}
	}
	if base := path.Base(urlStr); base != "" && !strings.HasSuffix(base, "/") {
		return base
	}

	ext := ""
	if ct := resp.Header.Get("Content-Type"); ct != "" {
		if exts, _ := mime.ExtensionsByType(ct); len(exts) > 0 {
			ext = exts[0]
		}
	}
	if ext == "" {
		ext = ".bin"
	}
	if fallbackBase == "" {
		fallbackBase = "download"
	}

	return fallbackBase + ext
}
