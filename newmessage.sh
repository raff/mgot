#!/bin/sh
bin=`dirname $0`
code=`$bin/mgot`
echo newmessage at `date "+%Y-%m-%d %H:%M:%S"` with code ${code}
echo ${code} | pbcopy
