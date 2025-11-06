#!/bin/bash

set -e

echo "Generating flow-wasm.js..."

# Check if WASM file exists
if [ ! -f "web/flow-transpiler.wasm" ]; then
    echo "Error: web/flow-transpiler.wasm not found. Run 'make wasm' first."
    exit 1
fi

# Convert WASM to base64
echo "Converting WASM to base64..."
WASM_BASE64=$(base64 -i web/flow-transpiler.wasm)

# Generate the JavaScript file
cat > web/flow-wasm.js << EOF
// Auto-generated file containing the Flow transpiler WASM module
// Generated at: $(date)
// DO NOT EDIT - regenerate with 'make wasm-js'

const FLOW_WASM_BASE64 = \`${WASM_BASE64}\`;
EOF

# Also copy wasm_exec.js to wasm-runtime.js
if [ -f "web/wasm_exec.js" ]; then
    echo "// Go WASM runtime" > web/wasm-runtime.js
    echo "// From: https://github.com/golang/go/blob/master/misc/wasm/wasm_exec.js" >> web/wasm-runtime.js
    cat web/wasm_exec.js >> web/wasm-runtime.js
else
    echo "Warning: wasm_exec.js not found"
fi

echo "Generated web/flow-wasm.js ($(ls -lh web/flow-wasm.js | awk '{print $5}'))"
echo "Generated web/wasm-runtime.js"