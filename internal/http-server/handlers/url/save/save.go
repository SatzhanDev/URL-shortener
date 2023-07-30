package save

import (
	"net/http"
	resp "url-shortener/internal/lib/api/response"

	"golang.org/x/exp/slog"
)

type Request struct {
	URL   string `json:"url" variable:"required,url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	resp.Response
	Alias string `json:"alias,omitempty"`
}

type URLSaver interface {
	SaveURL(urlToSave string, alias string) (int64, error)
}

func New(log *slog.Logger, urlSaver URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
