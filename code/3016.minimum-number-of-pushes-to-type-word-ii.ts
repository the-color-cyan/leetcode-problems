const BUTTONS = 8;

function minimumPushes(word: string): number {
    const lettersByFreq = groupLettersByFreq(word);
    const freqs = Array.from(
        lettersByFreq,
        ([freq, letters]) => [freq, letters.size] as const,
    ).sort(([a], [b]) => b - a);

    let pushTotal = 0;
    let assignedLetters = 0;

    for (const [freq, lettersInFreq] of freqs) {
        for (let remaining = lettersInFreq; remaining > 0; remaining--) {
            const pushesPerLetter = Math.floor(assignedLetters / BUTTONS) + 1;

            pushTotal += freq * pushesPerLetter;
            assignedLetters++;
        }
    }

    return pushTotal;
}

function groupLettersByFreq(word: string): Map<number, Set<string>> {
    const letterFreqs = countLetterFreqs(word);
    let lettersByFreq = new Map<number, Set<string>>();

    for (const [letter, freq] of letterFreqs) {
        const letters = lettersByFreq.get(freq);

        if (letters) letters.add(letter);
        else lettersByFreq.set(freq, new Set<string>([letter]));
    }

    return lettersByFreq;
}

function countLetterFreqs(word: string): Map<string, number> {
    let letterFreqs = new Map<string, number>();

    for (const letter of word) {
        letterFreqs.set(letter, (letterFreqs.get(letter) ?? 0) + 1);
    }

    return letterFreqs;
}
