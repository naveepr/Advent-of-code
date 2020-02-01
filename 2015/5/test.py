#!/usr/bin/python

import puzzle
import unittest

class TestBasic(unittest.TestCase):
    def setUp(self):
        pass

    def test_case_1(self):
        self.assertEqual(puzzle.solve("ugknbfddgicrmopn"), 1)

    def test_case_2(self):
        self.assertEqual(puzzle.solve("aaa"), 1)

    def test_case_3(self):
        self.assertEqual(puzzle.solve("jchzalrnumimnmhp"), 0)

    def test_case_4(self):
        self.assertEqual(puzzle.solve("haegwjzuvuyypxyu"), 0)

    def test_case_5(self):
        self.assertEqual(puzzle.solve("dvszwmarrgswjxmb"), 0)

if __name__ == '__main__':
    unittest.main()
