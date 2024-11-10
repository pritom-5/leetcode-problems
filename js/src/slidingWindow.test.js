import test, { describe } from "node:test";
import { sliding_window } from "./slidingWindow.js";
import { strictEqual } from "node:assert";

describe ("maxProfit", {only: false}, () => {
	/** 
	 * @typedef {{id: string, prices: number[], exp: number}} T_Max_Profit_Test_Data
	 * @type {T_Max_Profit_Test_Data[]}
	 * */
	const max_profit_test_data = [
		{id: "1", prices: [7,1,5,3,6,4], exp: 5}, 
		{id: "2", prices: [7,6,4,3,1], exp: 0}, 
		{id: "3", prices: [1, 2], exp: 1}, 
	]

	max_profit_test_data.forEach(({id, prices, exp}) => {

		test(`test no: ${id}`, {only: false}, () => {
			const output = sliding_window.maxProfit(prices);
			strictEqual(output, exp);
		})

	})

})


describe("longestSubstring", {only: false},() => {

	/** 
	 * @typedef {{id: string, s: string, exp: number}} T_Longest_Substring
	 * @type {T_Longest_Substring[]}
	 * */
	const longest_substring_test_data = [
		{id: "1", s: "abcabcbb", exp: 3}, 
		{id: "2", s: "bbbbb", exp: 1}, 
		{id: "3", s: "dvdf", exp: 3}, 
	]

	longest_substring_test_data.forEach(({id, s, exp}) => {
		test(`test no : ${id}`, {only: false},() => {
			const output = sliding_window.longestSubstring(s);
			strictEqual(output, exp);
		})
	})

})

describe("characterReplacement", {only: true}, () => {

	/** 
	 * @typedef {{id: string, s: string, k: number, exp: number}} T_Character_Replacement
	 * @type {T_Character_Replacement[]}
	 * */
	const character_replacement_test_data = [
		{id: "1", s: "ABAB", k: 2, exp: 4},
		{id: "2", s: "AABABBA", k: 1, exp: 4},
	]

	character_replacement_test_data.forEach(({id, s, k, exp}) => {
		test(`test id: ${id}`, {only: true}, () => {
			const output = sliding_window.characterReplacement(s, k);
			strictEqual(output, exp);
		})
		
	})
	
})
