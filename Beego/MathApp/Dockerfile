FROM golang:1.7

# Create the directory where the application will reside
RUN mkdir /app
#RUN echo "test"
#RUN ["/bin/bash", "-c", "echo hello"]


# Copy the application files (needed for production)
ADD MathApp /app/MathApp
ADD views /app/views
ADD conf /app/conf

# Set the working directory to the app directory
WORKDIR /app

#RUN chmod +x MathApp

# Expose the application on port 8080.
# This should be the same as in the app.conf file
EXPOSE 8080

# Set the entry point of the container to the application executable
ENTRYPOINT ["/app/MathApp"]

#RUN pwd

#CMD ["/app/MathApp"]

