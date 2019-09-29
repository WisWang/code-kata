#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Copyright (c) 2017 - hongzhi.wang <hongzhi.wang@moji.com> 
'''
Author: hongzhi.wang
Create Date: 2019-09-04
Modify Date: 2019-09-04
'''

import unittest

from .bowling_game import BowlingGame


class TestBowlingGame(unittest.TestCase):

    def setUp(self):
        self.g = BowlingGame()

    def test_game_all_zero(self):
        self.roll_many(20, 0)
        self.assertEqual(0, self.g.score())

    def test_game_all_one(self):
        self.roll_many(20, 1)
        self.assertEqual(20, self.g.score())

    def roll_many(self, times, pin):
        for i in range(times):
            self.g.roll(pin)

    def test_game_first_strike(self):
        self.g.roll(4)
        self.g.roll(6)
        self.roll_many(18, 1)
        self.assertEqual(29, self.g.score())

    def test_perfect_game(self):
        self.roll_many(12, 10)
        self.assertEqual(300, self.g.score())

    def test_game_first_two_strike(self):
        self.roll_many(4, 5)
        self.roll_many(16, 1)
        self.assertEqual(42, self.g.score())