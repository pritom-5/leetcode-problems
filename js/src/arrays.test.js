import { deepEqual, equal, strict, strictEqual } from "node:assert";
import test, { describe, only } from "node:test";
import { arrays } from "./arrays.js";

describe("test hasDuplicate", () => {
  /** @type {ArrayTestType[]}*/
  const has_duplicate_test_values = [
    { id: "1", nums: [1, 2, 3], exp: false },
    { id: "2", nums: [1, 2, 2, 3], exp: true },
    { id: "3", nums: [1, 2, 2, 3, 4], exp: true },
  ];

  has_duplicate_test_values.forEach(({ id, nums, exp }) => {
    test(`test no: ${id}`, () => {
      const got_value = arrays.containsDuplicate(nums);
      strictEqual(got_value, exp);
    });
  });
});

describe("validAnagramWithUnicode", () => {
  /**
   * @typedef {{id: string, s: string, t: string, exp: boolean}} ValidAnagramTestType
   * @type {ValidAnagramTestType[]}*/
  const valid_anagram_test_values = [
    { id: "1", s: "anagram", t: "nagaram", exp: true },
    { id: "2", s: "car", t: "rat", exp: false },
    { id: "3", s: "car☀️", t: "☀️rac", exp: true },
  ];

  valid_anagram_test_values.forEach(({ id, s, t, exp }) => {
    test(`test no: ${id}`, () => {
      const got_value = arrays.validAnagramWithUnicode(s, t);
      strictEqual(exp, got_value);
    });
  });
});

describe("validAnagram", () => {
  /**
   * @typedef {{id: string, s: string, t: string, exp: boolean}} ValidAnagramTestType
   * @type {ValidAnagramTestType[]}*/
  const valid_anagram_test_values = [
    { id: "1", s: "anagram", t: "nagaram", exp: true },
    { id: "2", s: "car", t: "rat", exp: false },
    { id: "3", s: "a", t: "ab", exp: false },
  ];

  valid_anagram_test_values.forEach(({ id, s, t, exp }) => {
    test(`test no: ${id}`, () => {
      const got_value = arrays.validAnagram(s, t);
      strictEqual(exp, got_value);
    });
  });
});

describe("twoSum", () => {
  /**
   * @typedef {{id: string, nums: number[], target: number, exp: number[]}} TwoSumType
   * @type {TwoSumType[]}
   * */
  const two_sum_test_values = [
    { id: "1", nums: [2, 7, 11, 15], target: 9, exp: [0, 1] },
    { id: "2", nums: [3, 2, 4], target: 6, exp: [1, 2] },
    { id: "3", nums: [3, 3], target: 6, exp: [0, 1] },
  ];

  two_sum_test_values.forEach(({ id, nums, target, exp }) => {
    test(`test no: ${id}`, () => {
      const _got_result = arrays.twoSum(nums, target);
      deepEqual(_got_result, exp);
    });
  });
});

/**
 * @param {string[][]} arr
 * @returns {string[][]}
 * */
function getSortedDeepArr(arr) {
  return arr.map((_inner) => _inner.sort()).sort();
}

describe("groupValidAnagrams", () => {
  /**
   * @typedef {{id: string, str: string[], exp: string[][]}} Group_valid_anagram_test_type
   * @type{Group_valid_anagram_test_type[]}
   * */
  const group_valid_anagram_test_data = [
    {
      id: "1",
      str: ["eat", "tea", "tan", "ate", "nat", "bat"],
      exp: [["bat"], ["nat", "tan"], ["ate", "eat", "tea"]],
    },
    { id: "2", str: ["a"], exp: [["a"]] },
  ];

  group_valid_anagram_test_data.forEach(({ id, str, exp }) => {
    test(`test no: ${id}`, () => {
      const _got_result = arrays.groupValidAnagrams(str);

      const _sorted_got_result = getSortedDeepArr(_got_result);
      const _sorted_exp = getSortedDeepArr(exp);

      deepEqual(_sorted_got_result, _sorted_exp);
    });
  });
});

describe("topKFrequent", () => {
  /**
   * @typedef {{id: string, nums: number[], k: number, exp: number[]}} T_Top_K_Frequent
   * @type {T_Top_K_Frequent[]}
   **/
  const top_k_frequent_test_data = [
    { id: "1", nums: [1, 1, 1, 2, 2, 3], k: 2, exp: [1, 2] },
    { id: "2", nums: [1], k: 1, exp: [1] },
  ];

  top_k_frequent_test_data.forEach(({ id, nums, k, exp }) => {
    test(`test id: ${id}`, () => {
      const _got_result = arrays.topKFrequent(nums, k);

      deepEqual(exp, _got_result);
    });
  });
});

