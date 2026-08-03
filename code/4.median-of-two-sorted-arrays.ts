function findMedianSortedArrays(numsA: number[], numsB: number[]): number {
    const smaller = numsA.length <= numsB.length ? numsA : numsB;
    const larger = numsA.length <= numsB.length ? numsB : numsA;

    const smallerSize = smaller.length;
    const largerSize = larger.length;

    const combinedMidpoint = Math.floor((smallerSize + largerSize + 1) / 2);
    let lowBound = Math.max(0, combinedMidpoint - largerSize);
    let highBound = Math.min(smallerSize, combinedMidpoint);

    while (lowBound <= highBound) {
        const smallerPartition = Math.floor((lowBound + highBound) / 2);
        const largerPartition = combinedMidpoint - smallerPartition;

        const smallerLeft =
            smallerPartition === 0 ? -Infinity : smaller[smallerPartition - 1];
        const smallerRight =
            smallerPartition === smallerSize
                ? Infinity
                : smaller[smallerPartition];
        const largerLeft =
            largerPartition === 0 ? -Infinity : larger[largerPartition - 1];
        const largerRight =
            largerPartition === largerSize ? Infinity : larger[largerPartition];

        if (smallerLeft > largerRight) {
            highBound = smallerPartition - 1;
        } else if (largerLeft > smallerRight) {
            lowBound = smallerPartition + 1;
        } else {
            const leftMax = Math.max(smallerLeft, largerLeft);

            if ((smallerSize + largerSize) % 2 === 1) {
                return leftMax;
            }

            const rightMin = Math.min(smallerRight, largerRight);

            return (leftMax + rightMin) / 2;
        }
    }

    return 0.0;
}
