# TODO - Ensure openapi tools are installed:
# npm install @openapitools/openapi-generator-cli -g
openapi-generator-cli generate \
    -i ../../cw-api.json \
    -g go -o ../../ \
    -c config.yaml \
    --git-user-id ai-connor \
    --git-repo-id connectwise