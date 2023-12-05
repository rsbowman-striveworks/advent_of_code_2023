#!/usr/bin/env pythonv3

import sys
from pprint import pprint

def main():
    total_score = 0
    for line in sys.stdin.readlines():
        winning_str, numbers_str = line[line.find(":") + 1:].split(" | ")
        winning = set(int(n) for n in winning_str.split())
        numbers = [int(n) for n in numbers_str.split()]
        score = 0
        for n in numbers:
            if n in winning:
                score = 1 if score == 0 else 2 * score
        total_score += score
    print(total_score)


def main2():
    scratchpads = {}

    for i, line in enumerate(sys.stdin.readlines()):
        winning_str, numbers_str = line[line.find(":") + 1:].split(" | ")
        winning = set(int(n) for n in winning_str.split())
        numbers = [int(n) for n in numbers_str.split()]
        n_matches = sum(1 if n in winning else 0 for n in numbers)
        scratchpads[i] = scratchpads.get(i, 1)
        n_extra_pads = scratchpads[i]

        # we get scratchpads[i] pads for each pad with i < j <= i + n_matches
        for j in range(i + 1, i + n_matches + 1):
            scratchpads[j] = scratchpads.get(j, 1) + n_extra_pads

    print(sum(scratchpads.values()))


main2()
