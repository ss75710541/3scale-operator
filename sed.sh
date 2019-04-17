#!/bin/bash
grep 'github.com/ss75710541/3scale-operator' -rl * |xargs sed -i '' 's#github.com/ss75710541/3scale-operator#github.com/ss75710541/3scale-operator#g'
