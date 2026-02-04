package messages

// Action repsonce
const (
	Start   = "Timer started."
	Pause   = "Timer paused."
	Resume  = "Timer resumed."
	Stop    = "Timer stopped."
	AddTime = "Adding time to database"
)

// Warnings
const (
	AlreadyStarted  = "Timer already started."
	NoTimer         = "No active timer."
	AlreadyPaused   = "Time already paused."
	NotPaused       = "Timer not paused."
	PausedCantStart = "Timer paused, enter 'stop' to clear current timer or 'resume' to continue current"
	Invalid         = "Invalid command, enter 'help' to see commands."
	StillRunning    = "Timer still running"
	InvalidDate     = "Invalid date format, please try again."
	NoSearch        = "Unable to find entry with provided date."
)

// Prompts
const (
	WantToStop  = "Do you want to stop? (y/n)"
	WantToSave  = "Do you want to save time? (y/n)"
	AnotherDate = "Try another date? (y/n)"
)

// Lists what each command does.
const CommandHelp = `
start - Starts timer
stop - Stops timer & prints final time
pause - Pauses timer
resume - Resumes timer after pausing
reveal - shows current time
search - search for specific time
times - prints all stored times
export - export stored times to CSV file
debug - toggle debug mode
`
