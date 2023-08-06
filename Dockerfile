# Base image
FROM ghcr.io/swuecho/chat:latest

# Expose port 8080
EXPOSE 8080

# Set environment variables
ENV OPENAI_API_KEY=thisisopenaikey
ENV CLAUDE_API_KEY=thisisclaudekey
ENV OPENAI_RATELIMIT=100
ENV PG_HOST=db
ENV PG_DB=postgres
ENV PG_USER=postgres
ENV PG_PASS=thisisapassword
ENV PG_PORT=5432

# Install any additional dependencies if required

# Set up healthcheck for the database
HEALTHCHECK --interval=5s --timeout=5s --retries=5 \
  CMD pg_isready -q -d postgres -U postgres

# Start the application
CMD ["/app/start.sh"]
