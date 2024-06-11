export default class ArrayProblems {


	/** 
	 * @param {number[]} arr 
	 * @returns {number[]}*/
	greatestElementOnRight(arr) {
		let g = -1

		/** @type {number[]}*/
		const new_arr = Array(arr.length).fill(-1)

		for (let i = arr.length - 1; i >= 0; i--) {
			new_arr[i] = g
			g = Math.max(g, arr[i])
		}


		return new_arr
	}

}
