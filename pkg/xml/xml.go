package xml

import (
	"encoding/xml"
	"time"

	"github.com/google/uuid"
)

const (
	layout = "2006-01-02T15:04:05.123Z"
)

// Event contains the whole of the CoT message
type Event struct {
	XMLName   xml.Name      `xml:"event"`
	Version   string        `xml:"version,attr"`
	Uid       string        `xml:"uid,attr"`
	Type      string        `xml:"type,attr"`
	Time      string        `xml:"time,attr"`
	Start     string        `xml:"start,attr"`
	Stale     string        `xml:"stale,attr"`
	How       string        `xml:"how,attr"`
	Point     *Point        `xml:"point"`
	Detail    *Detail       `xml:"detail"`
	DfltStale time.Duration `xml:"-"`
}

// Returns an event that keeps state of minimal info, normally dont use this event directly
func NewDefaultEvent(call string, tv *Takv) *Event {
	var takv *Takv
	if tv != nil {
		takv = tv
	} else {
		takv = &Takv{
			OS:       "1",
			Version:  "1",
			Device:   "go-tak",
			Platform: "Go",
		}
	}

	evt := &Event{
		DfltStale: 5 * time.Minute,
		Version:   "2.0",
		Uid:       uuid.New().String(),
		Type:      "a-f-G-U-C",
		How:       "m-g",
		Point:     &Point{},
		Detail: &Detail{
			Takv: takv,
			Contact: &Contact{
				Endpoint: "tcpsrcreply:4242:srctcp",
				Callsign: call,
			},
			UID:    &UID{Droid: call},
			Loc:    &Loc{AltSrc: "GPS", Geo: "GPS"},
			Group:  &Group{Role: "Team member", Name: "Cyan"},
			Status: &Status{Battery: 100},
			Track:  &Track{},
		},
	}
	return evt
}

// UpdateSelf returns a copy of the base event with a few fields updated, primary use case is to update your position
func (e *Event) UpdateSelf(pt *Point, tr *Track) *Event {
	eCpy := *e
	eCpy.Point = pt
	if tr != nil {
		eCpy.Detail.Track = tr
	}
	now := getTime(0)
	stale := getTime(eCpy.DfltStale)
	eCpy.Time = now
	eCpy.Start = now
	eCpy.Stale = stale
	return &eCpy
}

// Point in CoT
type Point struct {
	XMLName xml.Name `xml:"point"`
	Lat     float64  `xml:"lat,attr"`
	Long    float64  `xml:"lon,attr"`
	Hae     float64  `xml:"hae,attr"`
	CE      float64  `xml:"ce,attr"`
	LE      float64  `xml:"le,attr"`
}

// Detail
type Detail struct {
	XMLName xml.Name `xml:"detail"`
	Takv    *Takv    `xml:"takv"`
	Contact *Contact `xml:"contact"`
	UID     *UID     `xml:"uid"`
	Loc     *Loc     `xml:"precisionlocation"`
	Group   *Group   `xml:"__group"`
	Status  *Status  `xml:"status"`
	Track   *Track   `xml:"track"`
}

// Takv
type Takv struct {
	XMLName  xml.Name `xml:"takv"`
	OS       string   `xml:"os,attr"`
	Version  string   `xml:"version,attr"`
	Device   string   `xml:"device,attr"`
	Platform string   `xml:"platform,attr"`
}

// Contact
type Contact struct {
	XMLName  xml.Name `xml:"contact"`
	Endpoint string   `xml:"endpoint,attr"`
	Callsign string   `xml:"callsign,attr"`
}

// UID
type UID struct {
	XMLName xml.Name `xml:"uid"`
	Droid   string   `xml:"Droid,attr"`
}

// Loc is the precisionlocation
type Loc struct {
	XMLName xml.Name `xml:"precisionlocation"`
	AltSrc  string   `xml:"altsrc,attr"`
	Geo     string   `xml:"geopointsrc,attr"`
}

// Group
type Group struct {
	XMLName xml.Name `xml:"__group"`
	Role    string   `xml:"role,attr"`
	Name    string   `xml:"name,attr"`
}

// Status
type Status struct {
	XMLName xml.Name `xml:"status"`
	Battery int      `xml:"battery,attr"`
}

// Track
type Track struct {
	XMLName xml.Name `xml:"track"`
	Course  float64  `xml:"course,attr"`
	Speed   float64  `xml:"speed,attr"`
}

// Gets the current zulu time in mil format plus any added duration for future time
func getTime(d time.Duration) string {
	t := time.Now()
	if d != 0 {
		t = t.Add(d)
	}
	return t.UTC().Format(layout)
}

// MarshallEvent into XML bytes
func MarshallEvent(evt *Event) ([]byte, error) {
	str, err := xml.Marshal(evt)
	if err != nil {
		return nil, err
	}
	return str, nil
}
