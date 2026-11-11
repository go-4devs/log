/*
Package log writes structured logs at the eight RFC 5424 levels.

Attributes are slog.Attr. The entry passed to a Logger is valid only until
the call returns. Keep a copy with entry.Entry.Clone.
*/
package log
