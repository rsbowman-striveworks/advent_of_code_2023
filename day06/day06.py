#!/usr/bin/env python3

import sys
import math


def n_ways_to_win(t, d_min):
    """
    Number of ways we can beat `d_min` mm in `t` ms playing the game.

    Also the number of integers t_h so that

    (t - t_h) * t_h > d_min

    Rearranging we get -t_h^2 + T*t_h + d_min > 0; solve for the distance
    between the zeros t_0_1 and t_0_2 of this function

    Should work if discriminant is nonnegative
    """
    root_term = math.sqrt(t**2 - 4 * d_min)
    t_0_1 = int(math.ceil((t - root_term) / 2))
    t_0_2 = int(math.floor((t + root_term) / 2))

    if (t - t_0_1) * t_0_1 <= d_min:
        t_0_1 += 1
    if (t - t_0_2) * t_0_2 <= d_min:
        t_0_2 -= 1

    return t_0_2 - t_0_1 + 1


def main():
    lines = list(sys.stdin.readlines())

    def parse_ints(l):
        return [int(n) for n in l[l.find(":") + 1 :].split()]

    def parse_one_big_int(l: str):
        return int(l[l.find(":") + 1 :].replace(" ", ""))

    if False:  # do the first part of the problem
        ts = parse_ints(lines[0])
        ds = parse_ints(lines[1])
        assert len(ts) == len(ds)
    else:  # do part 2
        ts = [parse_one_big_int(lines[0])]
        ds = [parse_one_big_int(lines[1])]

    answer_prod = 1
    for t, d in zip(ts, ds):
        n = n_ways_to_win(t, d)
        answer_prod *= n
    print(answer_prod)


main()
