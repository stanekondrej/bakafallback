package bakafallback

import (
	_ "embed"
	"log"
	"net/http"
	"time"

	"github.com/stanekondrej/bakalari/pkg/bakalari"
)

type requestHandler struct {
	timetable  *bakalari.Timetable
	cachedHtml string
	api        *bakalari.Api

	accessToken  string
	refreshToken string
}

var updateInterval = time.Hour

func newRequestHandler(accessToken string, refreshToken string, api *bakalari.Api) *requestHandler {
	if api == nil {
		log.Fatal("API nesmí být nil")
	}

	rh := &requestHandler{
		api:          api,
		accessToken:  accessToken,
		refreshToken: refreshToken,
	}

	go func() {
		d := time.Now()

		for {
			timetable, err := api.FetchTimetable(accessToken, &d)
			if err != nil {
				log.Println(err)
				time.Sleep(updateInterval)

				continue
			}

			html := RenderPage(timetable)
			rh.timetable = timetable
			rh.cachedHtml = html

			time.Sleep(updateInterval)
		}
	}()

	return rh
}

func (rh *requestHandler) handleRequest(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(rh.cachedHtml))
	if err != nil {
		log.Println(err)
	}
}

func StartServer(accessToken string, refreshToken string, api *bakalari.Api) {
	if api == nil {
		log.Fatal("API nesmí být nil")
	}

	log.Println("Spouštím server bakafallback")

	rh := newRequestHandler(accessToken, refreshToken, api)

	http.HandleFunc("/", rh.handleRequest)
	log.Fatal(http.ListenAndServe("0.0.0.0:9999", nil))
}
