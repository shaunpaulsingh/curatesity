package sites

import "math/rand"

var sitesHeld = Sites{}

type Sites struct {
	Address []string
}

func Init() {
	sitesHeld.Address = append(sitesHeld.Address, "https://findtheinvisiblecow.com/")
	sitesHeld.Address = append(sitesHeld.Address, "https://www.mapcrunch.com/")
	sitesHeld.Address = append(sitesHeld.Address, "https://theuselessweb.com/")
	sitesHeld.Address = append(sitesHeld.Address, "http://hackertyper.com/")
	sitesHeld.Address = append(sitesHeld.Address, "http://papertoilet.com/")
	sitesHeld.Address = append(sitesHeld.Address, "http://newsoffuture.com/")
	sitesHeld.Address = append(sitesHeld.Address, "https://pointerpointer.com/")
	sitesHeld.Address = append(sitesHeld.Address, "http://beesbeesbees.com/")
	sitesHeld.Address = append(sitesHeld.Address, "http://www.shadyurl.com/")
	sitesHeld.Address = append(sitesHeld.Address, "https://archive.org/web/")
}

func GetSite() string {
	random := rand.Intn(len(sitesHeld.Address))
	return sitesHeld.Address[random]
}
