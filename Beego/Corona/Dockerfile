FROM golang:1.7

# Create the directory where the application will reside
RUN rm -frv /app
RUN mkdir /app

# Copy the application files (needed for production)
ADD Corona /app/Corona
ADD views /app/views
ADD conf /app/conf
ADD static /app/static

# Set the working directory to the app directory
WORKDIR /app



# Expose the application on port 8080.
# This should be the same as in the app.conf file
EXPOSE 8080

# Set the entry point of the container to the application executable
ENTRYPOINT ["/app/Corona"]

#RUN pwd

#CMD ["/app/Corona"]

