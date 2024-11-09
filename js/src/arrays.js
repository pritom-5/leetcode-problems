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

		/** @type {number[]}*/
		const final_results = [];


		const top_k_frequencies = Object.
			entries(frequency_counter).
			sort((a, b) => b[1] - a[1]).
			slice(0, k).
			map(item => final_results.push(Number(item[0])))
		;

		return final_results;
	}

	/** 
	 * @param {number[]} nums 
	 * @returns {number}
	 * */
	longestConsecutiveSequence (nums) {
		const sorted_nums = nums.sort((a, b) => a - b);

		let longest_sequence = 0;
		let temp_longest_sequence = 1;

		for (let i = 0; i < sorted_nums.length - 1; i++) {
			if (sorted_nums[i + 1] - sorted_nums[i] === 1) {
				temp_longest_sequence += 1;
			} else {
				temp_longest_sequence = 1;
			}

		 longest_sequence	= Math.max(longest_sequence, temp_longest_sequence);
		}

		return longest_sequence;
	}

	/** @param {T_Board_item[]} nums */
	#checkSingleLineOfSudoku (nums) {

		for (let i = 0; i < nums.length; i++) {

			/** @type {Record<string, number>}*/
			const _map = {}

			for (let j = 0; j < 9; j++) {
				const _num = nums[j];

				if (_num in _map && _num !== ".") {
					return false;
				} else {
					_map[_num] = 1;
				}
			}
		}

		return true;
	}


	/** @param {T_Board} board 
	 * @returns {boolean}*/
	#checkHorizontal (board) {
		for (let i = 0; i < 9; i++) {
			const _hor_line = board[i];
			const _output = this.#checkSingleLineOfSudoku(_hor_line);
			if (!_output) {
				return false;
			}
		}

		return true;
	}



	/** @param {T_Board} board 
	 * @returns {boolean}*/
	#checkVertical (board) {
		for (let i = 0; i < 9; i++) {
			const _ver_line = board[i];
			const _output = this.#checkSingleLineOfSudoku(_ver_line);
			if (!_output) {
				return false;
			}

		}

		return true;
	}

	/** @param {T_Board} board 
	 * @returns {boolean}
	 * */
	#check3x3Box (board) {
		for (let i = 0; i < 9; i += 3) {
			for (let j = 0; j < 9; j += 3) {

				const _nums = [];

				for (let k = i; k < i + 3; k++) {
					for (let l = i; l < i + 3; l++) {
						_nums.push(board[k][l]);
					}
				}

				console.log("lines:-----> ", _nums);

				const _output = this.#checkSingleLineOfSudoku(_nums);
				if (!_output)  {
					return false;
				}
				
			}
		}

		return true;

	}

	/** 
	 * @param {T_Board} board 
	 * @returns {boolean}
	 * */
	validSudoku (board) {

		return this.#checkVertical(board) && this.#checkHorizontal(board) && this.#check3x3Box(board);
	}


	testValue () {
		console.log("hello there");
	}

}

const arrays = new Arrays();

export {arrays};














