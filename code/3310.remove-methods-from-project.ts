function remainingMethods(
    n: number,
    k: number,
    invocations: number[][],
): number[] {
    let methods = mapMethods(invocations);
}

type MethodId = number;

type Method = {
    invokes: Set<MethodId>;
};

function mapMethods(invocations: number[][]): Map<MethodId, Method> {
    let methods = new Map<MethodId, Method>();

    for (const [id, invocation] of invocations) {
        methods.
    }
}

function findSuspicious(k: number, invocations: number[][]): number[] {}a

function upsert(key: number, map: Map<number, Method>) {

}
