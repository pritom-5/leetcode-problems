class TwoPointer {

	/** 
	 * @param {number[]} numbers 
	 * @param {number} target 
	 * @returns {number[]}
	 * */
	sortedTwoSum (numbers, target) {
		let left = 0, right = numbers.length - 1;


		while (left < right) {
			const left_num = numbers[left]; 
			const right_num = numbers[right]; 

			const sum = left_num + right_num;

			if (sum > target) {
				right -= 1;
			} else if (sum < target) {
				left += 1;
			} else {
				return [left + 1, right + 1];
			}
			
		}

		return [-1, -1];

	}

	/** 
	 * @param {number[]} height 
	 * @returns {number}
	 * */
	waterContainer (height) {
		let left = 0, right = height.length - 1, max = 0;

		while (right > left) {
			const width = right - left;
			const left_h = height[left], right_h = height[right];
			const curr_height = Math.min(right_h, left_h);
			const area = width * curr_height;

			max = Math.max(max, area);

			if (right_h > left_h) {
				left += 1;
			} else {
				right -= 1;
			}
		}
		
		return max;
	}
}

const two_pointer = new TwoPointer();

export {two_pointer}
