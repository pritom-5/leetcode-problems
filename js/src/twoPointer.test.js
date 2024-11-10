import test, { describe } from "node:test";
import { two_pointer } from "./twoPointer.js";
import { deepEqual } from "node:assert";

describe ("sortedTwoSum", {only: false}, () => {
	/**
	 * @typedef {{id: string, numbers: number[], target: number, exp: number[]}} T_Sorted_Two_Sum_Test 
	 * @type {T_Sorted_Two_Sum_Test[]}
	 * */
	const sorted_two_sum_test_data = [
		{id: "1", numbers: [2, 7, 11, 15], target: 9, exp: [1, 2]},
		{id: "2", numbers: [5,25,75], target: 100, exp: [2, 3]},
	]

	sorted_two_sum_test_data.forEach(({id, numbers, target, exp}) => {
		test(`test no: ${id}`, {only: false},() => {
			const res = two_pointer.sortedTwoSum(numbers, target);
			deepEqual(exp, res);
		})

	})
});


describe ("waterContainer", {only: true}, () => {
	/**
	 * @typedef {{id: string, height: number[], exp: number}} T_Water_Container_Test 
	 * @type {T_Water_Container_Test[]}
	 * */
	const water_container_test_data = [
		{id: "1", height: [1,8,6,2,5,4,8,3,7], exp: 49},
		{id: "2", height: [1, 1], exp: 1},
	]

	water_container_test_data.forEach(({id, height, exp}) => {
		test(`test no: ${id}`, {only: true},() => {
			const res = two_pointer.waterContainer(height);
			deepEqual(exp, res);
		})

	})
});
