#!/usr/bin/env python3

valid = 0
validanagram = 0

with open("input", "r") as f:
    for line in f:
        words = line.strip().split()
        checked = []
        checkedanagram = []
        isvalid = True
        isvalidana = True
        for word in words:
            if word in checked:
                isvalid = False
            checked.append(word)
            ana = ''.join(sorted(word))
            if ana in checkedanagram:
                isvalidana = False
            checkedanagram.append(ana)
        if isvalid:
            valid = valid + 1
        if isvalidana:
            validanagram = validanagram + 1
print(valid, validanagram)
