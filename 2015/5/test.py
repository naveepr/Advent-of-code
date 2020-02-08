#!/usr/bin/python

import puzzle
import unittest

class TestBasic(unittest.TestCase):
    def setUp(self):
        pass

    def test_case_1(self):
        self.assertEqual(puzzle.solve("ugknbfddgicrmopn", True), 1)

    def test_case_2(self):
        self.assertEqual(puzzle.solve("aaa", True), 1)

    def test_case_3(self):
        self.assertEqual(puzzle.solve("jchzalrnumimnmhp", True), 0)

    def test_case_4(self):
        self.assertEqual(puzzle.solve("haegwjzuvuyypxyu", True), 0)

    def test_case_5(self):
        self.assertEqual(puzzle.solve("dvszwmarrgswjxmb", True), 0)

    def test_case_2_1(self):
        self.assertEqual(puzzle.solve("qjhvhtzxzqqjkmpb", False), 1)

    def test_case_2_2(self):
        self.assertEqual(puzzle.solve("xxyxx", False), 1)

    def test_case_2_3(self):
        self.assertEqual(puzzle.solve("uurcxstgmygtbstg", False), 0)

    def test_case_2_4(self):
        self.assertEqual(puzzle.solve("ieodomkazucvgmuy", False), 0)

if __name__ == '__main__':
    unittest.main()
