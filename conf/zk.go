package conf

type Zookeeper struct {
	RackRoot   string
	VolumeRoot string
	GroupRoot  string
	PitchRoot  string
	Addrs      []string
}

// New new live zookeeper registry
func New(config *Zookeeper) {


}
