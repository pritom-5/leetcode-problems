import { log } from "node:console";

class Arrays {
	/** 
	 * @param {Array<number>} nums 
	 * @returns {boolean}
	 * */
	containsDuplicate (nums) {
		/** 
		 * @type {Record<string, number>}
		 * */
		const new_map = {}

		for (const num of nums) {
			if (num in new_map) {
				return true;
			}
			new_map[num] = 1;
		}

		return false;
	}

	/** 
	 * @param {string} s
	 * @param {string} t
	 * @returns {boolean}
	 * */
	validAnagram (s, t) {
		if (s.length !== t.length) {
			return false;
		}

		/** @type {number[]}*/
		const ch_arr = new Array(26).fill(0);

		s.split("").forEach(ch => {
			const _ch_code = ch.charCodeAt(0);
			ch_arr[_ch_code - 98] += 1;
		})

		for (const _ch of t.split("")) {
			const _ch_code = _ch.charCodeAt(0);
			ch_arr[_ch_code - 98] -= 1;
		}


		for (const _ch_nos of ch_arr) {
			if (_ch_nos !== 0) {
				return false;
			}
		}

		return true;
	}

	/** 
	 * @param {string} s 
	 * @param {string} t 
	 * @returns {boolean}
	 * */
	validAnagramWithUnicode (s, t) {
		/** @type {Record<string, number>} */
		const character_map = {}

		s.split("").forEach(_ch => {
			if (_ch in character_map) {
				character_map[_ch] += 1;
			} else {
				character_map[_ch] = 1;
			}
		})

		for (const _ch of t.split("")) {
			if (_ch in character_map) {
				character_map[_ch] -= 1;
			} else {
				character_map[_ch] = -1;
			} 
		}


		for (const _val of Object.values(character_map)) {
			if (_val !== 0) {
				return false;
			}
		}

		return true;
	}

	/** 
	 * @param {number[]} nums
	 * @param {number} target
	 * @returns {number[]}
	 * */
	twoSum (nums, target) {
		for (let i = 0; i < nums.length; i++) {
			const _num = nums[i];
			const _diff = target - _num;

			for (let j = i + 1; j < nums.length; j++) {
				if (nums[j] === _diff && i !== j) {
					return [i, j];
				}
			}
		}

		return [-1, -1];
	}

	/** 
	 * @param {string[]} strs 
	 * @returns {string[][]}
	 * */
	groupValidAnagrams (strs) {
		/** @type {Record<string, boolean>} */
		const included_already = {};

		/** @type {string[][]}*/
		const grouped_anagrams = [];


		for (let i = 0; i < strs.length; i++) {
			const _str = strs[i];


			if (_str in included_already) {
				continue;
			}

			const _groups = [_str];

			for (let j = i + 1; j < strs.length; j++) {
				const _str_1 = strs[j];


				if (this.validAnagram(_str, _str_1)) {
					_groups.push(_str_1);

					included_already[_str] = true;
					included_already[_str_1] = true;
				}
			}

			console.log(_groups);

			grouped_anagrams.push(_groups);


		}

		return grouped_anagrams;

	}

	/** 
	 * @param {number[]} nums 
	 * @param {number} k 
	 * @returns {number[]}
	 * */
	topKFrequent (nums, k) {

		/** @type {Record<number, number>}*/
		const frequency_counter = {}

		for (const _num of nums) {
			if (_num in frequency_counter) {
				frequency_counter[_num] += 1;
			} else {
				frequency_counter[_num] = 1;
			}
		}

		const top_k_frequencies = Object.
			values(frequency_counter).
			sort((a, b) => a - b).
			slice(0, k);

		return top_k_frequencies;
	}

}

const arrays = new Arrays();

export {arrays};














