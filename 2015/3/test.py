#!/usr/bin/env
import puzzle
import unittest

class TestBasic(unittest.TestCase):
    def setUp(self):
        pass

    def test_case_part1_1(self):
       content = ['^','>','v','<']
       ans=puzzle.solve(content, True)
       self.assertEqual(ans,4)

    def test_case_part1_2(self):
        with open("input.txt") as f:
            content = puzzle.parse(f)
            ans=puzzle.solve(content, True)
            self.assertNotEqual(ans,0)
        f.close()

    def test_case_part2_1(self):
       content = ['^','>','v','<']
       ans=puzzle.solve(content, False)
       self.assertEqual(ans,3)

    def test_case_part2_2(self):
        with open("input.txt") as f:
            content = puzzle.parse(f)
            ans=puzzle.solve(content, False)
            self.assertNotEqual(ans,0)
        f.close()



if __name__ == '__main__':
    unittest.main()
