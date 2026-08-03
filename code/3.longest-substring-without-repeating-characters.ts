function lengthOfLongestSubstring(s: string): number {
    // index of last occurance of a character
    let lastOccurance = new Map<string, number>();
    let longest = 0;
    let start = 0;

    for (let end = 0; end < s.length; end++) {
        if (lastOccurance.has(s[end])) {
            const currLastOccurance = lastOccurance.get(s[end]);
            if (currLastOccurance + 1 > start) start = currLastOccurance + 1;
        }

        lastOccurance.set(s[end], end);

        const currSubstringLength = end + 1 - start;
        if (currSubstringLength > longest) longest = currSubstringLength;
    }

    return longest;
}
