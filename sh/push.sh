#!/bin/bash

git add -A .


# if the first parameter is provided, use it as the commit message
if [ -n "$1" ]; then
  git commit -m "$1"
fi

git commit -m "."
git push

