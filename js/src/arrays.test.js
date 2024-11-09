
import { deepEqual, equal, strictEqual } from "node:assert";
import test, { describe,  } from "node:test";
import { arrays } from "./arrays.js";

describe("test hasDuplicate", () => {
	/** @type {ArrayTestType[]}*/
	const has_duplicate_test_values = [
		{id: "1", nums: [1, 2, 3], exp: false},
		{id: "2", nums: [1, 2, 2, 3], exp: true},
		{id: "3", nums: [1, 2, 2, 3, 4], exp: true}
	]

	has_duplicate_test_values.forEach(({id, nums, exp}) => {
		test(`test no: ${id}`, () => {
			const got_value = arrays.containsDuplicate(nums);
			strictEqual(got_value, exp);
		})
	})

})

describe ("validAnagramWithUnicode", () => {
	/** 
	 * @typedef {{id: string, s: string, t: string, exp: boolean}} ValidAnagramTestType
	 * @type {ValidAnagramTestType[]}*/
	const valid_anagram_test_values = [
		{id: "1", s: "anagram", t: "nagaram", exp: true},
		{id: "2", s: "car", t: "rat", exp: false},
		{id: "3", s: "car☀️", t: "☀️rac", exp: true},
	]

	valid_anagram_test_values.forEach(({id, s, t, exp}) => {
		test(`test no: ${id}`, () => {
			const got_value = arrays.validAnagramWithUnicode(s, t);
			strictEqual(exp, got_value);
		})

	})
})



describe ("validAnagram", () => {
	/** 
	 * @typedef {{id: string, s: string, t: string, exp: boolean}} ValidAnagramTestType
	 * @type {ValidAnagramTestType[]}*/
	const valid_anagram_test_values = [
		{id: "1", s: "anagram", t: "nagaram", exp: true},
		{id: "2", s: "car", t: "rat", exp: false},
		{id: "3", s: "a", t: "ab", exp: false},
	]

	valid_anagram_test_values.forEach(({id, s, t, exp}) => {
		test(`test no: ${id}`, () => {
			const got_value = arrays.validAnagram(s, t);
			strictEqual(exp, got_value);
		})

	})
})



describe ("twoSum", () => {

	/** 
	 * @typedef {{id: string, nums: number[], target: number, exp: number[]}} TwoSumType
	 * @type {TwoSumType[]}
	 * */
	const two_sum_test_values = [
		{id: "1", nums: [2, 7, 11, 15], target: 9, exp: [0, 1]},
		{id: "2", nums: [3, 2, 4], target: 6, exp: [1, 2]},
		{id: "3", nums: [3, 3], target: 6, exp: [0, 1]},
	]

	two_sum_test_values.forEach(({id, nums, target, exp}) => {
		test(`test no: ${id}`, () => {
			const _got_result = arrays.twoSum(nums, target);
			deepEqual(_got_result, exp);
		})
	})

})


/** 
 * @param {string[][]} arr 
 * @returns {string[][]}
 * */
function getSortedDeepArr (arr) {
	return arr.map(_inner => _inner.sort()).sort();
}


describe ("groupValidAnagrams", () => {
	/** 
	 * @typedef {{id: string, str: string[], exp: string[][]}} Group_valid_anagram_test_type
	 * @type{Group_valid_anagram_test_type[]}
	 * */
	const group_valid_anagram_test_data = [
		{id: "1", str: ["eat","tea","tan","ate","nat","bat"], exp: [["bat"],["nat","tan"],["ate","eat","tea"]]},
		{id: "2", str: ["a"], exp: [["a"]]},
	]

	group_valid_anagram_test_data.forEach(({id, str, exp}) => {

		test(`test no: ${id}`, () => {
			const _got_result = arrays.groupValidAnagrams(str);

			const _sorted_got_result = getSortedDeepArr(_got_result);
			const _sorted_exp = getSortedDeepArr(exp);

			deepEqual(_sorted_got_result, _sorted_exp);
		})
	})
	
})


describe ("topKFrequent", () => {
	/** 
	 * @typedef {{id: string, nums: number[], k: number, exp: number[]}} T_Top_K_Frequent
	 * @type {T_Top_K_Frequent[]}
	 **/
	const top_k_frequent_test_data = [
		{id: "1", nums: [1,1,1,2,2,3], k: 2,exp: [1, 2]},
		{id: "2", nums: [1], k: 1 ,exp: [1]},
	]

	top_k_frequent_test_data.forEach(({id, nums, k, exp}) => {
		test(`test id: ${id}`, () => {
			const _got_result = arrays.topKFrequent(nums, k);

			deepEqual(exp, _got_result);

		})
	})
})













