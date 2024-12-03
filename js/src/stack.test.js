import test, {describe, only} from "node:test";
import { stack } from "./stack.js";
import { deepEqual, strictEqual } from "node:assert";


describe("polishNotation", {only: false}, () => {

	/** 
	 * @typedef {{id: string, tokens: string[], exp: number}} T_Test_polish_natation
	 * @type {T_Test_polish_natation[]}
	 * */
	const test_data = [
		{id: "1", tokens: ["10","6","9","3","+","-11","*","/","*","17","+","5","+"], exp: 22}
	]

	test_data.forEach(({id, tokens, exp}) => {
		test(`test no: ${id}`, {only: false}, () => {
			const output = stack.polishNotation(tokens);
			strictEqual(output, exp);
		})
	})
})

describe("nextTemperatures", {only: false}, () => {

	/** 
	 * @typedef {{id: string, temps: number[], exp: number[]}} T_Temps_Test_data
	 * @type {T_Temps_Test_data[]}
	 * */
	const test_data = [
		{id: "1", temps: [73,74,75,71,69,72,76,73], exp: [1,1,4,2,1,1,0,0]},
		{id: "2", temps: [30, 40, 50, 60], exp: [1, 1, 1, 0]},
		{id: "3", temps: [30, 60, 90], exp: [1, 1, 0]},
	]

	test_data.forEach(({id, temps, exp}) => {
		test(`test id: ${id}`, {only: false}, () => {
			const output = stack.nextTemperatures(temps);
			deepEqual(exp, output);
		})
	})
	
})
