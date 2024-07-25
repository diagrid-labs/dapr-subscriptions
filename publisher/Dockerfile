# Use Maven image to build the jar
FROM maven:3.8.6-openjdk-11 as build

# Set the working directory inside the container
WORKDIR /app

# Copy the pom.xml and src directory to the container
COPY pom.xml .
COPY src ./src

# Run the Maven build command to generate the JAR file
RUN mvn clean package

# Use OpenJDK image for runtime
FROM openjdk:11

# Expose the application port
EXPOSE 5001

# Set the environment variable
ENV PORT 5001

# Create a volume for temporary files
VOLUME /tmp

# Copy the JAR file from the Maven build stage
COPY --from=build /app/target/*.jar app.jar

# Set the entry point to run the JAR file
ENTRYPOINT ["java","-jar","/app.jar"]
