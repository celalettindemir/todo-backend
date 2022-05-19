# Finally our multi-stage to build a small image
# Start a new stage from scratch
FROM scratch

COPY .config .config
# Copy the Pre-built binary file
COPY bin/main .
# Run executable
CMD ["./main"]