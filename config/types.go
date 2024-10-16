package config

// Cfg 配置信息
type Cfg struct {
	App   App   `yaml:"app"`
	MySQL MySQL `yaml:"mysql"`
}

// App 应用信息配置
type App struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
}

// MySQL mysql配置
type MySQL struct {
	DNS string `yaml:"dns"`
}