//describe("testing 124", {only: false}, () => {
//	test("test", {only: false} , () => {
//		console.log("hello there 2");
//		arrays.testValue();
//		strictEqual(1, 2);
//	})
//})

describe("longestConsecutiveSequence", () => {
  /**
   * @typedef {{id: string, nums: number[], exp: number}} T_Longest_Consecutive_Sequence_test_data
   * @type {T_Longest_Consecutive_Sequence_test_data[]}
   * */

  const longest_consecutive_cequence_test_data = [
    { id: "1", nums: [100, 4, 200, 1, 3, 2], exp: 4 },
    { id: "2", nums: [0, 3, 7, 2, 5, 8, 4, 6, 0, 1], exp: 9 },
    { id: "3", nums: [9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6], exp: 7 },
  ];

  longest_consecutive_cequence_test_data.forEach(({ id, nums, exp }) => {
    test(`test id: ${id}`, () => {
      const output = arrays.longestConsecutiveSequence(nums);
      strictEqual(output, exp);
    });
  });
});

describe("validSudoku", { only: false }, () => {
  /** @type {T_Board}*/
  const board_1 = [
    ["5", "3", ".", ".", "7", ".", ".", ".", "."],
    ["6", ".", ".", "1", "9", "5", ".", ".", "."],
    [".", "9", "8", ".", ".", ".", ".", "6", "."],
    ["8", ".", ".", ".", "6", ".", ".", ".", "3"],
    ["4", ".", ".", "8", ".", "3", ".", ".", "1"],
    ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
    [".", "6", ".", ".", ".", ".", "2", "8", "."],
    [".", ".", ".", "4", "1", "9", ".", ".", "5"],
    [".", ".", ".", ".", "8", ".", ".", "7", "9"],
  ];

  /** @type {T_Board}*/
  const board_2 = [
    ["8", "3", ".", ".", "7", ".", ".", ".", "."],
    ["6", ".", ".", "1", "9", "5", ".", ".", "."],
    [".", "9", "8", ".", ".", ".", ".", "6", "."],
    ["8", ".", ".", ".", "6", ".", ".", ".", "3"],
    ["4", ".", ".", "8", ".", "3", ".", ".", "1"],
    ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
    [".", "6", ".", ".", ".", ".", "2", "8", "."],
    [".", ".", ".", "4", "1", "9", ".", ".", "5"],
    [".", ".", ".", ".", "8", ".", ".", "7", "9"],
  ];

  /** @type {T_Board}*/
  const board_3 = [
    [".", ".", ".", ".", "5", ".", ".", "1", "."],
    [".", "4", ".", "3", ".", ".", ".", ".", "."],
    [".", ".", ".", ".", ".", "3", ".", ".", "1"],
    ["8", ".", ".", ".", ".", ".", ".", "2", "."],
    [".", ".", "2", ".", "7", ".", ".", ".", "."],
    [".", "1", "5", ".", ".", ".", ".", ".", "."],
    [".", ".", ".", ".", ".", "2", ".", ".", "."],
    [".", "2", ".", "9", ".", ".", ".", ".", "."],
    [".", ".", "4", ".", ".", ".", ".", ".", "."],
  ];

  /** @type {T_Board}*/
  const board_4 = [
    [".", ".", "4", ".", ".", ".", "6", "3", "."],
    [".", ".", ".", ".", ".", ".", ".", ".", "."],
    ["5", ".", ".", ".", ".", ".", ".", "9", "."],
    [".", ".", ".", "5", "6", ".", ".", ".", "."],
    ["4", ".", "3", ".", ".", ".", ".", ".", "1"],
    [".", ".", ".", "7", ".", ".", ".", ".", "."],
    [".", ".", ".", "5", ".", ".", ".", ".", "."],
    [".", ".", ".", ".", ".", ".", ".", ".", "."],
    [".", ".", ".", ".", ".", ".", ".", ".", "."],
  ];

  /**
   * @typedef {{id: string, board: T_Board, exp: boolean}} T_Valid_Sudoky_Test_Data
   * @type {T_Valid_Sudoky_Test_Data[]}
   * */
  const valid_sudoku_test_data = [
    { id: "4", board: board_4, exp: false },
    { id: "1", board: board_1, exp: true },
    { id: "2", board: board_2, exp: false },
    { id: "3", board: board_3, exp: false },
  ];

  valid_sudoku_test_data.forEach(({ id, board, exp }) => {
    test(`test id : ${id}`, { only: false }, () => {
      const output = arrays.validSudoku(board);
      strictEqual(output, exp);
    });
  });
});
