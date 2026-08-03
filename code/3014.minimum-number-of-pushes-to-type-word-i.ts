function minimumPushes(word: string): number {
    const l = word.length;
    let result = 0;

    for (let i = 0; i < l; i++) {
        result += Math.floor(i / 8) + 1;
    }

    return result;
}
