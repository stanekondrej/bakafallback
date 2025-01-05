package bakafallback

import (
	_ "embed"
	"log"
	"strings"
	"time"

	"github.com/stanekondrej/bakalari/pkg/bakalari"
)

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

	teachers := makeTeacherMap(timetable.Teachers)
	subjects := makeSubjectMap(timetable.Subjects)
	rooms := makeRoomMap(timetable.Rooms)

	result := "<table>"

	for _, day := range timetable.Days {
		atoms := makeAtomMap(day.Atoms)
		result += "<tr>"

		for _, hour := range timetable.Hours {
			result += "<td>"

			a, ok := atoms[hour.Id]
			if !ok {
				result += "</td>"
				continue
			}

			result += `<p class="room-abbrev">` + rooms[a.RoomId].Abbrev + "</p>"
			result += `<p class="subject-abbrev">` + subjects[a.SubjectId].Abbrev + "</p>"
			result += `<p class="teacher-abbrev">` + teachers[a.TeacherId].Abbrev + "</p>"

			result += "</td>"
		}

		result += "</tr>"
	}

	result += "</table>"

	return result
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
