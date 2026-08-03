function twoSum(nums: number[], target: number): number[] {
    const numMap: Map<number, number> = new Map();

    for (const [i, n] of nums.entries()) {
        const diff = target - n;
        const diffIndex = numMap.get(diff);

        if (diffIndex != undefined) return [diffIndex, i];

        numMap.set(n, i);
    }

    return [];
}
