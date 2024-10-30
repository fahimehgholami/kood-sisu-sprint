#!/bin/bash

file_count=$(find . -type f -o -type d | grep -v '/\.' | wc -l)

result=$((file_count * 5))

printf "\t\vTotal files * 5: %d\v\n" "$result"