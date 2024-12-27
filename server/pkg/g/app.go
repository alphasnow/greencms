// @author AlphaSnow

package g

var globalContainer Container

func init() {
	globalContainer = NewContainer()
}

func App() Container {
	return globalContainer
}
