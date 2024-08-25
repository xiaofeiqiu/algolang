package main

/*

Problem Explanation:
We are working with two strings: ref_string and src_string. A "cover" is a way to describe how portions of src_string map to portions of ref_string. This "cover" consists of pairs of indices (start, end) that define "blocks" of characters in ref_string that match portions of src_string.

For example:

The block (7, 11) in ref_string represents the substring "abcd" (since ref_string[7:11] == "abcd").
The block (3, 6) represents the substring "123" (since ref_string[3:6] == "123").
The cover describes how to reconstruct the src_string by piecing together these blocks. For instance, the cover [(7, 11), (3, 6)] represents the src_string "abcd123", and [(7, 10), (10, 11), (3, 6)] represents "abc"+"d"+"123" which also reconstructs src_string.

Task
You are required to implement the function delete(cover, index) which, given:

A valid cover (a list of pairs of indices),
An index in src_string,
The function will simulate deleting the character at index from src_string and return a new valid cover. This new cover should represent how the src_string would be compressed with the character at index removed.

Example:
Suppose the src_string is "abcd123", and the cover is [(7, 11), (3, 6)].

delete(cov1, 3): This simulates removing the character at index 3 of "abcd123", resulting in "abc123". We need to adjust the blocks of the cover to reflect this change. The resulting cover would be [(7, 10), (3, 6)], where the first block "abcd" is reduced by one character to "abc".

delete(cov1, 5): This simulates removing the character at index 5 from "abcd123", resulting in "abcd13". The cover must be adjusted to reflect the missing character.

delete(cov1, 0): This simulates removing the first character (index = 0), which results in "bcd123". The cover must be updated accordingly.

*/

/*
ref_string = "abc1234abcd!
src_string = "abcd123"
cov = [(7, 11), (3, 6)], cover can be used to find src string, each block in cov represents substring in ref_string
block (7,11) represents block (0, 0 + blockLen) = (0, 4)
delete(cov, 2): "c" will be deleted


idea:
for each given block
	check if index is before, after, within the block
		if before, shift left by one
		if after, no change
		if within
			first: shift start, end to left by 1
			somewhere middle: split
			end: shift only end to left by 1
*/

type Block struct {
	Start int
	End   int
}

func delete1(cover []Block, index int) []Block {
	if len(cover) == 0 {
		return []Block{}
	}

	srcStart := 0
	res := make([]Block, 0)

	for _, block := range cover {
		blockLen := block.End - block.Start

		// calculate src block end
		srcEnd := srcStart + blockLen

		// Handle cases where index is outside the current block
		if index >= srcEnd || index < srcStart {
			res = append(res, block)
			srcStart += blockLen
			continue
		}

		// Handle all cases where index is within the current block

		// delete first in src block
		if index == srcStart {
			if block.Start+1 != block.End {
				res = append(res, Block{Start: block.Start + 1, End: block.End})
			}
			srcStart += blockLen
			continue
		}

		// delete last in the src block
		if index == srcEnd-1 {
			if block.Start+1 != block.End {
				res = append(res, Block{Start: block.Start, End: block.End - 1})
			}
			srcStart += blockLen
			continue
		}

		// Handle all cases where index is within the current block
		deletionOffset := index - srcStart
		res = append(res, Block{Start: block.Start, End: block.Start + deletionOffset})
		res = append(res, Block{Start: block.Start + deletionOffset + 1, End: block.End})

		/*
				or we can write simpler way
			        // Handle all cases where index is within the current block
			        deletionOffset := index - srcStart
			        if deletionOffset > 0 {
			            res = append(res, Block{Start: block.Start, End: block.Start + deletionOffset})
			        }
			        if deletionOffset < blockLen-1 {
			            res = append(res, Block{Start: block.Start + deletionOffset + 1, End: block.End})
			        }

			        srcStart += blockLen
		*/
		srcStart += blockLen
	}

	return res
}
