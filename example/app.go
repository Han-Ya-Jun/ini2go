package example

type App struct {
	ConsoleWriterColor string `json:"console_writer_color" db:"console_writer_color"`
	ConsoleWriterOn    string `json:"console_writer_on" db:"console_writer_on"`
	FileWriterOn       string `json:"file_writer_on" db:"file_writer_on"`
	Level              string `json:"level" db:"level"`
	Path               string `json:"path" db:"path"`
	RotateLogPath      string `json:"rotate_log_path" db:"rotate_log_path"`
	RotateWfLogPath    string `json:"rotate_wf_log_path" db:"rotate_wf_log_path"`
	TimeFormat         string `json:"time_format" db:"time_format"`
	WfLogPath          string `json:"wf_log_path" db:"wf_log_path"`
}

type Server struct {
	HTTPPort     string `json:"http_port" db:"http_port"`
	ReadTimeout  string `json:"read_timeout" db:"read_timeout"`
	RunMode      string `json:"run_mode" db:"run_mode"`
	WriteTimeout string `json:"write_timeout" db:"write_timeout"`
}

type Mongo struct {
	Host     string `json:"host" db:"host"`
	Password string `json:"password" db:"password"`
	Source   string `json:"source" db:"source"`
	User     string `json:"user" db:"user"`
}

type Redis struct {
	Host        string `json:"host" db:"host"`
	IdleTimeout int    `json:"idle_timeout" db:"idle_timeout"`
	MaxActive   int    `json:"max_active" db:"max_active"`
	MaxIdle     int    `json:"max_idle" db:"max_idle"`
	Password    string `json:"password" db:"password"`
}
