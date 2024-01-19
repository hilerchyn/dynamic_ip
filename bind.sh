#!/bin/bash

set -e

accessKeyId="LTAI5tDJApv8zLPf99BAFam9"
accessKeySecret="QSpHdry3lfeixe7sIfxdS1natlHzHG"
recordId="866374301683228672"
recordName="live"
domainName="7shu.co"
cmd="/home/tao/Desktop/tools/dynamic_ip/dynamic_ip"
logPath="/var/log/dynamic_ip"
logFile="dynamic_ip_$(date +"%Y%m%d").log"

export ALIBABA_CLOUD_ACCESS_KEY_ID=$accessKeyId
export ALIBABA_CLOUD_ACCESS_KEY_SECRET=$accessKeySecret

fnCheck() {
    if [ ! -f "$cmd" ]; then
        echo "An error occurred: Command file [$cmd] not found." >&2
        exit 1
    fi

    if [ ! -d "$logPath" ]; then
        echo "An error occurred: Log path [$logPath] not found." >&2
        exit 1
    fi
}

fnLogger() {
    echo "$(date +'%Y-%m-%d %H:%M:%S'): $1" >> "$logPath/$logFile"
}

fnCheck

# query old IP
oldIp=$("$cmd" queryRecords $domainName -r $recordName)

# query dynamic IP
dynamicIp=$("$cmd" queryIp)

# ip changed
if [ "$oldIp" != "$dynamicIp" ] ; then
    result=$("$cmd" updateDomainRecord -v $dynamicIp -i $recordId -r $recordName)
    if requestId=$(echo "$result" | jq -r '.RequestId'); then
        fnLogger "IP changed from [$oldIp] to [$dynamicIp], RequestId: $requestId"
    else
        fnLogger "IP chang from [$oldIp] to [$dynamicIp] failed, response: $result"
    fi
fi
