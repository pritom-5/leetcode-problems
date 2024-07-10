import { describe, expect, it } from "bun:test";
import ArrayProblems from "./01";

describe("test validEmail", () => {
	const test_data: {id: string, emails: string[], exp: number}[] =  [
		{id: "1", emails: ["test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"], exp: 2},
		{id: "2", emails: ["a@leetcode.com","b@leetcode.com","c@leetcode.com"], exp: 3}
	]


	test_data.map(({id, emails, exp}) => {
		const array_class = new ArrayProblems();

		it(`test no : ${id}`, () => {
			const result = array_class.validEmails(emails);
			expect(result).toBe(exp);
		})

	})
})
