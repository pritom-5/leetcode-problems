import test, { describe } from "node:test";
import ArrayProblems from "./one.js";
import { deepEqual, deepStrictEqual, strictEqual } from "node:assert";

describe("test greatestElementOnRight", () => {
	const test_data = [
		{id: "01", arr: [17,18,5,4,6,1], exp: [18,6,6,6,1,-1]}
	]

	const new_arr_cls = new ArrayProblems()

	test_data.forEach(({id, arr, exp}) => {
		test(`Test: ${id}`, () => {
			const got = new_arr_cls.greatestElementOnRight(arr) 
			deepEqual(got, exp)
		})
	})
})

describe("test getPascalsTriangle", () => {
	const test_data = [
		{id: "01", numRows: 1, exp: [[1]]},
		{id: "02", numRows: 2, exp: [[1], [1, 1]]},
		{id: "03", numRows: 3, exp: [[1], [1, 1], [1, 2, 1]]},
	]

	const new_arr_cls = new ArrayProblems()

	test_data.forEach(({id, numRows, exp}) => {
		test(`Test: ${id}`, () => {
			const got = new_arr_cls.getPascalsTriangle(numRows)
			deepStrictEqual(got, exp)
		})

	})
})
