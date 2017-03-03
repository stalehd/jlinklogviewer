# JLink Log Viewer

The JLinkRTTClient command line tool is *really* annoying when debugging
embedded devices since it doesn't reconnect to the JLinkExe process when
it is terminated. Reconnecting the JLinkExe, then restarting the JLinkRTTClient
process gets annoying after a while.

I fixed it.

# Building 

Build with `go build`. There's no parameters on the command line.