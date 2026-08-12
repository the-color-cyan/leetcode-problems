function lengthOfLongestSubstring(s: string): number {
    // index of last occurrence of a character
    let lastOccurrence = new Map<string, number>();
    let longest = 0;
    let start = 0;

    for (let end = 0; end < s.length; end++) {
        if (lastOccurrence.has(s[end])) {
            const currLastOccurrence = lastOccurrence.get(s[end]);
            if (currLastOccurrence + 1 > start) start = currLastOccurrence + 1;
        }

        lastOccurrence.set(s[end], end);

        const currSubstringLength = end + 1 - start;
        if (currSubstringLength > longest) longest = currSubstringLength;
    }

    return longest;
}
