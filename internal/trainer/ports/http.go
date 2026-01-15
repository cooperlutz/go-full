package ports

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/oapi-codegen/runtime/types"

	"github.com/cooperlutz/go-full/internal/trainer/app"
	"github.com/cooperlutz/go-full/internal/trainer/app/command"
	"github.com/cooperlutz/go-full/internal/trainer/app/query"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(application app.Application) HttpServer {
	return HttpServer{
		app: application,
	}
}

func (h HttpServer) GetTrainerAvailableHours(w http.ResponseWriter, r *http.Request, params GetTrainerAvailableHoursParams) {
	dateModels, err := h.app.Queries.TrainerAvailableHours.Handle(r.Context(), query.AvailableHours{
		From: params.DateFrom,
		To:   params.DateTo,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	dates := dateModelsToResponse(dateModels)
	render.Respond(w, r, dates)
}

func dateModelsToResponse(models []query.Date) []Date {
	var dates []Date

	for _, d := range models {
		var hours []Hour
		for _, h := range d.Hours {
			hours = append(hours, Hour{
				Available:            h.Available,
				HasTrainingScheduled: h.HasTrainingScheduled,
				Hour:                 h.Hour,
			})
		}

		dates = append(dates, Date{
			Date: types.Date{
				Time: d.Date,
			},
			HasFreeHours: d.HasFreeHours,
			Hours:        hours,
		})
	}

	return dates
}

func (h HttpServer) MakeHourAvailable(w http.ResponseWriter, r *http.Request) {
	hourUpdate := &HourUpdate{}
	if err := render.Decode(r, hourUpdate); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	err := h.app.Commands.MakeHoursAvailable.Handle(r.Context(), command.MakeHoursAvailable{Hours: hourUpdate.Hours})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h HttpServer) MakeHourUnavailable(w http.ResponseWriter, r *http.Request) {
	hourUpdate := &HourUpdate{}
	if err := render.Decode(r, hourUpdate); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	err := h.app.Commands.MakeHoursUnavailable.Handle(r.Context(), command.MakeHoursUnavailable{Hours: hourUpdate.Hours})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
