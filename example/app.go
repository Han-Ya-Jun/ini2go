package example

type App struct {
	ConsoleWriterColor string
	ConsoleWriterOn    string
	FileWriterOn       string
	Level              string
	Path               string
	RotateLogPath      string
	RotateWfLogPath    string
	TimeFormat         string
	WfLogPath          string
}

type Server struct {
	HTTPPort     string
	ReadTimeout  string
	RunMode      string
	WriteTimeout string
}

type Mongo struct {
	Host     string
	Password string
	Source   string
	User     string
}

type Redis struct {
	Host        string
	IdleTimeout int
	MaxActive   int
	MaxIdle     int
	Password    string
}
