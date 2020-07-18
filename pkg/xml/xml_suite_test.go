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
			cot := NewCotXML("Joe1", nil)
			Expect(cot.Event.Detail.UID.Droid).Should(Equal("Joe1"))
		})
		It("UpdateSelf/Marshall should work", func() {
			cot := NewCotXML("Joe1", nil)
			expLat := 41.0
			expCourse := 123.4
			pt := &Point{
				Lat: expLat,
			}
			tr := &Track{
				Course: expCourse,
			}
			up := cot.UpdateSelfEvent(pt, tr)
			Expect(up.Point.Lat).Should(Equal(expLat))
			Expect(up.Point.Lat).ShouldNot(Equal(cot.Event.Point.Lat)) // Make sure its a copy
			Expect(up.Detail.Track.Course).Should(Equal(expCourse))

			_, err := up.MarshallEvent()
			Expect(err).Should(BeNil())
		})
	})
})
