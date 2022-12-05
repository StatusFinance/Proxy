# Proxy

Proxy is a Go program that acts as a proxy between the Status Finance API and your users' requests. It adds an additional layer of privacy by protecting the users' IP addresses and other data from being exposed to the API, and it also helps to protect the API key from being compromised.

## How it works

Proxy receives requests from clients and forwards them to the Status Finance API. It modifies the requests by removing certain headers and adding the API key in the `x-api-key` header. It then receives the response from the API, removes any sensitive data, and forwards it back to the client. This process effectively acts as a barrier between the Status Finance API and the clients, protecting both parties from potential privacy threats.

## Environment variables

The following environment variables must be set in order for the program to function properly:

- `PORT`: The port on which the proxy server will listen for incoming requests. If not set, the default value is `3000`.
- `API_BASE_URL`: The base URL of the API. If not set, the default value is `https://api.status.finance`.
- `API_KEY`: The API key that will be used in the `x-api-key` header of the requests. This key is required for authentication with the API.

## Dependencies

Proxy has the following dependency:

- [`github.com/joho/godotenv`](https://github.com/joho/godotenv): A package for loading environment variables from a `.env` file. This dependency is used to easily manage the environment variables required by the program.

## License

This project is licensed under the GNU General Public License v3.0. See the [LICENSE](LICENSE) file for details.
