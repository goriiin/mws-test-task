#!/bin/bash
set -e

COVERAGE=$(go tool cover -func=coverage.out | grep "total:" | awk '{print $3}' | sed 's/%//')
COVERAGE_INT=$(printf "%.0f\n" $COVERAGE)
echo "==> Total coverage: ${COVERAGE}%"


COLOR="red"
if (( $(echo "$COVERAGE_INT > 80" | bc -l) )); then
    COLOR="success"
elif (( $(echo "$COVERAGE_INT > 60" | bc -l) )); then
    COLOR="yellow"
fi
echo "==> Badge color: ${COLOR}"


BADGE_URL="https://img.shields.io/badge/coverage-${COVERAGE_INT}%25-${COLOR}"
echo "==> New badge URL: ${BADGE_URL}"


sed -i "s|!\[Coverage Status\].*|![Coverage Status](${BADGE_URL})|" README.md
echo "==> README.md updated."


git config --global user.name 'github-actions[bot]'
git config --global user.email 'github-actions[bot]@users.noreply.github.com'

if [[ `git status --porcelain` ]]; then
    echo "==> Committing and pushing changes..."
    git add README.md
    git commit -m "ci: обновить бейдж покрытия до ${COVERAGE}%"
    git push
else
    echo "==> No changes to commit."
fi

echo "==> Badge update complete."