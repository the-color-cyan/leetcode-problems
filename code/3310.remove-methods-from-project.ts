function remainingMethods(
    n: number,
    k: number,
    invocations: number[][],
): number[] {
    const methods = mapMethods(invocations);
    const suspicious = findSuspicious(k, methods);
}

type MethodId = number;

type Method = {
    invokes: Set<MethodId>;
};

function mapMethods(invocations: number[][]): Map<MethodId, Method> {
    let methods = new Map<MethodId, Method>();

    for (const [id, invocation] of invocations) {
        upsertInvocation(id, invocation, methods);
    }

    return methods;
}

function upsertInvocation(
    id: MethodId,
    invocation: MethodId,
    methods: Map<MethodId, Method>,
) {
    const method = methods.get(id);

    if (method !== undefined) {
        method.invokes.add(invocation);
    } else {
        methods.set(id, {
            invokes: new Set([invocation]),
        });
    }
}

function findSuspicious(
    buggedId: MethodId,
    methods: Map<MethodId, Method>,
): Set<MethodId> {
    const buggedMethod = methods.get(buggedId);

    let sus = buggedMethod.invokes;
    sus.add(buggedId);

    let allLooped = false;

    while (allLooped) {}
}
