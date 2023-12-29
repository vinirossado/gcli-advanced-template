# Use the official SQL Server 2019 image for Linux
FROM mcr.microsoft.com/mssql/server:2019-latest

# Set environment variables
ENV ACCEPT_EULA=Y \
    SA_PASSWORD=sua-maeEuOdeioReact2x \
    MSSQL_PID=Express \
    MSSQL_TCP_PORT=1433

# Expose SQL Server port
EXPOSE 1433

# Start SQL Server and run initialization script
CMD /opt/mssql/bin/sqlservr