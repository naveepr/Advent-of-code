#!/usr/bin/python

import puzzle
import unittest

class TestBasic(unittest.TestCase):
    def setUp(self):
        pass

    def test_case_1(self):
        ans = str(puzzle.solve('abcdef', True))
        self.assertEqual(ans,'609043')

    def test_case_2(self):
        ans = str(puzzle.solve('pqrstuv', True))
        self.assertEqual(ans,'1048970')

    def test_case_3(self):
        ans = str(puzzle.solve('bgvyzdsv', True))
        self.assertEqual(ans,'254575')

    def test_case_4(self):
        ans = str(puzzle.solve('bgvyzdsv', False))
        self.assertEqual(ans,'1038736')

if __name__ == '__main__':
    unittest.main()
