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


	/** 
	 * @param {number} numRows 
	 * @returns {number[][]}
	 * */
	getPascalsTriangle(numRows) {
		
		if (numRows <= 0) return []

		/**@type {number[][]} */
		const pascals_triangle = []
		pascals_triangle.push([1])

		for (let i = 1; i < numRows; i++) {
			/** @type {number[]}*/
			const temp_rows = []
			for (let j = 0; j <= i; j++) {
				let prev = 0; let next = 0

				if (j !== 0) {
					prev = pascals_triangle[i - 1][j - 1]
				}
				if (j != i) {
					next = pascals_triangle[i - 1][j]
				}

				temp_rows.push(prev + next)
			}
			pascals_triangle.push(temp_rows)
		}
		return pascals_triangle
	}

}
