PARAM_TYPES = {
    "string": {
        True: "string",
        False: "*optional.String",
    },
    "number": {
        True: "int",
        False: "*optional.Int",
    },
    "boolean": {
        True: "bool",
        False: "*optional.Bool",
    },
    "numbers": {
        True: "[]int",
        False: "*[]int",
    },
}
