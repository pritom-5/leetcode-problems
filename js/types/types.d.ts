export {};


declare global {
	type One = {
		firstName: string,
		secondName: string,
		age: number
	}

	type ArrayTestType = {
		id: string,
		nums: number[], 
		exp: boolean
	}

	type T_Board_item = string
	type T_Board = T_Board_item[][]

}
