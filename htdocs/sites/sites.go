package sites

import (
	"math/rand"
)

type Sites struct {
	Address []string
}

func (s *Sites) Init() {
	s.Address = append(s.Address, "https://findtheinvisiblecow.com/")
	s.Address = append(s.Address, "https://www.mapcrunch.com/")
	s.Address = append(s.Address, "https://theuselessweb.com/")
	s.Address = append(s.Address, "http://hackertyper.com/")
	s.Address = append(s.Address, "http://papertoilet.com/")
	s.Address = append(s.Address, "http://newsoffuture.com/")
	s.Address = append(s.Address, "https://pointerpointer.com/")
	s.Address = append(s.Address, "http://beesbeesbees.com/")
	s.Address = append(s.Address, "http://www.shadyurl.com/")
	s.Address = append(s.Address, "https://archive.org/web/")
}

func (s *Sites) GetSite() string {
	random := rand.Intn(len(s.Address))
	return s.Address[random]
}
