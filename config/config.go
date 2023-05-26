package config

type Server struct {
	Zap   Zap   `mapstructure:"zap" json:"zap" yaml:"zap"`
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Local Local `mapstructure:"local" json:"local" yaml:"local"`
	Excel Excel `mapstructure:"excel" json:"excel" yaml:"excel"`
	Timer Timer `mapstructure:"timer" json:"timer" yaml:"timer"`
	Bot   Bot   `mapstructure:"bot" json:"bot" yaml:"bot"`
}
