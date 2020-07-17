package xml

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestXml(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Xml Suite")
}

var _ = Describe("XML CoT", func() {
	Context("Event", func() {
		It("NewDefaultEvent work", func() {
			evt := NewDefaultEvent("Joe1", nil)
			Expect(evt.Detail.UID.Droid).Should(Equal("Joe1"))
		})
		It("UpdateSelf should work", func() {
			evt := NewDefaultEvent("Joe1", nil)
			expLat := 41.0
			expCourse := 123.4
			pt := &Point{
				Lat: expLat,
			}
			tr := &Track{
				Course: expCourse,
			}
			up := evt.UpdateSelf(pt, tr)
			Expect(up.Point.Lat).Should(Equal(expLat))
			Expect(up.Detail.Track.Course).Should(Equal(expCourse))
		})
	})
	Context("MarshallEvent", func() {
		It("should work", func() {
			evt := NewDefaultEvent("Joe1", nil)
			_, err := MarshallEvent(evt)
			Expect(err).Should(BeNil())
		})
	})
})
