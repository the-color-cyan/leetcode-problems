function remainingMethods(
    n: number,
    k: number,
    invocations: number[][],
): number[] {
    const methods: Method[] = mapMethods(n, invocations);
    const sus: Set<MethodId> = findSus(k, methods);
    const nonSus: Set<MethodId> = new Set(
        [...methods.keys()].filter((id) => !sus.has(id)),
    );

    for (const methodId of nonSus) {
        let method = methods[methodId];

        for (const invokedId of method.invokes) {
            if (sus.has(invokedId)) {
                return [...methods.keys()];
            }
        }
    }

    return [...nonSus];
}

type MethodId = number;

type Method = {
    invokes: Set<MethodId>;
};

function mapMethods(n: number, invocations: number[][]): Method[] {
    const methods: Method[] = Array.from({ length: n }, () => {
        return {
            invokes: new Set<MethodId>(),
        };
    });

    for (const [id, invocation] of invocations) {
        methods[id].invokes.add(invocation);
    }

    return methods;
}

function findSus(buggedId: MethodId, methods: Method[]): Set<MethodId> {
    let sus = new Set<MethodId>([buggedId]);

    for (const susId of sus) {
        const nextSus = methods[susId].invokes;

        for (const nextSusId of nextSus) {
            if (!sus.has(nextSusId)) sus.add(nextSusId);
        }
    }

    return sus;
}
