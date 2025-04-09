<div align="center">

# LaECa

#### A Simple and Lightweight Load Balancer in Go
  
<img src="https://res.cloudinary.com/dtr5bmnd0/image/upload/f_auto,q_auto/v1/LaECa/LaECa" alt="LaECa" width="380px" height="380px"/>
  
</div>

## Description
`laeca` is a lightweight and efficient load balancer developed in Go, leveraging only Go's native standard library (`stdlib`). It is designed to be simple to configure and deploy, making it an excellent choice for lightweight applications.

## Supported Features

### Protocols
- **Currently Supported:**
  - `http`: Load balancing for HTTP traffic.
  - `tcp`: Load balancing for raw TCP connections.

### Load Balancing Algorithms
- **Currently Supported:**
  - **Round-Robin**: Requests are distributed sequentially to all upstream servers.
  
- **Future Enhancements:**
  - **Least Connections**: Directs traffic to the server with the fewest active connections.
  - **Weighted Round-Robin**: Servers are assigned weights, and traffic distribution considers these weights.
  - **IP Hashing**: Routes requests based on the client's IP address to ensure session affinity.

## Configuration and Installation

1. **Configuration File:**
   Create a `laeca.yaml` file in the project directory with the desired settings. Below is an example configuration file:

   ```yaml
   listen: 9999
   protocol: "http"
   algorithm: "round-robin"
   upstream:
       - url: "0.0.0.0"
         port: 9998
       - url: "0.0.0.0"
         port: 9997
       - url: "0.0.0.0"
         port: 9996
       - url: "0.0.0.0"
         port: 9995
   ```

   - `listen`: The port on which the load balancer will listen for incoming traffic.
   - `protocol`: Protocol used for balancing traffic. Supported values: `http`, `tcp`.
   - `algorithm`: The load balancing algorithm to use. Supported value: `round-robin`.
   - `upstream`: A list of backend servers to which requests will be forwarded, including their `url` and `port`.

2. **Starting the Application:**
   Once the configuration is ready, run the application with the following command:

   ```bash
   laeca start
   ```

   Ensure the `laeca.yaml` file is located in the same directory where the command is executed.

## License
This project is licensed under a standard open-source license. Feel free to use, modify, and contribute to the project.
