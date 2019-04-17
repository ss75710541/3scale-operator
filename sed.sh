#!/bin/bash
grep 'github.com/3scale/3scale-operator' -rl *|grep -v sed.sh |xargs sed -i '' 's#github.com/3scale/3scale-operator#github.com/ss75710541/3scale-operator#g'
