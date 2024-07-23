# Use the official Node.js image as the base image
FROM node:18

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install application dependencies
RUN npm install

# Copy the application entry file
COPY index.js ./

# Expose the port the app runs on
EXPOSE 5004

# Command to run the application
CMD ["npm", "run", "start"]


