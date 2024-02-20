package domain

type ExtConf struct {
	App      *App      `json:"app"`
	Database *Database `json:"database"`
	Kafka    *Kafka    `json:"kafka"`
}
type App struct {
	Port int `json:"port"`
}
type Database struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Port     int    `json:"port"`
}

type Kafka struct {
	Host  string `json:"host"`
	Port  int    `json:"port"`
	Topic string `json:"topic"`
}
