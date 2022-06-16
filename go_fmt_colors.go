package go_fmt_colors

const (
	FmtFgBlack       string = "\033[30m"
	FmtFgRed         string = "\033[31m"
	FmtFgGreen       string = "\033[32m"
	FmtFgYellow      string = "\033[33m"
	FmtFgBlue        string = "\033[34m"
	FmtFgPurple      string = "\033[35m"
	FmtFgCyan        string = "\033[36m"
	FmtFgWhite       string = "\033[37m"
	FmtFgDarkGray    string = "\033[90m"
	FmtFgLightRed    string = "\033[91m"
	FmtFgLightGreen  string = "\033[92m"
	FmtFgLightYellow string = "\033[93m"
	FmtFgLightBlue   string = "\033[94m"
	FmtFgLightPurple string = "\033[95m"
	FmtFgLightCyan   string = "\033[96m"
	FmtFgLightWhite  string = "\033[97m"

	FmtBgBlack       string = "\033[40m"
	FmtBgRed         string = "\033[41m"
	FmtBgGreen       string = "\033[42m"
	FmtBgYellow      string = "\033[43m"
	FmtBgBlue        string = "\033[44m"
	FmtBgPurple      string = "\033[45m"
	FmtBgCyan        string = "\033[46m"
	FmtBgWhite       string = "\033[47m"
	FmtBgDarkGray    string = "\033[100m"
	FmtBgLightRed    string = "\033[101m"
	FmtBgLightGreen  string = "\033[102m"
	FmtBgLightYellow string = "\033[103m"
	FmtBgLightBlue   string = "\033[104m"
	FmtBgLightPurple string = "\033[105m"
	FmtBgLightCyan   string = "\033[106m"
	FmtBgLightWhite  string = "\033[107m"

	FmtFgBgBlackWhite   string = FmtFgBlack + FmtBgWhite
	FmtFgBgWhiteBlack   string = FmtFgWhite + FmtBgBlack
	FmtFgBgWhiteBlue    string = FmtFgWhite + FmtBgBlue
	FmtFgBgWhiteRed     string = FmtFgWhite + FmtBgRed
	FmtFgBgWhitePurple  string = FmtFgWhite + FmtFgPurple
	FmtFgBgWhiteLBlue   string = FmtFgWhite + FmtBgLightBlue
	FmtFgBgWhiteLRed    string = FmtFgWhite + FmtBgLightRed
	FmtFgBgWhiteLPurple string = FmtFgWhite + FmtFgLightPurple
	FmtFgBgBlackLBlue   string = FmtFgBlack + FmtBgLightBlue
	FmtFgBgBlackLRed    string = FmtFgBlack + FmtBgLightRed
	FmtFgBgBlackLPurple string = FmtFgBlack + FmtFgLightPurple

	FmtReset string = "\033[0m"
)
