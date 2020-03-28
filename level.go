package log

//go:generate stringer -type=Level -linecomment

// Level log.
type Level uint8

// awailable log levels.
const (
	LevelEmergency Level = iota //emergency
	LevelAlert                  //alert
	LevelCritical               //critical
	LevelError                  //error
	LevelWarning                //warning
	LevelNotice                 //notice
	LevelInfo                   //info
	LevelDebug                  //debug
)
