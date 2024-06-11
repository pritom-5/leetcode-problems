import test, { describe } from "node:test";
import ArrayProblems from "./one.js";
import { deepEqual } from "node:assert";

describe("test greatestElementOnRight", () => {
	const test_data = [
		{id: "01", arr: [17,18,5,4,6,1], exp: [18,6,6,6,1,-1]}
	]

	const new_arr_cls = new ArrayProblems()

	test_data.forEach(({id, arr, exp}) => {
		test(`Test: ${id}`, () => {
			const got = new_arr_cls.greatestElementOnRight(arr) 
			console.log(got)
			deepEqual(got, exp)
		})

	})
})
