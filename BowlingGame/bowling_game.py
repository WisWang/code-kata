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
        self.pins = [0] * 21
        self.current_pin_index = 0

    def roll(self, pin):
        self.pins[self.current_pin_index] = pin
        self.current_pin_index += 1

    def score(self):
        frame_index=0
        frame= []
        for pin_index in range(len(self.pins)):
            if frame_index >= 10:
                return self._score
            self._score += self.pins[pin_index]
            if self.pins[pin_index] == 10:
                self._score += self.pins[pin_index+1] + self.pins[pin_index+2]
                frame_index += 1
                frame = []
            else:
                frame.append(self.pins[pin_index])
                if len(frame) == 2:
                    if sum(frame) == 10:
                        self._score += self.pins[pin_index+1]
                    frame_index += 1
                    frame = []






