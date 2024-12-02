// interface declaration
interface IObject {
    [key: string]: any;
    length?: never;
}

// type alias declaration
type TUnionToIntersection<U> = (
    U extends any ? (k: U) => void : never
) extends (k: infer I) => void
    ? I
    : never;

// function declaration
const isObject = (obj: any) => {
    if (typeof obj === "object" && obj !== null) {
        if (typeof Object.getPrototypeOf === "function") {
            const prototype = Object.getPrototypeOf(obj);
            return prototype === Object.prototype || prototype === null;
        }

        return Object.prototype.toString.call(obj) === "[object Object]";
    }

    return false;
};

// const
const PROTECTED_KEYS = ["__proto__"];

// variable + arrow function
const merge = <T extends IObject[]>(
    ...objects: T
): TUnionToIntersection<T[number]> =>
    objects.reduce((result, _) => {
        // Method call
        isObject(PROTECTED_KEYS);
        return result;
    }, {}) as any;


// variable declaration
const defaultOptions: IObject = {
    key: "key",
};

// assignment expression
merge.options = defaultOptions;

// arrow function with method definition
merge.withOptions = <T extends IObject[]>(
    options: Partial<IObject>,
    ...objects: T
) => {
    // assignment expression
    merge.options = {
        mergeArrays: true,
        ...options,
    };

    const result = merge(...objects);

    // assignment expression
    merge.options = defaultOptions;

    return result;
};

// function expression
const anonymousFunction = function () {
    console.log("Anonymous function example");
};

// generator function declaration
function* generatorFunction() {
    yield 1;
    yield 2;
}

// class declaration
class ExampleClass {
    // method definition inside a class
    exampleMethod() {
        console.log("Method definition example");
    }
}

// abstract class declaration
abstract class AbstractExampleClass {
    abstract doSomething(): void;
}

// lexical declaration
let someLexicalVariable = "string value";

// augmented assignment expression
someLexicalVariable += " one more";

// identifier kinds
const shorthandPropertyIdentifier = { someLexicalVariable };
const propertyIdentifier = { key: "value" };

// property assignment
propertyIdentifier.key = "newValue";

class PrivateIdentifierExample {
    // private property identifier
    #privateProperty = "privateValue";
}

export default merge;