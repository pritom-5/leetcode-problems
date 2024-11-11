import test, {describe} from "node:test";
import { stack } from "./stack.js";
import { strictEqual } from "node:assert";


describe("polishNotation", {only: true}, () => {

	/** 
	 * @typedef {{id: string, tokens: string[], exp: number}} T_Test_polish_natation
	 * @type {T_Test_polish_natation[]}
	 * */
	const test_data = [
		{id: "1", tokens: ["10","6","9","3","+","-11","*","/","*","17","+","5","+"], exp: 22}
	]

	test_data.forEach(({id, tokens, exp}) => {
		test(`test no: ${id}`, {only: true}, () => {
			const output = stack.polishNotation(tokens);
			strictEqual(output, exp);
		})
	})
})
