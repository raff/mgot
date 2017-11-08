# mgot
Get last SMS message from MacOS Messages database

The main purpose of this command is to get the text of the last SMS sent and copy it to the clipboard 
(for example, to get the code sent for two-factor authentication and past it in the login form).

The script is pre-configured for Okta authentication, but should be adaptable to other forms of 2FA by passing 
the appripriate parameters.

## USAGE

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
