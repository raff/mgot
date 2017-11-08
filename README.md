# mgot
Get last SMS message from MacOS Messages database

The main purpose of this command is to get the text of the last SMS sent and copy it to the clipboard 
(for example, to get the code sent for two-factor authentication and past it in the login form).

The script is pre-configured for Okta authentication, but should be adaptable to other forms of 2FA by passing 
the appripriate parameters.

## Usage

      mgot [-db path] [-filter] [-pattern pattern] [-service service]

      where:

      -db string
            path to user Messages database (default "~/Library/Messages/chat.db")
      -filter
            filter out pattern (default true)
      -pattern string
            message text pattern (default "Your verification code is %.")
      -service string
            message service to query (default "SMS")

- The standard location for the Messages database is ~/Library/Messages/chat.db
- The pattern should be compatible with a SQL `LIKE` statement (i.e. `message LIKE "pattern"`)
- The Messages services for SMS is "SMS"
- filter=false will return the full message, filter=true will return the part matching '%' in pattern.

## Install

You can automate the full process by installing the associated scripts as following:

    cp newmessage.sh /usr/local/bin/
    chmod +x /usr/local/bin/newmessage.sh

    cp NewMessage.plist ~/Library/LaunchAgents/
    launchctl load ~/Library/LaunchAgents/NewMessage.plist

After that, every time you get a message matching the pattern, the `code` should be copied to the clipboard.
You can verify if the process works by checking the log file /tmp/newmessage.log (usage may vary - this works for me
on one of my MacBooks, running High Sierra, but not on the older one, running MacOS Sierra). If the automated process 
doesn't work you can manually run `newmessage.sh` to copy the latest code to the clipboard.
