package myLibary

var (
	Banner1 = `
{{.AnsiColor.BrightBlue}}████████████████████████████████████████████████████████████████████████████████
{{.AnsiColor.BrightBlue}}████████████████████████████████████████████████████████████████████████████████
{{.AnsiColor.Red}}██████████████████{{.AnsiColor.BrightBlue}}████████████████{{.AnsiColor.Black}}██████████████████████████████{{.AnsiColor.BrightBlue}}████████████████
{{.AnsiColor.Red}}████████████████████████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██████████████████████████████{{.AnsiColor.Black}}██{{.AnsiColor.BrightBlue}}██████████████
{{.AnsiColor.BrightRed}}████{{.AnsiColor.Red}}██████████████████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██████{{.AnsiColor.Magenta}}██████████████████████{{.AnsiColor.White}}██████{{.AnsiColor.Black}}██{{.AnsiColor.BrightBlue}}████████████
{{.AnsiColor.BrightRed}}██████████████████████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}████{{.AnsiColor.Magenta}}████████████████{{.AnsiColor.Black}}████{{.AnsiColor.Magenta}}██████{{.AnsiColor.White}}████{{.AnsiColor.Black}}██{{.AnsiColor.BrightBlue}}██{{.AnsiColor.Black}}████{{.AnsiColor.BrightBlue}}██████
{{.AnsiColor.BrightRed}}██████████████████████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██{{.AnsiColor.Magenta}}████████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}████{{.AnsiColor.Black}}██{{.AnsiColor.Magenta}}██████{{.AnsiColor.White}}██{{.AnsiColor.Black}}████{{.AnsiColor.White}}████{{.AnsiColor.Black}}██{{.AnsiColor.BrightBlue}}████
{{.AnsiColor.BrightYellow}}██████████████████{{.AnsiColor.BrightRed}}████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██{{.AnsiColor.Magenta}}████████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██████{{.AnsiColor.Magenta}}██████{{.AnsiColor.White}}██{{.AnsiColor.Black}}██{{.AnsiColor.White}}██████{{.AnsiColor.Black}}██{{.AnsiColor.BrightBlue}}████
{{.AnsiColor.BrightYellow}}██████████████████████{{.AnsiColor.Black}}██{{.AnsiColor.BrightYellow}}██████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██{{.AnsiColor.Magenta}}████████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██████{{.AnsiColor.Black}}████████{{.AnsiColor.White}}████████{{.AnsiColor.Black}}██{{.AnsiColor.BrightBlue}}████
{{.AnsiColor.BrightYellow}}████████████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██{{.AnsiColor.Black}}██{{.AnsiColor.BrightYellow}}████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██{{.AnsiColor.Magenta}}████████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██████████████████████{{.AnsiColor.Black}}██{{.AnsiColor.BrightBlue}}████
{{.AnsiColor.BrightGreen}}██████████████████{{.AnsiColor.BrightYellow}}██{{.AnsiColor.Black}}██{{.AnsiColor.White}}██{{.AnsiColor.Black}}████████{{.AnsiColor.White}}██{{.AnsiColor.Magenta}}██████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██████████████████████████{{.AnsiColor.Black}}██{{.AnsiColor.BrightBlue}}██
{{.AnsiColor.BrightGreen}}██████████████████████{{.AnsiColor.White}}████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██{{.AnsiColor.Magenta}}██████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██████{{.AnsiColor.BrightYellow}}██{{.AnsiColor.White}}██████████{{.AnsiColor.BrightYellow}}██{{.AnsiColor.Black}}██{{.AnsiColor.White}}████{{.AnsiColor.Black}}██{{.AnsiColor.BrightBlue}}██
{{.AnsiColor.BrightGreen}}██████████████████████{{.AnsiColor.Black}}████{{.AnsiColor.White}}████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██{{.AnsiColor.Magenta}}██████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██{{.AnsiColor.Black}}████{{.AnsiColor.White}}████{{.AnsiColor.Black}}██{{.AnsiColor.BrightBlue}}██
{{.AnsiColor.Blue}}██████████████████{{.AnsiColor.BrightGreen}}████████{{.AnsiColor.Black}}██████{{.AnsiColor.White}}██{{.AnsiColor.Magenta}}██████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██{{.AnsiColor.Magenta}}████{{.AnsiColor.White}}████████████████{{.AnsiColor.Magenta}}████{{.AnsiColor.Black}}██{{.AnsiColor.BrightBlue}}██
{{.AnsiColor.Blue}}██████████████████████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}████{{.AnsiColor.Magenta}}██████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██████{{.AnsiColor.Black}}████████████{{.AnsiColor.White}}████{{.AnsiColor.Black}}██{{.AnsiColor.BrightBlue}}████
{{.AnsiColor.BrightBlue}}██████████████████{{.AnsiColor.Blue}}████{{.AnsiColor.Blue}}██████{{.AnsiColor.Black}}████{{.AnsiColor.White}}██████{{.AnsiColor.Magenta}}██████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██████████████████{{.AnsiColor.Black}}██{{.AnsiColor.BrightBlue}}██████
{{.AnsiColor.BrightBlue}}██████████████████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██{{.AnsiColor.Black}}████{{.AnsiColor.White}}████████████████████{{.AnsiColor.Black}}██████████████████{{.AnsiColor.BrightBlue}}████████
{{.AnsiColor.BrightBlue}}████████████████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}██████{{.AnsiColor.Black}}████████████████████████████████{{.AnsiColor.White}}██{{.AnsiColor.Black}}██{{.AnsiColor.BrightBlue}}████████████
{{.AnsiColor.BrightBlue}}████████████████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}████{{.AnsiColor.Black}}██{{.AnsiColor.BrightBlue}}██{{.AnsiColor.Black}}██{{.AnsiColor.White}}████{{.AnsiColor.BrightBlue}}████████████{{.AnsiColor.Black}}██{{.AnsiColor.White}}████{{.AnsiColor.Black}}████{{.AnsiColor.White}}████{{.AnsiColor.Black}}██{{.AnsiColor.BrightBlue}}████████████
{{.AnsiColor.BrightBlue}}████████████████████████{{.AnsiColor.Black}}██████{{.AnsiColor.BrightBlue}}████{{.AnsiColor.Black}}██████{{.AnsiColor.BrightBlue}}████████████{{.AnsiColor.Black}}██████{{.AnsiColor.BrightBlue}}████{{.AnsiColor.Black}}██████{{.AnsiColor.BrightBlue}}████████████
████████████████████████████████████████████████████████████████████████████████
{{ .AnsiColor.Default }}
GoVersion: {{ .GoVersion }}
GOOS: {{ .GOOS }}
GOARCH: {{ .GOARCH }}
NumCPU: {{ .NumCPU }}
Compiler: {{ .Compiler }}
Now: {{ .Now "Monday, 2 Jan 2006" }}
`
	Banner2 = `
 {{ .AnsiColor.BrightRed }}██    ██ {{ .AnsiColor.BrightGreen }}███    ██ {{ .AnsiColor.BrightYellow }}██ {{ .AnsiColor.BrightMagenta }}████████ {{ .AnsiColor.BrightCyan }}██████  {{ .AnsiColor.BrightWhite }}██    ██ {{ .AnsiColor.BrightBlack }}███████ {{ .AnsiColor.Blue }}████████ 
 {{ .AnsiColor.BrightRed }}██    ██ {{ .AnsiColor.BrightGreen }}████   ██ {{ .AnsiColor.BrightYellow }}██    {{ .AnsiColor.BrightMagenta }}██    {{ .AnsiColor.BrightCyan }}██   ██ {{ .AnsiColor.BrightWhite }}██    ██ {{ .AnsiColor.BrightBlack }}██         {{ .AnsiColor.Blue }}██    
 {{ .AnsiColor.BrightRed }}██    ██ {{ .AnsiColor.BrightGreen }}██ ██  ██ {{ .AnsiColor.BrightYellow }}██    {{ .AnsiColor.BrightMagenta }}██    {{ .AnsiColor.BrightCyan }}██████  {{ .AnsiColor.BrightWhite }}██    ██ {{ .AnsiColor.BrightBlack }}███████    {{ .AnsiColor.Blue }}██    
 {{ .AnsiColor.BrightRed }}██    ██ {{ .AnsiColor.BrightGreen }}██  ██ ██ {{ .AnsiColor.BrightYellow }}██    {{ .AnsiColor.BrightMagenta }}██    {{ .AnsiColor.BrightCyan }}██   ██ {{ .AnsiColor.BrightWhite }}██    ██      {{ .AnsiColor.BrightBlack }}██    {{ .AnsiColor.Blue }}██    
  {{ .AnsiColor.BrightRed }}██████  {{ .AnsiColor.BrightGreen }}██   ████ {{ .AnsiColor.BrightYellow }}██    {{ .AnsiColor.BrightMagenta }}██    {{ .AnsiColor.BrightCyan }}██   ██  {{ .AnsiColor.BrightWhite }}██████  {{ .AnsiColor.BrightBlack }}███████    {{ .AnsiColor.Blue }}██    
{{ .AnsiColor.Default }}
GoVersion: {{ .GoVersion }}
GOOS: {{ .GOOS }}
GOARCH: {{ .GOARCH }}
NumCPU: {{ .NumCPU }}
Compiler: {{ .Compiler }}
Now: {{ .Now "Monday, 2 Jan 2006" }}
`
)

func Banner(str string) string {
	switch str {
	case "Banner1":
		return Banner1
	case "Banner2":
		return Banner2
	default:
		return ""
	}
}
