#!/bin/bash
set -e

# Run tests
./scripts/manual_test.sh inl_corp_1 650alldb2 300 300
./scripts/manual_test.sh inl_corp_1 dynamicntdw1 1200 627

./scripts/manual_test.sh inl_corp_2 650alldb2 300 300
./scripts/manual_test.sh inl_corp_2 dynamicntdw1 1200 627

./scripts/manual_test.sh keeta "" 300 300

echo ""
echo "--------------------------------"
echo "🎉 All tests passed successfully!"
