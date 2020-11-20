#!/bin/bash

host="node1"
user="dev"
pass="123"

source="."
destination="/app/bin/http-server"

echo "ci/cd listening..."
fswatch -0 -i ".go$" -e ".*" . | \
xargs   -0 -n 1 -I {} sshpass -p $pass \
rsync   -aruz --include="*.go" --exclude="*" --delete --progress $source $user@$host:$destination