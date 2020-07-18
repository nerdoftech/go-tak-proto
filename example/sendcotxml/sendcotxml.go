package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"

	cot "github.com/nerdoftech/go-tak-proto/pkg/xml"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	udpMulticast = "239.2.3.1:6969"
)

var (
	flgName = flag.String("name", "joe", "Basename of the targets")
	flgLat  = flag.Float64("lat", 0, "lattitude")
	flgLon  = flag.Float64("lon", 0, "longitude")
	flgHae  = flag.Float64("hae", 0, "hae altitude")
	flgNum  = flag.Int("num", 20, "number of targets")
)

func main() {
	flag.Parse()
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	addr, err := net.ResolveUDPAddr("udp", udpMulticast)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to resolve UDP address")
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to make UDP conn")
	}

	generateCots(*flgName, *flgLat, *flgLon, *flgHae, *flgNum, conn)
	conn.Close()
	log.Info().Msg("all done")
}

// Create random target around given coord
func generateCots(basename string, lat, lon, hae float64, num int, conn *net.UDPConn) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num; i++ {
		name := fmt.Sprintf("%s-%d", basename, i)
		c := cot.NewCotXML(name, nil)
		pt := &cot.Point{
			Lat:  lat + float64(rand.Int()%50-25)/100,
			Long: lon + float64(rand.Int()%50-25)/100,
			Hae:  hae,
			CE:   3.2,
			LE:   9999.0,
		}
		tr := &cot.Track{
			Course: float64(rand.Int() % 360),
			Speed:  float64(rand.Int() % 150),
		}
		b, err := c.UpdateSelfEvent(pt, tr).MarshallEvent()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to marshall CoT")
		}
		log.Debug().Bytes("cot", b).Msg("CoT contents")
		sz, err := conn.Write(b)
		if err != nil {
			log.Fatal().Err(err).Msg("could not send data")
		}
		log.Debug().Int("size", sz).Msg("bytes written to socket")
	}

}
