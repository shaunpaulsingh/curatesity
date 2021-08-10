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
	s.Address = append(s.Address, "https://radio.garden/")
	s.Address = append(s.Address, "https://outline.com/")
	s.Address = append(s.Address, "https://www.window-swap.com/")
	s.Address = append(s.Address, "https://boilthefrog.playlistmachinery.com/")
	s.Address = append(s.Address, "https://whatthefuckshouldimakefordinner.com/")
	s.Address = append(s.Address, "https://mynoise.net/")
	s.Address = append(s.Address, "https://www.howmanypeopleareinspacerightnow.com/")
	s.Address = append(s.Address, "https://chir.ag/projects/tip-of-my-tongue/")
	s.Address = append(s.Address, "https://www.futureme.org/")
	s.Address = append(s.Address, "https://erowid.org/")
	s.Address = append(s.Address, "https://thenounproject.com/")
	s.Address = append(s.Address, "https://earth.nullschool.net/")
	s.Address = append(s.Address, "https://www.naturalreaders.com/")
	s.Address = append(s.Address, "https://www.futuretimeline.net/")
	s.Address = append(s.Address, "https://imslp.org/wiki/Main_Page")
	s.Address = append(s.Address, "https://runpee.com/")
	s.Address = append(s.Address, "https://terriblerealestateagentphotos.com/")
	s.Address = append(s.Address, "https://asoftmurmur.com/")
	s.Address = append(s.Address, "https://www.removeddit.com/")
	s.Address = append(s.Address, "https://www.projectrho.com/public_html/rocket/")
}

func (s *Sites) GetSite() string {
	random := rand.Intn(len(s.Address))
	return s.Address[random]
}
