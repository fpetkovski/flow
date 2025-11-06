// PromQL+ mode for CodeMirror
(function(mod) {
    if (typeof exports == "object" && typeof module == "object") // CommonJS
        mod(require("codemirror"));
    else if (typeof define == "function" && define.amd) // AMD
        define(["codemirror"], mod);
    else // Plain browser env
        mod(CodeMirror);
})(function(CodeMirror) {
    "use strict";

    CodeMirror.defineMode("promqlplus", function() {
        // Keywords
        const keywords = ["let", "in", "by", "without"];

        // Aggregation operators
        const aggregationOps = [
            "sum", "min", "max", "avg", "stddev", "stdvar",
            "count", "count_values", "bottomk", "topk", "quantile"
        ];

        // Common function names (extensible)
        const functions = [
            "rate", "irate", "increase", "delta", "idelta", "deriv",
            "predict_linear", "resets", "changes", "avg_over_time",
            "min_over_time", "max_over_time", "sum_over_time",
            "count_over_time", "quantile_over_time", "stddev_over_time",
            "stdvar_over_time", "last_over_time", "present_over_time"
        ];

        // Match operators
        const matchOps = ["=", "!=", "=~", "!~"];

        // Duration units
        const durationUnits = ["ms", "s", "m", "h", "d", "w", "y"];

        function tokenBase(stream, state) {
            // Handle whitespace
            if (stream.eatSpace()) return null;

            // Handle comments
            if (stream.match(/^#.*/)) {
                return "comment";
            }

            // Handle strings
            if (stream.match(/^"([^"\\]|\\.)*"/)) {
                return "string";
            }
            if (stream.match(/^'([^'\\]|\\.)*'/)) {
                return "string";
            }
            if (stream.match(/^`([^`\\]|\\.)*`/)) {
                return "string";
            }

            // Handle numbers with duration units
            if (stream.match(/^\d+(\.\d+)?([eE][+-]?\d+)?/)) {
                // Check if followed by duration unit
                const pos = stream.pos;
                if (stream.match(/^(ms|s|m|h|d|w|y)\b/)) {
                    return "number";
                }
                stream.pos = pos;
                return "number";
            }

            // Handle operators
            if (stream.match(/^(=~|!~|!=|<=|>=|==)/)) {
                return "operator";
            }
            if (stream.match(/^[=<>!]/)) {
                return "operator";
            }
            if (stream.match(/^[+\-*/|(){},%]/)) {
                return "operator";
            }

            // Handle identifiers and keywords
            if (stream.match(/^[a-zA-Z_:][a-zA-Z0-9_:]*/)) {
                const word = stream.current();

                // Check for keywords
                if (keywords.includes(word)) {
                    return "keyword";
                }

                // Check for aggregation operators
                if (aggregationOps.includes(word)) {
                    return "builtin";
                }

                // Check for functions
                if (functions.includes(word)) {
                    return "variable-2";
                }

                // Check for boolean values
                if (word === "true" || word === "false") {
                    return "atom";
                }

                // Default to identifier
                return "variable";
            }

            // Consume any remaining character
            stream.next();
            return null;
        }

        return {
            startState: function() {
                return {
                    inComment: false
                };
            },
            token: function(stream, state) {
                return tokenBase(stream, state);
            },
            lineComment: "#"
        };
    });

    // Define MIME type
    CodeMirror.defineMIME("text/x-promqlplus", "promqlplus");
    CodeMirror.defineMIME("application/x-promqlplus", "promqlplus");
});