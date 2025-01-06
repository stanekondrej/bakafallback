package bakafallback

import (
	_ "embed"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/stanekondrej/bakalari/pkg/bakalari"
)

type Renderer struct {
	hours    []bakalari.Hour
	teachers map[string]bakalari.Teacher
	subjects map[string]bakalari.Subject
	rooms    map[string]bakalari.Room
	days     []bakalari.Day
}

func newRenderer(timetable *bakalari.Timetable) *Renderer {
	if timetable == nil {
		log.Fatal("Rozvrh byl nil (funkce newRenderer)")
	}

	return &Renderer{
		hours:    timetable.Hours,
		teachers: makeTeacherMap(timetable.Teachers),
		subjects: makeSubjectMap(timetable.Subjects),
		rooms:    makeRoomMap(timetable.Rooms),
		days:     timetable.Days,
	}
}

func (r *Renderer) renderTimetable() (html string) {
	html = "<table>"
	html += r.renderHourRow()

	for i, day := range r.days {
		html += `<tr class="day">`
		html += `<td class="day-label">` + dayOfWeekToString(uint(i)) + "</td>"

		atoms := makeAtomMap(day.Atoms)

		for _, hour := range r.hours {
			html += `<td class="hour">`

			a, ok := atoms[hour.Id]

			if ok {
				html += `<p class="room-abbrev">` + r.rooms[a.RoomId].Abbrev + "</p>"
				html += `<p class="subject-abbrev">` + r.subjects[a.SubjectId].Abbrev + "</p>"
				html += `<p class="teacher-abbrev">` + r.teachers[a.TeacherId].Abbrev + "</p>"
			}

			html += "</td>"
		}

		html += "</tr>"
	}

	return
}

func dayOfWeekToString(i uint) string {
	days := []string{"Pondělí", "Úterý", "Středa", "Čtvrtek", "Pátek"}

	return days[i]
}

func (r *Renderer) renderHourRow() (html string) {
	html = `<tr class="hours-row">`
	html += "<td></td>" // prázdná buňka

	for i, h := range r.hours {
		html += renderHour(&h, uint(i))
	}

	html += "</tr>"
	return
}

func renderHour(h *bakalari.Hour, index uint) (html string) {
	if h == nil {
		log.Fatal("Hodina byla nil (funkce renderHour)")
	}

	html = `<td class="hour">`
	html += fmt.Sprintf(`<p class="hour-id">%d</p>`, index)
	html += fmt.Sprintf(`<p class="hour-times">%s - %s</p>`, h.BeginTime, h.EndTime)
	html += "</td>"

	return
}

func makeAtomMap(atoms []bakalari.Atom) map[uint]bakalari.Atom {
	m := make(map[uint]bakalari.Atom)

	for _, atom := range atoms {
		m[atom.HourId] = atom
	}

	return m
}

func makeTeacherMap(teachers []bakalari.Teacher) map[string]bakalari.Teacher {
	m := make(map[string]bakalari.Teacher)

	for _, teacher := range teachers {
		m[teacher.Id] = teacher
	}

	return m
}

func makeSubjectMap(subjects []bakalari.Subject) map[string]bakalari.Subject {
	m := make(map[string]bakalari.Subject)

	for _, subject := range subjects {
		m[subject.Id] = subject
	}

	return m
}

func makeRoomMap(rooms []bakalari.Room) map[string]bakalari.Room {
	m := make(map[string]bakalari.Room)

	for _, room := range rooms {
		m[room.Id] = room
	}

	return m
}

func renderTimetable(timetable *bakalari.Timetable) string {
	if timetable == nil {
		log.Fatal("Rozvrh byl nil")
	}

	renderer := newRenderer(timetable)
	html := renderer.renderTimetable()

	return html
}

func renderStatus() string {
	result := ""

	d := time.Now()
	result += "Datum aktualizace: " + d.Format(time.DateTime)

	return result
}

//go:embed template.html
var template string

func RenderPage(timetable *bakalari.Timetable) string {
	if timetable == nil {
		log.Fatal("Rozvrh byl nil")
	}

	statusHtml := renderStatus()
	timetableHtml := renderTimetable(timetable)

	halves := strings.SplitN(template, "<!--STATUS-->", 2)
	html := halves[0] + statusHtml + halves[1]

	halves = strings.SplitN(html, "<!--MARKER-->", 2)

	return halves[0] + timetableHtml + halves[1]
}
