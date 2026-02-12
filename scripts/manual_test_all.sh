#!/bin/bash
set -e

# Run tests
run_test inl_corp_1 650alldb2 300 300
run_test inl_corp_1 dynamicntdw1 1200 627

run_test inl_corp_2 650alldb2 300 300
run_test inl_corp_2 dynamicntdw1 1200 627

# Keeta doesn't require proxy
./scripts/manual_test.sh keeta "" 300 300

echo ""
echo "--------------------------------"
echo "🎉 All tests passed successfully!"
