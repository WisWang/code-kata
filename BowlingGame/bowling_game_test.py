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
        for i in range(20):
            self.g.roll(0)
        self.assertEqual(0, self.g.score())

    def test_game_all_one(self):
        for i in range(20):
            self.g.roll(1)
        self.assertEqual(20, self.g.score())


    # def test_game_first_strike(self):
    #     self.g.roll(5)
    #     self.g.roll(5)
    #     for i in range(18):
    #         self.g.roll(1)
    #     self.assertEqual(29, self.g.score())
