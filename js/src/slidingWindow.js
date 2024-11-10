class SlidingWindow {
/** 
 * @param {number[]} prices 
 * @returns {number}
 * */
 maxProfit (prices) {
		let left = 0, right = 0, max_profit = 0;

		while (right < prices.length) {
			const diff = prices[right] - prices[left];

			if (diff > 0) {
				max_profit = Math.max(max_profit, diff);
			} else {
				left = right;
			}

			right += 1;
		}

		return max_profit;
	}

	/** 
	 * @param {string} s 
	 * @returns {number}
	 * */
	longestSubstring (s) {
		let left = 0, right = 1, max_substring = 1;

		while (right < s.length) {
			const tmp_substring = s.slice(left, right + 1);

			if (!this.#hasRepeat(tmp_substring)) {
				max_substring = Math.max(max_substring, tmp_substring.length);
				right+= 1;
			} else {
				left += 1;
			}

		}

		return max_substring;

	}

	/** 
	 * @param {string} s 
	 * @returns {boolean}
	 * */
	#hasRepeat (s) {
		/** @type {Record<string, boolean>}*/
		const map = {}

		for (let i = 0; i < s.length; i++) {
			const tmp_ch = s[i];

			if (tmp_ch in map) {
				return true;
			}
			map[tmp_ch] = true;
		}
		return false;
	}

	/** 
	 * @param {string} s 
	 * @param {number} k 
	 * @returns {number}
	 * */
	characterReplacement (s, k) {
		let left = 0, right = 0, max = 0;

		while (right < s.length) {
			const tmp_substring = s.slice(left, right + 1);

			if (this.#nosOfLessFrequentCharacters(tmp_substring) <= k) {
				max = Math.max(max, tmp_substring.length);
				right += 1;
			} else {
				left += 1;
			}
		}

		return max;
	}

	/** 
	* @param {string} s 
	*/
	#nosOfLessFrequentCharacters (s) {

		/** @type {Record<string, number>}*/
		const ch_map = {};

		for (let i = 0; i < s.length; i++) {
			const tmp_ch = s[i];
			if (tmp_ch in ch_map) {
				ch_map[tmp_ch] += 1;
			} else {
				ch_map[tmp_ch] = 1;
			}
		}

		const nos_of_most_frequent_ch = Math.max(...Object.values(ch_map))

		return s.length - nos_of_most_frequent_ch;


	}

}

const sliding_window = new SlidingWindow();
export {sliding_window};



