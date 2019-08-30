#!/bin/bash

source scripts/.variables.sh

cloc_test(){
        cd `pwd` && \
        echo "圈复杂度检查前20:" && \
        ls -d */ | grep -v vendor | xargs gocyclo -top 20 && \
        ls -d */ | grep -v vendor | xargs gocyclo | awk '{sum+=$$1}END{printf("总圈复杂度: %s\n", sum)}'
#        echo "统计代码行数:\n" && \
#        ls -d */ | grep -v vendor | xargs cloc --by-file
}
cloc_test