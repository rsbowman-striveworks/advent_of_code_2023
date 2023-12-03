#!/usr/bin/env python3

import sys

def parse(line: str):
    numbers, parts = [], []
    i = 0
    while i < len(line):
        if line[i].isdigit():
            i_start = i
            while i < len(line) and line[i].isdigit():
                i += 1
            i_end = i
            n = int(line[i_start:i_end])
            numbers.append((n, i_start, i_end))
            i -= 1
        elif line[i] != ".":
            parts.append(i)
        i += 1

    return numbers, parts

def discharge_numbers(numbers, parts):
    """
    `numbers` and `parts` both sorted
    """
    discharged, remaining = [], []
    i_part, i_number = 0, 0
    n_numbers = len(numbers)

    while i_part < len(parts):
        while i_number < n_numbers and numbers[i_number][2] < parts[i_part]: # end of number < part
            remaining.append(numbers[i_number])
            i_number += 1
        if i_number >= n_numbers:
            break
        elif numbers[i_number][1] - 1 <= parts[i_part]:
            discharged.append(numbers[i_number][0])
            i_number += 1
        else:
            i_part += 1

    while i_number < n_numbers:
        remaining.append(numbers[i_number])
        i_number += 1

    return discharged, remaining


# print(discharge_numbers([(75, 43, 45), (50, 54, 56), (89, 92, 94)], [38, 55, 62]))

def main():
    numbers_prev = []
    parts_prev = []

    part_number_sum = 0

    for line in sys.stdin.readlines():
        numbers, parts = parse(line.strip())
        # discharge any numbers in `numbers_prev` using `parts`
        discharged_numbers, _ = discharge_numbers(numbers_prev, parts)
        part_number_sum += sum(discharged_numbers)

        # discharge any numbers in `numbers` using `parts_prev`
        discharged_numbers, remaining_numbers = discharge_numbers(numbers, parts_prev)
        part_number_sum += sum(discharged_numbers)

        # discharge any numbers in `numbers` using `parts`
        discharged_numbers, remaining_numbers = discharge_numbers(remaining_numbers, parts)
        part_number_sum += sum(discharged_numbers)

        numbers_prev = remaining_numbers
        parts_prev = parts

    print(part_number_sum)


main()
