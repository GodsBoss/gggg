package maininit

// SimpleConfig is a simple implementation of Config.
type SimpleConfig struct {
	TPS int
}

func (cfg SimpleConfig) TicksPerSecond() int {
	return cfg.TPS
}
