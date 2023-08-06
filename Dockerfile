# Use an official Python runtime as a parent image
FROM python:3.7-slim

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Install any needed packages specified in requirements.txt
RUN pip install --trusted-host pypi.python.org -r requirements.txt

# Make port 8080 available to the world outside this container
EXPOSE 8080

# Define environment variable
ENV OPENAI_API_KEY=thisisopenaikey 
ENV CLAUDE_API_KEY=thisisclaudekey 
ENV OPENAI_RATELIMIT=100 
ENV PG_HOST=db
ENV PG_DB=postgres 
ENV PG_USER=postgres
ENV PG_PASS=thisisapassword
ENV PG_PORT=5432

# Run app.py when the container launches
CMD ["python", "app.py"]
