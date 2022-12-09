#!/bin/bash
p1="phase1/phase1.go"
p2="phase2/phase2.go"

if [ "$1" = "" ]; 
then
  read -p "Day#" n
fi;

if [ !-d $1 ]; then
  dir="day$n"
  mkdir -p $dir/phase1 $dir/phase2
  touch $dir/demo.txt $dir/input.txt $dir/$p1 $dir/$p2
fi;

echo "Successfully created ./$dir!"