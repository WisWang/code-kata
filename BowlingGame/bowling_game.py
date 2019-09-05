#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Copyright (c) 2017 - hongzhi.wang <hongzhi.wang@moji.com> 
'''
Author: hongzhi.wang
Create Date: 2019-09-04
Modify Date: 2019-09-04
'''


class BowlingGame:
    def __init__(self):
        self._score = 0
        self.pins = []
        self.current_pin_index = 0

    def roll(self, pin):
        self.pins.append(pin)

    def score(self):
        for pin in self.pins:
            self._score += pin

        return self._score
