package maininit

// SimpleConfig is a simple implementation of Config.
type SimpleConfig struct {
	GraphicsWidth  int
	GraphicsHeight int
	TPS            int
}

func (cfg SimpleConfig) GraphicsSize() (width int, height int) {
	return cfg.GraphicsWidth, cfg.GraphicsHeight
}

func (cfg SimpleConfig) TicksPerSecond() int {
	return cfg.TPS
}
